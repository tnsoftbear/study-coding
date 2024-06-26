#include <iostream>

/**
 * Двухфазное разрешение имён:
 *
 * Первая фаза: до инстанцирования. Шаблоны проходят общую синтаксическую проверку, а также разрешаются _независимые_ имена.
 * Вторая фаза: во время инстанцирования. Происходит специальная синтаксическая проверка и разрешаются _зависимые_ имена.
 * Зависимое имя - это имя, которое семантически зависит от шаблонного параметра. Шаблонный параметр может быть его типом, 
 * он может участвовать в формировании типа и так далее.
 *
 * Короче:
 * Сначала разрешение имен, потом инстанцирование и во время инстанцирования вторая фаза разрешения имён.
 * Золотое правило: разрешение зависимых имён откладывается до подстановки шаблонного параметра.
 */

template<typename T> void foo(T) { std::cout << "T"; }

struct S {};

// void foo(S) { std::cout << "S"; }

template<typename T> void call_foo(T t, S s) {
    // S s - независимое имя не откладывается, сразу биндится в шаблон foo(T) и выводит "T"
    // Если бы foo(S) определялось выше, тогда бы произошла перегрузка и вывело бы "S"
    foo(s);
    // T t - зависимое имя откладывается до точки инстанцирования call_foo(x, x)
    // В точке инстанцирования уже есть перегрузка foo(S) и оно лучше подходит чем foo(T)
    foo(t);
}

void foo(S) { std::cout << "S"; }

void bar (S x) {
    call_foo(x, x);
}

int main() {
    S x;
    bar(x);
}
