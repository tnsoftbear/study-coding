extern crate redis;

pub fn connect() -> redis::Connection {
    let client = redis::Client::open("redis://127.0.0.1:6379") // #todo take connection info from env
        .unwrap();
    client.get_connection().unwrap()
}

pub fn make_parcel_locker_key(id: &str) -> String {
    format!("parcel_locker:{id}")
}
