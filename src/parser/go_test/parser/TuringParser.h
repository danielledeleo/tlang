
// Generated from Turing.g4 by ANTLR 4.7.2

#pragma once


#include "antlr4-runtime.h"




class  TuringParser : public antlr4::Parser {
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

  enum {
    RuleProgram = 0, RuleTopLevel = 1, RuleProcHeader = 2, RuleFuncHeader = 3, 
    RuleSubprogramHeader = 4, RuleSubprogramDeclaration = 5, RuleSubprogramPrefix = 6, 
    RuleClassDeclaration = 7, RuleClassBody = 8, RuleExportStatement = 9, 
    RuleExportListItem = 10, RuleExportList = 11, RuleInheritStatement = 12, 
    RuleImplementStatement = 13, RuleIdOrFileItem = 14, RuleHowExport = 15, 
    RuleParamDeclaration = 16, RuleParamDeclarationList = 17, RuleProcBody = 18, 
    RuleStatementOrDeclaration = 19, RulePrimaryExpression = 20, RuleArgumentList = 21, 
    RuleExponentialExpression = 22, RulePointerExpression = 23, RulePostfixExpression = 24, 
    RulePrefixExpression = 25, RuleMultiplicativeExpression = 26, RuleMultiplicativeOperator = 27, 
    RuleAdditiveExpression = 28, RuleComparativeExpression = 29, RuleNotExpression = 30, 
    RuleAndExpression = 31, RuleOrExpression = 32, RuleImpliesExpression = 33, 
    RuleExpression = 34, RuleExpressionList = 35, RuleDeclaration = 36, 
    RuleStatements = 37, RuleAssignmentStatement = 38, RulePutStatement = 39, 
    RulePutItem = 40, RulePutItemList = 41, RuleBeginStatement = 42, RuleResultStatement = 43, 
    RuleNewStatement = 44, RuleFreeStatement = 45, RuleForkStatement = 46, 
    RuleTypeDeclaration = 47, RuleTypeSpec = 48, RuleIndexType = 49, RuleIndexTypeList = 50, 
    RuleStringType = 51, RuleRecordType = 52, RuleRecordField = 53, RuleVariableDeclaration = 54, 
    RuleVariableIdentifierList = 55, RuleVariableIdentifier = 56, RuleArrayDeclaration = 57, 
    RuleIdentifierList = 58, RuleLiteral = 59, RuleComment = 60
  };

  TuringParser(antlr4::TokenStream *input);
  ~TuringParser();

  virtual std::string getGrammarFileName() const override;
  virtual const antlr4::atn::ATN& getATN() const override { return _atn; };
  virtual const std::vector<std::string>& getTokenNames() const override { return _tokenNames; }; // deprecated: use vocabulary instead.
  virtual const std::vector<std::string>& getRuleNames() const override;
  virtual antlr4::dfa::Vocabulary& getVocabulary() const override;


  class ProgramContext;
  class TopLevelContext;
  class ProcHeaderContext;
  class FuncHeaderContext;
  class SubprogramHeaderContext;
  class SubprogramDeclarationContext;
  class SubprogramPrefixContext;
  class ClassDeclarationContext;
  class ClassBodyContext;
  class ExportStatementContext;
  class ExportListItemContext;
  class ExportListContext;
  class InheritStatementContext;
  class ImplementStatementContext;
  class IdOrFileItemContext;
  class HowExportContext;
  class ParamDeclarationContext;
  class ParamDeclarationListContext;
  class ProcBodyContext;
  class StatementOrDeclarationContext;
  class PrimaryExpressionContext;
  class ArgumentListContext;
  class ExponentialExpressionContext;
  class PointerExpressionContext;
  class PostfixExpressionContext;
  class PrefixExpressionContext;
  class MultiplicativeExpressionContext;
  class MultiplicativeOperatorContext;
  class AdditiveExpressionContext;
  class ComparativeExpressionContext;
  class NotExpressionContext;
  class AndExpressionContext;
  class OrExpressionContext;
  class ImpliesExpressionContext;
  class ExpressionContext;
  class ExpressionListContext;
  class DeclarationContext;
  class StatementsContext;
  class AssignmentStatementContext;
  class PutStatementContext;
  class PutItemContext;
  class PutItemListContext;
  class BeginStatementContext;
  class ResultStatementContext;
  class NewStatementContext;
  class FreeStatementContext;
  class ForkStatementContext;
  class TypeDeclarationContext;
  class TypeSpecContext;
  class IndexTypeContext;
  class IndexTypeListContext;
  class StringTypeContext;
  class RecordTypeContext;
  class RecordFieldContext;
  class VariableDeclarationContext;
  class VariableIdentifierListContext;
  class VariableIdentifierContext;
  class ArrayDeclarationContext;
  class IdentifierListContext;
  class LiteralContext;
  class CommentContext; 

