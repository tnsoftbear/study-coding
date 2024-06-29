use warp::http::StatusCode;
use crate::controller::rejection::errors::StorageError;
use crate::model::parcel_locker::ParcelLocker;
use crate::storage::saver;

pub async fn save_parcel_locker(parcel_locker: ParcelLocker) -> Result<impl warp::Reply, warp::Rejection> {
    match saver::save_parcel_locker(&parcel_locker) {
        Ok(is_new) => Ok(warp::reply::with_status(
            warp::reply::json::<ParcelLocker>(&parcel_locker),
            if is_new { StatusCode::CREATED } else { StatusCode::OK }
        )),
        Err(err) => Err(warp::reject::custom(StorageError(err)))
    }
}
