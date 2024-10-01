#pragma once

#include "fwd.hpp"

namespace hazard {

class Collector {
  private:
    Mutator* mutator_;
  public:
    Mutator MakeMutator();
};

}