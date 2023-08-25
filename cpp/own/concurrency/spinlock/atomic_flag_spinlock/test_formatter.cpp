#include <format>
#include <iostream>
#include <thread>
 
int main()
{
    std::thread::id this_id = std::this_thread::get_id();
    std::thread::id null_id;
 
    std::cout << std::format("current thread id: {}\n", this_id);
    std::cout << std::format("{:=^10}\n", null_id);
}