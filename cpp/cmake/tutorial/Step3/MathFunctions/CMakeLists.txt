add_library(MathFunctions MathFunctions.cxx)
target_link_libraries(MathFunctions PUBLIC tutorial_compiler_flags)

# By INTERFACE usage requirement we state that anybody linking to MathFunctions 
# needs to include the current source directory, while MathFunctions itself doesn't.
# INTERFACE means things that consumers require but the producer doesn't.
target_include_directories(MathFunctions INTERFACE "$(CMAKE_CURRENT_SOURCE_DIR)")

# should we use our own math functions
option(USE_MYMATH "Use tutorial provided math implementation" ON)

if(USE_MYMATH)
  target_compile_definitions(MathFunctions PRIVATE "USE_MYMATH")
  target_link_libraries(MathFunctions PRIVATE SqrtLibrary)

  # Add library that just does sqrt and link it to tutorial_compiler_flags
  add_library(SqrtLibrary STATIC mysqrt.cxx) 
  target_link_libraries(SqrtLibrary PUBLIC tutorial_compiler_flags)
endif(USE_MYMATH)
