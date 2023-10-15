# Tutorial Step 5: Installing and Testing

 #cmake #tutorial #install #test #ctest

## Часть 1 - инсталяция

### Run

```sh
rm -rf {build,install} \
&& cmake -B build \
&& cmake --build build \
&& cmake --install ./build --prefix install \
&& install/bin/Tutorial 123456
```

### Решение

В `./MathFunctions/CMakeLists.txt` добавим инструкции установки

```sh
set(installable_libs MathFunctions tutorial_compiler_flags)
if(TARGET SqrtLibrary)
  list(APPEND installable_libs SqrtLibrary)
endif()
install(TARGETS ${installable_libs} DESTINATION lib)
install(FILES Mathfunctions.h mysqrt.h DESTINATION include)
```

В `./CMakeLists.txt` добавим инструкции установки

```sh
install(TARGETS Tutorial DESTINATION bin)
install(FILES ${PROJECT_BINARY_DIR}/TutorialConfig.h DESTINATION include)
```

### Links

* [Step 5: Installing and Testing](https://cmake.org/cmake/help/latest/guide/tutorial/Installing%20and%20Testing.html)
* [install()](https://cmake.org/cmake/help/latest/command/install.html) ~ Укажите правила, которые будут выполняться во время установки.

## Часть 2 - Тестирование

### Run

```sh
rm -rf build && cmake -B build && cmake --build build && cd build && ctest -VV
```

### Решение

Добавьте тесты в `CMakeLists.txt`

```sh
enable_testing()

add_test(NAME Runs COMMAND Tutorial 25)

add_test(NAME Usage COMMAND Tutorial)
set_tests_properties(Usage PROPERTIES
  PASS_REGULAR_EXPRESSION "Usage.*number"
)

add_test(NAME Usage_2 COMMAND Tutorial 4)
set_tests_properties(Usage_2 PROPERTIES
  PASS_REGULAR_EXPRESSION "4 is 2"
)

function(do_test input expected)
  add_test(NAME Usage_${input} COMMAND Tutorial ${input})
  set_tests_properties(Usage_${input} PROPERTIES
    PASS_REGULAR_EXPRESSION "${input} is ${expected}"
  )
endfunction(do_test)
do_test(4 2)
do_test(9 3)
do_test(25 5)
do_test(0.0001 0.01)
do_test(5 2.23607)
do_test(7 2.64575)
do_test(-25 0)
```

### Links

* [Step 5: Installing and Testing](https://cmake.org/cmake/help/latest/guide/tutorial/Installing%20and%20Testing.html)
* [add_test()](https://cmake.org/cmake/help/latest/command/add_test.html) ~ Добавьте тест в проект, который будет запускаться `ctest`.
* [enable_testing()](https://cmake.org/cmake/help/latest/command/enable_testing.html) ~ Включите тестирование для текущего каталога и ниже.
* [PASS_REGULAR_EXPRESSION](https://cmake.org/cmake/help/latest/prop_test/PASS_REGULAR_EXPRESSION.html)
* [ctest](https://cmake.org/cmake/help/latest/manual/ctest.1.html)
* [set_tests_properties()](https://cmake.org/cmake/help/latest/command/set_tests_properties.html)
* [function()](https://cmake.org/cmake/help/latest/command/function.html)

---
