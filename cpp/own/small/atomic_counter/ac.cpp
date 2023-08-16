#include <iostream>
#include <atomic>
#include <thread>
#include <vector>

#include "helpers.hpp"

class AtomicCounter {
    public:
      void Increment() {
        value_.fetch_add(1);
      }

      size_t Get() {
        return value_.load();
      }

    private:
      std::atomic<size_t> value_{0};
};

void Stress() {
    AtomicCounter counter;
    std::vector<std::thread> threads;
    auto stop_watch = StopWatch();
    for (size_t i = 0; i < 2; ++i) {
        threads.emplace_back([&counter]() {
            for (size_t j = 0; j < 1'000'000; ++j) {
                counter.Increment();
            }
        });
    }
    std::cout << "Atomic counter: " << counter.Get() << "; Elapsed: " << stop_watch.ElapsedMillis() << "ms" << std::endl;
}

int main() {
  while (true) {
    Stress();
  }
}