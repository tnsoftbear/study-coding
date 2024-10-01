#include <iostream>

// https://en.wikipedia.org/wiki/Copy_elision

struct C {
  C() { std::cout << "Default ctor\n"; };
  C(const C&) { std::cout << "A copy was made.\n"; }
};

C f() {
  std::cout << "Now default ctor will be called:\n";
  C c = C();
  std::cout << "Return C object from f()\n";
  return c;
}

int main() {
  std::cout << "Hello World!\n";
  C obj = f(); // Мы не вызываем конструктор копирования здесь (NRVO)
  // Мы можем убрать NRVO при помощи -fno-elide-constructors и конструктор копирования будет вызван.
  // Тогда будет напечатано: "A copy was made.\n"
  // g++ -fno-elide-constructors copy_elision3.cpp && ./a.out
}