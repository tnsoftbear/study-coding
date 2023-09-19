#include <cstring>
#include <initializer_list>
#include <algorithm>
#include <iostream>

using std::cout;
using std::endl;

class String {
    private:
        char* str = nullptr;
        size_t size = 0;
        size_t cap = 0;
    public:
        // String() = default;
        // String() = delete;
        // String(String&) = default;
        // String(const String&) = default;
        // String(String&&) = default;
        // ~String() = default;

        String() {
            char* info = DebugInfo2_();
            printf("Default Constructor - String() %s\n", info);
            free(info);
        };
        
        String(size_t sz, const char c = '\0')
        : str(new char[sz])     // member initializer list
        , cap(sz)
        , size(sz)
        {
            // this->sz = sz;
            // str = new char[sz];
            memset(str, c, sz);
            ll_("Parameterized Constructor - String(size_t, const char)");
        }

        String(const String& s)
        : String(s.size, '\0')    // delegating constructor (since c++11)
        {
            memcpy(str, s.str, size);
            ll_("Copy Constructor - String(const String&)");
        }

        String(std::initializer_list<const char> lst) {
            size = lst.size();
            str = new char[size];
            std::copy(lst.begin(), lst.end(), str);
            ll_("Initializer List Constructor - String(std::initializer_list<const char>)");
        }

        String(const char* str) {
            size = strlen(str);
            this->str = new char[size];
            memcpy(this->str, str, size);
            ll_("Parameterized Constructor - String(const char*)");
        }

        ~String();

        // Передача аргумента по значению вместо ссылки (copy elision), т.е. вместо `operator=(const String& other)`:
        // "if you're going to make a copy of something in a function, let the compiler do it in the parameter list"
        String& operator=(String other) {
            if (this == &other) return *this;  // не имеет большого смысла при использовании Copy-and-swap().
            // this->~String(); // вызов деструктора это UB. После его вызова компилятор может считать, что объект всё.
            // Вместо такого применим Copy-and-Swap idiom
            // delete[] str; size = s.size; str = new char[size]; memcpy(str, s.str, size);
            Swap_(other);    // подменяем себя на copy, локальная переменная copy уничтожается по окончанию ф-ции (она на стеке).
            return *this;
        }

        String operator+(const String& s) {
            char* new_str = new char[size + s.size];
            strcat(new_str, str);
            strcat(new_str, s.str);
            return String(new_str);
        }

        String& operator+=(const String& s) {
            *this = *this + s;
            return *this;
        }

        char* data() {
            return str;
        }

    private:
        void DebugInfo_(char* info) {
            sprintf(info, "DebugInfo: str: %s, sz: %lu", str, size);
        }

        char* DebugInfo2_() {
            char* info = (char*)malloc(50 * sizeof(char));
            DebugInfo_(info);
            return info;
        }

        void ll_(std::string info) {
            cout << info << endl;
        }

        // Пишем свой своп, потому что std::swap(*this, copy) всебе вызывает присваивание.
        void Swap_(String& s) {
            std::swap(str, s.str);
            std::swap(size, s.size);
            std::swap(cap, s.cap);
        }
};

String::~String() {
    char di [50];
    DebugInfo_(di);
    printf("Destructor ~String() - %s\n", di);
    delete[] str;
};


int main() {
    String s0;              // Default constructor
    String s1(3, '1');      // Parameterized constructor
    String s2 = s1;         // Copy constructor
    String s3{'s', '3'};    // Initializer list constructor
    String s4;
    s4 = s2 + s3;
    String s5("s5");
    s5 += s3;
    printf("s0: %s, s1: %s, s2: %s, s3: %s, s4: %s, s5: %s\n",
        s0.data(), s1.data(), s2.data(), s3.data(), s4.data(), s5.data());
    // std::cout << s2 << std::endl;
}

/**
     * The rule of three:
     * If you need to explicitly declare either the destructor, copy constructor or copy assignment operator yourself, you probably need to explicitly declare all three of them.
 */