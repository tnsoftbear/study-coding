use screenshots::Screen;
use crate::date::clock::{Clock, CurrentDate};

pub fn make(storage_root_path: &str) {
    let screens = Screen::all().unwrap();
    let clock = Clock::new();
    for screen in screens {
        let image = screen.capture().unwrap();
        let filename = make_filename(&clock);
        let result = image.save(format!("{storage_root_path}/{filename}"));
        assert!(result.is_ok(), "ERROR: When grabbing screen: {:?}", result.err());
    }
}

fn make_filename(clock: &dyn CurrentDate) -> String {
    let now = clock.get_current_date_utc();
    let now_formatted = now.format("%Y-%m-%d_%H-%M-%S_%f");
    let filename = format!("{now_formatted}.png");
    filename
}

#[cfg(test)]
mod tests {
    use chrono::{DateTime, TimeZone, Utc};
    use super::*;

    #[test]
    fn test_make_filename() {
        let mut mock = crate::date::clock::MockCurrentDate::new();
        mock.expect_get_current_date_utc()
            .times(1)
            .returning(|| Utc.with_ymd_and_hms(2024, 5, 7, 12, 0, 0).unwrap());
        assert_eq!(make_filename(&mock), "2024-05-07_12-00-00_000000000.png");
    }

    #[test]
    fn test_make_filename_2() {
        struct ClockMock;
        impl CurrentDate for ClockMock {
            fn get_current_date_utc(&self) -> DateTime<Utc> {
                let fixed_date_time = Utc.with_ymd_and_hms(2024, 5, 7, 12, 0, 0).unwrap();
                fixed_date_time
            }
        }
        let result = make_filename(&ClockMock);
        assert_eq!(result, "2024-05-07_12-00-00_000000000.png");
    }
}