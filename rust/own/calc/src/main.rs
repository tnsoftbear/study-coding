#![warn(clippy::all, clippy::pedantic)]

use warp::{Filter, reply, reject, Rejection, Reply};
use warp::cors::CorsForbidden;
use warp::http::{Method, StatusCode};

#[derive(Debug)]
struct InvalidParam {
    value: String,
    message: String
}
impl reject::Reject for InvalidParam {}
impl InvalidParam {
    fn new(value: String, message: String) -> Self {
        InvalidParam {
            value,
            message
        }
    }
}

async fn sum_handler(first_param: String, second_param: String, request_id: String) -> Result<impl Reply, Rejection> {
    log::info!("request id: {request_id}");
    match first_param.parse::<i32>() {
        Ok(first_num) => {
            match second_param.parse::<i32>() {
                Ok(second_num) => {
                    let sum_num = first_num + second_num;
                    Ok(reply::json(&format!("{first_num} + {second_num} = {sum_num}")))
                }
                Err(err) => Err(reject::custom(
                    InvalidParam::new(
                        second_param.clone(),
                        format!("Problem with the second value (\"{second_param}\") - {err}")
                    )
                ))
            }
        }
        Err(err) => Err(reject::custom(
            InvalidParam::new(
                first_param.clone(),
                format!("Problem with the first value (\"{first_param}\") - {err}")
            )
        ))
    }
}

async fn return_error(r: Rejection) -> Result<impl Reply, Rejection> {
    if let Some(error) = r.find::<CorsForbidden>() {
        Ok(reply::with_status(
            error.to_string(),
            StatusCode::FORBIDDEN
        ))
    } else if let Some(err) = r.find::<InvalidParam>() {
        eprintln!("Invalid number error on value: {}. Error: {}", err.value, err.message.clone());
        Ok(reply::with_status(
            err.message.clone(),
            StatusCode::UNPROCESSABLE_ENTITY
        ))
    } else {
        Ok(reply::with_status(
            String::from("Route not found"),
            StatusCode::NOT_FOUND
        ))
    }
}

#[tokio::main]
async fn main() {
    log4rs::init_file("config/log4rs.yaml", Default::default()).unwrap();

    let id_filter = warp::any().map(|| uuid::Uuid::new_v4().to_string());

    let sum_route = warp::path("sum")
        .and(warp::path::param())
        .and(warp::path::param())
        .and(warp::path::end())
        .and(id_filter)
        .and_then(sum_handler)
        .recover(return_error);

    let cors = warp::cors()
        .allow_any_origin()
        .allow_header("content-type")
        .allow_methods(&[Method::PUT, Method::DELETE, Method::GET, Method::POST]);

    let log = warp::log::custom(|info| {
        eprintln!(
            "{} {} {} {:?} from {} with {:?}",
            info.method(),
            info.path(),
            info.status(),
            info.elapsed(),
            info.remote_addr().unwrap(),
            info.request_headers()
        )
    });

    let routes = sum_route
        .with(cors)
        .with(log);

    warp::serve(routes)
        .run(([127, 0, 0, 1], 8080))
        .await;
}
