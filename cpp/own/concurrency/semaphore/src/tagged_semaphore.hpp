#pragma once

#include "semaphore.hpp"

// --- TaggedSemaphore ---

template <class Tag>
class TaggedSemaphore {

 public:
  class Token;

  explicit TaggedSemaphore(size_t tokens);
  Token Acquire();
  void Release(Token&& token);

 private:
  Semaphore impl_;
};

// --- Token ---

template <class Tag>
class TaggedSemaphore<Tag>::Token {
  friend class TaggedSemaphore;

public:
  ~Token();
  Token(Token&& that);

  Token& operator=(Token&&) = delete;

private:
  Token() = default;
  void Invalidate();

private:
  bool valid_{true};
};