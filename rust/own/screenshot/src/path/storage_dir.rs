use std::path::{Path, PathBuf};
use std::{env, fs};
use std::io;
use std::io::Write;
use crate::cli::reader::CliInput;

const SCREENS_DIR_DEF: &str = "storage";

pub fn determine_root_path(cli_input: &CliInput) -> String {
    let mut screens_dir: String = cli_input.screens_dir.to_string();
    loop {
        match do_prepare(&screens_dir) {
            Ok(root_path) => {
                return root_path;
            },
            Err(err) => {
                println!("ERROR: Cannot create directory for screens. {err}");
                let screens_root_path_def = make_storage_root_path_string(SCREENS_DIR_DEF);
                print!("Provide screens directory (Press <Enter> for default: \"{screens_root_path_def}\"): ");
                io::stdout().flush().expect("ERROR: Unable to flush output buffer.");
                let mut screens_dir_input: String = String::new();
                io::stdin()
                    .read_line(&mut screens_dir_input)
                    .expect("ERROR: on reading directory from command line");
                screens_dir = screens_dir_input.trim().to_string();
            }
        }
    }
}

fn make_storage_root_path_buf(screens_dir: &str) -> PathBuf {
    let mut storage_path_buf = env::current_dir().unwrap();
    storage_path_buf.push(screens_dir);
    storage_path_buf
}

fn make_storage_root_path_string(screens_dir: &str) -> String {
    let storage_path_buf = make_storage_root_path_buf(screens_dir);
    let storage_path_string = storage_path_buf.to_str().unwrap().to_string();
    storage_path_string
}

fn do_prepare(input_screens_dir: &str) -> Result<String, String> {
    let mut screens_dir: &str = SCREENS_DIR_DEF;
    if !input_screens_dir.is_empty() {
        screens_dir = input_screens_dir;
    }
    let storage_path_buf = make_storage_root_path_buf(screens_dir);
    let storage_path_string = make_storage_root_path_string(screens_dir);

    if check_path_exists(&storage_path_buf) {
        check_dir_ready(&storage_path_buf)?;
        println!("Screens save path is: {storage_path_string}");
    } else {
        if let Err(err) = fs::create_dir(storage_path_buf.clone()) {
            return Err(format!("{err}"));
        }
        println!("Directory for screens is created: {storage_path_string}");
    }

    Ok(storage_path_string)
}

fn check_path_exists(path: &Path) -> bool {
    fs::metadata(path).is_ok()
}

fn check_dir_ready(path: &Path) -> Result<(), String> {
    let path_str = path.to_str().unwrap();
    if let Ok(path_meta) = fs::metadata(path_str) {
        if !path_meta.is_dir() {
            return Err(format!("Screens save path \"{path_str}\" is not a directory"));
        }
        if path_meta.permissions().readonly() {
            return Err(format!("Screens save directory \"{path_str}\" is read-only"));
        }
    }
    Ok(())
}
