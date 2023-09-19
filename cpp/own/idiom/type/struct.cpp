#include <string>
#include <iostream>

struct S {
    int x;
    double d;
    std::string s;
};

int main() {
    S s{3, 3.14, "Hello"};
    std::cout << s.s << std::endl;
}