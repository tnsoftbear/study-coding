#include <iostream>
#include <vector>
#include <typeinfo>
#include <typeindex>

struct ThreadState {
    std::vector<void*> v;
    std::vector<std::type_index> types; // Вектор типов
};

class Collector {
public:
    void deleteObjects(ThreadState& state) {
        for (size_t i = 0; i < state.v.size(); ++i) {
            const std::type_index& type = state.types[i];
            void* ptr = state.v[i];

            // В зависимости от типа, выполняйте корректное удаление
            if (type == typeid(int)) {
                delete static_cast<int*>(ptr);
            } else if (type == typeid(double)) {
                delete static_cast<double*>(ptr);
            } // и так далее для других типов

            state.v[i] = nullptr; // Обнулите указатель
        }
        state.v.clear();
        state.types.clear();
    }
};

int main() {
    ThreadState state;
    // Добавление объектов в вектор и их типов
    int* intPtr = new int(42);
    double* doublePtr = new double(3.14);
    state.v.push_back(intPtr);
    state.v.push_back(doublePtr);
    state.types.push_back(typeid(int));
    state.types.push_back(typeid(double));

    Collector collector;
    collector.deleteObjects(state);

    return 0;
}