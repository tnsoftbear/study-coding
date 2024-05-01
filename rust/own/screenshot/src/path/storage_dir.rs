use std::path::Path;
use std::{env, fs};

const TARGET_DIR_DEF: &str = "storage";

pub fn prepare(mut screens_dir: String) -> String {
    if screens_dir.is_empty() {
        screens_dir = TARGET_DIR_DEF.to_string();
    }
    let mut storage_path_buf = env::current_dir().unwrap();
    storage_path_buf.push(screens_dir);
    let storage_path_string = storage_path_buf.to_str().unwrap().to_string();
    if check_dir_exists(&storage_path_buf) {
        println!("Screens will be saved in directory: {storage_path_string}");
    } else {
        if let Err(err) = fs::create_dir(storage_path_buf.clone()) {
            panic!("ERROR: Cannot create directory for screens. Error: {err}");
        }
        println!("Directory for screens is created: {storage_path_string}");
    }
    verify_dir_ready(&storage_path_buf);
    storage_path_string
}

fn check_dir_exists(path: &Path) -> bool {
    fs::metadata(path).is_ok()
}

fn verify_dir_ready(path: &Path) {
    let path_str = path.to_str().unwrap();
    if let Ok(path_meta) = fs::metadata(path_str) {
        assert!(
            path_meta.is_dir(),
            "ERROR: Screens save path ({path_str}) is not a directory"
        );
        assert!(
            !path_meta.permissions().readonly(),
            "ERROR: Screens save directory ({path_str}) is read-only"
        );
    }
}
