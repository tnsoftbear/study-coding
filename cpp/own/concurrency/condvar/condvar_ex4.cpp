#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>

class MyClass {
public:
    MyClass() {
        // Инициализация мьютекса в конструкторе
        mutex_ = std::make_unique<std::mutex>();

        // Инициализация unique_lock в конструкторе
        lock_ = std::make_unique<std::unique_lock<std::mutex>>(*mutex_);
    }

    void MyMethod() {
        // Использование lock_ для захвата мьютекса
        // lock_ будет владеть мьютексом в течение времени жизни объекта MyClass
        // Здесь вы можете выполнять операции, защищенные мьютексом

        // Пример использования condition_variable
        condition_var_.wait(*lock_, [this] { return isReady_; });

        // Далее продолжайте выполнение метода
    }

    void SetReady() {
        // Выставление флага готовности и уведомление ожидающих потоков
        isReady_ = true;
        condition_var_.notify_all();
    }

private:
    std::unique_ptr<std::mutex> mutex_; // Мьютекс как член объекта
    std::unique_ptr<std::unique_lock<std::mutex>> lock_; // unique_lock как член объекта
    std::condition_variable condition_var_;
    bool isReady_ = false;
};

int main() {
    MyClass obj;

    // Создание и запуск потока, который вызовет MyMethod
    std::thread t([&obj]() {
        obj.MyMethod();
    });

    // Ждем некоторое время, прежде чем установить флаг готовности
    std::this_thread::sleep_for(std::chrono::seconds(2));

    // Устанавливаем флаг готовности
    obj.SetReady();

    // Дожидаемся завершения потока
    t.join();

    return 0;
}
