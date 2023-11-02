#include <iostream>
#include <utility>

int main() {
  std::pair<std::string, int> p1("hello", 42);
  std::pair p2("hello", 42);
  std::cout << "p1.first == p2.first" << " = " << (p1.first == p2.first) << '\n';
  
  // error: no match for ‘operator==’ (operand types are ‘std::pair<std::__cxx11::basic_string<char>, int>’ and ‘std::pair<const char*, int>’)
  // std::cout << "p1 == p2 = " << (p1 == p2) << '\n';
}