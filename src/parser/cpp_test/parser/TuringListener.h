
// Generated from Turing.g4 by ANTLR 4.7.2

#pragma once


#include "antlr4-runtime.h"
#include "TuringParser.h"


/**
 * This interface defines an abstract listener for a parse tree produced by TuringParser.
 */
class  TuringListener : public antlr4::tree::ParseTreeListener {
public:

  virtual void enterProgram(TuringParser::ProgramContext *ctx) = 0;
  virtual void exitProgram(TuringParser::ProgramContext *ctx) = 0;

  virtual void enterTopLevel(TuringParser::TopLevelContext *ctx) = 0;
  virtual void exitTopLevel(TuringParser::TopLevelContext *ctx) = 0;

  virtual void enterProcHeader(TuringParser::ProcHeaderContext *ctx) = 0;
  virtual void exitProcHeader(TuringParser::ProcHeaderContext *ctx) = 0;

  virtual void enterFuncHeader(TuringParser::FuncHeaderContext *ctx) = 0;
  virtual void exitFuncHeader(TuringParser::FuncHeaderContext *ctx) = 0;

  virtual void enterSubprogramHeader(TuringParser::SubprogramHeaderContext *ctx) = 0;
  virtual void exitSubprogramHeader(TuringParser::SubprogramHeaderContext *ctx) = 0;

  virtual void enterSubprogramDeclaration(TuringParser::SubprogramDeclarationContext *ctx) = 0;
  virtual void exitSubprogramDeclaration(TuringParser::SubprogramDeclarationContext *ctx) = 0;

  virtual void enterSubprogramPrefix(TuringParser::SubprogramPrefixContext *ctx) = 0;
  virtual void exitSubprogramPrefix(TuringParser::SubprogramPrefixContext *ctx) = 0;

  virtual void enterClassDeclaration(TuringParser::ClassDeclarationContext *ctx) = 0;
  virtual void exitClassDeclaration(TuringParser::ClassDeclarationContext *ctx) = 0;

  virtual void enterClassBody(TuringParser::ClassBodyContext *ctx) = 0;
  virtual void exitClassBody(TuringParser::ClassBodyContext *ctx) = 0;

  virtual void enterExportStatement(TuringParser::ExportStatementContext *ctx) = 0;
  virtual void exitExportStatement(TuringParser::ExportStatementContext *ctx) = 0;

  virtual void enterExportListItem(TuringParser::ExportListItemContext *ctx) = 0;
  virtual void exitExportListItem(TuringParser::ExportListItemContext *ctx) = 0;

  virtual void enterExportList(TuringParser::ExportListContext *ctx) = 0;
  virtual void exitExportList(TuringParser::ExportListContext *ctx) = 0;

  virtual void enterInheritStatement(TuringParser::InheritStatementContext *ctx) = 0;
  virtual void exitInheritStatement(TuringParser::InheritStatementContext *ctx) = 0;

  virtual void enterImplementStatement(TuringParser::ImplementStatementContext *ctx) = 0;
  virtual void exitImplementStatement(TuringParser::ImplementStatementContext *ctx) = 0;

  virtual void enterIdOrFileItem(TuringParser::IdOrFileItemContext *ctx) = 0;
  virtual void exitIdOrFileItem(TuringParser::IdOrFileItemContext *ctx) = 0;

  virtual void enterHowExport(TuringParser::HowExportContext *ctx) = 0;
  virtual void exitHowExport(TuringParser::HowExportContext *ctx) = 0;

  virtual void enterParamDeclaration(TuringParser::ParamDeclarationContext *ctx) = 0;
  virtual void exitParamDeclaration(TuringParser::ParamDeclarationContext *ctx) = 0;

  virtual void enterParamDeclarationList(TuringParser::ParamDeclarationListContext *ctx) = 0;
  virtual void exitParamDeclarationList(TuringParser::ParamDeclarationListContext *ctx) = 0;

  virtual void enterProcBody(TuringParser::ProcBodyContext *ctx) = 0;
  virtual void exitProcBody(TuringParser::ProcBodyContext *ctx) = 0;

  virtual void enterStatementOrDeclaration(TuringParser::StatementOrDeclarationContext *ctx) = 0;
  virtual void exitStatementOrDeclaration(TuringParser::StatementOrDeclarationContext *ctx) = 0;

  virtual void enterPrimaryExpression(TuringParser::PrimaryExpressionContext *ctx) = 0;
  virtual void exitPrimaryExpression(TuringParser::PrimaryExpressionContext *ctx) = 0;

  virtual void enterExpression(TuringParser::ExpressionContext *ctx) = 0;
  virtual void exitExpression(TuringParser::ExpressionContext *ctx) = 0;

  virtual void enterExpressionList(TuringParser::ExpressionListContext *ctx) = 0;
  virtual void exitExpressionList(TuringParser::ExpressionListContext *ctx) = 0;

  virtual void enterDeclaration(TuringParser::DeclarationContext *ctx) = 0;
  virtual void exitDeclaration(TuringParser::DeclarationContext *ctx) = 0;

