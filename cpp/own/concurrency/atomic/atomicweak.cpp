#include <iostream>
#include <memory>
#include <thread>
using namespace std::literals;

int main() {
  std::atomic<std::weak_ptr<int>> pShared;
  std::atomic<bool> done{false};
  std::jthread updates{[&] {
    for (int i = 0; i < 10; ++i) {
      {
        auto sp = std::make_shared<int>(i);
        pShared.store(sp);
        std::this_thread::sleep_for(0.1s);
      }
      std::this_thread::sleep_for(0.1s);
    }
    done.store(true);
  }};

  while (!done.load()) {
    if (auto sp = pShared.load().lock()) {
        std::cout << "shared: " << *sp << '\n';
    } else {
        std::cout << "shared: <no data>\n";
    }
    std::this_thread::sleep_for(0.07s);
  }
}
