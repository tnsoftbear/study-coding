#include <iostream>
#include <vector>

// Передача аргументов по константной ссылке
size_t find(const std::string& text, const std::string& str) {
    // ...
    return 0;
}

const int& f(const int& a) {
    return a;
}

int main() {
    // Non-const operations: =, op=, ++, --

    const int x = 0;
    // int* p1 = &x;    // error: invalid conversion from ‘const int*’ to ‘int*’

    const int* p = &x;  // указатель на константу, 
    ++p;                // сам указатель можно менять
    // *p = 1;          // но не значение error: assignment of read-only location ‘* p’

    int y = 1;
    int* const pp = &y; // константный указатель
    // ++pp;            // error: increment of read-only variable ‘pp’
    *pp = 2;            // значение можно менять

    const int* const ppp = &y;
    // ++ppp;           // error: increment of read-only variable ‘ppp’
    // *ppp = 3;        // error: assignment of read-only location ‘*(const int*)ppp’

    // int& p0 = x;         // error: binding reference of type ‘int&’ to ‘const int’ discards qualifiers
    const int& ry = y;      // Cсылка на константу
    // int& const rry = y;  // error: ‘const’ qualifiers cannot be applied to ‘int&’
    // ++ry;                // error: increment of read-only reference ‘ry’
    ++y;
    printf("ry: %d\n", ry);

    find("abcde", "abc");
    const int& r = 0;       // Разрешается инициализировать константную ссылку литералом
    // int& rn = 0;         // error: cannot bind non-const lvalue reference of type ‘int&’ to an rvalue of type ‘int’

    const std::vector<int> v1 = {0, 1, 2, 3, 4};
    {
        const std::vector<int>& v2 = v1;
        // Вектор v1, на который ссылается &v2 не уничтожается при выходе из ОВ
        
        // Lifetime expansion
        const std::vector<int>& v3 = {1, 2, 3};
        // v3 уничтожается при выходе из ОВ
    }

    int r1 = f(0);          // UB: dangling reference
    printf("r1: %d\n", r1);
}