#![warn(clippy::all, clippy::pedantic)]

extern crate redis;

use std::collections::HashMap;
use std::string::ToString;
use redis::{Commands, RedisError, RedisResult};
use redis::geo::{RadiusOptions, RadiusOrder, RadiusSearchResult, Unit};
use warp::{Filter, Rejection, Reply};
use serde::{Deserialize, Serialize};
use warp::filters::body::BodyDeserializeError;
use warp::http::{StatusCode};

#[derive(Debug)]
enum BadRequestError {
    ParameterNotNumeric(String),
    ParameterRequired(String),
}

impl warp::reject::Reject for BadRequestError {}

impl std::fmt::Display for BadRequestError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match *self {
            BadRequestError::ParameterNotNumeric(ref param) => write!(f, "{}", format!("\"{}\" parameter must be numeric value", param)),
            BadRequestError::ParameterRequired(ref param) => write!(f, "{}", format!("\"{}\" parameter required", param)),
        }
    }
}

#[derive(Debug)]
struct RedisErrorType(RedisError);

impl warp::reject::Reject for RedisErrorType {}

async fn handle_rejection(err: Rejection) -> Result<impl Reply, std::convert::Infallible> {
    #[derive(Serialize)]
    struct Response {
        error: String,
    }

    let error_message;
    let status_code;
    if err.is_not_found() {
        error_message = "404 page not found".to_string();
        status_code = StatusCode::NOT_FOUND;
    } else if let Some(err) = err.find::<BadRequestError>() {
        error_message = err.to_string();
        status_code = StatusCode::BAD_REQUEST;
    } else if let Some(err) = err.find::<RedisErrorType>() {
        let code = err.0.code().unwrap_or("");
        let category = err.0.category();
        let message = err.0.to_string();
        let status = format!("REDIS_ERROR: {} (code: {}, category: {})", message, code, category);
        error_message = status;
        status_code = StatusCode::INTERNAL_SERVER_ERROR;
    } else if let Some(err) = err.find::<BodyDeserializeError>() {
        error_message = err.to_string();
        status_code = StatusCode::UNPROCESSABLE_ENTITY;
    } else if let Some(_) = err.find::<warp::reject::MethodNotAllowed>() {
        // Silly hack for converting 405 -> 404
        // https://github.com/seanmonstar/warp/issues/77
        error_message = "404 page not found".to_string();
        status_code = StatusCode::NOT_FOUND;
    } else {
        eprintln!("unhandled rejection: {:?}", err);
        error_message = "INTERNAL_SERVER_ERROR".to_string();
        status_code = StatusCode::INTERNAL_SERVER_ERROR;
    }
    Ok(warp::reply::with_status(
        warp::reply::json(&Response { error: error_message }),
        status_code
    ))
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

async fn load_parcel_lockers_pagenated(params: HashMap<String, String>) -> Result<impl Reply, Rejection> {
    let page = match params.get("page") {
        Some(page_param) => {
            match page_param.parse::<isize>() {
                Ok(page) => page,
                Err(_) => return Err(warp::reject::custom(
                    BadRequestError::ParameterNotNumeric("page".to_string())
                ))
            }
        }
        None => 1,
    };
    let per_page = match params.get("per_page") {
        Some(per_page_param) => {
            match per_page_param.parse::<isize>() {
                Ok(per_page) => per_page,
                Err(_) => return Err(warp::reject::custom(
                    BadRequestError::ParameterNotNumeric("per_page".to_string())
                ))
            }
        }
        None => 10,
    };

    let start = (page - 1) * per_page;
    let stop = start + per_page - 1;

    let mut con = connect();
    let result: RedisResult<Vec<String>> = con.zrange("parcel_lockers", start, stop);
    match result {
        Ok(parcel_locker_ids) => {
            let mut parcel_lockers: Vec<ParcelLocker> = Vec::new();
            for id in parcel_locker_ids {
                let key = make_parcel_locker_key(&id);
                match con.hgetall::<String, HashMap<String, String>>(key) {
                    Ok(pl_hm) => parcel_lockers.push(pl_hm.into()),
                    Err(err) => return Err(warp::reject::custom(RedisErrorType(err))),
                }
            }
            Ok(warp::reply::json(&parcel_lockers))
        }
        Err(err) => Err(warp::reject::custom(RedisErrorType(err)))
    }
}

async fn find_parcel_lockers_by_distance(params: HashMap<String, String>) -> Result<impl Reply, Rejection> {
    let longitude = match params.get("longitude") {
        Some(longitude) => {
            match longitude.parse::<f64>() {
                Ok(longitude) => longitude,
                Err(_) => return Err(warp::reject::custom(
                    BadRequestError::ParameterNotNumeric("longitude".to_string())
                ))
            }
        },
        None => return Err(warp::reject::custom(
            BadRequestError::ParameterRequired("longitude".to_string())
        ))
    };
    let latitude = match params.get("latitude") {
        Some(latitude) => {
            match latitude.parse::<f64>() {
                Ok(latitude) => latitude,
                Err(_) => return Err(warp::reject::custom(
                    BadRequestError::ParameterNotNumeric("latitude".to_string())
                ))
            }
        },
        None => return Err(warp::reject::custom(
            BadRequestError::ParameterRequired("latitude".to_string())
        ))
    };
    let radius = match params.get("radius") {
        Some(radius) => {
            match radius.parse::<f64>() {
                Ok(radius) => radius,
                Err(_) => return Err(warp::reject::custom(
                    BadRequestError::ParameterNotNumeric("radius".to_string())
                ))
            }
        },
        None => return Err(warp::reject::custom(
            BadRequestError::ParameterRequired("radius".to_string())
        ))
    };

    #[derive(Serialize)]
    pub struct RadiusSearchResultSerializable {
        name: String,
        latitude: f64,
        longitude: f64,
        distance: Option<f64>,
    }

    impl From<&RadiusSearchResult> for RadiusSearchResultSerializable {
        fn from(rsr: &RadiusSearchResult) -> Self {
            RadiusSearchResultSerializable {
                name: rsr.name.clone(),
                latitude: rsr.coord.as_ref().unwrap().latitude,
                longitude: rsr.coord.as_ref().unwrap().longitude,
                distance: rsr.dist.clone(),
            }
        }
    }

    let mut con = connect();
    let opts = RadiusOptions::default()
        .with_dist()
        .with_coord()
        .order(RadiusOrder::Asc);
    let result: RedisResult<Vec<RadiusSearchResult>> = con
        .geo_radius("parcel_lockers", longitude, latitude, radius, Unit::Kilometers, opts);
    match result {
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

async fn delete_all_parcel_lockers() -> Result<impl Reply, Rejection> {
    #[derive(Serialize)]
    struct Response {
        message: String,
    }
    let mut con = connect();
    match con.zrange::<&str, Vec<String>>("parcel_lockers", 0, -1) {
        Ok(pl_ids) => {
            for id in pl_ids {
                if let Err(err) = con.del::<String, ()>(make_parcel_locker_key(&id)) {
                    return Err(warp::reject::custom(RedisErrorType(err)))
                }
            }
            Ok(warp::reply::json(&Response { message: "Parcel lockers deleted".to_string() }))
        }
        Err(err) => Err(warp::reject::custom(RedisErrorType(err)))
    }
}

fn make_parcel_locker_key(id: &str) -> String {
    format!("parcel_locker:{}", id)
}

#[tokio::main]
async fn main() {
    let ping_route = warp::path("ping")
        .and_then(ping_handler);

    let get_parcel_lockers_route = warp::get()
        .and(warp::path("parcel-lockers"))
        .and(warp::path::end())
        .and(warp::query())
        .and_then(load_parcel_lockers_pagenated);

    let get_parcel_locker_by_id_route = warp::get()
        .and(warp::path("parcel-locker"))
        .and(warp::path::param())
        .and(warp::path::end())
        .and_then(load_parcel_locker_by_id);

    let get_parcel_lockers_by_distance_route = warp::get()
        .and(warp::path("parcel-locker-distance-search"))
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

    let delete_all_parcel_lockers_route = warp::delete()
        .and(warp::path("all-parcel-lockers"))
        .and(warp::path::end())
        .and_then(delete_all_parcel_lockers);

    let routes = ping_route
        .or(get_parcel_lockers_route)
        .or(get_parcel_locker_by_id_route)
        .or(post_parcel_locker_route)
        .or(delete_parcel_locker_by_id_route)
        .or(get_parcel_lockers_by_distance_route)
        .or(delete_all_parcel_lockers_route)
        .recover(handle_rejection);

    warp::serve(routes)
        .run(([127, 0, 0, 1], 8081))
        .await;
}
