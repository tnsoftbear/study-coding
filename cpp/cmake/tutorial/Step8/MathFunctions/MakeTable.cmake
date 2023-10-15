add_executable(MakeTable MakeTable.cxx)
target_link_libraries(MakeTable PRIVATE tutorial_compiler_flags)

# добавляем специальную команду, которая определяет, как производить работу Table.h с помощью запуска MakeTable
add_custom_command(
  OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/Table.h
  COMMAND MakeTable ${CMAKE_CURRENT_BINARY_DIR}/Table.h
  DEPENDS MakeTable
)