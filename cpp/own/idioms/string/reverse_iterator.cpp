#include <string>
#include <iostream>

using std::cout;
using std::endl;

int main() {
  std::string str("microsoft");
  std::string::reverse_iterator r;
  for (r = str.rbegin(); r < str.rend(); r++) {
    cout << *r;
  }
  cout << endl;
  return 0;
}