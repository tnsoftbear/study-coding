cmake_minimum_required(VERSION 3.15)
project(Tutorial VERSION 1.0)

# specify the C++ standard
add_library(tutorial_compiler_flags INTERFACE)
target_compile_features(tutorial_compiler_flags INTERFACE cxx_std_11)

# Create a new variable gcc_like_cxx that is true if we are using CXX and
# any of the following compilers: ARMClang, AppleClang, Clang, GNU, LCC
SET(gcc_like_cxx $<COMPILE_LANG_AND_ID:CXX,ARMClang,AppleClang,Clang,GNU,LCC>)
# Create a new variable msvc_cxx that is true if we are using CXX and MSVC
SET(msvc_cxx $<COMPILE_LANG_AND_ID:CXX,MSVC>)

# Add warning flag compile options to the interface library tutorial_compiler_flags.
# BUILD_INTERFACE: Мы хотим, чтобы эти предупреждающие флаги использовались только во время сборок. 
# Потребители нашего установленного проекта не должны наследовать наши предупреждающие флаги.
target_compile_options(
    tutorial_compiler_flags INTERFACE
    $<${gcc_like_cxx}:$<BUILD_INTERFACE:-Wall;-Wextra;-Wshadow;-Wformat=2;-Wunused>>
    $<${msvc_cxx}:$<BUILD_INTERFACE:-W3>>
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
