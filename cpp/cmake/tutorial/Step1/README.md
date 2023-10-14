# Tutorial Step-1

## Run

```sh
rm -rf build && cmake -B build && cmake --build build && ./build/Tutorial
```

## Details

`configure_file()` скопировать файл в другое имя и изменить содержимое.

`#cmakedefine` работает плохо, используйте `#define`.

Я создал файл `TutorialConfig h.in`, который создает макропеременные на основе значений версии проекта.

```h
#define Tutorial_VMAJ "@Tutorial_VERSION_MAJOR@"
#define Tutorial_VMIN "@Tutorial_VERSION_MINOR@"
```

CMakeList.txt содержит:

```h
cmake_minimum_required(VERSION 3.10)
project(
    Tutorial
    VERSION 1.0
)
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)
configure_file(TutorialConfig.h.in TutorialConfig.h @ONLY)
add_executable(TutorialExe tutorial.cxx)
target_include_directories(TutorialExe PUBLIC "${PROJECT_BINARY_DIR}")
```

Это создаст файл `TutorialConfig.h` в директории сборки.

```h
#define Tutorial_VMAJ "1"
#define Tutorial_VMIN "0"
```

Можно иклудить `#include "TutorialConfig.h"` в код `tutorial.cxx` и использовать макро-переменные определяющие номер версии, заданные в `project()`.

## Links

* [CMake Tutorial Step-1](https://cmake.org/cmake/help/latest/guide/tutorial/A%20Basic%20Starting%20Point.html)
* [configure_file()](https://cmake.org/cmake/help/latest/command/configure_file.html)
* [target_include_directories()](https://cmake.org/cmake/help/latest/command/target_include_directories.html)
* [project()](https://cmake.org/cmake/help/latest/command/project.html)
* [set()](https://cmake.org/cmake/help/latest/command/set.html)

---
