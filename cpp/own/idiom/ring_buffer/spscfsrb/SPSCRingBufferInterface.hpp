#pragma once

#include <iostream>

static const size_t kCacheLineSize = 64;

template <typename T>
class SPSCRingBufferInterface {
    public:
        virtual bool TryProduce(T value) = 0;
        virtual bool TryConsume(T& value) = 0;
};
