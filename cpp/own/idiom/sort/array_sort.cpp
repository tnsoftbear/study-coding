#include <iostream>
#include <algorithm>

bool cmp(int a, int b) {
    if (a % 2 == 0) {
        return a > b;
    }
    return a < b;
}

int main() {
    // Сортировка
    int s[5] = {3, 6, 4, 2, 5};
    std::sort(s, s+5);
    for (int i = 0; i < 5; i++) std::cout << s[i] << " ";
    std::cout << std::endl;
    // Сортировка с компаратором
    std::sort(s, s+5, cmp);
    for (int i = 0; i < 5; i++) std::cout << s[i] << " ";
    std::cout << std::endl;
}