use serde::Serialize;
use warp::{Rejection, Reply};

pub async fn ping_handler(id: String) -> Result<impl Reply, Rejection> {
    #[derive(Serialize)]
    struct PingResponse {
        message: &'static str,
    }
    log::info!("Ping called, id: {id}");
    Ok(warp::reply::json(&PingResponse { message: "pong" }))
}