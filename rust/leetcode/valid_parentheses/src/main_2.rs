// https://leetcode.com/problems/valid-parentheses/
// Быстрое решение, но больше по памяти чем main_o.rs
// Runtime: 0ms Beats: 100.00% of users with Rust, Memory: 2.15MB BeatsL 39.34% of users with Rust
// https://leetcode.com/problems/valid-parentheses/submissions/1230655718/

struct Solution;

impl Solution {
    pub fn is_valid(s: String) -> bool {
        fn is_valid_recursive(chars: &[char], stack: &mut Vec<char>) -> bool {
            if chars.is_empty() {
                return stack.is_empty();
            }

            match chars[0] {
                '(' | '[' | '{' => {
                    stack.push(chars[0]);
                    is_valid_recursive(&chars[1..], stack)
                }
                ')' => {
                    if let Some(top) = stack.pop() {
                        top == '(' && is_valid_recursive(&chars[1..], stack)
                    } else {
                        false
                    }
                }
                ']' => {
                    if let Some(top) = stack.pop() {
                        top == '[' && is_valid_recursive(&chars[1..], stack)
                    } else {
                        false
                    }
                }
                '}' => {
                    if let Some(top) = stack.pop() {
                        top == '{' && is_valid_recursive(&chars[1..], stack)
                    } else {
                        false
                    }
                }
                _ => false,
            }
        }

        is_valid_recursive(&s.chars().collect::<Vec<_>>(), &mut Vec::new())
    }
}

fn main() {
    // let s = String::from("{[]}[]()");
    let s = String::from("(([]){})");
    let is = Solution::is_valid(s);
    println!("{is}");
}
