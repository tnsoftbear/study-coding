#include <functional>
#include <iostream>

std::function<double(int)> f = [&f](int k) {
  return k == 0 ? 1.0 : k * f(k-1);
};

int main() {
  std::cout << f(5) << std::endl;
}