#include <iostream>

int x = 1;

int main() {
    int x = 2;
    std::cout << "local x: " << x << "; global x: " << ::x << std::endl;
    {
        int x = 3;
        std::cout << "local x: " << x << "; global x: " << ::x << std::endl;
    }
}