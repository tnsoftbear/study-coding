struct Solution;

impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let mut map = std::collections::HashMap::new();
        map.insert('I', 1);
        map.insert('V', 5);
        map.insert('X', 10);
        map.insert('L', 50);
        map.insert('C', 100);
        map.insert('D', 500);
        map.insert('M', 1000);

        let mut ans = 0;

        let chars: Vec<char> = s.chars().collect();

        for i in 0..chars.len() {
            if i + 1 < chars.len() && map[&chars[i]] < map[&chars[i + 1]] {
                ans -= map[&chars[i]];
            } else {
                ans += map[&chars[i]];
            }
        }
        ans
    }
}

fn main() {
    let str = "MCMXCIV".to_string();
    let int = Solution::roman_to_int(str);
    println!("int: {int}");
}
