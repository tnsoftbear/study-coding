#include <atomic>
#include <cstdint>
#include <twist/ed/stdlike/atomic.hpp>

struct SplitCount {
  int32_t transient{0};
  int32_t strong{0};  // > 0
};

int main() {
    auto sc1 = SplitCount{1, 1};
    auto sc2 = SplitCount{2, 2};
    std::atomic<SplitCount> asc1{sc1};
    asc1.compare_exchange_strong(sc1, sc2);
}