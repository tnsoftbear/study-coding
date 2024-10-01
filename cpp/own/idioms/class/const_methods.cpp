#include <iostream>

class C {
public:
  void NonConst() { std::cout << "Hi, I'm not const" << std::endl; }
  void IsConst() const { std::cout << "Hi, I'm const" << std::endl; }
  // Перегрузка по признаку константности самого объекта (а не аргумента)
  void IsConstNonConstOverload() { std::cout << "IsConstNonConstOverload() - non const" << std::endl; }
  void IsConstNonConstOverload() const { std::cout << "IsConstNonConstOverload() - is const" << std::endl; }
};

int main() {
  const C isConst;
  // isConst.NonConst(); // error: passing ‘const C’ as ‘this’ argument discards qualifiers
  isConst.IsConst();
  isConst.IsConstNonConstOverload();

  C nonConst;
  nonConst.NonConst();
  nonConst.IsConst();
  nonConst.IsConstNonConstOverload();

  const C& constRefToNonConst = nonConst;
  // constRefToNonConst.NonConst(); // error: passing ‘const C’ as ‘this’ argument discards qualifiers
}

/**
 Глубокое понимание константности, подразумевает, что константные объекты - это не то что нельзя менять,
 а это то у чего отсутствует часть операций, которые есть у обычного типа.
 Вывод: Все методы в классе, которые предполагаются вызываться и у константных объектов тоже, должны быть помечены как const.
 
 (с) Илья Мещерин [Лекция 15. Const, mutable, static и explicit в методах классов](https://www.youtube.com/watch?v=pzv8zlW4Kc4)
*/
