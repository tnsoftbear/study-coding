use chrono::{DateTime, Local, Utc};
use mockall::automock;

pub struct Clock;

impl Clock {
    pub fn new() -> Self {
        Clock
    }
}

#[automock]
pub trait CurrentDateUtc {
    fn now_utc(&self) -> DateTime<Utc>;
}

impl CurrentDateUtc for Clock {
    fn now_utc(&self) -> DateTime<Utc> {
        Utc::now()
    }
}

#[automock]
pub trait CurrentDateLocal {
    fn now_local(&self) -> DateTime<Local>;
}

impl CurrentDateLocal for Clock {
    fn now_local(&self) -> DateTime<Local> {
        Local::now()
    }
}