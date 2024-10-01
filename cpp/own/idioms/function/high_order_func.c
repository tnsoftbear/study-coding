// Примеры ф-ций высшего порядка на языке С
// gcc high_order_func.c && ./a.out

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// объявим "int_action_t" - тип указателя на ф-цию, которая принимает целое и ничего не возвращает
typedef void (*int_action_t)(int);

/**
  --- a) Ф-ция высшего порядка принимает аргументом другую ф-цию ---
 */
void int_for_each(int* b, int* e, int_action_t f) {
  while (b != e) { // указатели в роли итераторов
    f(*b++);
  }
}

void int_print(int x) { printf("%d ", x); }

/**
  --- b) Свёркта по бинарной операции ---
  Недостаток алгоритма int_for_each - каждый элемент массива обрабатывается сам по себе,
  невозможно при обработке очередного элемента использовать результат обработки предыдущего.
 */
typedef int (*int_binop_t)(int, int);
int int_fold(int* b, int* e, int_binop_t f, int a) {
  while (b != e) {
    a = f(a, *b++);
  }
  return a;
}
int int_maxof(int x, int y) { return x > y ? x : y; }
int int_plus(int x, int y) { return x + y; }
int int_mult(int x, int y) { return x * y; }

/**
  --- c) Поэлементное применение ф-ций "наоборот" ---
  Вычисляется результат цепочки применения ф-ций к переменной x: fn(.. f1(f0(x)))
  Цепочку можно формировать динамически из имеющихся ф-ций.
  Можно видеть, что такие шаблоны как "композит" и "цепочка обязанностей" реализуются на языке С без объектов и классов.
 */
typedef int (*int_unop_t)(int);
int int_apply_chain(int_unop_t* b, int_unop_t* e, int x) {
  while (b != e) { // b, e - указатели на указатели на функции
    x = (*b++)(x);
  }
  return x;
}
int f1(int x) { return x * 2; }
int f2(int x) { return x + 1; }

/**
   --- d) Передача состояни ---
   Передавать объект данных между итерациями, который хранит промежуточное состояние.
   Обработка каждого очередного элемента изменяет это состояние.
 */
typedef void(int_stateful_op_t)(int x, void* s);
void int_apply_stateful(int* b, int* e, int_stateful_op_t f, void* s) {
  while (b != e) {
    f(*b++, s);
  }
}
typedef struct {
  int count;
  int sum;
} int_average_state_t;
void int_count_and_sum(int x, void* p) {
  int_average_state_t* state = (int_average_state_t*)p;
  ++state->count;
  state->sum += x;
}
double int_average(int* b, int* e) {
  int_average_state_t state = {0, 0};
  int_apply_stateful(b, e, int_count_and_sum, &state);
  return ((double)state.sum) / state.count;
}

/**
   --- e) Ф-ции высшего порядка в стандартной библиотеке
 */
int str_comparer(void const *p, void const *q) {
    char const *s = *(char const **)p;
    char const *t = *(char const **)q;
    return strcmp(s, t);
}

int main() {
  // a) Ф-ция высшего порядка принимает аргументом другую ф-цию
  int arr[] = {1, 2, 3, 4, 5, 8, 13, 21};
  printf("a) ");
  int_for_each(arr, arr + 8, int_print);
  printf("\n");
  // ---------------------------------------------------------------------

  // b) Свёркта по бинарной операции
  int max = int_fold(arr, arr + 8, int_maxof, 0);
  int sum = int_fold(arr, arr + 8, int_plus, 0);
  int product = int_fold(arr, arr + 8, int_mult, 1);
  printf("b) max: %d, sum: %d, product: %d\n", max, sum, product);
  // ---------------------------------------------------------------------

  // c) Поэлементное применение ф-ций
  int_unop_t funcs[] = {f1, f2, f1, f2, f1, f2};
  int res = int_apply_chain(funcs, funcs + 6, 10);
  printf("c) res: %d\n", res);
  // ---------------------------------------------------------------------

  // d) Вычисление с состоянием
  printf("d) average: %f\n", int_average(arr, arr + 8));
  // ---------------------------------------------------------------------

  // e) Ф-ции высшего порядка в стандартной библиотеке
  char *words[] = {"one", "two", "three", "four", "five", "six", "seven", "eight"};
  qsort(words, 8, sizeof(char *), str_comparer);
  printf("e) ");
  for (int i = 0; i < 8; ++i) {
    printf("%s ", words[i]);
  }
  printf("\n");
  // ---------------------------------------------------------------------
}
