#include <atomic>
#include <vector>
#include <new>
#include <thread>
#include <iostream>

#include "helpers.hpp"

/**
 * Single-Producer/Single-Consumer Fixed-Size Ring Buffer
 * 
 * Компиляция:
 * g++ ring_buffer1.cpp -pthread
 * 
 * buffer_: Циклический буфер.
 * head_: указывает на первую занятую ячейку.
 * tail_: указывает на первую свободную ячейку.
 * Когда Next(tail_) == head_, то буфер полон.
 * 
 * https://www.youtube.com/watch?v=3NmyqUmvlbQ
 */


static const size_t kCacheLineSize = 64;
static const int kValues = 10'000'000;

template <typename T>
class SPSCRingBuffer {
    struct Slot {
        // alignas(kCacheLineSize) T value; // не помогает
        T value;
    };
    public:
        explicit SPSCRingBuffer(const size_t capacity) : buffer_(capacity + 1){
        }
        bool TryProduce(T value) {
            const size_t curr_tail = tail_.load();
            
            // Оптимизация алгоритма:
            // Вместо того, чтобы каждый раз читать голову кольцевого буфера (head_) в продюсере, 
            // и тем самым провоцировать синхронизацию кэш-линии.
            // Мы будем перечитывать голову (сохраняя её в head_cached_) только когда хвост близок к концу, т.е. когда может срабоать условие в IsFull()
            // const size_t curr_head = head_.load();
            if (Next(curr_tail) == head_cached_) {
                head_cached_ = head_.load();
            }
            const size_t curr_head = head_cached_;

            if (IsFull(curr_head, curr_tail)) {
                // std::cout << "TryProduce(" << value << ")->IsFull(), curr_head: " << curr_head << " curr_tail: " << curr_tail << std::endl;
                return false;
            }

            // std::cout << "TryProduce(" << value << "), curr_head: " << curr_head << " curr_tail: " << curr_tail << std::endl;

            buffer_[curr_tail].value = std::move(value);
            tail_.store(Next(curr_tail));
            return true;
        }

        bool TryConsume(T& value) {
            const size_t curr_head = head_.load();

            // Оптимизация алгоритма: аналогично как для продюсера.
            // const size_t curr_tail = tail_.load();
            if (curr_head == tail_cached_) {
                tail_cached_ = tail_.load();
            }
            const size_t curr_tail = tail_cached_;


            if (IsEmpty(curr_head, curr_tail)) {
                // std::cout << "TryConsume()->IsEmpty(), curr_head: " << curr_head << " curr_tail: " << curr_tail << std::endl;
                return false;
            }

            value = std::move(buffer_[curr_head].value);

            // std::cout << "TryConsume(" << value << "), curr_head: " << curr_head << " curr_tail: " << curr_tail << std::endl;

            head_.store(Next(curr_head));
            return true;
        }

    private:
        size_t Next(size_t slot) {
            return (slot + 1) % buffer_.size();
        }

        bool IsEmpty(size_t head, size_t tail) {
            return head == tail;
        }

        bool IsFull(size_t head, size_t tail) {
            return Next(tail) == head;
        }

    private:
        std::vector<Slot> buffer_;
        // Здесь происходит лишная синхронизация кэш-линии (false sharing)
        // Можно попробовать разнести значения по разным кэш-линиям, добавив между ними 64-байта пустых данных
        // Но это бесполезная оптимизация, потому что к этим полям обращается как продюсер, так и консьюмер.
        // char padding_1[kCacheLineSize]; 
        std::atomic<size_t> head_{0}; // Owned by consumer
        size_t head_cached_{0};

        // char padding_2[kCacheLineSize];
        std::atomic<size_t> tail_{0}; // Owned by producer
        size_t tail_cached_{0};
        
};

class Backoff {
    public:
        void operator()() {
            std::this_thread::yield();
        }
};

class Digest {
  public:
    void Feed(int value) {
        digest_ = std::hash<int>()(value) ^ (digest_ << 1);
    }
    int Value() {
        return digest_;
    }
  private:
    size_t digest_ = 0;
};

void Stress() {
    SPSCRingBuffer<int> buffer(256);
    StopWatch stop_watch;
    Digest digest;

    // Single producer
    std::thread producer([&]() {
        for (int i = 0; i < kValues; ++i) {
            Backoff backoff;
            while (!buffer.TryProduce(i)) {
                backoff();
            }
        }
    });

    // Single consumer
    std::thread consumer([&]() {
        for (int i = 0; i < kValues; ++i) {
            Backoff backoff;
            int value;
            while (!buffer.TryConsume(value)) {
                backoff();
            }
            digest.Feed(value);
        }
    });

    producer.join();
    consumer.join();

    std::cout << "Digest = " << digest.Value() << std::endl;
    std::cout << "Elapsed: " << stop_watch.ElapsedMillis() << "ms" << std::endl;
}

int main() {
    while (true)
    {
        Stress();
    }
};