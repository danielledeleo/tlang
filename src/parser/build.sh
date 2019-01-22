antlr -Dlanguage=Cpp -listener -visitor Turing.g4 -o ./cpp_test/parser
antlr -Dlanguage=Go -listener -visitor Turing.g4 -o ./go_test/parser
