#include <iostream>
#include <atomic>
#include <thread>
#include <vector>

#include "helpers.hpp"

static const size_t kCacheLineSize = 64;
static const size_t kShards = 4;

class CounterInterface {
  public:
    virtual void Increment() = 0;
    virtual size_t Get() = 0;
};

/**
 * Счётчик не параллелен
 */
class AtomicCounter : public CounterInterface {
    public:
      void Increment() {
        value_.fetch_add(1);
      }

      size_t Get() {
        return value_.load();
      }

    private:
      std::atomic<size_t> value_{0};
};

/**
 * Шардированный счётчик выбирает для каждого потока свой элемент в массиве счётчиков
*/
class ShardedCounter : public CounterInterface {
    private:
      class Shard {
        public:
            void Increment() {
                value_.fetch_add(1);
            }

            size_t Get() {
                return value_.load();
            }
        private:
            alignas(kCacheLineSize) std::atomic<size_t> value_{0};
      };

    public:
      void Increment() {
        size_t shard_index = GetThisThreadShard();
        shards_[shard_index].Increment();
      }

      size_t Get() {
        size_t value = 0;
        for (size_t i = 0; i < kShards; ++i) {
            value += shards_[i].Get();
        }
        return value;
      }

    private:
      static size_t GetThisThreadShard() {
        static std::hash<std::thread::id> hasher;
        auto tid = std::this_thread::get_id();
        return hasher(tid) % kShards;
      }

    private:
      std::array<Shard, kShards> shards_;
};


void Stress(CounterInterface& counter) {
    std::vector<std::thread> threads;
    for (size_t i = 0; i < kShards; ++i) {
        threads.emplace_back([&counter]() {
            for (size_t j = 0; j < 10'000'000; ++j) {
                counter.Increment();
            }
        });
    }

    for (auto& t : threads) {
        t.join();
    }
}

int main() {
  AtomicCounter atomicCounter;
  ShardedCounter shardedCounter;
  StopWatch stop_watch;
  while (true) {
    stop_watch = StopWatch();
    Stress(atomicCounter);
    std::cout << "Atomic counter: " << atomicCounter.Get()
        << ", Elapsed: " << stop_watch.ElapsedMillis() << "ms" << std::endl;

    stop_watch = StopWatch();
    Stress(shardedCounter);
    std::cout << "Sharded counter: " << shardedCounter.Get()
        << ", Elapsed: " << stop_watch.ElapsedMillis() << "ms" << std::endl;
  }
}