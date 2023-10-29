#include <memory>
#include <iostream>

struct Widget {
};

int main() {
    Widget widget;
    Widget* widget_ptr = &widget;
    std::shared_ptr sptr = std::make_shared<Widget*>(widget_ptr);
    std::cout << "widget ptr; " << widget_ptr << "; sptr: " << &sptr << std::endl;
    return 0;
}