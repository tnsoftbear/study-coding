#![warn(clippy::all, clippy::pedantic)]

extern crate redis;

mod model;
mod storage;
mod controller;

use std::env;
use std::net::IpAddr;
use crate::controller::route::routing;

#[tokio::main]
async fn main() {
    log4rs::init_file("config/log4rs.yaml", Default::default()).unwrap();
    let routes = routing::build_routes();
    let host = env::var("APP_HOST")
        .unwrap_or("0.0.0.0".to_string())
        .parse::<IpAddr>()
        .expect("APP_HOST env variable must be valid IP address");
    let port = env::var("APP_PORT")
        .unwrap_or("8081".to_string())
        .parse::<u16>()
        .expect("APP_PORT env variable must be valid port number");
    warp::serve(routes)
        .run((host, port))
        .await;
}
