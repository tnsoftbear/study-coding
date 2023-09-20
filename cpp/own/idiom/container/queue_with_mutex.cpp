#include <cstdio>
#include <mutex>
#include <queue>

int main() {
  std::mutex m1, m2;
  std::queue<std::mutex*> q;
  q.push(&m1);
  q.push(&m2);
  printf("%lu\n", q.size());
  q.pop();
  q.pop();
}