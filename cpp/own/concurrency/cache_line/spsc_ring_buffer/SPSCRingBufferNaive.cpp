#pragma once

#include <atomic>
#include <vector>
#include <iostream>

#include "SPSCRingBufferInterface.hpp"

template <typename T>
class SPSCRingBufferNaive : public SPSCRingBufferInterface<T> {
    struct Slot {
        T value;
    };
    public:
        explicit SPSCRingBufferNaive(const size_t capacity) : buffer_(capacity + 1){
        }
        bool TryProduce(T value) {
            const size_t curr_tail = tail_.load();
            const size_t curr_head = head_.load(); // Убрали

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
            const size_t curr_tail = tail_.load();

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
        std::atomic<size_t> head_{0}; // Owned by consumer
        std::atomic<size_t> tail_{0}; // Owned by producer
};