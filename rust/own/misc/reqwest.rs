#![warn(clippy::all, clippy::pedantic)]

use std::collections::HashMap;

#[tokio::main]
async fn main() {
    let response = reqwest::get("https://httpbin.org/ip")
        .await.unwrap()
        .json::<HashMap<String, String>>()
        .await.unwrap();
    println!("{:#?}", response);
}
