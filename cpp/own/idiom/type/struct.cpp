#include <string>
#include <iostream>

struct S {
    int x;
    double d;
    std::string s;
    void f(char c) {
        std::cout << "f(char c): " << c << std::endl;
    }
};

int main() {
    int S::* ptr = &S::x; // Указатель на поле x можно объявить до создания объекта
    void (S::* pf)(char) = &S::f;
    S s{3, 3.14, "Hello"};
    std::cout << "s.s: " << s.s << ", s.*ptr: " << s.*ptr << std::endl;
    (s.*pf)('a');

    S* ps = new S{4, 4.14, "Bye"};
    (ps->*pf)('b');
}