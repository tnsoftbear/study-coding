#include <iostream>

template <int N>
struct Fibonacci {
    static const unsigned int value = Fibonacci<N-2>::value + Fibonacci<N-1>::value;
};

template <>
struct Fibonacci<1> {
    static const unsigned int value = 1;
};

template <>
struct Fibonacci<0> {
    static const unsigned int value = 0;
};

template <unsigned int N>
const unsigned int Fib = Fibonacci<N>::value;

int main() {
    // std::cout << Fibonacci<20>::value << std::endl;
    std::cout << Fib<50> << std::endl;
}