fn main() {
    let mut counter = 0;
    'loop_a: loop {
        counter += 1;
        if counter == 10 {
            // break;
            break 'loop_a;
        }
        'loop_b: loop {
            if counter < 5 {
                continue 'loop_a;
            } else {
                break 'loop_b;
            }
        }
        println!("counter: {}", counter);
    }
}