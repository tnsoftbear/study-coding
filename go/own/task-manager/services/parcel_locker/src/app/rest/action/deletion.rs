use crate::app::rest::rejection::app_errors::StorageError;
use crate::domain::model::ParcelLocker;
use crate::infra::storage::deleter;
use crate::infra::storage::deleter::DeletionError;
use serde::Serialize;
use tracing::instrument;
use warp::{Rejection, Reply};

#[instrument]
pub async fn delete_parcel_locker_by_id(id: String) -> Result<impl Reply, Rejection> {
    #[derive(Serialize)]
    struct Response {
        deleted: bool,
        message: String,
        parcel_locker: Option<ParcelLocker>,
    }

    match deleter::delete_parcel_locker_by_id(&id) {
        Ok(parcel_locker) => Ok(warp::reply::json(&Response {
            deleted: true,
            message: "Parcel locker deleted".to_string(),
            parcel_locker: Some(parcel_locker),
        })),
        Err(err) => match err {
            DeletionError::RedisErrorType(e) => Err(warp::reject::custom(StorageError(e))),
            DeletionError::EntryNotFound => Ok(warp::reply::json(&Response {
                deleted: false,
                message: format!("Parcel locker not found by id: {id}"),
                parcel_locker: None,
            })),
        },
    }
}

#[instrument]
pub async fn delete_all_parcel_lockers() -> Result<impl Reply, Rejection> {
    #[derive(Serialize)]
    struct Response {
        message: String,
    }

    match deleter::delete_all_parcel_lockers() {
        Ok(()) => Ok(warp::reply::json(&Response {
            message: "Parcel lockers deleted".to_string(),
        })),
        Err(err) => Err(warp::reject::custom(StorageError(err))),
    }
}
