use redis::RedisError;

#[derive(Debug)]
pub enum BadRequestError {
    ParameterNotNumeric(String),
    ParameterRequired(String),
    InvalidValue(String),
}

impl warp::reject::Reject for BadRequestError {}

impl std::fmt::Display for BadRequestError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match *self {
            BadRequestError::ParameterNotNumeric(ref param) => {
                write!(f, "\"{param}\" parameter must be numeric value")
            }
            BadRequestError::ParameterRequired(ref param) => {
                write!(f, "\"{param}\" parameter required")
            }
            BadRequestError::InvalidValue(ref message) => {
                write!(f, "{message}")
            }
        }
    }
}

#[derive(Debug)]
pub struct StorageError(pub RedisError);

impl warp::reject::Reject for StorageError {}
