use std::convert::Infallible;
use warp::{Filter, Reply};
use crate::controller::action::{deletion, health, loading, saving};
use crate::controller::rejection::handler::reject;

pub fn build_routes() -> impl Filter<Extract = impl Reply, Error = Infallible> + Clone {
    let ping_route = warp::path("ping")
        .and_then(health::ping_handler);

    let get_parcel_lockers_route = warp::get()
        .and(warp::path("parcel-lockers"))
        .and(warp::path::end())
        .and(warp::query())
        .and_then(loading::load_parcel_lockers_paginated);

    let get_parcel_locker_by_id_route = warp::get()
        .and(warp::path("parcel-locker"))
        .and(warp::path::param())
        .and(warp::path::end())
        .and_then(loading::load_parcel_locker_by_id);

    let get_parcel_lockers_by_distance_route = warp::get()
        .and(warp::path("parcel-locker-distance-search"))
        .and(warp::path::end())
        .and(warp::query())
        .and_then(loading::find_parcel_lockers_by_distance);

    let post_parcel_locker_route = warp::post()
        .and(warp::path("parcel-locker"))
        .and(warp::path::end())
        .and(warp::body::json())
        .and_then(saving::save_parcel_locker);

    let delete_parcel_locker_by_id_route = warp::delete()
        .and(warp::path("parcel-locker"))
        .and(warp::path::param())
        .and(warp::path::end())
        .and_then(deletion::delete_parcel_locker_by_id);

    let delete_all_parcel_lockers_route = warp::delete()
        .and(warp::path("all-parcel-lockers"))
        .and(warp::path::end())
        .and_then(deletion::delete_all_parcel_lockers);

    ping_route
        .or(get_parcel_lockers_route)
        .or(get_parcel_locker_by_id_route)
        .or(post_parcel_locker_route)
        .or(delete_parcel_locker_by_id_route)
        .or(get_parcel_lockers_by_distance_route)
        .or(delete_all_parcel_lockers_route)
        .recover(reject)
}