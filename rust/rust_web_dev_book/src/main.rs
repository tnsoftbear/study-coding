#![warn(clippy::all, clippy::pedantic)]

mod entity;

use std::str::FromStr;
use entity::{question::Question, question_id::QuestionId};
use warp::Filter;

#[tokio::main]
async fn main() {
    let question = Question::new(
        QuestionId::from_str("1").expect("No id provided"), 
        "To be or not to be?".to_string(),
        "To be".to_string(),
        Some(vec!("faq".to_string()))
    );
    println!("{question:#?}");
    println!("{question}");

    let hello = warp::path("hello").map(|| format!("Hello, man!"));
    warp::serve(hello).run(([127, 0, 0, 1], 3030)).await;
}
