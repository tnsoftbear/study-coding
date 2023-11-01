#include <iostream>
#include <cstdarg>

/**
 * Посмотрим манглированные имена ф-ций в ассембленом коде:
 * g++ -g0 -O1 -masm=intel -S overload.cpp
    
   c++filt _Z1fd
   > f(double)

   c++filt _Z1fi
   > f(int)

   c++filt _Z1ff
   > f(float)

   c++filt _Z1fPKcz
   _Z1fPKcz

   c++filt _Z1fl
   > f(long)
   
   c++filt _Z1fRl
   > f(long&)
 */

int f(double a) {       // _Z1fd
    printf("int f(double a)\n");
    return (int)a;
}

double f(int a) {       // _Z1fi
    printf("double f(int a)\n");
    return (double)a;
}

float f(float a) {      // _Z1ff
    printf("float f(float a)\n");
    return (float)a;
}

void f(const char* a, ...) {    // _Z1fPKcz
    printf("void f(...) / ");
    va_list args;
    va_start(args, a);
    vprintf(a, args);
    va_end(args);
}

void f(long a) {        // _Z1fl
    printf("void f(long)\n");
}

void f(long& a) {       // _Z1fRl
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

/**

Правила разрешения перегрузки
• Первое приближение (здесь много чего не хватает)
1. Точное совпадение (int → int, int → const int&, etc)
Обратите внимание: nullptr точно совпадает с любым указателем.
2. Точное совпадение с шаблоном (int →T)
3. Стандартные преобразования (int → char, float → unsigned short, etc)
4. Переменное число аргументов
5. Неправильно связанные ссылки (literal → int&, etc)

 */