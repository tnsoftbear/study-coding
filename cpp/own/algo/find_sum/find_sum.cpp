#include <iostream>
#include <limits>
#include <vector>

bool calc(std::vector<int>& nums, int& target, int sum, int idx) {
  if (idx == nums.size()) {
    printf("idx: %d, sum: %d\n", idx, sum);
    return sum == target;
  }
  printf("idx: %d, sum: %d, nums[idx]: %d\n", idx, sum, nums[idx]);
  return calc(nums, target, sum + nums[idx], idx + 1) || calc(nums, target, sum - nums[idx], idx + 1);
}

bool solve(std::vector<int>& nums, int& target) {
  if (nums.size() == 0) {
    return false;
  }
  return calc(nums, target, 0, 0);
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
Первая задача из видео [#59 | Адилет Жаксыбай - Разбор задач по программированию и алгоритмам для попадания в Google (Гугл)](https://www.youtube.com/watch?v=gbgiFVFhGkc)

g++ -g -Og find_sum.cpp
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
*/
