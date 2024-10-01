# Пример ошибки в типе

Типы с одним именем должны полексемно совпадать.
Реализуем компиляцию программы, где поле `y` пропадает из структуры `S` в `src2.cc`.

## Ручная сборка

```sh
# Компилируем объектные файлы
mkdir build
g++ -c src1.cc -o ./build/src1.o
g++ -c src2.cc -o ./build/src2.o
# Связываем объектные файлы в исполняемый
g++ ./build/src1.o ./build/src2.o -o ./build/prog
./build/prog
```

## Сборка CMake

```sh
rm -rf ./build && cmake -B build && cmake --build build && ./build/my_program
```