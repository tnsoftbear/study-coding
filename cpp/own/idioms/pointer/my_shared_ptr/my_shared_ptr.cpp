#include <cstddef>
#include <cstdint>
#include <iostream>
#include <stdexcept>
#include <utility>

using std::cout;
using std::endl;

namespace detail {

struct SplitCount {
  int32_t transient;
  int32_t strong; // > 0
};

static_assert(sizeof(SplitCount) == sizeof(size_t), "Not supported");

} // namespace detail

//////////////////////////////////////////////////////////////////////

template <typename T> class SharedPtr {
public:
  SharedPtr(T* data, detail::SplitCount* counter)
      : data_ptr_(data)
      , counter_ptr_(counter) {
    IncrementStrong();
    PrintMe("SharedPtr(T* data, detail::SplitCount* count)");
  }

  SharedPtr()
      : data_ptr_(nullptr)
      , counter_ptr_(nullptr) {
    PrintMe("SharedPtr()");
  }

  // Copy ctor
  SharedPtr(const SharedPtr<T>& that)
      : data_ptr_(that.data_ptr_)
      , counter_ptr_(that.counter_ptr_) // that.count_ + 1 ?
  {
    IncrementStrong();
    PrintMe("SharedPtr(const SharedPtr<T>& that)");
  }

  // Copy assignment
  SharedPtr<T>& operator=(const SharedPtr<T>& that) {
    if (this != &that) {
      Reset();
      data_ptr_ = that.data_ptr_;
      counter_ptr_ = that.counter_ptr_;
      IncrementStrong();
    }
    PrintMe("SharedPtr<T>& operator=(const SharedPtr<T>& that)");
    return *this;
  }

  SharedPtr(SharedPtr<T>&& that)
      : data_ptr_(that.data_ptr_)
      , counter_ptr_(that.counter_ptr_) {
    that.data_ptr_ = nullptr;
    that.counter_ptr_ = nullptr;
    PrintMe("SharedPtr(SharedPtr<T>&& that)");
  }

  SharedPtr<T>& operator=(SharedPtr<T>&& that) {
    if (this != &that) {
      Reset();
      data_ptr_ = std::exchange(that.data_ptr_, nullptr);
      counter_ptr_ = std::exchange(that.counter_ptr_, nullptr);
    }
    PrintMe("SharedPtr<T>& operator=(SharedPtr<T>&& that)");
    return *this;
  }

  ~SharedPtr() {
    PrintMe("~SharedPtr()");
    Reset();
  }

  T* operator->() const { return data_ptr_; }

  T& operator*() const { return *data_ptr_; }

  explicit operator bool() const { return data_ptr_ != nullptr; }

  void Reset() {
    if (counter_ptr_) {
      counter_ptr_->strong--;
      if (counter_ptr_->strong == 0) {
        delete data_ptr_;
        delete counter_ptr_;
      }
    }
    // data_ptr_ = nullptr;
    // counter_ptr_ = nullptr;
    PrintMe("Reset()");
  }

  void IncrementStrong(int32_t inc = 1) {
    if (counter_ptr_) {
      counter_ptr_->strong += inc;
    }
  }

  int32_t GetStrongCount() const { return counter_ptr_->strong; }

private:
  T* data_ptr_;
  detail::SplitCount* counter_ptr_;

private:
  void PrintMe(std::string prefix = "") {
    if (counter_ptr_) {
      cout << "SharedPtr::" << prefix << "; data_: " << data_ptr_
           << "; count_.strong: " << counter_ptr_->strong
           << "; count_.transient: " << counter_ptr_->transient << endl;
    } else {
      cout << "SharedPtr::" << prefix << " " << endl;
    }
  }
};

template <typename T, typename... Args>
SharedPtr<T> MakeShared(Args&&... args) {
  T* data = new T(std::forward<Args>(args)...);
  detail::SplitCount* count = new detail::SplitCount{0, 1};
  return SharedPtr<T>(data, count);
}
