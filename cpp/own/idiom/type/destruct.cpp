#include <iostream>

using std::cout;
using std::endl;

class Bar {
  public:
    Bar() {
        cout << "Construct Bar()" << endl;
    }
    Bar(const Bar& b) {
        cout << "Copy Construct Bar(const Bar& b)" << endl;
    }
    ~Bar() {
        cout << "Destruct ~Bar()" << endl;
    }
    void echo() {
        cout << "Bar::echo()" << endl;
    }
};

void f(Bar* b) {
    cout << "f(Bar b)" << endl;    
}

int main() {
    Bar* b = new Bar();
    cout << "Before f(b)" << endl;
    f(b);
    b->echo();
    cout << "Before return" << endl;
    return 0;
}