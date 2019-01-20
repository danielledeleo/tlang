
// Generated from Turing.g4 by ANTLR 4.7.2

#pragma once


#include "antlr4-runtime.h"




class  TuringLexer : public antlr4::Lexer {
public:
  enum {
    T__0 = 1, T__1 = 2, T__2 = 3, STRING_LITERAL = 4, INTEGER_LITERAL = 5, 
    END = 6, OF = 7, TO = 8, TYPE = 9, VAR = 10, PROCEDURE = 11, FUNCTION = 12, 
    CLASS = 13, PROCESS = 14, FOR = 15, LOOP = 16, EXIT = 17, IF = 18, ELSE = 19, 
    ELSIF = 20, CASE = 21, ASSERT = 22, BEGIN = 23, RETURN = 24, RESULT = 25, 
    NEW = 26, FREE = 27, TAG = 28, FORK = 29, SIGNAL = 30, WAIT = 31, PAUSE = 32, 
    QUIT = 33, UNCHECKED = 34, CHECKED = 35, EXPORT = 36, IMPORT = 37, INHERIT = 38, 
    IMPLEMENT = 39, BY = 40, PUT = 41, INT = 42, REAL = 43, BOOLEAN = 44, 
    STRING = 45, ARRAY = 46, SET = 47, RECORD = 48, UNION = 49, POINTER = 50, 
    NAT = 51, INTN = 52, NATN = 53, REALN = 54, CHAR = 55, DEFERRED = 56, 
    FORWARD = 57, BODY = 58, NOT = 59, CARET = 60, COLON = 61, L_PAREN = 62, 
    R_PAREN = 63, DOT = 64, RANGE = 65, COMMA = 66, CHEAT = 67, DEREFERENCE = 68, 
    ASSIGNMENT = 69, PLUS = 70, MINUS = 71, MULTIPLY = 72, DIVIDE = 73, 
    DIV = 74, MOD = 75, REM = 76, EXP = 77, LESSTHAN = 78, GREATERTHAN = 79, 
    EQUAL = 80, LESSTHANEQUAL = 81, GREATERTHANEQUAL = 82, NOTEQUAL = 83, 
    AND = 84, OR = 85, IMPLIES = 86, IN = 87, SHR = 88, SHL = 89, XOR = 90, 
    PLUSEQUALS = 91, MINUSEQUALS = 92, MULTIPLYEQUALS = 93, DIVIDEEQUALS = 94, 
    DIVEQUALS = 95, MODEQUALS = 96, ANDEQUALS = 97, OREQUALS = 98, IMPLIESEQUALS = 99, 
    SHREQUALS = 100, SHLEQUALS = 101, XOREQUALS = 102, IDENTIFIER = 103, 
    WHITESPACE = 104, BLOCK_COMMENT = 105, LINE_COMMENT = 106
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

