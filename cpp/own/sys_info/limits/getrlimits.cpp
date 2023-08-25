#include <iostream>
#include <map>
#include <sys/resource.h>
// #include <limits>
#include <cstdint>

/**
 * This example shows how to get the current resource limits.
 * https://man7.org/linux/man-pages/man2/getrlimit.2.html
 */

int main() {
    // uint64_t maxValue = std::numeric_limits<uint64_t>::max();
    // Создаем мапу с именами ресурсов и их идентификаторами
    std::map<int, std::string> resourceNames = {
        {RLIMIT_CPU, "RLIMIT_CPU - Per-process CPU limit, in seconds"},
        {RLIMIT_FSIZE, "RLIMIT_FSIZE - Largest file that can be created, in bytes"},
        {RLIMIT_DATA, "RLIMIT_DATA - Maximum size of data segment, in bytes"},
        {RLIMIT_STACK, "RLIMIT_STACK - Maximum size of stack segment, in bytes"},
        {RLIMIT_CORE, "RLIMIT_CORE - Largest core file that can be created, in bytes"},
        {RLIMIT_RSS, "RLIMIT_RSS - Largest resident set size, in bytes"},
        {RLIMIT_NPROC, "RLIMIT_NPROC - Number of processes"},
        {RLIMIT_NOFILE, "RLIMIT_NOFILE - Number of open files"},
        {RLIMIT_MEMLOCK, "RLIMIT_MEMLOCK - Locked-in-memory address space"},
        {RLIMIT_AS, "RLIMIT_AS - Address space limit"},
        {RLIMIT_LOCKS, "RLIMIT_LOCKS - Maximum number of file locks"},
        {RLIMIT_SIGPENDING, "RLIMIT_SIGPENDING - Maximum number of pending signals"},
        {RLIMIT_MSGQUEUE, "RLIMIT_MSGQUEUE - Maximum bytes in POSIX message queues"},
        {RLIMIT_NICE, "RLIMIT_NICE - Maximum nice priority allowed to raise to"},
        {RLIMIT_RTPRIO, "RLIMIT_RTPRIO - Maximum realtime priority allowed for non-priviledged processes"},
        {RLIMIT_RTTIME, "RLIMIT_RTTIME - Maximum CPU time in microseconds that a process scheduled under a real-time"},
    };

    // Выводим информацию о лимитах для всех ресурсов
    for (const auto& entry : resourceNames) {
        int resource = entry.first;
        const std::string& resourceName = entry.second;

        struct rlimit limit;
        if (getrlimit(resource, &limit) == 0) {
            std::cout << "Resource: " << resourceName << std::endl;
            std::cout << "  Soft Limit: " << (limit.rlim_cur == RLIM_INFINITY ? "Unlimited" : std::to_string(limit.rlim_cur)) << std::endl;
            std::cout << "  Hard Limit: " << (limit.rlim_max == RLIM_INFINITY ? "Unlimited" : std::to_string(limit.rlim_max)) << std::endl;
        } else {
            std::cerr << "Failed to get resource limit for: " << resourceName << std::endl;
        }
    }

    return 0;
}
