#include "collector.hpp"

namespace hazard {

ThreadState Collector::thread_state{53};

Mutator Collector::MakeMutator() {
  mutator_ = new Mutator(this);
  return *mutator_;
};

void Collector::Collect() {
  auto& ts = Collector::GetThreadState();
  std::cout << "I'm collector, thread state is: " << ts.data << std::endl;
}

ThreadState& Collector::GetThreadState() {
  return thread_state;
}

}