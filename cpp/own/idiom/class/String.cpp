#include <algorithm>
#include <cstring>
#include <initializer_list>
#include <iostream>
#include <utility>
#include <cassert>

using std::cout;
using std::endl;

class String {
private:
  size_t size = 0;
  size_t cap = 0;
  char* str = nullptr;

public:
  // String() = default;
  // String() = delete;
  // String(String&) = default;
  // String(const String&) = default;
  // String(String&&) = default;
  // ~String() = default;

  ~String(); // Деструктор определим вне класса (просто так)

  String() {
    char* info = DebugInfo2();
    printf("String::String() - Default Constructor, %s\n", info);
    free(info);
  };

  String(const size_t sz, const char c = '\0')
      : size(sz)
      , cap(RoundUpToPowerOfTwo(sz + 1))
      , str(new char[cap]) // member initializer list
  {
    std::fill(str, str + size, c); // Или: memset(str, c, sz);
    str[size] = '\0';              // Препод говорит, что нет гарантии завершающего нулевого символа
    printf("String::String(size_t, const char) - Parameterized Constructor (Delegating), %s\n", DebugInfo2());
  }

  // Copy constructor
  String(const String& other)
      : String(other.size) // delegating constructor (since c++11)
  {
    std::copy(other.str, other.str + size + 1, str); // Или: memcpy(str, s.str, size + 1); // + 1 для завершающего нуля
    printf("String::String(const String&) - Copy Constructor, %s\n", DebugInfo2());
  }

  /**
   * Тип списка инициализации передаётся по значению, а не по ссылке,
   * потому что гарантируется, что это лёгкий тип и его копирование тривиально.
   * Это такая захардкоженная легковестная обёртка реализованная компилятором.
   */
  String(std::initializer_list<const char> lst)
      : String(lst.size()) // delegating constructor
  {
    std::copy(lst.begin(), lst.end(), str);
    printf("String::String(std::initializer_list<const char>) - Initializer List Constructor, %s\n", DebugInfo2());
  }

  String(const char* str)
      : String(strlen(str)) // delegating constructor
  {
    memcpy(this->str, str, size);
    printf("String::String(const char*) - Parameterized Constructor, %s\n", DebugInfo2());
  }

  // Move constructor
  String(String&& other) noexcept
      : size(other.size)
      , cap(other.cap) {
    if (this == &other) {
      return;
    }
    str = std::move(other.str);
    other.str = nullptr;
    other.size = 0;
    other.cap = 0;
    printf("String::String(String&&) - Move Constructor, %s\n", DebugInfo2());
  }

  // --- Operators ---

  /**
   * Когда нет move assignment оператора, тогда copy assignment оператор реализуется таким образом
   * Здесь происходит передача аргумента по значению вместо ссылки, с целью оптимизации copy elision
   * и т.к. нам всё равно необходимо делать копию other для свопа.
   * "if you're going to make a copy of something in a function, let the compiler do it in the parameter list"
   *
   * Но т.к. у нас реализован Move Assignment оператор и его сигнатура совпадает с этой,
   * то возникает CE "ambiguous overload for ‘operator=’"
   * поэтому copy assignment оператор реализован с другой сигнатурой
   *
  String& operator=(String other) noexcept {
      Ll("Copy assignment operator - String& operator=(String other)");
      if (this == &other) return *this;  // не имеет большого смысла при использовании Copy-and-swap().
      // this->~String(); // вызов деструктора это UB. После его вызова компилятор может считать, что объект всё.
      // Вместо такого применим Copy-and-Swap idiom
      // delete[] str; size = s.size; str = new char[size]; memcpy(str, s.str, size);
      Swap(other);    // подменяем себя на copy, локальная переменная copy уничтожается по окончанию ф-ции (она на стеке).
      return *this;
  }
  */

  // Copy assignment operator
  // Возвращает String&, чтобы мы могли писать цепочку присваивания. Если бы был void, то не могли бы.
  // Без ссылок вызывался бы конструктор копирования (и присваивание происходило уже в адрес другой переменной).
  // Параметр константная ссылка, чтобы ф-ция работала как с неконстантными, так и с константными аргументами.
  String& operator=(const String& other) noexcept {
    if (this == &other) {
      return *this; // не имеет большого смысла при использовании Copy-and-swap().
    }
    // this->~String(); // вызов деструктора это UB. После его вызова компилятор
    // может считать, что объект всё. Вместо такого применим Copy-and-Swap idiom
    // delete[] str; size = s.size; str = new char[size]; memcpy(str, s.str, size);
    String copy = other;
    Swap(copy); // подменяем себя на copy, локальная переменная copy уничтожается по окончанию ф-ции (она на стеке).
    printf("String::operator=(String other) - Copy assignment operator, %s\n", DebugInfo2());
    return *this;
  }

  // Move assignment
  String& operator=(String&& other) noexcept {
    if (this == &other) {
      return *this;
    }
    delete[] str;
    str = std::exchange(other.str, nullptr);
    size = std::exchange(other.size, 0);
    cap = std::exchange(other.cap, 0);
    printf("String::operator=(String&& other) - Move assignment operator, %s\n", DebugInfo2());
    return *this;
  }

  String operator+(const String& other) {
    size_t cap = RoundUpToPowerOfTwo(size + other.size + 1);
    char* new_str = new char[cap];
    strcat(new_str, str);
    strcat(new_str, other.str);
    printf("String::operator+(const String& other), %s\n", DebugInfo2());
    return String(new_str);
  }

  String& operator+=(const String& other) {
    *this = *this + other;
    return *this;
  }

