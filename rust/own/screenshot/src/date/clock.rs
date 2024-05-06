use chrono::{DateTime, Utc};

pub struct Clock;

impl Clock {
    pub fn new() -> Self {
        Clock
    }
}

pub trait CurrentDate {
    fn get_current_date(&self) -> DateTime<Utc>;
}

impl CurrentDate for Clock {
    fn get_current_date(&self) -> DateTime<Utc> {
        Utc::now()
    }
}