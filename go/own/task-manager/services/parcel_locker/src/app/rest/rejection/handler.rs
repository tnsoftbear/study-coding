use crate::app::rest::rejection::app_errors::{BadRequestError, StorageError};
use serde::Serialize;
use warp::body::BodyDeserializeError;
use warp::http::StatusCode;
use warp::{Rejection, Reply};

pub async fn reject(err: Rejection) -> Result<impl Reply, std::convert::Infallible> {
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
    } else if let Some(err) = err.find::<StorageError>() {
        let code = err.0.code().unwrap_or("");
        let category = err.0.category();
        let message = err.0.to_string();
        let status = format!("REDIS_ERROR: {message} (code: {code}, category: {category})");
        error_message = status;
        status_code = StatusCode::INTERNAL_SERVER_ERROR;
    } else if let Some(err) = err.find::<BodyDeserializeError>() {
        error_message = err.to_string();
        status_code = StatusCode::UNPROCESSABLE_ENTITY;
    } else if err.find::<warp::reject::MethodNotAllowed>().is_some() {
        // Silly hack for converting 405 -> 404
        // https://github.com/seanmonstar/warp/issues/77
        error_message = "404 page not found".to_string();
        status_code = StatusCode::NOT_FOUND;
    } else {
        error_message = "INTERNAL_SERVER_ERROR".to_string();
        status_code = StatusCode::INTERNAL_SERVER_ERROR;
    }
    log::info!("Request rejected: {error_message}");
    Ok(warp::reply::with_status(
        warp::reply::json(&Response {
            error: error_message,
        }),
        status_code,
    ))
}
