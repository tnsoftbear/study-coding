# Prime numbers detection

Implemented with help of segmented sieve algorithm.

## PRINT - Prime Intervals

In this problem you have to print all primes from given interval.

### Input

`t` - the number of test cases, then `t` lines follows. `[t <= 150]`  
On each line are written two integers `L` and `U` separated by a blank.  
`L` - lower bound of interval, `U` - upper bound of interval. `[2 <= L < U <= 2147483647]` `[U-L <= 1000000]`.

### Output

For each test case output must contain all primes from interval `[L; U]` in increasing order.

### Example

#### Input

```sh
2
2 10
3 7
```

#### Output

```sh
2
3
5
7
3
5
7
```