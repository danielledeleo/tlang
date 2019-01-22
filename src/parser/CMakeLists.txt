cmake_minimum_required(VERSION 3.0.0)
project(tlangparsers VERSION 0.0.1)

set (CMAKE_CXX_STANDARD 11)

FILE(GLOB ANTLR_CPP cpp_test/parser/*.cpp)

add_custom_target(
    generate_antlr
    DEPENDS
        antlr.stamp
)

add_custom_command(
    OUTPUT
        antlr.stamp
    DEPENDS
        tlang.g4
    COMMAND 
        ./antlr_gen.sh
    COMMAND
        cmake -E touch antlr.stamp
    COMMENT
        "=== Generating ANTLR Parsers ==="
    VERBATIM
)

set(CPP_PARSER_SRCS
cpp_test/main.cpp
cpp_test/TTypes.cpp
${ANTLR_CPP}
)

add_executable(tlangcppparser ${CPP_PARSER_SRCS})
add_dependencies(tlangcppparser generate_antlr)

include_directories("cpp_test/parser")
include_directories("/usr/local/Cellar/antlr4-cpp-runtime/4.7.2/include/antlr4-runtime")
include_directories("/usr/local/include/antlr4-runtime")
target_link_libraries (tlangcppparser "/usr/local/lib/libantlr4-runtime.a")
