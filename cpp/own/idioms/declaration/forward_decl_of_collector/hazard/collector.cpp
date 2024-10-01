#include "mutator.hpp"
#include "collector.hpp"

namespace hazard {

Mutator Collector::MakeMutator() {
  mutator_ = new Mutator(this);
  return *mutator_;
};

void Collector::Collect() {
  std::cout << "Collect" << std::endl;
}

}