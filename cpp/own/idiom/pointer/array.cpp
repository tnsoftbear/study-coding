#include <iostream>
#include <vector>
#include <algorithm>

int main()  {
    // Создать массив на стеке
    int a[5] = {1, 2, 3, 4, 5};
    // void f(std::vector<double> v);
    // f(a); // error: could not convert ‘(int*)(& a)’ from ‘int*’ to ‘std::vector<double>’
    // f(&a); // error: could not convert ‘& a’ from ‘int (*)[10]’ to ‘std::vector<double>’

    int* p0 = a; // Array to pointer conversion
    int c[30] = {1, 2, 3, 4};
    int* p2 = a + 2;
    // *p0: 1; *(p0+2): 3; p2[-1]: 2; 3[a]: 4
    std::cout << "*p0: " << *p0 << "; *(p0+2): " << *(p0+2)
         << "; p2[-1]: " << p2[-1] << "; 3[a]: " << 3[a] << std::endl;

    // int b[5] = a; // error: array must be initialized with a brace-enclosed initializer
    // int b[5]; a = b; // error: invalid array assignment

    // sizeof(a): 20; sizeof(p0): 8; sizeof(c): 120
    std::cout << "sizeof(a): " << sizeof(a) << "; sizeof(p0): " << sizeof(p0) << "; sizeof(c): " << sizeof(c) << std::endl;

    int b[5];
    std::cout << b - a << std::endl;
    std::cout << b << std::endl;
    std::cout << a << std::endl;

    int size = 100;
    int* pa = new int[size]; // Выделение динамической памяти под массив
    std::cout << "sizeof(pa): " << sizeof(pa) << "; sizeof(*pa): " << sizeof(*pa) << "; size: " << sizeof(int) * size << std::endl;
    delete[] pa;            // [] указывает, что необходимо удалить массив заданной ранее длинны.

    int* pb = new int(100); // Указатель на int со значением 100, который аллоцирован в динамической памяти.
    // int* pb = new int{100}; // Аналогично
    std::cout << "sizeof(pb): " << sizeof(pb) << "; sizeof(*pb): " << sizeof(*pb) << std::endl;
    delete pb;

    // Двумерный динамический массив так лучше не делать, а использовать std::vector<int>.
    int h, w = 10;
    int** ca = new int*[h];
    for (int i = 0; i < h; i++) ca[i] = new int[w];
}
