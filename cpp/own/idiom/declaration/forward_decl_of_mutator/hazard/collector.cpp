#pragma once

#include "mutator.hpp"

namespace hazard {

Mutator Collector::MakeMutator() {
  mutator_ = new Mutator(*this);
  return *mutator_;
};

}