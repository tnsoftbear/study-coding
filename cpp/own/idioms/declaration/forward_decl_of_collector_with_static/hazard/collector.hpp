#pragma once

#include "mutator.hpp"
#include "thread_state.hpp"

namespace hazard {

class Collector {
  private:
    Mutator* mutator_;
  public:
    static ThreadState thread_state;
  public:
    Mutator MakeMutator();
    void Collect();
    static ThreadState& GetThreadState();
};

}