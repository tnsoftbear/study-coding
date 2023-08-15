#include <iostream>

/**
 * g++ -O2 --std=c++17 re_segfault2.cpp
 * Оптимизация приводит к бесконечному циклу.
 * Компилятор считает, что ввод коректен и что программа не содержит необпределенного поведения.
 * Компилятор видит умножение на 12345678 и делает вывод, что i не превышает 173.
 * Т.к. в цикле for есть условие i < 300, то 173 < 300 всегда true.
 * 
re_segfault2.cpp:5:42: warning: iteration 174 invokes undefined behavior [-Waggressive-loop-optimizations]
    5 |         std::cout << i << ":\t\t" << i * 12345678 << std::endl;
      |                                          ^~~~~~~~
re_segfault2.cpp:4:23: note: within this loop
    4 |     for (int i = 0; i < 300; i++) {
      |                     ~~^~~~~
 */

int main() {
    for (int i = 0; i < 300; i++) {
        std::cout << i << ":\t\t" << i * 12345678 << std::endl;
    }
}