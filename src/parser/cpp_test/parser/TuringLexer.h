
// Generated from Turing.g4 by ANTLR 4.7.2

#pragma once


#include "antlr4-runtime.h"




class  TuringLexer : public antlr4::Lexer {
public:
  enum {
    T__0 = 1, T__1 = 2, T__2 = 3, STRING_LITERAL = 4, INTEGER_LITERAL = 5, 
    REAL_LITERAL = 6, END = 7, OF = 8, TO = 9, TYPE = 10, VAR = 11, PROCEDURE = 12, 
    FUNCTION = 13, CLASS = 14, PROCESS = 15, FOR = 16, LOOP = 17, EXIT = 18, 
    IF = 19, ELSE = 20, ELSIF = 21, CASE = 22, ASSERT = 23, BEGIN = 24, 
    RETURN = 25, RESULT = 26, NEW = 27, FREE = 28, TAG = 29, FORK = 30, 
    SIGNAL = 31, WAIT = 32, PAUSE = 33, QUIT = 34, UNCHECKED = 35, CHECKED = 36, 
    EXPORT = 37, IMPORT = 38, INHERIT = 39, IMPLEMENT = 40, BY = 41, PUT = 42, 
    INT = 43, REAL = 44, BOOLEAN = 45, STRING = 46, ARRAY = 47, SET = 48, 
    RECORD = 49, UNION = 50, POINTER = 51, NAT = 52, INTN = 53, NATN = 54, 
    REALN = 55, CHAR = 56, DEFERRED = 57, FORWARD = 58, BODY = 59, NOT = 60, 
    CARET = 61, COLON = 62, L_PAREN = 63, R_PAREN = 64, DOT = 65, RANGE = 66, 
    COMMA = 67, CHEAT = 68, DEREFERENCE = 69, ASSIGNMENT = 70, PLUS = 71, 
    MINUS = 72, MULTIPLY = 73, DIVIDE = 74, DIV = 75, MOD = 76, REM = 77, 
    EXP = 78, LESSTHAN = 79, GREATERTHAN = 80, EQUAL = 81, LESSTHANEQUAL = 82, 
    GREATERTHANEQUAL = 83, NOTEQUAL = 84, AND = 85, OR = 86, IMPLIES = 87, 
    IN = 88, SHR = 89, SHL = 90, XOR = 91, PLUSEQUALS = 92, MINUSEQUALS = 93, 
    MULTIPLYEQUALS = 94, DIVIDEEQUALS = 95, DIVEQUALS = 96, MODEQUALS = 97, 
    ANDEQUALS = 98, OREQUALS = 99, IMPLIESEQUALS = 100, SHREQUALS = 101, 
    SHLEQUALS = 102, XOREQUALS = 103, IDENTIFIER = 104, WHITESPACE = 105, 
    BLOCK_COMMENT = 106, LINE_COMMENT = 107
  };

  TuringLexer(antlr4::CharStream *input);
  ~TuringLexer();

  virtual std::string getGrammarFileName() const override;
  virtual const std::vector<std::string>& getRuleNames() const override;

  virtual const std::vector<std::string>& getChannelNames() const override;
  virtual const std::vector<std::string>& getModeNames() const override;
  virtual const std::vector<std::string>& getTokenNames() const override; // deprecated, use vocabulary instead
  virtual antlr4::dfa::Vocabulary& getVocabulary() const override;

  virtual const std::vector<uint16_t> getSerializedATN() const override;
  virtual const antlr4::atn::ATN& getATN() const override;

private:
  static std::vector<antlr4::dfa::DFA> _decisionToDFA;
  static antlr4::atn::PredictionContextCache _sharedContextCache;
  static std::vector<std::string> _ruleNames;
  static std::vector<std::string> _tokenNames;
  static std::vector<std::string> _channelNames;
  static std::vector<std::string> _modeNames;

  static std::vector<std::string> _literalNames;
  static std::vector<std::string> _symbolicNames;
  static antlr4::dfa::Vocabulary _vocabulary;
  static antlr4::atn::ATN _atn;
  static std::vector<uint16_t> _serializedATN;


  // Individual action functions triggered by action() above.

  // Individual semantic predicate functions triggered by sempred() above.

  struct Initializer {
    Initializer();
  };
  static Initializer _init;
};

