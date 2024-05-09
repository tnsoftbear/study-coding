use std::ops::Deref;

fn main() {
    // deref1();
    // deref2();
    // deref3();
    // deref4();
    // deref5();
    // deref6();
    deref7();
}

#[allow(dead_code)]
fn deref1() {
    let mut x: Box<i32> = Box::new(1);
    let a: i32 = *x;         // *x reads the heap value, so a = 1
    *x += 1;                 // *x on the left-side modifies the heap value, so x points to the value 2

    let r1: &Box<i32> = &x;  // r1 points to x on the stack
    let b: i32 = **r1;       // two dereferences get us to the heap value

    let r2: &i32 = &*x;      // r2 points to the heap value directly
    let c: i32 = *r2;    // so only one dereference is needed to read it
    println!("a is {a}, b is {b}, c is {c}");
}

#[allow(dead_code)]
fn deref2() {
    let x: Box<i32> = Box::new(-1);
    let x_abs1 = i32::abs(*x); // explicit dereference
    let x_abs2 = x.abs();      // implicit dereference
    assert_eq!(x_abs1, x_abs2);

    let r: &Box<i32> = &x;
    let r_abs1 = i32::abs(**r); // explicit dereference (twice)
    let r_abs2 = r.abs();       // implicit dereference (twice)
    assert_eq!(r_abs1, r_abs2);

    let s = String::from("Hello");
    let s_len1 = str::len(&s); // explicit reference
    let s_len2 = s.len();      // implicit reference
    assert_eq!(s_len1, s_len2);
}

#[allow(dead_code)]
fn deref3() {
    let x = Box::new(42);
    let y = Box::new(&x);
    println!("{}, {}, {}, {}", y, *y, **y, ***y); // 42, 42, 42, 42
    // y has the type Box<&Box<i32>>. It is a heap pointer to a stack reference to a heap pointer.
    // Therefore y must be dereferenced three times, once for each layer of indirection.
}

#[allow(dead_code)]
fn deref4() {
    let mut v: Vec<i32> = vec![1, 2, 3];
    let num: &i32 = &v[2];
    println!("Third element is {}", num);
    println!("Third element is {}", *num);
    println!("Third element is {}", &v[2]);
    v[2] = 33;
    println!("Third element is {}", &v[2]);
    v.push(4);
}

#[allow(dead_code)]
#[allow(unused_assignments)]
fn deref5() {
    let mut v: Vec<i32> = vec![1, 2, 3];
    let num: &mut i32 = &mut v[2];
    *num += 1;
    println!("Third element is {}", *num);
    println!("Vector is now {:?}", v);

    let mut num1: &mut i32 = &mut v[2];
    num1 = &mut v[1];   // Можно присвоить другую ссылку
    *num1 += 10;
    println!("Vector is now {:?}", v);
}

#[allow(dead_code)]
fn deref6() {
    let mut s = String::from("Hello");
    let t = &mut s;
    /* here */
    // println!("{}", s);   // cannot borrow `s` as immutable because it is also borrowed as mutable
    t.push_str(" world");
    println!("{}", s);
}

#[allow(dead_code)]
fn deref7() {
    let a = Box::new(42);   // Box<i32>
    let b = &a; // &Box<i32>
    let c: &i32 = a.deref(); // &i32
    let d = *a; // i32
    // let e = *b; // error[E0507]: cannot move out of `*b` which is behind a shared reference
    let f = *b.clone();
    println!("c: {:?}, d: {:?}, e: {:?}", c, d, f);
}
