#include <iostream>

template <typename T, typename U>
void f(T, U) {
    std::cout << 1 << std::endl;
}

// Если специализация определена до шаблонного метода (f(T, T)), то она будет пропущена. Вывод 1 2
// Однако, если специализация всё же находится в этом месте, мы можем указать желание вызвать её явно: f<int, int>(1, 2);
// template <>
// void f(int, int) {
//     std::cout << 3 << std::endl;
// }

template <typename T>
void f(T, T) {
    std::cout << 2 << std::endl;
}

// Будет использована эта специализация. Вывод 1 3
// Сначала компилятор решает какая версия перегрузки побеждает (т.е. void f(T, T) vs void f(T, U)), 
// а потом уже принимать ли во внимание специализацию (принимать, касаемо f(1, 2)).
template <>
void f(int, int) {
    std::cout << 3 << std::endl;
}

int main() {
    f(1, 3.14);
    f(1, 2);
}