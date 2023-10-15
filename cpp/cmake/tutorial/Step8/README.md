# CMake Step 8: Adding a Custom Command and Generated File

 #cmake #custom-command #tutorial

Мы хотим создать таблицу предварительно вычисленных значений для использования в mysqrtфункции. Она генерируется файлом `MathFunctions/MakeTable.cxx` в файл `Table.h`.

Создадим файл `MathFunctions/MakeTable.cmake`. Таблица данных будет сгенерирована в `build/MathFunctions/Table.h`.

```sh
add_executable(MakeTable MakeTable.cxx)
target_link_libraries(MakeTable PRIVATE tutorial_compiler_flags)

# добавляем специальную команду, которая определяет, как производить работу Table.h с помощью запуска MakeTable
add_custom_command(
  OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/Table.h
  COMMAND MakeTable ${CMAKE_CURRENT_BINARY_DIR}/Table.h
  DEPENDS MakeTable
)
```

В `CMakeLists.txt` добавим инструкции установки

```sh
include(MakeTable.cmake)

  add_library(SqrtLibrary STATIC mysqrt.cxx
    # сообщить CMake, что mysqrt.cxx зависит от сгенерированного файла Table.h
    ${CMAKE_CURRENT_BINARY_DIR}/Table.h
  )
  # добавить текущий двоичный каталог в список подключаемых каталогов, чтобы его Table.h можно было найти и включить с помощью mysqrt.cxx
  target_include_directories(SqrtLibrary PRIVATE ${CMAKE_CURRENT_BINARY_DIR})
```

Далее переписать `MathFunctions/mysqrt.cxx` для использования таблицы и включить её `#include "Table.h"`.

## Run

```sh
rm -rf build && cmake -B build && cmake --build build && build/Tutorial 25
```

## Links

* [Step 8: Adding a Custom Command and Generated File](https://cmake.org/cmake/help/latest/guide/tutorial/Adding%20a%20Custom%20Command%20and%20Generated%20File.html)
* [add_custom_command()](https://cmake.org/cmake/help/latest/command/add_custom_command.html) ~ Добавьте пользовательское правило сборки в созданную систему сборки.

---
