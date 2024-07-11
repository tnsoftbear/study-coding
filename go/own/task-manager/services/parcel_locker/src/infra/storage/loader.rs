use crate::domain::model::parcel_locker::ParcelLocker;
use crate::redis::Commands;
use crate::infra::storage::common::{connect, make_parcel_locker_key};
use redis::{RedisError, RedisResult};
use std::collections::HashMap;

pub enum LoadError {
    NotFound,
    RedisError(RedisError),
}

pub fn load_parcel_locker_by_id(id: &str) -> Result<ParcelLocker, LoadError> {
    let mut con = connect();
    let key = make_parcel_locker_key(id);
    match con.hgetall::<String, HashMap<String, String>>(key) {
        Ok(pl_hm) => {
            if pl_hm.is_empty() {
                return Err(LoadError::NotFound);
            }
            Ok(pl_hm.into())
        }
        Err(err) => Err(LoadError::RedisError(err)),
    }
}

pub fn load_parcel_lockers(page: isize, per_page: isize) -> Result<Vec<ParcelLocker>, RedisError> {
    let start = (page - 1) * per_page;
    let stop = start + per_page - 1;

    let mut con = connect();
    let result: RedisResult<Vec<String>> = con.zrange("parcel_lockers", start, stop);
    match result {
        Ok(parcel_locker_ids) => {
            let mut parcel_lockers: Vec<ParcelLocker> = Vec::new();
            for id in parcel_locker_ids {
                let key = make_parcel_locker_key(&id);
                match con.hgetall::<String, HashMap<String, String>>(key) {
                    Ok(pl_hm) => parcel_lockers.push(pl_hm.into()),
                    Err(err) => return Err(err),
                }
            }
            Ok(parcel_lockers)
        }
        Err(err) => Err(err),
    }
}

