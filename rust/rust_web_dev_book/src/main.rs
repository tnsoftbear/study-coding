#![warn(clippy::all, clippy::pedantic)]

#[derive(Debug)]
struct QuestionId(String);

#[derive(Debug)]
struct Question {
    id: QuestionId,
    title: String,
    content: String,
    tags: Option<Vec<String>>
}

impl Question {
    fn new(
        id: QuestionId,
        title: String,
        content: String,
        tags: Option<Vec<String>>
    ) -> Self {
        Question{id, title, content, tags}
    }
}

impl std::fmt::Display for QuestionId {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> Result<(), std::fmt::Error> {
        write!(f, "id: {}", self.0)
    }
}

impl std::fmt::Display for Question {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> Result<(), std::fmt::Error> {
        write!(f, "{}, Q: {}, A: {}, tags: {:?}", self.id, self.title, self.content, self.tags.as_ref().unwrap())
    }
}

fn main() {
    let question = Question::new(
        QuestionId("1".to_string()), 
        "To be or not to be?".to_string(),
        "To be".to_string(),
        Some(vec!("faq".to_string()))
    );
    println!("{question:#?}");
    println!("{question}");
}
