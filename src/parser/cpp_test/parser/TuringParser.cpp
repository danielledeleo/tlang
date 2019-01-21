
// Generated from Turing.g4 by ANTLR 4.7.2


#include "TuringListener.h"
#include "TuringVisitor.h"

#include "TuringParser.h"


using namespace antlrcpp;
using namespace antlr4;

TuringParser::TuringParser(TokenStream *input) : Parser(input) {
  _interpreter = new atn::ParserATNSimulator(this, _atn, _decisionToDFA, _sharedContextCache);
}

TuringParser::~TuringParser() {
  delete _interpreter;
}

std::string TuringParser::getGrammarFileName() const {
  return "Turing.g4";
}

const std::vector<std::string>& TuringParser::getRuleNames() const {
  return _ruleNames;
}

dfa::Vocabulary& TuringParser::getVocabulary() const {
  return _vocabulary;
}


//----------------- ProgramContext ------------------------------------------------------------------

TuringParser::ProgramContext::ProgramContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<TuringParser::TopLevelContext *> TuringParser::ProgramContext::topLevel() {
  return getRuleContexts<TuringParser::TopLevelContext>();
}

TuringParser::TopLevelContext* TuringParser::ProgramContext::topLevel(size_t i) {
  return getRuleContext<TuringParser::TopLevelContext>(i);
}


size_t TuringParser::ProgramContext::getRuleIndex() const {
  return TuringParser::RuleProgram;
}

void TuringParser::ProgramContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterProgram(this);
}

void TuringParser::ProgramContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitProgram(this);
}


antlrcpp::Any TuringParser::ProgramContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitProgram(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ProgramContext* TuringParser::program() {
  ProgramContext *_localctx = _tracker.createInstance<ProgramContext>(_ctx, getState());
  enterRule(_localctx, 0, TuringParser::RuleProgram);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(101); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(100);
      topLevel();
      setState(103); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::STRING_LITERAL)
      | (1ULL << TuringParser::INTEGER_LITERAL)
      | (1ULL << TuringParser::REAL_LITERAL)
      | (1ULL << TuringParser::TYPE)
      | (1ULL << TuringParser::VAR)
      | (1ULL << TuringParser::PROCEDURE)
      | (1ULL << TuringParser::FUNCTION)
      | (1ULL << TuringParser::CLASS)
      | (1ULL << TuringParser::PROCESS)
      | (1ULL << TuringParser::EXIT)
      | (1ULL << TuringParser::BEGIN)
      | (1ULL << TuringParser::RETURN)
      | (1ULL << TuringParser::RESULT)
      | (1ULL << TuringParser::NEW)
      | (1ULL << TuringParser::FREE)
      | (1ULL << TuringParser::FORK)
      | (1ULL << TuringParser::PUT)
      | (1ULL << TuringParser::ARRAY)
      | (1ULL << TuringParser::DEFERRED)
      | (1ULL << TuringParser::FORWARD)
      | (1ULL << TuringParser::BODY)
      | (1ULL << TuringParser::NOT)
      | (1ULL << TuringParser::CARET))) != 0) || ((((_la - 64) & ~ 0x3fULL) == 0) &&
      ((1ULL << (_la - 64)) & ((1ULL << (TuringParser::L_PAREN - 64))
      | (1ULL << (TuringParser::CHEAT - 64))
      | (1ULL << (TuringParser::PLUS - 64))
      | (1ULL << (TuringParser::MINUS - 64))
      | (1ULL << (TuringParser::IDENTIFIER - 64)))) != 0));
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TopLevelContext ------------------------------------------------------------------

TuringParser::TopLevelContext::TopLevelContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::StatementOrDeclarationContext* TuringParser::TopLevelContext::statementOrDeclaration() {
  return getRuleContext<TuringParser::StatementOrDeclarationContext>(0);
}

TuringParser::ClassDeclarationContext* TuringParser::TopLevelContext::classDeclaration() {
  return getRuleContext<TuringParser::ClassDeclarationContext>(0);
}


size_t TuringParser::TopLevelContext::getRuleIndex() const {
  return TuringParser::RuleTopLevel;
}

void TuringParser::TopLevelContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTopLevel(this);
}

void TuringParser::TopLevelContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTopLevel(this);
}


antlrcpp::Any TuringParser::TopLevelContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitTopLevel(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::TopLevelContext* TuringParser::topLevel() {
  TopLevelContext *_localctx = _tracker.createInstance<TopLevelContext>(_ctx, getState());
  enterRule(_localctx, 2, TuringParser::RuleTopLevel);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(107);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 1, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(105);
      statementOrDeclaration();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(106);
      classDeclaration();
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ProcHeaderContext ------------------------------------------------------------------

TuringParser::ProcHeaderContext::ProcHeaderContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::ProcHeaderContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}

tree::TerminalNode* TuringParser::ProcHeaderContext::PROCEDURE() {
  return getToken(TuringParser::PROCEDURE, 0);
}

tree::TerminalNode* TuringParser::ProcHeaderContext::PROCESS() {
  return getToken(TuringParser::PROCESS, 0);
}

tree::TerminalNode* TuringParser::ProcHeaderContext::L_PAREN() {
  return getToken(TuringParser::L_PAREN, 0);
}

tree::TerminalNode* TuringParser::ProcHeaderContext::R_PAREN() {
  return getToken(TuringParser::R_PAREN, 0);
}

TuringParser::ParamDeclarationListContext* TuringParser::ProcHeaderContext::paramDeclarationList() {
  return getRuleContext<TuringParser::ParamDeclarationListContext>(0);
}


size_t TuringParser::ProcHeaderContext::getRuleIndex() const {
  return TuringParser::RuleProcHeader;
}

void TuringParser::ProcHeaderContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterProcHeader(this);
}

void TuringParser::ProcHeaderContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitProcHeader(this);
}


antlrcpp::Any TuringParser::ProcHeaderContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitProcHeader(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ProcHeaderContext* TuringParser::procHeader() {
  ProcHeaderContext *_localctx = _tracker.createInstance<ProcHeaderContext>(_ctx, getState());
  enterRule(_localctx, 4, TuringParser::RuleProcHeader);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(109);
    _la = _input->LA(1);
    if (!(_la == TuringParser::PROCEDURE

    || _la == TuringParser::PROCESS)) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
    setState(110);
    match(TuringParser::IDENTIFIER);
    setState(116);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 3, _ctx)) {
    case 1: {
      setState(111);
      match(TuringParser::L_PAREN);
      setState(113);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if ((((_la & ~ 0x3fULL) == 0) &&
        ((1ULL << _la) & ((1ULL << TuringParser::VAR)
        | (1ULL << TuringParser::PROCEDURE)
        | (1ULL << TuringParser::FUNCTION)
        | (1ULL << TuringParser::PROCESS)
        | (1ULL << TuringParser::DEFERRED)
        | (1ULL << TuringParser::FORWARD)
        | (1ULL << TuringParser::BODY))) != 0) || _la == TuringParser::IDENTIFIER) {
        setState(112);
        paramDeclarationList(0);
      }
      setState(115);
      match(TuringParser::R_PAREN);
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- FuncHeaderContext ------------------------------------------------------------------

TuringParser::FuncHeaderContext::FuncHeaderContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::FuncHeaderContext::FUNCTION() {
  return getToken(TuringParser::FUNCTION, 0);
}

std::vector<tree::TerminalNode *> TuringParser::FuncHeaderContext::IDENTIFIER() {
  return getTokens(TuringParser::IDENTIFIER);
}

tree::TerminalNode* TuringParser::FuncHeaderContext::IDENTIFIER(size_t i) {
  return getToken(TuringParser::IDENTIFIER, i);
}

tree::TerminalNode* TuringParser::FuncHeaderContext::COLON() {
  return getToken(TuringParser::COLON, 0);
}

TuringParser::TypeSpecContext* TuringParser::FuncHeaderContext::typeSpec() {
  return getRuleContext<TuringParser::TypeSpecContext>(0);
}

tree::TerminalNode* TuringParser::FuncHeaderContext::L_PAREN() {
  return getToken(TuringParser::L_PAREN, 0);
}

tree::TerminalNode* TuringParser::FuncHeaderContext::R_PAREN() {
  return getToken(TuringParser::R_PAREN, 0);
}

TuringParser::ParamDeclarationListContext* TuringParser::FuncHeaderContext::paramDeclarationList() {
  return getRuleContext<TuringParser::ParamDeclarationListContext>(0);
}


size_t TuringParser::FuncHeaderContext::getRuleIndex() const {
  return TuringParser::RuleFuncHeader;
}

void TuringParser::FuncHeaderContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterFuncHeader(this);
}

void TuringParser::FuncHeaderContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitFuncHeader(this);
}


antlrcpp::Any TuringParser::FuncHeaderContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitFuncHeader(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::FuncHeaderContext* TuringParser::funcHeader() {
  FuncHeaderContext *_localctx = _tracker.createInstance<FuncHeaderContext>(_ctx, getState());
  enterRule(_localctx, 6, TuringParser::RuleFuncHeader);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(118);
    match(TuringParser::FUNCTION);
    setState(119);
    match(TuringParser::IDENTIFIER);
    setState(125);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == TuringParser::L_PAREN) {
      setState(120);
      match(TuringParser::L_PAREN);
      setState(122);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if ((((_la & ~ 0x3fULL) == 0) &&
        ((1ULL << _la) & ((1ULL << TuringParser::VAR)
        | (1ULL << TuringParser::PROCEDURE)
        | (1ULL << TuringParser::FUNCTION)
        | (1ULL << TuringParser::PROCESS)
        | (1ULL << TuringParser::DEFERRED)
        | (1ULL << TuringParser::FORWARD)
        | (1ULL << TuringParser::BODY))) != 0) || _la == TuringParser::IDENTIFIER) {
        setState(121);
        paramDeclarationList(0);
      }
      setState(124);
      match(TuringParser::R_PAREN);
    }
    setState(128);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == TuringParser::IDENTIFIER) {
      setState(127);
      match(TuringParser::IDENTIFIER);
    }
    setState(130);
    match(TuringParser::COLON);
    setState(131);
    typeSpec();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- SubprogramHeaderContext ------------------------------------------------------------------

TuringParser::SubprogramHeaderContext::SubprogramHeaderContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::FuncHeaderContext* TuringParser::SubprogramHeaderContext::funcHeader() {
  return getRuleContext<TuringParser::FuncHeaderContext>(0);
}

TuringParser::ProcHeaderContext* TuringParser::SubprogramHeaderContext::procHeader() {
  return getRuleContext<TuringParser::ProcHeaderContext>(0);
}

TuringParser::SubprogramPrefixContext* TuringParser::SubprogramHeaderContext::subprogramPrefix() {
  return getRuleContext<TuringParser::SubprogramPrefixContext>(0);
}


size_t TuringParser::SubprogramHeaderContext::getRuleIndex() const {
  return TuringParser::RuleSubprogramHeader;
}

void TuringParser::SubprogramHeaderContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterSubprogramHeader(this);
}

void TuringParser::SubprogramHeaderContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitSubprogramHeader(this);
}


antlrcpp::Any TuringParser::SubprogramHeaderContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitSubprogramHeader(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::SubprogramHeaderContext* TuringParser::subprogramHeader() {
  SubprogramHeaderContext *_localctx = _tracker.createInstance<SubprogramHeaderContext>(_ctx, getState());
  enterRule(_localctx, 8, TuringParser::RuleSubprogramHeader);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(134);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::DEFERRED)
      | (1ULL << TuringParser::FORWARD)
      | (1ULL << TuringParser::BODY))) != 0)) {
      setState(133);
      subprogramPrefix();
    }
    setState(138);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::FUNCTION: {
        setState(136);
        funcHeader();
        break;
      }

      case TuringParser::PROCEDURE:
      case TuringParser::PROCESS: {
        setState(137);
        procHeader();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- SubprogramDeclarationContext ------------------------------------------------------------------

TuringParser::SubprogramDeclarationContext::SubprogramDeclarationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::SubprogramHeaderContext* TuringParser::SubprogramDeclarationContext::subprogramHeader() {
  return getRuleContext<TuringParser::SubprogramHeaderContext>(0);
}

TuringParser::ProcBodyContext* TuringParser::SubprogramDeclarationContext::procBody() {
  return getRuleContext<TuringParser::ProcBodyContext>(0);
}

tree::TerminalNode* TuringParser::SubprogramDeclarationContext::END() {
  return getToken(TuringParser::END, 0);
}

tree::TerminalNode* TuringParser::SubprogramDeclarationContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}


size_t TuringParser::SubprogramDeclarationContext::getRuleIndex() const {
  return TuringParser::RuleSubprogramDeclaration;
}

void TuringParser::SubprogramDeclarationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterSubprogramDeclaration(this);
}

void TuringParser::SubprogramDeclarationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitSubprogramDeclaration(this);
}


antlrcpp::Any TuringParser::SubprogramDeclarationContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitSubprogramDeclaration(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::SubprogramDeclarationContext* TuringParser::subprogramDeclaration() {
  SubprogramDeclarationContext *_localctx = _tracker.createInstance<SubprogramDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 10, TuringParser::RuleSubprogramDeclaration);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(140);
    subprogramHeader();
    setState(141);
    procBody();
    setState(142);
    match(TuringParser::END);
    setState(143);
    match(TuringParser::IDENTIFIER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- SubprogramPrefixContext ------------------------------------------------------------------

TuringParser::SubprogramPrefixContext::SubprogramPrefixContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::SubprogramPrefixContext::DEFERRED() {
  return getToken(TuringParser::DEFERRED, 0);
}

tree::TerminalNode* TuringParser::SubprogramPrefixContext::BODY() {
  return getToken(TuringParser::BODY, 0);
}

tree::TerminalNode* TuringParser::SubprogramPrefixContext::FORWARD() {
  return getToken(TuringParser::FORWARD, 0);
}


size_t TuringParser::SubprogramPrefixContext::getRuleIndex() const {
  return TuringParser::RuleSubprogramPrefix;
}

void TuringParser::SubprogramPrefixContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterSubprogramPrefix(this);
}

void TuringParser::SubprogramPrefixContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitSubprogramPrefix(this);
}


antlrcpp::Any TuringParser::SubprogramPrefixContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitSubprogramPrefix(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::SubprogramPrefixContext* TuringParser::subprogramPrefix() {
  SubprogramPrefixContext *_localctx = _tracker.createInstance<SubprogramPrefixContext>(_ctx, getState());
  enterRule(_localctx, 12, TuringParser::RuleSubprogramPrefix);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(145);
    _la = _input->LA(1);
    if (!((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::DEFERRED)
      | (1ULL << TuringParser::FORWARD)
      | (1ULL << TuringParser::BODY))) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ClassDeclarationContext ------------------------------------------------------------------

TuringParser::ClassDeclarationContext::ClassDeclarationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::ClassDeclarationContext::CLASS() {
  return getToken(TuringParser::CLASS, 0);
}

std::vector<tree::TerminalNode *> TuringParser::ClassDeclarationContext::IDENTIFIER() {
  return getTokens(TuringParser::IDENTIFIER);
}

tree::TerminalNode* TuringParser::ClassDeclarationContext::IDENTIFIER(size_t i) {
  return getToken(TuringParser::IDENTIFIER, i);
}

tree::TerminalNode* TuringParser::ClassDeclarationContext::END() {
  return getToken(TuringParser::END, 0);
}

std::vector<TuringParser::ClassBodyContext *> TuringParser::ClassDeclarationContext::classBody() {
  return getRuleContexts<TuringParser::ClassBodyContext>();
}

TuringParser::ClassBodyContext* TuringParser::ClassDeclarationContext::classBody(size_t i) {
  return getRuleContext<TuringParser::ClassBodyContext>(i);
}


size_t TuringParser::ClassDeclarationContext::getRuleIndex() const {
  return TuringParser::RuleClassDeclaration;
}

void TuringParser::ClassDeclarationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterClassDeclaration(this);
}

void TuringParser::ClassDeclarationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitClassDeclaration(this);
}


antlrcpp::Any TuringParser::ClassDeclarationContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitClassDeclaration(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ClassDeclarationContext* TuringParser::classDeclaration() {
  ClassDeclarationContext *_localctx = _tracker.createInstance<ClassDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 14, TuringParser::RuleClassDeclaration);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(147);
    match(TuringParser::CLASS);
    setState(148);
    match(TuringParser::IDENTIFIER);
    setState(150); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(149);
      classBody();
      setState(152); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::STRING_LITERAL)
      | (1ULL << TuringParser::INTEGER_LITERAL)
      | (1ULL << TuringParser::REAL_LITERAL)
      | (1ULL << TuringParser::TYPE)
      | (1ULL << TuringParser::VAR)
      | (1ULL << TuringParser::PROCEDURE)
      | (1ULL << TuringParser::FUNCTION)
      | (1ULL << TuringParser::CLASS)
      | (1ULL << TuringParser::PROCESS)
      | (1ULL << TuringParser::EXIT)
      | (1ULL << TuringParser::BEGIN)
      | (1ULL << TuringParser::RETURN)
      | (1ULL << TuringParser::RESULT)
      | (1ULL << TuringParser::NEW)
      | (1ULL << TuringParser::FREE)
      | (1ULL << TuringParser::FORK)
      | (1ULL << TuringParser::EXPORT)
      | (1ULL << TuringParser::INHERIT)
      | (1ULL << TuringParser::IMPLEMENT)
      | (1ULL << TuringParser::PUT)
      | (1ULL << TuringParser::ARRAY)
      | (1ULL << TuringParser::DEFERRED)
      | (1ULL << TuringParser::FORWARD)
      | (1ULL << TuringParser::BODY)
      | (1ULL << TuringParser::NOT)
      | (1ULL << TuringParser::CARET))) != 0) || ((((_la - 64) & ~ 0x3fULL) == 0) &&
      ((1ULL << (_la - 64)) & ((1ULL << (TuringParser::L_PAREN - 64))
      | (1ULL << (TuringParser::CHEAT - 64))
      | (1ULL << (TuringParser::PLUS - 64))
      | (1ULL << (TuringParser::MINUS - 64))
      | (1ULL << (TuringParser::IDENTIFIER - 64)))) != 0));
    setState(154);
    match(TuringParser::END);
    setState(155);
    match(TuringParser::IDENTIFIER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ClassBodyContext ------------------------------------------------------------------

TuringParser::ClassBodyContext::ClassBodyContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::ExportStatementContext* TuringParser::ClassBodyContext::exportStatement() {
  return getRuleContext<TuringParser::ExportStatementContext>(0);
}

TuringParser::InheritStatementContext* TuringParser::ClassBodyContext::inheritStatement() {
  return getRuleContext<TuringParser::InheritStatementContext>(0);
}

TuringParser::ImplementStatementContext* TuringParser::ClassBodyContext::implementStatement() {
  return getRuleContext<TuringParser::ImplementStatementContext>(0);
}

TuringParser::StatementOrDeclarationContext* TuringParser::ClassBodyContext::statementOrDeclaration() {
  return getRuleContext<TuringParser::StatementOrDeclarationContext>(0);
}


size_t TuringParser::ClassBodyContext::getRuleIndex() const {
  return TuringParser::RuleClassBody;
}

void TuringParser::ClassBodyContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterClassBody(this);
}

void TuringParser::ClassBodyContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitClassBody(this);
}


