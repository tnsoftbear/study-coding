use std::env;

pub struct CliInput {
    pub screens_dir: String
}

pub fn read() -> CliInput {
    let args: Vec<String> = env::args().collect();
    let screens_dir = args.get(1).unwrap_or(&String::new()).to_string();
    CliInput {
        screens_dir
    }
}