#![warn(clippy::all, clippy::pedantic)]

use std::{env, thread};
use std::io::Write;
use std::path::{Path, PathBuf};
use std::thread::JoinHandle;

fn collect_input_file_root_paths(input_root_path_buf: &Path) -> Vec<PathBuf> {
    let mut input_file_root_path_bufs: Vec<PathBuf> = Vec::new();
    if let Ok(entries) = std::fs::read_dir(input_root_path_buf) {
        for entry in entries.flatten() {
            match entry.metadata() {
                Ok(metadata) => {
                    if metadata.is_dir() {
                        println!("{} is directory. Skipped.", entry.path().to_str().unwrap());
                        continue;
                    }
                    input_file_root_path_bufs.push(entry.path());
                },
                Err(err) => println!("Couldn't get metadata type for {:?}, because of: {err}", entry.path())
            }
        }
    }

    assert!(!input_file_root_path_bufs.is_empty(), "Input directory is empty");

    input_file_root_path_bufs
}

fn convert_one(input_file_root_path_buf: &Path, mut output_root_path_buf: PathBuf) {
    if let Some(input_root_path_file_str) = input_file_root_path_buf.to_str() {
        match std::fs::read_to_string(input_root_path_file_str) {
            Ok(markdown_input) => {
                let parser = pulldown_cmark::Parser::new(&markdown_input);
                let mut html_output = String::new();
                pulldown_cmark::html::push_html(&mut html_output, parser);

                let output_file_name = input_file_root_path_buf
                    .file_name()
                    .unwrap()
                    .to_str()
                    .unwrap()
                    .to_owned()
                    .trim_end_matches(".md")
                    .to_owned()
                    + ".html";
                output_root_path_buf.push(output_file_name);
                let output_file_root_path_str = output_root_path_buf.to_str().unwrap();

                let output_file_create_result = std::fs::File::create(&output_root_path_buf);
                match output_file_create_result {
                    Ok(mut output_file) => {
                        if let Err(err) = output_file.write_all(&html_output.into_bytes()) {
                            eprintln!("Cannot write to file {output_file_root_path_str}, because of error: {err}");
                        }
                    },
                    Err(err) => eprintln!("Cannot create output file {output_file_root_path_str}, because of error: {err}")
                }
            },
            Err(err) => eprintln!("Cannot read input file {input_root_path_file_str}, because of error: {err}")
        }
    }
}

fn make_input_root_path() -> PathBuf {
    let mut input_root_path_buf = env::current_dir().unwrap();
    input_root_path_buf.push("playground");
    input_root_path_buf.push("md");
    assert!(input_root_path_buf.is_dir(), "{} is not directory", input_root_path_buf.to_str().unwrap());
    input_root_path_buf
}

fn make_output_root_path() -> PathBuf {
    let mut output_root_path_buf = env::current_dir().unwrap();
    output_root_path_buf.push("playground");
    output_root_path_buf.push("html");
    assert!(output_root_path_buf.is_dir(), "{} is not directory", output_root_path_buf.to_str().unwrap());
    output_root_path_buf
}

fn main() {
    let input_root_path_buf = make_input_root_path();
    let output_root_path_buf = make_output_root_path();
    let input_file_root_path_bufs = collect_input_file_root_paths(&input_root_path_buf);
    let mut join_handles: Vec<JoinHandle<()>> = vec![];
    for input_file_root_path_buf in input_file_root_path_bufs {
        let output_root_path_buf_cloned = output_root_path_buf.clone();
        let join_handle = thread::spawn(move || {
            convert_one(&input_file_root_path_buf, output_root_path_buf_cloned);
        });
        join_handles.push(join_handle);
    }
    join_handles.into_iter().for_each(|j| j.join().unwrap());
}

// Alternative solution with channel: https://github.com/bodrovis-learning/Rust-YT-Series/blob/master/lesson15/src/main.rs
