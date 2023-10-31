#pragma once

#include <deque>

#include "tagged_semaphore.hpp"

// Bounded Blocking Multi-Producer/Multi-Consumer (MPMC) Queue

template <typename T>
class BoundedBlockingQueue {
 private:
  struct SemTag {};
  using Token = typename TaggedSemaphore<SemTag>::Token;
 
 public:
  explicit BoundedBlockingQueue(size_t capacity);

  void Put(T value);

  T Take();

 private:
  std::deque<T> buffer_;
  TaggedSemaphore<SemTag> free_sem_;
  TaggedSemaphore<SemTag> avail_sem_;
  TaggedSemaphore<SemTag> mutex_sem_;
};
