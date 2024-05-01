mod deref;
mod borrow2;

fn main() {
    float1();
}

#[allow(dead_code)]
fn float1() {
    let x: f64 = 0.1;
    let y: f64 = 0.2;
    let z: f64 = 0.3;
    if (x + y - z).abs() < f64::EPSILON {
        println!("Yes");
    } else {
        println!("No");
    }
}