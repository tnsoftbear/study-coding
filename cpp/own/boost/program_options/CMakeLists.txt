# rm -rf build && cmake -B build && cmake --build build && ./build/ProgramOptionsExample --help

cmake_minimum_required(VERSION 3.27)

project(ProgramOptionsExample CXX)

add_executable(${PROJECT_NAME} main.cpp)

# вызывает FindBoost.cmake https://cmake.org/cmake/help/latest/module/FindBoost.html
set(Boost_DEBUG ON)
find_package(Boost COMPONENTS program_options REQUIRED)
target_link_libraries(${PROJECT_NAME} ${Boost_LIBRARIES})
