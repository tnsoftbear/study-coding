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

    // Solution without conversion number to string (3ms, longer)
    // https://leetcode.com/problems/palindrome-number/solutions/3651712/2-method-s-c-java-python-beginner-friendly/
    //
    // pub fn is_palindrome(x: i32) -> bool {
    //     if x < 0 || (x != 0 && x % 10 == 0) {
    //         return false;
    //     }
    //     let mut reversed = 0;
    //     let mut num = x;
    //     while num > reversed {
    //         reversed = reversed * 10 + num % 10;
    //         num /= 10;
    //     }
    //     return num == reversed || num == reversed / 10;
    // }
}

fn main() {
    let input = -1021;
    let res = Solution::is_palindrome(input);
    println!("{}", res);
}
