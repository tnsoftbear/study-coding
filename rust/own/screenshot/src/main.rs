#![warn(clippy::all, clippy::pedantic)]

mod path;
mod keyboard;
mod cli;
mod screen;

fn main() {
    let cli_input = cli::reader::read();
    let storage_root_path = path::storage_dir::determine_root_path(&cli_input);
    keyboard::key_binder::bind(&storage_root_path);
}
