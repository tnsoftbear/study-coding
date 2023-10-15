cmake_minimum_required(VERSION 3.15)
project(Tutorial VERSION 1.0)

# specify the C++ standard
add_library(tutorial_compiler_flags INTERFACE)
target_compile_features(tutorial_compiler_flags INTERFACE cxx_std_11)

# add compiler warning flags just when building this project via the BUILD_INTERFACE genex
set(gcc_like_cxx "$<COMPILE_LANG_AND_ID:CXX,ARMClang,AppleClang,Clang,GNU,LCC>")
set(msvc_cxx "$<COMPILE_LANG_AND_ID:CXX,MSVC>")
target_compile_options(tutorial_compiler_flags INTERFACE
  "$<${gcc_like_cxx}:$<BUILD_INTERFACE:-Wall;-Wextra;-Wshadow;-Wformat=2;-Wunused>>"
  "$<${msvc_cxx}:$<BUILD_INTERFACE:-W3>>"
)

# configure a header file to pass some of the CMake settings
# to the source code
configure_file(TutorialConfig.h.in TutorialConfig.h)

# add the MathFunctions library
add_subdirectory(MathFunctions)

# add the executable
add_executable(Tutorial tutorial.cxx)

target_link_libraries(Tutorial PUBLIC MathFunctions tutorial_compiler_flags)

# add the binary tree to the search path for include files
# so that we will find TutorialConfig.h
target_include_directories(Tutorial PUBLIC "${PROJECT_BINARY_DIR}")

# Install Tutorial in the bin directory
install(TARGETS Tutorial DESTINATION bin)

# Install TutorialConfig.h to the include directory
install(FILES ${PROJECT_BINARY_DIR}/TutorialConfig.h DESTINATION include)

# Enable testing
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

# Create a function called do_test to avoid copy + paste.
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