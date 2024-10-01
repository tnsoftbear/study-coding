#include <optional>
#include <utility>

struct MoveOnly {
  MoveOnly() = default;
  MoveOnly(const MoveOnly&) = delete;
  MoveOnly(MoveOnly&&) {};
};

template <typename T>
struct Promise {
  T value;
  std::optional<T> opt;
  void SetValue(T&& v) {
    //value = std::forward<T>(v);
    opt.emplace(std::move(v));
  }
};

int main() {
  Promise<MoveOnly> p1;
  p1.SetValue(std::move(MoveOnly()));
}