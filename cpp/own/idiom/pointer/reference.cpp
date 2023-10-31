#include <iostream>
#include <vector>

void swap(int& a, int& b) {
  int t = a;
  a = b;
  b = t;
}

int& staticRef() {
  // int x = 1;  // warning: reference to local variable ‘x’ returned, тк приводит к висячей ссылке (dangling)
  int static x = 1;
  return x;
}

int foo(int& x) { return x; }

int main() {
  using std::vector;
  vector<int> v = {0, 1, 2, 3, 4, 5};
  vector<int> vv = v;   // создать копию вектора v
  vector<int>& vvv = v; // ссылка на тот же объект v
  vv[0] = 100;          // нет эффекта на v
  vvv[1] = 11;          // есть эффект на v
  swap(v[2], v[3]);
  int& vref = v[4];
  int* vptr = &v[4];
  vref += 10; // теперь v[4] = 14
  vptr += 1;  // теперь vptr указывает на v[5] = 5
  size_t index = 0;
  for (const auto& value : v) {
    printf("v[%lu]=%d, ", index, value);
    index++;
  }
  printf("*vptr: %d\n", *vptr);

  int z = staticRef();
  printf("ref to static var: %d\n", z);

  int x = 100, y = 200;
  int& xref = x; // ссылка может быть связана лишь однажды, её больше нельзя привязать к y
  xref = y;      // то же что x=y, теперь х=200
  xref += 50;    // теперь x=250
  // Ссылки прозрачны для операций, включая взятие адреса.
  int* xptr = &xref; // тоже самое, что &x
  // Ссылки не имеют адреса. Нельзя сделать указатель на ссылку.
  // int&* xrefptr = &xref; // error: cannot declare pointer to ‘int&’
  int*& xptrref = xptr; // ок, ссылка на указатель
  printf("x: %d, y: %d, xref: %d, xptr: %d, xptrref: %d\n", x, y, xref, *xptr, *xptrref);

  // Неконстантные левые ссылки не создают временных объектов и просто отказываются связываться с литералами
  // foo(1); // error: cannot bind non-const lvalue reference of type ‘int&’ to an rvalue of type ‘int’
  // int& x1 = 1; // error: cannot bind non-const lvalue reference of type ‘int&’ to an rvalue of type ‘int’
}