  inline bool operator<(const String& other) const {
    if (size < other.size) {
      return true;
    }
    if (size > other.size) {
      return false;
    }
    return memcmp(str, other.str, size) < 0;
  }

  inline bool operator>(const String& other) const { return other < *this; }
  inline bool operator<=(const String& other) const { return !(*this > other); }
  inline bool operator>=(const String& other) const { return !(*this < other); }
  inline bool operator==(const String& other) const { return memcmp(str, other.str, size) == 0; }
  inline bool operator!=(const String& other) const { return !(*this == other); }
  
  inline char& operator[](const size_t index) { 
    printf("String::operator[%lu] (non-const), return: %c (full: %s)\n", index, str[index], str);
    return str[index];
  }
  inline const char& operator[](const size_t index) const { 
    printf("String::operator[%lu] const, return: %c (full: %s)\n", index, str[index], str);
    return str[index];
  }
  // inline char operator[](const size_t index) const { return str[index]; }

  String& operator->() {
    Ll("String& operator->()");
    return *this;
  }

  char* data() { return str; }

private:
  void DebugInfo(char* info) { sprintf(info, "DebugInfo: str: %s, sz: %lu, cap: %lu", str, size, cap); }

  char* DebugInfo2() {
    char* info = (char*)malloc(50 * sizeof(char));
    DebugInfo(info);
    return info;
  }

  void Ll(std::string info) { cout << info << " - " << DebugInfo2() << endl; }

  // Пишем свой своп, потому что std::swap(*this, copy) в себе вызывает присваивание.
  void Swap(String& other) {
    std::swap(str, other.str);
    std::swap(size, other.size);
    std::swap(cap, other.cap);
  }

  size_t RoundUpToPowerOfTwo(size_t n) {
    if (n <= 1) {
      return 1;
    }

    size_t result = 1;
    while (result < n) {
      result <<= 1;
    }

    return result;
  }
};

String::~String() {
  char di[150];
  DebugInfo(di);
  printf("String::~String() - Destructor, %s\n", di);
  delete[] str;
};

void Ll(std::string info) { cout << info << endl; }

int main() {
  Ll("main: Before - String s0;");
  String s0; // Default constructor
  Ll("main: Before - String s1(3, '1');");
  String s1(3, '1'); // Parameterized constructor
  Ll("main: Before - String s2 = s1;");
  String s2 = s1; // Copy constructor
  Ll("main: Before - String s3{'s', '3'};");
  String s3{'s', '3'}; // Initializer list constructor
  Ll("main: Before - String s4;");
  String s4;
  Ll("main: Before - s4 = s2 + s3;");
  s4 = s2 + s3; // s4.operator=(s2.operator+(s3));
  Ll("main: Before - String s5(\"s5\");");
  String s5("s5");
  Ll("main: Before - s5 += s3;");
  s5 += s3; // s5.operator+=(s3);
  Ll("main: Before - String s6(3, '6');");
  String s6(3, '6');
  Ll("main: Before - String s7(std::move(s6));");
  String s7(std::move(s6)); // Move constructor
  Ll("main: Before - s6 = std::move(s5);");
  s6 = std::move(s5); // Move assignment
  Ll("main: Before - String* s8 = new String(\"s8\");");
  String* s8 = new String("s8");

  printf("s0: %s, s1: %s, s2: %s, s3: %s, s4: %s, s5: %s, s6: %s, s6[2]: %c, "
         "s7: %s, s8: %s\n",
         s0.data(), s1.data(), s2.data(), s3.data(), s4.data(), s5.data(), s6.data(), s6[2], s7.data(), s8->data());

  String sc1("sc1");
  String sc2("sc2");
  String sc1again("sc1");
  printf("sc1 < sc2: %d, sc1 > sc2: %d, sc1 <= sc2: %d, sc1 >= sc2: %d, sc1 >= "
         "sc1again: %d, sc1 == sc1again: %d, sc1 != sc1again: %d\n",
         sc1<sc2, sc1> sc2, sc1 <= sc2, sc1 >= sc2, sc1 >= sc1again, sc1 == sc1again, sc1 != sc1again);

  String nc_str("nc_str");
  const String c_str("c_str");
  nc_str[1] = 'C';
  // c_str[1] = 'C'; // error: assignment of read-only location ‘c_str.String::operator[](1)’
  char& nc_char = nc_str[2];
  nc_char = '-';
  // char& c_char = c_str[2]; // error: binding reference of type ‘char&’ to ‘const char’ discards qualifiers
  const char& c_char = c_str[2];
  // c_char = '-'; // error: assignment of read-only reference ‘c_char’
  
  // Пример показывает, почему ф-ция возвращает char&, а не char,
  // eg: inline char operator[](const size_t index) const { return str[index]; }
  // изменение в строке должно влиять на переменную-ссылку с символом.
  const String& cr_str = nc_str;
  const char& c_char2 = cr_str[3];
  nc_str[3] = 'S';
  assert(c_char2 == 'S');
}

/**
 * The rule of three:
 * If you need to explicitly declare either the destructor, copy constructor or
 * copy assignment operator yourself, you probably need to explicitly declare
 * all three of them.
 * https://stackoverflow.com/questions/4172722/what-is-the-rule-of-three
 *
 * https://stackoverflow.com/questions/3279543/what-is-the-copy-and-swap-idiom
 *
 * https://en.cppreference.com/w/cpp/language/operators
 *
 * https://learn.microsoft.com/ru-ru/cpp/cpp/constructors-cpp?view=msvc-170
 */

// https://www.youtube.com/watch?v=FrCt1jcpXLA
