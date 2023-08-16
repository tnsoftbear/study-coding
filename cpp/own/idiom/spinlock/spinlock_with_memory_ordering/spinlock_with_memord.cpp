#include <iostream>
#include <chrono>
#include <atomic>
#include <vector>
#include <thread>

// g++ spinlock_with_memord.cpp -pthread

/**
 * Relax in favour of the CPU owning the lock
 * https://c9x.me/x86/html/file_module_x86_id_232.html
 */
inline void SpinLockPause() {
    asm volatile("pause\n" : : : "memory");
}

class SpinLockInterface {
  public:
    virtual void Lock() = 0;
    virtual void Unlock() = 0;
};

class SpinLockNaive : public SpinLockInterface {
  public:
    void Lock() override {
        while (locked_.exchange(true)) {    // <-- cache ping-pong
            SpinLockPause();
        }
    }

    void Unlock() override {
        locked_.store(false);
    }
  private:
    std::atomic<size_t> locked_{false};
};

class SpinLockOptimized : public SpinLockInterface {
  public:
    void Lock() override {
        while (locked_.exchange(true, std::memory_order_acquire)) {    // <-- cache ping-pong
            while (locked_.load()) {        // mitigate cache ping-pong (не вижу разницы, проверить на большем CPU)
                SpinLockPause();
            }
        }
    }

    void Unlock() override {
        locked_.store(false, std::memory_order_release);
    }
  private:
    std::atomic<size_t> locked_{false};
};

void Stress(SpinLockInterface& spinlock) {
    size_t counter = 0; // Guarded by spinlock
    std::vector<std::thread> threads;

    for (size_t i = 0; i < 10050; ++i) {
        threads.emplace_back([&]() {
            spinlock.Lock();
            ++counter;
            spinlock.Unlock();
        });
    }

    for (auto& t : threads) {
        t.join();
    }

    // std::cout << "Counter: " << counter << std::endl;
}

int main() {
    SpinLockOptimized spinlockOpt;
    SpinLockNaive spinlockNaive;
    auto start = std::chrono::steady_clock::now();
    auto elapsed = std::chrono::steady_clock::now() - start;
    while (true) {
        start = std::chrono::steady_clock::now();
        Stress(spinlockNaive);
        elapsed = std::chrono::steady_clock::now() - start;
        std::cout << "Elapsed for Naive spinlock: "
            << std::chrono::duration_cast<std::chrono::milliseconds>(elapsed).count()
            << "ms" << std::endl;

        start = std::chrono::steady_clock::now();
        Stress(spinlockOpt);
        elapsed = std::chrono::steady_clock::now() - start;
        std::cout << "Elapsed for Optimized spinlock: "
            << std::chrono::duration_cast<std::chrono::milliseconds>(elapsed).count()
            << "ms" << std::endl;
    }
}