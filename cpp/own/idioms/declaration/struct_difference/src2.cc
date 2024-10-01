#include <iostream>

#undef MYDEF
#include "header.h"

int foo(S *s);

int main() {
  S s = {1};
  std::cout << "s->y: " << foo(&s) << std::endl;
}

