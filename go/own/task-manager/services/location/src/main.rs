#![warn(clippy::all, clippy::pedantic)]

extern crate redis;

use std::collections::HashMap;
use redis::{Commands, RedisError, RedisResult};
use redis::geo::{RadiusOptions, RadiusOrder, RadiusSearchResult, Unit};
use warp::{Filter, Rejection, Reply};
use serde::{Deserialize, Serialize};
use warp::http::{StatusCode};

#[derive(Debug)]
struct InvalidParameter {
    message: String,
}

impl warp::reject::Reject for InvalidParameter {}

#[derive(Debug)]
struct RedisErrorType(RedisError);

impl warp::reject::Reject for RedisErrorType {}

#[derive(Debug)]
struct ParameterRequired {
    message: String,
}

impl warp::reject::Reject for ParameterRequired {}

async fn handle_rejection(err: Rejection) -> Result<impl Reply, std::convert::Infallible> {
    if err.is_not_found() {
        Ok(warp::reply::with_status("404 page not found".to_string(), StatusCode::NOT_FOUND))
    } else if let Some(err) = err.find::<InvalidParameter>() {
        Ok(warp::reply::with_status(err.message.clone(), StatusCode::BAD_REQUEST))
    } else if let Some(err) = err.find::<ParameterRequired>() {
        Ok(warp::reply::with_status(err.message.clone(), StatusCode::BAD_REQUEST))
    } else if let Some(err) = err.find::<RedisErrorType>() {
        let code = err.0.code().unwrap_or("");
        let category = err.0.category();
        let message = err.0.to_string();
        let status = format!("REDIS_ERROR: code: {}, category: {}, message: {}", code, category, message);
        Ok(warp::reply::with_status(status, StatusCode::INTERNAL_SERVER_ERROR))
    } else if let Some(_) = err.find::<warp::reject::MethodNotAllowed>() {
        // Silly hack for converting 405 -> 404
        // https://github.com/seanmonstar/warp/issues/77
        Ok(warp::reply::with_status("404 page not found".to_string(), StatusCode::NOT_FOUND))
    } else {
        eprintln!("unhandled rejection: {:?}", err);
        Ok(warp::reply::with_status("INTERNAL_SERVER_ERROR".to_string(), StatusCode::INTERNAL_SERVER_ERROR))
    }
}

async fn ping_handler() -> Result<impl Reply, Rejection> {
    #[derive(Serialize)]
    struct PingResponse {
        message: &'static str,
    }
    Ok(warp::reply::json(&PingResponse { message: "pong" }))
}

#[derive(Clone, Deserialize, Debug, Serialize)]
struct ParcelLocker {
    id: String,
    name: String,
    longitude: f64,
    latitude: f64,
}

impl ParcelLocker {
    fn to_tuples(&self) -> Vec<(&str, String)> {
        vec![
            ("id", self.id.clone()),
            ("name", self.name.clone()),
            ("longitude", self.longitude.to_string()),
            ("latitude", self.latitude.to_string()),
        ]
    }
}

impl From<HashMap<String, String>> for ParcelLocker {
    fn from(value: HashMap<String, String>) -> Self {
        ParcelLocker {
            id: value["id"].clone(),
            name: value["name"].clone(),
            longitude: value["longitude"].parse::<f64>().unwrap(),
            latitude: value["latitude"].parse::<f64>().unwrap(),
        }
    }
}

//async fn connect() -> redis::RedisResult<redis::Connection> {
fn connect() -> redis::Connection {
    let client = redis::Client::open("redis://127.0.0.1:6379").unwrap();
    let con = client.get_connection().unwrap();
    con
}

async fn load_parcel_locker_by_id(id: String) -> Result<impl Reply, Rejection> {
    let mut con = connect();
    let key = make_parcel_locker_key(&id);
    match con.hgetall::<String, HashMap<String, String>>(key) {
        Ok(pl_hm) => Ok(warp::reply::json::<ParcelLocker>(&pl_hm.into())),
        Err(err) => Err(warp::reject::custom(RedisErrorType(err)))
    }
}

