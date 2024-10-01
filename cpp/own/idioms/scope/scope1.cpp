#include <iostream>

int x = 1;

namespace N {
    int x = 2;
}

int main() {
    // scope resolution
    using namespace N;
    using std::cout;
    int x = 3;
    cout << "local x: " << x << "; N::x: " << N::x << "; global ::x: " << ::x << std::endl;
    {
        int x = 4;
        cout << "local x: " << x << "; N::x: " << N::x << "; global ::x: " << ::x << std::endl;
    }
}
