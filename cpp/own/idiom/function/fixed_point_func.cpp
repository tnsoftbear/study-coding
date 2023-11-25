#include <functional>
#include <algorithm>
#include <iostream>

template <typename R, typename A>
std::function<R(A)> fixed(std::function<R(std::function<R(A)>, A)> h)
{
    return [h](A x) { return h(fixed(h), x); };
}

int main() {
    std::vector<int> xs = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    std::vector<int> ys{};
    std::transform(
        xs.cbegin(),
        xs.cend(),
        std::back_inserter(ys),
        fixed<double, int>(
            [](std::function<double(int)> f, int k) { return k == 0 ? 1.0 : k * f(k-1); }
        )
    );
    for (auto y : ys) {
        std::cout << y << ' ';
    }
}

/**
  Универсальный движок рекурсии. Инкапсулирует в себе рекурсивность как таковую.
  Пользователю достаточно изобрести лишь нерекурсивную функцию h.
 */