#include <thread>
#include <iostream>

#include "helpers.hpp"
#include "SPSCRingBufferOptimized.cpp"
#include "SPSCRingBufferNaive.cpp"
#include "SPSCRingBufferInterface.hpp"

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

void Stress(SPSCRingBufferInterface<int>& buffer) {
    Digest digest;
    static const int kValues = 10'000'000;

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
}

int main() {
    SPSCRingBufferOptimized<int> bufferOpt(256);
    SPSCRingBufferNaive<int> bufferNaive(256);
    StopWatch stop_watch;
    while (true)
    {
        stop_watch = StopWatch();
        Stress(bufferNaive);
        std::cout << "Elapsed for Naive: " << stop_watch.ElapsedMillis() << "ms" << std::endl;

        stop_watch = StopWatch();
        Stress(bufferOpt);
        std::cout << "Elapsed for Optimized: " << stop_watch.ElapsedMillis() << "ms" << std::endl;
    }
};