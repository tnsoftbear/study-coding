#include <iostream>

/**
 * Препод утверждает, что это unspecified behavior - неспецифированное поведение, 
 * потому что стандарт не определяет порядок в котором должны быть вычеслено выражение f() + g().
 * Т.е. что выведут эти ф-ции: 12 или 21
 */

int f() {
    std::cout << 1;
    return 2;
}

int g() {
    std::cout << 2;
    return 3;
}

int main() {
    std::cout << f() + g() << std::endl;
    return 0;
}