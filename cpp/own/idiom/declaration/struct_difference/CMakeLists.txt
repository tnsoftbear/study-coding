cmake_minimum_required(VERSION 3.0)
project(MyProgram)

# Компиляция src1.cc в объектный файл
add_library(src1 OBJECT src1.cc)

# Компиляция src2.cc в объектный файл
add_library(src2 OBJECT src2.cc)

# Связывание объектных файлов в исполняемый файл
add_executable(my_program $<TARGET_OBJECTS:src1> $<TARGET_OBJECTS:src2>)