#include <iostream>

int main() {
    int arr[100];
    for (int i = 0; i < 100000; i++) {
        //arr[i] = i;
        std::cout << i << ":\t\t" << arr[i] << std::endl;
    }
}