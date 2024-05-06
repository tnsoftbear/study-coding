#![warn(clippy::all, clippy::pedantic)]

mod screenshot;
mod path;
mod keyboard;

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let screens_dir = args.get(1).unwrap_or(&String::new()).to_string();
    let storage_root_path = path::storage_dir::determine_root_path(&screens_dir);
    keyboard::bind_print_screen(&storage_root_path);
}
