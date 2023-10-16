#include <boost/asio.hpp>
#include <boost/bind/bind.hpp>
#include <iostream>

// https://think-async.com/Asio/asio-1.22.1/doc/asio/tutorial/tuttimer5.html
// g++ -pthread ./async_timer_multithread_object.cpp && ./a.out

using namespace boost;

class printer {
public:
  printer(asio::io_context &io)
      : strand_(asio::make_strand(io)), timer_1(io, asio::chrono::seconds(1)),
        timer_2(io, asio::chrono::seconds(1)), count_(0) {
    timer_1.async_wait(
        asio::bind_executor(strand_, boost::bind(&printer::print1, this)));
    timer_2.async_wait(
        asio::bind_executor(strand_, boost::bind(&printer::print2, this)));
  }

  ~printer() { std::cout << "Final count is " << count_ << std::endl; }

  void print1() {
    if (count_ < 10) {
      std::cout << "Timer-1: " << count_ << std::endl;
      ++count_;

      timer_1.expires_at(timer_1.expiry() + asio::chrono::seconds(1));
      timer_1.async_wait(
          asio::bind_executor(strand_, boost::bind(&printer::print1, this)));
    }
  }

  void print2() {
    if (count_ < 10) {
      std::cout << "Timer-2: " << count_ << std::endl;
      ++count_;

      timer_2.expires_at(timer_2.expiry() + asio::chrono::seconds(1));
      timer_2.async_wait(
          asio::bind_executor(strand_, boost::bind(&printer::print2, this)));
    }
  }

private:
  asio::steady_timer timer_1;
  asio::steady_timer timer_2;
  int count_;
  asio::strand<asio::io_context::executor_type> strand_;
};

int main() {
  asio::io_context io;
  printer p(io);
  io.run();

  return 0;
}