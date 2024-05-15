use super::question_id::QuestionId;

#[derive(Debug)]
pub struct Question {
    id: QuestionId,
    title: String,
    content: String,
    tags: Option<Vec<String>>
}

impl Question {
    pub fn new(
        id: QuestionId,
        title: String,
        content: String,
        tags: Option<Vec<String>>
    ) -> Self {
        Question{id, title, content, tags}
    }
}

impl std::fmt::Display for Question {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> Result<(), std::fmt::Error> {
        write!(f, "{}, Q: {}, A: {}, tags: {:?}", self.id, self.title, self.content, self.tags.as_ref().unwrap())
    }
}
