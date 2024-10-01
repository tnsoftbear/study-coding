#include "my_shared_ptr.cpp"

using std::cout;
using std::endl;

struct Widget {
  int data;
  Widget(int data)
      : data(data) {}
  ~Widget() { cout << "~Widget()" << endl; }
};

int main() {
  cout << "main: Before SharedPtr<Widget> p = MakeShared<Widget>(5);" << endl;
  SharedPtr<Widget> p = MakeShared<Widget>(5);
  cout << "main: Before SharedPtr<Widget> p2 = p;" << endl;
  SharedPtr<Widget> p2 = p;
  cout << "main: Before SharedPtr<Widget> p3 = std::move(p);" << endl;
  auto p3 = std::move(p);
  cout << "main: Before SharedPtr<Widget> p4;" << endl;
  SharedPtr<Widget> p4;
  cout << "main: Before p4 = p3;" << endl;
  p4 = p3;
  cout << "main: Before return" << endl;
}
