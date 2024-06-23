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
        Ok(warp::reply::with_status("NOT_FOUND".to_string(), StatusCode::NOT_FOUND))
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

#[derive(Clone, Deserialize, Debug)]
struct Location {
    id: String,
    name: String,
    longitude: f64,
    latitude: f64,
}

impl Location {
    fn to_tuples(&self) -> Vec<(&str, String)> {
        vec![
            ("id", self.id.clone()),
            ("name", self.name.clone()),
            ("longitude", self.longitude.to_string()),
            ("latitude", self.latitude.to_string()),
        ]
    }
}

//async fn connect() -> redis::RedisResult<redis::Connection> {
fn connect() -> redis::Connection {
    let client = redis::Client::open("redis://127.0.0.1:6379").unwrap();
    let con = client.get_connection().unwrap();
    con
}

async fn load_location_by_id(id: String) -> Result<impl Reply, Rejection> {
    let mut con = connect();
    let key = format!("location:{}", id);
    match con.hgetall(key) {
        Ok(l) => Ok(warp::reply::json::<HashMap<String, String>>(&l)),
        Err(err) => Err(warp::reject::custom(RedisErrorType(err)))
    }
}

async fn save_location(location: Location) -> Result<impl Reply, Rejection> {
    let mut con = connect();
    let location2 = location.clone();
    let location3 = location.clone();
    let key = format!("location:{}", location.id);
    let location_tuples = location.to_tuples();
    if let Err(err) = con.hset_multiple::<String, &str, String, ()>(key, &location_tuples) {
        return Err(warp::reject::custom(RedisErrorType(err)))
    }
    if let Err(err) = con.geo_add::<&str, (f64, f64, String), ()>("locations", (location3.longitude, location3.latitude, location3.id)) {
        return Err(warp::reject::custom(RedisErrorType(err)))
    }
    Ok(warp::reply::with_status(
        format!("Location added {:?}", location2),
        StatusCode::CREATED
    ))
}

async fn find_locations_by_distance(params: HashMap<String, String>) -> Result<impl Reply, Rejection> {
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
        lat: f64,
        lon: f64,
    }

    #[derive(Serialize)]
    pub struct RadiusSearchResultSerializable {
        name: String,
        coord: Option<CoordSerializable>,
        dist: Option<f64>,
    }

    impl From<&RadiusSearchResult> for RadiusSearchResultSerializable {
        fn from(rsr: &RadiusSearchResult) -> Self {
            RadiusSearchResultSerializable {
                name: rsr.name.clone(),
                coord: rsr.coord.as_ref().map(|c| CoordSerializable {
                    lat: c.latitude,
                    lon: c.longitude,
                }),
                dist: rsr.dist,
            }
        }
    }

    let mut con = connect();
    let opts = RadiusOptions::default().with_dist().order(RadiusOrder::Asc);
    let result: RedisResult<Vec<RadiusSearchResult>> = con.geo_radius("locations", longitude, latitude, radius, Unit::Kilometers, opts);
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

async fn delete_location_by_id(id: String) -> Result<impl Reply, Rejection> {
    let mut con = connect();
    let key = format!("location:{}", id);
    match con.hgetall::<String, HashMap<String, String>>(key.clone()) {
        Ok(loc) if loc.is_empty() => Ok(warp::reply::json(&format!("Location not found by id: {}", id))),
        Err(err) => Err(warp::reject::custom(RedisErrorType(err))),
        Ok(loc) => {
            match con.del::<String, ()>(key) {
                Ok(_) => Ok(warp::reply::json(&format!("Location deleted: {:?}", loc))),
                Err(err) => Err(warp::reject::custom(RedisErrorType(err)))
            }
        }
    }
}

#[tokio::main]
async fn main() {
    let ping_route = warp::path("ping")
        .and_then(ping_handler);

    let get_location_by_id_route = warp::get()
        .and(warp::path("location"))
        .and(warp::path::param())
        .and(warp::path::end())
        .and_then(load_location_by_id);

    let post_location_route = warp::post()
        .and(warp::path("location"))
        .and(warp::path::end())
        .and(warp::body::json())
        .and_then(save_location);

    let get_locations_by_distance_route = warp::get()
        .and(warp::path("location"))
        .and(warp::path("search"))
        .and(warp::path::end())
        .and(warp::query())
        .and_then(find_locations_by_distance);

    let delete_location_by_id_route = warp::delete()
        .and(warp::path("location"))
        .and(warp::path::param())
        .and(warp::path::end())
        .and_then(delete_location_by_id);

    let routes = ping_route
        .or(get_locations_by_distance_route)
        .or(get_location_by_id_route)
        .or(post_location_route)
        .or(delete_location_by_id_route)
        .recover(handle_rejection);

    warp::serve(routes)
        .run(([127, 0, 0, 1], 8080))
        .await;
}
