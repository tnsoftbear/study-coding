fn move_box(b: Box<i32>) {
    //
}

fn main() {
    // However, doing let b2 = b and then println is not undefined behavior.
    // Although b is moved, its data is not deallocated until move_a_box is called at the end.
    // Therefore this program is technically safe, although still rejected by Rust.
    let b = Box::new(0);
    let b2 = b;
    println!("{}", b);
    move_box(b2);
}