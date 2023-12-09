#include <iostream>
#include <limits>
#include <vector>

bool calc(const std::vector<int>& nums, const int& target, const int sum, const int idx) {
  if (idx == nums.size()) {
    // printf("idx: %d, sum: %d\n", idx, sum);
    return sum == target;
  }
  // printf("idx: %d, sum: %d, nums[idx]: %d\n", idx, sum, nums[idx]);
  return calc(nums, target, sum + nums[idx], idx + 1) || calc(nums, target, sum - nums[idx], idx + 1);
}

bool solve(const std::vector<int>& nums, const int& target) {
  if (nums.size() == 0) {
    return false;
  }
  return calc(nums, target, nums[0], 1);
}

int main() {
  std::vector<int> nums;

  std::cout << "Введите числа через пробел:" << std::endl;
  int num;
  while (std::cin >> num) {
    nums.push_back(num);
    if (std::cin.peek() == '\n') {
      break;
    }
  }

  std::cout << "Введите целевое число: " << std::endl;
  int target;
  std::cin >> target;

  std::cout << (solve(nums, target) ? "true" : "false") << std::endl;
}

/**
Первая задача из видео [#59 | Адилет Жаксыбай - Разбор задач по программированию и алгоритмам для попадания в Google
(Гугл)](https://www.youtube.com/watch?v=gbgiFVFhGkc)

## Условие

Дан массив целых чисел nums и целое число target.
Необходимо выяснить можно ли используя сложение и вычитание на массиве чисел nums получить значение target.

### Пример

nums: 9 3 7, target: 12

### Решение рекурсивным перебором (Backtracking)

9
9+3, 9-3
9+3+7=5, 9+3-7=19, 9-3+7=13, 9-3-7=-1
Ответ: false (значение 12 не получено)

Time complexity: O(2^N) - по сути мы проходим бинарное дерево высотою N (слева сумма при сложении с элементом, справа при вычитании элемента массива num): 1 + 2
+ 4 ... + 2^(N-1) = 2^N - 1 Space complexity: O(N) - это высота дерева. Стек вызовов увеличивается пока мы идём до последнего элемента.

## Отладка в gdb

g++ -g -Og find_sum_backtracking.cpp
gdb ./a.out
set args < test2.txt
display $rsp  # выводить значение RSP регистра после каждой команды
b 7
b 11
start
bt            # стек вызовов с текущей функцией вверх по цепочке вызовов.
c             # выполнять до следующего брейкпоинта
frame 0       # конкретный уровень стека
info args     # значения всех аргументов функции на текущем уровне стека
info locals   # значения всех локальных переменных на текущем уровне стека
info reg rsp
set output-radix 10 # отображать регистры в десятичной система счисления
info breakpoints

## Другие алгоритмы решения

* Knapsack DP
* "Backtracking" with sets

*/
