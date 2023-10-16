#include <boost/asio.hpp>
#include <boost/bind/bind.hpp>
#include <iostream>

// https://think-async.com/Asio/asio-1.22.1/doc/asio/tutorial/tuttimer3.html
// g++ -pthread ./async_timer_repeat.cpp && ./a.out

using namespace boost;

void print(const boost::system::error_code &e, asio::steady_timer *t,
           int *count) {
  (void)e;
  if (*count < 5) {
    std::cout << "Hello in " << *count << " seconds" << std::endl;
    ++(*count);
    t->expires_at(t->expiry() + asio::chrono::seconds(1));
    t->async_wait(boost::bind(print, e, t, count));
  }
}

int main() {
  asio::io_context io;
  int count = 0;
  asio::steady_timer t(io, asio::chrono::seconds(0));
  t.async_wait(boost::bind(print, asio::placeholders::error, &t, &count));
  io.run();
  return 0;
}