# PRINT - Prime Intervals

## My solution

[Task: PRINT - Prime Intervals](https://www.spoj.com/problems/PRINT/)

Time: 0.71
Mem: 15M

It is based on Segmentation Sieve algorithm, see [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes).  
IO reads and writes are optimized with help buffering provided by "bufio" library.  

Use for testing the next min, max ranges:

2146483647 2147483647  
9999000000 10000000000  
999900000 1000000000  
999800000 999900000  
999700000 999800000  

## Simplified solution

`segmented-sieve-optimized-4.go` - ispired by [C++ v14 solution](https://github.com/rajonaust/SPOJ/blob/master/PRINT%20-%20Prime%20Intervals)

Time: 0.47
Mem: 14M

## OOP solution

`segmented-sieve-oop.go` works slower to 10% and reads more memory

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
