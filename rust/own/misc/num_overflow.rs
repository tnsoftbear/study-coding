mod float;

use std::num::Wrapping;
use std::{thread, time::Duration};

fn main() {
    // quiz1();
    // quiz2_1();
    // quiz2_2();
    quiz3();
}

// Panic: attempt to add with overflow
#[allow(dead_code)]
fn quiz1() {
    let mut n: i8 = 0;
    loop {
        n += 1;
        println!("Number is {n}");
    }
}

#[allow(dead_code)]
fn quiz2_1() {
    let mut n: Wrapping<i8> = Wrapping(0i8);
    loop {
        n += Wrapping(1i8);
        println!("Number is {n}");
    }
}

#[allow(dead_code)]
fn quiz2_2() {
    let mut n: i8 = 0;
    loop {
        n = n.wrapping_add(1);
        println!("Number is {n}");
    }
}

#[allow(dead_code)]
fn quiz3() {
    let mut n: i8 = 0i8;
    loop {
        n = n.checked_add(1).unwrap_or(0);
        println!("Number is {n}");
        //thread::sleep(Duration::from_millis(1000));
    }
}