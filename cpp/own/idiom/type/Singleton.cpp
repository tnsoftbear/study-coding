#include <iostream>

class Singleton {
    static Singleton* instance;
    Singleton() {}
    Singleton(const Singleton&) = delete;
    Singleton& operator=(const Singleton&) = delete;
    Singleton(Singleton&&) = delete;
    Singleton& operator=(Singleton&&) = delete;
    ~Singleton() {}
public:
    static Singleton& getInstance() {
        if (instance == nullptr) {
            instance = new Singleton();
        }
        return *instance;
    }
    static void destroy() {
        delete instance;
    }
};

Singleton* Singleton::instance = nullptr;

int main() {
    Singleton& s1 = Singleton::getInstance();
}