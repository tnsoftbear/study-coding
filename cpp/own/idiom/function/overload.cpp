#include <iostream>
#include <cstdarg>

int f(double a) {
    printf("int f(double a)\n");
    return (int)a;
}

double f(int a) {
    printf("double f(int a)\n");
    return (double)a;
}

float f(float a) {
    printf("float f(float a)\n");
    return (float)a;
}

void f(const char* a, ...) {
    printf("void f(...) / ");
    va_list args;
    va_start(args, a);
    vprintf(a, args);
    va_end(args);
}

void f(long a) {
    printf("void f(long)\n");
}

void f(long& a) {
    printf("void f(long&)\n");
}

int main() {
    double x = f(0.0);          // Output: int f(double a)
    int y = f(1);               // Output: double f(int a)
    float z = f(2.0f);          // Output: float f(float a)
    f("Nums: %d %d \n", 3, 4);  // Output: void f(...) / Nums: 3 4
    void (*pf)(long) = f;       // Указатель на функцию с long аргументом - это адрес из области инструкций кода
    pf(5);                      // Output: void f(long)

    // long v = 0;
    // long& vr = v;
    // f(vr);                   // error: call of overloaded ‘f(long int&)’ is ambiguous
}

// https://en.cppreference.com/w/cpp/language/overload_resolution
// 1. exact match
// 2. promotion
// 3. built-in conversion
// 4. user-defined conversion
