
// Generated from Turing.g4 by ANTLR 4.7.2


#include "TuringListener.h"

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

TuringParser::ProgramContext* TuringParser::program() {
  ProgramContext *_localctx = _tracker.createInstance<ProgramContext>(_ctx, getState());
  enterRule(_localctx, 0, TuringParser::RuleProgram);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(123); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(122);
      topLevel();
      setState(125); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::STRING_LITERAL)
      | (1ULL << TuringParser::INTEGER_LITERAL)
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
      | (1ULL << TuringParser::CARET)
      | (1ULL << TuringParser::L_PAREN))) != 0) || ((((_la - 67) & ~ 0x3fULL) == 0) &&
      ((1ULL << (_la - 67)) & ((1ULL << (TuringParser::CHEAT - 67))
      | (1ULL << (TuringParser::PLUS - 67))
      | (1ULL << (TuringParser::MINUS - 67))
      | (1ULL << (TuringParser::IDENTIFIER - 67)))) != 0));
   
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

TuringParser::TopLevelContext* TuringParser::topLevel() {
  TopLevelContext *_localctx = _tracker.createInstance<TopLevelContext>(_ctx, getState());
  enterRule(_localctx, 2, TuringParser::RuleTopLevel);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(129);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL:
      case TuringParser::TYPE:
      case TuringParser::VAR:
      case TuringParser::PROCEDURE:
      case TuringParser::FUNCTION:
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
        enterOuterAlt(_localctx, 1);
        setState(127);
        statementOrDeclaration();
        break;
      }

      case TuringParser::CLASS: {
        enterOuterAlt(_localctx, 2);
        setState(128);
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

TuringParser::ProcHeaderContext* TuringParser::procHeader() {
  ProcHeaderContext *_localctx = _tracker.createInstance<ProcHeaderContext>(_ctx, getState());
  enterRule(_localctx, 4, TuringParser::RuleProcHeader);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(131);
    _la = _input->LA(1);
    if (!(_la == TuringParser::PROCEDURE

    || _la == TuringParser::PROCESS)) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
    setState(132);
    match(TuringParser::IDENTIFIER);
    setState(138);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 3, _ctx)) {
    case 1: {
      setState(133);
      match(TuringParser::L_PAREN);
      setState(135);
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
        setState(134);
        paramDeclarationList(0);
      }
      setState(137);
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

TuringParser::FuncHeaderContext* TuringParser::funcHeader() {
  FuncHeaderContext *_localctx = _tracker.createInstance<FuncHeaderContext>(_ctx, getState());
  enterRule(_localctx, 6, TuringParser::RuleFuncHeader);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(140);
    match(TuringParser::FUNCTION);
    setState(141);
    match(TuringParser::IDENTIFIER);
    setState(147);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == TuringParser::L_PAREN) {
      setState(142);
      match(TuringParser::L_PAREN);
      setState(144);
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
        setState(143);
        paramDeclarationList(0);
      }
      setState(146);
      match(TuringParser::R_PAREN);
    }
    setState(150);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == TuringParser::IDENTIFIER) {
      setState(149);
      match(TuringParser::IDENTIFIER);
    }
    setState(152);
    match(TuringParser::COLON);
    setState(153);
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

TuringParser::SubprogramHeaderContext* TuringParser::subprogramHeader() {
  SubprogramHeaderContext *_localctx = _tracker.createInstance<SubprogramHeaderContext>(_ctx, getState());
  enterRule(_localctx, 8, TuringParser::RuleSubprogramHeader);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(156);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::DEFERRED)
      | (1ULL << TuringParser::FORWARD)
      | (1ULL << TuringParser::BODY))) != 0)) {
      setState(155);
      subprogramPrefix();
    }
    setState(160);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::FUNCTION: {
        setState(158);
        funcHeader();
        break;
      }

      case TuringParser::PROCEDURE:
      case TuringParser::PROCESS: {
        setState(159);
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

TuringParser::SubprogramDeclarationContext* TuringParser::subprogramDeclaration() {
  SubprogramDeclarationContext *_localctx = _tracker.createInstance<SubprogramDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 10, TuringParser::RuleSubprogramDeclaration);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(162);
    subprogramHeader();
    setState(163);
    procBody();
    setState(164);
    match(TuringParser::END);
    setState(165);
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

TuringParser::SubprogramPrefixContext* TuringParser::subprogramPrefix() {
  SubprogramPrefixContext *_localctx = _tracker.createInstance<SubprogramPrefixContext>(_ctx, getState());
  enterRule(_localctx, 12, TuringParser::RuleSubprogramPrefix);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(167);
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

TuringParser::ClassDeclarationContext* TuringParser::classDeclaration() {
  ClassDeclarationContext *_localctx = _tracker.createInstance<ClassDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 14, TuringParser::RuleClassDeclaration);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(169);
    match(TuringParser::CLASS);
    setState(170);
    match(TuringParser::IDENTIFIER);
    setState(172); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(171);
      classBody();
      setState(174); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::STRING_LITERAL)
      | (1ULL << TuringParser::INTEGER_LITERAL)
      | (1ULL << TuringParser::TYPE)
      | (1ULL << TuringParser::VAR)
      | (1ULL << TuringParser::PROCEDURE)
      | (1ULL << TuringParser::FUNCTION)
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
      | (1ULL << TuringParser::CARET)
      | (1ULL << TuringParser::L_PAREN))) != 0) || ((((_la - 67) & ~ 0x3fULL) == 0) &&
      ((1ULL << (_la - 67)) & ((1ULL << (TuringParser::CHEAT - 67))
      | (1ULL << (TuringParser::PLUS - 67))
      | (1ULL << (TuringParser::MINUS - 67))
      | (1ULL << (TuringParser::IDENTIFIER - 67)))) != 0));
    setState(176);
    match(TuringParser::END);
    setState(177);
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

TuringParser::ClassBodyContext* TuringParser::classBody() {
  ClassBodyContext *_localctx = _tracker.createInstance<ClassBodyContext>(_ctx, getState());
  enterRule(_localctx, 16, TuringParser::RuleClassBody);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(183);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::EXPORT: {
        enterOuterAlt(_localctx, 1);
        setState(179);
        exportStatement();
        break;
      }

      case TuringParser::INHERIT: {
        enterOuterAlt(_localctx, 2);
        setState(180);
        inheritStatement();
        break;
      }

      case TuringParser::IMPLEMENT: {
        enterOuterAlt(_localctx, 3);
        setState(181);
        implementStatement();
        break;
      }

      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL:
      case TuringParser::TYPE:
      case TuringParser::VAR:
      case TuringParser::PROCEDURE:
      case TuringParser::FUNCTION:
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
        setState(182);
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

TuringParser::ExportStatementContext* TuringParser::exportStatement() {
  ExportStatementContext *_localctx = _tracker.createInstance<ExportStatementContext>(_ctx, getState());
  enterRule(_localctx, 18, TuringParser::RuleExportStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(192);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 11, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(185);
      match(TuringParser::EXPORT);
      setState(186);
      exportList(0);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(187);
      match(TuringParser::EXPORT);
      setState(188);
      match(TuringParser::L_PAREN);
      setState(189);
      exportList(0);
      setState(190);
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

TuringParser::ExportListItemContext* TuringParser::exportListItem() {
  ExportListItemContext *_localctx = _tracker.createInstance<ExportListItemContext>(_ctx, getState());
  enterRule(_localctx, 20, TuringParser::RuleExportListItem);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(195);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::T__0)
      | (1ULL << TuringParser::T__1)
      | (1ULL << TuringParser::T__2)
      | (1ULL << TuringParser::VAR))) != 0)) {
      setState(194);
      howExport();
    }
    setState(197);
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
    setState(200);
    exportListItem();
    _ctx->stop = _input->LT(-1);
    setState(207);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 13, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ExportListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleExportList);
        setState(202);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(203);
        match(TuringParser::COMMA);
        setState(204);
        exportListItem(); 
      }
      setState(209);
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

TuringParser::InheritStatementContext* TuringParser::inheritStatement() {
  InheritStatementContext *_localctx = _tracker.createInstance<InheritStatementContext>(_ctx, getState());
  enterRule(_localctx, 24, TuringParser::RuleInheritStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(217);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 14, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(210);
      match(TuringParser::INHERIT);
      setState(211);
      idOrFileItem();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(212);
      match(TuringParser::INHERIT);
      setState(213);
      match(TuringParser::L_PAREN);
      setState(214);
      idOrFileItem();
      setState(215);
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

TuringParser::ImplementStatementContext* TuringParser::implementStatement() {
  ImplementStatementContext *_localctx = _tracker.createInstance<ImplementStatementContext>(_ctx, getState());
  enterRule(_localctx, 26, TuringParser::RuleImplementStatement);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(232);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 17, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(219);
      match(TuringParser::IMPLEMENT);
      setState(221);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == TuringParser::BY) {
        setState(220);
        match(TuringParser::BY);
      }
      setState(223);
      idOrFileItem();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(224);
      match(TuringParser::IMPLEMENT);
      setState(226);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == TuringParser::BY) {
        setState(225);
        match(TuringParser::BY);
      }
      setState(228);
      match(TuringParser::L_PAREN);
      setState(229);
      idOrFileItem();
      setState(230);
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

TuringParser::IdOrFileItemContext* TuringParser::idOrFileItem() {
  IdOrFileItemContext *_localctx = _tracker.createInstance<IdOrFileItemContext>(_ctx, getState());
  enterRule(_localctx, 28, TuringParser::RuleIdOrFileItem);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(238);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 18, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(234);
      match(TuringParser::IDENTIFIER);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(235);
      match(TuringParser::IDENTIFIER);
      setState(236);
      match(TuringParser::IN);
      setState(237);
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

TuringParser::HowExportContext* TuringParser::howExport() {
  HowExportContext *_localctx = _tracker.createInstance<HowExportContext>(_ctx, getState());
  enterRule(_localctx, 30, TuringParser::RuleHowExport);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(240);
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

TuringParser::ParamDeclarationContext* TuringParser::paramDeclaration() {
  ParamDeclarationContext *_localctx = _tracker.createInstance<ParamDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 32, TuringParser::RuleParamDeclaration);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(250);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::VAR:
      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 1);
        setState(243);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == TuringParser::VAR) {
          setState(242);
          match(TuringParser::VAR);
        }
        setState(245);
        identifierList(0);
        setState(246);
        match(TuringParser::COLON);
        setState(247);
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
        setState(249);
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
    setState(253);
    paramDeclaration();
    _ctx->stop = _input->LT(-1);
    setState(260);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 21, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ParamDeclarationListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleParamDeclarationList);
        setState(255);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(256);
        match(TuringParser::COMMA);
        setState(257);
        paramDeclaration(); 
      }
      setState(262);
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

TuringParser::ProcBodyContext* TuringParser::procBody() {
  ProcBodyContext *_localctx = _tracker.createInstance<ProcBodyContext>(_ctx, getState());
  enterRule(_localctx, 36, TuringParser::RuleProcBody);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(264); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(263);
      statementOrDeclaration();
      setState(266); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << TuringParser::STRING_LITERAL)
      | (1ULL << TuringParser::INTEGER_LITERAL)
      | (1ULL << TuringParser::TYPE)
      | (1ULL << TuringParser::VAR)
      | (1ULL << TuringParser::PROCEDURE)
      | (1ULL << TuringParser::FUNCTION)
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
      | (1ULL << TuringParser::CARET)
      | (1ULL << TuringParser::L_PAREN))) != 0) || ((((_la - 67) & ~ 0x3fULL) == 0) &&
      ((1ULL << (_la - 67)) & ((1ULL << (TuringParser::CHEAT - 67))
      | (1ULL << (TuringParser::PLUS - 67))
      | (1ULL << (TuringParser::MINUS - 67))
      | (1ULL << (TuringParser::IDENTIFIER - 67)))) != 0));
   
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

TuringParser::StatementOrDeclarationContext* TuringParser::statementOrDeclaration() {
  StatementOrDeclarationContext *_localctx = _tracker.createInstance<StatementOrDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 38, TuringParser::RuleStatementOrDeclaration);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(270);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL:
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
        setState(268);
        statements();
        break;
      }

      case TuringParser::TYPE:
      case TuringParser::VAR:
      case TuringParser::PROCEDURE:
      case TuringParser::FUNCTION:
      case TuringParser::PROCESS:
      case TuringParser::ARRAY:
      case TuringParser::DEFERRED:
      case TuringParser::FORWARD:
      case TuringParser::BODY: {
        enterOuterAlt(_localctx, 2);
        setState(269);
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

TuringParser::PrimaryExpressionContext* TuringParser::primaryExpression() {
  PrimaryExpressionContext *_localctx = _tracker.createInstance<PrimaryExpressionContext>(_ctx, getState());
  enterRule(_localctx, 40, TuringParser::RulePrimaryExpression);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(278);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL: {
        enterOuterAlt(_localctx, 1);
        setState(272);
        literal();
        break;
      }

      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 2);
        setState(273);
        match(TuringParser::IDENTIFIER);
        break;
      }

      case TuringParser::L_PAREN: {
        enterOuterAlt(_localctx, 3);
        setState(274);
        match(TuringParser::L_PAREN);
        setState(275);
        expression();
        setState(276);
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

//----------------- ArgumentListContext ------------------------------------------------------------------

TuringParser::ArgumentListContext::ArgumentListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::ExpressionContext* TuringParser::ArgumentListContext::expression() {
  return getRuleContext<TuringParser::ExpressionContext>(0);
}

std::vector<TuringParser::ArgumentListContext *> TuringParser::ArgumentListContext::argumentList() {
  return getRuleContexts<TuringParser::ArgumentListContext>();
}

TuringParser::ArgumentListContext* TuringParser::ArgumentListContext::argumentList(size_t i) {
  return getRuleContext<TuringParser::ArgumentListContext>(i);
}

tree::TerminalNode* TuringParser::ArgumentListContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
}


size_t TuringParser::ArgumentListContext::getRuleIndex() const {
  return TuringParser::RuleArgumentList;
}

void TuringParser::ArgumentListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterArgumentList(this);
}

void TuringParser::ArgumentListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitArgumentList(this);
}


TuringParser::ArgumentListContext* TuringParser::argumentList() {
   return argumentList(0);
}

TuringParser::ArgumentListContext* TuringParser::argumentList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::ArgumentListContext *_localctx = _tracker.createInstance<ArgumentListContext>(_ctx, parentState);
  TuringParser::ArgumentListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 42;
  enterRecursionRule(_localctx, 42, TuringParser::RuleArgumentList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(281);
    expression();
    _ctx->stop = _input->LT(-1);
    setState(288);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 25, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ArgumentListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleArgumentList);
        setState(283);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(284);
        match(TuringParser::COMMA);
        setState(285);
        argumentList(2); 
      }
      setState(290);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 25, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- ExponentialExpressionContext ------------------------------------------------------------------

TuringParser::ExponentialExpressionContext::ExponentialExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<TuringParser::PrimaryExpressionContext *> TuringParser::ExponentialExpressionContext::primaryExpression() {
  return getRuleContexts<TuringParser::PrimaryExpressionContext>();
}

TuringParser::PrimaryExpressionContext* TuringParser::ExponentialExpressionContext::primaryExpression(size_t i) {
  return getRuleContext<TuringParser::PrimaryExpressionContext>(i);
}

tree::TerminalNode* TuringParser::ExponentialExpressionContext::EXP() {
  return getToken(TuringParser::EXP, 0);
}


size_t TuringParser::ExponentialExpressionContext::getRuleIndex() const {
  return TuringParser::RuleExponentialExpression;
}

void TuringParser::ExponentialExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExponentialExpression(this);
}

void TuringParser::ExponentialExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExponentialExpression(this);
}

TuringParser::ExponentialExpressionContext* TuringParser::exponentialExpression() {
  ExponentialExpressionContext *_localctx = _tracker.createInstance<ExponentialExpressionContext>(_ctx, getState());
  enterRule(_localctx, 44, TuringParser::RuleExponentialExpression);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(291);
    primaryExpression();
    setState(292);
    match(TuringParser::EXP);
    setState(293);
    primaryExpression();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PointerExpressionContext ------------------------------------------------------------------

TuringParser::PointerExpressionContext::PointerExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::PointerExpressionContext::CARET() {
  return getToken(TuringParser::CARET, 0);
}

TuringParser::PrimaryExpressionContext* TuringParser::PointerExpressionContext::primaryExpression() {
  return getRuleContext<TuringParser::PrimaryExpressionContext>(0);
}


size_t TuringParser::PointerExpressionContext::getRuleIndex() const {
  return TuringParser::RulePointerExpression;
}

void TuringParser::PointerExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterPointerExpression(this);
}

void TuringParser::PointerExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitPointerExpression(this);
}

