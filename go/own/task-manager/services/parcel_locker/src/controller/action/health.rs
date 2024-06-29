use serde::Serialize;
use warp::{Rejection, Reply};

pub async fn ping_handler() -> Result<impl Reply, Rejection> {
    #[derive(Serialize)]
    struct PingResponse {
        message: &'static str,
    }
    Ok(warp::reply::json(&PingResponse { message: "pong" }))
}