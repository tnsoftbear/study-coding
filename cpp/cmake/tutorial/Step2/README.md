# CMake add_library (Tutorial Step-2)

 #cmake #add-library #tutorial

## Часть 1

Редактируем базовый `CMakeList.txt`

* `add_subdirectory(MathFunctions)` - Зарегистрировали под-директорию MathFunctions, в ней в `MathFunctions/CMakeLists.txt` объявлена библиотека: `add_library(MathFunctions MathFunctions.cxx mysqrt.cxx)`.  
* `target_link_libraries(Tutorial PUBLIC MathFunctions)` - Слинковали её к исполняемому файлу.
* Добавили путь библиотеки к цели спомощью `target_include_directories()`. Можно было написать отдельным вызовом `target_include_directories(Tutorial PUBLIC "${PROJECT_SOURCE_DIR}/MathFunctions")`.

```sh
cmake_minimum_required(VERSION 3.10)
project(Tutorial VERSION 1.0)
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)
configure_file(TutorialConfig.h.in TutorialConfig.h)
add_subdirectory(MathFunctions)
add_executable(Tutorial tutorial.cxx)
target_link_libraries(Tutorial PUBLIC MathFunctions) # Link the library to our executable

# Add the binary tree to the search path for include files so that we will find TutorialConfig.h
# Add MathFunctions to Tutorial's target_include_directories()
target_include_directories(
    Tutorial PUBLIC 
    "${PROJECT_BINARY_DIR}"
    "${PROJECT_SOURCE_DIR}/MathFunctions"
)
```

## Links

* [add_library](https://cmake.org/cmake/help/latest/command/add_library.html) ~ Добавьте в проект библиотеку, используя указанные исходные файлы.
* [target_include_directories](https://cmake.org/cmake/help/latest/command/target_include_directories.html) ~ Указывает каталоги включения, которые будут использоваться при компиляции заданной цели.
* [add_subdirectory](https://cmake.org/cmake/help/latest/command/add_subdirectory.html) ~ Добавляет подкаталог в сборку
* [target_link_libraries](https://cmake.org/cmake/help/latest/command/target_link_libraries.html) ~ Укажите библиотеки или флаги, которые будут использоваться при связывании данной цели и/или ее зависимых объектов.
* [cmake-buildsystem](https://cmake.org/cmake/help/latest/manual/cmake-buildsystem.7.html)

## Часть 2

Добавим в библиотеку `MathFunctions` опцию `USE_MYMATH`, позволяющую разработчикам выбирать либо собственную реализацию квадратного корня, либо встроенную стандартную реализацию.

```sh
# Исключим mysqrt.cxx из библиотеки MathFunctions, этот источник кода будет добавляться только в случае указания USE_MYMATH
add_library(MathFunctions MathFunctions.cxx)
option(USE_MYMATH "Use tutorial provided math implementatio" On) # Можно: option(USE_MYMATH On)
if(USE_MYMATH)
    # Передать USE_MYMATH как прекомпилированное определение в код
    target_compile_definitions(MathFunctions PRIVATE "USE_MYMATH")
    # When USE_MYMATH is ON, add a library for SqrtLibrary with source mysqrt.cxx
    add_library(SqrtLibrary mysqrt.cxx)
    # Link SqrtLibrary to the MathFunctions Library
    target_link_libraries(MathFunctions PUBLIC SqrtLibrary)
endif()
```

В C++ коде `MathFunctions.cxx` использование опции `USE_MYMATH` выглядит так:

```cpp
#include "MathFunctions.h"

#ifdef USE_MYMATH
  #include "mysqrt.h"
#else
  #include <cmath>
#endif

namespace mathfunctions {
  double sqrt(double x)
  {
#ifdef USE_MYMATH
    return detail::mysqrt(x);
#else
    return std::sqrt(x);
#endif
  }
}
```

```sh
# Собрать с опцией USE_MYMATH=OFF
rm -rf build && cmake -B build -DUSE_MYMATH=OFF && cmake --build build && ./build/Tutorial 1234567891
```

## Links

* [Step 2: Adding a Library](https://cmake.org/cmake/help/latest/guide/tutorial/Adding%20a%20Library.html)
* [if](https://cmake.org/cmake/help/latest/command/if.html)
* [option](https://cmake.org/cmake/help/latest/command/option.html)
* [target_compile_definitions](https://cmake.org/cmake/help/latest/command/target_compile_definitions.html) ~ Добавьте определения компиляции в цель.

---
