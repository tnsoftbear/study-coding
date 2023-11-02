#include <iostream>

int main() {
  bool b = double(1 / 2);
  bool c = double(1.0 / 2);
  std::cout << "b[" << b << "] + c[" << c << "] = " << (b + c) << std::endl;
  // b[0] + c[1] = 1
}