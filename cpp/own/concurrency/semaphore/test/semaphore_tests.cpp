#include <chrono>
#include <thread>

#include "../src/semaphore.hpp"
#include "assert_macros.hpp"

using namespace std::chrono_literals;

void testNonBlocking() {
  Semaphore semaphore(2);

  semaphore.Acquire(); // -1
  semaphore.Release(); // +1

  semaphore.Acquire(); // -1
  semaphore.Acquire(); // -1
  semaphore.Release(); // +1
  semaphore.Release(); // +1
}

void testBlocking() {
  Semaphore semaphore(0);

  bool touched = false;

  std::thread touch([&]() {
    semaphore.Acquire();
    touched = true;
  });

  std::this_thread::sleep_for(250ms);

  ASSERT_FALSE(touched);

  semaphore.Release();
  touch.join();

  ASSERT_TRUE(touched);
}

void testPingPong() {
  Semaphore my{1};
  Semaphore that{0};

  int step = 0;

  std::thread opponent([&]() {
    that.Acquire();
    ASSERT_EQ(step, 1);
    step = 0;
    my.Release();
  });

  my.Acquire();
  ASSERT_EQ(step, 0);
  step = 1;
  that.Release();

  my.Acquire();
  ASSERT_TRUE(step == 0);

  opponent.join();
}