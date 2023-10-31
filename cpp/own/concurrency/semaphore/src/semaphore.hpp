#pragma once

#include <condition_variable>
#include <mutex>

class Semaphore {
public:
  explicit Semaphore(size_t tokens);

  void Acquire();

  void Release();

private:
  size_t tokens_;
  std::unique_ptr<std::mutex> wait_mtx_;
  std::condition_variable wait_cv_;

private:
  void Ll(const char* format, ...);
};
