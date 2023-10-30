#include "hazard/collector.cpp"
#include "hazard/mutator.cpp"

/**
 * Сейчас происходит следующее разрешение зависимостей:
 * main.cpp -> collector.cpp -> collector.hpp -> mutator.hpp -> Forward declaration of class Collector.
 * main.cpp -> mutator.cpp -> т.к. pragma once, то повторного включения collector.hpp, mutator.hpp не происходит.

 * Если же я здесь включаю хедеры collector.hpp, mutator.hpp вместо .cpp файлов,
 * тогда необходимо скомпилировать и слинковать файлы .cpp вместе с main.cpp, 
 * чтобы компилятор мог найти определение метода MakeMutator() во время линковки.
 * g++ main.cpp hazard/collector.cpp hazard/mutator.cpp && ./a.out
 */

int main() {
    auto collector = new hazard::Collector();
    auto mutator = collector->MakeMutator();
    mutator.PrintMe();
}