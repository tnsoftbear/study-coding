use std::ops::RangeInclusive;
use std::path::PathBuf;
use clap::{arg, Command, ArgAction, ArgGroup, value_parser, Arg};

const PORT_RANGE: RangeInclusive<usize> = 1..=65535;

fn port_in_range(s: &str) -> Result<u16, String> {
    let port: usize = s
        .parse()
        .map_err(|_| format!("`{s}` isn't a port number"))?;
    if PORT_RANGE.contains(&port) {
        Ok(port as u16)
    } else {
        Err(format!(
            "port not in range {}-{}",
            PORT_RANGE.start(),
            PORT_RANGE.end()
        ))
    }
}

fn cmd() -> Command {
    Command::new("Example App")
        .version(env!("CARGO_PKG_VERSION"))
        .about("My example app about")
        .arg(arg!([name] "Optional name to operate on"))
        .subcommand(
            // .\target\debug\parse_args.exe help dev
            Command::new("dev")
                .about("Do development things")
                // ... (three consecutive dots/periods) specifies that this argument may occur multiple times
                .arg(arg!(-d --debug ... "Turn debugging on")) // это флаг, можно повторять -ddd
                .arg(
                    Arg::new("VERBOSEID")
                        .short('v')
                        .action(ArgAction::Count)   // -vvv
                )
                // Только один аргумент из группы может присутствовать во вводе
                .args([
                    arg!(--"set-ver" <ver> "set the version manually").required(false),
                    arg!(--major "auto increase major"),
                    arg!(--minor "auto increase minor"),
                    arg!(--patch "auto increase patch")
                ])
                .group(ArgGroup::new("vers")
                    .args(["set-ver", "major", "minor","patch"])
                    .required(true))
        )
        .args([
            arg!(-l --list "list test values").action(ArgAction::SetTrue),
            arg!(-c --config <FILE> "Set config file, <FILE> is required") // <FILE> - required
                .required(false)
                .value_parser(value_parser!(PathBuf)),
            arg!(--"config-more" [FILE] "Set config file, [FILE] is optional") // [FILE] - optional
                .required(false)
                .value_parser(value_parser!(PathBuf)),
            arg!(--"config-def" [FILE] "Set config default file, [FILE] is optional") // [FILE] - optional
                .required(false)
                .default_value("~/.config") // это значение вернётся, если аргумент --config-more не передан, не путать с пустым значением, когда только config-more передан
                .value_parser(value_parser!(PathBuf)),
            Arg::new("PODID")
                .short('p')
                .long("pod")
                .num_args(1)
                .action(ArgAction::Append) // -p POD123 -p POD321
                .help("Set POD number")
        ])
        .arg(
            Arg::new("FALSEYID")
                .short('f')
                .long("falsey")
                .value_parser(clap::builder::FalseyValueParser::new()) // false for "", No, oFF, 0, etc
        )
        .arg(
            arg!(-m --mode <MODE> "What mode to run program?")
                .value_parser(["slow", "normal", "fast"])
                .default_value("normal")
        )
        .arg(
            arg!(--port <PORT>)
                .help("Network port to use")
                .value_parser(value_parser!(u16).range(1..)),
        )
        .arg(
            arg!(--hostport <HOSTPORT>)
                .value_parser(port_in_range)
        )
}

fn main() {
    // Usage: parse_args.exe [OPTIONS] [name] [COMMAND]
    // .\target\debug\parse_args.exe -l -c E:\config.cfg -p POD123 -p POD321 -f oFF Joe test -ddd
    let matches = cmd().get_matches();

    if let Some(name) = matches.get_one::<String>("name") {
        println!("Value for name: {name}");
    }

    if let Some(config_path) = matches.get_one::<PathBuf>("config") {
        let filename = config_path.file_name().unwrap_or_default().to_str().unwrap();
        println!("Value for config path: {}, filename: {}", config_path.display(), filename);
    }

    if let Some(config_more_path) = matches.get_one::<PathBuf>("config-more") {
        let filename = config_more_path.file_name().unwrap_or_default().to_str().unwrap();
        println!("Value for config-more path: {}, filename: {}", config_more_path.display(), filename);
    }

    if let Some(config_def_path) = matches.get_one::<PathBuf>("config-def") {
        let filename = config_def_path.file_name().unwrap_or_default().to_str().unwrap();
        println!("Value for config-def path: {}, filename: {}", config_def_path.display(), filename);
    }

    if let Some(dev_cmd_matches) = matches.subcommand_matches("dev") {
        if let Some(debug) = dev_cmd_matches.get_one::<u8>("debug") {
            println!("Count of debug flags: {}", debug);
            match debug {
                0 => println!("Debug if off"),
                1..=2 => println!("Debug is on"),
                _ => println!("Too much debug")
            }
        }

        if let Some(debug) = dev_cmd_matches.get_one::<u8>("VERBOSEID") {
            println!("Count of verbose flags: {}", debug);
        }
    }

    if let Some(pod) = matches.get_one::<String>("PODID") {
        println!("Pod is {pod}");
        let pods = matches
            .get_many::<String>("PODID")
            .unwrap_or_default()
            .map(|v| v.as_str())
            .collect::<Vec<_>>();
        println!("Pods: {:?}", &pods);
    }

    if matches.get_flag("list") {
        println!("Printing list enabled");
    }

    if let Some(falsey) = matches.get_one::<bool>("FALSEYID") {
        println!("Falsey: {falsey}");
    }

    match matches
        .get_one::<String>("mode")
        .expect("MODE required")
        .as_str()
    {
        "fast" => println!("Fast as eagle"),
        "slow" => println!("Slow as turtle"),
        "normal" => println!("Normal speed"),
        _ => unreachable!(),
    }

    if let Some(port) = matches.get_one::<u16>("port") {
        println!("Port: {port}");
    }

    if let Some(hostport) = matches.get_one::<u16>("hostport") {
        println!("Host port: {hostport}");
    }

}

#[test]
fn verify_cmd() {
    cmd().debug_assert();
}