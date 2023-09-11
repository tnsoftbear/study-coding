#include <thread>
#include <vector>
#include <mutex>
#include <iostream>

class ForwardList {
    struct Node {
        Node* next;
    };
  public:
    void Push() {
        Node* new_node = new Node{head_};
        head_ = new_node;
    }

    void Pop() {
        Node* old_head = head_;
        head_ = head_->next;
        delete old_head;
    }
  private:
    Node* head_{nullptr};
};

int main() {
    ForwardList list;
    std::vector<std::thread> threads;
    std::mutex mutex;

    for (size_t i = 0; i < 5; ++i) {
        threads.emplace_back([&list, &mutex]() {
            for (size_t k = 0; k < 100500; ++k) {
                // Comment out the line and compile with address sanitizer: clang++ -fsanitize=address forward_list.cpp && ./a.out
                std::lock_guard<std::mutex> guard(mutex);
                list.Push();
                list.Pop();
            }
        });
    }

    for (auto& t : threads) {
        t.join();
    }
}