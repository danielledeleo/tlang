
// Generated from Turing.g4 by ANTLR 4.7.2

#pragma once


#include "antlr4-runtime.h"




class  TuringLexer : public antlr4::Lexer {
public:
  enum {
    T__0 = 1, T__1 = 2, T__2 = 3, NOT_IN = 4, STRING_LITERAL = 5, INTEGER_LITERAL = 6, 
    REAL_LITERAL = 7, END = 8, OF = 9, TO = 10, TYPE = 11, VAR = 12, PROCEDURE = 13, 
    FUNCTION = 14, CLASS = 15, PROCESS = 16, FOR = 17, LOOP = 18, EXIT = 19, 
    IF = 20, ELSE = 21, ELSIF = 22, CASE = 23, ASSERT = 24, BEGIN = 25, 
    RETURN = 26, RESULT = 27, NEW = 28, FREE = 29, TAG = 30, FORK = 31, 
    SIGNAL = 32, WAIT = 33, PAUSE = 34, QUIT = 35, UNCHECKED = 36, CHECKED = 37, 
    EXPORT = 38, IMPORT = 39, INHERIT = 40, IMPLEMENT = 41, BY = 42, PUT = 43, 
    INT = 44, REAL = 45, BOOLEAN = 46, STRING = 47, ARRAY = 48, SET = 49, 
    RECORD = 50, UNION = 51, POINTER = 52, NAT = 53, INTN = 54, NATN = 55, 
    REALN = 56, CHAR = 57, DEFERRED = 58, FORWARD = 59, BODY = 60, NOT = 61, 
    CARET = 62, COLON = 63, L_PAREN = 64, R_PAREN = 65, DOT = 66, RANGE = 67, 
    COMMA = 68, CHEAT = 69, DEREFERENCE = 70, ASSIGNMENT = 71, PLUS = 72, 
    MINUS = 73, MULTIPLY = 74, DIVIDE = 75, DIV = 76, MOD = 77, REM = 78, 
    EXP = 79, LESSTHAN = 80, GREATERTHAN = 81, EQUAL = 82, LESSTHANEQUAL = 83, 
    GREATERTHANEQUAL = 84, NOTEQUAL = 85, AND = 86, OR = 87, IMPLIES = 88, 
    IN = 89, SHR = 90, SHL = 91, XOR = 92, PLUSEQUALS = 93, MINUSEQUALS = 94, 
    MULTIPLYEQUALS = 95, DIVIDEEQUALS = 96, DIVEQUALS = 97, MODEQUALS = 98, 
    ANDEQUALS = 99, OREQUALS = 100, IMPLIESEQUALS = 101, SHREQUALS = 102, 
    SHLEQUALS = 103, XOREQUALS = 104, IDENTIFIER = 105, WHITESPACE = 106, 
    BLOCK_COMMENT = 107, LINE_COMMENT = 108
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