  class  ProgramContext : public antlr4::ParserRuleContext {
  public:
    ProgramContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<TopLevelContext *> topLevel();
    TopLevelContext* topLevel(size_t i);

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ProgramContext* program();

  class  TopLevelContext : public antlr4::ParserRuleContext {
  public:
    TopLevelContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    StatementOrDeclarationContext *statementOrDeclaration();
    ClassDeclarationContext *classDeclaration();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  TopLevelContext* topLevel();

  class  ProcHeaderContext : public antlr4::ParserRuleContext {
  public:
    ProcHeaderContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *IDENTIFIER();
    antlr4::tree::TerminalNode *PROCEDURE();
    antlr4::tree::TerminalNode *PROCESS();
    antlr4::tree::TerminalNode *L_PAREN();
    antlr4::tree::TerminalNode *R_PAREN();
    ParamDeclarationListContext *paramDeclarationList();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ProcHeaderContext* procHeader();

  class  FuncHeaderContext : public antlr4::ParserRuleContext {
  public:
    FuncHeaderContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *FUNCTION();
    std::vector<antlr4::tree::TerminalNode *> IDENTIFIER();
    antlr4::tree::TerminalNode* IDENTIFIER(size_t i);
    antlr4::tree::TerminalNode *COLON();
    TypeSpecContext *typeSpec();
    antlr4::tree::TerminalNode *L_PAREN();
    antlr4::tree::TerminalNode *R_PAREN();
    ParamDeclarationListContext *paramDeclarationList();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  FuncHeaderContext* funcHeader();

  class  SubprogramHeaderContext : public antlr4::ParserRuleContext {
  public:
    SubprogramHeaderContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    FuncHeaderContext *funcHeader();
    ProcHeaderContext *procHeader();
    SubprogramPrefixContext *subprogramPrefix();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  SubprogramHeaderContext* subprogramHeader();

  class  SubprogramDeclarationContext : public antlr4::ParserRuleContext {
  public:
    SubprogramDeclarationContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    SubprogramHeaderContext *subprogramHeader();
    ProcBodyContext *procBody();
    antlr4::tree::TerminalNode *END();
    antlr4::tree::TerminalNode *IDENTIFIER();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  SubprogramDeclarationContext* subprogramDeclaration();

  class  SubprogramPrefixContext : public antlr4::ParserRuleContext {
  public:
    SubprogramPrefixContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *DEFERRED();
    antlr4::tree::TerminalNode *BODY();
    antlr4::tree::TerminalNode *FORWARD();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  SubprogramPrefixContext* subprogramPrefix();

  class  ClassDeclarationContext : public antlr4::ParserRuleContext {
  public:
    ClassDeclarationContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *CLASS();
    std::vector<antlr4::tree::TerminalNode *> IDENTIFIER();
    antlr4::tree::TerminalNode* IDENTIFIER(size_t i);
    antlr4::tree::TerminalNode *END();
    std::vector<ClassBodyContext *> classBody();
    ClassBodyContext* classBody(size_t i);

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ClassDeclarationContext* classDeclaration();

  class  ClassBodyContext : public antlr4::ParserRuleContext {
  public:
    ClassBodyContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    ExportStatementContext *exportStatement();
    InheritStatementContext *inheritStatement();
    ImplementStatementContext *implementStatement();
    StatementOrDeclarationContext *statementOrDeclaration();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ClassBodyContext* classBody();

