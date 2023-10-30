#include <iostream>
#include <vector>
#include <functional>

class Widget {
 public:
    int data;
};

struct ThreadState {
    std::vector<void*> v;
    std::function<void(void*)> deleteFunction;

    ThreadState() : deleteFunction(nullptr) {}

    void setDeleteFunction(std::function<void(void*)> func) {
        deleteFunction = func;
    }

    void deleteObjects() {
        if (deleteFunction) {
            for (void* ptr : v) {
                deleteFunction(ptr);
            }
        }
        v.clear();
    }
};

int main() {
    ThreadState state;

    // Установите лямбда-функцию для удаления объектов
    state.setDeleteFunction([](void* ptr) {
        // Здесь может быть логика удаления, в зависимости от типа объекта
        // Пример для int
        if (Widget* widgetPtr = static_cast<Widget*>(ptr)) {
            delete widgetPtr;
        }
    });

    // Добавление объектов в вектор
    Widget* wPtr = new Widget{42};
    state.v.push_back(wPtr);

    state.deleteObjects();

    return 0;
}