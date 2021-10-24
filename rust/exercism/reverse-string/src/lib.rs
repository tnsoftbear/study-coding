// https://exercism.io/my/solutions/370ab77caaff44d2836ff8b98c90ec44
pub fn reverse(input: &str) -> String {
    // let string = input.to_string(); // now its type is String
    let string = String::from(input);
    let mut result = String::new();
    for c in string.chars().rev() {
        result.push(c);
    }
    return result;
}

// pub fn reverse(data: &str) -> String {
//     data.chars().rev().collect()
// }