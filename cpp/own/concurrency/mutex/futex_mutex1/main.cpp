#include <thread>
#include <iostream>
#include <chrono>
#include <mutex>

#include "futex_mutex.cpp"

/**
 * g++ -std=c++20 main.cpp
 */

// typedef std::mutex FutexMutex;

void doJob(FutexMutex& mutex, int workerId) {
    printf("Worker (%d): preparing for work.\n", workerId);
    mutex.lock();

    printf("Worker (%d): working in critical section.\n", workerId);
    std::this_thread::sleep_for(std::chrono::milliseconds(2000));

    mutex.unlock();
    printf("Worker (%d): work is completed.\n", workerId);
}

int main() {
    FutexMutex mutex;
    std::thread t1([&]() { doJob(mutex, 1); });
    std::thread t2([&]() { doJob(mutex, 2); });
    std::thread t3([&]() { doJob(mutex, 3); });
    t1.join();
    t2.join();
    t3.join();
}