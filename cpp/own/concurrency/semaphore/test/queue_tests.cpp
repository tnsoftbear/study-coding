#include <atomic>
#include <chrono>
#include <random>
#include <string>
#include <thread>
#include <vector>

#include "../src/queue.cpp"
#include "assert_macros.hpp"

using namespace std::chrono_literals;

void testPutThenTake() {
  BoundedBlockingQueue<int> queue{1};
  queue.Put(42);
  ASSERT_EQ(queue.Take(), 42);
}

void testSmallFifo() {
  BoundedBlockingQueue<std::string> queue(2);

  std::thread producer([&queue]() {
    queue.Put("hello");
    queue.Put("world");
    queue.Put("!");
  });

  ASSERT_EQ(queue.Take(), "hello");
  ASSERT_EQ(queue.Take(), "world");
  ASSERT_EQ(queue.Take(), "!");

  producer.join();
}

void testBigFifo() {
  BoundedBlockingQueue<int> queue{3};

  static const int kItems = 1024;

  std::thread producer([&]() {
    for (int i = 0; i < kItems; ++i) {
      queue.Put(i);
    }
    queue.Put(-1); // Poison pill
  });

  // Consumer

  for (int i = 0; i < kItems; ++i) {
    ASSERT_EQ(queue.Take(), i);
  }
  ASSERT_EQ(queue.Take(), -1);

  producer.join();
}

void testCapacity() {
  BoundedBlockingQueue<int> queue{3};
  std::atomic<size_t> send_count{0};

  std::thread producer([&]() {
    for (size_t i = 0; i < 100; ++i) {
      queue.Put(i);
      send_count.store(i);
    }
    queue.Put(-1);
  });

  std::this_thread::sleep_for(100ms);

  ASSERT_TRUE(send_count.load() <= 3);

  for (size_t i = 0; i < 14; ++i) {
    (void)queue.Take();
  }

  std::this_thread::sleep_for(100ms);

  ASSERT_TRUE(send_count.load() <= 17);

  while (queue.Take() != -1) {
    // Pass
  }

  producer.join();
}

void testPill() {
  static const size_t kThreads = 10;
  BoundedBlockingQueue<int> queue{1};

  std::vector<std::thread> threads;

  std::mt19937 twister;

  for (size_t i = 0; i < kThreads; ++i) {
    threads.emplace_back([&]() {
      std::this_thread::sleep_for(1ms * (twister() % 1000));

      ASSERT_EQ(queue.Take(), -1);
      queue.Put(-1);
    });
  }

  queue.Put(-1);

  for (auto& t : threads) {
    t.join();
  }
}

// --- Move only test ---

struct MoveOnly {
  MoveOnly() = default;

  MoveOnly(const MoveOnly& that) = delete;
  MoveOnly& operator=(const MoveOnly& that) = delete;

  MoveOnly(MoveOnly&& that) = default;
  MoveOnly& operator=(MoveOnly&& that) = default;
};

void testMoveOnly() {
  BoundedBlockingQueue<MoveOnly> queue{1};

  queue.Put(MoveOnly{});
  queue.Take();
}

// --- End of Move only test ---