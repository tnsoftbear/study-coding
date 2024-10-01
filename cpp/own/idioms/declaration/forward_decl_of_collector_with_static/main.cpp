#include "hazard/collector.cpp"
#include "hazard/mutator.cpp"

int main() {
    auto collector = new hazard::Collector();
    auto mutator = collector->MakeMutator();
    mutator.PrintMe();
    //collector->Collect();
}