TuringParser::PointerExpressionContext* TuringParser::pointerExpression() {
  PointerExpressionContext *_localctx = _tracker.createInstance<PointerExpressionContext>(_ctx, getState());
  enterRule(_localctx, 46, TuringParser::RulePointerExpression);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(295);
    match(TuringParser::CARET);
    setState(296);
    primaryExpression();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PostfixExpressionContext ------------------------------------------------------------------

TuringParser::PostfixExpressionContext::PostfixExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::PrimaryExpressionContext* TuringParser::PostfixExpressionContext::primaryExpression() {
  return getRuleContext<TuringParser::PrimaryExpressionContext>(0);
}

TuringParser::PointerExpressionContext* TuringParser::PostfixExpressionContext::pointerExpression() {
  return getRuleContext<TuringParser::PointerExpressionContext>(0);
}

TuringParser::ExponentialExpressionContext* TuringParser::PostfixExpressionContext::exponentialExpression() {
  return getRuleContext<TuringParser::ExponentialExpressionContext>(0);
}

TuringParser::PostfixExpressionContext* TuringParser::PostfixExpressionContext::postfixExpression() {
  return getRuleContext<TuringParser::PostfixExpressionContext>(0);
}

tree::TerminalNode* TuringParser::PostfixExpressionContext::L_PAREN() {
  return getToken(TuringParser::L_PAREN, 0);
}

tree::TerminalNode* TuringParser::PostfixExpressionContext::R_PAREN() {
  return getToken(TuringParser::R_PAREN, 0);
}

TuringParser::ArgumentListContext* TuringParser::PostfixExpressionContext::argumentList() {
  return getRuleContext<TuringParser::ArgumentListContext>(0);
}

tree::TerminalNode* TuringParser::PostfixExpressionContext::DOT() {
  return getToken(TuringParser::DOT, 0);
}

tree::TerminalNode* TuringParser::PostfixExpressionContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
}

tree::TerminalNode* TuringParser::PostfixExpressionContext::DEREFERENCE() {
  return getToken(TuringParser::DEREFERENCE, 0);
}


size_t TuringParser::PostfixExpressionContext::getRuleIndex() const {
  return TuringParser::RulePostfixExpression;
}

void TuringParser::PostfixExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterPostfixExpression(this);
}

void TuringParser::PostfixExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitPostfixExpression(this);
}


TuringParser::PostfixExpressionContext* TuringParser::postfixExpression() {
   return postfixExpression(0);
}

TuringParser::PostfixExpressionContext* TuringParser::postfixExpression(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::PostfixExpressionContext *_localctx = _tracker.createInstance<PostfixExpressionContext>(_ctx, parentState);
  TuringParser::PostfixExpressionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 48;
  enterRecursionRule(_localctx, 48, TuringParser::RulePostfixExpression, precedence);

    size_t _la = 0;

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(302);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 26, _ctx)) {
    case 1: {
      setState(299);
      primaryExpression();
      break;
    }

    case 2: {
      setState(300);
      pointerExpression();
      break;
    }

    case 3: {
      setState(301);
      exponentialExpression();
      break;
    }

    }
    _ctx->stop = _input->LT(-1);
    setState(318);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 29, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        setState(316);
        _errHandler->sync(this);
        switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 28, _ctx)) {
        case 1: {
          _localctx = _tracker.createInstance<PostfixExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RulePostfixExpression);
          setState(304);

          if (!(precpred(_ctx, 3))) throw FailedPredicateException(this, "precpred(_ctx, 3)");
          setState(305);
          match(TuringParser::L_PAREN);
          setState(307);
          _errHandler->sync(this);

          _la = _input->LA(1);
          if ((((_la & ~ 0x3fULL) == 0) &&
            ((1ULL << _la) & ((1ULL << TuringParser::STRING_LITERAL)
            | (1ULL << TuringParser::INTEGER_LITERAL)
            | (1ULL << TuringParser::NOT)
            | (1ULL << TuringParser::CARET)
            | (1ULL << TuringParser::L_PAREN))) != 0) || ((((_la - 67) & ~ 0x3fULL) == 0) &&
            ((1ULL << (_la - 67)) & ((1ULL << (TuringParser::CHEAT - 67))
            | (1ULL << (TuringParser::PLUS - 67))
            | (1ULL << (TuringParser::MINUS - 67))
            | (1ULL << (TuringParser::IDENTIFIER - 67)))) != 0)) {
            setState(306);
            argumentList(0);
          }
          setState(309);
          match(TuringParser::R_PAREN);
          break;
        }

        case 2: {
          _localctx = _tracker.createInstance<PostfixExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RulePostfixExpression);
          setState(310);

          if (!(precpred(_ctx, 2))) throw FailedPredicateException(this, "precpred(_ctx, 2)");
          setState(311);
          match(TuringParser::DOT);
          setState(312);
          match(TuringParser::IDENTIFIER);
          break;
        }

        case 3: {
          _localctx = _tracker.createInstance<PostfixExpressionContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RulePostfixExpression);
          setState(313);

          if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
          setState(314);
          match(TuringParser::DEREFERENCE);
          setState(315);
          match(TuringParser::IDENTIFIER);
          break;
        }

        } 
      }
      setState(320);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 29, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- PrefixExpressionContext ------------------------------------------------------------------

TuringParser::PrefixExpressionContext::PrefixExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::PostfixExpressionContext* TuringParser::PrefixExpressionContext::postfixExpression() {
  return getRuleContext<TuringParser::PostfixExpressionContext>(0);
}

tree::TerminalNode* TuringParser::PrefixExpressionContext::PLUS() {
  return getToken(TuringParser::PLUS, 0);
}

TuringParser::PrefixExpressionContext* TuringParser::PrefixExpressionContext::prefixExpression() {
  return getRuleContext<TuringParser::PrefixExpressionContext>(0);
}

tree::TerminalNode* TuringParser::PrefixExpressionContext::MINUS() {
  return getToken(TuringParser::MINUS, 0);
}

tree::TerminalNode* TuringParser::PrefixExpressionContext::CHEAT() {
  return getToken(TuringParser::CHEAT, 0);
}


size_t TuringParser::PrefixExpressionContext::getRuleIndex() const {
  return TuringParser::RulePrefixExpression;
}

void TuringParser::PrefixExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterPrefixExpression(this);
}

void TuringParser::PrefixExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitPrefixExpression(this);
}

TuringParser::PrefixExpressionContext* TuringParser::prefixExpression() {
  PrefixExpressionContext *_localctx = _tracker.createInstance<PrefixExpressionContext>(_ctx, getState());
  enterRule(_localctx, 50, TuringParser::RulePrefixExpression);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(328);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL:
      case TuringParser::CARET:
      case TuringParser::L_PAREN:
      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 1);
        setState(321);
        postfixExpression(0);
        break;
      }

      case TuringParser::PLUS: {
        enterOuterAlt(_localctx, 2);
        setState(322);
        match(TuringParser::PLUS);
        setState(323);
        prefixExpression();
        break;
      }

      case TuringParser::MINUS: {
        enterOuterAlt(_localctx, 3);
        setState(324);
        match(TuringParser::MINUS);
        setState(325);
        prefixExpression();
        break;
      }

      case TuringParser::CHEAT: {
        enterOuterAlt(_localctx, 4);
        setState(326);
        match(TuringParser::CHEAT);
        setState(327);
        prefixExpression();
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

//----------------- MultiplicativeExpressionContext ------------------------------------------------------------------

TuringParser::MultiplicativeExpressionContext::MultiplicativeExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::PostfixExpressionContext* TuringParser::MultiplicativeExpressionContext::postfixExpression() {
  return getRuleContext<TuringParser::PostfixExpressionContext>(0);
}

TuringParser::PrefixExpressionContext* TuringParser::MultiplicativeExpressionContext::prefixExpression() {
  return getRuleContext<TuringParser::PrefixExpressionContext>(0);
}

TuringParser::MultiplicativeExpressionContext* TuringParser::MultiplicativeExpressionContext::multiplicativeExpression() {
  return getRuleContext<TuringParser::MultiplicativeExpressionContext>(0);
}

TuringParser::MultiplicativeOperatorContext* TuringParser::MultiplicativeExpressionContext::multiplicativeOperator() {
  return getRuleContext<TuringParser::MultiplicativeOperatorContext>(0);
}


size_t TuringParser::MultiplicativeExpressionContext::getRuleIndex() const {
  return TuringParser::RuleMultiplicativeExpression;
}

void TuringParser::MultiplicativeExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterMultiplicativeExpression(this);
}

void TuringParser::MultiplicativeExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitMultiplicativeExpression(this);
}


TuringParser::MultiplicativeExpressionContext* TuringParser::multiplicativeExpression() {
   return multiplicativeExpression(0);
}

TuringParser::MultiplicativeExpressionContext* TuringParser::multiplicativeExpression(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::MultiplicativeExpressionContext *_localctx = _tracker.createInstance<MultiplicativeExpressionContext>(_ctx, parentState);
  TuringParser::MultiplicativeExpressionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 52;
  enterRecursionRule(_localctx, 52, TuringParser::RuleMultiplicativeExpression, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(333);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 31, _ctx)) {
    case 1: {
      setState(331);
      postfixExpression(0);
      break;
    }

    case 2: {
      setState(332);
      prefixExpression();
      break;
    }

    }
    _ctx->stop = _input->LT(-1);
    setState(343);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 33, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<MultiplicativeExpressionContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleMultiplicativeExpression);
        setState(335);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(336);
        multiplicativeOperator();
        setState(339);
        _errHandler->sync(this);
        switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 32, _ctx)) {
        case 1: {
          setState(337);
          postfixExpression(0);
          break;
        }

        case 2: {
          setState(338);
          prefixExpression();
          break;
        }

        } 
      }
      setState(345);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 33, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- MultiplicativeOperatorContext ------------------------------------------------------------------

TuringParser::MultiplicativeOperatorContext::MultiplicativeOperatorContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* TuringParser::MultiplicativeOperatorContext::MULTIPLY() {
  return getToken(TuringParser::MULTIPLY, 0);
}

tree::TerminalNode* TuringParser::MultiplicativeOperatorContext::DIVIDE() {
  return getToken(TuringParser::DIVIDE, 0);
}

tree::TerminalNode* TuringParser::MultiplicativeOperatorContext::DIV() {
  return getToken(TuringParser::DIV, 0);
}

tree::TerminalNode* TuringParser::MultiplicativeOperatorContext::MOD() {
  return getToken(TuringParser::MOD, 0);
}

tree::TerminalNode* TuringParser::MultiplicativeOperatorContext::REM() {
  return getToken(TuringParser::REM, 0);
}

tree::TerminalNode* TuringParser::MultiplicativeOperatorContext::SHR() {
  return getToken(TuringParser::SHR, 0);
}

tree::TerminalNode* TuringParser::MultiplicativeOperatorContext::SHL() {
  return getToken(TuringParser::SHL, 0);
}


size_t TuringParser::MultiplicativeOperatorContext::getRuleIndex() const {
  return TuringParser::RuleMultiplicativeOperator;
}

void TuringParser::MultiplicativeOperatorContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterMultiplicativeOperator(this);
}

void TuringParser::MultiplicativeOperatorContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitMultiplicativeOperator(this);
}

TuringParser::MultiplicativeOperatorContext* TuringParser::multiplicativeOperator() {
  MultiplicativeOperatorContext *_localctx = _tracker.createInstance<MultiplicativeOperatorContext>(_ctx, getState());
  enterRule(_localctx, 54, TuringParser::RuleMultiplicativeOperator);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(346);
    _la = _input->LA(1);
    if (!(((((_la - 72) & ~ 0x3fULL) == 0) &&
      ((1ULL << (_la - 72)) & ((1ULL << (TuringParser::MULTIPLY - 72))
      | (1ULL << (TuringParser::DIVIDE - 72))
      | (1ULL << (TuringParser::DIV - 72))
      | (1ULL << (TuringParser::MOD - 72))
      | (1ULL << (TuringParser::REM - 72))
      | (1ULL << (TuringParser::SHR - 72))
      | (1ULL << (TuringParser::SHL - 72)))) != 0))) {
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

//----------------- AdditiveExpressionContext ------------------------------------------------------------------

TuringParser::AdditiveExpressionContext::AdditiveExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::MultiplicativeExpressionContext* TuringParser::AdditiveExpressionContext::multiplicativeExpression() {
  return getRuleContext<TuringParser::MultiplicativeExpressionContext>(0);
}

TuringParser::AdditiveExpressionContext* TuringParser::AdditiveExpressionContext::additiveExpression() {
  return getRuleContext<TuringParser::AdditiveExpressionContext>(0);
}

tree::TerminalNode* TuringParser::AdditiveExpressionContext::PLUS() {
  return getToken(TuringParser::PLUS, 0);
}

tree::TerminalNode* TuringParser::AdditiveExpressionContext::MINUS() {
  return getToken(TuringParser::MINUS, 0);
}

tree::TerminalNode* TuringParser::AdditiveExpressionContext::XOR() {
  return getToken(TuringParser::XOR, 0);
}


size_t TuringParser::AdditiveExpressionContext::getRuleIndex() const {
  return TuringParser::RuleAdditiveExpression;
}

void TuringParser::AdditiveExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAdditiveExpression(this);
}

void TuringParser::AdditiveExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAdditiveExpression(this);
}


TuringParser::AdditiveExpressionContext* TuringParser::additiveExpression() {
   return additiveExpression(0);
}

TuringParser::AdditiveExpressionContext* TuringParser::additiveExpression(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::AdditiveExpressionContext *_localctx = _tracker.createInstance<AdditiveExpressionContext>(_ctx, parentState);
  TuringParser::AdditiveExpressionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 56;
  enterRecursionRule(_localctx, 56, TuringParser::RuleAdditiveExpression, precedence);

    size_t _la = 0;

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(349);
    multiplicativeExpression(0);
    _ctx->stop = _input->LT(-1);
    setState(356);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 34, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<AdditiveExpressionContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleAdditiveExpression);
        setState(351);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(352);
        _la = _input->LA(1);
        if (!(((((_la - 70) & ~ 0x3fULL) == 0) &&
          ((1ULL << (_la - 70)) & ((1ULL << (TuringParser::PLUS - 70))
          | (1ULL << (TuringParser::MINUS - 70))
          | (1ULL << (TuringParser::XOR - 70)))) != 0))) {
        _errHandler->recoverInline(this);
        }
        else {
          _errHandler->reportMatch(this);
          consume();
        }
        setState(353);
        multiplicativeExpression(0); 
      }
      setState(358);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 34, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- ComparativeExpressionContext ------------------------------------------------------------------

TuringParser::ComparativeExpressionContext::ComparativeExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::AdditiveExpressionContext* TuringParser::ComparativeExpressionContext::additiveExpression() {
  return getRuleContext<TuringParser::AdditiveExpressionContext>(0);
}

TuringParser::ComparativeExpressionContext* TuringParser::ComparativeExpressionContext::comparativeExpression() {
  return getRuleContext<TuringParser::ComparativeExpressionContext>(0);
}

tree::TerminalNode* TuringParser::ComparativeExpressionContext::LESSTHAN() {
  return getToken(TuringParser::LESSTHAN, 0);
}

tree::TerminalNode* TuringParser::ComparativeExpressionContext::GREATERTHAN() {
  return getToken(TuringParser::GREATERTHAN, 0);
}

tree::TerminalNode* TuringParser::ComparativeExpressionContext::EQUAL() {
  return getToken(TuringParser::EQUAL, 0);
}

tree::TerminalNode* TuringParser::ComparativeExpressionContext::LESSTHANEQUAL() {
  return getToken(TuringParser::LESSTHANEQUAL, 0);
}

tree::TerminalNode* TuringParser::ComparativeExpressionContext::GREATERTHANEQUAL() {
  return getToken(TuringParser::GREATERTHANEQUAL, 0);
}

tree::TerminalNode* TuringParser::ComparativeExpressionContext::NOTEQUAL() {
  return getToken(TuringParser::NOTEQUAL, 0);
}

tree::TerminalNode* TuringParser::ComparativeExpressionContext::IN() {
  return getToken(TuringParser::IN, 0);
}

tree::TerminalNode* TuringParser::ComparativeExpressionContext::NOT() {
  return getToken(TuringParser::NOT, 0);
}


size_t TuringParser::ComparativeExpressionContext::getRuleIndex() const {
  return TuringParser::RuleComparativeExpression;
}

void TuringParser::ComparativeExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterComparativeExpression(this);
}

void TuringParser::ComparativeExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitComparativeExpression(this);
}


TuringParser::ComparativeExpressionContext* TuringParser::comparativeExpression() {
   return comparativeExpression(0);
}

