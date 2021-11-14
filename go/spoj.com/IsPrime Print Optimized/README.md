# PRINT - Prime Intervals

## My solution

Time: 0.95
Mem: 50M

It is based on Segmentation Sieve algorithm, see [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes).  
IO reads and writes are optimized with help buffering provided by "bufio" library.  

Use for testing the next min, max ranges:

2146483647 2147483647  
9999000000 10000000000  
999900000 1000000000  
999800000 999900000  
999700000 999800000  

## Task Description

In this problem you have to print all primes from given interval.

### Input

t - the number of test cases, then t lines follows. [t <= 150]  
On each line are written two integers L and U separated by a blank.  
L - lower bound of interval,  
U - upper bound of interval.  
[2 <= L < U <= 2147483647]  
[U-L <= 1000000].

### Output

For each test case output must contain all primes from interval [L; U] in increasing order.

### Example

```sh
Input:

2
2 10
3 7
Output:

2
3
5
7
3
5
7
```
