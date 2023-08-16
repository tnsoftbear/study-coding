#pragma once

#include <atomic>
#include <vector>
#include <iostream>

#include "SPSCRingBufferInterface.hpp"

template <typename T>
class SPSCRingBufferOptimized : public SPSCRingBufferInterface<T> {
    struct Slot {
        // alignas(kCacheLineSize) T value; // не помогает
        T value;
    };
    public:
        explicit SPSCRingBufferOptimized(const size_t capacity) : buffer_(capacity + 1){
        }
        bool TryProduce(T value) {
            const size_t curr_tail = tail_.load();
            
            // Оптимизация-Б:
            // Вместо того, чтобы каждый раз читать голову кольцевого буфера (head_) в продюсере, 
            // и тем самым провоцировать синхронизацию кэш-линии.
            // Мы будем перечитывать голову (сохраняя её в head_cached_) только когда хвост близок к концу, т.е. когда может срабоать условие в IsFull()
            // const size_t curr_head = head_.load(); // Убрали
            if (Next(curr_tail) == head_cached_) {  // Оптимизация-Б
                head_cached_ = head_.load();
            }
            const size_t curr_head = head_cached_;  // Оптимизация-Б

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

            // Оптимизация-Б алгоритма: аналогично как для продюсера.
            // const size_t curr_tail = tail_.load();
            if (curr_head == tail_cached_) {        // Оптимизация-Б
                tail_cached_ = tail_.load();
            }
            const size_t curr_tail = tail_cached_;  // Оптимизация-Б

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
        // Эта оптимизация-А обретает смысл только с учётом оптимизации-Б
        // Потому что без оптимизации-Б, к этим полям обращаются оба - как продюсер, так и консьюмер, что приводит к синхронизации их кэш-линий.
        // char padding_1[kCacheLineSize]; // оптимизация-А (альтернативный паддинг)
        alignas(kCacheLineSize) std::atomic<size_t> head_{0}; // оптимизация-А (alignas)
        size_t head_cached_{0};

        // char padding_2[kCacheLineSize]; // оптимизация-А (альтернативный паддинг)
        alignas(kCacheLineSize) std::atomic<size_t> tail_{0}; // оптимизация-А
        size_t tail_cached_{0};
};