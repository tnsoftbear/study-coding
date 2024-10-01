#include <iostream>
#include <string>
#include <variant>
 
int main()
{
    std::variant<std::string> v1;
    v1.emplace<0>("abc"); // OK
    std::cout << std::get<0>(v1) << '\n';
    v1.emplace<std::string>("def"); // OK
    std::cout << std::get<std::string>(v1) << '\n';
 
    std::variant<std::string, std::string> v2;
    v2.emplace<1>("ghi"); // OK
    std::cout << std::get<1>(v2) << '\n';
    // v2.emplace<std::string>("abc"); -> Error

    std::variant<std::monostate, int, std::string> v3;
    std::cout << "v3.index(): " << v3.index() << '\n'; // v3.index(): 0
    v3 = "abc";
    std::cout << "v3.index(): " << v3.index() << '\n'; // v3.index(): 2
    std::cout << "std::holds_alternative<int>(v3): " << std::holds_alternative<int>(v3) 
        << ", std::holds_alternative<std::string>(v3): " << std::holds_alternative<std::string>(v3)
        << '\n'; // std::holds_alternative<int>(v3): 0, std::holds_alternative<std::string>(v3): 1
}