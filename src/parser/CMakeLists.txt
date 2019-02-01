cmake_minimum_required(VERSION 3.0.0)
project(tlangparsers VERSION 0.0.1)

set (CMAKE_CXX_STANDARD 11)
set (CMAKE_BUILD_TYPE Debug)

set (LLVM_DIR /usr/local/opt/llvm/lib/cmake/llvm)
find_package(LLVM REQUIRED CONFIG)

message(STATUS "Found LLVM ${LLVM_PACKAGE_VERSION}")
message(STATUS "Using LLVMConfig.cmake in: ${LLVM_DIR}")

include_directories(${LLVM_INCLUDE_DIRS})
add_definitions(${LLVM_DEFINITIONS})

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

llvm_map_components_to_libnames(llvm_libs support core irreader)
include_directories("cpp_test/parser")
include_directories("/usr/local/Cellar/antlr4-cpp-runtime/4.7.2/include/antlr4-runtime")
include_directories("/usr/local/include/antlr4-runtime")
target_link_libraries (tlangcppparser "/usr/local/lib/libantlr4-runtime.a" ${llvm_libs})
