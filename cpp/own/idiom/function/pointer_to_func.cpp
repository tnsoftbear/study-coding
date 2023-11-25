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

  char* (*pf21)(char*) = &f2; // pf21 - объявление и определение переменной-указателя на функцию f2
  char* r21 = (*pf21)((char*)"hi"); // косвенный вызов ф-ции по указателю
  // Тоже самое с использованием упрощенного синтаксиса
  char* (*pf22)(char*); // объявление переменной-указателя на функцию
  pf22 = f2; // присвоение указателя на функцию f2
  char* r22 = pf22((char*)"hi"); // косвенный вызов ф-ции по указателю

  const int (*pf3)() = f3;
  int r3 = pf3();
  r3 += 100; // auto режет константность

  printf("r1: %d, r21: %s, r22: %s, r3: %d\n", r1, r21, r22, r3);

  // -- Ссылки и cdecl --
  char* (*arrayOfFunctionPointers[10])(int*&);
  // Ссылка на массив из 10 указателей на функцию принимающую ссылку на указатель и возвращающую указатель на char.
  char* (*(&c)[10])(int*& p) = arrayOfFunctionPointers;
}