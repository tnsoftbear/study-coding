#include <iostream>

using std::cout;
using std::endl;

struct Base {
  virtual int f() = 0;
  int g();
};

int Base::f() { return 1; }
int Base::g() { return 3; }

struct Derived : Base {
  int f() override;
  int g();
};

int Derived::f() { return 2; }
int Derived::g() { return 4; }

int main() {
  Derived d;
  cout << "d.f(): " << d.f() << ", ((Base&)d).f(): " << ((Base&)d).f() << endl; // 2, 2
  cout << "d.g(): " << d.g() << ", ((Base&)d).g(): " << ((Base&)d).g() << endl; // 4, 3
}

/**
 Несмотря на то, что объект d приводится к типу Base, происходит динамическое связывание (полиморфизм), 
 и вызывается версия функции f(), определенная в производном классе Derived.

 Когда объект типа Derived приводится к ссылке на Base и вызывается виртуальная функция,
 компилятор на самом деле использует тип объекта во время выполнения (тип Derived),
 чтобы определить, какую версию функции вызывать. Это называется полиморфным вызовом функции.
 */