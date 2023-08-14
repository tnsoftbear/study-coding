// Память не должна утекать, потому что при уничтожении RAII-объекта SmartPtr,
// так же подчищается динамическая память занятая его ресурсом.

template <typename T>
class SmartPtr {
private:
    T* ptr;
public:
    SmartPtr(T* ptr) : ptr(ptr) {}
    ~SmartPtr() { delete ptr; }
    T& operator*() { return *ptr; }
    T* operator->() { return ptr; }
};

struct S
{
    SmartPtr<int> p = nullptr;
    S(): p(new int (5)) {
        throw 1;
    }
    ~S() {}
};

void f() {
    SmartPtr<int> p = new int(5);
    throw 1;
}

int main() {
  for (;;) {
    try {
        f();
    } catch(...) {
        // return 1;
    }
  }
}