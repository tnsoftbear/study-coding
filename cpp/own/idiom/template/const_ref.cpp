#include <iostream>

template <typename T>
void f(T& z) {
    std::cout << "f(T& z)" << std::endl;
}

template <typename T>
void f(const T& z) {
    std::cout << "f(const T& z)" << std::endl;
}

int main() {
    f(1); // f(const T& z)
    int x = 0;
    int& y = x;
    f(y); // f(T& z)
    const int& z = x;
    f(z); // f(const T& z)
}