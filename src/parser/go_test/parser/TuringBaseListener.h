
// Generated from Turing.g4 by ANTLR 4.7.2

#pragma once


#include "antlr4-runtime.h"
#include "TuringListener.h"


/**
 * This class provides an empty implementation of TuringListener,
 * which can be extended to create a listener which only needs to handle a subset
 * of the available methods.
 */
class  TuringBaseListener : public TuringListener {
public:

  virtual void enterProgram(TuringParser::ProgramContext * /*ctx*/) override { }
  virtual void exitProgram(TuringParser::ProgramContext * /*ctx*/) override { }

  virtual void enterTopLevel(TuringParser::TopLevelContext * /*ctx*/) override { }
  virtual void exitTopLevel(TuringParser::TopLevelContext * /*ctx*/) override { }

  virtual void enterProcHeader(TuringParser::ProcHeaderContext * /*ctx*/) override { }
  virtual void exitProcHeader(TuringParser::ProcHeaderContext * /*ctx*/) override { }

  virtual void enterFuncHeader(TuringParser::FuncHeaderContext * /*ctx*/) override { }
  virtual void exitFuncHeader(TuringParser::FuncHeaderContext * /*ctx*/) override { }

  virtual void enterSubprogramHeader(TuringParser::SubprogramHeaderContext * /*ctx*/) override { }
  virtual void exitSubprogramHeader(TuringParser::SubprogramHeaderContext * /*ctx*/) override { }

  virtual void enterSubprogramDeclaration(TuringParser::SubprogramDeclarationContext * /*ctx*/) override { }
  virtual void exitSubprogramDeclaration(TuringParser::SubprogramDeclarationContext * /*ctx*/) override { }

  virtual void enterSubprogramPrefix(TuringParser::SubprogramPrefixContext * /*ctx*/) override { }
  virtual void exitSubprogramPrefix(TuringParser::SubprogramPrefixContext * /*ctx*/) override { }

  virtual void enterClassDeclaration(TuringParser::ClassDeclarationContext * /*ctx*/) override { }
  virtual void exitClassDeclaration(TuringParser::ClassDeclarationContext * /*ctx*/) override { }

  virtual void enterClassBody(TuringParser::ClassBodyContext * /*ctx*/) override { }
  virtual void exitClassBody(TuringParser::ClassBodyContext * /*ctx*/) override { }

  virtual void enterExportStatement(TuringParser::ExportStatementContext * /*ctx*/) override { }
  virtual void exitExportStatement(TuringParser::ExportStatementContext * /*ctx*/) override { }

  virtual void enterExportListItem(TuringParser::ExportListItemContext * /*ctx*/) override { }
  virtual void exitExportListItem(TuringParser::ExportListItemContext * /*ctx*/) override { }

  virtual void enterExportList(TuringParser::ExportListContext * /*ctx*/) override { }
  virtual void exitExportList(TuringParser::ExportListContext * /*ctx*/) override { }

  virtual void enterInheritStatement(TuringParser::InheritStatementContext * /*ctx*/) override { }
  virtual void exitInheritStatement(TuringParser::InheritStatementContext * /*ctx*/) override { }

  virtual void enterImplementStatement(TuringParser::ImplementStatementContext * /*ctx*/) override { }
  virtual void exitImplementStatement(TuringParser::ImplementStatementContext * /*ctx*/) override { }

  virtual void enterIdOrFileItem(TuringParser::IdOrFileItemContext * /*ctx*/) override { }
  virtual void exitIdOrFileItem(TuringParser::IdOrFileItemContext * /*ctx*/) override { }

  virtual void enterHowExport(TuringParser::HowExportContext * /*ctx*/) override { }
  virtual void exitHowExport(TuringParser::HowExportContext * /*ctx*/) override { }

  virtual void enterParamDeclaration(TuringParser::ParamDeclarationContext * /*ctx*/) override { }
  virtual void exitParamDeclaration(TuringParser::ParamDeclarationContext * /*ctx*/) override { }

  virtual void enterParamDeclarationList(TuringParser::ParamDeclarationListContext * /*ctx*/) override { }
  virtual void exitParamDeclarationList(TuringParser::ParamDeclarationListContext * /*ctx*/) override { }

  virtual void enterProcBody(TuringParser::ProcBodyContext * /*ctx*/) override { }
  virtual void exitProcBody(TuringParser::ProcBodyContext * /*ctx*/) override { }

  virtual void enterStatementOrDeclaration(TuringParser::StatementOrDeclarationContext * /*ctx*/) override { }
  virtual void exitStatementOrDeclaration(TuringParser::StatementOrDeclarationContext * /*ctx*/) override { }

