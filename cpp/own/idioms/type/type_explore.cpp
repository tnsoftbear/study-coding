#include <typeinfo>
#include <iostream>

using std::cout;
using std::endl;

struct Empty {};
struct NonEmpty {
    Empty e;
};

class EmptyClass {};

namespace ec {
class NamespacedEmptyClass {};
}

//template <typename T> struct TD;

template<typename T> class Type2Type {};

template <typename T>
void foo (const T& t) {
    //TD<decltype(t)> t; // Показывает действительный тип error: declaration of ‘TD<const Empty* const&> t’ shadows a parameter
    Type2Type<decltype(t)> t2t;
    cout << typeid(t).name() << endl; // врёт
    cout << typeid(t2t).name() << endl; // выводит правильное деманглирование
    cout << endl; // врёт
}

int main () {
    const Empty* a;
    foo(a);
    // PK5Empty
    // 9Type2TypeIRKPK5EmptyE

    const NonEmpty* b;
    foo(b);
    // PK8NonEmpty
    // 9Type2TypeIRKPK8NonEmptyE

    const EmptyClass* c;
    foo(c);
    // PK10EmptyClass
    // 9Type2TypeIRKPK10EmptyClassE

    const ec::NamespacedEmptyClass* d;
    foo(d);
    // PKN2ec20NamespacedEmptyClassE
    // 9Type2TypeIRKPKN2ec20NamespacedEmptyClassEE

    ec::NamespacedEmptyClass* e;
    foo(e);
    // PN2ec20NamespacedEmptyClassE
    // 9Type2TypeIRKPN2ec20NamespacedEmptyClassEE
    

    // sudo pacman -S binutils
    // -t означает, деманглировать то, что выдал typeid(). (без -t деманглировать манглирование ассемблера)
    // echo "PKN2ec20NamespacedEmptyClassE" | c++filt -t
    // Вывод:
    // ec::NamespacedEmptyClass const*
    // Можно видеть, что стёрли ссылку.

    // Проверяем через decltype:
    // error: declaration of ‘TD<const ec::NamespacedEmptyClass* const&> t’ shadows a parameter

    // Проверим правильное деманглирование
    // echo "9Type2TypeIRKPN2ec20NamespacedEmptyClassEE" | c++filt -t
    // Вывод:
    // Type2Type<ec::NamespacedEmptyClass* const&>

    // [C++ lectures at MIPT (in Russian). Lecture 6. Rvalue references, part 1](https://www.youtube.com/watch?v=pjo8iZQWLMY&ab_channel=KonstantinVladimirov)
}