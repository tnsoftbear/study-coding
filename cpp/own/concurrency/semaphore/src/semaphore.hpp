#pragma once

#include <condition_variable>
#include <mutex>

#include <cstdlib>

//#include <twist/ed/stdlike/thread.hpp>
//#include <sstream>

//const bool kShouldPrint = true;
const bool kShouldPrint = false;

class Semaphore {
 public:
  explicit Semaphore(size_t tokens) {
    tokens_ = tokens;
  }

  ~Semaphore() {
    wait_cv_.notify_all();
  }

  void Acquire() {
    std::unique_lock<std::mutex> lock(*wait_mtx_);
    wait_cv_.wait(lock, [this]{
      return tokens_ != 0;
    });
    --tokens_;
  }

  void Release() {
    std::lock_guard<std::mutex> lock(*wait_mtx_);
    ++tokens_;
    if (tokens_ == 1) {
      wait_cv_.notify_all();
    }
  }

 private:
  size_t tokens_;
  std::unique_ptr<std::mutex> wait_mtx_ = std::make_unique<std::mutex>();
  std::condition_variable wait_cv_;

 private:
  // void Ll(const char* format, ...) {
  //   if (!kShouldPrint) {
  //     return;
  //   }

  //   char buf [250];
  //   std::ostringstream pid;
  //   pid << "[" << twist::ed::stdlike::this_thread::get_id() << "]";
  //   sprintf(buf, "Semaphore: %s %s\n", pid.str().c_str(), format);
  //   va_list args;
  //   va_start(args, format);
  //   vprintf(buf, args);
  //   va_end(args);
  // }
};
