#pragma once

#include <iostream>

namespace hazard {

// Forward declaraction of Collector type
class Collector;

class Mutator {
  private:
    Collector* collector_;
  public:
    Mutator(Collector* collector);
    void PrintMe();
};

}