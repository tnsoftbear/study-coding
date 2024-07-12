use crate::infra::storage::common::connect;
use crate::redis::Commands;
use redis::geo::{RadiusOptions, RadiusOrder, RadiusSearchResult, Unit};
use redis::{RedisError, RedisResult};
use serde::Serialize;

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
    distance: f64,
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
        distance,
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
