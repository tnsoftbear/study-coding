#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>

std::mutex mtx;
std::condition_variable cv;
bool isReady = false;

void worker_thread() {
    std::this_thread::sleep_for(std::chrono::seconds(2));
    {
        std::lock_guard<std::mutex> lock(mtx);
        isReady = true;
    }
    cv.notify_one();
}

int main() {
    std::thread t(worker_thread);

    std::unique_lock<std::mutex> lock(mtx);
    cv.wait(lock, [] { return isReady; });  // Ожидаем, пока не будет isReady == true

    std::cout << "Worker thread is ready!" << std::endl;

    t.join();

    return 0;
}