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
