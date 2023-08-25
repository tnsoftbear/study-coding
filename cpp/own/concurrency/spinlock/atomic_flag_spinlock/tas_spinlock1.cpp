#include <atomic>
#include <thread>
#include <iostream>
#include <chrono>

class TasSpinLock {
    std::atomic_flag flag = ATOMIC_FLAG_INIT;
public:
    void lock() {
        while(flag.test_and_set(std::memory_order_acquire)) {}
    }

    void unlock() {
        flag.clear(std::memory_order_release);
    }
};

TasSpinLock spin;

void workOnResource() {
    while (true) {
        spin.lock();
        std::cout << "Locked Thread ID: " << std::this_thread::get_id() << std::endl;
        std::this_thread::sleep_for(std::chrono::seconds(2));
        spin.unlock();
        std::cout << "Unlocked Thread ID: " << std::this_thread::get_id() << std::endl;
    }
}

int main() {
    std::thread t1(workOnResource);
    std::thread t2(workOnResource);

    t1.join();
    t2.join();
}