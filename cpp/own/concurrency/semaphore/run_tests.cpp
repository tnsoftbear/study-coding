#include "test/semaphore_tests.cpp"
#include "test/queue_tests.cpp"

// Run options:
// 1) g++ run_tests.cpp && ./a.out
// 2) rm -rf ./build && cmake -B build && cmake --build build && ./build/run_tests

int main(int argc, char* argv[]) {
  testNonBlocking();
  testBlocking();
  testPingPong();

  testPutThenTake();
  testSmallFifo();
  testBigFifo();
  testCapacity();
  testPill();
  testMoveOnly();

  return 0;
}