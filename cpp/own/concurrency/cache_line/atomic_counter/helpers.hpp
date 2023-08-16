#pragma once

#include <chrono>

class StopWatch {
 public:
  using Clock = std::chrono::steady_clock;
  using TimePoint = std::chrono::time_point<Clock>;
  using Duration = std::chrono::nanoseconds;

 public:
  StopWatch() : start_time_(Now()) {
  }

  Duration Elapsed() const {
    return Now() - start_time_;
  }

  int ElapsedMillis() const {
    auto elapsed = Elapsed();
    return std::chrono::duration_cast<std::chrono::milliseconds>(elapsed).count();
  }

  Duration Restart() {
    auto elapsed = Elapsed();
    start_time_ = Now();
    return elapsed;
  }

 private:
  static TimePoint Now() {
    return Clock::now();
  }

 private:
  TimePoint start_time_;
};

void ll(char* arg) {
  std::cout << arg << std::endl;
}