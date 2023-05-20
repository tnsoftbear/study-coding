/**
 * Celebrity Problem
 * https://practice.geeksforgeeks.org/problems/the-celebrity-problem/1
 *
 * @param {number[][]} M
 * @param {number} n
 * @returns {number}
 */
class Solution {
  // Function to find if there is a celebrity in the party or not.
  celebrity(M, n) {
    let noFriends = [];
    for (let i = 0; i < n; i++) {
      let hasFriend = false;
      for (let j = 0; j < n; j++) {
        if (M[i][j] === 1) {
          hasFriend = true;
          break;
        }
      }
      if (!hasFriend) {
        noFriends.push(i);
      }
    }

    for (let p = 0; p < noFriends.length; p++) {
      let tryHim = noFriends[p];
      let allKnow = true;
      for (let y = 0; y < n; y++) {
        if (tryHim !== y && M[y][tryHim] !== 1) {
          allKnow = false;
          break;
        }
      }
      if (allKnow) {
        return tryHim;
      }
    }
    return -1;
  }
}

function main() {
  let M = [
    [0, 1, 0],
    [0, 0, 0],
    [0, 1, 0],
  ];
  let n = 3;

  const solution = new Solution();
  console.log(solution.celebrity(M, n));
}

main();