  class  ExportStatementContext : public antlr4::ParserRuleContext {
  public:
    ExportStatementContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *EXPORT();
    ExportListContext *exportList();
    antlr4::tree::TerminalNode *L_PAREN();
    antlr4::tree::TerminalNode *R_PAREN();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ExportStatementContext* exportStatement();

  class  ExportListItemContext : public antlr4::ParserRuleContext {
  public:
    ExportListItemContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *IDENTIFIER();
    HowExportContext *howExport();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ExportListItemContext* exportListItem();

  class  ExportListContext : public antlr4::ParserRuleContext {
  public:
    ExportListContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    ExportListItemContext *exportListItem();
    ExportListContext *exportList();
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ExportListContext* exportList();
  ExportListContext* exportList(int precedence);
  class  InheritStatementContext : public antlr4::ParserRuleContext {
  public:
    InheritStatementContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *INHERIT();
    IdOrFileItemContext *idOrFileItem();
    antlr4::tree::TerminalNode *L_PAREN();
    antlr4::tree::TerminalNode *R_PAREN();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  InheritStatementContext* inheritStatement();

  class  ImplementStatementContext : public antlr4::ParserRuleContext {
  public:
    ImplementStatementContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *IMPLEMENT();
    IdOrFileItemContext *idOrFileItem();
    antlr4::tree::TerminalNode *BY();
    antlr4::tree::TerminalNode *L_PAREN();
    antlr4::tree::TerminalNode *R_PAREN();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ImplementStatementContext* implementStatement();

  class  IdOrFileItemContext : public antlr4::ParserRuleContext {
  public:
    IdOrFileItemContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *IDENTIFIER();
    antlr4::tree::TerminalNode *IN();
    antlr4::tree::TerminalNode *STRING_LITERAL();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  IdOrFileItemContext* idOrFileItem();

  class  HowExportContext : public antlr4::ParserRuleContext {
  public:
    HowExportContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *VAR();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  HowExportContext* howExport();

  class  ParamDeclarationContext : public antlr4::ParserRuleContext {
  public:
    ParamDeclarationContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    IdentifierListContext *identifierList();
    antlr4::tree::TerminalNode *COLON();
    TypeSpecContext *typeSpec();
    antlr4::tree::TerminalNode *VAR();
    SubprogramHeaderContext *subprogramHeader();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ParamDeclarationContext* paramDeclaration();

  class  ParamDeclarationListContext : public antlr4::ParserRuleContext {
  public:
    ParamDeclarationListContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    ParamDeclarationContext *paramDeclaration();
    ParamDeclarationListContext *paramDeclarationList();
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ParamDeclarationListContext* paramDeclarationList();
  ParamDeclarationListContext* paramDeclarationList(int precedence);
  class  ProcBodyContext : public antlr4::ParserRuleContext {
  public:
    ProcBodyContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<StatementOrDeclarationContext *> statementOrDeclaration();
    StatementOrDeclarationContext* statementOrDeclaration(size_t i);

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ProcBodyContext* procBody();

  class  StatementOrDeclarationContext : public antlr4::ParserRuleContext {
  public:
    StatementOrDeclarationContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    StatementsContext *statements();
    DeclarationContext *declaration();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  StatementOrDeclarationContext* statementOrDeclaration();

  class  PrimaryExpressionContext : public antlr4::ParserRuleContext {
  public:
    PrimaryExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    LiteralContext *literal();
    antlr4::tree::TerminalNode *IDENTIFIER();
    antlr4::tree::TerminalNode *L_PAREN();
    ExpressionContext *expression();
    antlr4::tree::TerminalNode *R_PAREN();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  PrimaryExpressionContext* primaryExpression();

  class  ArgumentListContext : public antlr4::ParserRuleContext {
  public:
    ArgumentListContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    ExpressionContext *expression();
    std::vector<ArgumentListContext *> argumentList();
    ArgumentListContext* argumentList(size_t i);
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ArgumentListContext* argumentList();
  ArgumentListContext* argumentList(int precedence);
  class  ExponentialExpressionContext : public antlr4::ParserRuleContext {
  public:
    ExponentialExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<PrimaryExpressionContext *> primaryExpression();
    PrimaryExpressionContext* primaryExpression(size_t i);
    antlr4::tree::TerminalNode *EXP();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ExponentialExpressionContext* exponentialExpression();

