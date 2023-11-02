#include <iostream>
class Class {
  static double static_double;
  char c;
};

double Class::static_double = 0.0;

int main() {
  std::cout << sizeof(Class) << std::endl; // 1
}

// Статические поля не принадлежат объектам, а принадлежат классу в целом. 
// И будут храниться в bss или data секции, т.е. не на стеке/куче где хранится объект.