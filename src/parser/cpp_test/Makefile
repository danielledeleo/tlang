OUTPUT=output
GENERATED=parser

# here is where you plug in the runtime for your OS
INCLUDES=/usr/local/include
CC=g++
CCARGS=-c -I $(INCLUDES)/antlr4-runtime/ -I $(GENERATED) -std=c++11
LDARGS=-g
LIBDIR=/usr/local/lib
LIBS=$(LIBDIR)/libantlr4-runtime.a
ANTLRGEN=BaseListener Lexer Listener Parser
OBJS=$(addsuffix .o,$(addprefix $(OUTPUT)/Turing,$(ANTLRGEN)))

all: main

main: $(OBJS) main.cpp
	$(CC) $(CCARGS) main.cpp  -o $(OUTPUT)/main.o 
	$(CC) $(LDARGS) $(OUTPUT)/main.o $(OUTPUT)/TuringBaseListener.o $(OUTPUT)/TuringLexer.o $(OUTPUT)/TuringListener.o $(OUTPUT)/TuringParser.o $(LIBS) -o tlangparser

$(OUTPUT)/%.o : $(GENERATED)/%.cpp
	$(CC) $(CCARGS) $< -o $@

$(GENERATED)/%.cpp: $(GENERATED)/;