TuringParser::ComparativeExpressionContext* TuringParser::comparativeExpression(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::ComparativeExpressionContext *_localctx = _tracker.createInstance<ComparativeExpressionContext>(_ctx, parentState);
  TuringParser::ComparativeExpressionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 58;
  enterRecursionRule(_localctx, 58, TuringParser::RuleComparativeExpression, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(360);
    additiveExpression(0);
    _ctx->stop = _input->LT(-1);
    setState(377);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 36, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ComparativeExpressionContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleComparativeExpression);
        setState(362);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(372);
        _errHandler->sync(this);
        switch (_input->LA(1)) {
          case TuringParser::LESSTHAN: {
            setState(363);
            match(TuringParser::LESSTHAN);
            break;
          }

          case TuringParser::GREATERTHAN: {
            setState(364);
            match(TuringParser::GREATERTHAN);
            break;
          }

          case TuringParser::EQUAL: {
            setState(365);
            match(TuringParser::EQUAL);
            break;
          }

          case TuringParser::LESSTHANEQUAL: {
            setState(366);
            match(TuringParser::LESSTHANEQUAL);
            break;
          }

          case TuringParser::GREATERTHANEQUAL: {
            setState(367);
            match(TuringParser::GREATERTHANEQUAL);
            break;
          }

          case TuringParser::NOTEQUAL: {
            setState(368);
            match(TuringParser::NOTEQUAL);
            break;
          }

          case TuringParser::IN: {
            setState(369);
            match(TuringParser::IN);
            break;
          }

          case TuringParser::NOT: {
            setState(370);
            match(TuringParser::NOT);
            setState(371);
            match(TuringParser::IN);
            break;
          }

        default:
          throw NoViableAltException(this);
        }
        setState(374);
        additiveExpression(0); 
      }
      setState(379);
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

//----------------- NotExpressionContext ------------------------------------------------------------------

TuringParser::NotExpressionContext::NotExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::ComparativeExpressionContext* TuringParser::NotExpressionContext::comparativeExpression() {
  return getRuleContext<TuringParser::ComparativeExpressionContext>(0);
}

tree::TerminalNode* TuringParser::NotExpressionContext::NOT() {
  return getToken(TuringParser::NOT, 0);
}

TuringParser::NotExpressionContext* TuringParser::NotExpressionContext::notExpression() {
  return getRuleContext<TuringParser::NotExpressionContext>(0);
}


size_t TuringParser::NotExpressionContext::getRuleIndex() const {
  return TuringParser::RuleNotExpression;
}

void TuringParser::NotExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterNotExpression(this);
}

void TuringParser::NotExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitNotExpression(this);
}

TuringParser::NotExpressionContext* TuringParser::notExpression() {
  NotExpressionContext *_localctx = _tracker.createInstance<NotExpressionContext>(_ctx, getState());
  enterRule(_localctx, 60, TuringParser::RuleNotExpression);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(383);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::STRING_LITERAL:
      case TuringParser::INTEGER_LITERAL:
      case TuringParser::CARET:
      case TuringParser::L_PAREN:
      case TuringParser::CHEAT:
      case TuringParser::PLUS:
      case TuringParser::MINUS:
      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 1);
        setState(380);
        comparativeExpression(0);
        break;
      }

      case TuringParser::NOT: {
        enterOuterAlt(_localctx, 2);
        setState(381);
        match(TuringParser::NOT);
        setState(382);
        notExpression();
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

//----------------- AndExpressionContext ------------------------------------------------------------------

TuringParser::AndExpressionContext::AndExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::NotExpressionContext* TuringParser::AndExpressionContext::notExpression() {
  return getRuleContext<TuringParser::NotExpressionContext>(0);
}

TuringParser::AndExpressionContext* TuringParser::AndExpressionContext::andExpression() {
  return getRuleContext<TuringParser::AndExpressionContext>(0);
}

tree::TerminalNode* TuringParser::AndExpressionContext::AND() {
  return getToken(TuringParser::AND, 0);
}


size_t TuringParser::AndExpressionContext::getRuleIndex() const {
  return TuringParser::RuleAndExpression;
}

void TuringParser::AndExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAndExpression(this);
}

void TuringParser::AndExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAndExpression(this);
}


TuringParser::AndExpressionContext* TuringParser::andExpression() {
   return andExpression(0);
}

TuringParser::AndExpressionContext* TuringParser::andExpression(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::AndExpressionContext *_localctx = _tracker.createInstance<AndExpressionContext>(_ctx, parentState);
  TuringParser::AndExpressionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 62;
  enterRecursionRule(_localctx, 62, TuringParser::RuleAndExpression, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(386);
    notExpression();
    _ctx->stop = _input->LT(-1);
    setState(393);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 38, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<AndExpressionContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleAndExpression);
        setState(388);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(389);
        match(TuringParser::AND);
        setState(390);
        notExpression(); 
      }
      setState(395);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 38, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- OrExpressionContext ------------------------------------------------------------------

TuringParser::OrExpressionContext::OrExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::AndExpressionContext* TuringParser::OrExpressionContext::andExpression() {
  return getRuleContext<TuringParser::AndExpressionContext>(0);
}

TuringParser::OrExpressionContext* TuringParser::OrExpressionContext::orExpression() {
  return getRuleContext<TuringParser::OrExpressionContext>(0);
}

tree::TerminalNode* TuringParser::OrExpressionContext::OR() {
  return getToken(TuringParser::OR, 0);
}


size_t TuringParser::OrExpressionContext::getRuleIndex() const {
  return TuringParser::RuleOrExpression;
}

void TuringParser::OrExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterOrExpression(this);
}

void TuringParser::OrExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitOrExpression(this);
}


TuringParser::OrExpressionContext* TuringParser::orExpression() {
   return orExpression(0);
}

TuringParser::OrExpressionContext* TuringParser::orExpression(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::OrExpressionContext *_localctx = _tracker.createInstance<OrExpressionContext>(_ctx, parentState);
  TuringParser::OrExpressionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 64;
  enterRecursionRule(_localctx, 64, TuringParser::RuleOrExpression, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(397);
    andExpression(0);
    _ctx->stop = _input->LT(-1);
    setState(404);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 39, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<OrExpressionContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleOrExpression);
        setState(399);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(400);
        match(TuringParser::OR);
        setState(401);
        andExpression(0); 
      }
      setState(406);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 39, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- ImpliesExpressionContext ------------------------------------------------------------------

TuringParser::ImpliesExpressionContext::ImpliesExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::OrExpressionContext* TuringParser::ImpliesExpressionContext::orExpression() {
  return getRuleContext<TuringParser::OrExpressionContext>(0);
}

TuringParser::ImpliesExpressionContext* TuringParser::ImpliesExpressionContext::impliesExpression() {
  return getRuleContext<TuringParser::ImpliesExpressionContext>(0);
}

tree::TerminalNode* TuringParser::ImpliesExpressionContext::IMPLIES() {
  return getToken(TuringParser::IMPLIES, 0);
}


size_t TuringParser::ImpliesExpressionContext::getRuleIndex() const {
  return TuringParser::RuleImpliesExpression;
}

void TuringParser::ImpliesExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterImpliesExpression(this);
}

void TuringParser::ImpliesExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitImpliesExpression(this);
}


TuringParser::ImpliesExpressionContext* TuringParser::impliesExpression() {
   return impliesExpression(0);
}

TuringParser::ImpliesExpressionContext* TuringParser::impliesExpression(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::ImpliesExpressionContext *_localctx = _tracker.createInstance<ImpliesExpressionContext>(_ctx, parentState);
  TuringParser::ImpliesExpressionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 66;
  enterRecursionRule(_localctx, 66, TuringParser::RuleImpliesExpression, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(408);
    orExpression(0);
    _ctx->stop = _input->LT(-1);
    setState(415);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 40, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ImpliesExpressionContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleImpliesExpression);
        setState(410);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(411);
        match(TuringParser::IMPLIES);
        setState(412);
        orExpression(0); 
      }
      setState(417);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 40, _ctx);
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

TuringParser::ImpliesExpressionContext* TuringParser::ExpressionContext::impliesExpression() {
  return getRuleContext<TuringParser::ImpliesExpressionContext>(0);
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

TuringParser::ExpressionContext* TuringParser::expression() {
  ExpressionContext *_localctx = _tracker.createInstance<ExpressionContext>(_ctx, getState());
  enterRule(_localctx, 68, TuringParser::RuleExpression);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(418);
    impliesExpression(0);
   
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

TuringParser::ExpressionContext* TuringParser::ExpressionListContext::expression() {
  return getRuleContext<TuringParser::ExpressionContext>(0);
}

TuringParser::ExpressionListContext* TuringParser::ExpressionListContext::expressionList() {
  return getRuleContext<TuringParser::ExpressionListContext>(0);
}

tree::TerminalNode* TuringParser::ExpressionListContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
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


TuringParser::ExpressionListContext* TuringParser::expressionList() {
   return expressionList(0);
}

TuringParser::ExpressionListContext* TuringParser::expressionList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::ExpressionListContext *_localctx = _tracker.createInstance<ExpressionListContext>(_ctx, parentState);
  TuringParser::ExpressionListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 70;
  enterRecursionRule(_localctx, 70, TuringParser::RuleExpressionList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(421);
    expression();
    _ctx->stop = _input->LT(-1);
    setState(428);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 41, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ExpressionListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleExpressionList);
        setState(423);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(424);
        match(TuringParser::COMMA);
        setState(425);
        expression(); 
      }
      setState(430);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 41, _ctx);
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

