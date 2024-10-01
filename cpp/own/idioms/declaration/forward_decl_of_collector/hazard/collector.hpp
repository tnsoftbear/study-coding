#pragma once

#include "mutator.hpp"

namespace hazard {

class Collector {
  private:
    Mutator* mutator_;
  public:
    Mutator MakeMutator();
    void Collect();
};

}