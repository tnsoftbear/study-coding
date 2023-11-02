#include <iostream>

using std::cout;
using std::endl;

class Bar {
public:
  Bar() { cout << "Bar::Bar() - def-ctor" << endl; }
  Bar(const Bar& b) { cout << "Bar::Bar(const Bar& b) - copy-ctor" << endl; }
  Bar(Bar&& b) { cout << "Bar::Bar(Bar&& b) - move-ctor" << endl; }
  Bar& operator=(const Bar& b) {
    cout << "Bar::operator=(const Bar& b) - copy-assign" << endl;
    return *this;
  }
  Bar& operator=(Bar&& b) {
    cout << "Bar::operator=(Bar&& b) - move-assign" << endl;
    return *this;
  }
  ~Bar() { cout << "Bar::~Bar() - dtor" << endl; }
  void echo() { cout << "Bar::echo()" << endl; }
};

void f(Bar b) { cout << "Bar::f(Bar b)" << endl; }

int main() {
  Bar* b = new Bar();
  Bar b1, b2;
  cout << "main: Before f(*b)" << endl;
  f(*b);
  cout << "main: Before f(std::move(*b))" << endl;
  f(std::move(*b));
  cout << "main: Before b->echo()" << endl;
  b->echo();
  cout << "main: Before Bar b2 = b1;" << endl;
  b2 = b1;
  cout << "main: Before b1 = std::move(b2);" << endl;
  b1 = std::move(b2);
  cout << "main: Before b2.echo()" << endl;
  b2.echo();
  cout << "main: Before delete" << endl;
  delete b; // Без delete, деструктор не вызывается для b
  cout << "main: Before return" << endl;
  return 0;
}