  virtual void enterPrimaryExpression(TuringParser::PrimaryExpressionContext * /*ctx*/) override { }
  virtual void exitPrimaryExpression(TuringParser::PrimaryExpressionContext * /*ctx*/) override { }

  virtual void enterArgumentList(TuringParser::ArgumentListContext * /*ctx*/) override { }
  virtual void exitArgumentList(TuringParser::ArgumentListContext * /*ctx*/) override { }

  virtual void enterExponentialExpression(TuringParser::ExponentialExpressionContext * /*ctx*/) override { }
  virtual void exitExponentialExpression(TuringParser::ExponentialExpressionContext * /*ctx*/) override { }

  virtual void enterPointerExpression(TuringParser::PointerExpressionContext * /*ctx*/) override { }
  virtual void exitPointerExpression(TuringParser::PointerExpressionContext * /*ctx*/) override { }

  virtual void enterPostfixExpression(TuringParser::PostfixExpressionContext * /*ctx*/) override { }
  virtual void exitPostfixExpression(TuringParser::PostfixExpressionContext * /*ctx*/) override { }

  virtual void enterPrefixExpression(TuringParser::PrefixExpressionContext * /*ctx*/) override { }
  virtual void exitPrefixExpression(TuringParser::PrefixExpressionContext * /*ctx*/) override { }

  virtual void enterMultiplicativeExpression(TuringParser::MultiplicativeExpressionContext * /*ctx*/) override { }
  virtual void exitMultiplicativeExpression(TuringParser::MultiplicativeExpressionContext * /*ctx*/) override { }

  virtual void enterMultiplicativeOperator(TuringParser::MultiplicativeOperatorContext * /*ctx*/) override { }
  virtual void exitMultiplicativeOperator(TuringParser::MultiplicativeOperatorContext * /*ctx*/) override { }

  virtual void enterAdditiveExpression(TuringParser::AdditiveExpressionContext * /*ctx*/) override { }
  virtual void exitAdditiveExpression(TuringParser::AdditiveExpressionContext * /*ctx*/) override { }

  virtual void enterComparativeExpression(TuringParser::ComparativeExpressionContext * /*ctx*/) override { }
  virtual void exitComparativeExpression(TuringParser::ComparativeExpressionContext * /*ctx*/) override { }

  virtual void enterNotExpression(TuringParser::NotExpressionContext * /*ctx*/) override { }
  virtual void exitNotExpression(TuringParser::NotExpressionContext * /*ctx*/) override { }

  virtual void enterAndExpression(TuringParser::AndExpressionContext * /*ctx*/) override { }
  virtual void exitAndExpression(TuringParser::AndExpressionContext * /*ctx*/) override { }

  virtual void enterOrExpression(TuringParser::OrExpressionContext * /*ctx*/) override { }
  virtual void exitOrExpression(TuringParser::OrExpressionContext * /*ctx*/) override { }

  virtual void enterImpliesExpression(TuringParser::ImpliesExpressionContext * /*ctx*/) override { }
  virtual void exitImpliesExpression(TuringParser::ImpliesExpressionContext * /*ctx*/) override { }

  virtual void enterExpression(TuringParser::ExpressionContext * /*ctx*/) override { }
  virtual void exitExpression(TuringParser::ExpressionContext * /*ctx*/) override { }

  virtual void enterExpressionList(TuringParser::ExpressionListContext * /*ctx*/) override { }
  virtual void exitExpressionList(TuringParser::ExpressionListContext * /*ctx*/) override { }

  virtual void enterDeclaration(TuringParser::DeclarationContext * /*ctx*/) override { }
  virtual void exitDeclaration(TuringParser::DeclarationContext * /*ctx*/) override { }

  virtual void enterStatements(TuringParser::StatementsContext * /*ctx*/) override { }
  virtual void exitStatements(TuringParser::StatementsContext * /*ctx*/) override { }

  virtual void enterAssignmentStatement(TuringParser::AssignmentStatementContext * /*ctx*/) override { }
  virtual void exitAssignmentStatement(TuringParser::AssignmentStatementContext * /*ctx*/) override { }

  virtual void enterPutStatement(TuringParser::PutStatementContext * /*ctx*/) override { }
  virtual void exitPutStatement(TuringParser::PutStatementContext * /*ctx*/) override { }

  virtual void enterPutItem(TuringParser::PutItemContext * /*ctx*/) override { }
  virtual void exitPutItem(TuringParser::PutItemContext * /*ctx*/) override { }

