// https://leetcode.com/problems/palindrome-number/

struct Solution;

impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        let s = x.to_string();
        for i in 0..(s.len() as f64 / 2.0).floor() as usize {
            if let Some(c) = s.chars().nth(i) {
                if let Some(c2) = s.chars().nth(s.len()-i-1) {
                    if c != c2 {
                        return false
                    }
                }
            }
        }
        return true
    }
}

fn main() {
    let input = -1021;
    let res = Solution::is_palindrome(input);
    println!("{}", res);
}
