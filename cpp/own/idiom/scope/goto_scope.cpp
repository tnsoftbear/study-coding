#include <iostream>

using std::cout;

class A {
public:
  A() { cout << "Ctor "; }
  ~A() { cout << "Dtor "; }
};

int i = 1;

int main() {
label:
  A a;
  if (i--) {
    goto label;
  }
}

// Ctor Dtor Ctor Dtor