#include <vector>
#include <iostream>

void swap(int& a, int& b) {
    int t = a;
    a = b;
    b = t;
}

int& staticRef() {
    // int x = 1;  // warning: reference to local variable ‘x’ returned, тк приводит к висячей ссылке (dangling)
    int static x = 1;
    return x;
}

int main() {
    using std::vector;
    vector<int> v = {1, 2, 3, 4, 5};
    vector<int> vv = v;     // создать копию вектора v
    vector<int>& vvv = v;   // ссылка на тот же объект v
    vv[0] = 0;              // нет эффекта
    vvv[1] = 0;             // есть эффект
    swap(v[3], v[4]);
    for (const auto& n: v) {
        printf("%d ", n);
    }
    printf("\n");

    int z = staticRef();
    printf("ref to static var: %d\n", z);
}