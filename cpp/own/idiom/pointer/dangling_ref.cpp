#include <iostream>

// Временные объекты живут на стеке функции, как стек закончился, временный объект тоже.
// Оба варианта приводят к UB.

const int& f1(const int& a) { return a; }

int& f2() {
  int x = 42;
  return x; // warning: reference to local variable ‘x’ returned
}

// В этом примере всё ок. Константные ссылки безопасно брать, но небезопасно возвращаеть
int f3(const int& a) { return a; }

int& g1() {
  int* p = new int(1);
  return *p;
}

int main() {
  int r1, r2, r3;
  r1 = f1(0); // UB: dangling reference
  // r2 = f2();  // UB, dangling reference, приводит к Segmentation fault (core dumped)
  r3 = f3(3);

  printf("r1: %d, r2: %d, r3: %d\n", r1, r2, r3);

  int& h1 = g1();
  delete &h1; // ok
  // int h11 = g1();  // Здесь мы скопировали *p полученное из g1(), это значение лежит на стеке main(), а не в динамической памяти выделенной через new int(1).
  // delete &h11;     // munmap_chunk(): invalid pointer.
}
