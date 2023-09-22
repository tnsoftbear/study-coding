#include <iostream>
#include <mutex>
#include <condition_variable>
#include <thread>
#include <ctime>
#include <sstream>

std::mutex mtx;
std::condition_variable condVar;
bool isReady = false;
int counter = 0;
bool areWorkersFinished = false;

std::string makePrefix() {
    std::time_t currentTime = std::time(nullptr);
    char timeString[100];
    std::strftime(timeString, sizeof(timeString), "[%H:%M:%S] ", std::localtime(&currentTime));
    std::thread::id this_id = std::this_thread::get_id();

    std::ostringstream threadIdStream;
    threadIdStream << std::this_thread::get_id();
    std::string threadId = threadIdStream.str();

    return timeString + threadId + ", c: " + std::to_string(counter) + " ";
}

void workerThread() {
    while (counter < 5) {
        ++counter;
        std::this_thread::sleep_for(std::chrono::seconds(1));
        {
            std::cout << makePrefix() << "Before lock(mtx)" << std::endl;
             // Здесь захватывается мьютекс, второй поток будет заблокирован
            std::unique_lock<std::mutex> lock(mtx);
            std::cout << makePrefix() << "After lock(mtx)" << std::endl;
            std::this_thread::sleep_for(std::chrono::seconds(1));
            std::cout << makePrefix() << "Before cv.wait" << std::endl;
            // Здесь сначала мьютекс освобождается тем самым второй поток продолжается на After lock(mtx),
            // в первом потоке проверяется условие предиката wait(), и т.к. оно false, то мьютекс захватывается вновь первым потоком.
            condVar.wait(lock, [] { return isReady; });
            std::cout << makePrefix() << "After cv.wait" << std::endl;
        }
    }
}

void controlThread() {
    while (!areWorkersFinished) {
        std::this_thread::sleep_for(std::chrono::seconds(4));
        isReady = true;
        condVar.notify_one();
        std::cout << makePrefix() <<  "isReady = true" << std::endl;
        std::this_thread::sleep_for(std::chrono::seconds(4));
        isReady = false;
        std::cout << makePrefix() <<  "isReady = false" << std::endl;
    }
}

int main() {
    std::thread t1(workerThread);
    std::thread t2(workerThread);
    std::thread ctl(controlThread);
    std::cout << makePrefix() <<  "All threads created" << std::endl;

    t1.join();
    std::cout << makePrefix() <<  "1st worker thread joined" << std::endl;
    t2.join();
    std::cout << makePrefix() <<  "2st worker thread joined" << std::endl;
    areWorkersFinished = true;

    ctl.join();
    std::cout << makePrefix() <<  "Control thread joined" << std::endl;

    return 0;
}