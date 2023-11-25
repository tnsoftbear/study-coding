/**
  Адаптер функционального объекта превращает функциональный объект в функцию, принимающую состояние через нетипизированный указатель.
  Функция приводит указатель к типу функционального объекта и применяет его к остальным аргументами.
  Такой адаптер может помочь, если необходимо взаимодействовать коду на языке С++ с библиотекой на С.
*/

#include <iostream>
#include <locale>
#include <string>
#include <utility>

template <typename F, typename... A> decltype(std::declval<F>()(std::declval<A>()...)) fptr_adaptor(A... args, void* p) {
  auto& f = *static_cast<F*>(p);
  return f(args...);
}

// Функциональный объект для адаптера
class pcstring_cmp {
  std::locale m_locale;

public:
  pcstring_cmp(std::string const& name)
      : m_locale(name) {}
  // принимает 2 нетипизированных указателя
  int operator()(void const* p, void const* q) const {
    // приводит их к указателю на указатель на строку
    auto const& s = **static_cast<std::string* const*>(p);
    auto const& t = **static_cast<std::string* const*>(q);
    // сравнивает строки и возвращает -1, 1, 0 в зависимости от их лексигрофического порядка
    return m_locale(s, t) ? -1 : m_locale(t, s) ? 1 : 0;
  }
};

int main() {
  std::string* items[] = {new std::string("c"), new std::string("b"), new std::string("a")};
  pcstring_cmp comparer{"en_US.UTF-8"};
  int count = sizeof(items) / sizeof(items[0]);
  // Приминение функции qsort_r из библиотеки С к массиву указателей на С++ строки std::string
  qsort_r(items, count, sizeof(items[0]), &fptr_adaptor<pcstring_cmp, void const*, void const*>, &comparer);
  for (int i = 0; i < count; ++i) {
    std::cout << *items[i] << std::endl;
  }
}
