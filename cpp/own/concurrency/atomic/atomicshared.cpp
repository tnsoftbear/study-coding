#include <iostream>
#include <memory>
#include <thread>
#include <vector>
using namespace std::literals;

template <typename T> class AtomicList {
private:
  struct Node {
    T val;
    std::shared_ptr<Node> next;
  };
  std::atomic<std::shared_ptr<Node>> head;

public:
  AtomicList() = default;
  void insert(T v) {
    auto p = std::make_shared<Node>();
    p->val = v;
    p->next = head;
    while (!head.compare_exchange_weak(p->next, p)) {
    }
  }
  void print() const {
    std::cout << "HEAD";
    for (auto p = head.load(); p; p = p->next) {
      std::cout << "->" << p->val;
    }
    std::cout << std::endl;
  }
};

int main() {
  AtomicList<std::string> alist;
  {
    std::vector<std::jthread> threads;
    for (int i = 0; i < 100; ++i) {
      threads.push_back(std::jthread{[&, i] {
        for (auto s : {"hi", "hey", "ho", "last"}) {
          alist.insert(std::to_string(i) + s);
          std::this_thread::sleep_for(5ns);
        }
      }});
    }
  }
  alist.print();
}
