#include <iostream>

namespace N {
    int x;
}

using namespace N;

int x; // Causes error: reference to 'x' is ambiguous

int main() {
    std::cin >> x;
    std::cout << x + 5;
}