#include <atomic>

int main() {
    int n = 3;
    std::atomic<int> **a;
    a = new std::atomic<int>* [n];
    for(int i = 0; i < n; i++)
    {
        a[i] = new std::atomic<int>(-1);
    }
}