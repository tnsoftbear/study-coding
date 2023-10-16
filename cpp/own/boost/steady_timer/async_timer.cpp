#include <boost/asio.hpp>
#include <iostream>

// https://think-async.com/Asio/asio-1.22.1/doc/asio/tutorial/tuttimer2.html
// g++ -pthread ./async_timer.cpp && ./a.out

using namespace boost;

void print(const boost::system::error_code &e) {
  (void)e;
  std::cout << "Hello in 3 seconds" << std::endl;
}

int main() {
  asio::io_context io;
  asio::steady_timer t(io, asio::chrono::seconds(3));
  // Библиотека asio обеспечивает гарантию того, что обработчики обратного вызова будут вызываться только из потоков, которые в данный момент вызывают io_context::run(). 
  t.async_wait(&print);
  // Функция io_context::run() также будет продолжать работать, пока еще есть «работа». В этом примере работа заключается в асинхронном ожидании таймера, поэтому вызов не вернется до тех пор, пока не истечет время таймера и не завершится обратный вызов.
  io.run();
  return 0;
}