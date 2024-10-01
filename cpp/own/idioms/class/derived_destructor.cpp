#include <iostream>

using std::cout;
using std::endl;

/**
 * Обычный деструктор удаляет объект по типу указателя, 
 * виртуальный удаляет объект с учётом типа самого объекта.
 * Все как с обычными виртуальными функциями.
 */

class Base {
public:
  // Так только ~Base будет вызван
  ~Base() { cout << "Base dtor" << endl; };
  // Так будут вызваны ~Derived(), потом ~Base()
  // virtual ~Base() { cout << "Base virtual dtor" << endl; };
};

class Derived : public Base {
public:
  ~Derived() /* override */ { cout << "Derived dtor" << endl; };
};

int main() {
  Base* b = new Derived;
  delete b;
  return 0;
}