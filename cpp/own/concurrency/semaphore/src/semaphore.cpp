#pragma once

#include "semaphore.hpp"

// For log method: Ll()
#include <cstdarg>
#include <sstream>
#include <thread>

// const bool kShouldPrint = true;
const bool kShouldPrint = false;

Semaphore::Semaphore(size_t tokens) : tokens_(tokens) {
  wait_mtx_ = std::make_unique<std::mutex>();
}

void Semaphore::Acquire() {
  std::unique_lock<std::mutex> lock(*wait_mtx_);
  wait_cv_.wait(lock, [this] { return tokens_ != 0; });
  --tokens_;
}

void Semaphore::Release() {
  std::lock_guard<std::mutex> lock(*wait_mtx_);
  ++tokens_;
  if (tokens_ == 1) {
    wait_cv_.notify_all();
  }
}

void Semaphore::Ll(const char* format, ...) {
  if (!kShouldPrint) {
    return;
  }

  char buf[250];
  std::ostringstream pid;
  pid << "[" << std::this_thread::get_id() << "]";
  sprintf(buf, "Semaphore: %s %s\n", pid.str().c_str(), format);
  va_list args;
  va_start(args, format);
  vprintf(buf, args);
  va_end(args);
}
