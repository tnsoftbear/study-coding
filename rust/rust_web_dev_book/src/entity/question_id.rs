use std::{io::{Error, ErrorKind}};

#[derive(Debug)]
pub struct QuestionId(String);

impl std::fmt::Display for QuestionId {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> Result<(), std::fmt::Error> {
        write!(f, "id: {}", self.0)
    }
}

impl std::str::FromStr for QuestionId {
    type Err = std::io::Error;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s.is_empty() {
            false => Ok(QuestionId(s.to_string())),
            true => Err(Error::new(ErrorKind::InvalidInput, "No id provided"))
        }
    }
}