antlrcpp::Any TuringParser::ClassBodyContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitClassBody(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ClassBodyContext* TuringParser::classBody() {
  ClassBodyContext *_localctx = _tracker.createInstance<ClassBodyContext>(_ctx, getState());
  enterRule(_localctx, 16, TuringParser::RuleClassBody);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(161);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::EXPORT: {
        enterOuterAlt(_localctx, 1);
        setState(157);
        exportStatement();
        break;
      }

      case TuringParser::INHERIT: {
        enterOuterAlt(_localctx, 2);
        setState(158);
        inheritStatement();
        break;
      }

      case TuringParser::IMPLEMENT: {
        enterOuterAlt(_localctx, 3);
        setState(159);
        implementStatement();
        break;
      }

      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL:
      case TuringParser::REAL_LITERAL:
      case TuringParser::TYPE:
      case TuringParser::VAR:
      case TuringParser::PROCEDURE:
      case TuringParser::FUNCTION:
      case TuringParser::CLASS:
      case TuringParser::PROCESS:
      case TuringParser::EXIT:
      case TuringParser::BEGIN:
      case TuringParser::RETURN:
      case TuringParser::RESULT:
      case TuringParser::NEW:
      case TuringParser::FREE:
      case TuringParser::FORK:
      case TuringParser::PUT:
      case TuringParser::ARRAY:
      case TuringParser::DEFERRED:
      case TuringParser::FORWARD:
      case TuringParser::BODY:
      case TuringParser::NOT:
      case TuringParser::CARET:
      case TuringParser::L_PAREN:
      case TuringParser::CHEAT:
      case TuringParser::PLUS:
      case TuringParser::MINUS:
      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 4);
        setState(160);
        statementOrDeclaration();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExportStatementContext ------------------------------------------------------------------

TuringParser::ExportStatementContext::ExportStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::ExportStatementContext::EXPORT() {
  return getToken(TuringParser::EXPORT, 0);
}

TuringParser::ExportListContext* TuringParser::ExportStatementContext::exportList() {
  return getRuleContext<TuringParser::ExportListContext>(0);
}

tree::TerminalNode* TuringParser::ExportStatementContext::L_PAREN() {
  return getToken(TuringParser::L_PAREN, 0);
}

tree::TerminalNode* TuringParser::ExportStatementContext::R_PAREN() {
  return getToken(TuringParser::R_PAREN, 0);
}


size_t TuringParser::ExportStatementContext::getRuleIndex() const {
  return TuringParser::RuleExportStatement;
}

void TuringParser::ExportStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExportStatement(this);
}

void TuringParser::ExportStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExportStatement(this);
}


antlrcpp::Any TuringParser::ExportStatementContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitExportStatement(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ExportStatementContext* TuringParser::exportStatement() {
  ExportStatementContext *_localctx = _tracker.createInstance<ExportStatementContext>(_ctx, getState());
  enterRule(_localctx, 18, TuringParser::RuleExportStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(170);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 11, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(163);
      match(TuringParser::EXPORT);
      setState(164);
      exportList(0);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(165);
      match(TuringParser::EXPORT);
      setState(166);
      match(TuringParser::L_PAREN);
      setState(167);
      exportList(0);
      setState(168);
      match(TuringParser::R_PAREN);
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExportListItemContext ------------------------------------------------------------------

TuringParser::ExportListItemContext::ExportListItemContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::ExportListItemContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}

TuringParser::HowExportContext* TuringParser::ExportListItemContext::howExport() {
  return getRuleContext<TuringParser::HowExportContext>(0);
}


size_t TuringParser::ExportListItemContext::getRuleIndex() const {
  return TuringParser::RuleExportListItem;
}

void TuringParser::ExportListItemContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExportListItem(this);
}

void TuringParser::ExportListItemContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExportListItem(this);
}


antlrcpp::Any TuringParser::ExportListItemContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitExportListItem(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ExportListItemContext* TuringParser::exportListItem() {
  ExportListItemContext *_localctx = _tracker.createInstance<ExportListItemContext>(_ctx, getState());
  enterRule(_localctx, 20, TuringParser::RuleExportListItem);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(173);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::T__0)
      | (1ULL << TuringParser::T__1)
      | (1ULL << TuringParser::T__2)
      | (1ULL << TuringParser::VAR))) != 0)) {
      setState(172);
      howExport();
    }
    setState(175);
    match(TuringParser::IDENTIFIER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExportListContext ------------------------------------------------------------------

TuringParser::ExportListContext::ExportListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::ExportListItemContext* TuringParser::ExportListContext::exportListItem() {
  return getRuleContext<TuringParser::ExportListItemContext>(0);
}

TuringParser::ExportListContext* TuringParser::ExportListContext::exportList() {
  return getRuleContext<TuringParser::ExportListContext>(0);
}

tree::TerminalNode* TuringParser::ExportListContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
}


size_t TuringParser::ExportListContext::getRuleIndex() const {
  return TuringParser::RuleExportList;
}

void TuringParser::ExportListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExportList(this);
}

void TuringParser::ExportListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExportList(this);
}


antlrcpp::Any TuringParser::ExportListContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitExportList(this);
  else
    return visitor->visitChildren(this);
}


TuringParser::ExportListContext* TuringParser::exportList() {
   return exportList(0);
}

TuringParser::ExportListContext* TuringParser::exportList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::ExportListContext *_localctx = _tracker.createInstance<ExportListContext>(_ctx, parentState);
  TuringParser::ExportListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 22;
  enterRecursionRule(_localctx, 22, TuringParser::RuleExportList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(178);
    exportListItem();
    _ctx->stop = _input->LT(-1);
    setState(185);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 13, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ExportListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleExportList);
        setState(180);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(181);
        match(TuringParser::COMMA);
        setState(182);
        exportListItem(); 
      }
      setState(187);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 13, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- InheritStatementContext ------------------------------------------------------------------

TuringParser::InheritStatementContext::InheritStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::InheritStatementContext::INHERIT() {
  return getToken(TuringParser::INHERIT, 0);
}

TuringParser::IdOrFileItemContext* TuringParser::InheritStatementContext::idOrFileItem() {
  return getRuleContext<TuringParser::IdOrFileItemContext>(0);
}

tree::TerminalNode* TuringParser::InheritStatementContext::L_PAREN() {
  return getToken(TuringParser::L_PAREN, 0);
}

tree::TerminalNode* TuringParser::InheritStatementContext::R_PAREN() {
  return getToken(TuringParser::R_PAREN, 0);
}


size_t TuringParser::InheritStatementContext::getRuleIndex() const {
  return TuringParser::RuleInheritStatement;
}

void TuringParser::InheritStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterInheritStatement(this);
}

void TuringParser::InheritStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitInheritStatement(this);
}


antlrcpp::Any TuringParser::InheritStatementContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitInheritStatement(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::InheritStatementContext* TuringParser::inheritStatement() {
  InheritStatementContext *_localctx = _tracker.createInstance<InheritStatementContext>(_ctx, getState());
  enterRule(_localctx, 24, TuringParser::RuleInheritStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(195);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 14, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(188);
      match(TuringParser::INHERIT);
      setState(189);
      idOrFileItem();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(190);
      match(TuringParser::INHERIT);
      setState(191);
      match(TuringParser::L_PAREN);
      setState(192);
      idOrFileItem();
      setState(193);
      match(TuringParser::R_PAREN);
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ImplementStatementContext ------------------------------------------------------------------

TuringParser::ImplementStatementContext::ImplementStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::ImplementStatementContext::IMPLEMENT() {
  return getToken(TuringParser::IMPLEMENT, 0);
}

TuringParser::IdOrFileItemContext* TuringParser::ImplementStatementContext::idOrFileItem() {
  return getRuleContext<TuringParser::IdOrFileItemContext>(0);
}

tree::TerminalNode* TuringParser::ImplementStatementContext::BY() {
  return getToken(TuringParser::BY, 0);
}

tree::TerminalNode* TuringParser::ImplementStatementContext::L_PAREN() {
  return getToken(TuringParser::L_PAREN, 0);
}

tree::TerminalNode* TuringParser::ImplementStatementContext::R_PAREN() {
  return getToken(TuringParser::R_PAREN, 0);
}


size_t TuringParser::ImplementStatementContext::getRuleIndex() const {
  return TuringParser::RuleImplementStatement;
}

void TuringParser::ImplementStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterImplementStatement(this);
}

void TuringParser::ImplementStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitImplementStatement(this);
}


antlrcpp::Any TuringParser::ImplementStatementContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitImplementStatement(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ImplementStatementContext* TuringParser::implementStatement() {
  ImplementStatementContext *_localctx = _tracker.createInstance<ImplementStatementContext>(_ctx, getState());
  enterRule(_localctx, 26, TuringParser::RuleImplementStatement);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(210);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 17, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(197);
      match(TuringParser::IMPLEMENT);
      setState(199);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == TuringParser::BY) {
        setState(198);
        match(TuringParser::BY);
      }
      setState(201);
      idOrFileItem();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(202);
      match(TuringParser::IMPLEMENT);
      setState(204);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == TuringParser::BY) {
        setState(203);
        match(TuringParser::BY);
      }
      setState(206);
      match(TuringParser::L_PAREN);
      setState(207);
      idOrFileItem();
      setState(208);
      match(TuringParser::R_PAREN);
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- IdOrFileItemContext ------------------------------------------------------------------

TuringParser::IdOrFileItemContext::IdOrFileItemContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::IdOrFileItemContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}

tree::TerminalNode* TuringParser::IdOrFileItemContext::IN() {
  return getToken(TuringParser::IN, 0);
}

tree::TerminalNode* TuringParser::IdOrFileItemContext::STRING_LITERAL() {
  return getToken(TuringParser::STRING_LITERAL, 0);
}


size_t TuringParser::IdOrFileItemContext::getRuleIndex() const {
  return TuringParser::RuleIdOrFileItem;
}

void TuringParser::IdOrFileItemContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterIdOrFileItem(this);
}

void TuringParser::IdOrFileItemContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitIdOrFileItem(this);
}


antlrcpp::Any TuringParser::IdOrFileItemContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitIdOrFileItem(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::IdOrFileItemContext* TuringParser::idOrFileItem() {
  IdOrFileItemContext *_localctx = _tracker.createInstance<IdOrFileItemContext>(_ctx, getState());
  enterRule(_localctx, 28, TuringParser::RuleIdOrFileItem);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(216);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 18, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(212);
      match(TuringParser::IDENTIFIER);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(213);
      match(TuringParser::IDENTIFIER);
      setState(214);
      match(TuringParser::IN);
      setState(215);
      match(TuringParser::STRING_LITERAL);
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- HowExportContext ------------------------------------------------------------------

TuringParser::HowExportContext::HowExportContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::HowExportContext::VAR() {
  return getToken(TuringParser::VAR, 0);
}


size_t TuringParser::HowExportContext::getRuleIndex() const {
  return TuringParser::RuleHowExport;
}

void TuringParser::HowExportContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterHowExport(this);
}

void TuringParser::HowExportContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitHowExport(this);
}


antlrcpp::Any TuringParser::HowExportContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitHowExport(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::HowExportContext* TuringParser::howExport() {
  HowExportContext *_localctx = _tracker.createInstance<HowExportContext>(_ctx, getState());
  enterRule(_localctx, 30, TuringParser::RuleHowExport);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(218);
    _la = _input->LA(1);
    if (!((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::T__0)
      | (1ULL << TuringParser::T__1)
      | (1ULL << TuringParser::T__2)
      | (1ULL << TuringParser::VAR))) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ParamDeclarationContext ------------------------------------------------------------------

TuringParser::ParamDeclarationContext::ParamDeclarationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::IdentifierListContext* TuringParser::ParamDeclarationContext::identifierList() {
  return getRuleContext<TuringParser::IdentifierListContext>(0);
}

tree::TerminalNode* TuringParser::ParamDeclarationContext::COLON() {
  return getToken(TuringParser::COLON, 0);
}

TuringParser::TypeSpecContext* TuringParser::ParamDeclarationContext::typeSpec() {
  return getRuleContext<TuringParser::TypeSpecContext>(0);
}

tree::TerminalNode* TuringParser::ParamDeclarationContext::VAR() {
  return getToken(TuringParser::VAR, 0);
}

TuringParser::SubprogramHeaderContext* TuringParser::ParamDeclarationContext::subprogramHeader() {
  return getRuleContext<TuringParser::SubprogramHeaderContext>(0);
}


size_t TuringParser::ParamDeclarationContext::getRuleIndex() const {
  return TuringParser::RuleParamDeclaration;
}

void TuringParser::ParamDeclarationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterParamDeclaration(this);
}

void TuringParser::ParamDeclarationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitParamDeclaration(this);
}


antlrcpp::Any TuringParser::ParamDeclarationContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitParamDeclaration(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ParamDeclarationContext* TuringParser::paramDeclaration() {
  ParamDeclarationContext *_localctx = _tracker.createInstance<ParamDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 32, TuringParser::RuleParamDeclaration);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(228);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::VAR:
      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 1);
        setState(221);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == TuringParser::VAR) {
          setState(220);
          match(TuringParser::VAR);
        }
        setState(223);
        identifierList(0);
        setState(224);
        match(TuringParser::COLON);
        setState(225);
        typeSpec();
        break;
      }

      case TuringParser::PROCEDURE:
      case TuringParser::FUNCTION:
      case TuringParser::PROCESS:
      case TuringParser::DEFERRED:
      case TuringParser::FORWARD:
      case TuringParser::BODY: {
        enterOuterAlt(_localctx, 2);
        setState(227);
        subprogramHeader();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ParamDeclarationListContext ------------------------------------------------------------------

TuringParser::ParamDeclarationListContext::ParamDeclarationListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::ParamDeclarationContext* TuringParser::ParamDeclarationListContext::paramDeclaration() {
  return getRuleContext<TuringParser::ParamDeclarationContext>(0);
}

TuringParser::ParamDeclarationListContext* TuringParser::ParamDeclarationListContext::paramDeclarationList() {
  return getRuleContext<TuringParser::ParamDeclarationListContext>(0);
}

tree::TerminalNode* TuringParser::ParamDeclarationListContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
}


size_t TuringParser::ParamDeclarationListContext::getRuleIndex() const {
  return TuringParser::RuleParamDeclarationList;
}

void TuringParser::ParamDeclarationListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterParamDeclarationList(this);
}

void TuringParser::ParamDeclarationListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitParamDeclarationList(this);
}


antlrcpp::Any TuringParser::ParamDeclarationListContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitParamDeclarationList(this);
  else
    return visitor->visitChildren(this);
}


TuringParser::ParamDeclarationListContext* TuringParser::paramDeclarationList() {
   return paramDeclarationList(0);
}

TuringParser::ParamDeclarationListContext* TuringParser::paramDeclarationList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::ParamDeclarationListContext *_localctx = _tracker.createInstance<ParamDeclarationListContext>(_ctx, parentState);
  TuringParser::ParamDeclarationListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 34;
  enterRecursionRule(_localctx, 34, TuringParser::RuleParamDeclarationList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(231);
    paramDeclaration();
    _ctx->stop = _input->LT(-1);
    setState(238);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 21, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ParamDeclarationListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleParamDeclarationList);
        setState(233);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(234);
        match(TuringParser::COMMA);
        setState(235);
        paramDeclaration(); 
      }
      setState(240);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 21, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- ProcBodyContext ------------------------------------------------------------------

TuringParser::ProcBodyContext::ProcBodyContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<TuringParser::StatementOrDeclarationContext *> TuringParser::ProcBodyContext::statementOrDeclaration() {
  return getRuleContexts<TuringParser::StatementOrDeclarationContext>();
}

TuringParser::StatementOrDeclarationContext* TuringParser::ProcBodyContext::statementOrDeclaration(size_t i) {
  return getRuleContext<TuringParser::StatementOrDeclarationContext>(i);
}


size_t TuringParser::ProcBodyContext::getRuleIndex() const {
  return TuringParser::RuleProcBody;
}

void TuringParser::ProcBodyContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterProcBody(this);
}

void TuringParser::ProcBodyContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitProcBody(this);
}


antlrcpp::Any TuringParser::ProcBodyContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitProcBody(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ProcBodyContext* TuringParser::procBody() {
  ProcBodyContext *_localctx = _tracker.createInstance<ProcBodyContext>(_ctx, getState());
  enterRule(_localctx, 36, TuringParser::RuleProcBody);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(242); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(241);
      statementOrDeclaration();
      setState(244); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::STRING_LITERAL)
      | (1ULL << TuringParser::INTEGER_LITERAL)
      | (1ULL << TuringParser::REAL_LITERAL)
      | (1ULL << TuringParser::TYPE)
      | (1ULL << TuringParser::VAR)
      | (1ULL << TuringParser::PROCEDURE)
      | (1ULL << TuringParser::FUNCTION)
      | (1ULL << TuringParser::CLASS)
      | (1ULL << TuringParser::PROCESS)
      | (1ULL << TuringParser::EXIT)
      | (1ULL << TuringParser::BEGIN)
      | (1ULL << TuringParser::RETURN)
      | (1ULL << TuringParser::RESULT)
      | (1ULL << TuringParser::NEW)
      | (1ULL << TuringParser::FREE)
      | (1ULL << TuringParser::FORK)
      | (1ULL << TuringParser::PUT)
      | (1ULL << TuringParser::ARRAY)
      | (1ULL << TuringParser::DEFERRED)
      | (1ULL << TuringParser::FORWARD)
      | (1ULL << TuringParser::BODY)
      | (1ULL << TuringParser::NOT)
      | (1ULL << TuringParser::CARET))) != 0) || ((((_la - 64) & ~ 0x3fULL) == 0) &&
      ((1ULL << (_la - 64)) & ((1ULL << (TuringParser::L_PAREN - 64))
      | (1ULL << (TuringParser::CHEAT - 64))
      | (1ULL << (TuringParser::PLUS - 64))
      | (1ULL << (TuringParser::MINUS - 64))
      | (1ULL << (TuringParser::IDENTIFIER - 64)))) != 0));
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- StatementOrDeclarationContext ------------------------------------------------------------------

TuringParser::StatementOrDeclarationContext::StatementOrDeclarationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::StatementsContext* TuringParser::StatementOrDeclarationContext::statements() {
  return getRuleContext<TuringParser::StatementsContext>(0);
}

TuringParser::DeclarationContext* TuringParser::StatementOrDeclarationContext::declaration() {
  return getRuleContext<TuringParser::DeclarationContext>(0);
}


size_t TuringParser::StatementOrDeclarationContext::getRuleIndex() const {
  return TuringParser::RuleStatementOrDeclaration;
}

void TuringParser::StatementOrDeclarationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterStatementOrDeclaration(this);
}

void TuringParser::StatementOrDeclarationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitStatementOrDeclaration(this);
}


antlrcpp::Any TuringParser::StatementOrDeclarationContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitStatementOrDeclaration(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::StatementOrDeclarationContext* TuringParser::statementOrDeclaration() {
  StatementOrDeclarationContext *_localctx = _tracker.createInstance<StatementOrDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 38, TuringParser::RuleStatementOrDeclaration);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(248);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL:
      case TuringParser::REAL_LITERAL:
      case TuringParser::EXIT:
      case TuringParser::BEGIN:
      case TuringParser::RETURN:
      case TuringParser::RESULT:
      case TuringParser::NEW:
      case TuringParser::FREE:
      case TuringParser::FORK:
      case TuringParser::PUT:
      case TuringParser::NOT:
      case TuringParser::CARET:
      case TuringParser::L_PAREN:
      case TuringParser::CHEAT:
      case TuringParser::PLUS:
      case TuringParser::MINUS:
      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 1);
        setState(246);
        statements();
        break;
      }

      case TuringParser::TYPE:
      case TuringParser::VAR:
      case TuringParser::PROCEDURE:
      case TuringParser::FUNCTION:
      case TuringParser::CLASS:
      case TuringParser::PROCESS:
      case TuringParser::ARRAY:
      case TuringParser::DEFERRED:
      case TuringParser::FORWARD:
      case TuringParser::BODY: {
        enterOuterAlt(_localctx, 2);
        setState(247);
        declaration();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PrimaryExpressionContext ------------------------------------------------------------------

TuringParser::PrimaryExpressionContext::PrimaryExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::LiteralContext* TuringParser::PrimaryExpressionContext::literal() {
  return getRuleContext<TuringParser::LiteralContext>(0);
}

tree::TerminalNode* TuringParser::PrimaryExpressionContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}

tree::TerminalNode* TuringParser::PrimaryExpressionContext::L_PAREN() {
  return getToken(TuringParser::L_PAREN, 0);
}

TuringParser::ExpressionContext* TuringParser::PrimaryExpressionContext::expression() {
  return getRuleContext<TuringParser::ExpressionContext>(0);
}

