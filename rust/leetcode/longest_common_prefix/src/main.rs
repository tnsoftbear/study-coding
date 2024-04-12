// https://leetcode.com/problems/longest-common-prefix

struct Solution;

impl Solution {
    pub fn longest_common_prefix(mut strs: Vec<String>) -> String {
        strs.sort_by(|a, b| a.len().cmp(&b.len()));
        if let Some(smallest) = strs.get(0) {
            let smallest_len = smallest.len();
            for j in 1..=smallest_len {
                if let Some(prefix) = smallest.get(0..j) {
                    for i in 1..strs.len() {
                        if !strs[i].starts_with(prefix) {
                           if let Some(result) = prefix.get(0..prefix.len()-1) {
                               return result.to_string()
                           }
                           return String::new()
                        }
                    }
                }
            }
            return smallest.to_string()
        }
        return String::new()
    }
}

fn main() {
    let strs: Vec<String> = ["flowers", "flow", "florida"]
        .iter()
        .map(|&s| s.to_string())
        .collect();
    let prefix = Solution::longest_common_prefix(strs);
    println!("Prefix: {prefix}");
}
