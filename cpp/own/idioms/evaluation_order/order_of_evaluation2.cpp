#include <iostream>

int main() {
    int x = 0;
    ++x = x++;
    std::cout << "++x = x++ -> " << x << std::endl;
    
    int y = 0;
    ++y = ++y;
    std::cout << "++y = ++y -> " << y << std::endl; // 2 in <= c++14, 1 in >= c++17
}