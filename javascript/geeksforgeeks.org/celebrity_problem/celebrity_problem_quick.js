/**
 * O(2n) => O(n)
 * https://www.youtube.com/watch?v=xGvQN_g-JCI
 */
class Solution {
    // Function to find if there is a celebrity in the party or not.
    celebrity(M, n) {
      let l = 0;
      let r = n - 1;
      while (l < r) {
          if (this.knows(M, l, r)) {
              l++;
          } else {
              r--;
          }
      }
  
      for (let i = 0; i < n; i++) {
          if (i !== l && (this.knows(M, l, i) || !this.knows(M, i, l))) {
              return -1;
          }
      }
      return l;
    }
  
    knows(M, a, b) {
      return M[a][b];
    }
  }
  