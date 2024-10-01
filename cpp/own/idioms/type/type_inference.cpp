#include <iostream>

int main() {
    const int& cx1 = 10;
    decltype(cx1) cx2 = 20; // const int &
    decltype(cx1 + 1) cx3 = 30; // int
    std::cout << "cx2: " << cx2 << ", cx3: " << cx3 << std::endl;
}