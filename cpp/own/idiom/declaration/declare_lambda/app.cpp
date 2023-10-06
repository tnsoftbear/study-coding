#include "lib.cpp"
#include <iostream>

int main() {
    foo([]() { std::cout << "Hello, world!" << std::endl; });
}
