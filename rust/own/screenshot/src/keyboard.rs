use rdev::{Event, EventType, grab, Key};
use crate::screenshot;

pub fn bind_print_screen(storage_root_path: &str) {
    let s = storage_root_path.to_string().clone();
    if let Err(error) = grab(move |event| grab_cb(event, &s)) {
        panic!("ERROR: {error:?}");
    }
}

fn grab_cb(event: Event, storage_root_path: &str) -> Option<Event> {
    match event.event_type {
        EventType::KeyPress(Key::PrintScreen) => {
            screenshot::make(storage_root_path);
            None
        }
        _ => Some(event),
    }
}