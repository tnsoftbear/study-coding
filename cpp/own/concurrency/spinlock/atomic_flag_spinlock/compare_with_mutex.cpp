#include <chrono>
#include <mutex>
#include <thread>

std::mutex mut;

void workOnResource() {
    while (true) {
        mut.lock();
        std::this_thread::sleep_for(std::chrono::seconds(2));
        mut.unlock();
    }
}

int main()  {
    std::thread t1(workOnResource);
    std::thread t2(workOnResource);

    t1.join();
    t2.join();
}