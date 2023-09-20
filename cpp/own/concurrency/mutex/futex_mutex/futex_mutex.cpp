#include <cstdint>
#include <atomic>

class FutexMutex {
 public:
  void lock() {
    uint32_t c;
    if ((c = Cmpxchg(Status::Unlocked, Status::Locked)) != Status::Unlocked) {
      do {
        if (c == Status::Sleeping
          || Cmpxchg(Status::Locked, Status::Sleeping) != Status::Unlocked
        ) {
          m_.wait(Status::Sleeping);
        }
      } while ((c = Cmpxchg(Status::Unlocked, Status::Sleeping)) != Status::Unlocked);
    }
  }

  void unlock() {
    if (m_.exchange(Status::Unlocked) != Status::Locked) {
      m_.notify_one();
    }
  }

 private:
  uint32_t Cmpxchg(uint32_t old_st, uint32_t new_st) {
    m_.compare_exchange_strong(old_st, new_st);
    return old_st;
  }

 private:
  std::atomic<uint32_t> m_{Status::Unlocked};
  enum Status {
    Unlocked = 0,
    Locked = 1,   // Mutex is locked without waiters
    Sleeping = 2  // Mutex is locked with waiters
  };
};
