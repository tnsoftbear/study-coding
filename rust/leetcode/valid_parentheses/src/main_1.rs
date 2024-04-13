// Недоделанное решение

struct Solution;

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let chars: Vec<char> = s.chars().collect();
        for i in 0..chars.len() {
            if chars[i] == '(' {
                if let Some(next) = chars.get(i+1) {
                    if *next != ')' {
                        return false
                    }
                } else {
                    return false
                }
            }
            if chars[i] == '{' {
                if let Some(next) = chars.get(i+1) {
                    if *next != '}' {
                        return false
                    }
                } else {
                    return false
                }
            }
            if chars[i] == '[' {
                if let Some(next) = chars.get(i+1) {
                    if *next != ']' {
                        return false
                    }
                } else {
                    return false
                }
            }
        }
        return true
    }
}

fn main() {
    let s = String::from("{[]}[]()(");
    let is = Solution::is_valid(s);
    println!("{is}");
}
