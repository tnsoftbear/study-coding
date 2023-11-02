#include <iostream>

struct S {
    //S()=default;
    S(const int ii) : i(ii) {}
    int i;
};

struct Accessor {
    Accessor(S& s) /* : s_(s) */  {
        s_ = S(s.i);
    };
    void printi() {
        std::cout << "s.i: " << s_.i << std::endl;
    }

    S s_{0};
    // S s_; // Без инициализации значения требуется конструктор по-умолчанию S()=default;
};


int main() {
  S s = {10}; // S(10); // Str(10);
  auto a = new Accessor(s);
  a->printi();
}