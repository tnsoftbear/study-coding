#pragma once

#include <iostream>

#include "collector.hpp"
#include "mutator.hpp"

namespace hazard {

Mutator::Mutator(Collector* collector) : collector_(collector) {
  collector->Collect();
}

void Mutator::PrintMe() {
  std::cout << "I'm mutator, thread state is " << Mutator::GetThreadState().data << std::endl;
}

ThreadState& Mutator::GetThreadState() {
  return Collector::GetThreadState(); // thread_state;
}

}