#include <iostream>

int main() {
    int x1;
    double y = 1.0;
    x1 = static_cast<int>(y);
    const int *p = &x1;
    int *q = const_cast<int*>(p);
    long long uq = *reinterpret_cast<long long*>(q);
    std::cout << "uq: " << uq << std::endl; // 1
}