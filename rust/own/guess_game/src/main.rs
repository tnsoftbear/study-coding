use std::io;
use rand::Rng;
use std::cmp::Ordering;

fn main() {
    println!("Отгадай число!");
    
    let rand_num = rand::thread_rng()
        .gen_range(1..=100);
    
    loop {
        let mut input_str: String = String::new();

        io::stdin()
            .read_line(&mut input_str)
            .expect("Error on reading input");

        println!("guess: {input_str}");

        let result: Result<i32, _> = input_str
            .trim()
            .parse();

        let input_num: i32 = match result {
            Ok(n) => n,
            Err(error) => {
                println!("Error on number parsing: {error}");
                continue
            }
        };

        match input_num.cmp(&rand_num) {
            Ordering::Less => println!("Ваше число меньше"),
            Ordering::Greater => println!("Ваше число больше"),
            Ordering::Equal => {
                println!("Вы угадали число {input_num}!");
                break;
            }
        }
    }
}