TuringParser::DeclarationContext* TuringParser::declaration() {
  DeclarationContext *_localctx = _tracker.createInstance<DeclarationContext>(_ctx, getState());
  enterRule(_localctx, 72, TuringParser::RuleDeclaration);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(435);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::TYPE: {
        enterOuterAlt(_localctx, 1);
        setState(431);
        typeDeclaration();
        break;
      }

      case TuringParser::VAR: {
        enterOuterAlt(_localctx, 2);
        setState(432);
        variableDeclaration();
        break;
      }

      case TuringParser::ARRAY: {
        enterOuterAlt(_localctx, 3);
        setState(433);
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
        setState(434);
        subprogramDeclaration();
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

TuringParser::AssignmentStatementContext* TuringParser::StatementsContext::assignmentStatement() {
  return getRuleContext<TuringParser::AssignmentStatementContext>(0);
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

TuringParser::StatementsContext* TuringParser::statements() {
  StatementsContext *_localctx = _tracker.createInstance<StatementsContext>(_ctx, getState());
  enterRule(_localctx, 74, TuringParser::RuleStatements);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(447);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 43, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(437);
      expression();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(438);
      assignmentStatement();
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);
      setState(439);
      putStatement();
      break;
    }

    case 4: {
      enterOuterAlt(_localctx, 4);
      setState(440);
      match(TuringParser::EXIT);
      break;
    }

    case 5: {
      enterOuterAlt(_localctx, 5);
      setState(441);
      beginStatement();
      break;
    }

    case 6: {
      enterOuterAlt(_localctx, 6);
      setState(442);
      match(TuringParser::RETURN);
      break;
    }

    case 7: {
      enterOuterAlt(_localctx, 7);
      setState(443);
      resultStatement();
      break;
    }

    case 8: {
      enterOuterAlt(_localctx, 8);
      setState(444);
      newStatement();
      break;
    }

    case 9: {
      enterOuterAlt(_localctx, 9);
      setState(445);
      freeStatement();
      break;
    }

    case 10: {
      enterOuterAlt(_localctx, 10);
      setState(446);
      forkStatement();
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

//----------------- AssignmentStatementContext ------------------------------------------------------------------

TuringParser::AssignmentStatementContext::AssignmentStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

TuringParser::PostfixExpressionContext* TuringParser::AssignmentStatementContext::postfixExpression() {
  return getRuleContext<TuringParser::PostfixExpressionContext>(0);
}

TuringParser::ExpressionContext* TuringParser::AssignmentStatementContext::expression() {
  return getRuleContext<TuringParser::ExpressionContext>(0);
}

tree::TerminalNode* TuringParser::AssignmentStatementContext::ASSIGNMENT() {
  return getToken(TuringParser::ASSIGNMENT, 0);
}

tree::TerminalNode* TuringParser::AssignmentStatementContext::PLUSEQUALS() {
  return getToken(TuringParser::PLUSEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentStatementContext::MINUSEQUALS() {
  return getToken(TuringParser::MINUSEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentStatementContext::MULTIPLYEQUALS() {
  return getToken(TuringParser::MULTIPLYEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentStatementContext::DIVIDEEQUALS() {
  return getToken(TuringParser::DIVIDEEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentStatementContext::DIVEQUALS() {
  return getToken(TuringParser::DIVEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentStatementContext::SHLEQUALS() {
  return getToken(TuringParser::SHLEQUALS, 0);
}

tree::TerminalNode* TuringParser::AssignmentStatementContext::SHREQUALS() {
  return getToken(TuringParser::SHREQUALS, 0);
}


size_t TuringParser::AssignmentStatementContext::getRuleIndex() const {
  return TuringParser::RuleAssignmentStatement;
}

void TuringParser::AssignmentStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAssignmentStatement(this);
}

void TuringParser::AssignmentStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<TuringListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAssignmentStatement(this);
}

TuringParser::AssignmentStatementContext* TuringParser::assignmentStatement() {
  AssignmentStatementContext *_localctx = _tracker.createInstance<AssignmentStatementContext>(_ctx, getState());
  enterRule(_localctx, 76, TuringParser::RuleAssignmentStatement);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(449);
    postfixExpression(0);
    setState(450);
    _la = _input->LA(1);
    if (!(((((_la - 69) & ~ 0x3fULL) == 0) &&
      ((1ULL << (_la - 69)) & ((1ULL << (TuringParser::ASSIGNMENT - 69))
      | (1ULL << (TuringParser::PLUSEQUALS - 69))
      | (1ULL << (TuringParser::MINUSEQUALS - 69))
      | (1ULL << (TuringParser::MULTIPLYEQUALS - 69))
      | (1ULL << (TuringParser::DIVIDEEQUALS - 69))
      | (1ULL << (TuringParser::DIVEQUALS - 69))
      | (1ULL << (TuringParser::SHREQUALS - 69))
      | (1ULL << (TuringParser::SHLEQUALS - 69)))) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
    setState(451);
    expression();
   
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

TuringParser::PutStatementContext* TuringParser::putStatement() {
  PutStatementContext *_localctx = _tracker.createInstance<PutStatementContext>(_ctx, getState());
  enterRule(_localctx, 78, TuringParser::RulePutStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(453);
    match(TuringParser::PUT);
    setState(454);
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

TuringParser::PutItemContext* TuringParser::putItem() {
  PutItemContext *_localctx = _tracker.createInstance<PutItemContext>(_ctx, getState());
  enterRule(_localctx, 80, TuringParser::RulePutItem);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(456);
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


TuringParser::PutItemListContext* TuringParser::putItemList() {
   return putItemList(0);
}

TuringParser::PutItemListContext* TuringParser::putItemList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::PutItemListContext *_localctx = _tracker.createInstance<PutItemListContext>(_ctx, parentState);
  TuringParser::PutItemListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 82;
  enterRecursionRule(_localctx, 82, TuringParser::RulePutItemList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(459);
    putItem();
    _ctx->stop = _input->LT(-1);
    setState(466);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 44, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<PutItemListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RulePutItemList);
        setState(461);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(462);
        match(TuringParser::COMMA);
        setState(463);
        putItem(); 
      }
      setState(468);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 44, _ctx);
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

TuringParser::BeginStatementContext* TuringParser::beginStatement() {
  BeginStatementContext *_localctx = _tracker.createInstance<BeginStatementContext>(_ctx, getState());
  enterRule(_localctx, 84, TuringParser::RuleBeginStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(469);
    match(TuringParser::BEGIN);
    setState(470);
    statementOrDeclaration();
    setState(471);
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

TuringParser::ResultStatementContext* TuringParser::resultStatement() {
  ResultStatementContext *_localctx = _tracker.createInstance<ResultStatementContext>(_ctx, getState());
  enterRule(_localctx, 86, TuringParser::RuleResultStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(473);
    match(TuringParser::RESULT);
    setState(474);
    expression();
   
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

TuringParser::NewStatementContext* TuringParser::newStatement() {
  NewStatementContext *_localctx = _tracker.createInstance<NewStatementContext>(_ctx, getState());
  enterRule(_localctx, 88, TuringParser::RuleNewStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(476);
    match(TuringParser::NEW);
    setState(479);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 45, _ctx)) {
    case 1: {
      setState(477);
      match(TuringParser::IDENTIFIER);
      setState(478);
      match(TuringParser::COMMA);
      break;
    }

    }
    setState(481);
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

TuringParser::FreeStatementContext* TuringParser::freeStatement() {
  FreeStatementContext *_localctx = _tracker.createInstance<FreeStatementContext>(_ctx, getState());
  enterRule(_localctx, 90, TuringParser::RuleFreeStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(483);
    match(TuringParser::FREE);
    setState(486);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 46, _ctx)) {
    case 1: {
      setState(484);
      match(TuringParser::IDENTIFIER);
      setState(485);
      match(TuringParser::COMMA);
      break;
    }

    }
    setState(488);
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

TuringParser::ForkStatementContext* TuringParser::forkStatement() {
  ForkStatementContext *_localctx = _tracker.createInstance<ForkStatementContext>(_ctx, getState());
  enterRule(_localctx, 92, TuringParser::RuleForkStatement);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(490);
    match(TuringParser::FORK);
    setState(491);
    expression();
   
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

TuringParser::TypeDeclarationContext* TuringParser::typeDeclaration() {
  TypeDeclarationContext *_localctx = _tracker.createInstance<TypeDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 94, TuringParser::RuleTypeDeclaration);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(493);
    match(TuringParser::TYPE);
    setState(494);
    match(TuringParser::IDENTIFIER);
    setState(495);
    match(TuringParser::COLON);
    setState(496);
    typeSpec();
   
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

tree::TerminalNode* TuringParser::TypeSpecContext::INT() {
  return getToken(TuringParser::INT, 0);
}

tree::TerminalNode* TuringParser::TypeSpecContext::REAL() {
  return getToken(TuringParser::REAL, 0);
}

tree::TerminalNode* TuringParser::TypeSpecContext::BOOLEAN() {
  return getToken(TuringParser::BOOLEAN, 0);
}

tree::TerminalNode* TuringParser::TypeSpecContext::NAT() {
  return getToken(TuringParser::NAT, 0);
}

tree::TerminalNode* TuringParser::TypeSpecContext::INTN() {
  return getToken(TuringParser::INTN, 0);
}

tree::TerminalNode* TuringParser::TypeSpecContext::NATN() {
  return getToken(TuringParser::NATN, 0);
}

tree::TerminalNode* TuringParser::TypeSpecContext::REALN() {
  return getToken(TuringParser::REALN, 0);
}

tree::TerminalNode* TuringParser::TypeSpecContext::CHAR() {
  return getToken(TuringParser::CHAR, 0);
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

tree::TerminalNode* TuringParser::TypeSpecContext::IDENTIFIER() {
  return getToken(TuringParser::IDENTIFIER, 0);
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

TuringParser::TypeSpecContext* TuringParser::typeSpec() {
  TypeSpecContext *_localctx = _tracker.createInstance<TypeSpecContext>(_ctx, getState());
  enterRule(_localctx, 96, TuringParser::RuleTypeSpec);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(510);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case TuringParser::INT: {
        enterOuterAlt(_localctx, 1);
        setState(498);
        match(TuringParser::INT);
        break;
      }

      case TuringParser::REAL: {
        enterOuterAlt(_localctx, 2);
        setState(499);
        match(TuringParser::REAL);
        break;
      }

      case TuringParser::BOOLEAN: {
        enterOuterAlt(_localctx, 3);
        setState(500);
        match(TuringParser::BOOLEAN);
        break;
      }

      case TuringParser::NAT: {
        enterOuterAlt(_localctx, 4);
        setState(501);
        match(TuringParser::NAT);
        break;
      }

      case TuringParser::INTN: {
        enterOuterAlt(_localctx, 5);
        setState(502);
        match(TuringParser::INTN);
        break;
      }

      case TuringParser::NATN: {
        enterOuterAlt(_localctx, 6);
        setState(503);
        match(TuringParser::NATN);
        break;
      }

      case TuringParser::REALN: {
        enterOuterAlt(_localctx, 7);
        setState(504);
        match(TuringParser::REALN);
        break;
      }

      case TuringParser::CHAR: {
        enterOuterAlt(_localctx, 8);
        setState(505);
        match(TuringParser::CHAR);
        break;
      }

      case TuringParser::INTEGER_LITERAL: {
        enterOuterAlt(_localctx, 9);
        setState(506);
        indexType();
        break;
      }

      case TuringParser::STRING: {
        enterOuterAlt(_localctx, 10);
        setState(507);
        stringType();
        break;
      }

      case TuringParser::RECORD: {
        enterOuterAlt(_localctx, 11);
        setState(508);
        recordType();
        break;
      }

      case TuringParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 12);
        setState(509);
        match(TuringParser::IDENTIFIER);
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

TuringParser::IndexTypeContext* TuringParser::indexType() {
  IndexTypeContext *_localctx = _tracker.createInstance<IndexTypeContext>(_ctx, getState());
  enterRule(_localctx, 98, TuringParser::RuleIndexType);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(512);
    match(TuringParser::INTEGER_LITERAL);
    setState(513);
    match(TuringParser::RANGE);
    setState(514);
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


TuringParser::IndexTypeListContext* TuringParser::indexTypeList() {
   return indexTypeList(0);
}

TuringParser::IndexTypeListContext* TuringParser::indexTypeList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::IndexTypeListContext *_localctx = _tracker.createInstance<IndexTypeListContext>(_ctx, parentState);
  TuringParser::IndexTypeListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 100;
  enterRecursionRule(_localctx, 100, TuringParser::RuleIndexTypeList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(517);
    indexType();
    _ctx->stop = _input->LT(-1);
    setState(524);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 48, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<IndexTypeListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleIndexTypeList);
        setState(519);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(520);
        match(TuringParser::COMMA);
        setState(521);
        indexType(); 
      }
      setState(526);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 48, _ctx);
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

TuringParser::StringTypeContext* TuringParser::stringType() {
  StringTypeContext *_localctx = _tracker.createInstance<StringTypeContext>(_ctx, getState());
  enterRule(_localctx, 102, TuringParser::RuleStringType);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(527);
    match(TuringParser::STRING);
    setState(531);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 49, _ctx)) {
    case 1: {
      setState(528);
      match(TuringParser::L_PAREN);
      setState(529);
      match(TuringParser::INTEGER_LITERAL);
      setState(530);
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

TuringParser::RecordTypeContext* TuringParser::recordType() {
  RecordTypeContext *_localctx = _tracker.createInstance<RecordTypeContext>(_ctx, getState());
  enterRule(_localctx, 104, TuringParser::RuleRecordType);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(533);
    match(TuringParser::RECORD);
    setState(535); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(534);
      recordField();
      setState(537); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while (_la == TuringParser::IDENTIFIER);
    setState(539);
    match(TuringParser::END);
    setState(540);
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

TuringParser::RecordFieldContext* TuringParser::recordField() {
  RecordFieldContext *_localctx = _tracker.createInstance<RecordFieldContext>(_ctx, getState());
  enterRule(_localctx, 106, TuringParser::RuleRecordField);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(542);
    identifierList(0);
    setState(543);
    match(TuringParser::COLON);
    setState(544);
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

TuringParser::VariableDeclarationContext* TuringParser::variableDeclaration() {
  VariableDeclarationContext *_localctx = _tracker.createInstance<VariableDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 108, TuringParser::RuleVariableDeclaration);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(559);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 52, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(546);
      match(TuringParser::VAR);
      setState(547);
      variableIdentifierList(0);
      setState(548);
      match(TuringParser::COLON);
      setState(549);
      typeSpec();
      setState(552);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == TuringParser::ASSIGNMENT) {
        setState(550);
        match(TuringParser::ASSIGNMENT);
        setState(551);
        expression();
      }
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(554);
      match(TuringParser::VAR);
      setState(555);
      variableIdentifierList(0);
      setState(556);
      match(TuringParser::ASSIGNMENT);
      setState(557);
      expression();
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

TuringParser::VariableIdentifierContext* TuringParser::VariableIdentifierListContext::variableIdentifier() {
  return getRuleContext<TuringParser::VariableIdentifierContext>(0);
}

TuringParser::VariableIdentifierListContext* TuringParser::VariableIdentifierListContext::variableIdentifierList() {
  return getRuleContext<TuringParser::VariableIdentifierListContext>(0);
}

tree::TerminalNode* TuringParser::VariableIdentifierListContext::COMMA() {
  return getToken(TuringParser::COMMA, 0);
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


TuringParser::VariableIdentifierListContext* TuringParser::variableIdentifierList() {
   return variableIdentifierList(0);
}

TuringParser::VariableIdentifierListContext* TuringParser::variableIdentifierList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::VariableIdentifierListContext *_localctx = _tracker.createInstance<VariableIdentifierListContext>(_ctx, parentState);
  TuringParser::VariableIdentifierListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 110;
  enterRecursionRule(_localctx, 110, TuringParser::RuleVariableIdentifierList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(562);
    variableIdentifier();
    _ctx->stop = _input->LT(-1);
    setState(569);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 53, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<VariableIdentifierListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleVariableIdentifierList);
        setState(564);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(565);
        match(TuringParser::COMMA);
        setState(566);
        variableIdentifier(); 
      }
      setState(571);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 53, _ctx);
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

TuringParser::VariableIdentifierContext* TuringParser::variableIdentifier() {
  VariableIdentifierContext *_localctx = _tracker.createInstance<VariableIdentifierContext>(_ctx, getState());
  enterRule(_localctx, 112, TuringParser::RuleVariableIdentifier);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(572);
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

TuringParser::ArrayDeclarationContext* TuringParser::arrayDeclaration() {
  ArrayDeclarationContext *_localctx = _tracker.createInstance<ArrayDeclarationContext>(_ctx, getState());
  enterRule(_localctx, 114, TuringParser::RuleArrayDeclaration);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(583);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 54, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(574);
      match(TuringParser::ARRAY);
      setState(575);
      indexType();
      setState(576);
      match(TuringParser::OF);
      setState(577);
      typeSpec();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(579);
      match(TuringParser::ARRAY);
      setState(580);
      indexType();
      setState(581);
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


TuringParser::IdentifierListContext* TuringParser::identifierList() {
   return identifierList(0);
}

TuringParser::IdentifierListContext* TuringParser::identifierList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  TuringParser::IdentifierListContext *_localctx = _tracker.createInstance<IdentifierListContext>(_ctx, parentState);
  TuringParser::IdentifierListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 116;
  enterRecursionRule(_localctx, 116, TuringParser::RuleIdentifierList, precedence);

    

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(586);
    match(TuringParser::IDENTIFIER);
    _ctx->stop = _input->LT(-1);
    setState(593);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 55, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<IdentifierListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleIdentifierList);
        setState(588);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(589);
        match(TuringParser::COMMA);
        setState(590);
        match(TuringParser::IDENTIFIER); 
      }
      setState(595);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 55, _ctx);
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

TuringParser::LiteralContext* TuringParser::literal() {
  LiteralContext *_localctx = _tracker.createInstance<LiteralContext>(_ctx, getState());
  enterRule(_localctx, 118, TuringParser::RuleLiteral);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(596);
    _la = _input->LA(1);
    if (!(_la == TuringParser::STRING_LITERAL

    || _la == TuringParser::INTEGER_LITERAL)) {
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

TuringParser::CommentContext* TuringParser::comment() {
  CommentContext *_localctx = _tracker.createInstance<CommentContext>(_ctx, getState());
  enterRule(_localctx, 120, TuringParser::RuleComment);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(598);
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
    case 21: return argumentListSempred(dynamic_cast<ArgumentListContext *>(context), predicateIndex);
    case 24: return postfixExpressionSempred(dynamic_cast<PostfixExpressionContext *>(context), predicateIndex);
    case 26: return multiplicativeExpressionSempred(dynamic_cast<MultiplicativeExpressionContext *>(context), predicateIndex);
    case 28: return additiveExpressionSempred(dynamic_cast<AdditiveExpressionContext *>(context), predicateIndex);
    case 29: return comparativeExpressionSempred(dynamic_cast<ComparativeExpressionContext *>(context), predicateIndex);
    case 31: return andExpressionSempred(dynamic_cast<AndExpressionContext *>(context), predicateIndex);
    case 32: return orExpressionSempred(dynamic_cast<OrExpressionContext *>(context), predicateIndex);
    case 33: return impliesExpressionSempred(dynamic_cast<ImpliesExpressionContext *>(context), predicateIndex);
    case 35: return expressionListSempred(dynamic_cast<ExpressionListContext *>(context), predicateIndex);
    case 41: return putItemListSempred(dynamic_cast<PutItemListContext *>(context), predicateIndex);
    case 50: return indexTypeListSempred(dynamic_cast<IndexTypeListContext *>(context), predicateIndex);
    case 55: return variableIdentifierListSempred(dynamic_cast<VariableIdentifierListContext *>(context), predicateIndex);
    case 58: return identifierListSempred(dynamic_cast<IdentifierListContext *>(context), predicateIndex);

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

bool TuringParser::argumentListSempred(ArgumentListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 2: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::postfixExpressionSempred(PostfixExpressionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 3: return precpred(_ctx, 3);
    case 4: return precpred(_ctx, 2);
    case 5: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::multiplicativeExpressionSempred(MultiplicativeExpressionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 6: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::additiveExpressionSempred(AdditiveExpressionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 7: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::comparativeExpressionSempred(ComparativeExpressionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 8: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::andExpressionSempred(AndExpressionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 9: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::orExpressionSempred(OrExpressionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 10: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::impliesExpressionSempred(ImpliesExpressionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 11: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::expressionListSempred(ExpressionListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 12: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::putItemListSempred(PutItemListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 13: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::indexTypeListSempred(IndexTypeListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 14: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::variableIdentifierListSempred(VariableIdentifierListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 15: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool TuringParser::identifierListSempred(IdentifierListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 16: return precpred(_ctx, 1);

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
  "argumentList", "exponentialExpression", "pointerExpression", "postfixExpression", 
  "prefixExpression", "multiplicativeExpression", "multiplicativeOperator", 
  "additiveExpression", "comparativeExpression", "notExpression", "andExpression", 
  "orExpression", "impliesExpression", "expression", "expressionList", "declaration", 
  "statements", "assignmentStatement", "putStatement", "putItem", "putItemList", 
  "beginStatement", "resultStatement", "newStatement", "freeStatement", 
  "forkStatement", "typeDeclaration", "typeSpec", "indexType", "indexTypeList", 
  "stringType", "recordType", "recordField", "variableDeclaration", "variableIdentifierList", 
  "variableIdentifier", "arrayDeclaration", "identifierList", "literal", 
  "comment"
};

std::vector<std::string> TuringParser::_literalNames = {
  "", "'unqualified'", "'pervasive'", "'opaque'", "", "", "'end'", "'of'", 
  "'to'", "'type'", "'var'", "", "", "'class'", "'process'", "'for'", "'loop'", 
  "'exit'", "'if'", "'else'", "'elsif'", "'case'", "'assert'", "'begin'", 
  "'return'", "'result'", "'new'", "'free'", "'tag'", "'fork'", "'signal'", 
  "'wait'", "'pause'", "'quit'", "'unchecked'", "'checked'", "'export'", 
  "'import'", "'inherit'", "'implement'", "'by'", "'put'", "'int'", "'real'", 
  "'boolean'", "'string'", "'array'", "'set'", "'record'", "'union'", "'pointer'", 
  "'nat'", "'intn'", "'natn'", "'realn'", "'char'", "'deferred'", "'forward'", 
  "'body'", "'not'", "'^'", "':'", "'('", "')'", "'.'", "'..'", "','", "'#'", 
  "'->'", "':='", "'+'", "'-'", "'*'", "'/'", "'div'", "'mod'", "'rem'", 
  "'**'", "'<'", "'>'", "'='", "'<='", "'>='", "'not='", "'and'", "'or'", 
  "'=>'", "'in'", "'shr'", "'shl'", "'xor'", "'+='", "'-='", "'*='", "'/='", 
  "'div='", "'mod='", "'and='", "'or='", "'=>='", "'shr='", "'shl='", "'xor='"
};

std::vector<std::string> TuringParser::_symbolicNames = {
  "", "", "", "", "STRING_LITERAL", "INTEGER_LITERAL", "END", "OF", "TO", 
  "TYPE", "VAR", "PROCEDURE", "FUNCTION", "CLASS", "PROCESS", "FOR", "LOOP", 
  "EXIT", "IF", "ELSE", "ELSIF", "CASE", "ASSERT", "BEGIN", "RETURN", "RESULT", 
  "NEW", "FREE", "TAG", "FORK", "SIGNAL", "WAIT", "PAUSE", "QUIT", "UNCHECKED", 
  "CHECKED", "EXPORT", "IMPORT", "INHERIT", "IMPLEMENT", "BY", "PUT", "INT", 
  "REAL", "BOOLEAN", "STRING", "ARRAY", "SET", "RECORD", "UNION", "POINTER", 
  "NAT", "INTN", "NATN", "REALN", "CHAR", "DEFERRED", "FORWARD", "BODY", 
  "NOT", "CARET", "COLON", "L_PAREN", "R_PAREN", "DOT", "RANGE", "COMMA", 
  "CHEAT", "DEREFERENCE", "ASSIGNMENT", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", 
  "DIV", "MOD", "REM", "EXP", "LESSTHAN", "GREATERTHAN", "EQUAL", "LESSTHANEQUAL", 
  "GREATERTHANEQUAL", "NOTEQUAL", "AND", "OR", "IMPLIES", "IN", "SHR", "SHL", 
  "XOR", "PLUSEQUALS", "MINUSEQUALS", "MULTIPLYEQUALS", "DIVIDEEQUALS", 
  "DIVEQUALS", "MODEQUALS", "ANDEQUALS", "OREQUALS", "IMPLIESEQUALS", "SHREQUALS", 
  "SHLEQUALS", "XOREQUALS", "IDENTIFIER", "WHITESPACE", "BLOCK_COMMENT", 
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
    0x3, 0x6c, 0x25b, 0x4, 0x2, 0x9, 0x2, 0x4, 0x3, 0x9, 0x3, 0x4, 0x4, 
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
    0x32, 0x4, 0x33, 0x9, 0x33, 0x4, 0x34, 0x9, 0x34, 0x4, 0x35, 0x9, 0x35, 
    0x4, 0x36, 0x9, 0x36, 0x4, 0x37, 0x9, 0x37, 0x4, 0x38, 0x9, 0x38, 0x4, 
    0x39, 0x9, 0x39, 0x4, 0x3a, 0x9, 0x3a, 0x4, 0x3b, 0x9, 0x3b, 0x4, 0x3c, 
    0x9, 0x3c, 0x4, 0x3d, 0x9, 0x3d, 0x4, 0x3e, 0x9, 0x3e, 0x3, 0x2, 0x6, 
    0x2, 0x7e, 0xa, 0x2, 0xd, 0x2, 0xe, 0x2, 0x7f, 0x3, 0x3, 0x3, 0x3, 0x5, 
    0x3, 0x84, 0xa, 0x3, 0x3, 0x4, 0x3, 0x4, 0x3, 0x4, 0x3, 0x4, 0x5, 0x4, 
    0x8a, 0xa, 0x4, 0x3, 0x4, 0x5, 0x4, 0x8d, 0xa, 0x4, 0x3, 0x5, 0x3, 0x5, 
    0x3, 0x5, 0x3, 0x5, 0x5, 0x5, 0x93, 0xa, 0x5, 0x3, 0x5, 0x5, 0x5, 0x96, 
    0xa, 0x5, 0x3, 0x5, 0x5, 0x5, 0x99, 0xa, 0x5, 0x3, 0x5, 0x3, 0x5, 0x3, 
    0x5, 0x3, 0x6, 0x5, 0x6, 0x9f, 0xa, 0x6, 0x3, 0x6, 0x3, 0x6, 0x5, 0x6, 
    0xa3, 0xa, 0x6, 0x3, 0x7, 0x3, 0x7, 0x3, 0x7, 0x3, 0x7, 0x3, 0x7, 0x3, 
    0x8, 0x3, 0x8, 0x3, 0x9, 0x3, 0x9, 0x3, 0x9, 0x6, 0x9, 0xaf, 0xa, 0x9, 
    0xd, 0x9, 0xe, 0x9, 0xb0, 0x3, 0x9, 0x3, 0x9, 0x3, 0x9, 0x3, 0xa, 0x3, 
    0xa, 0x3, 0xa, 0x3, 0xa, 0x5, 0xa, 0xba, 0xa, 0xa, 0x3, 0xb, 0x3, 0xb, 
    0x3, 0xb, 0x3, 0xb, 0x3, 0xb, 0x3, 0xb, 0x3, 0xb, 0x5, 0xb, 0xc3, 0xa, 
    0xb, 0x3, 0xc, 0x5, 0xc, 0xc6, 0xa, 0xc, 0x3, 0xc, 0x3, 0xc, 0x3, 0xd, 
    0x3, 0xd, 0x3, 0xd, 0x3, 0xd, 0x3, 0xd, 0x3, 0xd, 0x7, 0xd, 0xd0, 0xa, 
    0xd, 0xc, 0xd, 0xe, 0xd, 0xd3, 0xb, 0xd, 0x3, 0xe, 0x3, 0xe, 0x3, 0xe, 
    0x3, 0xe, 0x3, 0xe, 0x3, 0xe, 0x3, 0xe, 0x5, 0xe, 0xdc, 0xa, 0xe, 0x3, 
    0xf, 0x3, 0xf, 0x5, 0xf, 0xe0, 0xa, 0xf, 0x3, 0xf, 0x3, 0xf, 0x3, 0xf, 
    0x5, 0xf, 0xe5, 0xa, 0xf, 0x3, 0xf, 0x3, 0xf, 0x3, 0xf, 0x3, 0xf, 0x5, 
    0xf, 0xeb, 0xa, 0xf, 0x3, 0x10, 0x3, 0x10, 0x3, 0x10, 0x3, 0x10, 0x5, 
    0x10, 0xf1, 0xa, 0x10, 0x3, 0x11, 0x3, 0x11, 0x3, 0x12, 0x5, 0x12, 0xf6, 
    0xa, 0x12, 0x3, 0x12, 0x3, 0x12, 0x3, 0x12, 0x3, 0x12, 0x3, 0x12, 0x5, 
    0x12, 0xfd, 0xa, 0x12, 0x3, 0x13, 0x3, 0x13, 0x3, 0x13, 0x3, 0x13, 0x3, 
    0x13, 0x3, 0x13, 0x7, 0x13, 0x105, 0xa, 0x13, 0xc, 0x13, 0xe, 0x13, 
    0x108, 0xb, 0x13, 0x3, 0x14, 0x6, 0x14, 0x10b, 0xa, 0x14, 0xd, 0x14, 
    0xe, 0x14, 0x10c, 0x3, 0x15, 0x3, 0x15, 0x5, 0x15, 0x111, 0xa, 0x15, 
    0x3, 0x16, 0x3, 0x16, 0x3, 0x16, 0x3, 0x16, 0x3, 0x16, 0x3, 0x16, 0x5, 
    0x16, 0x119, 0xa, 0x16, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 
    0x3, 0x17, 0x3, 0x17, 0x7, 0x17, 0x121, 0xa, 0x17, 0xc, 0x17, 0xe, 0x17, 
    0x124, 0xb, 0x17, 0x3, 0x18, 0x3, 0x18, 0x3, 0x18, 0x3, 0x18, 0x3, 0x19, 
    0x3, 0x19, 0x3, 0x19, 0x3, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x5, 
    0x1a, 0x131, 0xa, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x5, 0x1a, 
    0x136, 0xa, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 0x3, 0x1a, 
    0x3, 0x1a, 0x3, 0x1a, 0x7, 0x1a, 0x13f, 0xa, 0x1a, 0xc, 0x1a, 0xe, 0x1a, 
    0x142, 0xb, 0x1a, 0x3, 0x1b, 0x3, 0x1b, 0x3, 0x1b, 0x3, 0x1b, 0x3, 0x1b, 
    0x3, 0x1b, 0x3, 0x1b, 0x5, 0x1b, 0x14b, 0xa, 0x1b, 0x3, 0x1c, 0x3, 0x1c, 
    0x3, 0x1c, 0x5, 0x1c, 0x150, 0xa, 0x1c, 0x3, 0x1c, 0x3, 0x1c, 0x3, 0x1c, 
    0x3, 0x1c, 0x5, 0x1c, 0x156, 0xa, 0x1c, 0x7, 0x1c, 0x158, 0xa, 0x1c, 
    0xc, 0x1c, 0xe, 0x1c, 0x15b, 0xb, 0x1c, 0x3, 0x1d, 0x3, 0x1d, 0x3, 0x1e, 
    0x3, 0x1e, 0x3, 0x1e, 0x3, 0x1e, 0x3, 0x1e, 0x3, 0x1e, 0x7, 0x1e, 0x165, 
    0xa, 0x1e, 0xc, 0x1e, 0xe, 0x1e, 0x168, 0xb, 0x1e, 0x3, 0x1f, 0x3, 0x1f, 
    0x3, 0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x3, 
    0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x5, 0x1f, 0x177, 
    0xa, 0x1f, 0x3, 0x1f, 0x7, 0x1f, 0x17a, 0xa, 0x1f, 0xc, 0x1f, 0xe, 0x1f, 
    0x17d, 0xb, 0x1f, 0x3, 0x20, 0x3, 0x20, 0x3, 0x20, 0x5, 0x20, 0x182, 
    0xa, 0x20, 0x3, 0x21, 0x3, 0x21, 0x3, 0x21, 0x3, 0x21, 0x3, 0x21, 0x3, 
    0x21, 0x7, 0x21, 0x18a, 0xa, 0x21, 0xc, 0x21, 0xe, 0x21, 0x18d, 0xb, 
    0x21, 0x3, 0x22, 0x3, 0x22, 0x3, 0x22, 0x3, 0x22, 0x3, 0x22, 0x3, 0x22, 
    0x7, 0x22, 0x195, 0xa, 0x22, 0xc, 0x22, 0xe, 0x22, 0x198, 0xb, 0x22, 
    0x3, 0x23, 0x3, 0x23, 0x3, 0x23, 0x3, 0x23, 0x3, 0x23, 0x3, 0x23, 0x7, 
    0x23, 0x1a0, 0xa, 0x23, 0xc, 0x23, 0xe, 0x23, 0x1a3, 0xb, 0x23, 0x3, 
    0x24, 0x3, 0x24, 0x3, 0x25, 0x3, 0x25, 0x3, 0x25, 0x3, 0x25, 0x3, 0x25, 
    0x3, 0x25, 0x7, 0x25, 0x1ad, 0xa, 0x25, 0xc, 0x25, 0xe, 0x25, 0x1b0, 
    0xb, 0x25, 0x3, 0x26, 0x3, 0x26, 0x3, 0x26, 0x3, 0x26, 0x5, 0x26, 0x1b6, 
    0xa, 0x26, 0x3, 0x27, 0x3, 0x27, 0x3, 0x27, 0x3, 0x27, 0x3, 0x27, 0x3, 
    0x27, 0x3, 0x27, 0x3, 0x27, 0x3, 0x27, 0x3, 0x27, 0x5, 0x27, 0x1c2, 
    0xa, 0x27, 0x3, 0x28, 0x3, 0x28, 0x3, 0x28, 0x3, 0x28, 0x3, 0x29, 0x3, 
    0x29, 0x3, 0x29, 0x3, 0x2a, 0x3, 0x2a, 0x3, 0x2b, 0x3, 0x2b, 0x3, 0x2b, 
    0x3, 0x2b, 0x3, 0x2b, 0x3, 0x2b, 0x7, 0x2b, 0x1d3, 0xa, 0x2b, 0xc, 0x2b, 
    0xe, 0x2b, 0x1d6, 0xb, 0x2b, 0x3, 0x2c, 0x3, 0x2c, 0x3, 0x2c, 0x3, 0x2c, 
    0x3, 0x2d, 0x3, 0x2d, 0x3, 0x2d, 0x3, 0x2e, 0x3, 0x2e, 0x3, 0x2e, 0x5, 
    0x2e, 0x1e2, 0xa, 0x2e, 0x3, 0x2e, 0x3, 0x2e, 0x3, 0x2f, 0x3, 0x2f, 
    0x3, 0x2f, 0x5, 0x2f, 0x1e9, 0xa, 0x2f, 0x3, 0x2f, 0x3, 0x2f, 0x3, 0x30, 
    0x3, 0x30, 0x3, 0x30, 0x3, 0x31, 0x3, 0x31, 0x3, 0x31, 0x3, 0x31, 0x3, 
    0x31, 0x3, 0x32, 0x3, 0x32, 0x3, 0x32, 0x3, 0x32, 0x3, 0x32, 0x3, 0x32, 
    0x3, 0x32, 0x3, 0x32, 0x3, 0x32, 0x3, 0x32, 0x3, 0x32, 0x3, 0x32, 0x5, 
    0x32, 0x201, 0xa, 0x32, 0x3, 0x33, 0x3, 0x33, 0x3, 0x33, 0x3, 0x33, 
    0x3, 0x34, 0x3, 0x34, 0x3, 0x34, 0x3, 0x34, 0x3, 0x34, 0x3, 0x34, 0x7, 
    0x34, 0x20d, 0xa, 0x34, 0xc, 0x34, 0xe, 0x34, 0x210, 0xb, 0x34, 0x3, 
    0x35, 0x3, 0x35, 0x3, 0x35, 0x3, 0x35, 0x5, 0x35, 0x216, 0xa, 0x35, 
    0x3, 0x36, 0x3, 0x36, 0x6, 0x36, 0x21a, 0xa, 0x36, 0xd, 0x36, 0xe, 0x36, 
    0x21b, 0x3, 0x36, 0x3, 0x36, 0x3, 0x36, 0x3, 0x37, 0x3, 0x37, 0x3, 0x37, 
    0x3, 0x37, 0x3, 0x38, 0x3, 0x38, 0x3, 0x38, 0x3, 0x38, 0x3, 0x38, 0x3, 
    0x38, 0x5, 0x38, 0x22b, 0xa, 0x38, 0x3, 0x38, 0x3, 0x38, 0x3, 0x38, 
    0x3, 0x38, 0x3, 0x38, 0x5, 0x38, 0x232, 0xa, 0x38, 0x3, 0x39, 0x3, 0x39, 
    0x3, 0x39, 0x3, 0x39, 0x3, 0x39, 0x3, 0x39, 0x7, 0x39, 0x23a, 0xa, 0x39, 
    0xc, 0x39, 0xe, 0x39, 0x23d, 0xb, 0x39, 0x3, 0x3a, 0x3, 0x3a, 0x3, 0x3b, 
    0x3, 0x3b, 0x3, 0x3b, 0x3, 0x3b, 0x3, 0x3b, 0x3, 0x3b, 0x3, 0x3b, 0x3, 
    0x3b, 0x3, 0x3b, 0x5, 0x3b, 0x24a, 0xa, 0x3b, 0x3, 0x3c, 0x3, 0x3c, 
    0x3, 0x3c, 0x3, 0x3c, 0x3, 0x3c, 0x3, 0x3c, 0x7, 0x3c, 0x252, 0xa, 0x3c, 
    0xc, 0x3c, 0xe, 0x3c, 0x255, 0xb, 0x3c, 0x3, 0x3d, 0x3, 0x3d, 0x3, 0x3e, 
    0x3, 0x3e, 0x3, 0x3e, 0x2, 0x11, 0x18, 0x24, 0x2c, 0x32, 0x36, 0x3a, 
    0x3c, 0x40, 0x42, 0x44, 0x48, 0x54, 0x66, 0x70, 0x76, 0x3f, 0x2, 0x4, 
    0x6, 0x8, 0xa, 0xc, 0xe, 0x10, 0x12, 0x14, 0x16, 0x18, 0x1a, 0x1c, 0x1e, 
    0x20, 0x22, 0x24, 0x26, 0x28, 0x2a, 0x2c, 0x2e, 0x30, 0x32, 0x34, 0x36, 
    0x38, 0x3a, 0x3c, 0x3e, 0x40, 0x42, 0x44, 0x46, 0x48, 0x4a, 0x4c, 0x4e, 
    0x50, 0x52, 0x54, 0x56, 0x58, 0x5a, 0x5c, 0x5e, 0x60, 0x62, 0x64, 0x66, 
    0x68, 0x6a, 0x6c, 0x6e, 0x70, 0x72, 0x74, 0x76, 0x78, 0x7a, 0x2, 0xa, 
    0x4, 0x2, 0xd, 0xd, 0x10, 0x10, 0x3, 0x2, 0x3a, 0x3c, 0x4, 0x2, 0x3, 
    0x5, 0xc, 0xc, 0x4, 0x2, 0x4a, 0x4e, 0x5a, 0x5b, 0x4, 0x2, 0x48, 0x49, 
    0x5c, 0x5c, 0x5, 0x2, 0x47, 0x47, 0x5d, 0x61, 0x66, 0x67, 0x3, 0x2, 
    0x6, 0x7, 0x3, 0x2, 0x6b, 0x6c, 0x2, 0x276, 0x2, 0x7d, 0x3, 0x2, 0x2, 
    0x2, 0x4, 0x83, 0x3, 0x2, 0x2, 0x2, 0x6, 0x85, 0x3, 0x2, 0x2, 0x2, 0x8, 
    0x8e, 0x3, 0x2, 0x2, 0x2, 0xa, 0x9e, 0x3, 0x2, 0x2, 0x2, 0xc, 0xa4, 
    0x3, 0x2, 0x2, 0x2, 0xe, 0xa9, 0x3, 0x2, 0x2, 0x2, 0x10, 0xab, 0x3, 
    0x2, 0x2, 0x2, 0x12, 0xb9, 0x3, 0x2, 0x2, 0x2, 0x14, 0xc2, 0x3, 0x2, 
    0x2, 0x2, 0x16, 0xc5, 0x3, 0x2, 0x2, 0x2, 0x18, 0xc9, 0x3, 0x2, 0x2, 
    0x2, 0x1a, 0xdb, 0x3, 0x2, 0x2, 0x2, 0x1c, 0xea, 0x3, 0x2, 0x2, 0x2, 
    0x1e, 0xf0, 0x3, 0x2, 0x2, 0x2, 0x20, 0xf2, 0x3, 0x2, 0x2, 0x2, 0x22, 
    0xfc, 0x3, 0x2, 0x2, 0x2, 0x24, 0xfe, 0x3, 0x2, 0x2, 0x2, 0x26, 0x10a, 
    0x3, 0x2, 0x2, 0x2, 0x28, 0x110, 0x3, 0x2, 0x2, 0x2, 0x2a, 0x118, 0x3, 
    0x2, 0x2, 0x2, 0x2c, 0x11a, 0x3, 0x2, 0x2, 0x2, 0x2e, 0x125, 0x3, 0x2, 
    0x2, 0x2, 0x30, 0x129, 0x3, 0x2, 0x2, 0x2, 0x32, 0x130, 0x3, 0x2, 0x2, 
    0x2, 0x34, 0x14a, 0x3, 0x2, 0x2, 0x2, 0x36, 0x14f, 0x3, 0x2, 0x2, 0x2, 
    0x38, 0x15c, 0x3, 0x2, 0x2, 0x2, 0x3a, 0x15e, 0x3, 0x2, 0x2, 0x2, 0x3c, 
    0x169, 0x3, 0x2, 0x2, 0x2, 0x3e, 0x181, 0x3, 0x2, 0x2, 0x2, 0x40, 0x183, 
    0x3, 0x2, 0x2, 0x2, 0x42, 0x18e, 0x3, 0x2, 0x2, 0x2, 0x44, 0x199, 0x3, 
    0x2, 0x2, 0x2, 0x46, 0x1a4, 0x3, 0x2, 0x2, 0x2, 0x48, 0x1a6, 0x3, 0x2, 
    0x2, 0x2, 0x4a, 0x1b5, 0x3, 0x2, 0x2, 0x2, 0x4c, 0x1c1, 0x3, 0x2, 0x2, 
    0x2, 0x4e, 0x1c3, 0x3, 0x2, 0x2, 0x2, 0x50, 0x1c7, 0x3, 0x2, 0x2, 0x2, 
    0x52, 0x1ca, 0x3, 0x2, 0x2, 0x2, 0x54, 0x1cc, 0x3, 0x2, 0x2, 0x2, 0x56, 
    0x1d7, 0x3, 0x2, 0x2, 0x2, 0x58, 0x1db, 0x3, 0x2, 0x2, 0x2, 0x5a, 0x1de, 
    0x3, 0x2, 0x2, 0x2, 0x5c, 0x1e5, 0x3, 0x2, 0x2, 0x2, 0x5e, 0x1ec, 0x3, 
    0x2, 0x2, 0x2, 0x60, 0x1ef, 0x3, 0x2, 0x2, 0x2, 0x62, 0x200, 0x3, 0x2, 
    0x2, 0x2, 0x64, 0x202, 0x3, 0x2, 0x2, 0x2, 0x66, 0x206, 0x3, 0x2, 0x2, 
    0x2, 0x68, 0x211, 0x3, 0x2, 0x2, 0x2, 0x6a, 0x217, 0x3, 0x2, 0x2, 0x2, 
    0x6c, 0x220, 0x3, 0x2, 0x2, 0x2, 0x6e, 0x231, 0x3, 0x2, 0x2, 0x2, 0x70, 
    0x233, 0x3, 0x2, 0x2, 0x2, 0x72, 0x23e, 0x3, 0x2, 0x2, 0x2, 0x74, 0x249, 
    0x3, 0x2, 0x2, 0x2, 0x76, 0x24b, 0x3, 0x2, 0x2, 0x2, 0x78, 0x256, 0x3, 
    0x2, 0x2, 0x2, 0x7a, 0x258, 0x3, 0x2, 0x2, 0x2, 0x7c, 0x7e, 0x5, 0x4, 
    0x3, 0x2, 0x7d, 0x7c, 0x3, 0x2, 0x2, 0x2, 0x7e, 0x7f, 0x3, 0x2, 0x2, 
    0x2, 0x7f, 0x7d, 0x3, 0x2, 0x2, 0x2, 0x7f, 0x80, 0x3, 0x2, 0x2, 0x2, 
    0x80, 0x3, 0x3, 0x2, 0x2, 0x2, 0x81, 0x84, 0x5, 0x28, 0x15, 0x2, 0x82, 
    0x84, 0x5, 0x10, 0x9, 0x2, 0x83, 0x81, 0x3, 0x2, 0x2, 0x2, 0x83, 0x82, 
    0x3, 0x2, 0x2, 0x2, 0x84, 0x5, 0x3, 0x2, 0x2, 0x2, 0x85, 0x86, 0x9, 
    0x2, 0x2, 0x2, 0x86, 0x8c, 0x7, 0x69, 0x2, 0x2, 0x87, 0x89, 0x7, 0x40, 
    0x2, 0x2, 0x88, 0x8a, 0x5, 0x24, 0x13, 0x2, 0x89, 0x88, 0x3, 0x2, 0x2, 
    0x2, 0x89, 0x8a, 0x3, 0x2, 0x2, 0x2, 0x8a, 0x8b, 0x3, 0x2, 0x2, 0x2, 
    0x8b, 0x8d, 0x7, 0x41, 0x2, 0x2, 0x8c, 0x87, 0x3, 0x2, 0x2, 0x2, 0x8c, 
    0x8d, 0x3, 0x2, 0x2, 0x2, 0x8d, 0x7, 0x3, 0x2, 0x2, 0x2, 0x8e, 0x8f, 
    0x7, 0xe, 0x2, 0x2, 0x8f, 0x95, 0x7, 0x69, 0x2, 0x2, 0x90, 0x92, 0x7, 
    0x40, 0x2, 0x2, 0x91, 0x93, 0x5, 0x24, 0x13, 0x2, 0x92, 0x91, 0x3, 0x2, 
    0x2, 0x2, 0x92, 0x93, 0x3, 0x2, 0x2, 0x2, 0x93, 0x94, 0x3, 0x2, 0x2, 
    0x2, 0x94, 0x96, 0x7, 0x41, 0x2, 0x2, 0x95, 0x90, 0x3, 0x2, 0x2, 0x2, 
    0x95, 0x96, 0x3, 0x2, 0x2, 0x2, 0x96, 0x98, 0x3, 0x2, 0x2, 0x2, 0x97, 
    0x99, 0x7, 0x69, 0x2, 0x2, 0x98, 0x97, 0x3, 0x2, 0x2, 0x2, 0x98, 0x99, 
    0x3, 0x2, 0x2, 0x2, 0x99, 0x9a, 0x3, 0x2, 0x2, 0x2, 0x9a, 0x9b, 0x7, 
    0x3f, 0x2, 0x2, 0x9b, 0x9c, 0x5, 0x62, 0x32, 0x2, 0x9c, 0x9, 0x3, 0x2, 
    0x2, 0x2, 0x9d, 0x9f, 0x5, 0xe, 0x8, 0x2, 0x9e, 0x9d, 0x3, 0x2, 0x2, 
    0x2, 0x9e, 0x9f, 0x3, 0x2, 0x2, 0x2, 0x9f, 0xa2, 0x3, 0x2, 0x2, 0x2, 
    0xa0, 0xa3, 0x5, 0x8, 0x5, 0x2, 0xa1, 0xa3, 0x5, 0x6, 0x4, 0x2, 0xa2, 
    0xa0, 0x3, 0x2, 0x2, 0x2, 0xa2, 0xa1, 0x3, 0x2, 0x2, 0x2, 0xa3, 0xb, 
    0x3, 0x2, 0x2, 0x2, 0xa4, 0xa5, 0x5, 0xa, 0x6, 0x2, 0xa5, 0xa6, 0x5, 
    0x26, 0x14, 0x2, 0xa6, 0xa7, 0x7, 0x8, 0x2, 0x2, 0xa7, 0xa8, 0x7, 0x69, 
    0x2, 0x2, 0xa8, 0xd, 0x3, 0x2, 0x2, 0x2, 0xa9, 0xaa, 0x9, 0x3, 0x2, 
    0x2, 0xaa, 0xf, 0x3, 0x2, 0x2, 0x2, 0xab, 0xac, 0x7, 0xf, 0x2, 0x2, 
    0xac, 0xae, 0x7, 0x69, 0x2, 0x2, 0xad, 0xaf, 0x5, 0x12, 0xa, 0x2, 0xae, 
    0xad, 0x3, 0x2, 0x2, 0x2, 0xaf, 0xb0, 0x3, 0x2, 0x2, 0x2, 0xb0, 0xae, 
    0x3, 0x2, 0x2, 0x2, 0xb0, 0xb1, 0x3, 0x2, 0x2, 0x2, 0xb1, 0xb2, 0x3, 
    0x2, 0x2, 0x2, 0xb2, 0xb3, 0x7, 0x8, 0x2, 0x2, 0xb3, 0xb4, 0x7, 0x69, 
    0x2, 0x2, 0xb4, 0x11, 0x3, 0x2, 0x2, 0x2, 0xb5, 0xba, 0x5, 0x14, 0xb, 
    0x2, 0xb6, 0xba, 0x5, 0x1a, 0xe, 0x2, 0xb7, 0xba, 0x5, 0x1c, 0xf, 0x2, 
    0xb8, 0xba, 0x5, 0x28, 0x15, 0x2, 0xb9, 0xb5, 0x3, 0x2, 0x2, 0x2, 0xb9, 
    0xb6, 0x3, 0x2, 0x2, 0x2, 0xb9, 0xb7, 0x3, 0x2, 0x2, 0x2, 0xb9, 0xb8, 
    0x3, 0x2, 0x2, 0x2, 0xba, 0x13, 0x3, 0x2, 0x2, 0x2, 0xbb, 0xbc, 0x7, 
    0x26, 0x2, 0x2, 0xbc, 0xc3, 0x5, 0x18, 0xd, 0x2, 0xbd, 0xbe, 0x7, 0x26, 
    0x2, 0x2, 0xbe, 0xbf, 0x7, 0x40, 0x2, 0x2, 0xbf, 0xc0, 0x5, 0x18, 0xd, 
    0x2, 0xc0, 0xc1, 0x7, 0x41, 0x2, 0x2, 0xc1, 0xc3, 0x3, 0x2, 0x2, 0x2, 
    0xc2, 0xbb, 0x3, 0x2, 0x2, 0x2, 0xc2, 0xbd, 0x3, 0x2, 0x2, 0x2, 0xc3, 
    0x15, 0x3, 0x2, 0x2, 0x2, 0xc4, 0xc6, 0x5, 0x20, 0x11, 0x2, 0xc5, 0xc4, 
    0x3, 0x2, 0x2, 0x2, 0xc5, 0xc6, 0x3, 0x2, 0x2, 0x2, 0xc6, 0xc7, 0x3, 
    0x2, 0x2, 0x2, 0xc7, 0xc8, 0x7, 0x69, 0x2, 0x2, 0xc8, 0x17, 0x3, 0x2, 
    0x2, 0x2, 0xc9, 0xca, 0x8, 0xd, 0x1, 0x2, 0xca, 0xcb, 0x5, 0x16, 0xc, 
    0x2, 0xcb, 0xd1, 0x3, 0x2, 0x2, 0x2, 0xcc, 0xcd, 0xc, 0x3, 0x2, 0x2, 
    0xcd, 0xce, 0x7, 0x44, 0x2, 0x2, 0xce, 0xd0, 0x5, 0x16, 0xc, 0x2, 0xcf, 
    0xcc, 0x3, 0x2, 0x2, 0x2, 0xd0, 0xd3, 0x3, 0x2, 0x2, 0x2, 0xd1, 0xcf, 
    0x3, 0x2, 0x2, 0x2, 0xd1, 0xd2, 0x3, 0x2, 0x2, 0x2, 0xd2, 0x19, 0x3, 
    0x2, 0x2, 0x2, 0xd3, 0xd1, 0x3, 0x2, 0x2, 0x2, 0xd4, 0xd5, 0x7, 0x28, 
    0x2, 0x2, 0xd5, 0xdc, 0x5, 0x1e, 0x10, 0x2, 0xd6, 0xd7, 0x7, 0x28, 0x2, 
    0x2, 0xd7, 0xd8, 0x7, 0x40, 0x2, 0x2, 0xd8, 0xd9, 0x5, 0x1e, 0x10, 0x2, 
    0xd9, 0xda, 0x7, 0x41, 0x2, 0x2, 0xda, 0xdc, 0x3, 0x2, 0x2, 0x2, 0xdb, 
    0xd4, 0x3, 0x2, 0x2, 0x2, 0xdb, 0xd6, 0x3, 0x2, 0x2, 0x2, 0xdc, 0x1b, 
    0x3, 0x2, 0x2, 0x2, 0xdd, 0xdf, 0x7, 0x29, 0x2, 0x2, 0xde, 0xe0, 0x7, 
    0x2a, 0x2, 0x2, 0xdf, 0xde, 0x3, 0x2, 0x2, 0x2, 0xdf, 0xe0, 0x3, 0x2, 
    0x2, 0x2, 0xe0, 0xe1, 0x3, 0x2, 0x2, 0x2, 0xe1, 0xeb, 0x5, 0x1e, 0x10, 
    0x2, 0xe2, 0xe4, 0x7, 0x29, 0x2, 0x2, 0xe3, 0xe5, 0x7, 0x2a, 0x2, 0x2, 
    0xe4, 0xe3, 0x3, 0x2, 0x2, 0x2, 0xe4, 0xe5, 0x3, 0x2, 0x2, 0x2, 0xe5, 
    0xe6, 0x3, 0x2, 0x2, 0x2, 0xe6, 0xe7, 0x7, 0x40, 0x2, 0x2, 0xe7, 0xe8, 
    0x5, 0x1e, 0x10, 0x2, 0xe8, 0xe9, 0x7, 0x41, 0x2, 0x2, 0xe9, 0xeb, 0x3, 
    0x2, 0x2, 0x2, 0xea, 0xdd, 0x3, 0x2, 0x2, 0x2, 0xea, 0xe2, 0x3, 0x2, 
    0x2, 0x2, 0xeb, 0x1d, 0x3, 0x2, 0x2, 0x2, 0xec, 0xf1, 0x7, 0x69, 0x2, 
    0x2, 0xed, 0xee, 0x7, 0x69, 0x2, 0x2, 0xee, 0xef, 0x7, 0x59, 0x2, 0x2, 
    0xef, 0xf1, 0x7, 0x6, 0x2, 0x2, 0xf0, 0xec, 0x3, 0x2, 0x2, 0x2, 0xf0, 
    0xed, 0x3, 0x2, 0x2, 0x2, 0xf1, 0x1f, 0x3, 0x2, 0x2, 0x2, 0xf2, 0xf3, 
    0x9, 0x4, 0x2, 0x2, 0xf3, 0x21, 0x3, 0x2, 0x2, 0x2, 0xf4, 0xf6, 0x7, 
    0xc, 0x2, 0x2, 0xf5, 0xf4, 0x3, 0x2, 0x2, 0x2, 0xf5, 0xf6, 0x3, 0x2, 
    0x2, 0x2, 0xf6, 0xf7, 0x3, 0x2, 0x2, 0x2, 0xf7, 0xf8, 0x5, 0x76, 0x3c, 
    0x2, 0xf8, 0xf9, 0x7, 0x3f, 0x2, 0x2, 0xf9, 0xfa, 0x5, 0x62, 0x32, 0x2, 
    0xfa, 0xfd, 0x3, 0x2, 0x2, 0x2, 0xfb, 0xfd, 0x5, 0xa, 0x6, 0x2, 0xfc, 
    0xf5, 0x3, 0x2, 0x2, 0x2, 0xfc, 0xfb, 0x3, 0x2, 0x2, 0x2, 0xfd, 0x23, 
    0x3, 0x2, 0x2, 0x2, 0xfe, 0xff, 0x8, 0x13, 0x1, 0x2, 0xff, 0x100, 0x5, 
    0x22, 0x12, 0x2, 0x100, 0x106, 0x3, 0x2, 0x2, 0x2, 0x101, 0x102, 0xc, 
    0x3, 0x2, 0x2, 0x102, 0x103, 0x7, 0x44, 0x2, 0x2, 0x103, 0x105, 0x5, 
    0x22, 0x12, 0x2, 0x104, 0x101, 0x3, 0x2, 0x2, 0x2, 0x105, 0x108, 0x3, 
    0x2, 0x2, 0x2, 0x106, 0x104, 0x3, 0x2, 0x2, 0x2, 0x106, 0x107, 0x3, 
    0x2, 0x2, 0x2, 0x107, 0x25, 0x3, 0x2, 0x2, 0x2, 0x108, 0x106, 0x3, 0x2, 
    0x2, 0x2, 0x109, 0x10b, 0x5, 0x28, 0x15, 0x2, 0x10a, 0x109, 0x3, 0x2, 
    0x2, 0x2, 0x10b, 0x10c, 0x3, 0x2, 0x2, 0x2, 0x10c, 0x10a, 0x3, 0x2, 
    0x2, 0x2, 0x10c, 0x10d, 0x3, 0x2, 0x2, 0x2, 0x10d, 0x27, 0x3, 0x2, 0x2, 
    0x2, 0x10e, 0x111, 0x5, 0x4c, 0x27, 0x2, 0x10f, 0x111, 0x5, 0x4a, 0x26, 
    0x2, 0x110, 0x10e, 0x3, 0x2, 0x2, 0x2, 0x110, 0x10f, 0x3, 0x2, 0x2, 
    0x2, 0x111, 0x29, 0x3, 0x2, 0x2, 0x2, 0x112, 0x119, 0x5, 0x78, 0x3d, 
    0x2, 0x113, 0x119, 0x7, 0x69, 0x2, 0x2, 0x114, 0x115, 0x7, 0x40, 0x2, 
    0x2, 0x115, 0x116, 0x5, 0x46, 0x24, 0x2, 0x116, 0x117, 0x7, 0x41, 0x2, 
    0x2, 0x117, 0x119, 0x3, 0x2, 0x2, 0x2, 0x118, 0x112, 0x3, 0x2, 0x2, 
    0x2, 0x118, 0x113, 0x3, 0x2, 0x2, 0x2, 0x118, 0x114, 0x3, 0x2, 0x2, 
    0x2, 0x119, 0x2b, 0x3, 0x2, 0x2, 0x2, 0x11a, 0x11b, 0x8, 0x17, 0x1, 
    0x2, 0x11b, 0x11c, 0x5, 0x46, 0x24, 0x2, 0x11c, 0x122, 0x3, 0x2, 0x2, 
    0x2, 0x11d, 0x11e, 0xc, 0x3, 0x2, 0x2, 0x11e, 0x11f, 0x7, 0x44, 0x2, 
    0x2, 0x11f, 0x121, 0x5, 0x2c, 0x17, 0x4, 0x120, 0x11d, 0x3, 0x2, 0x2, 
    0x2, 0x121, 0x124, 0x3, 0x2, 0x2, 0x2, 0x122, 0x120, 0x3, 0x2, 0x2, 
    0x2, 0x122, 0x123, 0x3, 0x2, 0x2, 0x2, 0x123, 0x2d, 0x3, 0x2, 0x2, 0x2, 
    0x124, 0x122, 0x3, 0x2, 0x2, 0x2, 0x125, 0x126, 0x5, 0x2a, 0x16, 0x2, 
    0x126, 0x127, 0x7, 0x4f, 0x2, 0x2, 0x127, 0x128, 0x5, 0x2a, 0x16, 0x2, 
    0x128, 0x2f, 0x3, 0x2, 0x2, 0x2, 0x129, 0x12a, 0x7, 0x3e, 0x2, 0x2, 
    0x12a, 0x12b, 0x5, 0x2a, 0x16, 0x2, 0x12b, 0x31, 0x3, 0x2, 0x2, 0x2, 
    0x12c, 0x12d, 0x8, 0x1a, 0x1, 0x2, 0x12d, 0x131, 0x5, 0x2a, 0x16, 0x2, 
    0x12e, 0x131, 0x5, 0x30, 0x19, 0x2, 0x12f, 0x131, 0x5, 0x2e, 0x18, 0x2, 
    0x130, 0x12c, 0x3, 0x2, 0x2, 0x2, 0x130, 0x12e, 0x3, 0x2, 0x2, 0x2, 
    0x130, 0x12f, 0x3, 0x2, 0x2, 0x2, 0x131, 0x140, 0x3, 0x2, 0x2, 0x2, 
    0x132, 0x133, 0xc, 0x5, 0x2, 0x2, 0x133, 0x135, 0x7, 0x40, 0x2, 0x2, 
    0x134, 0x136, 0x5, 0x2c, 0x17, 0x2, 0x135, 0x134, 0x3, 0x2, 0x2, 0x2, 
    0x135, 0x136, 0x3, 0x2, 0x2, 0x2, 0x136, 0x137, 0x3, 0x2, 0x2, 0x2, 
    0x137, 0x13f, 0x7, 0x41, 0x2, 0x2, 0x138, 0x139, 0xc, 0x4, 0x2, 0x2, 
    0x139, 0x13a, 0x7, 0x42, 0x2, 0x2, 0x13a, 0x13f, 0x7, 0x69, 0x2, 0x2, 
    0x13b, 0x13c, 0xc, 0x3, 0x2, 0x2, 0x13c, 0x13d, 0x7, 0x46, 0x2, 0x2, 
    0x13d, 0x13f, 0x7, 0x69, 0x2, 0x2, 0x13e, 0x132, 0x3, 0x2, 0x2, 0x2, 
    0x13e, 0x138, 0x3, 0x2, 0x2, 0x2, 0x13e, 0x13b, 0x3, 0x2, 0x2, 0x2, 
    0x13f, 0x142, 0x3, 0x2, 0x2, 0x2, 0x140, 0x13e, 0x3, 0x2, 0x2, 0x2, 
    0x140, 0x141, 0x3, 0x2, 0x2, 0x2, 0x141, 0x33, 0x3, 0x2, 0x2, 0x2, 0x142, 
    0x140, 0x3, 0x2, 0x2, 0x2, 0x143, 0x14b, 0x5, 0x32, 0x1a, 0x2, 0x144, 
    0x145, 0x7, 0x48, 0x2, 0x2, 0x145, 0x14b, 0x5, 0x34, 0x1b, 0x2, 0x146, 
    0x147, 0x7, 0x49, 0x2, 0x2, 0x147, 0x14b, 0x5, 0x34, 0x1b, 0x2, 0x148, 
    0x149, 0x7, 0x45, 0x2, 0x2, 0x149, 0x14b, 0x5, 0x34, 0x1b, 0x2, 0x14a, 
    0x143, 0x3, 0x2, 0x2, 0x2, 0x14a, 0x144, 0x3, 0x2, 0x2, 0x2, 0x14a, 
    0x146, 0x3, 0x2, 0x2, 0x2, 0x14a, 0x148, 0x3, 0x2, 0x2, 0x2, 0x14b, 
    0x35, 0x3, 0x2, 0x2, 0x2, 0x14c, 0x14d, 0x8, 0x1c, 0x1, 0x2, 0x14d, 
    0x150, 0x5, 0x32, 0x1a, 0x2, 0x14e, 0x150, 0x5, 0x34, 0x1b, 0x2, 0x14f, 
    0x14c, 0x3, 0x2, 0x2, 0x2, 0x14f, 0x14e, 0x3, 0x2, 0x2, 0x2, 0x150, 
    0x159, 0x3, 0x2, 0x2, 0x2, 0x151, 0x152, 0xc, 0x3, 0x2, 0x2, 0x152, 
    0x155, 0x5, 0x38, 0x1d, 0x2, 0x153, 0x156, 0x5, 0x32, 0x1a, 0x2, 0x154, 
    0x156, 0x5, 0x34, 0x1b, 0x2, 0x155, 0x153, 0x3, 0x2, 0x2, 0x2, 0x155, 
    0x154, 0x3, 0x2, 0x2, 0x2, 0x156, 0x158, 0x3, 0x2, 0x2, 0x2, 0x157, 
    0x151, 0x3, 0x2, 0x2, 0x2, 0x158, 0x15b, 0x3, 0x2, 0x2, 0x2, 0x159, 
    0x157, 0x3, 0x2, 0x2, 0x2, 0x159, 0x15a, 0x3, 0x2, 0x2, 0x2, 0x15a, 
    0x37, 0x3, 0x2, 0x2, 0x2, 0x15b, 0x159, 0x3, 0x2, 0x2, 0x2, 0x15c, 0x15d, 
    0x9, 0x5, 0x2, 0x2, 0x15d, 0x39, 0x3, 0x2, 0x2, 0x2, 0x15e, 0x15f, 0x8, 
    0x1e, 0x1, 0x2, 0x15f, 0x160, 0x5, 0x36, 0x1c, 0x2, 0x160, 0x166, 0x3, 
    0x2, 0x2, 0x2, 0x161, 0x162, 0xc, 0x3, 0x2, 0x2, 0x162, 0x163, 0x9, 
    0x6, 0x2, 0x2, 0x163, 0x165, 0x5, 0x36, 0x1c, 0x2, 0x164, 0x161, 0x3, 
    0x2, 0x2, 0x2, 0x165, 0x168, 0x3, 0x2, 0x2, 0x2, 0x166, 0x164, 0x3, 
    0x2, 0x2, 0x2, 0x166, 0x167, 0x3, 0x2, 0x2, 0x2, 0x167, 0x3b, 0x3, 0x2, 
    0x2, 0x2, 0x168, 0x166, 0x3, 0x2, 0x2, 0x2, 0x169, 0x16a, 0x8, 0x1f, 
    0x1, 0x2, 0x16a, 0x16b, 0x5, 0x3a, 0x1e, 0x2, 0x16b, 0x17b, 0x3, 0x2, 
    0x2, 0x2, 0x16c, 0x176, 0xc, 0x3, 0x2, 0x2, 0x16d, 0x177, 0x7, 0x50, 
    0x2, 0x2, 0x16e, 0x177, 0x7, 0x51, 0x2, 0x2, 0x16f, 0x177, 0x7, 0x52, 
    0x2, 0x2, 0x170, 0x177, 0x7, 0x53, 0x2, 0x2, 0x171, 0x177, 0x7, 0x54, 
    0x2, 0x2, 0x172, 0x177, 0x7, 0x55, 0x2, 0x2, 0x173, 0x177, 0x7, 0x59, 
    0x2, 0x2, 0x174, 0x175, 0x7, 0x3d, 0x2, 0x2, 0x175, 0x177, 0x7, 0x59, 
    0x2, 0x2, 0x176, 0x16d, 0x3, 0x2, 0x2, 0x2, 0x176, 0x16e, 0x3, 0x2, 
    0x2, 0x2, 0x176, 0x16f, 0x3, 0x2, 0x2, 0x2, 0x176, 0x170, 0x3, 0x2, 
    0x2, 0x2, 0x176, 0x171, 0x3, 0x2, 0x2, 0x2, 0x176, 0x172, 0x3, 0x2, 
    0x2, 0x2, 0x176, 0x173, 0x3, 0x2, 0x2, 0x2, 0x176, 0x174, 0x3, 0x2, 
    0x2, 0x2, 0x177, 0x178, 0x3, 0x2, 0x2, 0x2, 0x178, 0x17a, 0x5, 0x3a, 
    0x1e, 0x2, 0x179, 0x16c, 0x3, 0x2, 0x2, 0x2, 0x17a, 0x17d, 0x3, 0x2, 
    0x2, 0x2, 0x17b, 0x179, 0x3, 0x2, 0x2, 0x2, 0x17b, 0x17c, 0x3, 0x2, 
    0x2, 0x2, 0x17c, 0x3d, 0x3, 0x2, 0x2, 0x2, 0x17d, 0x17b, 0x3, 0x2, 0x2, 
    0x2, 0x17e, 0x182, 0x5, 0x3c, 0x1f, 0x2, 0x17f, 0x180, 0x7, 0x3d, 0x2, 
    0x2, 0x180, 0x182, 0x5, 0x3e, 0x20, 0x2, 0x181, 0x17e, 0x3, 0x2, 0x2, 
    0x2, 0x181, 0x17f, 0x3, 0x2, 0x2, 0x2, 0x182, 0x3f, 0x3, 0x2, 0x2, 0x2, 
    0x183, 0x184, 0x8, 0x21, 0x1, 0x2, 0x184, 0x185, 0x5, 0x3e, 0x20, 0x2, 
    0x185, 0x18b, 0x3, 0x2, 0x2, 0x2, 0x186, 0x187, 0xc, 0x3, 0x2, 0x2, 
    0x187, 0x188, 0x7, 0x56, 0x2, 0x2, 0x188, 0x18a, 0x5, 0x3e, 0x20, 0x2, 
    0x189, 0x186, 0x3, 0x2, 0x2, 0x2, 0x18a, 0x18d, 0x3, 0x2, 0x2, 0x2, 
    0x18b, 0x189, 0x3, 0x2, 0x2, 0x2, 0x18b, 0x18c, 0x3, 0x2, 0x2, 0x2, 
    0x18c, 0x41, 0x3, 0x2, 0x2, 0x2, 0x18d, 0x18b, 0x3, 0x2, 0x2, 0x2, 0x18e, 
    0x18f, 0x8, 0x22, 0x1, 0x2, 0x18f, 0x190, 0x5, 0x40, 0x21, 0x2, 0x190, 
    0x196, 0x3, 0x2, 0x2, 0x2, 0x191, 0x192, 0xc, 0x3, 0x2, 0x2, 0x192, 
    0x193, 0x7, 0x57, 0x2, 0x2, 0x193, 0x195, 0x5, 0x40, 0x21, 0x2, 0x194, 
    0x191, 0x3, 0x2, 0x2, 0x2, 0x195, 0x198, 0x3, 0x2, 0x2, 0x2, 0x196, 
    0x194, 0x3, 0x2, 0x2, 0x2, 0x196, 0x197, 0x3, 0x2, 0x2, 0x2, 0x197, 
    0x43, 0x3, 0x2, 0x2, 0x2, 0x198, 0x196, 0x3, 0x2, 0x2, 0x2, 0x199, 0x19a, 
    0x8, 0x23, 0x1, 0x2, 0x19a, 0x19b, 0x5, 0x42, 0x22, 0x2, 0x19b, 0x1a1, 
    0x3, 0x2, 0x2, 0x2, 0x19c, 0x19d, 0xc, 0x3, 0x2, 0x2, 0x19d, 0x19e, 
    0x7, 0x58, 0x2, 0x2, 0x19e, 0x1a0, 0x5, 0x42, 0x22, 0x2, 0x19f, 0x19c, 
    0x3, 0x2, 0x2, 0x2, 0x1a0, 0x1a3, 0x3, 0x2, 0x2, 0x2, 0x1a1, 0x19f, 
    0x3, 0x2, 0x2, 0x2, 0x1a1, 0x1a2, 0x3, 0x2, 0x2, 0x2, 0x1a2, 0x45, 0x3, 
    0x2, 0x2, 0x2, 0x1a3, 0x1a1, 0x3, 0x2, 0x2, 0x2, 0x1a4, 0x1a5, 0x5, 
    0x44, 0x23, 0x2, 0x1a5, 0x47, 0x3, 0x2, 0x2, 0x2, 0x1a6, 0x1a7, 0x8, 
    0x25, 0x1, 0x2, 0x1a7, 0x1a8, 0x5, 0x46, 0x24, 0x2, 0x1a8, 0x1ae, 0x3, 
    0x2, 0x2, 0x2, 0x1a9, 0x1aa, 0xc, 0x3, 0x2, 0x2, 0x1aa, 0x1ab, 0x7, 
    0x44, 0x2, 0x2, 0x1ab, 0x1ad, 0x5, 0x46, 0x24, 0x2, 0x1ac, 0x1a9, 0x3, 
    0x2, 0x2, 0x2, 0x1ad, 0x1b0, 0x3, 0x2, 0x2, 0x2, 0x1ae, 0x1ac, 0x3, 
    0x2, 0x2, 0x2, 0x1ae, 0x1af, 0x3, 0x2, 0x2, 0x2, 0x1af, 0x49, 0x3, 0x2, 
    0x2, 0x2, 0x1b0, 0x1ae, 0x3, 0x2, 0x2, 0x2, 0x1b1, 0x1b6, 0x5, 0x60, 
    0x31, 0x2, 0x1b2, 0x1b6, 0x5, 0x6e, 0x38, 0x2, 0x1b3, 0x1b6, 0x5, 0x74, 
    0x3b, 0x2, 0x1b4, 0x1b6, 0x5, 0xc, 0x7, 0x2, 0x1b5, 0x1b1, 0x3, 0x2, 
    0x2, 0x2, 0x1b5, 0x1b2, 0x3, 0x2, 0x2, 0x2, 0x1b5, 0x1b3, 0x3, 0x2, 
    0x2, 0x2, 0x1b5, 0x1b4, 0x3, 0x2, 0x2, 0x2, 0x1b6, 0x4b, 0x3, 0x2, 0x2, 
    0x2, 0x1b7, 0x1c2, 0x5, 0x46, 0x24, 0x2, 0x1b8, 0x1c2, 0x5, 0x4e, 0x28, 
    0x2, 0x1b9, 0x1c2, 0x5, 0x50, 0x29, 0x2, 0x1ba, 0x1c2, 0x7, 0x13, 0x2, 
    0x2, 0x1bb, 0x1c2, 0x5, 0x56, 0x2c, 0x2, 0x1bc, 0x1c2, 0x7, 0x1a, 0x2, 
    0x2, 0x1bd, 0x1c2, 0x5, 0x58, 0x2d, 0x2, 0x1be, 0x1c2, 0x5, 0x5a, 0x2e, 
    0x2, 0x1bf, 0x1c2, 0x5, 0x5c, 0x2f, 0x2, 0x1c0, 0x1c2, 0x5, 0x5e, 0x30, 
    0x2, 0x1c1, 0x1b7, 0x3, 0x2, 0x2, 0x2, 0x1c1, 0x1b8, 0x3, 0x2, 0x2, 
    0x2, 0x1c1, 0x1b9, 0x3, 0x2, 0x2, 0x2, 0x1c1, 0x1ba, 0x3, 0x2, 0x2, 
    0x2, 0x1c1, 0x1bb, 0x3, 0x2, 0x2, 0x2, 0x1c1, 0x1bc, 0x3, 0x2, 0x2, 
    0x2, 0x1c1, 0x1bd, 0x3, 0x2, 0x2, 0x2, 0x1c1, 0x1be, 0x3, 0x2, 0x2, 
    0x2, 0x1c1, 0x1bf, 0x3, 0x2, 0x2, 0x2, 0x1c1, 0x1c0, 0x3, 0x2, 0x2, 
    0x2, 0x1c2, 0x4d, 0x3, 0x2, 0x2, 0x2, 0x1c3, 0x1c4, 0x5, 0x32, 0x1a, 
    0x2, 0x1c4, 0x1c5, 0x9, 0x7, 0x2, 0x2, 0x1c5, 0x1c6, 0x5, 0x46, 0x24, 
    0x2, 0x1c6, 0x4f, 0x3, 0x2, 0x2, 0x2, 0x1c7, 0x1c8, 0x7, 0x2b, 0x2, 
    0x2, 0x1c8, 0x1c9, 0x5, 0x54, 0x2b, 0x2, 0x1c9, 0x51, 0x3, 0x2, 0x2, 
    0x2, 0x1ca, 0x1cb, 0x5, 0x4c, 0x27, 0x2, 0x1cb, 0x53, 0x3, 0x2, 0x2, 
    0x2, 0x1cc, 0x1cd, 0x8, 0x2b, 0x1, 0x2, 0x1cd, 0x1ce, 0x5, 0x52, 0x2a, 
    0x2, 0x1ce, 0x1d4, 0x3, 0x2, 0x2, 0x2, 0x1cf, 0x1d0, 0xc, 0x3, 0x2, 
    0x2, 0x1d0, 0x1d1, 0x7, 0x44, 0x2, 0x2, 0x1d1, 0x1d3, 0x5, 0x52, 0x2a, 
    0x2, 0x1d2, 0x1cf, 0x3, 0x2, 0x2, 0x2, 0x1d3, 0x1d6, 0x3, 0x2, 0x2, 
    0x2, 0x1d4, 0x1d2, 0x3, 0x2, 0x2, 0x2, 0x1d4, 0x1d5, 0x3, 0x2, 0x2, 
    0x2, 0x1d5, 0x55, 0x3, 0x2, 0x2, 0x2, 0x1d6, 0x1d4, 0x3, 0x2, 0x2, 0x2, 
    0x1d7, 0x1d8, 0x7, 0x19, 0x2, 0x2, 0x1d8, 0x1d9, 0x5, 0x28, 0x15, 0x2, 
    0x1d9, 0x1da, 0x7, 0x8, 0x2, 0x2, 0x1da, 0x57, 0x3, 0x2, 0x2, 0x2, 0x1db, 
    0x1dc, 0x7, 0x1b, 0x2, 0x2, 0x1dc, 0x1dd, 0x5, 0x46, 0x24, 0x2, 0x1dd, 
    0x59, 0x3, 0x2, 0x2, 0x2, 0x1de, 0x1e1, 0x7, 0x1c, 0x2, 0x2, 0x1df, 
    0x1e0, 0x7, 0x69, 0x2, 0x2, 0x1e0, 0x1e2, 0x7, 0x44, 0x2, 0x2, 0x1e1, 
    0x1df, 0x3, 0x2, 0x2, 0x2, 0x1e1, 0x1e2, 0x3, 0x2, 0x2, 0x2, 0x1e2, 
    0x1e3, 0x3, 0x2, 0x2, 0x2, 0x1e3, 0x1e4, 0x7, 0x69, 0x2, 0x2, 0x1e4, 
    0x5b, 0x3, 0x2, 0x2, 0x2, 0x1e5, 0x1e8, 0x7, 0x1d, 0x2, 0x2, 0x1e6, 
    0x1e7, 0x7, 0x69, 0x2, 0x2, 0x1e7, 0x1e9, 0x7, 0x44, 0x2, 0x2, 0x1e8, 
    0x1e6, 0x3, 0x2, 0x2, 0x2, 0x1e8, 0x1e9, 0x3, 0x2, 0x2, 0x2, 0x1e9, 
    0x1ea, 0x3, 0x2, 0x2, 0x2, 0x1ea, 0x1eb, 0x7, 0x69, 0x2, 0x2, 0x1eb, 
    0x5d, 0x3, 0x2, 0x2, 0x2, 0x1ec, 0x1ed, 0x7, 0x1f, 0x2, 0x2, 0x1ed, 
    0x1ee, 0x5, 0x46, 0x24, 0x2, 0x1ee, 0x5f, 0x3, 0x2, 0x2, 0x2, 0x1ef, 
    0x1f0, 0x7, 0xb, 0x2, 0x2, 0x1f0, 0x1f1, 0x7, 0x69, 0x2, 0x2, 0x1f1, 
    0x1f2, 0x7, 0x3f, 0x2, 0x2, 0x1f2, 0x1f3, 0x5, 0x62, 0x32, 0x2, 0x1f3, 
    0x61, 0x3, 0x2, 0x2, 0x2, 0x1f4, 0x201, 0x7, 0x2c, 0x2, 0x2, 0x1f5, 
    0x201, 0x7, 0x2d, 0x2, 0x2, 0x1f6, 0x201, 0x7, 0x2e, 0x2, 0x2, 0x1f7, 
    0x201, 0x7, 0x35, 0x2, 0x2, 0x1f8, 0x201, 0x7, 0x36, 0x2, 0x2, 0x1f9, 
    0x201, 0x7, 0x37, 0x2, 0x2, 0x1fa, 0x201, 0x7, 0x38, 0x2, 0x2, 0x1fb, 
    0x201, 0x7, 0x39, 0x2, 0x2, 0x1fc, 0x201, 0x5, 0x64, 0x33, 0x2, 0x1fd, 
    0x201, 0x5, 0x68, 0x35, 0x2, 0x1fe, 0x201, 0x5, 0x6a, 0x36, 0x2, 0x1ff, 
    0x201, 0x7, 0x69, 0x2, 0x2, 0x200, 0x1f4, 0x3, 0x2, 0x2, 0x2, 0x200, 
    0x1f5, 0x3, 0x2, 0x2, 0x2, 0x200, 0x1f6, 0x3, 0x2, 0x2, 0x2, 0x200, 
    0x1f7, 0x3, 0x2, 0x2, 0x2, 0x200, 0x1f8, 0x3, 0x2, 0x2, 0x2, 0x200, 
    0x1f9, 0x3, 0x2, 0x2, 0x2, 0x200, 0x1fa, 0x3, 0x2, 0x2, 0x2, 0x200, 
    0x1fb, 0x3, 0x2, 0x2, 0x2, 0x200, 0x1fc, 0x3, 0x2, 0x2, 0x2, 0x200, 
    0x1fd, 0x3, 0x2, 0x2, 0x2, 0x200, 0x1fe, 0x3, 0x2, 0x2, 0x2, 0x200, 
    0x1ff, 0x3, 0x2, 0x2, 0x2, 0x201, 0x63, 0x3, 0x2, 0x2, 0x2, 0x202, 0x203, 
    0x7, 0x7, 0x2, 0x2, 0x203, 0x204, 0x7, 0x43, 0x2, 0x2, 0x204, 0x205, 
    0x7, 0x7, 0x2, 0x2, 0x205, 0x65, 0x3, 0x2, 0x2, 0x2, 0x206, 0x207, 0x8, 
    0x34, 0x1, 0x2, 0x207, 0x208, 0x5, 0x64, 0x33, 0x2, 0x208, 0x20e, 0x3, 
    0x2, 0x2, 0x2, 0x209, 0x20a, 0xc, 0x3, 0x2, 0x2, 0x20a, 0x20b, 0x7, 
    0x44, 0x2, 0x2, 0x20b, 0x20d, 0x5, 0x64, 0x33, 0x2, 0x20c, 0x209, 0x3, 
    0x2, 0x2, 0x2, 0x20d, 0x210, 0x3, 0x2, 0x2, 0x2, 0x20e, 0x20c, 0x3, 
    0x2, 0x2, 0x2, 0x20e, 0x20f, 0x3, 0x2, 0x2, 0x2, 0x20f, 0x67, 0x3, 0x2, 
    0x2, 0x2, 0x210, 0x20e, 0x3, 0x2, 0x2, 0x2, 0x211, 0x215, 0x7, 0x2f, 
    0x2, 0x2, 0x212, 0x213, 0x7, 0x40, 0x2, 0x2, 0x213, 0x214, 0x7, 0x7, 
    0x2, 0x2, 0x214, 0x216, 0x7, 0x41, 0x2, 0x2, 0x215, 0x212, 0x3, 0x2, 
    0x2, 0x2, 0x215, 0x216, 0x3, 0x2, 0x2, 0x2, 0x216, 0x69, 0x3, 0x2, 0x2, 
    0x2, 0x217, 0x219, 0x7, 0x32, 0x2, 0x2, 0x218, 0x21a, 0x5, 0x6c, 0x37, 
    0x2, 0x219, 0x218, 0x3, 0x2, 0x2, 0x2, 0x21a, 0x21b, 0x3, 0x2, 0x2, 
    0x2, 0x21b, 0x219, 0x3, 0x2, 0x2, 0x2, 0x21b, 0x21c, 0x3, 0x2, 0x2, 
    0x2, 0x21c, 0x21d, 0x3, 0x2, 0x2, 0x2, 0x21d, 0x21e, 0x7, 0x8, 0x2, 
    0x2, 0x21e, 0x21f, 0x7, 0x32, 0x2, 0x2, 0x21f, 0x6b, 0x3, 0x2, 0x2, 
    0x2, 0x220, 0x221, 0x5, 0x76, 0x3c, 0x2, 0x221, 0x222, 0x7, 0x3f, 0x2, 
    0x2, 0x222, 0x223, 0x5, 0x62, 0x32, 0x2, 0x223, 0x6d, 0x3, 0x2, 0x2, 
    0x2, 0x224, 0x225, 0x7, 0xc, 0x2, 0x2, 0x225, 0x226, 0x5, 0x70, 0x39, 
    0x2, 0x226, 0x227, 0x7, 0x3f, 0x2, 0x2, 0x227, 0x22a, 0x5, 0x62, 0x32, 
    0x2, 0x228, 0x229, 0x7, 0x47, 0x2, 0x2, 0x229, 0x22b, 0x5, 0x46, 0x24, 
    0x2, 0x22a, 0x228, 0x3, 0x2, 0x2, 0x2, 0x22a, 0x22b, 0x3, 0x2, 0x2, 
    0x2, 0x22b, 0x232, 0x3, 0x2, 0x2, 0x2, 0x22c, 0x22d, 0x7, 0xc, 0x2, 
    0x2, 0x22d, 0x22e, 0x5, 0x70, 0x39, 0x2, 0x22e, 0x22f, 0x7, 0x47, 0x2, 
    0x2, 0x22f, 0x230, 0x5, 0x46, 0x24, 0x2, 0x230, 0x232, 0x3, 0x2, 0x2, 
    0x2, 0x231, 0x224, 0x3, 0x2, 0x2, 0x2, 0x231, 0x22c, 0x3, 0x2, 0x2, 
    0x2, 0x232, 0x6f, 0x3, 0x2, 0x2, 0x2, 0x233, 0x234, 0x8, 0x39, 0x1, 
    0x2, 0x234, 0x235, 0x5, 0x72, 0x3a, 0x2, 0x235, 0x23b, 0x3, 0x2, 0x2, 
    0x2, 0x236, 0x237, 0xc, 0x3, 0x2, 0x2, 0x237, 0x238, 0x7, 0x44, 0x2, 
    0x2, 0x238, 0x23a, 0x5, 0x72, 0x3a, 0x2, 0x239, 0x236, 0x3, 0x2, 0x2, 
    0x2, 0x23a, 0x23d, 0x3, 0x2, 0x2, 0x2, 0x23b, 0x239, 0x3, 0x2, 0x2, 
    0x2, 0x23b, 0x23c, 0x3, 0x2, 0x2, 0x2, 0x23c, 0x71, 0x3, 0x2, 0x2, 0x2, 
    0x23d, 0x23b, 0x3, 0x2, 0x2, 0x2, 0x23e, 0x23f, 0x7, 0x69, 0x2, 0x2, 
    0x23f, 0x73, 0x3, 0x2, 0x2, 0x2, 0x240, 0x241, 0x7, 0x30, 0x2, 0x2, 
    0x241, 0x242, 0x5, 0x64, 0x33, 0x2, 0x242, 0x243, 0x7, 0x9, 0x2, 0x2, 
    0x243, 0x244, 0x5, 0x62, 0x32, 0x2, 0x244, 0x24a, 0x3, 0x2, 0x2, 0x2, 
    0x245, 0x246, 0x7, 0x30, 0x2, 0x2, 0x246, 0x247, 0x5, 0x64, 0x33, 0x2, 
    0x247, 0x248, 0x7, 0x44, 0x2, 0x2, 0x248, 0x24a, 0x3, 0x2, 0x2, 0x2, 
    0x249, 0x240, 0x3, 0x2, 0x2, 0x2, 0x249, 0x245, 0x3, 0x2, 0x2, 0x2, 
    0x24a, 0x75, 0x3, 0x2, 0x2, 0x2, 0x24b, 0x24c, 0x8, 0x3c, 0x1, 0x2, 
    0x24c, 0x24d, 0x7, 0x69, 0x2, 0x2, 0x24d, 0x253, 0x3, 0x2, 0x2, 0x2, 
    0x24e, 0x24f, 0xc, 0x3, 0x2, 0x2, 0x24f, 0x250, 0x7, 0x44, 0x2, 0x2, 
    0x250, 0x252, 0x7, 0x69, 0x2, 0x2, 0x251, 0x24e, 0x3, 0x2, 0x2, 0x2, 
    0x252, 0x255, 0x3, 0x2, 0x2, 0x2, 0x253, 0x251, 0x3, 0x2, 0x2, 0x2, 
    0x253, 0x254, 0x3, 0x2, 0x2, 0x2, 0x254, 0x77, 0x3, 0x2, 0x2, 0x2, 0x255, 
    0x253, 0x3, 0x2, 0x2, 0x2, 0x256, 0x257, 0x9, 0x8, 0x2, 0x2, 0x257, 
    0x79, 0x3, 0x2, 0x2, 0x2, 0x258, 0x259, 0x9, 0x9, 0x2, 0x2, 0x259, 0x7b, 
    0x3, 0x2, 0x2, 0x2, 0x3a, 0x7f, 0x83, 0x89, 0x8c, 0x92, 0x95, 0x98, 
    0x9e, 0xa2, 0xb0, 0xb9, 0xc2, 0xc5, 0xd1, 0xdb, 0xdf, 0xe4, 0xea, 0xf0, 
    0xf5, 0xfc, 0x106, 0x10c, 0x110, 0x118, 0x122, 0x130, 0x135, 0x13e, 
    0x140, 0x14a, 0x14f, 0x155, 0x159, 0x166, 0x176, 0x17b, 0x181, 0x18b, 
    0x196, 0x1a1, 0x1ae, 0x1b5, 0x1c1, 0x1d4, 0x1e1, 0x1e8, 0x200, 0x20e, 
    0x215, 0x21b, 0x22a, 0x231, 0x23b, 0x249, 0x253, 
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
