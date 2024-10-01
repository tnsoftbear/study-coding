#include <iostream>

int main() {
    int i = 13;     // 001101
    int j = 60;     // 111100
    i ^= j;         // 110001
    j ^= i;         // 001101
    i ^= j;         // 111100
    std::cout << "i: " << i << "; j: " << j << std::endl; // i: 60; j: 13
}