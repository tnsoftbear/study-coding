use crate::model::parcel_locker::ParcelLocker;
use crate::redis::Commands;
use crate::storage::common::{connect, make_parcel_locker_key};
use redis::geo::{RadiusOptions, RadiusOrder, RadiusSearchResult, Unit};
use redis::{RedisError, RedisResult};
use serde::Serialize;
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

///////////////////////////////////////////////////////////////////////////////////////////////////

#[derive(Serialize)]
pub struct RadiusSearchResultSerializable {
    name: String,
    latitude: f64,
    longitude: f64,
    distance: Option<f64>,
}

impl From<&RadiusSearchResult> for RadiusSearchResultSerializable {
    fn from(rsr: &RadiusSearchResult) -> Self {
        RadiusSearchResultSerializable {
            name: rsr.name.clone(),
            latitude: rsr.coord.as_ref().unwrap().latitude,
            longitude: rsr.coord.as_ref().unwrap().longitude,
            distance: rsr.dist,
        }
    }
}

pub fn find_parcel_lockers_by_distance(
    longitude: f64,
    latitude: f64,
    radius: f64,
) -> Result<Vec<RadiusSearchResultSerializable>, RedisError> {
    let mut con = connect();
    let opts = RadiusOptions::default()
        .with_dist()
        .with_coord()
        .order(RadiusOrder::Asc);
    let result: RedisResult<Vec<RadiusSearchResult>> = con.geo_radius(
        "parcel_lockers",
        longitude,
        latitude,
        radius,
        Unit::Kilometers,
        opts,
    );
    match result {
        Ok(search_results) => {
            let serializable_results: Vec<RadiusSearchResultSerializable> = search_results
                .iter()
                .map(RadiusSearchResultSerializable::from)
                .collect();
            Ok(serializable_results)
        }
        Err(err) => Err(err),
    }
}
