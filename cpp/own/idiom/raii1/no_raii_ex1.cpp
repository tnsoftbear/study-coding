// Запустить и смотреть в top как утекает память

struct S {
    int* p = nullptr;
    S(): p(new int (5)) {
        throw 1;
    }
    ~S() {
        delete p;
    }
};

void f() {
    int* p = new int(5);
    throw 1;
    delete p;
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