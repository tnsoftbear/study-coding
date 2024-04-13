// https://leetcode.com/problems/valid-parentheses/
// Самое оптимальное решение
// Runtime: 0ms, Beats: 100.00% of users with Rust, Memory: 2.00MB, Beats: 99.90% of users with Rust
// https://leetcode.com/problems/valid-parentheses/submissions/1230570421/

struct Solution;

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack = Vec::new();
        let bracket_pairs: std::collections::HashMap<char, char> = [ // Это массив кортежей
            ('(', ')'),
            ('{', '}'),
            ('[', ']')
        ]
            .iter() // создать итератор, который позволяет нам перебирать элементы массива
            .cloned() // хэш-мапа должна владеть своими значениями, иначе они остались бы связанными с массивом, который может быть уничтожен после создания хэш-мапы
            .collect(); // создания хэш-мапы из итератора, который содержит клонированные кортежи с парами скобок

        for c in s.chars() {
            match c {
                '(' | '{' | '[' => stack.push(c),
                ')' | '}' | ']' => {
                    if let Some(last) = stack.pop() {
                        // Ф-ция гет принимает ссылку на ключ k: & и возврщает ссылку на значение &V: pub fn get<Q: ?Sized>(&self, k: &Q) -> Option<&V>
                        if let Some(&expected) = bracket_pairs.get(&last) {
                            if c != expected {
                                return false;
                            }
                        } else {
                            return false;
                        }
                    } else {
                        return false;
                    }
                }
                _ => {} // ничего не выполнять в запасном случае
            }
        }

        stack.is_empty()
    }
}

fn main() {
    let s = String::from("{[]}[]()");
    let is = Solution::is_valid(s);
    println!("{is}");
}
