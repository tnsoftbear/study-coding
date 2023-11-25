// g++ -std=c++20 find_primes_by_ranges.cpp && ./a.out

#include <iostream>
#include <ranges>
#include <algorithm>

int main() {
  // range позволяют императивный код преобразовать в декларативный
  auto primes = std::views::iota(1)                                    // бесконечная последовательность начиная с 1
                | std::ranges::views::filter([](auto x) {              // фильтрация по критерию
                    return 2 == std::ranges::count_if(                 // должно быть только два элемента в контейнере
                                    std::views::iota(1, x + 1),        // в диапазоне от 1 до x (включительно)
                                    [x](auto k) { return x % k == 0; } // на которые х делится без остатка
                                );
                  })                    // получим бесконечную последовательность всех простых чисел
                | std::views::take(100) // взять первые 100 элементов
                | std::views::transform([](auto x) { return x * x; }); // возведем каждое число в квадрат
  for (auto x : primes) {
    std::cout << x << " ";
  }
  std::cout << std::endl;
}