tree::TerminalNode* TuringParser::PrimaryExpressionContext::R_PAREN() {
  return getToken(TuringParser::R_PAREN, 0);
}


size_t TuringParser::PrimaryExpressionContext::getRuleIndex() const {
  return TuringParser::RulePrimaryExpression;
}

void TuringParser::PrimaryExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterPrimaryExpression(this);
}

void TuringParser::PrimaryExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitPrimaryExpression(this);
}


antlrcpp::Any TuringParser::PrimaryExpressionContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitPrimaryExpression(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::PrimaryExpressionContext* TuringParser::primaryExpression() {
  PrimaryExpressionContext *_localctx = _tracker.createInstance<PrimaryExpressionContext>(_ctx, getState());
  enterRule(_localctx, 40, TuringParser::RulePrimaryExpression);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(256);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL:
      case TuringParser::REAL_LITERAL: {
        enterOuterAlt(_localctx, 1);
        setState(250);
        literal();
        break;
      }

      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 2);
        setState(251);
        match(TuringParser::IDENTIFIER);
        break;
      }

      case TuringParser::L_PAREN: {
        enterOuterAlt(_localctx, 3);
        setState(252);
        match(TuringParser::L_PAREN);
        setState(253);
        expression(0);
        setState(254);
        match(TuringParser::R_PAREN);
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExpressionContext ------------------------------------------------------------------

TuringParser::ExpressionContext::ExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::PrimaryExpressionContext* TuringParser::ExpressionContext::primaryExpression() {
  return getRuleContext<TuringParser::PrimaryExpressionContext>(0);
}

std::vector<TuringParser::ExpressionContext *> TuringParser::ExpressionContext::expression() {
  return getRuleContexts<TuringParser::ExpressionContext>();
}

TuringParser::ExpressionContext* TuringParser::ExpressionContext::expression(size_t i) {
  return getRuleContext<TuringParser::ExpressionContext>(i);
}

tree::TerminalNode* TuringParser::ExpressionContext::CARET() {
  return getToken(TuringParser::CARET, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::PLUS() {
  return getToken(TuringParser::PLUS, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::MINUS() {
  return getToken(TuringParser::MINUS, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::CHEAT() {
  return getToken(TuringParser::CHEAT, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::NOT() {
  return getToken(TuringParser::NOT, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::EXP() {
  return getToken(TuringParser::EXP, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::MULTIPLY() {
  return getToken(TuringParser::MULTIPLY, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::DIVIDE() {
  return getToken(TuringParser::DIVIDE, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::DIV() {
  return getToken(TuringParser::DIV, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::MOD() {
  return getToken(TuringParser::MOD, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::REM() {
  return getToken(TuringParser::REM, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::SHR() {
  return getToken(TuringParser::SHR, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::SHL() {
  return getToken(TuringParser::SHL, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::XOR() {
  return getToken(TuringParser::XOR, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::LESSTHAN() {
  return getToken(TuringParser::LESSTHAN, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::GREATERTHAN() {
  return getToken(TuringParser::GREATERTHAN, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::EQUAL() {
  return getToken(TuringParser::EQUAL, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::LESSTHANEQUAL() {
  return getToken(TuringParser::LESSTHANEQUAL, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::GREATERTHANEQUAL() {
  return getToken(TuringParser::GREATERTHANEQUAL, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::NOTEQUAL() {
  return getToken(TuringParser::NOTEQUAL, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::IN() {
  return getToken(TuringParser::IN, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::NOT_IN() {
  return getToken(TuringParser::NOT_IN, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::AND() {
  return getToken(TuringParser::AND, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::OR() {
  return getToken(TuringParser::OR, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::IMPLIES() {
  return getToken(TuringParser::IMPLIES, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::ASSIGNMENT() {
  return getToken(TuringParser::ASSIGNMENT, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::PLUSEQUALS() {
  return getToken(TuringParser::PLUSEQUALS, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::MINUSEQUALS() {
  return getToken(TuringParser::MINUSEQUALS, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::MULTIPLYEQUALS() {
  return getToken(TuringParser::MULTIPLYEQUALS, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::DIVIDEEQUALS() {
  return getToken(TuringParser::DIVIDEEQUALS, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::DIVEQUALS() {
  return getToken(TuringParser::DIVEQUALS, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::SHLEQUALS() {
  return getToken(TuringParser::SHLEQUALS, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::SHREQUALS() {
  return getToken(TuringParser::SHREQUALS, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::DEREFERENCE() {
  return getToken(TuringParser::DEREFERENCE, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::DOT() {
  return getToken(TuringParser::DOT, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::L_PAREN() {
  return getToken(TuringParser::L_PAREN, 0);
}

tree::TerminalNode* TuringParser::ExpressionContext::R_PAREN() {
  return getToken(TuringParser::R_PAREN, 0);
}

TuringParser::ExpressionListContext* TuringParser::ExpressionContext::expressionList() {
  return getRuleContext<TuringParser::ExpressionListContext>(0);
}


size_t TuringParser::ExpressionContext::getRuleIndex() const {
  return TuringParser::RuleExpression;
}

void TuringParser::ExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExpression(this);
}

void TuringParser::ExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExpression(this);
}


antlrcpp::Any TuringParser::ExpressionContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitExpression(this);
  else
    return visitor->visitChildren(this);
}


TuringParser::ExpressionContext* TuringParser::expression() {
   return expression(0);
}

TuringParser::ExpressionContext* TuringParser::expression(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::ExpressionContext *_localctx = _tracker.createInstance<ExpressionContext>(_ctx, parentState);
  TuringParser::ExpressionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 42;
  enterRecursionRule(_localctx, 42, TuringParser::RuleExpression, precedence);

    size_t _la = 0;

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(266);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL:
      case TuringParser::REAL_LITERAL:
      case TuringParser::L_PAREN:
      case TuringParser::IDENTIFIER: {
        setState(259);
        primaryExpression();
        break;
      }

      case TuringParser::CARET: {
        setState(260);
        dynamic_cast<ExpressionContext *>(_localctx)->prefix = match(TuringParser::CARET);
        setState(261);
        expression(13);
        break;
      }

      case TuringParser::CHEAT:
      case TuringParser::PLUS:
      case TuringParser::MINUS: {
        setState(262);
        dynamic_cast<ExpressionContext *>(_localctx)->prefix = _input->LT(1);
        _la = _input->LA(1);
        if (!(((((_la - 69) & ~ 0x3fULL) == 0) &&
          ((1ULL << (_la - 69)) & ((1ULL << (TuringParser::CHEAT - 69))
          | (1ULL << (TuringParser::PLUS - 69))
          | (1ULL << (TuringParser::MINUS - 69)))) != 0))) {
          dynamic_cast<ExpressionContext *>(_localctx)->prefix = _errHandler->recoverInline(this);
        }
        else {
          _errHandler->reportMatch(this);
          consume();
        }
        setState(263);
        expression(9);
        break;
      }

      case TuringParser::NOT: {
        setState(264);
        dynamic_cast<ExpressionContext *>(_localctx)->prefix = match(TuringParser::NOT);
        setState(265);
        expression(5);
        break;
      }

    default:
      throw NoViableAltException(this);
    }
    _ctx->stop = _input->LT(-1);
    setState(303);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 28, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        setState(301);
        _errHandler->sync(this);
        switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 27, _ctx)) {
        case 1: {
          _localctx = _tracker.createInstance<ExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpression);
          setState(268);

          if (!(precpred(_ctx, 11))) throw FailedPredicateException(this, "precpred(_ctx, 11)");
          setState(269);
          dynamic_cast<ExpressionContext *>(_localctx)->bop = match(TuringParser::EXP);
          setState(270);
          expression(12);
          break;
        }

        case 2: {
          _localctx = _tracker.createInstance<ExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpression);
          setState(271);

          if (!(precpred(_ctx, 8))) throw FailedPredicateException(this, "precpred(_ctx, 8)");
          setState(272);
          dynamic_cast<ExpressionContext *>(_localctx)->bop = _input->LT(1);
          _la = _input->LA(1);
          if (!(((((_la - 74) & ~ 0x3fULL) == 0) &&
            ((1ULL << (_la - 74)) & ((1ULL << (TuringParser::MULTIPLY - 74))
            | (1ULL << (TuringParser::DIVIDE - 74))
            | (1ULL << (TuringParser::DIV - 74))
            | (1ULL << (TuringParser::MOD - 74))
            | (1ULL << (TuringParser::REM - 74))
            | (1ULL << (TuringParser::SHR - 74))
            | (1ULL << (TuringParser::SHL - 74)))) != 0))) {
            dynamic_cast<ExpressionContext *>(_localctx)->bop = _errHandler->recoverInline(this);
          }
          else {
            _errHandler->reportMatch(this);
            consume();
          }
          setState(273);
          expression(9);
          break;
        }

        case 3: {
          _localctx = _tracker.createInstance<ExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpression);
          setState(274);

          if (!(precpred(_ctx, 7))) throw FailedPredicateException(this, "precpred(_ctx, 7)");
          setState(275);
          dynamic_cast<ExpressionContext *>(_localctx)->bop = _input->LT(1);
          _la = _input->LA(1);
          if (!(((((_la - 72) & ~ 0x3fULL) == 0) &&
            ((1ULL << (_la - 72)) & ((1ULL << (TuringParser::PLUS - 72))
            | (1ULL << (TuringParser::MINUS - 72))
            | (1ULL << (TuringParser::XOR - 72)))) != 0))) {
            dynamic_cast<ExpressionContext *>(_localctx)->bop = _errHandler->recoverInline(this);
          }
          else {
            _errHandler->reportMatch(this);
            consume();
          }
          setState(276);
          expression(8);
          break;
        }

        case 4: {
          _localctx = _tracker.createInstance<ExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpression);
          setState(277);

          if (!(precpred(_ctx, 6))) throw FailedPredicateException(this, "precpred(_ctx, 6)");
          setState(278);
          dynamic_cast<ExpressionContext *>(_localctx)->bop = _input->LT(1);
          _la = _input->LA(1);
          if (!(_la == TuringParser::NOT_IN || ((((_la - 80) & ~ 0x3fULL) == 0) &&
            ((1ULL << (_la - 80)) & ((1ULL << (TuringParser::LESSTHAN - 80))
            | (1ULL << (TuringParser::GREATERTHAN - 80))
            | (1ULL << (TuringParser::EQUAL - 80))
            | (1ULL << (TuringParser::LESSTHANEQUAL - 80))
            | (1ULL << (TuringParser::GREATERTHANEQUAL - 80))
            | (1ULL << (TuringParser::NOTEQUAL - 80))
            | (1ULL << (TuringParser::IN - 80)))) != 0))) {
            dynamic_cast<ExpressionContext *>(_localctx)->bop = _errHandler->recoverInline(this);
          }
          else {
            _errHandler->reportMatch(this);
            consume();
          }
          setState(279);
          expression(7);
          break;
        }

        case 5: {
          _localctx = _tracker.createInstance<ExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpression);
          setState(280);

          if (!(precpred(_ctx, 4))) throw FailedPredicateException(this, "precpred(_ctx, 4)");
          setState(281);
          dynamic_cast<ExpressionContext *>(_localctx)->bop = match(TuringParser::AND);
          setState(282);
          expression(5);
          break;
        }

        case 6: {
          _localctx = _tracker.createInstance<ExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpression);
          setState(283);

          if (!(precpred(_ctx, 3))) throw FailedPredicateException(this, "precpred(_ctx, 3)");
          setState(284);
          dynamic_cast<ExpressionContext *>(_localctx)->bop = match(TuringParser::OR);
          setState(285);
          expression(4);
          break;
        }

        case 7: {
          _localctx = _tracker.createInstance<ExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpression);
          setState(286);

          if (!(precpred(_ctx, 2))) throw FailedPredicateException(this, "precpred(_ctx, 2)");
          setState(287);
          dynamic_cast<ExpressionContext *>(_localctx)->bop = match(TuringParser::IMPLIES);
          setState(288);
          expression(3);
          break;
        }

        case 8: {
          _localctx = _tracker.createInstance<ExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpression);
          setState(289);

          if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
          setState(290);
          dynamic_cast<ExpressionContext *>(_localctx)->bop = _input->LT(1);
          _la = _input->LA(1);
          if (!(((((_la - 71) & ~ 0x3fULL) == 0) &&
            ((1ULL << (_la - 71)) & ((1ULL << (TuringParser::ASSIGNMENT - 71))
            | (1ULL << (TuringParser::PLUSEQUALS - 71))
            | (1ULL << (TuringParser::MINUSEQUALS - 71))
            | (1ULL << (TuringParser::MULTIPLYEQUALS - 71))
            | (1ULL << (TuringParser::DIVIDEEQUALS - 71))
            | (1ULL << (TuringParser::DIVEQUALS - 71))
            | (1ULL << (TuringParser::SHREQUALS - 71))
            | (1ULL << (TuringParser::SHLEQUALS - 71)))) != 0))) {
            dynamic_cast<ExpressionContext *>(_localctx)->bop = _errHandler->recoverInline(this);
          }
          else {
            _errHandler->reportMatch(this);
            consume();
          }
          setState(291);
          expression(1);
          break;
        }

        case 9: {
          _localctx = _tracker.createInstance<ExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpression);
          setState(292);

          if (!(precpred(_ctx, 12))) throw FailedPredicateException(this, "precpred(_ctx, 12)");
          setState(293);
          dynamic_cast<ExpressionContext *>(_localctx)->bop = _input->LT(1);
          _la = _input->LA(1);
          if (!(_la == TuringParser::DOT

          || _la == TuringParser::DEREFERENCE)) {
            dynamic_cast<ExpressionContext *>(_localctx)->bop = _errHandler->recoverInline(this);
          }
          else {
            _errHandler->reportMatch(this);
            consume();
          }
          setState(294);
          match(TuringParser::IDENTIFIER);
          break;
        }

        case 10: {
          _localctx = _tracker.createInstance<ExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpression);
          setState(295);

          if (!(precpred(_ctx, 10))) throw FailedPredicateException(this, "precpred(_ctx, 10)");
          setState(296);
          match(TuringParser::L_PAREN);
          setState(298);
          _errHandler->sync(this);

          _la = _input->LA(1);
          if ((((_la & ~ 0x3fULL) == 0) &&
            ((1ULL << _la) & ((1ULL << TuringParser::STRING_LITERAL)
            | (1ULL << TuringParser::INTEGER_LITERAL)
            | (1ULL << TuringParser::REAL_LITERAL)
            | (1ULL << TuringParser::NOT)
            | (1ULL << TuringParser::CARET))) != 0) || ((((_la - 64) & ~ 0x3fULL) == 0) &&
            ((1ULL << (_la - 64)) & ((1ULL << (TuringParser::L_PAREN - 64))
            | (1ULL << (TuringParser::CHEAT - 64))
            | (1ULL << (TuringParser::PLUS - 64))
            | (1ULL << (TuringParser::MINUS - 64))
            | (1ULL << (TuringParser::IDENTIFIER - 64)))) != 0)) {
            setState(297);
            expressionList();
          }
          setState(300);
          match(TuringParser::R_PAREN);
          break;
        }

        } 
      }
      setState(305);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 28, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- ExpressionListContext ------------------------------------------------------------------

TuringParser::ExpressionListContext::ExpressionListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<TuringParser::ExpressionContext *> TuringParser::ExpressionListContext::expression() {
  return getRuleContexts<TuringParser::ExpressionContext>();
}

TuringParser::ExpressionContext* TuringParser::ExpressionListContext::expression(size_t i) {
  return getRuleContext<TuringParser::ExpressionContext>(i);
}

std::vector<tree::TerminalNode *> TuringParser::ExpressionListContext::COMMA() {
  return getTokens(TuringParser::COMMA);
}

tree::TerminalNode* TuringParser::ExpressionListContext::COMMA(size_t i) {
  return getToken(TuringParser::COMMA, i);
}


size_t TuringParser::ExpressionListContext::getRuleIndex() const {
  return TuringParser::RuleExpressionList;
}

void TuringParser::ExpressionListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExpressionList(this);
}

void TuringParser::ExpressionListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExpressionList(this);
}


antlrcpp::Any TuringParser::ExpressionListContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitExpressionList(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ExpressionListContext* TuringParser::expressionList() {
  ExpressionListContext *_localctx = _tracker.createInstance<ExpressionListContext>(_ctx, getState());
  enterRule(_localctx, 44, TuringParser::RuleExpressionList);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(306);
    expression(0);
    setState(311);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while (_la == TuringParser::COMMA) {
      setState(307);
      match(TuringParser::COMMA);
      setState(308);
      expression(0);
      setState(313);
      _errHandler->sync(this);
      _la = _input->LA(1);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- DeclarationContext ------------------------------------------------------------------

TuringParser::DeclarationContext::DeclarationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::TypeDeclarationContext* TuringParser::DeclarationContext::typeDeclaration() {
  return getRuleContext<TuringParser::TypeDeclarationContext>(0);
}

TuringParser::VariableDeclarationContext* TuringParser::DeclarationContext::variableDeclaration() {
  return getRuleContext<TuringParser::VariableDeclarationContext>(0);
}

TuringParser::ArrayDeclarationContext* TuringParser::DeclarationContext::arrayDeclaration() {
  return getRuleContext<TuringParser::ArrayDeclarationContext>(0);
}

TuringParser::SubprogramDeclarationContext* TuringParser::DeclarationContext::subprogramDeclaration() {
  return getRuleContext<TuringParser::SubprogramDeclarationContext>(0);
}

TuringParser::ClassDeclarationContext* TuringParser::DeclarationContext::classDeclaration() {
  return getRuleContext<TuringParser::ClassDeclarationContext>(0);
}


size_t TuringParser::DeclarationContext::getRuleIndex() const {
  return TuringParser::RuleDeclaration;
}

void TuringParser::DeclarationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterDeclaration(this);
}

void TuringParser::DeclarationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitDeclaration(this);
}


antlrcpp::Any TuringParser::DeclarationContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitDeclaration(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::DeclarationContext* TuringParser::declaration() {
  DeclarationContext *_localctx = _tracker.createInstance<DeclarationContext>(_ctx, getState());
  enterRule(_localctx, 46, TuringParser::RuleDeclaration);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(319);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::TYPE: {
        enterOuterAlt(_localctx, 1);
        setState(314);
        typeDeclaration();
        break;
      }

      case TuringParser::VAR: {
        enterOuterAlt(_localctx, 2);
        setState(315);
        variableDeclaration();
        break;
      }

      case TuringParser::ARRAY: {
        enterOuterAlt(_localctx, 3);
        setState(316);
        arrayDeclaration();
        break;
      }

      case TuringParser::PROCEDURE:
      case TuringParser::FUNCTION:
      case TuringParser::PROCESS:
      case TuringParser::DEFERRED:
      case TuringParser::FORWARD:
      case TuringParser::BODY: {
        enterOuterAlt(_localctx, 4);
        setState(317);
        subprogramDeclaration();
        break;
      }

      case TuringParser::CLASS: {
        enterOuterAlt(_localctx, 5);
        setState(318);
        classDeclaration();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- StatementsContext ------------------------------------------------------------------

TuringParser::StatementsContext::StatementsContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::ExpressionContext* TuringParser::StatementsContext::expression() {
  return getRuleContext<TuringParser::ExpressionContext>(0);
}

TuringParser::PutStatementContext* TuringParser::StatementsContext::putStatement() {
  return getRuleContext<TuringParser::PutStatementContext>(0);
}

tree::TerminalNode* TuringParser::StatementsContext::EXIT() {
  return getToken(TuringParser::EXIT, 0);
}

TuringParser::BeginStatementContext* TuringParser::StatementsContext::beginStatement() {
  return getRuleContext<TuringParser::BeginStatementContext>(0);
}

tree::TerminalNode* TuringParser::StatementsContext::RETURN() {
  return getToken(TuringParser::RETURN, 0);
}

TuringParser::ResultStatementContext* TuringParser::StatementsContext::resultStatement() {
  return getRuleContext<TuringParser::ResultStatementContext>(0);
}

TuringParser::NewStatementContext* TuringParser::StatementsContext::newStatement() {
  return getRuleContext<TuringParser::NewStatementContext>(0);
}

TuringParser::FreeStatementContext* TuringParser::StatementsContext::freeStatement() {
  return getRuleContext<TuringParser::FreeStatementContext>(0);
}

TuringParser::ForkStatementContext* TuringParser::StatementsContext::forkStatement() {
  return getRuleContext<TuringParser::ForkStatementContext>(0);
}


size_t TuringParser::StatementsContext::getRuleIndex() const {
  return TuringParser::RuleStatements;
}

void TuringParser::StatementsContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterStatements(this);
}

void TuringParser::StatementsContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitStatements(this);
}


antlrcpp::Any TuringParser::StatementsContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitStatements(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::StatementsContext* TuringParser::statements() {
  StatementsContext *_localctx = _tracker.createInstance<StatementsContext>(_ctx, getState());
  enterRule(_localctx, 48, TuringParser::RuleStatements);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(330);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL:
      case TuringParser::REAL_LITERAL:
      case TuringParser::NOT:
      case TuringParser::CARET:
      case TuringParser::L_PAREN:
      case TuringParser::CHEAT:
      case TuringParser::PLUS:
      case TuringParser::MINUS:
      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 1);
        setState(321);
        expression(0);
        break;
      }

      case TuringParser::PUT: {
        enterOuterAlt(_localctx, 2);
        setState(322);
        putStatement();
        break;
      }

      case TuringParser::EXIT: {
        enterOuterAlt(_localctx, 3);
        setState(323);
        match(TuringParser::EXIT);
        break;
      }

      case TuringParser::BEGIN: {
        enterOuterAlt(_localctx, 4);
        setState(324);
        beginStatement();
        break;
      }

      case TuringParser::RETURN: {
        enterOuterAlt(_localctx, 5);
        setState(325);
        match(TuringParser::RETURN);
        break;
      }

      case TuringParser::RESULT: {
        enterOuterAlt(_localctx, 6);
        setState(326);
        resultStatement();
        break;
      }

      case TuringParser::NEW: {
        enterOuterAlt(_localctx, 7);
        setState(327);
        newStatement();
        break;
      }

      case TuringParser::FREE: {
        enterOuterAlt(_localctx, 8);
        setState(328);
        freeStatement();
        break;
      }

      case TuringParser::FORK: {
        enterOuterAlt(_localctx, 9);
        setState(329);
        forkStatement();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- AssignmentOperatorContext ------------------------------------------------------------------

TuringParser::AssignmentOperatorContext::AssignmentOperatorContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::AssignmentOperatorContext::ASSIGNMENT() {
  return getToken(TuringParser::ASSIGNMENT, 0);
}

tree::TerminalNode* TuringParser::AssignmentOperatorContext::PLUSEQUALS() {
  return getToken(TuringParser::PLUSEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentOperatorContext::MINUSEQUALS() {
  return getToken(TuringParser::MINUSEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentOperatorContext::MULTIPLYEQUALS() {
  return getToken(TuringParser::MULTIPLYEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentOperatorContext::DIVIDEEQUALS() {
  return getToken(TuringParser::DIVIDEEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentOperatorContext::DIVEQUALS() {
  return getToken(TuringParser::DIVEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentOperatorContext::SHLEQUALS() {
  return getToken(TuringParser::SHLEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentOperatorContext::SHREQUALS() {
  return getToken(TuringParser::SHREQUALS, 0);
}


size_t TuringParser::AssignmentOperatorContext::getRuleIndex() const {
  return TuringParser::RuleAssignmentOperator;
}

void TuringParser::AssignmentOperatorContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAssignmentOperator(this);
}

void TuringParser::AssignmentOperatorContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAssignmentOperator(this);
}


antlrcpp::Any TuringParser::AssignmentOperatorContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitAssignmentOperator(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::AssignmentOperatorContext* TuringParser::assignmentOperator() {
  AssignmentOperatorContext *_localctx = _tracker.createInstance<AssignmentOperatorContext>(_ctx, getState());
  enterRule(_localctx, 50, TuringParser::RuleAssignmentOperator);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(332);
    _la = _input->LA(1);
    if (!(((((_la - 71) & ~ 0x3fULL) == 0) &&
      ((1ULL << (_la - 71)) & ((1ULL << (TuringParser::ASSIGNMENT - 71))
      | (1ULL << (TuringParser::PLUSEQUALS - 71))
      | (1ULL << (TuringParser::MINUSEQUALS - 71))
      | (1ULL << (TuringParser::MULTIPLYEQUALS - 71))
      | (1ULL << (TuringParser::DIVIDEEQUALS - 71))
      | (1ULL << (TuringParser::DIVEQUALS - 71))
      | (1ULL << (TuringParser::SHREQUALS - 71))
      | (1ULL << (TuringParser::SHLEQUALS - 71)))) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PutStatementContext ------------------------------------------------------------------

TuringParser::PutStatementContext::PutStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::PutStatementContext::PUT() {
  return getToken(TuringParser::PUT, 0);
}

TuringParser::PutItemListContext* TuringParser::PutStatementContext::putItemList() {
  return getRuleContext<TuringParser::PutItemListContext>(0);
}


size_t TuringParser::PutStatementContext::getRuleIndex() const {
  return TuringParser::RulePutStatement;
}

void TuringParser::PutStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterPutStatement(this);
}

void TuringParser::PutStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitPutStatement(this);
}


antlrcpp::Any TuringParser::PutStatementContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitPutStatement(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::PutStatementContext* TuringParser::putStatement() {
  PutStatementContext *_localctx = _tracker.createInstance<PutStatementContext>(_ctx, getState());
  enterRule(_localctx, 52, TuringParser::RulePutStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(334);
    match(TuringParser::PUT);
    setState(335);
    putItemList(0);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PutItemContext ------------------------------------------------------------------

TuringParser::PutItemContext::PutItemContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::StatementsContext* TuringParser::PutItemContext::statements() {
  return getRuleContext<TuringParser::StatementsContext>(0);
}


size_t TuringParser::PutItemContext::getRuleIndex() const {
  return TuringParser::RulePutItem;
}

void TuringParser::PutItemContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterPutItem(this);
}

void TuringParser::PutItemContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitPutItem(this);
}


antlrcpp::Any TuringParser::PutItemContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitPutItem(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::PutItemContext* TuringParser::putItem() {
  PutItemContext *_localctx = _tracker.createInstance<PutItemContext>(_ctx, getState());
  enterRule(_localctx, 54, TuringParser::RulePutItem);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(337);
    statements();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PutItemListContext ------------------------------------------------------------------

TuringParser::PutItemListContext::PutItemListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::PutItemContext* TuringParser::PutItemListContext::putItem() {
  return getRuleContext<TuringParser::PutItemContext>(0);
}

TuringParser::PutItemListContext* TuringParser::PutItemListContext::putItemList() {
  return getRuleContext<TuringParser::PutItemListContext>(0);
}

tree::TerminalNode* TuringParser::PutItemListContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
}


size_t TuringParser::PutItemListContext::getRuleIndex() const {
  return TuringParser::RulePutItemList;
}

void TuringParser::PutItemListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterPutItemList(this);
}

void TuringParser::PutItemListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitPutItemList(this);
}


antlrcpp::Any TuringParser::PutItemListContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitPutItemList(this);
  else
    return visitor->visitChildren(this);
}


TuringParser::PutItemListContext* TuringParser::putItemList() {
   return putItemList(0);
}

TuringParser::PutItemListContext* TuringParser::putItemList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::PutItemListContext *_localctx = _tracker.createInstance<PutItemListContext>(_ctx, parentState);
  TuringParser::PutItemListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 56;
  enterRecursionRule(_localctx, 56, TuringParser::RulePutItemList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(340);
    putItem();
    _ctx->stop = _input->LT(-1);
    setState(347);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 32, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<PutItemListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RulePutItemList);
        setState(342);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(343);
        match(TuringParser::COMMA);
        setState(344);
        putItem(); 
      }
      setState(349);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 32, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- BeginStatementContext ------------------------------------------------------------------

TuringParser::BeginStatementContext::BeginStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::BeginStatementContext::BEGIN() {
  return getToken(TuringParser::BEGIN, 0);
}

TuringParser::StatementOrDeclarationContext* TuringParser::BeginStatementContext::statementOrDeclaration() {
  return getRuleContext<TuringParser::StatementOrDeclarationContext>(0);
}

tree::TerminalNode* TuringParser::BeginStatementContext::END() {
  return getToken(TuringParser::END, 0);
}


size_t TuringParser::BeginStatementContext::getRuleIndex() const {
  return TuringParser::RuleBeginStatement;
}

void TuringParser::BeginStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterBeginStatement(this);
}

void TuringParser::BeginStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitBeginStatement(this);
}


antlrcpp::Any TuringParser::BeginStatementContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitBeginStatement(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::BeginStatementContext* TuringParser::beginStatement() {
  BeginStatementContext *_localctx = _tracker.createInstance<BeginStatementContext>(_ctx, getState());
  enterRule(_localctx, 58, TuringParser::RuleBeginStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(350);
    match(TuringParser::BEGIN);
    setState(351);
    statementOrDeclaration();
    setState(352);
    match(TuringParser::END);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ResultStatementContext ------------------------------------------------------------------

TuringParser::ResultStatementContext::ResultStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::ResultStatementContext::RESULT() {
  return getToken(TuringParser::RESULT, 0);
}

TuringParser::ExpressionContext* TuringParser::ResultStatementContext::expression() {
  return getRuleContext<TuringParser::ExpressionContext>(0);
}


size_t TuringParser::ResultStatementContext::getRuleIndex() const {
  return TuringParser::RuleResultStatement;
}

void TuringParser::ResultStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterResultStatement(this);
}

void TuringParser::ResultStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitResultStatement(this);
}


antlrcpp::Any TuringParser::ResultStatementContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitResultStatement(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ResultStatementContext* TuringParser::resultStatement() {
  ResultStatementContext *_localctx = _tracker.createInstance<ResultStatementContext>(_ctx, getState());
  enterRule(_localctx, 60, TuringParser::RuleResultStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(354);
    match(TuringParser::RESULT);
    setState(355);
    expression(0);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- NewStatementContext ------------------------------------------------------------------

TuringParser::NewStatementContext::NewStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::NewStatementContext::NEW() {
  return getToken(TuringParser::NEW, 0);
}

std::vector<tree::TerminalNode *> TuringParser::NewStatementContext::IDENTIFIER() {
  return getTokens(TuringParser::IDENTIFIER);
}

tree::TerminalNode* TuringParser::NewStatementContext::IDENTIFIER(size_t i) {
  return getToken(TuringParser::IDENTIFIER, i);
}

tree::TerminalNode* TuringParser::NewStatementContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
}


size_t TuringParser::NewStatementContext::getRuleIndex() const {
  return TuringParser::RuleNewStatement;
}

void TuringParser::NewStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterNewStatement(this);
}

void TuringParser::NewStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitNewStatement(this);
}


antlrcpp::Any TuringParser::NewStatementContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitNewStatement(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::NewStatementContext* TuringParser::newStatement() {
  NewStatementContext *_localctx = _tracker.createInstance<NewStatementContext>(_ctx, getState());
  enterRule(_localctx, 62, TuringParser::RuleNewStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(357);
    match(TuringParser::NEW);
    setState(360);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 33, _ctx)) {
    case 1: {
      setState(358);
      match(TuringParser::IDENTIFIER);
      setState(359);
      match(TuringParser::COMMA);
      break;
    }

    }
    setState(362);
    match(TuringParser::IDENTIFIER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- FreeStatementContext ------------------------------------------------------------------

TuringParser::FreeStatementContext::FreeStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::FreeStatementContext::FREE() {
  return getToken(TuringParser::FREE, 0);
}

std::vector<tree::TerminalNode *> TuringParser::FreeStatementContext::IDENTIFIER() {
  return getTokens(TuringParser::IDENTIFIER);
}

tree::TerminalNode* TuringParser::FreeStatementContext::IDENTIFIER(size_t i) {
  return getToken(TuringParser::IDENTIFIER, i);
}

tree::TerminalNode* TuringParser::FreeStatementContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
}


size_t TuringParser::FreeStatementContext::getRuleIndex() const {
  return TuringParser::RuleFreeStatement;
}

void TuringParser::FreeStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterFreeStatement(this);
}

void TuringParser::FreeStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitFreeStatement(this);
}


antlrcpp::Any TuringParser::FreeStatementContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitFreeStatement(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::FreeStatementContext* TuringParser::freeStatement() {
  FreeStatementContext *_localctx = _tracker.createInstance<FreeStatementContext>(_ctx, getState());
  enterRule(_localctx, 64, TuringParser::RuleFreeStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(364);
    match(TuringParser::FREE);
    setState(367);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 34, _ctx)) {
    case 1: {
      setState(365);
      match(TuringParser::IDENTIFIER);
      setState(366);
      match(TuringParser::COMMA);
      break;
    }

    }
    setState(369);
    match(TuringParser::IDENTIFIER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ForkStatementContext ------------------------------------------------------------------

TuringParser::ForkStatementContext::ForkStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::ForkStatementContext::FORK() {
  return getToken(TuringParser::FORK, 0);
}

TuringParser::ExpressionContext* TuringParser::ForkStatementContext::expression() {
  return getRuleContext<TuringParser::ExpressionContext>(0);
}


size_t TuringParser::ForkStatementContext::getRuleIndex() const {
  return TuringParser::RuleForkStatement;
}

void TuringParser::ForkStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterForkStatement(this);
}

void TuringParser::ForkStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitForkStatement(this);
}


antlrcpp::Any TuringParser::ForkStatementContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitForkStatement(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ForkStatementContext* TuringParser::forkStatement() {
  ForkStatementContext *_localctx = _tracker.createInstance<ForkStatementContext>(_ctx, getState());
  enterRule(_localctx, 66, TuringParser::RuleForkStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(371);
    match(TuringParser::FORK);
    setState(372);
    expression(0);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TypeDeclarationContext ------------------------------------------------------------------

TuringParser::TypeDeclarationContext::TypeDeclarationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::TypeDeclarationContext::TYPE() {
  return getToken(TuringParser::TYPE, 0);
}

tree::TerminalNode* TuringParser::TypeDeclarationContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}

tree::TerminalNode* TuringParser::TypeDeclarationContext::COLON() {
  return getToken(TuringParser::COLON, 0);
}

TuringParser::TypeSpecContext* TuringParser::TypeDeclarationContext::typeSpec() {
  return getRuleContext<TuringParser::TypeSpecContext>(0);
}


size_t TuringParser::TypeDeclarationContext::getRuleIndex() const {
  return TuringParser::RuleTypeDeclaration;
}

void TuringParser::TypeDeclarationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTypeDeclaration(this);
}

void TuringParser::TypeDeclarationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTypeDeclaration(this);
}


antlrcpp::Any TuringParser::TypeDeclarationContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitTypeDeclaration(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::TypeDeclarationContext* TuringParser::typeDeclaration() {
  TypeDeclarationContext *_localctx = _tracker.createInstance<TypeDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 68, TuringParser::RuleTypeDeclaration);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(374);
    match(TuringParser::TYPE);
    setState(375);
    match(TuringParser::IDENTIFIER);
    setState(376);
    match(TuringParser::COLON);
    setState(377);
    typeSpec();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- BasicTypeContext ------------------------------------------------------------------

TuringParser::BasicTypeContext::BasicTypeContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::BasicTypeContext::INT() {
  return getToken(TuringParser::INT, 0);
}

tree::TerminalNode* TuringParser::BasicTypeContext::REAL() {
  return getToken(TuringParser::REAL, 0);
}

tree::TerminalNode* TuringParser::BasicTypeContext::BOOLEAN() {
  return getToken(TuringParser::BOOLEAN, 0);
}

tree::TerminalNode* TuringParser::BasicTypeContext::NAT() {
  return getToken(TuringParser::NAT, 0);
}

tree::TerminalNode* TuringParser::BasicTypeContext::INTN() {
  return getToken(TuringParser::INTN, 0);
}

tree::TerminalNode* TuringParser::BasicTypeContext::NATN() {
  return getToken(TuringParser::NATN, 0);
}

tree::TerminalNode* TuringParser::BasicTypeContext::REALN() {
  return getToken(TuringParser::REALN, 0);
}

tree::TerminalNode* TuringParser::BasicTypeContext::CHAR() {
  return getToken(TuringParser::CHAR, 0);
}


size_t TuringParser::BasicTypeContext::getRuleIndex() const {
  return TuringParser::RuleBasicType;
}

void TuringParser::BasicTypeContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterBasicType(this);
}

void TuringParser::BasicTypeContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitBasicType(this);
}


antlrcpp::Any TuringParser::BasicTypeContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitBasicType(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::BasicTypeContext* TuringParser::basicType() {
  BasicTypeContext *_localctx = _tracker.createInstance<BasicTypeContext>(_ctx, getState());
  enterRule(_localctx, 70, TuringParser::RuleBasicType);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(379);
    _la = _input->LA(1);
    if (!((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::INT)
      | (1ULL << TuringParser::REAL)
      | (1ULL << TuringParser::BOOLEAN)
      | (1ULL << TuringParser::NAT)
      | (1ULL << TuringParser::INTN)
      | (1ULL << TuringParser::NATN)
      | (1ULL << TuringParser::REALN)
      | (1ULL << TuringParser::CHAR))) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ReferenceTypeContext ------------------------------------------------------------------

TuringParser::ReferenceTypeContext::ReferenceTypeContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::ReferenceTypeContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}


size_t TuringParser::ReferenceTypeContext::getRuleIndex() const {
  return TuringParser::RuleReferenceType;
}

void TuringParser::ReferenceTypeContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterReferenceType(this);
}

void TuringParser::ReferenceTypeContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitReferenceType(this);
}


antlrcpp::Any TuringParser::ReferenceTypeContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitReferenceType(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ReferenceTypeContext* TuringParser::referenceType() {
  ReferenceTypeContext *_localctx = _tracker.createInstance<ReferenceTypeContext>(_ctx, getState());
  enterRule(_localctx, 72, TuringParser::RuleReferenceType);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(381);
    match(TuringParser::IDENTIFIER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TypeSpecContext ------------------------------------------------------------------

TuringParser::TypeSpecContext::TypeSpecContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::BasicTypeContext* TuringParser::TypeSpecContext::basicType() {
  return getRuleContext<TuringParser::BasicTypeContext>(0);
}

TuringParser::IndexTypeContext* TuringParser::TypeSpecContext::indexType() {
  return getRuleContext<TuringParser::IndexTypeContext>(0);
}

TuringParser::StringTypeContext* TuringParser::TypeSpecContext::stringType() {
  return getRuleContext<TuringParser::StringTypeContext>(0);
}

TuringParser::RecordTypeContext* TuringParser::TypeSpecContext::recordType() {
  return getRuleContext<TuringParser::RecordTypeContext>(0);
}

TuringParser::ArrayDeclarationContext* TuringParser::TypeSpecContext::arrayDeclaration() {
  return getRuleContext<TuringParser::ArrayDeclarationContext>(0);
}

TuringParser::ClassDeclarationContext* TuringParser::TypeSpecContext::classDeclaration() {
  return getRuleContext<TuringParser::ClassDeclarationContext>(0);
}

TuringParser::ReferenceTypeContext* TuringParser::TypeSpecContext::referenceType() {
  return getRuleContext<TuringParser::ReferenceTypeContext>(0);
}


size_t TuringParser::TypeSpecContext::getRuleIndex() const {
  return TuringParser::RuleTypeSpec;
}

void TuringParser::TypeSpecContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTypeSpec(this);
}

void TuringParser::TypeSpecContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTypeSpec(this);
}


antlrcpp::Any TuringParser::TypeSpecContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitTypeSpec(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::TypeSpecContext* TuringParser::typeSpec() {
  TypeSpecContext *_localctx = _tracker.createInstance<TypeSpecContext>(_ctx, getState());
  enterRule(_localctx, 74, TuringParser::RuleTypeSpec);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(390);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::INT:
      case TuringParser::REAL:
      case TuringParser::BOOLEAN:
      case TuringParser::NAT:
      case TuringParser::INTN:
      case TuringParser::NATN:
      case TuringParser::REALN:
      case TuringParser::CHAR: {
        enterOuterAlt(_localctx, 1);
        setState(383);
        basicType();
        break;
      }

      case TuringParser::INTEGER_LITERAL: {
        enterOuterAlt(_localctx, 2);
        setState(384);
        indexType();
        break;
      }

      case TuringParser::STRING: {
        enterOuterAlt(_localctx, 3);
        setState(385);
        stringType();
        break;
      }

      case TuringParser::RECORD: {
        enterOuterAlt(_localctx, 4);
        setState(386);
        recordType();
        break;
      }

      case TuringParser::ARRAY: {
        enterOuterAlt(_localctx, 5);
        setState(387);
        arrayDeclaration();
        break;
      }

      case TuringParser::CLASS: {
        enterOuterAlt(_localctx, 6);
        setState(388);
        classDeclaration();
        break;
      }

      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 7);
        setState(389);
        referenceType();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- IndexTypeContext ------------------------------------------------------------------

TuringParser::IndexTypeContext::IndexTypeContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<tree::TerminalNode *> TuringParser::IndexTypeContext::INTEGER_LITERAL() {
  return getTokens(TuringParser::INTEGER_LITERAL);
}

tree::TerminalNode* TuringParser::IndexTypeContext::INTEGER_LITERAL(size_t i) {
  return getToken(TuringParser::INTEGER_LITERAL, i);
}

tree::TerminalNode* TuringParser::IndexTypeContext::RANGE() {
  return getToken(TuringParser::RANGE, 0);
}


size_t TuringParser::IndexTypeContext::getRuleIndex() const {
  return TuringParser::RuleIndexType;
}

void TuringParser::IndexTypeContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterIndexType(this);
}

void TuringParser::IndexTypeContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitIndexType(this);
}


antlrcpp::Any TuringParser::IndexTypeContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitIndexType(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::IndexTypeContext* TuringParser::indexType() {
  IndexTypeContext *_localctx = _tracker.createInstance<IndexTypeContext>(_ctx, getState());
  enterRule(_localctx, 76, TuringParser::RuleIndexType);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(392);
    match(TuringParser::INTEGER_LITERAL);
    setState(393);
    match(TuringParser::RANGE);
    setState(394);
    match(TuringParser::INTEGER_LITERAL);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- IndexTypeListContext ------------------------------------------------------------------

TuringParser::IndexTypeListContext::IndexTypeListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::IndexTypeContext* TuringParser::IndexTypeListContext::indexType() {
  return getRuleContext<TuringParser::IndexTypeContext>(0);
}

TuringParser::IndexTypeListContext* TuringParser::IndexTypeListContext::indexTypeList() {
  return getRuleContext<TuringParser::IndexTypeListContext>(0);
}

tree::TerminalNode* TuringParser::IndexTypeListContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
}


size_t TuringParser::IndexTypeListContext::getRuleIndex() const {
  return TuringParser::RuleIndexTypeList;
}

void TuringParser::IndexTypeListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterIndexTypeList(this);
}

void TuringParser::IndexTypeListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitIndexTypeList(this);
}


antlrcpp::Any TuringParser::IndexTypeListContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitIndexTypeList(this);
  else
    return visitor->visitChildren(this);
}


TuringParser::IndexTypeListContext* TuringParser::indexTypeList() {
   return indexTypeList(0);
}

TuringParser::IndexTypeListContext* TuringParser::indexTypeList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::IndexTypeListContext *_localctx = _tracker.createInstance<IndexTypeListContext>(_ctx, parentState);
  TuringParser::IndexTypeListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 78;
  enterRecursionRule(_localctx, 78, TuringParser::RuleIndexTypeList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(397);
    indexType();
    _ctx->stop = _input->LT(-1);
    setState(404);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 36, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<IndexTypeListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleIndexTypeList);
        setState(399);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(400);
        match(TuringParser::COMMA);
        setState(401);
        indexType(); 
      }
      setState(406);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 36, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- StringTypeContext ------------------------------------------------------------------

TuringParser::StringTypeContext::StringTypeContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::StringTypeContext::STRING() {
  return getToken(TuringParser::STRING, 0);
}

tree::TerminalNode* TuringParser::StringTypeContext::L_PAREN() {
  return getToken(TuringParser::L_PAREN, 0);
}

tree::TerminalNode* TuringParser::StringTypeContext::INTEGER_LITERAL() {
  return getToken(TuringParser::INTEGER_LITERAL, 0);
}

tree::TerminalNode* TuringParser::StringTypeContext::R_PAREN() {
  return getToken(TuringParser::R_PAREN, 0);
}


size_t TuringParser::StringTypeContext::getRuleIndex() const {
  return TuringParser::RuleStringType;
}

void TuringParser::StringTypeContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterStringType(this);
}

void TuringParser::StringTypeContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitStringType(this);
}


antlrcpp::Any TuringParser::StringTypeContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitStringType(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::StringTypeContext* TuringParser::stringType() {
  StringTypeContext *_localctx = _tracker.createInstance<StringTypeContext>(_ctx, getState());
  enterRule(_localctx, 80, TuringParser::RuleStringType);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(407);
    match(TuringParser::STRING);
    setState(411);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 37, _ctx)) {
    case 1: {
      setState(408);
      match(TuringParser::L_PAREN);
      setState(409);
      match(TuringParser::INTEGER_LITERAL);
      setState(410);
      match(TuringParser::R_PAREN);
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- RecordTypeContext ------------------------------------------------------------------

TuringParser::RecordTypeContext::RecordTypeContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<tree::TerminalNode *> TuringParser::RecordTypeContext::RECORD() {
  return getTokens(TuringParser::RECORD);
}

tree::TerminalNode* TuringParser::RecordTypeContext::RECORD(size_t i) {
  return getToken(TuringParser::RECORD, i);
}

tree::TerminalNode* TuringParser::RecordTypeContext::END() {
  return getToken(TuringParser::END, 0);
}

std::vector<TuringParser::RecordFieldContext *> TuringParser::RecordTypeContext::recordField() {
  return getRuleContexts<TuringParser::RecordFieldContext>();
}

TuringParser::RecordFieldContext* TuringParser::RecordTypeContext::recordField(size_t i) {
  return getRuleContext<TuringParser::RecordFieldContext>(i);
}


size_t TuringParser::RecordTypeContext::getRuleIndex() const {
  return TuringParser::RuleRecordType;
}

void TuringParser::RecordTypeContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterRecordType(this);
}

void TuringParser::RecordTypeContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitRecordType(this);
}


antlrcpp::Any TuringParser::RecordTypeContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitRecordType(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::RecordTypeContext* TuringParser::recordType() {
  RecordTypeContext *_localctx = _tracker.createInstance<RecordTypeContext>(_ctx, getState());
  enterRule(_localctx, 82, TuringParser::RuleRecordType);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(413);
    match(TuringParser::RECORD);
    setState(415); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(414);
      recordField();
      setState(417); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while (_la == TuringParser::IDENTIFIER);
    setState(419);
    match(TuringParser::END);
    setState(420);
    match(TuringParser::RECORD);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- RecordFieldContext ------------------------------------------------------------------

TuringParser::RecordFieldContext::RecordFieldContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::IdentifierListContext* TuringParser::RecordFieldContext::identifierList() {
  return getRuleContext<TuringParser::IdentifierListContext>(0);
}

tree::TerminalNode* TuringParser::RecordFieldContext::COLON() {
  return getToken(TuringParser::COLON, 0);
}

TuringParser::TypeSpecContext* TuringParser::RecordFieldContext::typeSpec() {
  return getRuleContext<TuringParser::TypeSpecContext>(0);
}


size_t TuringParser::RecordFieldContext::getRuleIndex() const {
  return TuringParser::RuleRecordField;
}

void TuringParser::RecordFieldContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterRecordField(this);
}

void TuringParser::RecordFieldContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitRecordField(this);
}


antlrcpp::Any TuringParser::RecordFieldContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitRecordField(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::RecordFieldContext* TuringParser::recordField() {
  RecordFieldContext *_localctx = _tracker.createInstance<RecordFieldContext>(_ctx, getState());
  enterRule(_localctx, 84, TuringParser::RuleRecordField);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(422);
    identifierList(0);
    setState(423);
    match(TuringParser::COLON);
    setState(424);
    typeSpec();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- VariableDeclarationContext ------------------------------------------------------------------

TuringParser::VariableDeclarationContext::VariableDeclarationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::VariableDeclarationContext::VAR() {
  return getToken(TuringParser::VAR, 0);
}

TuringParser::VariableIdentifierListContext* TuringParser::VariableDeclarationContext::variableIdentifierList() {
  return getRuleContext<TuringParser::VariableIdentifierListContext>(0);
}

tree::TerminalNode* TuringParser::VariableDeclarationContext::COLON() {
  return getToken(TuringParser::COLON, 0);
}

TuringParser::TypeSpecContext* TuringParser::VariableDeclarationContext::typeSpec() {
  return getRuleContext<TuringParser::TypeSpecContext>(0);
}

tree::TerminalNode* TuringParser::VariableDeclarationContext::ASSIGNMENT() {
  return getToken(TuringParser::ASSIGNMENT, 0);
}

TuringParser::ExpressionContext* TuringParser::VariableDeclarationContext::expression() {
  return getRuleContext<TuringParser::ExpressionContext>(0);
}


size_t TuringParser::VariableDeclarationContext::getRuleIndex() const {
  return TuringParser::RuleVariableDeclaration;
}

void TuringParser::VariableDeclarationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterVariableDeclaration(this);
}

void TuringParser::VariableDeclarationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitVariableDeclaration(this);
}


antlrcpp::Any TuringParser::VariableDeclarationContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitVariableDeclaration(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::VariableDeclarationContext* TuringParser::variableDeclaration() {
  VariableDeclarationContext *_localctx = _tracker.createInstance<VariableDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 86, TuringParser::RuleVariableDeclaration);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(439);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 40, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(426);
      match(TuringParser::VAR);
      setState(427);
      variableIdentifierList();
      setState(428);
      match(TuringParser::COLON);
      setState(429);
      typeSpec();
      setState(432);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == TuringParser::ASSIGNMENT) {
        setState(430);
        match(TuringParser::ASSIGNMENT);
        setState(431);
        expression(0);
      }
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(434);
      match(TuringParser::VAR);
      setState(435);
      variableIdentifierList();
      setState(436);
      match(TuringParser::ASSIGNMENT);
      setState(437);
      expression(0);
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- VariableIdentifierListContext ------------------------------------------------------------------

TuringParser::VariableIdentifierListContext::VariableIdentifierListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<TuringParser::VariableIdentifierContext *> TuringParser::VariableIdentifierListContext::variableIdentifier() {
  return getRuleContexts<TuringParser::VariableIdentifierContext>();
}

TuringParser::VariableIdentifierContext* TuringParser::VariableIdentifierListContext::variableIdentifier(size_t i) {
  return getRuleContext<TuringParser::VariableIdentifierContext>(i);
}

std::vector<tree::TerminalNode *> TuringParser::VariableIdentifierListContext::COMMA() {
  return getTokens(TuringParser::COMMA);
}

tree::TerminalNode* TuringParser::VariableIdentifierListContext::COMMA(size_t i) {
  return getToken(TuringParser::COMMA, i);
}


size_t TuringParser::VariableIdentifierListContext::getRuleIndex() const {
  return TuringParser::RuleVariableIdentifierList;
}

void TuringParser::VariableIdentifierListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterVariableIdentifierList(this);
}

void TuringParser::VariableIdentifierListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitVariableIdentifierList(this);
}


antlrcpp::Any TuringParser::VariableIdentifierListContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitVariableIdentifierList(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::VariableIdentifierListContext* TuringParser::variableIdentifierList() {
  VariableIdentifierListContext *_localctx = _tracker.createInstance<VariableIdentifierListContext>(_ctx, getState());
  enterRule(_localctx, 88, TuringParser::RuleVariableIdentifierList);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(441);
    variableIdentifier();
    setState(446);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while (_la == TuringParser::COMMA) {
      setState(442);
      match(TuringParser::COMMA);
      setState(443);
      variableIdentifier();
      setState(448);
      _errHandler->sync(this);
      _la = _input->LA(1);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- VariableIdentifierContext ------------------------------------------------------------------

TuringParser::VariableIdentifierContext::VariableIdentifierContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::VariableIdentifierContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}


size_t TuringParser::VariableIdentifierContext::getRuleIndex() const {
  return TuringParser::RuleVariableIdentifier;
}

void TuringParser::VariableIdentifierContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterVariableIdentifier(this);
}

void TuringParser::VariableIdentifierContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitVariableIdentifier(this);
}


antlrcpp::Any TuringParser::VariableIdentifierContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitVariableIdentifier(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::VariableIdentifierContext* TuringParser::variableIdentifier() {
  VariableIdentifierContext *_localctx = _tracker.createInstance<VariableIdentifierContext>(_ctx, getState());
  enterRule(_localctx, 90, TuringParser::RuleVariableIdentifier);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(449);
    match(TuringParser::IDENTIFIER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ArrayDeclarationContext ------------------------------------------------------------------

TuringParser::ArrayDeclarationContext::ArrayDeclarationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::ArrayDeclarationContext::ARRAY() {
  return getToken(TuringParser::ARRAY, 0);
}

TuringParser::IndexTypeContext* TuringParser::ArrayDeclarationContext::indexType() {
  return getRuleContext<TuringParser::IndexTypeContext>(0);
}

tree::TerminalNode* TuringParser::ArrayDeclarationContext::OF() {
  return getToken(TuringParser::OF, 0);
}

TuringParser::TypeSpecContext* TuringParser::ArrayDeclarationContext::typeSpec() {
  return getRuleContext<TuringParser::TypeSpecContext>(0);
}

tree::TerminalNode* TuringParser::ArrayDeclarationContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
}


size_t TuringParser::ArrayDeclarationContext::getRuleIndex() const {
  return TuringParser::RuleArrayDeclaration;
}

void TuringParser::ArrayDeclarationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterArrayDeclaration(this);
}

void TuringParser::ArrayDeclarationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitArrayDeclaration(this);
}


antlrcpp::Any TuringParser::ArrayDeclarationContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitArrayDeclaration(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::ArrayDeclarationContext* TuringParser::arrayDeclaration() {
  ArrayDeclarationContext *_localctx = _tracker.createInstance<ArrayDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 92, TuringParser::RuleArrayDeclaration);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(460);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 42, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(451);
      match(TuringParser::ARRAY);
      setState(452);
      indexType();
      setState(453);
      match(TuringParser::OF);
      setState(454);
      typeSpec();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(456);
      match(TuringParser::ARRAY);
      setState(457);
      indexType();
      setState(458);
      match(TuringParser::COMMA);
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- IdentifierListContext ------------------------------------------------------------------

TuringParser::IdentifierListContext::IdentifierListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::IdentifierListContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}

TuringParser::IdentifierListContext* TuringParser::IdentifierListContext::identifierList() {
  return getRuleContext<TuringParser::IdentifierListContext>(0);
}

tree::TerminalNode* TuringParser::IdentifierListContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
}


size_t TuringParser::IdentifierListContext::getRuleIndex() const {
  return TuringParser::RuleIdentifierList;
}

void TuringParser::IdentifierListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterIdentifierList(this);
}

void TuringParser::IdentifierListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitIdentifierList(this);
}


antlrcpp::Any TuringParser::IdentifierListContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitIdentifierList(this);
  else
    return visitor->visitChildren(this);
}


TuringParser::IdentifierListContext* TuringParser::identifierList() {
   return identifierList(0);
}

TuringParser::IdentifierListContext* TuringParser::identifierList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::IdentifierListContext *_localctx = _tracker.createInstance<IdentifierListContext>(_ctx, parentState);
  TuringParser::IdentifierListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 94;
  enterRecursionRule(_localctx, 94, TuringParser::RuleIdentifierList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(463);
    match(TuringParser::IDENTIFIER);
    _ctx->stop = _input->LT(-1);
    setState(470);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 43, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<IdentifierListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleIdentifierList);
        setState(465);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(466);
        match(TuringParser::COMMA);
        setState(467);
        match(TuringParser::IDENTIFIER); 
      }
      setState(472);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 43, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- LiteralContext ------------------------------------------------------------------

TuringParser::LiteralContext::LiteralContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::LiteralContext::STRING_LITERAL() {
  return getToken(TuringParser::STRING_LITERAL, 0);
}

tree::TerminalNode* TuringParser::LiteralContext::INTEGER_LITERAL() {
  return getToken(TuringParser::INTEGER_LITERAL, 0);
}

tree::TerminalNode* TuringParser::LiteralContext::REAL_LITERAL() {
  return getToken(TuringParser::REAL_LITERAL, 0);
}


size_t TuringParser::LiteralContext::getRuleIndex() const {
  return TuringParser::RuleLiteral;
}

void TuringParser::LiteralContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterLiteral(this);
}

void TuringParser::LiteralContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitLiteral(this);
}


antlrcpp::Any TuringParser::LiteralContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitLiteral(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::LiteralContext* TuringParser::literal() {
  LiteralContext *_localctx = _tracker.createInstance<LiteralContext>(_ctx, getState());
  enterRule(_localctx, 96, TuringParser::RuleLiteral);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(473);
    _la = _input->LA(1);
    if (!((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::STRING_LITERAL)
      | (1ULL << TuringParser::INTEGER_LITERAL)
      | (1ULL << TuringParser::REAL_LITERAL))) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- CommentContext ------------------------------------------------------------------

TuringParser::CommentContext::CommentContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::CommentContext::BLOCK_COMMENT() {
  return getToken(TuringParser::BLOCK_COMMENT, 0);
}

tree::TerminalNode* TuringParser::CommentContext::LINE_COMMENT() {
  return getToken(TuringParser::LINE_COMMENT, 0);
}


size_t TuringParser::CommentContext::getRuleIndex() const {
  return TuringParser::RuleComment;
}

void TuringParser::CommentContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterComment(this);
}

void TuringParser::CommentContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitComment(this);
}


antlrcpp::Any TuringParser::CommentContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<TuringVisitor*>(visitor))
    return parserVisitor->visitComment(this);
  else
    return visitor->visitChildren(this);
}

TuringParser::CommentContext* TuringParser::comment() {
  CommentContext *_localctx = _tracker.createInstance<CommentContext>(_ctx, getState());
  enterRule(_localctx, 98, TuringParser::RuleComment);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(475);
    _la = _input->LA(1);
    if (!(_la == TuringParser::BLOCK_COMMENT

    || _la == TuringParser::LINE_COMMENT)) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

bool TuringParser::sempred(RuleContext *context, size_t ruleIndex, size_t predicateIndex) {
  switch (ruleIndex) {
    case 11: return exportListSempred(dynamic_cast<ExportListContext *>(context), predicateIndex);
    case 17: return paramDeclarationListSempred(dynamic_cast<ParamDeclarationListContext *>(context), predicateIndex);
    case 21: return expressionSempred(dynamic_cast<ExpressionContext *>(context), predicateIndex);
    case 28: return putItemListSempred(dynamic_cast<PutItemListContext *>(context), predicateIndex);
    case 39: return indexTypeListSempred(dynamic_cast<IndexTypeListContext *>(context), predicateIndex);
    case 47: return identifierListSempred(dynamic_cast<IdentifierListContext *>(context), predicateIndex);

  default:
    break;
  }
  return true;
}

bool TuringParser::exportListSempred(ExportListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 0: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::paramDeclarationListSempred(ParamDeclarationListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 1: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::expressionSempred(ExpressionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 2: return precpred(_ctx, 11);
    case 3: return precpred(_ctx, 8);
    case 4: return precpred(_ctx, 7);
    case 5: return precpred(_ctx, 6);
    case 6: return precpred(_ctx, 4);
    case 7: return precpred(_ctx, 3);
    case 8: return precpred(_ctx, 2);
    case 9: return precpred(_ctx, 1);
    case 10: return precpred(_ctx, 12);
    case 11: return precpred(_ctx, 10);

  default:
    break;
  }
  return true;
}

bool TuringParser::putItemListSempred(PutItemListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 12: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::indexTypeListSempred(IndexTypeListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 13: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::identifierListSempred(IdentifierListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 14: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

// Static vars and initialization.
std::vector<dfa::DFA> TuringParser::_decisionToDFA;
atn::PredictionContextCache TuringParser::_sharedContextCache;

// We own the ATN which in turn owns the ATN states.
atn::ATN TuringParser::_atn;
std::vector<uint16_t> TuringParser::_serializedATN;

std::vector<std::string> TuringParser::_ruleNames = {
  "program", "topLevel", "procHeader", "funcHeader", "subprogramHeader", 
  "subprogramDeclaration", "subprogramPrefix", "classDeclaration", "classBody", 
  "exportStatement", "exportListItem", "exportList", "inheritStatement", 
  "implementStatement", "idOrFileItem", "howExport", "paramDeclaration", 
  "paramDeclarationList", "procBody", "statementOrDeclaration", "primaryExpression", 
  "expression", "expressionList", "declaration", "statements", "assignmentOperator", 
  "putStatement", "putItem", "putItemList", "beginStatement", "resultStatement", 
  "newStatement", "freeStatement", "forkStatement", "typeDeclaration", "basicType", 
  "referenceType", "typeSpec", "indexType", "indexTypeList", "stringType", 
  "recordType", "recordField", "variableDeclaration", "variableIdentifierList", 
  "variableIdentifier", "arrayDeclaration", "identifierList", "literal", 
  "comment"
};

std::vector<std::string> TuringParser::_literalNames = {
  "", "'unqualified'", "'pervasive'", "'opaque'", "", "", "", "", "'end'", 
  "'of'", "'to'", "'type'", "'var'", "", "", "'class'", "'process'", "'for'", 
  "'loop'", "'exit'", "'if'", "'else'", "'elsif'", "'case'", "'assert'", 
  "'begin'", "'return'", "'result'", "'new'", "'free'", "'tag'", "'fork'", 
  "'signal'", "'wait'", "'pause'", "'quit'", "'unchecked'", "'checked'", 
  "'export'", "'import'", "'inherit'", "'implement'", "'by'", "'put'", "'int'", 
  "'real'", "'boolean'", "'string'", "'array'", "'set'", "'record'", "'union'", 
  "'pointer'", "'nat'", "'intn'", "'natn'", "'realn'", "'char'", "'deferred'", 
  "'forward'", "'body'", "'not'", "'^'", "':'", "'('", "')'", "'.'", "'..'", 
  "','", "'#'", "'->'", "':='", "'+'", "'-'", "'*'", "'/'", "'div'", "'mod'", 
  "'rem'", "'**'", "'<'", "'>'", "'='", "'<='", "'>='", "'not='", "'and'", 
  "'or'", "'=>'", "'in'", "'shr'", "'shl'", "'xor'", "'+='", "'-='", "'*='", 
  "'/='", "'div='", "'mod='", "'and='", "'or='", "'=>='", "'shr='", "'shl='", 
  "'xor='"
};

std::vector<std::string> TuringParser::_symbolicNames = {
  "", "", "", "", "NOT_IN", "STRING_LITERAL", "INTEGER_LITERAL", "REAL_LITERAL", 
  "END", "OF", "TO", "TYPE", "VAR", "PROCEDURE", "FUNCTION", "CLASS", "PROCESS", 
  "FOR", "LOOP", "EXIT", "IF", "ELSE", "ELSIF", "CASE", "ASSERT", "BEGIN", 
  "RETURN", "RESULT", "NEW", "FREE", "TAG", "FORK", "SIGNAL", "WAIT", "PAUSE", 
  "QUIT", "UNCHECKED", "CHECKED", "EXPORT", "IMPORT", "INHERIT", "IMPLEMENT", 
  "BY", "PUT", "INT", "REAL", "BOOLEAN", "STRING", "ARRAY", "SET", "RECORD", 
  "UNION", "POINTER", "NAT", "INTN", "NATN", "REALN", "CHAR", "DEFERRED", 
  "FORWARD", "BODY", "NOT", "CARET", "COLON", "L_PAREN", "R_PAREN", "DOT", 
  "RANGE", "COMMA", "CHEAT", "DEREFERENCE", "ASSIGNMENT", "PLUS", "MINUS", 
  "MULTIPLY", "DIVIDE", "DIV", "MOD", "REM", "EXP", "LESSTHAN", "GREATERTHAN", 
  "EQUAL", "LESSTHANEQUAL", "GREATERTHANEQUAL", "NOTEQUAL", "AND", "OR", 
  "IMPLIES", "IN", "SHR", "SHL", "XOR", "PLUSEQUALS", "MINUSEQUALS", "MULTIPLYEQUALS", 
  "DIVIDEEQUALS", "DIVEQUALS", "MODEQUALS", "ANDEQUALS", "OREQUALS", "IMPLIESEQUALS", 
  "SHREQUALS", "SHLEQUALS", "XOREQUALS", "IDENTIFIER", "WHITESPACE", "BLOCK_COMMENT", 
  "LINE_COMMENT"
};

dfa::Vocabulary TuringParser::_vocabulary(_literalNames, _symbolicNames);

std::vector<std::string> TuringParser::_tokenNames;

TuringParser::Initializer::Initializer() {
	for (size_t i = 0; i < _symbolicNames.size(); ++i) {
		std::string name = _vocabulary.getLiteralName(i);
		if (name.empty()) {
			name = _vocabulary.getSymbolicName(i);
		}

		if (name.empty()) {
			_tokenNames.push_back("<INVALID>");
		} else {
      _tokenNames.push_back(name);
    }
	}

  _serializedATN = {
    0x3, 0x608b, 0xa72a, 0x8133, 0xb9ed, 0x417c, 0x3be7, 0x7786, 0x5964, 
    0x3, 0x6e, 0x1e0, 0x4, 0x2, 0x9, 0x2, 0x4, 0x3, 0x9, 0x3, 0x4, 0x4, 
    0x9, 0x4, 0x4, 0x5, 0x9, 0x5, 0x4, 0x6, 0x9, 0x6, 0x4, 0x7, 0x9, 0x7, 
    0x4, 0x8, 0x9, 0x8, 0x4, 0x9, 0x9, 0x9, 0x4, 0xa, 0x9, 0xa, 0x4, 0xb, 
    0x9, 0xb, 0x4, 0xc, 0x9, 0xc, 0x4, 0xd, 0x9, 0xd, 0x4, 0xe, 0x9, 0xe, 
    0x4, 0xf, 0x9, 0xf, 0x4, 0x10, 0x9, 0x10, 0x4, 0x11, 0x9, 0x11, 0x4, 
    0x12, 0x9, 0x12, 0x4, 0x13, 0x9, 0x13, 0x4, 0x14, 0x9, 0x14, 0x4, 0x15, 
    0x9, 0x15, 0x4, 0x16, 0x9, 0x16, 0x4, 0x17, 0x9, 0x17, 0x4, 0x18, 0x9, 
    0x18, 0x4, 0x19, 0x9, 0x19, 0x4, 0x1a, 0x9, 0x1a, 0x4, 0x1b, 0x9, 0x1b, 
    0x4, 0x1c, 0x9, 0x1c, 0x4, 0x1d, 0x9, 0x1d, 0x4, 0x1e, 0x9, 0x1e, 0x4, 
    0x1f, 0x9, 0x1f, 0x4, 0x20, 0x9, 0x20, 0x4, 0x21, 0x9, 0x21, 0x4, 0x22, 
    0x9, 0x22, 0x4, 0x23, 0x9, 0x23, 0x4, 0x24, 0x9, 0x24, 0x4, 0x25, 0x9, 
    0x25, 0x4, 0x26, 0x9, 0x26, 0x4, 0x27, 0x9, 0x27, 0x4, 0x28, 0x9, 0x28, 
    0x4, 0x29, 0x9, 0x29, 0x4, 0x2a, 0x9, 0x2a, 0x4, 0x2b, 0x9, 0x2b, 0x4, 
    0x2c, 0x9, 0x2c, 0x4, 0x2d, 0x9, 0x2d, 0x4, 0x2e, 0x9, 0x2e, 0x4, 0x2f, 
    0x9, 0x2f, 0x4, 0x30, 0x9, 0x30, 0x4, 0x31, 0x9, 0x31, 0x4, 0x32, 0x9, 
    0x32, 0x4, 0x33, 0x9, 0x33, 0x3, 0x2, 0x6, 0x2, 0x68, 0xa, 0x2, 0xd, 
    0x2, 0xe, 0x2, 0x69, 0x3, 0x3, 0x3, 0x3, 0x5, 0x3, 0x6e, 0xa, 0x3, 0x3, 
    0x4, 0x3, 0x4, 0x3, 0x4, 0x3, 0x4, 0x5, 0x4, 0x74, 0xa, 0x4, 0x3, 0x4, 
    0x5, 0x4, 0x77, 0xa, 0x4, 0x3, 0x5, 0x3, 0x5, 0x3, 0x5, 0x3, 0x5, 0x5, 
    0x5, 0x7d, 0xa, 0x5, 0x3, 0x5, 0x5, 0x5, 0x80, 0xa, 0x5, 0x3, 0x5, 0x5, 
    0x5, 0x83, 0xa, 0x5, 0x3, 0x5, 0x3, 0x5, 0x3, 0x5, 0x3, 0x6, 0x5, 0x6, 
    0x89, 0xa, 0x6, 0x3, 0x6, 0x3, 0x6, 0x5, 0x6, 0x8d, 0xa, 0x6, 0x3, 0x7, 
    0x3, 0x7, 0x3, 0x7, 0x3, 0x7, 0x3, 0x7, 0x3, 0x8, 0x3, 0x8, 0x3, 0x9, 
    0x3, 0x9, 0x3, 0x9, 0x6, 0x9, 0x99, 0xa, 0x9, 0xd, 0x9, 0xe, 0x9, 0x9a, 
    0x3, 0x9, 0x3, 0x9, 0x3, 0x9, 0x3, 0xa, 0x3, 0xa, 0x3, 0xa, 0x3, 0xa, 
    0x5, 0xa, 0xa4, 0xa, 0xa, 0x3, 0xb, 0x3, 0xb, 0x3, 0xb, 0x3, 0xb, 0x3, 
    0xb, 0x3, 0xb, 0x3, 0xb, 0x5, 0xb, 0xad, 0xa, 0xb, 0x3, 0xc, 0x5, 0xc, 
    0xb0, 0xa, 0xc, 0x3, 0xc, 0x3, 0xc, 0x3, 0xd, 0x3, 0xd, 0x3, 0xd, 0x3, 
    0xd, 0x3, 0xd, 0x3, 0xd, 0x7, 0xd, 0xba, 0xa, 0xd, 0xc, 0xd, 0xe, 0xd, 
    0xbd, 0xb, 0xd, 0x3, 0xe, 0x3, 0xe, 0x3, 0xe, 0x3, 0xe, 0x3, 0xe, 0x3, 
    0xe, 0x3, 0xe, 0x5, 0xe, 0xc6, 0xa, 0xe, 0x3, 0xf, 0x3, 0xf, 0x5, 0xf, 
    0xca, 0xa, 0xf, 0x3, 0xf, 0x3, 0xf, 0x3, 0xf, 0x5, 0xf, 0xcf, 0xa, 0xf, 
    0x3, 0xf, 0x3, 0xf, 0x3, 0xf, 0x3, 0xf, 0x5, 0xf, 0xd5, 0xa, 0xf, 0x3, 
    0x10, 0x3, 0x10, 0x3, 0x10, 0x3, 0x10, 0x5, 0x10, 0xdb, 0xa, 0x10, 0x3, 
    0x11, 0x3, 0x11, 0x3, 0x12, 0x5, 0x12, 0xe0, 0xa, 0x12, 0x3, 0x12, 0x3, 
    0x12, 0x3, 0x12, 0x3, 0x12, 0x3, 0x12, 0x5, 0x12, 0xe7, 0xa, 0x12, 0x3, 
    0x13, 0x3, 0x13, 0x3, 0x13, 0x3, 0x13, 0x3, 0x13, 0x3, 0x13, 0x7, 0x13, 
    0xef, 0xa, 0x13, 0xc, 0x13, 0xe, 0x13, 0xf2, 0xb, 0x13, 0x3, 0x14, 0x6, 
    0x14, 0xf5, 0xa, 0x14, 0xd, 0x14, 0xe, 0x14, 0xf6, 0x3, 0x15, 0x3, 0x15, 
    0x5, 0x15, 0xfb, 0xa, 0x15, 0x3, 0x16, 0x3, 0x16, 0x3, 0x16, 0x3, 0x16, 
    0x3, 0x16, 0x3, 0x16, 0x5, 0x16, 0x103, 0xa, 0x16, 0x3, 0x17, 0x3, 0x17, 
    0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x5, 
    0x17, 0x10d, 0xa, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 
    0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 
    0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 
    0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 
    0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 
    0x5, 0x17, 0x12d, 0xa, 0x17, 0x3, 0x17, 0x7, 0x17, 0x130, 0xa, 0x17, 
    0xc, 0x17, 0xe, 0x17, 0x133, 0xb, 0x17, 0x3, 0x18, 0x3, 0x18, 0x3, 0x18, 
    0x7, 0x18, 0x138, 0xa, 0x18, 0xc, 0x18, 0xe, 0x18, 0x13b, 0xb, 0x18, 
    0x3, 0x19, 0x3, 0x19, 0x3, 0x19, 0x3, 0x19, 0x3, 0x19, 0x5, 0x19, 0x142, 
    0xa, 0x19, 0x3, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x3, 
    0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x5, 0x1a, 0x14d, 0xa, 0x1a, 
    0x3, 0x1b, 0x3, 0x1b, 0x3, 0x1c, 0x3, 0x1c, 0x3, 0x1c, 0x3, 0x1d, 0x3, 
    0x1d, 0x3, 0x1e, 0x3, 0x1e, 0x3, 0x1e, 0x3, 0x1e, 0x3, 0x1e, 0x3, 0x1e, 
    0x7, 0x1e, 0x15c, 0xa, 0x1e, 0xc, 0x1e, 0xe, 0x1e, 0x15f, 0xb, 0x1e, 
    0x3, 0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x3, 0x20, 0x3, 0x20, 0x3, 
    0x20, 0x3, 0x21, 0x3, 0x21, 0x3, 0x21, 0x5, 0x21, 0x16b, 0xa, 0x21, 
    0x3, 0x21, 0x3, 0x21, 0x3, 0x22, 0x3, 0x22, 0x3, 0x22, 0x5, 0x22, 0x172, 
    0xa, 0x22, 0x3, 0x22, 0x3, 0x22, 0x3, 0x23, 0x3, 0x23, 0x3, 0x23, 0x3, 
    0x24, 0x3, 0x24, 0x3, 0x24, 0x3, 0x24, 0x3, 0x24, 0x3, 0x25, 0x3, 0x25, 
    0x3, 0x26, 0x3, 0x26, 0x3, 0x27, 0x3, 0x27, 0x3, 0x27, 0x3, 0x27, 0x3, 
    0x27, 0x3, 0x27, 0x3, 0x27, 0x5, 0x27, 0x189, 0xa, 0x27, 0x3, 0x28, 
    0x3, 0x28, 0x3, 0x28, 0x3, 0x28, 0x3, 0x29, 0x3, 0x29, 0x3, 0x29, 0x3, 
    0x29, 0x3, 0x29, 0x3, 0x29, 0x7, 0x29, 0x195, 0xa, 0x29, 0xc, 0x29, 
    0xe, 0x29, 0x198, 0xb, 0x29, 0x3, 0x2a, 0x3, 0x2a, 0x3, 0x2a, 0x3, 0x2a, 
    0x5, 0x2a, 0x19e, 0xa, 0x2a, 0x3, 0x2b, 0x3, 0x2b, 0x6, 0x2b, 0x1a2, 
    0xa, 0x2b, 0xd, 0x2b, 0xe, 0x2b, 0x1a3, 0x3, 0x2b, 0x3, 0x2b, 0x3, 0x2b, 
    0x3, 0x2c, 0x3, 0x2c, 0x3, 0x2c, 0x3, 0x2c, 0x3, 0x2d, 0x3, 0x2d, 0x3, 
    0x2d, 0x3, 0x2d, 0x3, 0x2d, 0x3, 0x2d, 0x5, 0x2d, 0x1b3, 0xa, 0x2d, 
    0x3, 0x2d, 0x3, 0x2d, 0x3, 0x2d, 0x3, 0x2d, 0x3, 0x2d, 0x5, 0x2d, 0x1ba, 
    0xa, 0x2d, 0x3, 0x2e, 0x3, 0x2e, 0x3, 0x2e, 0x7, 0x2e, 0x1bf, 0xa, 0x2e, 
    0xc, 0x2e, 0xe, 0x2e, 0x1c2, 0xb, 0x2e, 0x3, 0x2f, 0x3, 0x2f, 0x3, 0x30, 
    0x3, 0x30, 0x3, 0x30, 0x3, 0x30, 0x3, 0x30, 0x3, 0x30, 0x3, 0x30, 0x3, 
    0x30, 0x3, 0x30, 0x5, 0x30, 0x1cf, 0xa, 0x30, 0x3, 0x31, 0x3, 0x31, 
    0x3, 0x31, 0x3, 0x31, 0x3, 0x31, 0x3, 0x31, 0x7, 0x31, 0x1d7, 0xa, 0x31, 
    0xc, 0x31, 0xe, 0x31, 0x1da, 0xb, 0x31, 0x3, 0x32, 0x3, 0x32, 0x3, 0x33, 
    0x3, 0x33, 0x3, 0x33, 0x2, 0x8, 0x18, 0x24, 0x2c, 0x3a, 0x50, 0x60, 
    0x34, 0x2, 0x4, 0x6, 0x8, 0xa, 0xc, 0xe, 0x10, 0x12, 0x14, 0x16, 0x18, 
    0x1a, 0x1c, 0x1e, 0x20, 0x22, 0x24, 0x26, 0x28, 0x2a, 0x2c, 0x2e, 0x30, 
    0x32, 0x34, 0x36, 0x38, 0x3a, 0x3c, 0x3e, 0x40, 0x42, 0x44, 0x46, 0x48, 
    0x4a, 0x4c, 0x4e, 0x50, 0x52, 0x54, 0x56, 0x58, 0x5a, 0x5c, 0x5e, 0x60, 
    0x62, 0x64, 0x2, 0xe, 0x4, 0x2, 0xf, 0xf, 0x12, 0x12, 0x3, 0x2, 0x3c, 
    0x3e, 0x4, 0x2, 0x3, 0x5, 0xe, 0xe, 0x4, 0x2, 0x47, 0x47, 0x4a, 0x4b, 
    0x4, 0x2, 0x4c, 0x50, 0x5c, 0x5d, 0x4, 0x2, 0x4a, 0x4b, 0x5e, 0x5e, 
    0x5, 0x2, 0x6, 0x6, 0x52, 0x57, 0x5b, 0x5b, 0x5, 0x2, 0x49, 0x49, 0x5f, 
    0x63, 0x68, 0x69, 0x4, 0x2, 0x44, 0x44, 0x48, 0x48, 0x4, 0x2, 0x2e, 
    0x30, 0x37, 0x3b, 0x3, 0x2, 0x7, 0x9, 0x3, 0x2, 0x6d, 0x6e, 0x2, 0x1f5, 
    0x2, 0x67, 0x3, 0x2, 0x2, 0x2, 0x4, 0x6d, 0x3, 0x2, 0x2, 0x2, 0x6, 0x6f, 
    0x3, 0x2, 0x2, 0x2, 0x8, 0x78, 0x3, 0x2, 0x2, 0x2, 0xa, 0x88, 0x3, 0x2, 
    0x2, 0x2, 0xc, 0x8e, 0x3, 0x2, 0x2, 0x2, 0xe, 0x93, 0x3, 0x2, 0x2, 0x2, 
    0x10, 0x95, 0x3, 0x2, 0x2, 0x2, 0x12, 0xa3, 0x3, 0x2, 0x2, 0x2, 0x14, 
    0xac, 0x3, 0x2, 0x2, 0x2, 0x16, 0xaf, 0x3, 0x2, 0x2, 0x2, 0x18, 0xb3, 
    0x3, 0x2, 0x2, 0x2, 0x1a, 0xc5, 0x3, 0x2, 0x2, 0x2, 0x1c, 0xd4, 0x3, 
    0x2, 0x2, 0x2, 0x1e, 0xda, 0x3, 0x2, 0x2, 0x2, 0x20, 0xdc, 0x3, 0x2, 
    0x2, 0x2, 0x22, 0xe6, 0x3, 0x2, 0x2, 0x2, 0x24, 0xe8, 0x3, 0x2, 0x2, 
    0x2, 0x26, 0xf4, 0x3, 0x2, 0x2, 0x2, 0x28, 0xfa, 0x3, 0x2, 0x2, 0x2, 
    0x2a, 0x102, 0x3, 0x2, 0x2, 0x2, 0x2c, 0x10c, 0x3, 0x2, 0x2, 0x2, 0x2e, 
    0x134, 0x3, 0x2, 0x2, 0x2, 0x30, 0x141, 0x3, 0x2, 0x2, 0x2, 0x32, 0x14c, 
    0x3, 0x2, 0x2, 0x2, 0x34, 0x14e, 0x3, 0x2, 0x2, 0x2, 0x36, 0x150, 0x3, 
    0x2, 0x2, 0x2, 0x38, 0x153, 0x3, 0x2, 0x2, 0x2, 0x3a, 0x155, 0x3, 0x2, 
    0x2, 0x2, 0x3c, 0x160, 0x3, 0x2, 0x2, 0x2, 0x3e, 0x164, 0x3, 0x2, 0x2, 
    0x2, 0x40, 0x167, 0x3, 0x2, 0x2, 0x2, 0x42, 0x16e, 0x3, 0x2, 0x2, 0x2, 
    0x44, 0x175, 0x3, 0x2, 0x2, 0x2, 0x46, 0x178, 0x3, 0x2, 0x2, 0x2, 0x48, 
    0x17d, 0x3, 0x2, 0x2, 0x2, 0x4a, 0x17f, 0x3, 0x2, 0x2, 0x2, 0x4c, 0x188, 
    0x3, 0x2, 0x2, 0x2, 0x4e, 0x18a, 0x3, 0x2, 0x2, 0x2, 0x50, 0x18e, 0x3, 
    0x2, 0x2, 0x2, 0x52, 0x199, 0x3, 0x2, 0x2, 0x2, 0x54, 0x19f, 0x3, 0x2, 
    0x2, 0x2, 0x56, 0x1a8, 0x3, 0x2, 0x2, 0x2, 0x58, 0x1b9, 0x3, 0x2, 0x2, 
    0x2, 0x5a, 0x1bb, 0x3, 0x2, 0x2, 0x2, 0x5c, 0x1c3, 0x3, 0x2, 0x2, 0x2, 
    0x5e, 0x1ce, 0x3, 0x2, 0x2, 0x2, 0x60, 0x1d0, 0x3, 0x2, 0x2, 0x2, 0x62, 
    0x1db, 0x3, 0x2, 0x2, 0x2, 0x64, 0x1dd, 0x3, 0x2, 0x2, 0x2, 0x66, 0x68, 
    0x5, 0x4, 0x3, 0x2, 0x67, 0x66, 0x3, 0x2, 0x2, 0x2, 0x68, 0x69, 0x3, 
    0x2, 0x2, 0x2, 0x69, 0x67, 0x3, 0x2, 0x2, 0x2, 0x69, 0x6a, 0x3, 0x2, 
    0x2, 0x2, 0x6a, 0x3, 0x3, 0x2, 0x2, 0x2, 0x6b, 0x6e, 0x5, 0x28, 0x15, 
    0x2, 0x6c, 0x6e, 0x5, 0x10, 0x9, 0x2, 0x6d, 0x6b, 0x3, 0x2, 0x2, 0x2, 
    0x6d, 0x6c, 0x3, 0x2, 0x2, 0x2, 0x6e, 0x5, 0x3, 0x2, 0x2, 0x2, 0x6f, 
    0x70, 0x9, 0x2, 0x2, 0x2, 0x70, 0x76, 0x7, 0x6b, 0x2, 0x2, 0x71, 0x73, 
    0x7, 0x42, 0x2, 0x2, 0x72, 0x74, 0x5, 0x24, 0x13, 0x2, 0x73, 0x72, 0x3, 
    0x2, 0x2, 0x2, 0x73, 0x74, 0x3, 0x2, 0x2, 0x2, 0x74, 0x75, 0x3, 0x2, 
    0x2, 0x2, 0x75, 0x77, 0x7, 0x43, 0x2, 0x2, 0x76, 0x71, 0x3, 0x2, 0x2, 
    0x2, 0x76, 0x77, 0x3, 0x2, 0x2, 0x2, 0x77, 0x7, 0x3, 0x2, 0x2, 0x2, 
    0x78, 0x79, 0x7, 0x10, 0x2, 0x2, 0x79, 0x7f, 0x7, 0x6b, 0x2, 0x2, 0x7a, 
    0x7c, 0x7, 0x42, 0x2, 0x2, 0x7b, 0x7d, 0x5, 0x24, 0x13, 0x2, 0x7c, 0x7b, 
    0x3, 0x2, 0x2, 0x2, 0x7c, 0x7d, 0x3, 0x2, 0x2, 0x2, 0x7d, 0x7e, 0x3, 
    0x2, 0x2, 0x2, 0x7e, 0x80, 0x7, 0x43, 0x2, 0x2, 0x7f, 0x7a, 0x3, 0x2, 
    0x2, 0x2, 0x7f, 0x80, 0x3, 0x2, 0x2, 0x2, 0x80, 0x82, 0x3, 0x2, 0x2, 
    0x2, 0x81, 0x83, 0x7, 0x6b, 0x2, 0x2, 0x82, 0x81, 0x3, 0x2, 0x2, 0x2, 
    0x82, 0x83, 0x3, 0x2, 0x2, 0x2, 0x83, 0x84, 0x3, 0x2, 0x2, 0x2, 0x84, 
    0x85, 0x7, 0x41, 0x2, 0x2, 0x85, 0x86, 0x5, 0x4c, 0x27, 0x2, 0x86, 0x9, 
    0x3, 0x2, 0x2, 0x2, 0x87, 0x89, 0x5, 0xe, 0x8, 0x2, 0x88, 0x87, 0x3, 
    0x2, 0x2, 0x2, 0x88, 0x89, 0x3, 0x2, 0x2, 0x2, 0x89, 0x8c, 0x3, 0x2, 
    0x2, 0x2, 0x8a, 0x8d, 0x5, 0x8, 0x5, 0x2, 0x8b, 0x8d, 0x5, 0x6, 0x4, 
    0x2, 0x8c, 0x8a, 0x3, 0x2, 0x2, 0x2, 0x8c, 0x8b, 0x3, 0x2, 0x2, 0x2, 
    0x8d, 0xb, 0x3, 0x2, 0x2, 0x2, 0x8e, 0x8f, 0x5, 0xa, 0x6, 0x2, 0x8f, 
    0x90, 0x5, 0x26, 0x14, 0x2, 0x90, 0x91, 0x7, 0xa, 0x2, 0x2, 0x91, 0x92, 
    0x7, 0x6b, 0x2, 0x2, 0x92, 0xd, 0x3, 0x2, 0x2, 0x2, 0x93, 0x94, 0x9, 
    0x3, 0x2, 0x2, 0x94, 0xf, 0x3, 0x2, 0x2, 0x2, 0x95, 0x96, 0x7, 0x11, 
    0x2, 0x2, 0x96, 0x98, 0x7, 0x6b, 0x2, 0x2, 0x97, 0x99, 0x5, 0x12, 0xa, 
    0x2, 0x98, 0x97, 0x3, 0x2, 0x2, 0x2, 0x99, 0x9a, 0x3, 0x2, 0x2, 0x2, 
    0x9a, 0x98, 0x3, 0x2, 0x2, 0x2, 0x9a, 0x9b, 0x3, 0x2, 0x2, 0x2, 0x9b, 
    0x9c, 0x3, 0x2, 0x2, 0x2, 0x9c, 0x9d, 0x7, 0xa, 0x2, 0x2, 0x9d, 0x9e, 
    0x7, 0x6b, 0x2, 0x2, 0x9e, 0x11, 0x3, 0x2, 0x2, 0x2, 0x9f, 0xa4, 0x5, 
    0x14, 0xb, 0x2, 0xa0, 0xa4, 0x5, 0x1a, 0xe, 0x2, 0xa1, 0xa4, 0x5, 0x1c, 
    0xf, 0x2, 0xa2, 0xa4, 0x5, 0x28, 0x15, 0x2, 0xa3, 0x9f, 0x3, 0x2, 0x2, 
    0x2, 0xa3, 0xa0, 0x3, 0x2, 0x2, 0x2, 0xa3, 0xa1, 0x3, 0x2, 0x2, 0x2, 
    0xa3, 0xa2, 0x3, 0x2, 0x2, 0x2, 0xa4, 0x13, 0x3, 0x2, 0x2, 0x2, 0xa5, 
    0xa6, 0x7, 0x28, 0x2, 0x2, 0xa6, 0xad, 0x5, 0x18, 0xd, 0x2, 0xa7, 0xa8, 
    0x7, 0x28, 0x2, 0x2, 0xa8, 0xa9, 0x7, 0x42, 0x2, 0x2, 0xa9, 0xaa, 0x5, 
    0x18, 0xd, 0x2, 0xaa, 0xab, 0x7, 0x43, 0x2, 0x2, 0xab, 0xad, 0x3, 0x2, 
    0x2, 0x2, 0xac, 0xa5, 0x3, 0x2, 0x2, 0x2, 0xac, 0xa7, 0x3, 0x2, 0x2, 
    0x2, 0xad, 0x15, 0x3, 0x2, 0x2, 0x2, 0xae, 0xb0, 0x5, 0x20, 0x11, 0x2, 
    0xaf, 0xae, 0x3, 0x2, 0x2, 0x2, 0xaf, 0xb0, 0x3, 0x2, 0x2, 0x2, 0xb0, 
    0xb1, 0x3, 0x2, 0x2, 0x2, 0xb1, 0xb2, 0x7, 0x6b, 0x2, 0x2, 0xb2, 0x17, 
    0x3, 0x2, 0x2, 0x2, 0xb3, 0xb4, 0x8, 0xd, 0x1, 0x2, 0xb4, 0xb5, 0x5, 
    0x16, 0xc, 0x2, 0xb5, 0xbb, 0x3, 0x2, 0x2, 0x2, 0xb6, 0xb7, 0xc, 0x3, 
    0x2, 0x2, 0xb7, 0xb8, 0x7, 0x46, 0x2, 0x2, 0xb8, 0xba, 0x5, 0x16, 0xc, 
    0x2, 0xb9, 0xb6, 0x3, 0x2, 0x2, 0x2, 0xba, 0xbd, 0x3, 0x2, 0x2, 0x2, 
    0xbb, 0xb9, 0x3, 0x2, 0x2, 0x2, 0xbb, 0xbc, 0x3, 0x2, 0x2, 0x2, 0xbc, 
    0x19, 0x3, 0x2, 0x2, 0x2, 0xbd, 0xbb, 0x3, 0x2, 0x2, 0x2, 0xbe, 0xbf, 
    0x7, 0x2a, 0x2, 0x2, 0xbf, 0xc6, 0x5, 0x1e, 0x10, 0x2, 0xc0, 0xc1, 0x7, 
    0x2a, 0x2, 0x2, 0xc1, 0xc2, 0x7, 0x42, 0x2, 0x2, 0xc2, 0xc3, 0x5, 0x1e, 
    0x10, 0x2, 0xc3, 0xc4, 0x7, 0x43, 0x2, 0x2, 0xc4, 0xc6, 0x3, 0x2, 0x2, 
    0x2, 0xc5, 0xbe, 0x3, 0x2, 0x2, 0x2, 0xc5, 0xc0, 0x3, 0x2, 0x2, 0x2, 
    0xc6, 0x1b, 0x3, 0x2, 0x2, 0x2, 0xc7, 0xc9, 0x7, 0x2b, 0x2, 0x2, 0xc8, 
    0xca, 0x7, 0x2c, 0x2, 0x2, 0xc9, 0xc8, 0x3, 0x2, 0x2, 0x2, 0xc9, 0xca, 
    0x3, 0x2, 0x2, 0x2, 0xca, 0xcb, 0x3, 0x2, 0x2, 0x2, 0xcb, 0xd5, 0x5, 
    0x1e, 0x10, 0x2, 0xcc, 0xce, 0x7, 0x2b, 0x2, 0x2, 0xcd, 0xcf, 0x7, 0x2c, 
    0x2, 0x2, 0xce, 0xcd, 0x3, 0x2, 0x2, 0x2, 0xce, 0xcf, 0x3, 0x2, 0x2, 
    0x2, 0xcf, 0xd0, 0x3, 0x2, 0x2, 0x2, 0xd0, 0xd1, 0x7, 0x42, 0x2, 0x2, 
    0xd1, 0xd2, 0x5, 0x1e, 0x10, 0x2, 0xd2, 0xd3, 0x7, 0x43, 0x2, 0x2, 0xd3, 
    0xd5, 0x3, 0x2, 0x2, 0x2, 0xd4, 0xc7, 0x3, 0x2, 0x2, 0x2, 0xd4, 0xcc, 
    0x3, 0x2, 0x2, 0x2, 0xd5, 0x1d, 0x3, 0x2, 0x2, 0x2, 0xd6, 0xdb, 0x7, 
    0x6b, 0x2, 0x2, 0xd7, 0xd8, 0x7, 0x6b, 0x2, 0x2, 0xd8, 0xd9, 0x7, 0x5b, 
    0x2, 0x2, 0xd9, 0xdb, 0x7, 0x7, 0x2, 0x2, 0xda, 0xd6, 0x3, 0x2, 0x2, 
    0x2, 0xda, 0xd7, 0x3, 0x2, 0x2, 0x2, 0xdb, 0x1f, 0x3, 0x2, 0x2, 0x2, 
    0xdc, 0xdd, 0x9, 0x4, 0x2, 0x2, 0xdd, 0x21, 0x3, 0x2, 0x2, 0x2, 0xde, 
    0xe0, 0x7, 0xe, 0x2, 0x2, 0xdf, 0xde, 0x3, 0x2, 0x2, 0x2, 0xdf, 0xe0, 
    0x3, 0x2, 0x2, 0x2, 0xe0, 0xe1, 0x3, 0x2, 0x2, 0x2, 0xe1, 0xe2, 0x5, 
    0x60, 0x31, 0x2, 0xe2, 0xe3, 0x7, 0x41, 0x2, 0x2, 0xe3, 0xe4, 0x5, 0x4c, 
    0x27, 0x2, 0xe4, 0xe7, 0x3, 0x2, 0x2, 0x2, 0xe5, 0xe7, 0x5, 0xa, 0x6, 
    0x2, 0xe6, 0xdf, 0x3, 0x2, 0x2, 0x2, 0xe6, 0xe5, 0x3, 0x2, 0x2, 0x2, 
    0xe7, 0x23, 0x3, 0x2, 0x2, 0x2, 0xe8, 0xe9, 0x8, 0x13, 0x1, 0x2, 0xe9, 
    0xea, 0x5, 0x22, 0x12, 0x2, 0xea, 0xf0, 0x3, 0x2, 0x2, 0x2, 0xeb, 0xec, 
    0xc, 0x3, 0x2, 0x2, 0xec, 0xed, 0x7, 0x46, 0x2, 0x2, 0xed, 0xef, 0x5, 
    0x22, 0x12, 0x2, 0xee, 0xeb, 0x3, 0x2, 0x2, 0x2, 0xef, 0xf2, 0x3, 0x2, 
    0x2, 0x2, 0xf0, 0xee, 0x3, 0x2, 0x2, 0x2, 0xf0, 0xf1, 0x3, 0x2, 0x2, 
    0x2, 0xf1, 0x25, 0x3, 0x2, 0x2, 0x2, 0xf2, 0xf0, 0x3, 0x2, 0x2, 0x2, 
    0xf3, 0xf5, 0x5, 0x28, 0x15, 0x2, 0xf4, 0xf3, 0x3, 0x2, 0x2, 0x2, 0xf5, 
    0xf6, 0x3, 0x2, 0x2, 0x2, 0xf6, 0xf4, 0x3, 0x2, 0x2, 0x2, 0xf6, 0xf7, 
    0x3, 0x2, 0x2, 0x2, 0xf7, 0x27, 0x3, 0x2, 0x2, 0x2, 0xf8, 0xfb, 0x5, 
    0x32, 0x1a, 0x2, 0xf9, 0xfb, 0x5, 0x30, 0x19, 0x2, 0xfa, 0xf8, 0x3, 
    0x2, 0x2, 0x2, 0xfa, 0xf9, 0x3, 0x2, 0x2, 0x2, 0xfb, 0x29, 0x3, 0x2, 
    0x2, 0x2, 0xfc, 0x103, 0x5, 0x62, 0x32, 0x2, 0xfd, 0x103, 0x7, 0x6b, 
    0x2, 0x2, 0xfe, 0xff, 0x7, 0x42, 0x2, 0x2, 0xff, 0x100, 0x5, 0x2c, 0x17, 
    0x2, 0x100, 0x101, 0x7, 0x43, 0x2, 0x2, 0x101, 0x103, 0x3, 0x2, 0x2, 
    0x2, 0x102, 0xfc, 0x3, 0x2, 0x2, 0x2, 0x102, 0xfd, 0x3, 0x2, 0x2, 0x2, 
    0x102, 0xfe, 0x3, 0x2, 0x2, 0x2, 0x103, 0x2b, 0x3, 0x2, 0x2, 0x2, 0x104, 
    0x105, 0x8, 0x17, 0x1, 0x2, 0x105, 0x10d, 0x5, 0x2a, 0x16, 0x2, 0x106, 
    0x107, 0x7, 0x40, 0x2, 0x2, 0x107, 0x10d, 0x5, 0x2c, 0x17, 0xf, 0x108, 
    0x109, 0x9, 0x5, 0x2, 0x2, 0x109, 0x10d, 0x5, 0x2c, 0x17, 0xb, 0x10a, 
    0x10b, 0x7, 0x3f, 0x2, 0x2, 0x10b, 0x10d, 0x5, 0x2c, 0x17, 0x7, 0x10c, 
    0x104, 0x3, 0x2, 0x2, 0x2, 0x10c, 0x106, 0x3, 0x2, 0x2, 0x2, 0x10c, 
    0x108, 0x3, 0x2, 0x2, 0x2, 0x10c, 0x10a, 0x3, 0x2, 0x2, 0x2, 0x10d, 
    0x131, 0x3, 0x2, 0x2, 0x2, 0x10e, 0x10f, 0xc, 0xd, 0x2, 0x2, 0x10f, 
    0x110, 0x7, 0x51, 0x2, 0x2, 0x110, 0x130, 0x5, 0x2c, 0x17, 0xe, 0x111, 
    0x112, 0xc, 0xa, 0x2, 0x2, 0x112, 0x113, 0x9, 0x6, 0x2, 0x2, 0x113, 
    0x130, 0x5, 0x2c, 0x17, 0xb, 0x114, 0x115, 0xc, 0x9, 0x2, 0x2, 0x115, 
    0x116, 0x9, 0x7, 0x2, 0x2, 0x116, 0x130, 0x5, 0x2c, 0x17, 0xa, 0x117, 
    0x118, 0xc, 0x8, 0x2, 0x2, 0x118, 0x119, 0x9, 0x8, 0x2, 0x2, 0x119, 
    0x130, 0x5, 0x2c, 0x17, 0x9, 0x11a, 0x11b, 0xc, 0x6, 0x2, 0x2, 0x11b, 
    0x11c, 0x7, 0x58, 0x2, 0x2, 0x11c, 0x130, 0x5, 0x2c, 0x17, 0x7, 0x11d, 
    0x11e, 0xc, 0x5, 0x2, 0x2, 0x11e, 0x11f, 0x7, 0x59, 0x2, 0x2, 0x11f, 
    0x130, 0x5, 0x2c, 0x17, 0x6, 0x120, 0x121, 0xc, 0x4, 0x2, 0x2, 0x121, 
    0x122, 0x7, 0x5a, 0x2, 0x2, 0x122, 0x130, 0x5, 0x2c, 0x17, 0x5, 0x123, 
    0x124, 0xc, 0x3, 0x2, 0x2, 0x124, 0x125, 0x9, 0x9, 0x2, 0x2, 0x125, 
    0x130, 0x5, 0x2c, 0x17, 0x3, 0x126, 0x127, 0xc, 0xe, 0x2, 0x2, 0x127, 
    0x128, 0x9, 0xa, 0x2, 0x2, 0x128, 0x130, 0x7, 0x6b, 0x2, 0x2, 0x129, 
    0x12a, 0xc, 0xc, 0x2, 0x2, 0x12a, 0x12c, 0x7, 0x42, 0x2, 0x2, 0x12b, 
    0x12d, 0x5, 0x2e, 0x18, 0x2, 0x12c, 0x12b, 0x3, 0x2, 0x2, 0x2, 0x12c, 
    0x12d, 0x3, 0x2, 0x2, 0x2, 0x12d, 0x12e, 0x3, 0x2, 0x2, 0x2, 0x12e, 
    0x130, 0x7, 0x43, 0x2, 0x2, 0x12f, 0x10e, 0x3, 0x2, 0x2, 0x2, 0x12f, 
    0x111, 0x3, 0x2, 0x2, 0x2, 0x12f, 0x114, 0x3, 0x2, 0x2, 0x2, 0x12f, 
    0x117, 0x3, 0x2, 0x2, 0x2, 0x12f, 0x11a, 0x3, 0x2, 0x2, 0x2, 0x12f, 
    0x11d, 0x3, 0x2, 0x2, 0x2, 0x12f, 0x120, 0x3, 0x2, 0x2, 0x2, 0x12f, 
    0x123, 0x3, 0x2, 0x2, 0x2, 0x12f, 0x126, 0x3, 0x2, 0x2, 0x2, 0x12f, 
    0x129, 0x3, 0x2, 0x2, 0x2, 0x130, 0x133, 0x3, 0x2, 0x2, 0x2, 0x131, 
    0x12f, 0x3, 0x2, 0x2, 0x2, 0x131, 0x132, 0x3, 0x2, 0x2, 0x2, 0x132, 
    0x2d, 0x3, 0x2, 0x2, 0x2, 0x133, 0x131, 0x3, 0x2, 0x2, 0x2, 0x134, 0x139, 
    0x5, 0x2c, 0x17, 0x2, 0x135, 0x136, 0x7, 0x46, 0x2, 0x2, 0x136, 0x138, 
    0x5, 0x2c, 0x17, 0x2, 0x137, 0x135, 0x3, 0x2, 0x2, 0x2, 0x138, 0x13b, 
    0x3, 0x2, 0x2, 0x2, 0x139, 0x137, 0x3, 0x2, 0x2, 0x2, 0x139, 0x13a, 
    0x3, 0x2, 0x2, 0x2, 0x13a, 0x2f, 0x3, 0x2, 0x2, 0x2, 0x13b, 0x139, 0x3, 
    0x2, 0x2, 0x2, 0x13c, 0x142, 0x5, 0x46, 0x24, 0x2, 0x13d, 0x142, 0x5, 
    0x58, 0x2d, 0x2, 0x13e, 0x142, 0x5, 0x5e, 0x30, 0x2, 0x13f, 0x142, 0x5, 
    0xc, 0x7, 0x2, 0x140, 0x142, 0x5, 0x10, 0x9, 0x2, 0x141, 0x13c, 0x3, 
    0x2, 0x2, 0x2, 0x141, 0x13d, 0x3, 0x2, 0x2, 0x2, 0x141, 0x13e, 0x3, 
    0x2, 0x2, 0x2, 0x141, 0x13f, 0x3, 0x2, 0x2, 0x2, 0x141, 0x140, 0x3, 
    0x2, 0x2, 0x2, 0x142, 0x31, 0x3, 0x2, 0x2, 0x2, 0x143, 0x14d, 0x5, 0x2c, 
    0x17, 0x2, 0x144, 0x14d, 0x5, 0x36, 0x1c, 0x2, 0x145, 0x14d, 0x7, 0x15, 
    0x2, 0x2, 0x146, 0x14d, 0x5, 0x3c, 0x1f, 0x2, 0x147, 0x14d, 0x7, 0x1c, 
    0x2, 0x2, 0x148, 0x14d, 0x5, 0x3e, 0x20, 0x2, 0x149, 0x14d, 0x5, 0x40, 
    0x21, 0x2, 0x14a, 0x14d, 0x5, 0x42, 0x22, 0x2, 0x14b, 0x14d, 0x5, 0x44, 
    0x23, 0x2, 0x14c, 0x143, 0x3, 0x2, 0x2, 0x2, 0x14c, 0x144, 0x3, 0x2, 
    0x2, 0x2, 0x14c, 0x145, 0x3, 0x2, 0x2, 0x2, 0x14c, 0x146, 0x3, 0x2, 
    0x2, 0x2, 0x14c, 0x147, 0x3, 0x2, 0x2, 0x2, 0x14c, 0x148, 0x3, 0x2, 
    0x2, 0x2, 0x14c, 0x149, 0x3, 0x2, 0x2, 0x2, 0x14c, 0x14a, 0x3, 0x2, 
    0x2, 0x2, 0x14c, 0x14b, 0x3, 0x2, 0x2, 0x2, 0x14d, 0x33, 0x3, 0x2, 0x2, 
    0x2, 0x14e, 0x14f, 0x9, 0x9, 0x2, 0x2, 0x14f, 0x35, 0x3, 0x2, 0x2, 0x2, 
    0x150, 0x151, 0x7, 0x2d, 0x2, 0x2, 0x151, 0x152, 0x5, 0x3a, 0x1e, 0x2, 
    0x152, 0x37, 0x3, 0x2, 0x2, 0x2, 0x153, 0x154, 0x5, 0x32, 0x1a, 0x2, 
    0x154, 0x39, 0x3, 0x2, 0x2, 0x2, 0x155, 0x156, 0x8, 0x1e, 0x1, 0x2, 
    0x156, 0x157, 0x5, 0x38, 0x1d, 0x2, 0x157, 0x15d, 0x3, 0x2, 0x2, 0x2, 
    0x158, 0x159, 0xc, 0x3, 0x2, 0x2, 0x159, 0x15a, 0x7, 0x46, 0x2, 0x2, 
    0x15a, 0x15c, 0x5, 0x38, 0x1d, 0x2, 0x15b, 0x158, 0x3, 0x2, 0x2, 0x2, 
    0x15c, 0x15f, 0x3, 0x2, 0x2, 0x2, 0x15d, 0x15b, 0x3, 0x2, 0x2, 0x2, 
    0x15d, 0x15e, 0x3, 0x2, 0x2, 0x2, 0x15e, 0x3b, 0x3, 0x2, 0x2, 0x2, 0x15f, 
    0x15d, 0x3, 0x2, 0x2, 0x2, 0x160, 0x161, 0x7, 0x1b, 0x2, 0x2, 0x161, 
    0x162, 0x5, 0x28, 0x15, 0x2, 0x162, 0x163, 0x7, 0xa, 0x2, 0x2, 0x163, 
    0x3d, 0x3, 0x2, 0x2, 0x2, 0x164, 0x165, 0x7, 0x1d, 0x2, 0x2, 0x165, 
    0x166, 0x5, 0x2c, 0x17, 0x2, 0x166, 0x3f, 0x3, 0x2, 0x2, 0x2, 0x167, 
    0x16a, 0x7, 0x1e, 0x2, 0x2, 0x168, 0x169, 0x7, 0x6b, 0x2, 0x2, 0x169, 
    0x16b, 0x7, 0x46, 0x2, 0x2, 0x16a, 0x168, 0x3, 0x2, 0x2, 0x2, 0x16a, 
    0x16b, 0x3, 0x2, 0x2, 0x2, 0x16b, 0x16c, 0x3, 0x2, 0x2, 0x2, 0x16c, 
    0x16d, 0x7, 0x6b, 0x2, 0x2, 0x16d, 0x41, 0x3, 0x2, 0x2, 0x2, 0x16e, 
    0x171, 0x7, 0x1f, 0x2, 0x2, 0x16f, 0x170, 0x7, 0x6b, 0x2, 0x2, 0x170, 
    0x172, 0x7, 0x46, 0x2, 0x2, 0x171, 0x16f, 0x3, 0x2, 0x2, 0x2, 0x171, 
    0x172, 0x3, 0x2, 0x2, 0x2, 0x172, 0x173, 0x3, 0x2, 0x2, 0x2, 0x173, 
    0x174, 0x7, 0x6b, 0x2, 0x2, 0x174, 0x43, 0x3, 0x2, 0x2, 0x2, 0x175, 
    0x176, 0x7, 0x21, 0x2, 0x2, 0x176, 0x177, 0x5, 0x2c, 0x17, 0x2, 0x177, 
    0x45, 0x3, 0x2, 0x2, 0x2, 0x178, 0x179, 0x7, 0xd, 0x2, 0x2, 0x179, 0x17a, 
    0x7, 0x6b, 0x2, 0x2, 0x17a, 0x17b, 0x7, 0x41, 0x2, 0x2, 0x17b, 0x17c, 
    0x5, 0x4c, 0x27, 0x2, 0x17c, 0x47, 0x3, 0x2, 0x2, 0x2, 0x17d, 0x17e, 
    0x9, 0xb, 0x2, 0x2, 0x17e, 0x49, 0x3, 0x2, 0x2, 0x2, 0x17f, 0x180, 0x7, 
    0x6b, 0x2, 0x2, 0x180, 0x4b, 0x3, 0x2, 0x2, 0x2, 0x181, 0x189, 0x5, 
    0x48, 0x25, 0x2, 0x182, 0x189, 0x5, 0x4e, 0x28, 0x2, 0x183, 0x189, 0x5, 
    0x52, 0x2a, 0x2, 0x184, 0x189, 0x5, 0x54, 0x2b, 0x2, 0x185, 0x189, 0x5, 
    0x5e, 0x30, 0x2, 0x186, 0x189, 0x5, 0x10, 0x9, 0x2, 0x187, 0x189, 0x5, 
    0x4a, 0x26, 0x2, 0x188, 0x181, 0x3, 0x2, 0x2, 0x2, 0x188, 0x182, 0x3, 
    0x2, 0x2, 0x2, 0x188, 0x183, 0x3, 0x2, 0x2, 0x2, 0x188, 0x184, 0x3, 
    0x2, 0x2, 0x2, 0x188, 0x185, 0x3, 0x2, 0x2, 0x2, 0x188, 0x186, 0x3, 
    0x2, 0x2, 0x2, 0x188, 0x187, 0x3, 0x2, 0x2, 0x2, 0x189, 0x4d, 0x3, 0x2, 
    0x2, 0x2, 0x18a, 0x18b, 0x7, 0x8, 0x2, 0x2, 0x18b, 0x18c, 0x7, 0x45, 
    0x2, 0x2, 0x18c, 0x18d, 0x7, 0x8, 0x2, 0x2, 0x18d, 0x4f, 0x3, 0x2, 0x2, 
    0x2, 0x18e, 0x18f, 0x8, 0x29, 0x1, 0x2, 0x18f, 0x190, 0x5, 0x4e, 0x28, 
    0x2, 0x190, 0x196, 0x3, 0x2, 0x2, 0x2, 0x191, 0x192, 0xc, 0x3, 0x2, 
    0x2, 0x192, 0x193, 0x7, 0x46, 0x2, 0x2, 0x193, 0x195, 0x5, 0x4e, 0x28, 
    0x2, 0x194, 0x191, 0x3, 0x2, 0x2, 0x2, 0x195, 0x198, 0x3, 0x2, 0x2, 
    0x2, 0x196, 0x194, 0x3, 0x2, 0x2, 0x2, 0x196, 0x197, 0x3, 0x2, 0x2, 
    0x2, 0x197, 0x51, 0x3, 0x2, 0x2, 0x2, 0x198, 0x196, 0x3, 0x2, 0x2, 0x2, 
    0x199, 0x19d, 0x7, 0x31, 0x2, 0x2, 0x19a, 0x19b, 0x7, 0x42, 0x2, 0x2, 
    0x19b, 0x19c, 0x7, 0x8, 0x2, 0x2, 0x19c, 0x19e, 0x7, 0x43, 0x2, 0x2, 
    0x19d, 0x19a, 0x3, 0x2, 0x2, 0x2, 0x19d, 0x19e, 0x3, 0x2, 0x2, 0x2, 
    0x19e, 0x53, 0x3, 0x2, 0x2, 0x2, 0x19f, 0x1a1, 0x7, 0x34, 0x2, 0x2, 
    0x1a0, 0x1a2, 0x5, 0x56, 0x2c, 0x2, 0x1a1, 0x1a0, 0x3, 0x2, 0x2, 0x2, 
    0x1a2, 0x1a3, 0x3, 0x2, 0x2, 0x2, 0x1a3, 0x1a1, 0x3, 0x2, 0x2, 0x2, 
    0x1a3, 0x1a4, 0x3, 0x2, 0x2, 0x2, 0x1a4, 0x1a5, 0x3, 0x2, 0x2, 0x2, 
    0x1a5, 0x1a6, 0x7, 0xa, 0x2, 0x2, 0x1a6, 0x1a7, 0x7, 0x34, 0x2, 0x2, 
    0x1a7, 0x55, 0x3, 0x2, 0x2, 0x2, 0x1a8, 0x1a9, 0x5, 0x60, 0x31, 0x2, 
    0x1a9, 0x1aa, 0x7, 0x41, 0x2, 0x2, 0x1aa, 0x1ab, 0x5, 0x4c, 0x27, 0x2, 
    0x1ab, 0x57, 0x3, 0x2, 0x2, 0x2, 0x1ac, 0x1ad, 0x7, 0xe, 0x2, 0x2, 0x1ad, 
    0x1ae, 0x5, 0x5a, 0x2e, 0x2, 0x1ae, 0x1af, 0x7, 0x41, 0x2, 0x2, 0x1af, 
    0x1b2, 0x5, 0x4c, 0x27, 0x2, 0x1b0, 0x1b1, 0x7, 0x49, 0x2, 0x2, 0x1b1, 
    0x1b3, 0x5, 0x2c, 0x17, 0x2, 0x1b2, 0x1b0, 0x3, 0x2, 0x2, 0x2, 0x1b2, 
    0x1b3, 0x3, 0x2, 0x2, 0x2, 0x1b3, 0x1ba, 0x3, 0x2, 0x2, 0x2, 0x1b4, 
    0x1b5, 0x7, 0xe, 0x2, 0x2, 0x1b5, 0x1b6, 0x5, 0x5a, 0x2e, 0x2, 0x1b6, 
    0x1b7, 0x7, 0x49, 0x2, 0x2, 0x1b7, 0x1b8, 0x5, 0x2c, 0x17, 0x2, 0x1b8, 
    0x1ba, 0x3, 0x2, 0x2, 0x2, 0x1b9, 0x1ac, 0x3, 0x2, 0x2, 0x2, 0x1b9, 
    0x1b4, 0x3, 0x2, 0x2, 0x2, 0x1ba, 0x59, 0x3, 0x2, 0x2, 0x2, 0x1bb, 0x1c0, 
    0x5, 0x5c, 0x2f, 0x2, 0x1bc, 0x1bd, 0x7, 0x46, 0x2, 0x2, 0x1bd, 0x1bf, 
    0x5, 0x5c, 0x2f, 0x2, 0x1be, 0x1bc, 0x3, 0x2, 0x2, 0x2, 0x1bf, 0x1c2, 
    0x3, 0x2, 0x2, 0x2, 0x1c0, 0x1be, 0x3, 0x2, 0x2, 0x2, 0x1c0, 0x1c1, 
    0x3, 0x2, 0x2, 0x2, 0x1c1, 0x5b, 0x3, 0x2, 0x2, 0x2, 0x1c2, 0x1c0, 0x3, 
    0x2, 0x2, 0x2, 0x1c3, 0x1c4, 0x7, 0x6b, 0x2, 0x2, 0x1c4, 0x5d, 0x3, 
    0x2, 0x2, 0x2, 0x1c5, 0x1c6, 0x7, 0x32, 0x2, 0x2, 0x1c6, 0x1c7, 0x5, 
    0x4e, 0x28, 0x2, 0x1c7, 0x1c8, 0x7, 0xb, 0x2, 0x2, 0x1c8, 0x1c9, 0x5, 
    0x4c, 0x27, 0x2, 0x1c9, 0x1cf, 0x3, 0x2, 0x2, 0x2, 0x1ca, 0x1cb, 0x7, 
    0x32, 0x2, 0x2, 0x1cb, 0x1cc, 0x5, 0x4e, 0x28, 0x2, 0x1cc, 0x1cd, 0x7, 
    0x46, 0x2, 0x2, 0x1cd, 0x1cf, 0x3, 0x2, 0x2, 0x2, 0x1ce, 0x1c5, 0x3, 
    0x2, 0x2, 0x2, 0x1ce, 0x1ca, 0x3, 0x2, 0x2, 0x2, 0x1cf, 0x5f, 0x3, 0x2, 
    0x2, 0x2, 0x1d0, 0x1d1, 0x8, 0x31, 0x1, 0x2, 0x1d1, 0x1d2, 0x7, 0x6b, 
    0x2, 0x2, 0x1d2, 0x1d8, 0x3, 0x2, 0x2, 0x2, 0x1d3, 0x1d4, 0xc, 0x3, 
    0x2, 0x2, 0x1d4, 0x1d5, 0x7, 0x46, 0x2, 0x2, 0x1d5, 0x1d7, 0x7, 0x6b, 
    0x2, 0x2, 0x1d6, 0x1d3, 0x3, 0x2, 0x2, 0x2, 0x1d7, 0x1da, 0x3, 0x2, 
    0x2, 0x2, 0x1d8, 0x1d6, 0x3, 0x2, 0x2, 0x2, 0x1d8, 0x1d9, 0x3, 0x2, 
    0x2, 0x2, 0x1d9, 0x61, 0x3, 0x2, 0x2, 0x2, 0x1da, 0x1d8, 0x3, 0x2, 0x2, 
    0x2, 0x1db, 0x1dc, 0x9, 0xc, 0x2, 0x2, 0x1dc, 0x63, 0x3, 0x2, 0x2, 0x2, 
    0x1dd, 0x1de, 0x9, 0xd, 0x2, 0x2, 0x1de, 0x65, 0x3, 0x2, 0x2, 0x2, 0x2e, 
    0x69, 0x6d, 0x73, 0x76, 0x7c, 0x7f, 0x82, 0x88, 0x8c, 0x9a, 0xa3, 0xac, 
    0xaf, 0xbb, 0xc5, 0xc9, 0xce, 0xd4, 0xda, 0xdf, 0xe6, 0xf0, 0xf6, 0xfa, 
    0x102, 0x10c, 0x12c, 0x12f, 0x131, 0x139, 0x141, 0x14c, 0x15d, 0x16a, 
    0x171, 0x188, 0x196, 0x19d, 0x1a3, 0x1b2, 0x1b9, 0x1c0, 0x1ce, 0x1d8, 
  };

  atn::ATNDeserializer deserializer;
  _atn = deserializer.deserialize(_serializedATN);

  size_t count = _atn.getNumberOfDecisions();
  _decisionToDFA.reserve(count);
  for (size_t i = 0; i < count; i++) { 
    _decisionToDFA.emplace_back(_atn.getDecisionState(i), i);
  }
}

TuringParser::Initializer TuringParser::_init;
