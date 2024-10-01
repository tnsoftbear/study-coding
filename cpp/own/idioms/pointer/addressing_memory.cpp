#include <iostream>

/**
 *
 * [ data | text | stack ]
 * data - static memory
 * text - code
 * stack - automatic memory (`ulimit -s` for stack size)
 *
 * dynamic memory
 */

int a = 0; // data

void f(int a) {
  int b; // stack
  std::cout << "Stack b: " << &b << std::endl;
  static int x = 0; // data
  std::cout << "Static x: " << &x << "; Value x: " << x << std::endl;
  ++x;
}

void d() {
  int* p = new int(5); // dynamic memory
  std::cout << "Dynamic p: " << p << "; Value *p: " << *p << std::endl;
  delete p;
  std::cout << "Dynamic p: " << p << "; Value *p: " << *p << std::endl; // UB
                                                                        // delete p; // UB/Error: "free(): double free detected in tcache 2"
}

int main() {
  int a; // stack
  std::cout << &a;
  {
    double x = 10.0; // stack
    std::cout << &x << std::endl;
  }

  std::cout << *(double*)(&a + 1) << std::endl;   // UB
  std::cout << "Stack a: " << &a << std::endl;    // Near b in stack
  std::cout << "Global a: " << &::a << std::endl; // Near x in static data

  f(a);
  f(a);

  d();

  int *p1, p2; // здесь p2 имеет тип int, а не указатель на int. Для 2х указателей надо так: int *p1, *p2;
  p1 = &a;
  p2 = a;
}
