#include <iostream>

using std::cout;
using std::endl;

void f(int&) {
    cout << "f(int&)" << endl;
}

void f(int&&) {
    cout << "f(int&&)" << endl;
}

int main() {
    int a = 1;
    // int &&b = a; // error: cannot bind rvalue reference of type ‘int&&’ to lvalue of type ‘int’
    int&& b = a + 1;
    b += 1; // ok
    // int&& c = c + 1; // ok, но значение рандомное, полагаю это UB, приводит к Segmentation Fault
    // cout << "a: " << a << "; b: " << b << "; c: " << c << endl;

    int&& d = 4;
    int& e = d; // ok, потому что d - это lvalue (идентификатор), lvalue reference инициализируется по lvalue
    // int&& g = d; // error: cannot bind rvalue reference of type ‘int&&’ to lvalue of type ‘int’
    cout << "d: " << d << "; e: " << e << "; &d: " << &d << "; &e: " << &e << endl;

    const int&& g = d + 1;
    // const int&& h = d; // error: cannot bind rvalue reference of type ‘const int&&’ to lvalue of type ‘int’
    const int&& h = (const int&&)d;
    cout << "g: " << g << "; h: " << h << "; g: " << &g << "; h: " << &h << endl;

    f(a);
    f(b);
    f(e);
    // The parameter cast to an rvalue-reference to allow moving it.
    f(std::move(a));
    f(std::move(b));
    f(std::move(e));
}