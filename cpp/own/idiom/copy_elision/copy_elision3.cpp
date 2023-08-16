#include <iostream>

// https://en.wikipedia.org/wiki/Copy_elision

struct C {
  C() = default;
  C(const C&) { std::cout << "A copy was made.\n"; }
};

C f() {
  return C();
}

int main() {
  std::cout << "Hello World!\n";
  C obj = f();
}