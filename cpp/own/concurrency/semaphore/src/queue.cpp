#include "queue.hpp"
#include "tagged_semaphore.cpp"

// Bounded Blocking Multi-Producer/Multi-Consumer (MPMC) Queue

template <typename T>
BoundedBlockingQueue<T>::BoundedBlockingQueue(size_t capacity)
  : free_sem_(capacity)
  , avail_sem_(0)
  , mutex_sem_(1)
{
}

template <typename T>
void BoundedBlockingQueue<T>::Put(T value) {
  Token token(std::move(free_sem_.Acquire()));

  Token mutex_token(std::move(mutex_sem_.Acquire()));
  buffer_.push_back(std::move(value));
  mutex_sem_.Release(std::move(mutex_token));

  avail_sem_.Release(std::move(token));
}

template <typename T>
T BoundedBlockingQueue<T>::Take() {
  Token token(std::move(avail_sem_.Acquire()));

  Token mutex_token(std::move(mutex_sem_.Acquire()));
  T value = std::move(buffer_.front());
  buffer_.pop_front();
  mutex_sem_.Release(std::move(mutex_token));

  free_sem_.Release(std::move(token));
  
  return value;
}

