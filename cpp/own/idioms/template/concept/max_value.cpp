// g++ -std=c++20 max_value.cpp && ./a.out

#include <iostream>
using std::cout;
using std::endl;

template <typename T> requires(!std::is_pointer_v<T>) T maxValue(T a, T b) {
  return b < a ? a : b;
}

int main() {
  int x = 42;
  int y = 77;
  cout << "maxValue(x, y): " << maxValue(x, y) << endl;
  // int max = maxValue(&x, &y);   // error: no matching function for call to
  // ‘maxValue(int*, int*)’
}
