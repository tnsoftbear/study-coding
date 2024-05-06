use chrono::Utc;
use screenshots::Screen;

pub fn make(storage_path_string: &str) {
    let screens = Screen::all().unwrap();
    for screen in screens {
        let image = screen.capture().unwrap();
        let now = Utc::now();
        let now_formatted = now.format("%Y-%m-%d_%H_%M_%S_%f");
        let result = image.save(format!("{storage_path_string}/{now_formatted}.png"));
        assert!(result.is_ok(), "ERROR: When grabbing screen: {:?}", result.err());
    }
}