  virtual void enterStatements(TuringParser::StatementsContext *ctx) = 0;
  virtual void exitStatements(TuringParser::StatementsContext *ctx) = 0;

  virtual void enterAssignmentOperator(TuringParser::AssignmentOperatorContext *ctx) = 0;
  virtual void exitAssignmentOperator(TuringParser::AssignmentOperatorContext *ctx) = 0;

  virtual void enterPutStatement(TuringParser::PutStatementContext *ctx) = 0;
  virtual void exitPutStatement(TuringParser::PutStatementContext *ctx) = 0;

  virtual void enterPutItem(TuringParser::PutItemContext *ctx) = 0;
  virtual void exitPutItem(TuringParser::PutItemContext *ctx) = 0;

  virtual void enterPutItemList(TuringParser::PutItemListContext *ctx) = 0;
  virtual void exitPutItemList(TuringParser::PutItemListContext *ctx) = 0;

  virtual void enterBeginStatement(TuringParser::BeginStatementContext *ctx) = 0;
  virtual void exitBeginStatement(TuringParser::BeginStatementContext *ctx) = 0;

  virtual void enterResultStatement(TuringParser::ResultStatementContext *ctx) = 0;
  virtual void exitResultStatement(TuringParser::ResultStatementContext *ctx) = 0;

  virtual void enterNewStatement(TuringParser::NewStatementContext *ctx) = 0;
  virtual void exitNewStatement(TuringParser::NewStatementContext *ctx) = 0;

  virtual void enterFreeStatement(TuringParser::FreeStatementContext *ctx) = 0;
  virtual void exitFreeStatement(TuringParser::FreeStatementContext *ctx) = 0;

  virtual void enterForkStatement(TuringParser::ForkStatementContext *ctx) = 0;
  virtual void exitForkStatement(TuringParser::ForkStatementContext *ctx) = 0;

  virtual void enterTypeDeclaration(TuringParser::TypeDeclarationContext *ctx) = 0;
  virtual void exitTypeDeclaration(TuringParser::TypeDeclarationContext *ctx) = 0;

  virtual void enterBasicType(TuringParser::BasicTypeContext *ctx) = 0;
  virtual void exitBasicType(TuringParser::BasicTypeContext *ctx) = 0;

  virtual void enterReferenceType(TuringParser::ReferenceTypeContext *ctx) = 0;
  virtual void exitReferenceType(TuringParser::ReferenceTypeContext *ctx) = 0;

  virtual void enterTypeSpec(TuringParser::TypeSpecContext *ctx) = 0;
  virtual void exitTypeSpec(TuringParser::TypeSpecContext *ctx) = 0;

  virtual void enterIndexType(TuringParser::IndexTypeContext *ctx) = 0;
  virtual void exitIndexType(TuringParser::IndexTypeContext *ctx) = 0;

  virtual void enterIndexTypeList(TuringParser::IndexTypeListContext *ctx) = 0;
  virtual void exitIndexTypeList(TuringParser::IndexTypeListContext *ctx) = 0;

  virtual void enterStringType(TuringParser::StringTypeContext *ctx) = 0;
  virtual void exitStringType(TuringParser::StringTypeContext *ctx) = 0;

  virtual void enterRecordType(TuringParser::RecordTypeContext *ctx) = 0;
  virtual void exitRecordType(TuringParser::RecordTypeContext *ctx) = 0;

  virtual void enterRecordField(TuringParser::RecordFieldContext *ctx) = 0;
  virtual void exitRecordField(TuringParser::RecordFieldContext *ctx) = 0;

  virtual void enterVariableDeclaration(TuringParser::VariableDeclarationContext *ctx) = 0;
  virtual void exitVariableDeclaration(TuringParser::VariableDeclarationContext *ctx) = 0;

  virtual void enterVariableIdentifierList(TuringParser::VariableIdentifierListContext *ctx) = 0;
  virtual void exitVariableIdentifierList(TuringParser::VariableIdentifierListContext *ctx) = 0;

  virtual void enterVariableIdentifier(TuringParser::VariableIdentifierContext *ctx) = 0;
  virtual void exitVariableIdentifier(TuringParser::VariableIdentifierContext *ctx) = 0;

  virtual void enterArrayDeclaration(TuringParser::ArrayDeclarationContext *ctx) = 0;
  virtual void exitArrayDeclaration(TuringParser::ArrayDeclarationContext *ctx) = 0;

  virtual void enterIdentifierList(TuringParser::IdentifierListContext *ctx) = 0;
  virtual void exitIdentifierList(TuringParser::IdentifierListContext *ctx) = 0;

  virtual void enterLiteral(TuringParser::LiteralContext *ctx) = 0;
  virtual void exitLiteral(TuringParser::LiteralContext *ctx) = 0;

  virtual void enterComment(TuringParser::CommentContext *ctx) = 0;
  virtual void exitComment(TuringParser::CommentContext *ctx) = 0;


};