  class  PointerExpressionContext : public antlr4::ParserRuleContext {
  public:
    PointerExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *CARET();
    PrimaryExpressionContext *primaryExpression();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  PointerExpressionContext* pointerExpression();

  class  PostfixExpressionContext : public antlr4::ParserRuleContext {
  public:
    PostfixExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    PrimaryExpressionContext *primaryExpression();
    PointerExpressionContext *pointerExpression();
    ExponentialExpressionContext *exponentialExpression();
    PostfixExpressionContext *postfixExpression();
    antlr4::tree::TerminalNode *L_PAREN();
    antlr4::tree::TerminalNode *R_PAREN();
    ArgumentListContext *argumentList();
    antlr4::tree::TerminalNode *DOT();
    antlr4::tree::TerminalNode *IDENTIFIER();
    antlr4::tree::TerminalNode *DEREFERENCE();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  PostfixExpressionContext* postfixExpression();
  PostfixExpressionContext* postfixExpression(int precedence);
  class  PrefixExpressionContext : public antlr4::ParserRuleContext {
  public:
    PrefixExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    PostfixExpressionContext *postfixExpression();
    antlr4::tree::TerminalNode *PLUS();
    PrefixExpressionContext *prefixExpression();
    antlr4::tree::TerminalNode *MINUS();
    antlr4::tree::TerminalNode *CHEAT();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  PrefixExpressionContext* prefixExpression();

  class  MultiplicativeExpressionContext : public antlr4::ParserRuleContext {
  public:
    MultiplicativeExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    PostfixExpressionContext *postfixExpression();
    PrefixExpressionContext *prefixExpression();
    MultiplicativeExpressionContext *multiplicativeExpression();
    MultiplicativeOperatorContext *multiplicativeOperator();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  MultiplicativeExpressionContext* multiplicativeExpression();
  MultiplicativeExpressionContext* multiplicativeExpression(int precedence);
  class  MultiplicativeOperatorContext : public antlr4::ParserRuleContext {
  public:
    MultiplicativeOperatorContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *MULTIPLY();
    antlr4::tree::TerminalNode *DIVIDE();
    antlr4::tree::TerminalNode *DIV();
    antlr4::tree::TerminalNode *MOD();
    antlr4::tree::TerminalNode *REM();
    antlr4::tree::TerminalNode *SHR();
    antlr4::tree::TerminalNode *SHL();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  MultiplicativeOperatorContext* multiplicativeOperator();

  class  AdditiveExpressionContext : public antlr4::ParserRuleContext {
  public:
    AdditiveExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    MultiplicativeExpressionContext *multiplicativeExpression();
    AdditiveExpressionContext *additiveExpression();
    antlr4::tree::TerminalNode *PLUS();
    antlr4::tree::TerminalNode *MINUS();
    antlr4::tree::TerminalNode *XOR();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  AdditiveExpressionContext* additiveExpression();
  AdditiveExpressionContext* additiveExpression(int precedence);
  class  ComparativeExpressionContext : public antlr4::ParserRuleContext {
  public:
    ComparativeExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    AdditiveExpressionContext *additiveExpression();
    ComparativeExpressionContext *comparativeExpression();
    antlr4::tree::TerminalNode *LESSTHAN();
    antlr4::tree::TerminalNode *GREATERTHAN();
    antlr4::tree::TerminalNode *EQUAL();
    antlr4::tree::TerminalNode *LESSTHANEQUAL();
    antlr4::tree::TerminalNode *GREATERTHANEQUAL();
    antlr4::tree::TerminalNode *NOTEQUAL();
    antlr4::tree::TerminalNode *IN();
    antlr4::tree::TerminalNode *NOT();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ComparativeExpressionContext* comparativeExpression();
  ComparativeExpressionContext* comparativeExpression(int precedence);
  class  NotExpressionContext : public antlr4::ParserRuleContext {
  public:
    NotExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    ComparativeExpressionContext *comparativeExpression();
    antlr4::tree::TerminalNode *NOT();
    NotExpressionContext *notExpression();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  NotExpressionContext* notExpression();