  virtual void enterPutItemList(TuringParser::PutItemListContext * /*ctx*/) override { }
  virtual void exitPutItemList(TuringParser::PutItemListContext * /*ctx*/) override { }

  virtual void enterBeginStatement(TuringParser::BeginStatementContext * /*ctx*/) override { }
  virtual void exitBeginStatement(TuringParser::BeginStatementContext * /*ctx*/) override { }

  virtual void enterResultStatement(TuringParser::ResultStatementContext * /*ctx*/) override { }
  virtual void exitResultStatement(TuringParser::ResultStatementContext * /*ctx*/) override { }

  virtual void enterNewStatement(TuringParser::NewStatementContext * /*ctx*/) override { }
  virtual void exitNewStatement(TuringParser::NewStatementContext * /*ctx*/) override { }

  virtual void enterFreeStatement(TuringParser::FreeStatementContext * /*ctx*/) override { }
  virtual void exitFreeStatement(TuringParser::FreeStatementContext * /*ctx*/) override { }

  virtual void enterForkStatement(TuringParser::ForkStatementContext * /*ctx*/) override { }
  virtual void exitForkStatement(TuringParser::ForkStatementContext * /*ctx*/) override { }

  virtual void enterTypeDeclaration(TuringParser::TypeDeclarationContext * /*ctx*/) override { }
  virtual void exitTypeDeclaration(TuringParser::TypeDeclarationContext * /*ctx*/) override { }

  virtual void enterTypeSpec(TuringParser::TypeSpecContext * /*ctx*/) override { }
  virtual void exitTypeSpec(TuringParser::TypeSpecContext * /*ctx*/) override { }

  virtual void enterIndexType(TuringParser::IndexTypeContext * /*ctx*/) override { }
  virtual void exitIndexType(TuringParser::IndexTypeContext * /*ctx*/) override { }

  virtual void enterIndexTypeList(TuringParser::IndexTypeListContext * /*ctx*/) override { }
  virtual void exitIndexTypeList(TuringParser::IndexTypeListContext * /*ctx*/) override { }

  virtual void enterStringType(TuringParser::StringTypeContext * /*ctx*/) override { }
  virtual void exitStringType(TuringParser::StringTypeContext * /*ctx*/) override { }

  virtual void enterRecordType(TuringParser::RecordTypeContext * /*ctx*/) override { }
  virtual void exitRecordType(TuringParser::RecordTypeContext * /*ctx*/) override { }

  virtual void enterRecordField(TuringParser::RecordFieldContext * /*ctx*/) override { }
  virtual void exitRecordField(TuringParser::RecordFieldContext * /*ctx*/) override { }

  virtual void enterVariableDeclaration(TuringParser::VariableDeclarationContext * /*ctx*/) override { }
  virtual void exitVariableDeclaration(TuringParser::VariableDeclarationContext * /*ctx*/) override { }

  virtual void enterVariableIdentifierList(TuringParser::VariableIdentifierListContext * /*ctx*/) override { }
  virtual void exitVariableIdentifierList(TuringParser::VariableIdentifierListContext * /*ctx*/) override { }

  virtual void enterVariableIdentifier(TuringParser::VariableIdentifierContext * /*ctx*/) override { }
  virtual void exitVariableIdentifier(TuringParser::VariableIdentifierContext * /*ctx*/) override { }

  virtual void enterArrayDeclaration(TuringParser::ArrayDeclarationContext * /*ctx*/) override { }
  virtual void exitArrayDeclaration(TuringParser::ArrayDeclarationContext * /*ctx*/) override { }

  virtual void enterIdentifierList(TuringParser::IdentifierListContext * /*ctx*/) override { }
  virtual void exitIdentifierList(TuringParser::IdentifierListContext * /*ctx*/) override { }

  virtual void enterLiteral(TuringParser::LiteralContext * /*ctx*/) override { }
  virtual void exitLiteral(TuringParser::LiteralContext * /*ctx*/) override { }

  virtual void enterComment(TuringParser::CommentContext * /*ctx*/) override { }
  virtual void exitComment(TuringParser::CommentContext * /*ctx*/) override { }


  virtual void enterEveryRule(antlr4::ParserRuleContext * /*ctx*/) override { }
  virtual void exitEveryRule(antlr4::ParserRuleContext * /*ctx*/) override { }
  virtual void visitTerminal(antlr4::tree::TerminalNode * /*node*/) override { }
  virtual void visitErrorNode(antlr4::tree::ErrorNode * /*node*/) override { }

};

