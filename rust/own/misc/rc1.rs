use std::{borrow::Borrow, rc::Rc};

#[derive(Debug)]
pub struct PRef<T> {
    value: T
}

fn main() {
    let pref1 = PRef {
        value: 10
    };
    let pref_rc = Rc::new(pref1);
    // println!("{:?}, value: {}", pref_rc, pref_rc.value);

    let pref2 = create(pref_rc);
}

fn create<T>(pref_rc: Rc<PRef<T>>) -> Option<PRef<T>> {
    let pref1 = PRef {
        value: 10
    };
    Some(...)
}

fn log<T: std::fmt::Debug + std::fmt::Display>(pref: PRef<T>) {
    println!("{:?}, value: {}", pref, pref.value);
}