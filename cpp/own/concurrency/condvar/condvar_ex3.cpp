#include <iostream>
#include <string>
#include <thread>
#include <mutex>
#include <condition_variable>
 
std::mutex m;
std::condition_variable condVar;
std::string data;
bool ready = false;
bool processed = false;
 
void workerThread()
{
    // Wait until main() sends data
    std::unique_lock lk(m);
    condVar.wait(lk, []{return ready;});
 
    // after the wait, we own the lock.
    std::cout << "Worker thread is processing data\n";
    data += " after processing";
 
    // Send data back to main()
    processed = true;
    std::cout << "Worker thread signals data processing completed\n";
 
    // Manual unlocking is done before notifying, to avoid waking up
    // the waiting thread only to block again (see notify_one for details)
    lk.unlock();
    condVar.notify_one();
}
 
int main()
{
    std::thread worker(workerThread);
 
    data = "Example data";
    // send data to the worker thread
    {
        std::lock_guard lk(m);
        ready = true;
        std::cout << "main() signals data ready for processing\n";
    }
    condVar.notify_one();
 
    // wait for the worker
    {
        std::unique_lock lk(m);
        condVar.wait(lk, [] {return processed;});
    }
    std::cout << "Back in main(), data = " << data << '\n';
 
    worker.join();
}