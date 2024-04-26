struct MyStruct<'a, 'b: 'a> { // lifetime 'b must be at least same long as lifetime 'a
    some_data: Vec<i32>,
    some_ref_data: &'a Vec<i32>,
    some_ref_data2: &'b Vec<i32>,
}

fn main() {
    let a: Vec<i32> = vec![1, 2, 3, 4, 5];
    let b: Vec<i32> = vec![6, 7, 8, 9];
    let c = get_vec_slice(&a, &b);
    println!("{:?}", c);

    let fl1 = 1.1;
    let fl2 = 2.2;
    println!("Smaller: {}", get_smaller(&fl1, &fl2));
}

fn get_vec_slice<'a>(param_1: &'a [i32], param_2: &'a [i32]) -> &'a [i32] {
    if param_1.len() > param_2.len() {
        &param_1[0..2]
    } else {
        &param_2[0..2]
    }
}

fn get_smaller<'a, T: std::cmp::PartialOrd>(param_1: &'a T, param_2: &'a T) -> &'a T {
    if param_1 < param_2 {
        param_1
    } else {
        param_2
    }
}

#[allow(dead_code)]
fn test_1(param_1: Vec<f64>) -> Vec<f64> {  // Lifetimes don't apply because there are no reference inputs or output
    param_1
}

#[allow(dead_code)]
fn test_2(param_1: &Vec<f64>) -> Vec<f64> {  // Lifetimes aren't an issue because there is no reference output
// Compiler sees lifetimes this way: fn test_2<'a>(param_1: &'a Vec<f64>) -> Vec<f64> {
    param_1.clone()
}

#[allow(dead_code)]
// fn test_3<'a>(param_1: Vec<f64>) -> &'a Vec<f64> {  // Lifetimes don't apply because there are no reference inputs
//     // &param_1    // error[E0515]: cannot return reference to function parameter `param_1`
// }

#[allow(dead_code)]
fn test_4<'a>(param_1: i32, param_2: &'a str, param_3: &'a str, param_4: f64) -> &'a str {
    #[allow(unused_parens)]
    if (param_1 == 7 && param_4 > 10.) {
        param_2
    } else {
        param_3
    }
}