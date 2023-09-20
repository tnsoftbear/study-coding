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

        // Move assignment
        String& operator=(String&& other) noexcept {
            ll_("Move assignment operator - String& operator=(String&& other)");
            if (this == &other) return *this;
            delete[] str;
            str = std::exchange(other.str, nullptr);
            size = std::exchange(other.size, 0);
            cap = std::exchange(other.cap, 0);
            return *this;
        }

        /**
         * Когда нет move assignment оператора, тогда copy assignment оператор реализуется таким образом
         * Здесь происходит передача аргумента по значению вместо ссылки, с целью оптимизации copy elision 
         * и т.к. нам всё равно необходимо делать копию other для свопа.
         * "if you're going to make a copy of something in a function, let the compiler do it in the parameter list"
         * 
         * Но т.к. у нас реализован move assignment оператор и его сигнатура совпадает с этой, то возникает CE "ambiguous overload for ‘operator=’"
         * поэтому copy assignment оператор реализован с другой сигнатурой
         * 
        String& operator=(String other) noexcept {
            ll_("Copy assignment operator - String& operator=(String other)");
            if (this == &other) return *this;  // не имеет большого смысла при использовании Copy-and-swap().
            // this->~String(); // вызов деструктора это UB. После его вызова компилятор может считать, что объект всё.
            // Вместо такого применим Copy-and-Swap idiom
            // delete[] str; size = s.size; str = new char[size]; memcpy(str, s.str, size);
            Swap_(other);    // подменяем себя на copy, локальная переменная copy уничтожается по окончанию ф-ции (она на стеке).
            return *this;
        }
         */

        // Copy assignment operator
        String& operator=(const String& other) noexcept {
            ll_("Copy assignment operator - String& operator=(String other)");
            if (this == &other) return *this;  // не имеет большого смысла при использовании Copy-and-swap().
            // this->~String(); // вызов деструктора это UB. После его вызова компилятор может считать, что объект всё.
            // Вместо такого применим Copy-and-Swap idiom
            // delete[] str; size = s.size; str = new char[size]; memcpy(str, s.str, size);
            String copy = other;
            Swap_(copy);    // подменяем себя на copy, локальная переменная copy уничтожается по окончанию ф-ции (она на стеке).
            return *this;
        }

        String operator+(const String& other) {
            char* new_str = new char[size + other.size];
            strcat(new_str, str);
            strcat(new_str, other.str);
            return String(new_str);
        }

        String& operator+=(const String& s) {
            *this = *this + s;
            return *this;
        }

        inline bool operator<(const String& other) const {
            if (size < other.size) return true;
            if (size > other.size) return false;
            return memcmp(str, other.str, size) < 0;
        }

        inline bool operator>(const String& other) const { 
            return other < *this; 
        }

        inline bool operator<=(const String& other) const {
            return !(*this > other);
        }

        inline bool operator>=(const String& other) const {
            return !(*this < other);
        }

        inline bool operator==(const String& other) const {
            return memcmp(str, other.str, size) == 0;
        }

        inline bool operator!=(const String& other) const {
            return !(*this == other);
        }

        inline char operator[](const int index) {
            return str[index];
        }

        String& operator->() {
            ll_("String& operator->()");
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
            cout << info << " - " << DebugInfo2_() << endl;
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
    s4 = s2 + s3;           // s4.operator=(s2.operator+(s3));
    String s5("s5");
    s5 += s3;               // s5.operator+=(s3);
    String s6(3, '6');
    s6 = std::move(s5);
    String* s7 = new String("s7");
    
    printf("s0: %s, s1: %s, s2: %s, s3: %s, s4: %s, s5: %s, s6: %s, s6[2]: %c, s7: %s\n",
        s0.data(), s1.data(), s2.data(), s3.data(), s4.data(), s5.data(), s6.data(), s6[2], s7->data());
    
    String sc1("sc1");
    String sc2("sc2");
    String sc1again("sc1");
    printf("sc1 < sc2: %d, sc1 > sc2: %d, sc1 <= sc2: %d, sc1 >= sc2: %d, sc1 >= sc1again: %d, sc1 == sc1again: %d, sc1 != sc1again: %d\n",
        sc1 < sc2, sc1 > sc2, sc1 <= sc2, sc1 >= sc2, sc1 >= sc1again, sc1 == sc1again, sc1 != sc1again);
}

/**
     * The rule of three:
     * If you need to explicitly declare either the destructor, copy constructor or copy assignment operator yourself, you probably need to explicitly declare all three of them.
     * https://stackoverflow.com/questions/4172722/what-is-the-rule-of-three
     * 
     * https://stackoverflow.com/questions/3279543/what-is-the-copy-and-swap-idiom
     * 
     * https://en.cppreference.com/w/cpp/language/operators
 */