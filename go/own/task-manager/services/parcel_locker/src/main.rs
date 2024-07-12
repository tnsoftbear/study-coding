#![warn(clippy::all, clippy::pedantic)]

extern crate redis;

mod app;
mod domain;
mod infra;

use std::env;
use std::net::IpAddr;

#[tokio::main]
async fn main() {
    infra::trace::tracing::init();
    let routes = app::rest::route::routing::build_routes();
    let host = env::var("APP_HOST")
        .unwrap_or("0.0.0.0".to_string())
        .parse::<IpAddr>()
        .expect("APP_HOST env variable must be valid IP address");
    let port = env::var("APP_PORT")
        .unwrap_or("8081".to_string())
        .parse::<u16>()
        .expect("APP_PORT env variable must be valid port number");
    warp::serve(routes).run((host, port)).await;
}
