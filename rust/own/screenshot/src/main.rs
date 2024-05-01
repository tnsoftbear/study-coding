#![warn(clippy::all, clippy::pedantic)]

mod screenshot;
mod path;
mod keyboard;

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let screens_dir = args.get(1).unwrap_or(&String::new()).to_string(); // _or(&storage_dir::TARGET_DIR.to_string()).to_string();
    let storage_root_path = path::storage_dir::prepare(screens_dir);
    keyboard::bind_print_screen(&storage_root_path);
}


