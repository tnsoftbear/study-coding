// https://leetcode.com/problems/roman-to-integer

// При переходе на вектор стало быстрее

struct Solution;

impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let mut sum = 0;
        let mut skip_next = false;
        let chars: Vec<char> = s.chars().collect();

        for i in 0..chars.len() {
            if skip_next {
                skip_next = false;
                continue;
            }
            let c = chars[i];

            if c == 'I' {
                if i + 1 != s.len() {
                    let c2 = chars[i + 1];
                    if c2 == 'X' {
                        sum += 9;
                        skip_next = true;
                        continue;
                    } else if c2 == 'V' {
                        sum += 4;
                        skip_next = true;
                        continue;
                    }
                }
                sum += 1;
                continue;
            }

            if c == 'X' {
                if i + 1 != s.len() {
                    let c2 = chars[i + 1];
                    if c2 == 'L' {
                        sum += 40;
                        skip_next = true;
                        continue;
                    } else if c2 == 'C' {
                        sum += 90;
                        skip_next = true;
                        continue;
                    }
                }
                sum += 10;
                continue;
            }

            if c == 'C' {
                if i + 1 != s.len() {
                    let c2 = chars[i + 1];
                    if c2 == 'D' {
                        sum += 400;
                        skip_next = true;
                        continue;
                    } else if c2 == 'M' {
                        sum += 900;
                        skip_next = true;
                        continue;
                    }
                }
                sum += 100;
                continue;
            }

            if c == 'L' {
                sum += 50;
                continue;
            }

            if c == 'V' {
                sum += 5;
                continue;
            }

            if c == 'D' {
                sum += 500;
                continue;
            }

            if c == 'M' {
                sum += 1000;
                continue;
            }
        }
        sum
    }
}

fn main() {
    let str = "MCMXCIV".to_string();
    let int = Solution::roman_to_int(str);
    println!("int: {int}");
}
