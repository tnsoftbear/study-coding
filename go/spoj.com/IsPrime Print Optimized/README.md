# PRINT - Prime Intervals

## My solution

[Task: PRINT - Prime Intervals](https://www.spoj.com/problems/PRINT/)

`segmented-sieve-optimized.go`

Time: 0.71
Mem: 15M

It is based on Segmentation Sieve algorithm, see [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes).  
IO reads and writes are optimized with help buffering provided by "bufio" library.  

Use for testing the next min, max ranges:

2146483647 2147483647  
999000000 1000000000  
998000000 9999000000  
997000000 9998000000  

## Simplified solution

`segmented-sieve-optimized-4.go` - ispired by [C++ v14 solution](https://github.com/rajonaust/SPOJ/blob/master/PRINT%20-%20Prime%20Intervals)

Time: 0.47
Mem: 14M

Run with inputs: `go run segmented-sieve-optimized-4.go < input.txt > primes.txt`

## Experiments

### OOP solution

This is OOP style solution for `segmented-sieve-optimized.go`. I want to check overhead of decoupling logic to objects.

Time: 0.89
Mem: 22M

### Predefined primes

`segmented-sieve-optimized-3.go` - I prebuilt primes and put them into array. Soultion isn't accepted because exceed acceptable program size

### Parallel calculations

`segmented-sieve-parallel.go` is parallel version of `segmented-sieve-optimized-4.go`

Time: 0.58, 0.61, 0.80
Mem: 168M, 178M, 180M

## Other solutions

* [C++ v14, Time: 0.30, Mem: 5.3M](http://spoj-solutions.blogspot.com/2015/01/prime1-prime-generatorprint-prime.html)

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
