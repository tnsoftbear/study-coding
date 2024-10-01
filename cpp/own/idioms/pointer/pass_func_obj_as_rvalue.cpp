#include <iostream>
#include <utility>

template <typename F>
void caller(F&& f) {
    std::cout << "Caller is calling the function." << std::endl;
    f();
}

int main() {
    auto someFunction = []() {
        std::cout << "Some function is called." << std::endl;
    };

    caller(static_cast<decltype(someFunction)>(someFunction));
    caller(someFunction);
    return 0;
}
