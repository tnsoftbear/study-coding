/**
 * Посмотрим ассемблерный код:
 * g++ -O1 -g0 -masm=intel -S ptr_and_ref_in_func_param.cpp
 * Он одинаковый для обеих функций:
 *
        mov	eax, DWORD PTR [rdi]
        mov	DWORD PTR g[rip], eax
        ret
 *
 * ---
 * g адресуется относительно RIP, чтобы недалеко ходить за глобальной переменной.
 * Поэтому мы адресуемся не от начала памяти в 64-битном адресном пространстве, а от того места, где мы сейчас.
 */

int g;

void foo(int* x) { g = *x; }

void bar(int& x) { g = x; }

int main() {
  int x = 123;
  foo(&x);
  bar(x);
}