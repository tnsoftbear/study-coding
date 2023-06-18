/**
 * Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
 * Hash Map - 1 Pass
 * Time O(N) | Space O(N)
 * https://leetcode.com/problems/two-sum/
 * @param {number[]} nums
 * @param {number} target
 * @returns {number[]}
 */
function twoSum(nums, target) {
  let map = new Map();
  for (let i = 0; i < nums.length; i++) { /* Time O(N) */
    let num = nums[i];
    let complement = target - num;
    if (map.has(complement)) {
      return [map.get(complement), i];
    }
    map.set(num, i);	/* Space O(N) */
  }
}

console.log(twoSum([2, 7, 11, 15], 9)); // [0, 1]
