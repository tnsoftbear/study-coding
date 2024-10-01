#include <iostream>

struct Base {
    std::string s1 = "Base s1";
    Base()=default;
    Base(std::string s1): s1(s1) {};
    std::string f() {
        return "Base f(), s1=" + s1;
    }
    
    std::string ff(int i) {
        return "Base ff(" + std::to_string(i) + ")";
    }

    // Виртуальные ф-ции замедляют исполнение (работа с таблицей виртуальных функций),
    // поэтому этот механизм не присутствуе по-умолчанию как в Java и требует явного определения.
    virtual std::string fff() {
        return "Base fff(), s1=" + s1;
    }

    virtual std::string fFinal() final { // final запрещает перегрузку ф-ции
        return "Base fFinal(), s1=" + s1;
    }
};

struct Derived: public Base {
    std::string s1 = "Derived s1";
    std::string f() {
        return "Derived f()";
    }
    using Base::ff;
    std::string ff(double d) {
        return "Derived ff(" + std::to_string(d) + ")";
    }
};

struct DerivedBro: public Base {
    std::string s1 = "DerivedBro s1";
    DerivedBro(std::string s1): Base(s1) {}
    std::string fff() override {
        return "DerivedBro fff(), s1=" + s1;
    }
};

struct DerivedSis: public Base {
    std::string s1 = "DerivedSis s1";
    std::string f() {
        return "DerivedSis f(), s1=" + s1;
    }
    std::string fff() override {
        return "DerivedSis fff(), s1=" + s1;
    }
};

int main() {
    auto b1 = Base();
    auto d1 = Derived();
    printf("sizeof(Base): %lu, sizeof(Derived): %lu\n", sizeof(Base), sizeof(Derived));
    printf("b1.s1: %s, d1.s1: %s, d1.Base::s1: %s \n", 
        b1.s1.c_str(), d1.s1.c_str(), d1.Base::s1.c_str());
    // b1.s1: Base s1, d1.s1: Derived s1, d1.Base::s1: Base s1
    printf("b1.f(): %s, d1.f(): %s, d1.Base::f(): %s \n", 
        b1.f().c_str(), d1.f().c_str(), d1.Base::f().c_str());
    // b1.f(): Base f(), s1=Base s1, d1.f(): Derived f(), d1.Base::f(): Base f(), s1=Base s1
    printf("b1.ff(3.14): %s, d1.ff(3.14): %s, d1.ff(3): %s\n",
        b1.ff(3.14).c_str(), d1.ff(3.14).c_str(), d1.ff(3).c_str());
    // b1.ff(3.14): Base ff(3), d1.ff(3.14): Derived ff(3.140000), d1.ff(3): Base ff(3)
    std::cout << std::endl;

    Base* d1bro = new DerivedBro("Constructed DerivedBro s1"); // Здесь инициализируется Base::s1, а не DerivedBro::s1
    printf("d1bro->s1: %s, d1bro->Base::s1: %s\n", d1bro->s1.c_str(), d1bro->Base::s1.c_str());
    // d2bro.fff(): Base fff(), s1=Constructed DerivedBro s1

    printf("d1bro->f(): %s\n", d1bro->f().c_str());
    // d1bro->f(): Base f(), s1=Constructed DerivedBro s1

    printf("d1bro->fff(): %s\n", d1bro->fff().c_str()); // вызов перегруженного виртуального метода, он покажет DerivedBro::s1
    // d1bro->fff(): DerivedBro fff(), s1=DerivedBro s1

    printf("d1bro->Base::fff(): %s\n", d1bro->Base::fff().c_str());
    // d1bro->Base::fff(): Base fff(), s1=Constructed DerivedBro s1

    Base d2bro = DerivedBro("Constructed DerivedBro s1");
    printf("d2bro.fff(): %s\n", d2bro.fff().c_str());
    // d2bro.fff(): Base fff(), s1=Constructed DerivedBro s1

    std::cout << std::endl;
    Base* d1sis = new DerivedSis();
    // Здесь видим, что d1sis->fff() вызвалась в контексте DerivedSis объекта потому что виртуальная перегруженная ф-ция (dynamic late binding), 
    // а d1sis->f() вызвана в контексте Base, потому что статическое связывание (ealry static binding), несмотря на то, что DerivedSis::f() тоже определена.
    printf("d1sis->fff(): %s, d1sis->f(): %s\n", d1sis->fff().c_str(), d1sis->f().c_str());
    // d1sis->fff(): DerivedSis fff(), s1=DerivedSis s1, d1sis->f(): Base f(), s1=Base s1
}
