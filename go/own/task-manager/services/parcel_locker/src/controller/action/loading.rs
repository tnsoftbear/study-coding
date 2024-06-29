use std::collections::HashMap;
use warp::{Rejection, Reply};
use crate::controller::rejection::errors::{BadRequestError, StorageError};
use crate::storage::loader;

pub async fn load_parcel_locker_by_id(id: String) -> Result<impl Reply, Rejection> {
    match loader::load_parcel_locker_by_id(&id) {
        Ok(parcel_locker) => Ok(warp::reply::json(&parcel_locker)),
        Err(err) => Err(warp::reject::custom(StorageError(err)))
    }
}

pub async fn load_parcel_lockers_paginated(params: HashMap<String, String>) -> Result<impl Reply, Rejection> {
    let page = match params.get("page") {
        Some(page_param) => {
            match page_param.parse::<isize>() {
                Ok(page) => page,
                Err(_) => return Err(warp::reject::custom(
                    BadRequestError::ParameterNotNumeric("page".to_string())
                ))
            }
        }
        None => 1,
    };
    let per_page = match params.get("per_page") {
        Some(per_page_param) => {
            match per_page_param.parse::<isize>() {
                Ok(per_page) => per_page,
                Err(_) => return Err(warp::reject::custom(
                    BadRequestError::ParameterNotNumeric("per_page".to_string())
                ))
            }
        }
        None => 10,
    };

    match loader::load_parcel_lockers(page, per_page) {
        Ok(parcel_lockers) => Ok(warp::reply::json(&parcel_lockers)),
        Err(err) => Err(warp::reject::custom(StorageError(err)))
    }
}

pub async fn find_parcel_lockers_by_distance(params: HashMap<String, String>) -> Result<impl Reply, Rejection> {
    let longitude = match params.get("longitude") {
        Some(longitude) => {
            match longitude.parse::<f64>() {
                Ok(longitude) => longitude,
                Err(_) => return Err(warp::reject::custom(
                    BadRequestError::ParameterNotNumeric("longitude".to_string())
                ))
            }
        },
        None => return Err(warp::reject::custom(
            BadRequestError::ParameterRequired("longitude".to_string())
        ))
    };
    let latitude = match params.get("latitude") {
        Some(latitude) => {
            match latitude.parse::<f64>() {
                Ok(latitude) => latitude,
                Err(_) => return Err(warp::reject::custom(
                    BadRequestError::ParameterNotNumeric("latitude".to_string())
                ))
            }
        },
        None => return Err(warp::reject::custom(
            BadRequestError::ParameterRequired("latitude".to_string())
        ))
    };
    let radius = match params.get("radius") {
        Some(radius) => {
            match radius.parse::<f64>() {
                Ok(radius) => radius,
                Err(_) => return Err(warp::reject::custom(
                    BadRequestError::ParameterNotNumeric("radius".to_string())
                ))
            }
        },
        None => return Err(warp::reject::custom(
            BadRequestError::ParameterRequired("radius".to_string())
        ))
    };

    match loader::find_parcel_lockers_by_distance(longitude, latitude, radius) {
        Ok(serializable_results) => Ok(warp::reply::json(&serializable_results)),
        Err(err) => Err(warp::reject::custom(StorageError(err)))
    }
}