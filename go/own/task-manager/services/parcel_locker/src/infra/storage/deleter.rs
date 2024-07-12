use crate::domain::model::ParcelLocker;
use crate::infra::storage::common::{connect, make_parcel_locker_key};
use crate::redis::Commands;
use redis::RedisError;
use std::collections::HashMap;

pub enum DeletionError {
    RedisErrorType(RedisError),
    EntryNotFound,
}

pub fn delete_parcel_locker_by_id(id: &str) -> Result<ParcelLocker, DeletionError> {
    let mut con = connect();
    let key = make_parcel_locker_key(id);
    match con.hgetall::<String, HashMap<String, String>>(key.clone()) {
        Ok(pl_hm) if pl_hm.is_empty() => Err(DeletionError::EntryNotFound),
        Err(err) => Err(DeletionError::RedisErrorType(err)),
        Ok(pl_hm) => match con.del::<&str, ()>(&key) {
            Ok(()) => match con.zrem::<&str, &str, ()>("parcel_lockers", id) {
                Ok(()) => Ok(pl_hm.into()),
                Err(err) => Err(DeletionError::RedisErrorType(err)),
            },
            Err(err) => Err(DeletionError::RedisErrorType(err)),
        },
    }
}

pub fn delete_all_parcel_lockers() -> Result<(), RedisError> {
    let mut con = connect();
    match con.zrange::<&str, Vec<String>>("parcel_lockers", 0, -1) {
        Ok(pl_ids) => {
            for id in pl_ids {
                con.del::<String, ()>(make_parcel_locker_key(&id))?;
            }
            Ok(())
        }
        Err(err) => Err(err),
    }
}
