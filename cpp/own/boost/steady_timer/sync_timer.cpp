#include <boost/asio.hpp>
#include <iostream>

// g++ -pthread ./sync_timer.cpp && ./a.out

int main() {
  using namespace boost;
  asio::io_context io;
  asio::steady_timer t(io, asio::chrono::seconds(5));
  t.wait();
  std::cout << "Hello, world!" << std::endl;
  return 0;
}