#include <thread>
#include <iostream>
#include <mutex>
#include <condition_variable>

std::mutex mtx;
std::condition_variable condVar;
bool isReady{false};

void waitingForWork() {
    std::cout << "Worker: Waiting for work." << std::endl;
    std::unique_lock<std::mutex> lck(mtx);
    
    condVar.wait(lck, []{ return isReady; });
    // The same as:
    // while (!isReady) {
    //     condVar.wait(lck);
    // }

    std::cout << "Worker: Processing shared data." << std::endl;
}

void setDataReady() {
    std::this_thread::sleep_for(std::chrono::seconds(1));
    {
        std::unique_lock<std::mutex> lock(mtx);
        isReady = true;
    }
    std::cout << "Sender: Data is ready." << std::endl;
    condVar.notify_one();
    std::this_thread::sleep_for(std::chrono::seconds(1));
    std::cout << "Sender: Completed." << std::endl;
}


int main() {
    std::thread t1(waitingForWork);
    std::thread t2(setDataReady);
    t1.join();
    t2.join();
}