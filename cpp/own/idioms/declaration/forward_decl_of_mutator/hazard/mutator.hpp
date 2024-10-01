#pragma once

#include <iostream>

#include "collector.hpp"

namespace hazard {

class Mutator {
  private:
    Collector& collector_;
  public:
    Mutator(Collector& collector) : collector_(collector) {};

    void PrintMe() {
        std::cout << "I'm mutator" << std::endl;
    }
};

}