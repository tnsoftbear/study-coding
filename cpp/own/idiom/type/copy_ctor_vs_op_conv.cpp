/**
 * Copy-constructor vs. operator conversion
 * Конфликт конструирования f2 в строчке: Foo f2 = b;
 * Мы видим, что в C++17 выигрывает конструктор копирования: Foo(const Bar &),
 * а в C++14 выигрывал оператор преобразования: operator Foo().
 * https://www.youtube.com/watch?v=YkyHCP-IZdk&ab_channel=KonstantinVladimirov

g++ --std=c++17 copy_ctor_vs_op_conv.cpp && ./a.out
> Ctor Bar
> Op Bar -> Foo
> Op Bar -> Foo

g++ --std=c++14 copy_ctor_vs_op_conv.cpp && ./a.out
> Ctor Bar
> Ctor Bar -> Foo
> Op Bar -> Foo

 */

#include <iostream>

struct Foo;
struct Bar;

struct Foo {
  Foo() {}
  Foo(const Bar&) {                                     // (!)
    std::cout << "Ctor Bar -> Foo" << std::endl;
  }
};

struct Bar {
  Bar() { std::cout << "Ctor Bar" << std::endl; }
  Bar(const Foo&) { std::cout << "Ctor Foo -> Bar" << std::endl; }
  operator Foo() {                                      // (!)
    std::cout << "Op Bar -> Foo" << std::endl;
    return Foo{};
  }
};

int main() {
  Bar b;
  Foo f1{b};  // direct-init, ctor
  Foo f2 = b; // copy-init, ctor vs op, op wins
}