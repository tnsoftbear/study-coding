#pragma once

#include <iostream>

#include "thread_state.hpp"

namespace hazard {

// Forward declaraction of Collector type
class Collector;

class Mutator {
  private:
    Collector* collector_;
  public:
    Mutator(Collector* collector);
    void PrintMe();
    static ThreadState& GetThreadState();
};

}