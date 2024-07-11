use serde::Serialize;
use tracing::instrument;
use warp::{Rejection, Reply};

#[instrument]
pub async fn ping_handler() -> Result<impl Reply, Rejection> {
    #[derive(Serialize)]
    struct PingResponse {
        message: &'static str,
    }
    Ok(warp::reply::json(&PingResponse { message: "pong" }))
}
