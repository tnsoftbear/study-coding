#include <iostream>

// Временные объекты живут на стеке функции, как стек закончился, временный объект тоже.
// Оба варианта приводят к UB.

const int& f1(const int& a) { return a; }

int& f2() {
  int x = 42;
  return x;
}

// В этом примере всё ок. Константные ссылки безопасно брать, но небезопасно возвращаеть
int f3(const int& a) { return a; }

int main() {
  int r1, r2, r3;
  r1 = f1(0); // UB: dangling reference
  // r2 = f2();  // UB, dangling reference, приводит к Segmentation fault (core dumped)
  r3 = f3(3);

  printf("r1: %d, r2: %d, r3: %d\n", r1, r2, r3);
}