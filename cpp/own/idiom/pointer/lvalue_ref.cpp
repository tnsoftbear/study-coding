#include <iostream>

using std::cout;
using std::endl;

int& foo(int& a) { return a; }
const int& cfoo(int& a) { return a; }
int ifoo(int& a) { return a; }

int main() {
    int a = 10;
    int b = 20;
    int c = a * b;         // ok, c is lvalue, a * b is rvalue
    // a * b = 42;         // error: lvalue required as left operand of assignment
    foo(a) = 42;
    // cfoo(a) = 42;       // error: assignment of read-only location ‘cfoo(a)’
    cout << "a: " << a << endl;

    int* p1 = &foo(a);
    const int* p2 = &cfoo(a);
    // int* p3 = &ifoo(a); // error: lvalue required as unary ‘&’ operand
    cout << "p1: " << p1 << "; p2: " << p2 << "; &a: " << &a << endl;

    int* p4 = &a;
    int a4 = *(p4 + 1);        // выражение (p4 + 1) можно разыменовывать, но нельзя взять его адрес
    // int** pp4 = &(p4 + 1);  // error: lvalue required as unary ‘&’ operand

    int x5 = 1;
    int &a5 = x5;
    int const &b5 = x5;
    int const &c5 = x5 + 1;
    // int &d5 = x5 + 1; // error: cannot bind non-const lvalue reference of type ‘int&’ to an rvalue of type ‘int’
    cout << "a5: " << x5 << "; a6: " << a5 << "; a7: " << b5 << "; a8: " << c5 << endl;
    cout << "x5: " << x5 << "; a5: " << a5 << "; b5: " << b5 << "; c5: " << c5 << endl;
    // x5: 1; a5: 1; b5: 1; c5: 2
    x5 = a5 + b5 + c5;
    cout << "x5: " << x5 << "; a5: " << a5 << "; b5: " << b5 << "; c5: " << c5 << endl;
    // x5: 4; a5: 4; b5: 4; c5: 2
    // Почему c5 = 2, а не 5? Вроде бы ссылка c5 связана со значением x5 + 1. 
    // На самоме деле ссылка с5 была связана с временным объектом и автоматически продлила время его жизни.
    // Важно, что для левых ссылок так работают только константные левые ссылки
    // и получившийся временный объект является неизменяемым без специальных хаков.
}