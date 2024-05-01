#![warn(clippy::all, clippy::pedantic)]

mod storage_dir;

use chrono::Utc;
use rdev::{grab, Event, EventType, Key};
use screenshots::Screen;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let screens_dir = args.get(1).unwrap_or(&String::new()).to_string(); // _or(&storage_dir::TARGET_DIR.to_string()).to_string();
    let storage_path_string = storage_dir::prepare(screens_dir);

    if let Err(error) = grab(move |e| grab_cb(e, &storage_path_string)) {
        panic!("ERROR: {error:?}");
    }
}

fn grab_cb(event: Event, storage_path_string: &str) -> Option<Event> {
    match event.event_type {
        EventType::KeyPress(Key::PrintScreen) => {
            make_screenshot(storage_path_string);
            None
        }
        _ => Some(event),
    }
}

fn make_screenshot(storage_path_string: &str) {
    let screens = Screen::all().unwrap();
    for screen in screens {
        let image = screen.capture().unwrap();
        let now = Utc::now();
        let now_formatted = now.format("%Y-%m-%d_%H_%M_%S_%f");
        let result = image.save(format!("{}/{}.png", storage_path_string, now_formatted));
        if result.is_err() {
            panic!("ERROR: When grabbing screen: {:?}", result.err());
        }
    }
}