  class  AndExpressionContext : public antlr4::ParserRuleContext {
  public:
    AndExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    NotExpressionContext *notExpression();
    AndExpressionContext *andExpression();
    antlr4::tree::TerminalNode *AND();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  AndExpressionContext* andExpression();
  AndExpressionContext* andExpression(int precedence);
  class  OrExpressionContext : public antlr4::ParserRuleContext {
  public:
    OrExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    AndExpressionContext *andExpression();
    OrExpressionContext *orExpression();
    antlr4::tree::TerminalNode *OR();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  OrExpressionContext* orExpression();
  OrExpressionContext* orExpression(int precedence);
  class  ImpliesExpressionContext : public antlr4::ParserRuleContext {
  public:
    ImpliesExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    OrExpressionContext *orExpression();
    ImpliesExpressionContext *impliesExpression();
    antlr4::tree::TerminalNode *IMPLIES();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ImpliesExpressionContext* impliesExpression();
  ImpliesExpressionContext* impliesExpression(int precedence);
  class  ExpressionContext : public antlr4::ParserRuleContext {
  public:
    ExpressionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    ImpliesExpressionContext *impliesExpression();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ExpressionContext* expression();

  class  ExpressionListContext : public antlr4::ParserRuleContext {
  public:
    ExpressionListContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    ExpressionContext *expression();
    ExpressionListContext *expressionList();
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ExpressionListContext* expressionList();
  ExpressionListContext* expressionList(int precedence);
  class  DeclarationContext : public antlr4::ParserRuleContext {
  public:
    DeclarationContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    TypeDeclarationContext *typeDeclaration();
    VariableDeclarationContext *variableDeclaration();
    ArrayDeclarationContext *arrayDeclaration();
    SubprogramDeclarationContext *subprogramDeclaration();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  DeclarationContext* declaration();

  class  StatementsContext : public antlr4::ParserRuleContext {
  public:
    StatementsContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    ExpressionContext *expression();
    AssignmentStatementContext *assignmentStatement();
    PutStatementContext *putStatement();
    antlr4::tree::TerminalNode *EXIT();
    BeginStatementContext *beginStatement();
    antlr4::tree::TerminalNode *RETURN();
    ResultStatementContext *resultStatement();
    NewStatementContext *newStatement();
    FreeStatementContext *freeStatement();
    ForkStatementContext *forkStatement();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  StatementsContext* statements();

  class  AssignmentStatementContext : public antlr4::ParserRuleContext {
  public:
    AssignmentStatementContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    PostfixExpressionContext *postfixExpression();
    ExpressionContext *expression();
    antlr4::tree::TerminalNode *ASSIGNMENT();
    antlr4::tree::TerminalNode *PLUSEQUALS();
    antlr4::tree::TerminalNode *MINUSEQUALS();
    antlr4::tree::TerminalNode *MULTIPLYEQUALS();
    antlr4::tree::TerminalNode *DIVIDEEQUALS();
    antlr4::tree::TerminalNode *DIVEQUALS();
    antlr4::tree::TerminalNode *SHLEQUALS();
    antlr4::tree::TerminalNode *SHREQUALS();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  AssignmentStatementContext* assignmentStatement();

  class  PutStatementContext : public antlr4::ParserRuleContext {
  public:
    PutStatementContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *PUT();
    PutItemListContext *putItemList();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  PutStatementContext* putStatement();

  class  PutItemContext : public antlr4::ParserRuleContext {
  public:
    PutItemContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    StatementsContext *statements();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  PutItemContext* putItem();

  class  PutItemListContext : public antlr4::ParserRuleContext {
  public:
    PutItemListContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    PutItemContext *putItem();
    PutItemListContext *putItemList();
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  PutItemListContext* putItemList();
  PutItemListContext* putItemList(int precedence);
  class  BeginStatementContext : public antlr4::ParserRuleContext {
  public:
    BeginStatementContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *BEGIN();
    StatementOrDeclarationContext *statementOrDeclaration();
    antlr4::tree::TerminalNode *END();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  BeginStatementContext* beginStatement();

