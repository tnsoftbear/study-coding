cmake_minimum_required(VERSION 3.15)
project(forward_list)
set(CMAKE_CXX_STANDARD 20)

add_compile_options(-fsanitize=address)
add_link_options(-fsanitize=address)

add_executable(forward_list forward_list.cpp)