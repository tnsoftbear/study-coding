#include <iostream>

const int& f(const int& a) { return a; }
char* f2(char* a) { return a; }
const int f3() { return 42; }

int main() {
  // --- Указатели на функцию ---
  const int& (*pf)(const int&) = f; // rf - указатель на функцию f
  int r1 = pf(1);
  r1 += 100; // возвращаемое значение можно менять, потому что объявлено как int (отрезали константность)
  auto r11 = pf(1);
  r11 += 200; // auto тоже режет константность
  const int r111 = pf(1) + 300;
  // r111 += 400;                   // error: assignment of read-only variable ‘r111'

  char* (*pf2)(char*) = f2; // rf2 - указатель на функцию f2
  char* r2 = pf2((char*)"hi");

  const int (*pf3)() = f3;
  int r3 = pf3();
  r3 += 100; // auto режет константность

  printf("r1: %d, r2: %s, r3: %d\n", r1, r2, r3);
}