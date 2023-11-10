#include <iostream>

class Refci {
  const int& ri = 1; // константная ссылка занимает размер указателя - 8 байт
};

static int si = 1;

class Refi {
  int& ri = si; // ссылка занимает размер указателя - 8 байт
};

// Т.к. ссылка не инициализирована, то необходим явный конструктор по-умолчанию (компилятор не может сгенерировать его)
class Refi2 {
  int& ri2;
  Refi2(int& ri2)
      : ri2(ri2){};
};

class Coni {
  const int ci = 1;
};

int globy = 5;
struct C {
  int& rx;
  const int& cr;
  int& ry = globy;
  C(int argx, int& argy) // Параметр должен быть ссылкой `int& x`, иначе предупреждение:
      : rx(argx)         //  warning: storing the address of local variable ‘argx’ in ‘*this.C::rx’
      , cr(5) // Так нельзя, несмотря на то, что `cr` - это константная ссылка, и она продлевает жизнь rvalue
              // warning: storing the address of local variable ‘<anonymous>’ in ‘*this.C::cr’
      , ry(argy) // Так правильно, инициализация ссылочной переменной в списке инициализации не повлияет на глобальный `globy`
  {
    // ry = argy;       // Так неправильно, будет изменена глобальная переменная `globy`, потому что она уже привязана к ссылке `ry`
  }
};

int main() {
  std::cout << "sizeof(Refci): " << sizeof(Refci)   // 8
            << ", sizeof(Refi): " << sizeof(Refi)   // 8
            << ", sizeof(Refi2): " << sizeof(Refi2) // 8
            << ", sizeof(Coni): " << sizeof(Coni)   // 4
            << std::endl;

  int xx = 10;
  C c(xx, xx);
  std::cout << "c.rx: " << c.rx // случайное значение
            << "; c.ry: " << c.ry
            << "; c.cr: " << c.cr // случайное значение
            << "; globy: " << globy << std::endl;
}

// g++ -Wall reference_members.cpp && ./a.out
