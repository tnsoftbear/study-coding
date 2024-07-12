use crate::domain::model::ParcelLocker;
use crate::domain::repository::{LoadError, Loading};
use crate::infra::storage::common::{connect, make_parcel_locker_key};
use crate::redis::Commands;
use redis::RedisResult;
use std::collections::HashMap;

#[derive(Debug)]
pub struct Loader {}

impl Loading for Loader {
    fn load_parcel_locker_by_id(&self, id: &str) -> Result<ParcelLocker, LoadError> {
        let mut con = connect();
        let key = make_parcel_locker_key(id);
        match con.hgetall::<String, HashMap<String, String>>(key) {
            Ok(pl_hm) => {
                if pl_hm.is_empty() {
                    return Err(LoadError::NotFound);
                }
                Ok(pl_hm.into())
            }
            Err(err) => Err(LoadError::StorageError(err)),
        }
    }

    fn load_parcel_lockers(
        &self,
        page: usize,
        per_page: usize,
    ) -> Result<Vec<ParcelLocker>, LoadError> {
        let start = ((page - 1) * per_page) as isize;
        let stop = start + (per_page as isize) - 1;

        let mut con = connect();
        let result: RedisResult<Vec<String>> = con.zrange("parcel_lockers", start, stop);
        match result {
            Ok(parcel_locker_ids) => {
                let mut parcel_lockers: Vec<ParcelLocker> = Vec::new();
                for id in parcel_locker_ids {
                    let key = make_parcel_locker_key(&id);
                    match con.hgetall::<String, HashMap<String, String>>(key) {
                        Ok(pl_hm) => parcel_lockers.push(pl_hm.into()),
                        Err(err) => return Err(LoadError::StorageError(err)),
                    }
                }
                Ok(parcel_lockers)
            }
            Err(err) => Err(LoadError::StorageError(err)),
        }
    }
}
