// https://leetcode.com/problems/two-sum/

use std::collections::HashMap;

struct Solution;

impl Solution {

    // 1. Brute force (15ms)

    // pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    //     for index in 0..nums.len()-1 {
    //         let checking_num = nums[index];
    //         for second_index in index+1..nums.len() {
    //             if checking_num + nums[second_index] == target {
    //                 return vec![index as i32, second_index as i32];
    //             }
    //         }
    //     }
    //     vec![0]
    // }

    // 2. One-pass Hash Table (2ms)

    // pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    //     let mut num_indices: HashMap<i32, usize> = HashMap::new();
    //     for (index, &num) in nums.iter().enumerate() {
    //         let complement = target - num;
    //         let result = num_indices.get(&complement);
    //         if let Some(&prev_index) = result {
    //             return vec![prev_index as i32, index as i32];
    //         }
    //         num_indices.insert(num, index);
    //     }
    //     vec![]
    // }


    // 3. One-pass Hash Table optimized (0 - 1ms)

    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut num_indices: HashMap<i32, i32> = HashMap::new();
        for index in 0..nums.len() {
            let complement = target - nums[index];
            let data = num_indices.get_key_value(&complement);
            if data.is_some() {
                return vec![index as i32, *(data.unwrap().1)];
            }
            num_indices.insert(nums[index], index as i32);
        }
    
        vec![]
    }
}

fn main() {
    let nums = vec![0, 1, 2, 7, 11, 15];
    let target = 9;
    // let nums = vec![3, 2, 3, 3];
    // let target = 6;

    let result = Solution::two_sum(nums, target);
    println!("Индексы двух чисел: {:?}", result);
}
 