  class  ResultStatementContext : public antlr4::ParserRuleContext {
  public:
    ResultStatementContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *RESULT();
    ExpressionContext *expression();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ResultStatementContext* resultStatement();

  class  NewStatementContext : public antlr4::ParserRuleContext {
  public:
    NewStatementContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *NEW();
    std::vector<antlr4::tree::TerminalNode *> IDENTIFIER();
    antlr4::tree::TerminalNode* IDENTIFIER(size_t i);
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  NewStatementContext* newStatement();

  class  FreeStatementContext : public antlr4::ParserRuleContext {
  public:
    FreeStatementContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *FREE();
    std::vector<antlr4::tree::TerminalNode *> IDENTIFIER();
    antlr4::tree::TerminalNode* IDENTIFIER(size_t i);
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  FreeStatementContext* freeStatement();

  class  ForkStatementContext : public antlr4::ParserRuleContext {
  public:
    ForkStatementContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *FORK();
    ExpressionContext *expression();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ForkStatementContext* forkStatement();

  class  TypeDeclarationContext : public antlr4::ParserRuleContext {
  public:
    TypeDeclarationContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *TYPE();
    antlr4::tree::TerminalNode *IDENTIFIER();
    antlr4::tree::TerminalNode *COLON();
    TypeSpecContext *typeSpec();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  TypeDeclarationContext* typeDeclaration();

  class  TypeSpecContext : public antlr4::ParserRuleContext {
  public:
    TypeSpecContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *INT();
    antlr4::tree::TerminalNode *REAL();
    antlr4::tree::TerminalNode *BOOLEAN();
    antlr4::tree::TerminalNode *NAT();
    antlr4::tree::TerminalNode *INTN();
    antlr4::tree::TerminalNode *NATN();
    antlr4::tree::TerminalNode *REALN();
    antlr4::tree::TerminalNode *CHAR();
    IndexTypeContext *indexType();
    StringTypeContext *stringType();
    RecordTypeContext *recordType();
    antlr4::tree::TerminalNode *IDENTIFIER();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  TypeSpecContext* typeSpec();

  class  IndexTypeContext : public antlr4::ParserRuleContext {
  public:
    IndexTypeContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<antlr4::tree::TerminalNode *> INTEGER_LITERAL();
    antlr4::tree::TerminalNode* INTEGER_LITERAL(size_t i);
    antlr4::tree::TerminalNode *RANGE();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  IndexTypeContext* indexType();

  class  IndexTypeListContext : public antlr4::ParserRuleContext {
  public:
    IndexTypeListContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    IndexTypeContext *indexType();
    IndexTypeListContext *indexTypeList();
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  IndexTypeListContext* indexTypeList();
  IndexTypeListContext* indexTypeList(int precedence);
  class  StringTypeContext : public antlr4::ParserRuleContext {
  public:
    StringTypeContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *STRING();
    antlr4::tree::TerminalNode *L_PAREN();
    antlr4::tree::TerminalNode *INTEGER_LITERAL();
    antlr4::tree::TerminalNode *R_PAREN();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  StringTypeContext* stringType();

  class  RecordTypeContext : public antlr4::ParserRuleContext {
  public:
    RecordTypeContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<antlr4::tree::TerminalNode *> RECORD();
    antlr4::tree::TerminalNode* RECORD(size_t i);
    antlr4::tree::TerminalNode *END();
    std::vector<RecordFieldContext *> recordField();
    RecordFieldContext* recordField(size_t i);

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  RecordTypeContext* recordType();

  class  RecordFieldContext : public antlr4::ParserRuleContext {
  public:
    RecordFieldContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    IdentifierListContext *identifierList();
    antlr4::tree::TerminalNode *COLON();
    TypeSpecContext *typeSpec();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  RecordFieldContext* recordField();

  class  VariableDeclarationContext : public antlr4::ParserRuleContext {
  public:
    VariableDeclarationContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *VAR();
    VariableIdentifierListContext *variableIdentifierList();
    antlr4::tree::TerminalNode *COLON();
    TypeSpecContext *typeSpec();
    antlr4::tree::TerminalNode *ASSIGNMENT();
    ExpressionContext *expression();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  VariableDeclarationContext* variableDeclaration();

