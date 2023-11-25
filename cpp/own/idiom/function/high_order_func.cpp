// g++ high_order_func.cpp && ./a.out

#include <iostream>
#include <functional>
#include <numeric>

using std::cout;
using std::endl;

template <class T>
struct sum_and_power {
    constexpr T operator()(T const & accum, T const & el) const {
        return accum + el * el;
    }
};

int main() {
    std::vector<std::string> names{"foo", "bar", "baz", "qux"};
    std::sort(names.begin(), names.end(), std::greater<std::string>{});
    cout << "sorted names: ";
    for (auto&& name : names) {
        cout << name << " ";
    }
    cout << endl;

    std::vector<int> numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    auto const product = std::accumulate(numbers.begin(), numbers.end(), 1, std::multiplies<int>{});
    cout << "product of numbers: " << product << endl;

    int sp = std::accumulate(numbers.begin(), numbers.end(), 0, sum_and_power<int>{});
}

/**

Особенности функциональных объектов:
В отличие от указателей на функции, хранят состояние внутри себя. Поэтому отпадает необходимость передавать состояние через нетепизированный указатель void*.
Состояние может быть как неизменяемым, так и изменяемым. Соответственно, operator() может быть константным или нет.
В функциях стандартной библиотеки передаются по значению. Следовательно, необходим конструктор копирования.
Рекомендуется делать легковесными.
Не могут преобразовываться к типу указателя на функцию. Поэтому несовместимы с традиционным для языка С механизмом, если не прибегнуть к следующему ухищрению.

Адаптер функционального объекта превращает функциональный объект в функцию, принимающую состояние через нетипизированный указатель.
Функция приводит указатель к типу функционального объекта и применяет его к остальным аргументами.

template <typename F, typename... A>
decltype(std::declval<F>()(std::declval<A>()...))
fptr_adaptor(A... args, void *p) {
    auto &f = *static_cast<F*>(p);
    return f(args...);
}

Такой адаптер может помочь, если необходимо взаимодействовать коду на языке С++ с библиотекой на С.

*/