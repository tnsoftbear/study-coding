#include <iostream>

template <int N>
void f(int x) {
    std::cout << x + N << std::endl;
}

int main() {
    f<1>(10);
    f<2>(10);
}