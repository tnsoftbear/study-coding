#include <iostream>

int main() {
    int x = 3;
    // Static cast (корректность проверяется на этапе компиляции)
    double d1 = static_cast<double>(x); // C++-style static cast

    // Reinterpret cast
    double d2 = *reinterpret_cast<double*>(&x); // UB: int и double типы имеют разный размер
    std::cout << d2 << std::endl;

    double d3 = 3.14;
    std::cout << std::hex << reinterpret_cast<int*>(&d3) << std::endl;

    // Const cast (обход защиты константности)
    {
        const int cx = 1;
        int& x = const_cast<int&>(cx);  // UB
        x = 2;
        std::cout << cx << std::endl;
    }

    // С-style cast пробует все касты подряд (const_cast -> static_cast -> reinterpret_cast)
    // Не использовать!
    double d4 = (double)x;              // С-style static cast

    // Dynamic cast
}
