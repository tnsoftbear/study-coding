#include <iostream>

template <typename T>
const T& max(const T& a, const T& b) {
    return a > b ? a : b;
}

// Специализация
template <> // Надо ли?
const float& max(const float& a, const float& b) {
    const float& c = (a > b ? a : b) * 100;
    return c;
}

template <typename T, typename U>
T min(const T& a, const U& b) { // произведёт каст U к возвращаемому типу T
    // Возвращаемый тип здесь не "const T&"", потому что (a : b) - это rvalue,
    // т.е. временное значение, поэтому мы не можем его биндить к константной ссылке.
    return a < b ? a : b;
}

template <typename T>
T pi = 3.14;

// template <typename U>
// using minInt = min<U, U>;

int main() {
    std::cout << "max(1, 2) = " << max(1, 2) << std::endl;
    std::cout << "max(1.99, 1.99) = " << max(1.99, 1.99) << std::endl;
    std::cout << "max(1.99f, 1.98f) * 100 = " << max(1.99f, 1.98f) << std::endl;
    std::cout << "max(\"abc\", \"def\") = " << max("def", "abc") << std::endl;
    std::cout << "max(2, 2.99) = " << max<int>(2, 2.99) << std::endl;
    std::cout << "min(1.99, 2) = " << min(1.99, 2) << std::endl;
    std::cout << "min(2, 1.99) = " << min(2, 1.99) << std::endl;
    // std::cout << "minInt(1, 1.99) = " << minInt(1, 2) << std::endl;
    std::cout << "pi<int> = " << pi<int> << "; pi<float> = " << pi<float> << std::endl;
}