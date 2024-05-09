#![warn(clippy::all, clippy::pedantic)]

use std::rc::Rc;
use std::thread;
use std::time::Duration;
use std::sync::{Arc, mpsc, Mutex};

fn main() {
    // thread1();
    // thread2();
    // thread3();
    thread4();
}

#[allow(dead_code)]
fn thread1() {
    let a = 42;
    let handle = thread::spawn(move || {
        for i in 1..10 {
            println!("Hi number {} from the spawned thread!", i);
            thread::sleep(Duration::from_millis(1))
        }
        println!("{a}");
    });

    for i in 1..5 {
        println!("Hi number {} from the main thread!", i);
        thread::sleep(Duration::from_millis(1))
    }

    handle.join().unwrap();
}

#[allow(dead_code)]
fn thread2() {
    let (tx, rx) = mpsc::channel();
    thread::spawn(move || {
        thread::sleep(Duration::from_secs(1));
        tx.send("hi".to_string()).unwrap();
        thread::sleep(Duration::from_secs(3));
        tx.send("bye".to_string()).unwrap();
    });

    let received: String = rx.recv().unwrap();
    println!("Go-1: {received}");
    let received: String = rx.recv().unwrap();
    println!("Go-2: {received}");
}

#[allow(dead_code)]
fn thread3() {
    let (tx, rx) = mpsc::channel();
    let tx_clone = tx.clone();
    thread::spawn(move || {
        let vals = vec![
            "hi".to_string(),
            "from".to_string(),
            "the".to_string(),
            "thread".to_string(),
        ];

        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    thread::spawn(move || {
        let vals = vec![
            "...".to_string(),
            "~~~".to_string(),
            "^^^".to_string(),
            ",,,".to_string(),
        ];

        for val in vals {
            tx_clone.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    for received in rx {
        println!("Got: {received}");
    }
}

#[allow(dead_code)]
fn thread4() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];
    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();
            *num += 1;
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }
}