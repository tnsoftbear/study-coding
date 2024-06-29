use redis::RedisError;
use crate::model::parcel_locker::ParcelLocker;
use crate::storage::common::{connect, make_parcel_locker_key};
use crate::redis::Commands;

pub fn save_parcel_locker(parcel_locker: &ParcelLocker) -> Result<bool, RedisError> {
    let mut con = connect();
    let parcel_locker_clone = parcel_locker.clone();
    let key = make_parcel_locker_key(&parcel_locker.id);

    let is_new = match con.exists::<String, bool>(key.clone()) {
        Ok(exists) => !exists,
        Err(err) => return Err(err)
    };

    let parcel_locker_tuples = parcel_locker.to_tuples();
    con.hset_multiple::<String, &str, String, ()>(key, &parcel_locker_tuples)?;

    con.geo_add::<&str, (f64, f64, String), ()>(
        "parcel_lockers",
        (parcel_locker_clone.longitude, parcel_locker_clone.latitude, parcel_locker_clone.id)
    )?;

    Ok(is_new)
}
