#include <cassert>
#include <iostream>

using std::cout;
using std::endl;

int f(int&) {
  cout << "f(int&)" << endl;
  return 1;
}

int f(int&&) {
  cout << "f(int&&)" << endl;
  return 2;
}

// --- Провисание ссылок ---
// Провиснуть (dangle) могут все типы ссылок
// const int& clref (int p) { return p + 0; }      // p + 0 is dead. Возвращаемому типу const int& нечего продливать, потому что "коробочка" (p + 0) умирает вместе с фреймом ф-ции
// const int& clref2 (int p) { return p; }         // p rests in peace
// int& lref (int p) { return p; }                 // p is p no more
// int&& rref (int p) { return p + 0; }            // p + 0 has expired and gone
// int&& rvref (int &&p) { return p + 0; }         // p + 0 is an ex-p + 0
// --- Когда не провисает ссылка ---
int& lvref(int& a) { return a; }
const int& clvref(const int& a) { return a; }

int main() {
  int a = 1;
  // int &&b = a; // error: cannot bind rvalue reference of type ‘int&&’ to lvalue of type ‘int’
  int&& b = a + 1;
  b += 1; // ok
  // int&& c = c + 1; // ok, но значение рандомное, полагаю это UB, приводит к Segmentation Fault
  // cout << "a: " << a << "; b: " << b << "; c: " << c << endl;

  int&& d = 4;
  int& e = d; // ok, потому что d - это lvalue (идентификатор), lvalue reference инициализируется по lvalue
  // int&& g = d; // error: cannot bind rvalue reference of type ‘int&&’ to lvalue of type ‘int’
  cout << "d: " << d << "; e: " << e << "; &d: " << &d << "; &e: " << &e << endl;

  const int&& g = d + 1;
  // const int&& h = d; // error: cannot bind rvalue reference of type ‘const int&&’ to lvalue of type ‘int’
  const int&& h = (const int&&)d;
  cout << "g: " << g << "; h: " << h << "; g: " << &g << "; h: " << &h << endl;

  f(a); // f(int&)
  f(b); // f(int&)
  f(e); // f(int&)

  // The parameter cast to an rvalue-reference to allow moving it.
  f(std::move(a)); // f(int&&)
  f(std::move(b)); // f(int&&)
  f(std::move(e));

  f(a + 1); // f(int&&)
  f(b + 1); // f(int&&)
  f(f(a));  // f(int&) f(int&&)

  // --- Провисание ссылок ---
  // int i = clref(2);              // warning: returning reference to temporary [-Wreturn-local-addr]
  // cout << "i: " << i << endl;    // значение i не определено (Segmentation fault)
  // const int& j = clref(3);       // значение j не определено
  // cout << "j: " << j << endl;    // Segmentation fault
  // const int& k = clref2(4);      // значение k не определено
  // cout << "k: " << k << endl;    // Segmentation fault
  // int& l = lref(5);
  // cout << "l: " << l << endl;    // Segmentation fault
  // int&& m = rref(6);
  // cout << "m: " << m << endl;    // Segmentation fault
  // int&& n = rvref(7);
  // cout << "n: " << n << endl;    // Segmentation fault
  // --- Не провисают ---
  int& o = lvref(a);
  cout << "o: " << o << endl;
  const int& p = clvref(8);
  cout << "p: " << p << endl;

  // --- Анонимизация переменных ---
  // Чтобы получить rvalue ref на значение существующего объекта, нужно использовать std::move
  int q = 9;
  int&& r = std::move(q); // q и r ссылкаются на одно и то же место со значением 9
  // Значение не убирается из q при std::move(), потому q - это int - built-in тип.
  q += 2;
  assert(q == 11);
  assert(r == 11);
  assert(r == q);
  r = 0;
  assert(q == 0);
  assert(r == q);

  // --- Методы над rvalues ---
  struct S {
    int n = 5;
    int& access() { return n; }
    int foo() & { return 10; };  // метод с lvalue-квалификатором
    int foo() && { return 20; }; // метод с rvalue-квалификатором
  };
  S s1;
  int& s2 = s1.access();
  // Метод может быть вызван для rvalue-expression
  int& s3 = S{}.access(); // Получили провисшую ссылку s3
  std::cout << "s3: " << s3 << ", s1.foo(): " << s1.foo() << ", S{}.foo(): " << S{}.foo() << std::endl;
}