# Adding System Introspection (Tutorial Step 7)

 #cmake #system-introspection #tutorial

В этой задаче мы сначала в среде CMake проверяем присутствие библиотеки спомощью С кода. Если найдено, тогда устанавливаем переменные и передаём их в проект, где проверяем их наличие спомощью макросов предкомпиляции, и соответственно реагируем.

## Решение

Добавим в `./MathFunctions/CMakeLists.txt`

```sh
  include(CheckCXXSourceCompiles)

  # Use check_cxx_source_compiles with simple C++ code to verify availability of std::log() and std::exp().
  # Store the results in HAVE_LOG and HAVE_EXP respectively.
  check_cxx_source_compiles("
    #include <cmath>
    int main() {
      std::log(1.0);
      return 0;
    }
  " HAVE_LOG)
  check_cxx_source_compiles("
    #include <cmath>
    int main() {
      std::exp(1.0);
      return 0;
    }
  " HAVE_EXP)
  
  # Conditionally on HAVE_LOG and HAVE_EXP, add private compile
  # definitions "HAVE_LOG" and "HAVE_EXP" to the SqrtLibrary target.
  if(HAVE_LOG AND HAVE_EXP)
    target_compile_definitions(SqrtLibrary PRIVATE HAVE_LOG HAVE_EXP)
  endif()
```

В С++ коде `MathFunctions/mysqrt.cxx` проверим эти макро-переменные

```cpp
#include <cmath>
...
#if defined(HAVE_LOG) && defined(HAVE_EXP)
  double result = std::exp(std::log(x) * 0.5);
  std::cout << "Computing sqrt of " << x << " to be " << result
            << " using log and exp" << std::endl;
#else
  ...
#endif
```

## Links

* [Adding System Introspection (Tutorial Step 7)](https://cmake.org/cmake/help/latest/guide/tutorial/Adding%20System%20Introspection.html)
* [CheckCXXSourceCompiles](https://cmake.org/cmake/help/latest/module/CheckCXXSourceCompiles.html) ~ Проверьте, компилируется ли данный исходный код C++ и связывается ли он с исполняемым файлом.

---
