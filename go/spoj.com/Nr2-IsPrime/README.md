# PRIME1 - Prime Generator

 #number-theory

* `simple-sieve.go` - not fast enough
* `segmented-sieve.go` - accepted solution

## Links

* [Task: PRIME1 - Prime Generator](https://www.spoj.com/problems/PRIME1/)
* [Primality test](https://en.wikipedia.org/wiki/Primality_test)
* [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)
* [Волшебное решето Эратосфена](https://habr.com/ru/post/91112/)

## Task

Peter wants to generate some prime numbers for his cryptosystem. Help him! Your task is to generate all prime numbers between two given numbers!

Input
The input begins with the number t of test cases in a single line (t<=10). In each of the next t lines there are two numbers m and n (1 <= m <= n <= 1000000000, n-m<=100000) separated by a space.

Output
For every test case print all prime numbers p such that m <= p <= n, one number per line, test cases separated by an empty line.

Example

```sh
Input:
2
1 10
3 5

Output:
2
3
5
7

3
5
```

Warning: large Input/Output data, be careful with certain languages (though most should be OK if the algorithm is well designed)