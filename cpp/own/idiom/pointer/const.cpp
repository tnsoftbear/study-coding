#include <iostream>
#include <vector>

// Передача аргументов по константной ссылке
size_t find(const std::string& text, const std::string& str) {
  // ...
  return 0;
}

const int& f(const int& a) { return a; }

int main() {
  // Non-const operations: =, op=, ++, --

  const int x = 0;   // константная переменная
  int const cx = 10; // так тоже можно объявлять константную переменную
  // int* p1 = &x;   // error: invalid conversion from ‘const int*’ to ‘int*’

  const int* p = &x;   // указатель на константные данные с типом int (west const)
  int const* cp = &cx; // указатель на константные данные с типом int (east const - тоже самое что west const)
  ++p;                 // сам указатель можно менять
  --cp;
  // *p = 1;              // но не значение error: "assignment of read-only location ‘* p’"
  // *cp = 1;             // аналогичная ошибка: "assignment of read-only location ‘* cp’"

  int y = 1;
  int* const pp = &y; // константный указатель на изменяемые данные
  *pp = 2;            // значение можно менять
  // ++pp;            // error: increment of read-only variable ‘pp’
  // int* const cpp;  // константый указатель должен быть связан с адресом данных: "error: uninitialized ‘const cpp’"

  const int* const ppp = &y;  // константный указатель на константные данные
  int const* const ppp2 = &y; // тоже самое
  // ++ppp;                   // error: increment of read-only variable ‘ppp’
  // *ppp = 3;                // error: assignment of read-only location ‘*(const int*)ppp’

  // Ссылки сами по себе связываются с данными однажды, их нельзя переназначить, т.е. буквально они всегда являются константными ссылками.
  // Но здесь происходит терминологическая путаница, и "константными ссылками" называют ссылки объявленные с модификатором const.
  // Модификатор const означает, что данные на которые ссылается ссылка менять нельзя.
  // Ссылка без модификатора const именуется неконстантной.
  int& rx = y; // Ссылка на изменяемые данные (неконстантная ссылка)
  ++rx;
  // int& p0 = x;              // error: binding reference of type ‘int&’ to ‘const int’ discards qualifiers
  const int& ry = y; // Константная ссылка на изменяемые данные является read-only ссылкой.
  // int& const rry = y;       // error: ‘const’ qualifiers cannot be applied to ‘int&’
  // const int& const rry = y; // error: ‘const’ qualifiers cannot be applied to ‘const int&’
  // ++ry;                     // error: increment of read-only reference ‘ry’
  ++y;
  printf("ry: %d\n", ry);

  const int& сrx = x; // Cсылка на константу x
  // int& rx = x;              // error: binding reference of type ‘int&’ to ‘const int’ discards qualifiers

  find("abcde", "abc");
  const int& r = 0; // Разрешается инициализировать константную ссылку литералом
  // int& rn = 0;              // error: cannot bind non-const lvalue reference of type ‘int&’ to an rvalue of type ‘int’

  const std::vector<int> v1 = {0, 1, 2, 3, 4};
  {
    const std::vector<int>& v2 = v1;
    // Вектор v1, на который ссылается &v2 не уничтожается при выходе из ОВ

    // Lifetime expansion
    const std::vector<int>& v3 = {1, 2, 3};
    // v3 уничтожается при выходе из ОВ
  }

  int r0 = f(0); // UB: dangling reference
  printf("r0: %d\n", r0);
}