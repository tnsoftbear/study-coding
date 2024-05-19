mod trait1;
mod trait2;

use std::ops::Deref;

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    type Target = T;
    fn deref(&self) -> &Self::Target { // тоже что и: -> &T
        &self.0
    }
}

fn hello(name: &str) {
    println!("Hi, {name}");
}

fn main() {
    let x = 5;
    let y = MyBox::new(x);
    assert_eq!(5, x);
    assert_eq!(5, *y);  // произошло: *(y.deref())

    let s1 = "hello".to_string();
    let s2 = MyBox::new(&s1);
    assert_eq!("hello".to_string(), s1);
    assert_eq!("hello".to_string(), **s2);
    hello(&s1);
    hello(&s2);

    let str1 = "hello";
    let str2 = MyBox::new(str1);
    assert_eq!("hello".to_string(), str1);
    assert_eq!("hello".to_string(), *str2);
    hello(str1);
    hello(&str2);
}