#include "hazard/collector.cpp"

int main() {
    auto collector = new hazard::Collector();
    auto mutator = collector->MakeMutator();
    mutator.PrintMe();
}