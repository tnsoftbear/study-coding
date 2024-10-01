char* baz(int*& p);

// Распутаем сложное выражение
void (*bar(int x, void (*func)(int&)))(int&);

// Используем typedef:
typedef void (*ptr_to_fref)(int&);
ptr_to_fref bar(int x, ptr_to_fref func);

// Однако, typedef не позволяет создавать шаблонные алиасы, поэтому мы испольуем using
using ptr_to_fref = void (*)(int&);
ptr_to_fref bar(int x, ptr_to_fref func);

// С using можно создавать параметризованные синонимы
template <typename T> using ptr_to_fref2 = void (*)(T&);