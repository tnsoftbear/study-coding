# TODO 1: Add a library called MathFunctions with sources MathFunctions.cxx
# and mysqrt.cxx
add_library(MathFunctions MathFunctions.cxx)

# TODO 7: Create a variable USE_MYMATH using option and set default to ON
option(USE_MYMATH "Use tutorial provided math implementatio" On)

if(USE_MYMATH)
    # TIf USE_MYMATH is ON, use target_compile_definitions to pass
    # USE_MYMATH as a precompiled definition to our source files
    target_compile_definitions(MathFunctions PRIVATE "USE_MYMATH")
    # Add a library for SqrtLibrary with source mysqrt.cxx
    add_library(SqrtLibrary mysqrt.cxx)
    # Link SqrtLibrary to the MathFunctions Library
    target_link_libraries(MathFunctions PUBLIC SqrtLibrary)
endif()


