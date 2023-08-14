#include <iostream>
#include <string>

class Adder {
public:
    int operator()(int a, int b) {
        return a + b;
    }
};

void divideToZero() {
    int x = 10;
    int y = 0;
    int result = x / y;
    std::cout << "Result: " << result << std::endl;
}

int table[4];
bool existInTable(int v)
{
    return (table[5] == v);
    //for (int i = 0; i <= 4; i++) {
    //    if (table[i] == v) return true;
    //}
    //return false;
}

int main()
{
    std::cout << "Hello World!\n";
    // divideToZero();
    int v1 = 0;
    int v2 = 1;
    std::cout << std::to_string(v1) + " " + (existInTable(v1) ? "exists" : "not exists") << std::endl;
    std::cout << std::to_string(v2) + " " + (existInTable(v2) ? "exists" : "not exists") << std::endl;
    for (int i = 0; i < 6; ++i) {
        printf("%d ", table[i]);
    }
    std::cout << std::endl;

    Adder add;
    std::cout << "Result: " << add(5, 3) << std::endl; // Вызываем объект add как функцию

    std::cout << "Game over!\n";
    system("pause>0");
}
