use crate::app::rest::rejection::errors::{BadRequestError, StorageError};
use crate::infra::storage::loader;
use std::collections::HashMap;
use tracing::instrument;
use warp::{Rejection, Reply};

#[instrument]
pub async fn load_parcel_locker_by_id(id: String) -> Result<impl Reply, Rejection> {
    match loader::load_parcel_locker_by_id(&id) {
        Ok(parcel_locker) => Ok(warp::reply::json(&parcel_locker)),
        Err(loader::LoadError::NotFound) => Ok(warp::reply::json(&())),
        Err(loader::LoadError::RedisError(err)) => Err(warp::reject::custom(StorageError(err))),
    }
}

#[instrument]
pub async fn load_parcel_lockers_paginated(
    params: HashMap<String, String>,
) -> Result<impl Reply, Rejection> {
    let page = match params.get("page") {
        Some(page_param) => match page_param.parse::<isize>() {
            Ok(page) => page,
            Err(_) => {
                return Err(warp::reject::custom(BadRequestError::ParameterNotNumeric(
                    "page".to_string(),
                )))
            }
        },
        None => 1,
    };
    let per_page = match params.get("per_page") {
        Some(per_page_param) => match per_page_param.parse::<isize>() {
            Ok(per_page) => per_page,
            Err(_) => {
                return Err(warp::reject::custom(BadRequestError::ParameterNotNumeric(
                    "per_page".to_string(),
                )))
            }
        },
        None => 10,
    };

    match loader::load_parcel_lockers(page, per_page) {
        Ok(parcel_lockers) => Ok(warp::reply::json(&parcel_lockers)),
        Err(err) => Err(warp::reject::custom(StorageError(err))),
    }
}
