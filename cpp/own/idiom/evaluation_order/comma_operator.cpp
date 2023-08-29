#include <iostream>

int main() {
    using std::cout;
    int a = 5, b = 6, c;
    c = a + b, ++a, ++b; // comma operator guarantes evaluation order (left first, then right)
    cout << "1] a: " << a << "; b: " << b << "; c: " << c << std::endl;
    c = (a + b, ++a, ++b); // increment b and assign b to c
    cout << "2] a: " << a << "; b: " << b << "; c: " << c << std::endl;
    (b = 2, a = b) = c; // assign c to a
    cout << "3] a: " << a << "; b: " << b << "; c: " << c << std::endl;
}