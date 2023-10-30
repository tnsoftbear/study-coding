#include <iostream>

#include "collector.hpp"
#include "mutator.hpp"

namespace hazard {

Mutator::Mutator(Collector* collector) : collector_(collector) {
  collector->Collect();
}

void Mutator::PrintMe() {
  std::cout << "I'm mutator" << std::endl;
}

}