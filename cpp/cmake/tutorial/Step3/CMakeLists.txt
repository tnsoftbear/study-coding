# Re-build and run:
# rm -rf build && cmake -B build && cmake --build build && ./build/Tutorial 123456

cmake_minimum_required(VERSION 3.10)

# set the project name and version
project(Tutorial VERSION 1.0)

# Create an interface library called tutorial_compiler_flags
add_library(tutorial_compiler_flags INTERFACE)

# Add compiler feature cxx_std_11 to tutorial_compiler_flags
# It replaces: set(CMAKE_CXX_STANDARD 11) and set(CMAKE_CXX_STANDARD_REQUIRED True)
target_compile_features(tutorial_compiler_flags INTERFACE cxx_std_11)

# configure a header file to pass some of the CMake settings to the source code
configure_file(TutorialConfig.h.in TutorialConfig.h)

# add the executable
add_executable(Tutorial tutorial.cxx)

# add the binary tree to the search path for include files so that we will find TutorialConfig.h
target_include_directories(Tutorial PUBLIC "${PROJECT_BINARY_DIR}")

# add the MathFunctions library
add_subdirectory(MathFunctions)

# Link Tutorial to tutorial_compiler_flags
target_link_libraries(Tutorial 
    PUBLIC MathFunctions
    tutorial_compiler_flags
)