  class  VariableIdentifierListContext : public antlr4::ParserRuleContext {
  public:
    VariableIdentifierListContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    VariableIdentifierContext *variableIdentifier();
    VariableIdentifierListContext *variableIdentifierList();
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  VariableIdentifierListContext* variableIdentifierList();
  VariableIdentifierListContext* variableIdentifierList(int precedence);
  class  VariableIdentifierContext : public antlr4::ParserRuleContext {
  public:
    VariableIdentifierContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *IDENTIFIER();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  VariableIdentifierContext* variableIdentifier();

  class  ArrayDeclarationContext : public antlr4::ParserRuleContext {
  public:
    ArrayDeclarationContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *ARRAY();
    IndexTypeContext *indexType();
    antlr4::tree::TerminalNode *OF();
    TypeSpecContext *typeSpec();
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  ArrayDeclarationContext* arrayDeclaration();

  class  IdentifierListContext : public antlr4::ParserRuleContext {
  public:
    IdentifierListContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *IDENTIFIER();
    IdentifierListContext *identifierList();
    antlr4::tree::TerminalNode *COMMA();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  IdentifierListContext* identifierList();
  IdentifierListContext* identifierList(int precedence);
  class  LiteralContext : public antlr4::ParserRuleContext {
  public:
    LiteralContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *STRING_LITERAL();
    antlr4::tree::TerminalNode *INTEGER_LITERAL();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  LiteralContext* literal();

  class  CommentContext : public antlr4::ParserRuleContext {
  public:
    CommentContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *BLOCK_COMMENT();
    antlr4::tree::TerminalNode *LINE_COMMENT();

    virtual void enterRule(antlr4::tree::ParseTreeListener *listener) override;
    virtual void exitRule(antlr4::tree::ParseTreeListener *listener) override;
   
  };

  CommentContext* comment();


  virtual bool sempred(antlr4::RuleContext *_localctx, size_t ruleIndex, size_t predicateIndex) override;
  bool exportListSempred(ExportListContext *_localctx, size_t predicateIndex);
  bool paramDeclarationListSempred(ParamDeclarationListContext *_localctx, size_t predicateIndex);
  bool argumentListSempred(ArgumentListContext *_localctx, size_t predicateIndex);
  bool postfixExpressionSempred(PostfixExpressionContext *_localctx, size_t predicateIndex);
  bool multiplicativeExpressionSempred(MultiplicativeExpressionContext *_localctx, size_t predicateIndex);
  bool additiveExpressionSempred(AdditiveExpressionContext *_localctx, size_t predicateIndex);
  bool comparativeExpressionSempred(ComparativeExpressionContext *_localctx, size_t predicateIndex);
  bool andExpressionSempred(AndExpressionContext *_localctx, size_t predicateIndex);
  bool orExpressionSempred(OrExpressionContext *_localctx, size_t predicateIndex);
  bool impliesExpressionSempred(ImpliesExpressionContext *_localctx, size_t predicateIndex);
  bool expressionListSempred(ExpressionListContext *_localctx, size_t predicateIndex);
  bool putItemListSempred(PutItemListContext *_localctx, size_t predicateIndex);
  bool indexTypeListSempred(IndexTypeListContext *_localctx, size_t predicateIndex);
  bool variableIdentifierListSempred(VariableIdentifierListContext *_localctx, size_t predicateIndex);
  bool identifierListSempred(IdentifierListContext *_localctx, size_t predicateIndex);

private:
  static std::vector<antlr4::dfa::DFA> _decisionToDFA;
  static antlr4::atn::PredictionContextCache _sharedContextCache;
  static std::vector<std::string> _ruleNames;
  static std::vector<std::string> _tokenNames;

  static std::vector<std::string> _literalNames;
  static std::vector<std::string> _symbolicNames;
  static antlr4::dfa::Vocabulary _vocabulary;
  static antlr4::atn::ATN _atn;
  static std::vector<uint16_t> _serializedATN;


  struct Initializer {
    Initializer();
  };
  static Initializer _init;
};

