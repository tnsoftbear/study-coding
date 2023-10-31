#pragma once

#include <cassert>

#include "tagged_semaphore.hpp"

/**
 * --- Token ---
 *
 * Префикс TaggedSemaphore<Tag>::Token:: перед методами класса Token указывает на то,
 * что эти методы принадлежат классу Token, который является вложенным в класс TaggedSemaphore с параметром Tag. 
 * Вложенный класс имеет область видимости и зависимость от параметра типа (Tag), 
 * и методы этого вложенного класса могут быть реализованы вне объявления класса TaggedSemaphore.
 */

template <class Tag>
TaggedSemaphore<Tag>::Token::~Token() {
    assert(!valid_);
}

template <class Tag>
TaggedSemaphore<Tag>::Token::Token(Token&& that) {
    that.Invalidate();
}

template <class Tag>
void TaggedSemaphore<Tag>::Token::Invalidate() {
    assert(valid_);
    valid_ = false;
}

// --- TaggedSemaphore ---

template <class Tag>
TaggedSemaphore<Tag>::TaggedSemaphore(size_t tokens)
  : impl_(tokens) {
}

/**
 * Ключевое слово typename используется для указания компилятору на то, 
 * что TaggedSemaphore<Tag>::Token представляет тип, а не статический член класса. 
 */
template <class Tag>
typename TaggedSemaphore<Tag>::Token TaggedSemaphore<Tag>::Acquire() {
    impl_.Acquire();
    return Token{};
}

template <class Tag>
void TaggedSemaphore<Tag>::Release(Token&& token) {
    impl_.Release();
    token.Invalidate();
}

