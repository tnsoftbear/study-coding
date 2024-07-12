use crate::app::rest::rejection::app_errors::StorageError;
use crate::domain::model::ParcelLocker;
use crate::infra::storage::saver;
use tracing::instrument;
use warp::http::StatusCode;

#[instrument]
pub async fn save_parcel_locker(
    parcel_locker: ParcelLocker,
) -> Result<impl warp::Reply, warp::Rejection> {
    match saver::save_parcel_locker(&parcel_locker) {
        Ok(is_new) => Ok(warp::reply::with_status(
            warp::reply::json::<ParcelLocker>(&parcel_locker),
            if is_new {
                StatusCode::CREATED
            } else {
                StatusCode::OK
            },
        )),
        Err(err) => Err(warp::reject::custom(StorageError(err))),
    }
}
