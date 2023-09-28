#include <mutex>
#include <thread>
#include <chrono>
#include <cassert>

#include "../futex_mutex/futex_mutex.cpp" // Мой мьютекс в деле

/**
 * Compile: g++ -std=c++20 mutexed.cpp
 *
 * Safe API for mutual exclusion
 *
 * Usage:
 *
 * Mutexed<std::vector<Apple>> apples;
 *
 * {
 *   auto owner_ref = apples->Acquire();
 *   owner_ref->push_back(Apple{});
 * }  // <- release ownership
 *
 */

template<typename T, typename Lock, class Mutex>
struct Accessor {
  Accessor(T& object, Mutex& mutex) : object_(object), lock_(mutex) {}

  T* operator->() {
    return &object_;
  }

  const T* operator->() const {
    return &object_;
  }

  const T& operator*() const {
    return object_;
  }

  T& operator*() {
    return object_;
  }

  T& object_;
  Lock lock_;
};

template <typename T, class Mutex = std::mutex>
class Mutexed {

 public:
  // https://eli.thegreenplace.net/2014/perfect-forwarding-and-universal-references-in-c/
  template <typename... Args>
  explicit Mutexed(Args&&... args)
      : object_(std::forward<Args>(args)...) {
  }

  auto Acquire() {
    return Accessor<T, std::lock_guard<Mutex>, Mutex>{ object_, mutex_ };
  }

 private:
  T object_;
  Mutex mutex_;  // Guards access to object_
};

//////////////////////////////////////////////////////////////////////

class Counter {
  public:
    void Increment() {
        size_t value = value_;
        std::this_thread::sleep_for(std::chrono::seconds(1));
        value_ = value + 1;
    }

    size_t Value() const {
        return value_;
    }

  private:
    size_t value_{0};
};

int main() {
    // Mutexed<Counter> counter;
    Mutexed<Counter, FutexMutex> counter; // Работает так же с моим мьютексом

    std::thread t1([&] {
      counter.Acquire()->Increment();
    });
    std::thread t2([&] {
      counter.Acquire()->Increment();
    });

    t1.join();
    t2.join();

    assert(counter.Acquire()->Value() == 2);
}
