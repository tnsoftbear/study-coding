use crate::domain::model::ParcelLocker;
use crate::infra::storage::common::{connect, make_parcel_locker_key};
use crate::redis::Commands;
use redis::{Connection, RedisError};

pub fn save_parcel_locker(parcel_locker: &ParcelLocker) -> Result<bool, RedisError> {
    let mut con = connect();
    let parcel_locker_clone = parcel_locker.clone();
    let key = make_parcel_locker_key(&parcel_locker.id);
    let is_new = match con.exists::<String, bool>(key.clone()) {
        Ok(exists) => !exists,
        Err(err) => return Err(err),
    };
    let parcel_locker_tuples = parcel_locker.to_tuples();
    do_atomic_save(
        &mut con,
        parcel_locker_clone,
        parcel_locker_tuples.as_slice(),
        &key,
    )?;
    Ok(is_new)
}

fn do_atomic_save(
    con: &mut Connection,
    parcel_locker: ParcelLocker,
    parcel_locker_tuples: &[(&str, String)],
    key: &str,
) -> Result<(), RedisError> {
    redis::pipe()
        .atomic()
        .cmd("HSET")
        .arg(key)
        .arg(parcel_locker_tuples)
        .cmd("GEOADD")
        .arg("parcel_lockers")
        .arg(parcel_locker.longitude)
        .arg(parcel_locker.latitude)
        .arg(parcel_locker.id)
        .query(con)?;
    Ok(())
}

// Alternative implementation without transaction
#[allow(dead_code)]
fn do_regular_save(
    con: &mut Connection,
    parcel_locker: ParcelLocker,
    parcel_locker_tuples: &[(&str, String)],
    key: &str,
) -> Result<(), RedisError> {
    con.hset_multiple::<&str, &str, String, ()>(key, parcel_locker_tuples)?;
    con.geo_add::<&str, (f64, f64, String), ()>(
        "parcel_lockers",
        (
            parcel_locker.longitude,
            parcel_locker.latitude,
            parcel_locker.id,
        ),
    )?;
    Ok(())
}
