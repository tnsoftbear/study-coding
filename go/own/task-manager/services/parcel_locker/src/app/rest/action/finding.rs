use crate::app::rest::rejection::app_errors::{BadRequestError, StorageError};
use crate::infra::storage::finder;
use std::collections::HashMap;
use tracing::instrument;
use warp::{Rejection, Reply};

#[instrument]
pub async fn find_parcel_lockers_by_distance(
    params: HashMap<String, String>,
) -> Result<impl Reply, Rejection> {
    let longitude = match params.get("longitude") {
        Some(longitude) => match longitude.parse::<f64>() {
            Ok(longitude) if longitude >= -180.0 && longitude <= 180.0 => longitude,
            Err(_) => {
                return Err(warp::reject::custom(BadRequestError::ParameterNotNumeric(
                    "longitude".to_string(),
                )))
            }
            _ => {
                return Err(warp::reject::custom(BadRequestError::InvalidValue(
                    "longitude must be between -180 and 180 degrees".to_string(),
                )))
            }
        },
        None => {
            return Err(warp::reject::custom(BadRequestError::ParameterRequired(
                "longitude".to_string(),
            )))
        }
    };
    let latitude = match params.get("latitude") {
        Some(latitude) => match latitude.parse::<f64>() {
            Ok(latitude) if latitude >= -90.0 && latitude <= 90.0 => latitude,
            Err(_) => {
                return Err(warp::reject::custom(BadRequestError::ParameterNotNumeric(
                    "latitude".to_string(),
                )))
            }
            _ => {
                return Err(warp::reject::custom(BadRequestError::InvalidValue(
                    "latitude must be between -90 and 90 degrees".to_string(),
                )))
            }
        },
        None => {
            return Err(warp::reject::custom(BadRequestError::ParameterRequired(
                "latitude".to_string(),
            )))
        }
    };
    let distance = match params.get("distance") {
        Some(distance) => match distance.parse::<f64>() {
            Ok(distance) => distance,
            Err(_) => {
                return Err(warp::reject::custom(BadRequestError::ParameterNotNumeric(
                    "distance".to_string(),
                )))
            }
        },
        None => {
            return Err(warp::reject::custom(BadRequestError::ParameterRequired(
                "distance".to_string(),
            )))
        }
    };

    match finder::find_parcel_lockers_by_distance(longitude, latitude, distance) {
        Ok(serializable_results) => Ok(warp::reply::json(&serializable_results)),
        Err(err) => Err(warp::reject::custom(StorageError(err))),
    }
}