async fn find_parcel_lockers_by_distance(params: HashMap<String, String>) -> Result<impl Reply, Rejection> {
    let longitude = match params.get("longitude") {
        Some(longitude) => {
            match longitude.parse::<f64>() {
                Ok(longitude) => longitude,
                Err(_) => return Err(warp::reject::custom(
                    InvalidParameter{message: "\"longitude\" parameter must be numeric value".to_string()}
                ))
            }
        },
        None => return Err(warp::reject::custom(
            ParameterRequired{message: "\"longitude\" parameter required".to_string()}
        ))
    };
    let latitude = match params.get("latitude") {
        Some(latitude) => {
            match latitude.parse::<f64>() {
                Ok(latitude) => latitude,
                Err(_) => return Err(warp::reject::custom(
                    InvalidParameter{message: "\"latitude\" parameter must be numeric value".to_string()}
                ))
            }
        },
        None => return Err(warp::reject::custom(
            ParameterRequired{message: "\"latitude\" parameter required".to_string()}
        ))
    };
    let radius = match params.get("radius") {
        Some(radius) => {
            match radius.parse::<f64>() {
                Ok(radius) => radius,
                Err(_) => return Err(warp::reject::custom(
                    InvalidParameter{message: "\"radius\" parameter must be numeric value".to_string()}
                ))
            }
        },
        None => return Err(warp::reject::custom(
            ParameterRequired{message: "\"radius\" parameter required".to_string()}
        ))
    };

    #[derive(Serialize)]
    pub struct CoordSerializable {
        latitude: f64,
        longitude: f64,
    }

    #[derive(Serialize)]
    pub struct RadiusSearchResultSerializable {
        name: String,
        // coord: Option<CoordSerializable>,
        latitude: f64,
        longitude: f64,
        distance: Option<f64>,
    }

    impl From<&RadiusSearchResult> for RadiusSearchResultSerializable {
        fn from(rsr: &RadiusSearchResult) -> Self {
            let mut rsrs = RadiusSearchResultSerializable {
                name: rsr.name.clone(),
                // coord: rsr.coord.as_ref().map(|c| CoordSerializable {
                //     latitude: c.latitude,
                //     longitude: c.longitude,
                // }),
                latitude: 0f64,
                longitude: 0f64,
                distance: rsr.dist.clone(),
            };
            // TODO: Разобраться, почему пустые координаты
            if let Some(coord) = rsr.coord.as_ref() {
                rsrs.latitude = coord.latitude;
                rsrs.longitude = coord.longitude;
            }
            rsrs
        }
    }

    let mut con = connect();
    let opts = RadiusOptions::default().with_dist().order(RadiusOrder::Asc);
    let result: RedisResult<Vec<RadiusSearchResult>> = con.geo_radius("parcel_lockers", longitude, latitude, radius, Unit::Kilometers, opts);
    return match result {
        Ok(search_results) => {
            let serializable_results: Vec<RadiusSearchResultSerializable> = search_results.iter()
                .map(|result| RadiusSearchResultSerializable::from(result))
                .collect();
            Ok(warp::reply::json(&serializable_results))
        }
        Err(err) => Err(warp::reject::custom(RedisErrorType(err)))
    }
}

async fn save_parcel_locker(parcel_locker: ParcelLocker) -> Result<impl Reply, Rejection> {
    let mut con = connect();
    let parcel_locker2 = parcel_locker.clone();
    let parcel_locker3 = parcel_locker.clone();
    let key = make_parcel_locker_key(&parcel_locker.id);

    let is_new = match con.exists::<String, bool>(key.clone()) {
        Ok(exists) => !exists,
        Err(err) => return Err(warp::reject::custom(RedisErrorType(err)))
    };

    let parcel_locker_tuples = parcel_locker.to_tuples();
    if let Err(err) = con.hset_multiple::<String, &str, String, ()>(key, &parcel_locker_tuples) {
        return Err(warp::reject::custom(RedisErrorType(err)))
    }

    if let Err(err) = con.geo_add::<&str, (f64, f64, String), ()>("parcel_lockers", (parcel_locker3.longitude, parcel_locker3.latitude, parcel_locker3.id)) {
        return Err(warp::reject::custom(RedisErrorType(err)))
    }

    Ok(warp::reply::with_status(
        warp::reply::json::<ParcelLocker>(&parcel_locker2),
        if is_new { StatusCode::CREATED } else { StatusCode::OK }
    ))
}

async fn delete_parcel_locker_by_id(id: String) -> Result<impl Reply, Rejection> {
    #[derive(Serialize)]
    struct Response {
        deleted: bool,
        message: String,
        parcel_locker: Option<ParcelLocker>,
    }

    let mut con = connect();
    let key = make_parcel_locker_key(&id);
    match con.hgetall::<String, HashMap<String, String>>(key.clone()) {
        Ok(pl_hm) if pl_hm.is_empty() => Ok(warp::reply::json(
            &Response {
                deleted: false,
                message: format!("Parcel locker not found by id: {}", id),
                parcel_locker: None
            })),
        Err(err) => Err(warp::reject::custom(RedisErrorType(err))),
        Ok(pl_hm) => {
            match con.del::<&str, ()>(&key) {
                Ok(_) => {
                    match con.zrem::<&str, &str, ()>("parcel_lockers", &id) {
                        Ok(_) => Ok(warp::reply::json(
                            &Response {
                                deleted: true,
                                message: "Parcel locker deleted".to_string(),
                                parcel_locker: Some(pl_hm.into()),
                            }
                        )),
                        Err(err) => Err(warp::reject::custom(RedisErrorType(err)))
                    }
                },
                Err(err) => Err(warp::reject::custom(RedisErrorType(err)))
            }
        }
    }
}

fn make_parcel_locker_key(id: &str) -> String {
    format!("parcel_locker:{}", id)
}

#[tokio::main]
async fn main() {
    let ping_route = warp::path("ping")
        .and_then(ping_handler);

    let get_parcel_locker_by_id_route = warp::get()
        .and(warp::path("parcel-locker"))
        .and(warp::path::param())
        .and(warp::path::end())
        .and_then(load_parcel_locker_by_id);

    let get_parcel_lockers_by_distance_route = warp::get()
        .and(warp::path("parcel-locker"))
        .and(warp::path("distance-search"))
        .and(warp::path::end())
        .and(warp::query())
        .and_then(find_parcel_lockers_by_distance);

    let post_parcel_locker_route = warp::post()
        .and(warp::path("parcel-locker"))
        .and(warp::path::end())
        .and(warp::body::json())
        .and_then(save_parcel_locker);

    let delete_parcel_locker_by_id_route = warp::delete()
        .and(warp::path("parcel-locker"))
        .and(warp::path::param())
        .and(warp::path::end())
        .and_then(delete_parcel_locker_by_id);

    let routes = ping_route
        .or(get_parcel_lockers_by_distance_route)
        .or(get_parcel_locker_by_id_route)
        .or(post_parcel_locker_route)
        .or(delete_parcel_locker_by_id_route)
        .recover(handle_rejection);

    warp::serve(routes)
        .run(([127, 0, 0, 1], 8081))
        .await;
}
