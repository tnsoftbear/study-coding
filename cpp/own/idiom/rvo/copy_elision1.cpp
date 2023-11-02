#include <iostream>

// https://en.wikipedia.org/wiki/Copy_elision

int n = 0;

struct C {
  explicit C(int) { std::cout << "Default ctor\n"; }
  C(const C&) { ++n; std::cout << "A copy was made.\n"; }  // the copy constructor has a visible side effect
};                                                         // it modifies an object with static storage duration

int main() {
  C c1(42);      // direct-initialization, calls C::C(int)
  C c2 = C(42);  // copy-initialization, calls C::C(const C&)

  std::cout << n << std::endl;  // prints 0 if the copy was elided, 1 otherwise
}

/**
 * Добиться вызова конструктора копирования можно так:
 * g++ --std=c++14 -fno-elide-constructors copy_elision1.cpp && ./a.out
 * Иначе он не вызывается, и сайд-эффекта (++n) не происходит
 * Ни так: g++ --std=c++14 copy_elision1.cpp && ./a.out
 * Ни так: g++ -fno-elide-constructors copy_elision1.cpp && ./a.out
 * Ни так: g++ --std=c++17 -fno-elide-constructors copy_elision1.cpp && ./a.out
 */