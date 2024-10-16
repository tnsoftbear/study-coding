// RVO demo
// compile with: g++ rvo_ex1.cpp && ./a.out
// compile with: g++ -fno-elide-constructors rvo_ex1.cpp && ./a.out

#include <iostream>

using std::cout;
using std::endl;

struct foo {
  foo() { cout << "foo::foo()" << endl; }
  foo(const foo&) { cout << "foo::foo(const foo&)" << endl; }
  ~foo() { cout << "foo::~foo()" << endl; }
};

// Без NRVO здесь создастся объект foo, а потом вызовется конструктор копирования для инициализации f1
foo nrvo() {
  foo local_foo;
  return local_foo;
}

foo rvo() { return foo(); }

int main() {
  cout << "rvo:" << endl;
  // для этого примера нет разницы как компилировать - с флагом или без, потому что происходит "mandatory rvalue elision"
  foo f1 = rvo();
  cout << "nrvo:" << endl;
  foo f2 = nrvo(); // копи-конструирование пропускается благодаря NRVO (если без флага -fno-elide-constructors)
}

/**
Флаг -fno-elide-constructors в компиляторе g++ используется для отключения оптимизации, называемой "конструкторами устранения".
Эта оптимизация, также известная как Return Value Optimization (RVO), позволяет компилятору оптимизировать процесс возврата значений из функций,
особенно когда значения возвращаются через конструктор копирования.

Когда флаг -fno-elide-constructors включен, это означает, что компилятор не будет применять оптимизацию RVO,
и будут вызываться конструкторы копирования или перемещения для объектов, возвращаемых из функций.
Это может привести к дополнительным вызовам конструкторов и копирования данных.

Иногда отключение оптимизации RVO может быть полезно для отладки или для более точного контроля над тем,
как объекты копируются и перемещаются при возврате из функций.
Однако в большинстве случаев оптимизация RVO улучшает производительность и снижает издержки на копирование объектов,
поэтому по умолчанию эта оптимизация включена в компиляторе g++.


"Mandatory Rvalue Elision" (обязательное опущение rvalue) - это оптимизация, предоставляемая стандартом C++11 
и более поздними версиями, которая позволяет компилятору опустить создание и разрушение временных объектов (rvalue), 
что приводит к улучшению производительности и снижению накладных расходов.

Суть этой оптимизации заключается в том, что компилятор может опустить конструктор копирования 
и деструктор для временных объектов, которые создаются и разрушаются в рамках выражения,
и это сделано обязательно согласно стандарту, если выполнены определенные условия.
Это позволяет избежать избыточных операций копирования и разрушения временных объектов,
что может быть очень полезным для оптимизации кода.
*/