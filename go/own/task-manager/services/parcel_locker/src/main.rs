#![warn(clippy::all, clippy::pedantic)]

extern crate redis;

mod model;
mod storage;
mod controller;

use crate::controller::route::routing;

#[tokio::main]
async fn main() {
    let routes = routing::build_routes();
    warp::serve(routes)
        .run(([127, 0, 0, 1], 8081))
        .await;
}
