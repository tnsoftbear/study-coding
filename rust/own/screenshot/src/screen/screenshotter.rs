use screenshots::Screen;
use crate::date::clock::{Clock, CurrentDateLocal};

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

fn make_filename(clock: &dyn CurrentDateLocal) -> String {
    let now = clock.now_local();
    let now_formatted = now.format("%Y-%m-%d_%H-%M-%S_%f");
    let filename = format!("{now_formatted}.png");
    filename
}

#[cfg(test)]
mod tests {
    use chrono::{DateTime, Local, TimeZone};
    use super::*;

    #[test]
    fn test_make_filename() {
        let mut mock = crate::date::clock::MockCurrentDateLocal::new();
        mock.expect_now_local()
            .times(1)
            .returning(|| Local.with_ymd_and_hms(2024, 5, 7, 12, 0, 0).unwrap());
        assert_eq!(make_filename(&mock), "2024-05-07_12-00-00_000000000.png");
    }

    #[test]
    fn test_make_filename_2() {
        struct ClockMock;
        impl CurrentDateLocal for ClockMock {
            fn now_local(&self) -> DateTime<Local> {
                Local.with_ymd_and_hms(2024, 5, 7, 12, 0, 0).unwrap()
            }
        }
        let result = make_filename(&ClockMock);
        assert_eq!(result, "2024-05-07_12-00-00_000000000.png");
    }
}