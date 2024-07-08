extern crate redis;

use std::env;

pub fn connect() -> redis::Connection {
    let redis_host = env::var("REDIS_HOST").unwrap_or_else(|_| "127.0.0.1".to_string());
    let redis_port = env::var("REDIS_PORT").unwrap_or_else(|_| "6379".to_string());
    let redis_url = format!("redis://{redis_host}:{redis_port}");
    let client = redis::Client::open(redis_url).unwrap();
    client.get_connection().unwrap()
}

pub fn make_parcel_locker_key(id: &str) -> String {
    format!("parcel_locker:{id}")
}
