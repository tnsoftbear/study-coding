use chrono::{DateTime, Utc};
use mockall::*;

pub struct Clock;

impl Clock {
    pub fn new() -> Self {
        Clock
    }
}

#[automock]
pub trait CurrentDate {
    fn get_current_date_utc(&self) -> DateTime<Utc>;
}

impl CurrentDate for Clock {
    fn get_current_date_utc(&self) -> DateTime<Utc> {
        Utc::now()
    }
}