#include <iostream>
#include <thread>
#include <atomic>

inline void CompilerBarrier() {
    asm volatile("" ::: "memory");
}

void StoreBuffering(size_t iter) {
    std::atomic<int> x{0};
    std::atomic<int> y{0};

    int r1 = 0, r2 = 0;

    std::thread t1([&]() {
        x = 1;
        CompilerBarrier();
        r1 = y;
    });

    std::thread t2([&]() {
        y = 1;
        CompilerBarrier();
        r2 = x;
    });

    t1.join();
    t2.join();

    if (r1 == 0 && r2 == 0) {
        std::cout << "Broken on Iteration #" << iter << std::endl;
        std::abort();
    }
}

int main() {
    for (int i = 0; i < 100500; i++) {
        StoreBuffering(i);
    }
}