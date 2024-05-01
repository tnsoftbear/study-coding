fn main() {
    // borrow1();
    borrow2();
}

// #[allow(dead_code)]
// #[allow(unused_variables)]
// fn borrow1() {
//     // However, doing let b2 = b and then println is not undefined behavior.
//     // Although b is moved, its data is not deallocated until move_a_box is called at the end.
//     // Therefore this program is technically safe, although still rejected by Rust.
//     let b = Box::new(0);
//     let b2 = b; // Fix: let b2 = b.clone();
//     println!("{}", b);  // error[E0382]: borrow of moved value: `b`
//     borrow1_move_box(b2);
// }
//
// #[allow(unused_variables)]
// fn borrow1_move_box(b: Box<i32>) {
//     //
// }

// #[allow(dead_code)]
// fn borrow2() {
//     let mut s = String::from("hello");
//     let s2 = &s;
//     let s3 = &mut s;    // error[E0502]: cannot borrow `s` as mutable because it is also borrowed as immutable
//     s3.push_str(" world");
//     println!("{s2}");   // Fix: Передвинуть наверх
// }