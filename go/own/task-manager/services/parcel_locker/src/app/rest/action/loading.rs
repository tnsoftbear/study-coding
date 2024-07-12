use crate::app::rest::rejection::app_errors;
use crate::app::rest::request::param_extractor;
use crate::domain::repository::{LoadError, Loading};
use std::collections::HashMap;
use std::fmt::Debug;
use tracing::instrument;
use warp::{Rejection, Reply};

#[instrument]
pub async fn load_parcel_locker_by_id(
    id: String,
    loader: impl Loading + Debug,
) -> Result<impl Reply, Rejection> {
    match loader.load_parcel_locker_by_id(&id) {
        Ok(parcel_locker) => Ok(warp::reply::json(&parcel_locker)),
        Err(LoadError::NotFound) => Ok(warp::reply::json(&())),
        Err(LoadError::StorageError(err)) => {
            Err(warp::reject::custom(app_errors::StorageError(err)))
        }
    }
}

#[instrument]
pub async fn load_parcel_lockers_paginated(
    params: HashMap<String, String>,
    loader: impl Loading + Debug,
) -> Result<impl Reply, Rejection> {
    let pagination_params = match param_extractor::extract_pagination(params) {
        Ok(pagination_params) => pagination_params,
        Err(param_extractor::PaginationParamError::NotPositiveIntError((key, _)))
            if key == param_extractor::PAGE_PARAM =>
        {
            return Err(warp::reject::custom(
                app_errors::BadRequestError::ParameterNotNumeric(
                    param_extractor::PAGE_PARAM.to_string(),
                ),
            ))
        }
        Err(param_extractor::PaginationParamError::NotPositiveIntError(_)) => {
            return Err(warp::reject::custom(
                app_errors::BadRequestError::ParameterNotNumeric(
                    param_extractor::PER_PAGE_PARAM.to_string(),
                ),
            ))
        }
    };
    match loader.load_parcel_lockers(pagination_params.page, pagination_params.per_page) {
        Ok(parcel_lockers) => Ok(warp::reply::json(&parcel_lockers)),
        Err(LoadError::NotFound) => Ok(warp::reply::json(&())),
        Err(LoadError::StorageError(err)) => {
            Err(warp::reject::custom(app_errors::StorageError(err)))
        }
    }
}
