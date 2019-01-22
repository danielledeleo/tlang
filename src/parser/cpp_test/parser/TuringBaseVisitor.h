
// Generated from Turing.g4 by ANTLR 4.7.2

#pragma once


#include "antlr4-runtime.h"
#include "TuringVisitor.h"


/**
 * This class provides an empty implementation of TuringVisitor, which can be
 * extended to create a visitor which only needs to handle a subset of the available methods.
 */
class  TuringBaseVisitor : public TuringVisitor {
public:

  virtual antlrcpp::Any visitProgram(TuringParser::ProgramContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitTopLevel(TuringParser::TopLevelContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitProcHeader(TuringParser::ProcHeaderContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitFuncHeader(TuringParser::FuncHeaderContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitSubprogramHeader(TuringParser::SubprogramHeaderContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitSubprogramDeclaration(TuringParser::SubprogramDeclarationContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitSubprogramPrefix(TuringParser::SubprogramPrefixContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitClassDeclaration(TuringParser::ClassDeclarationContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitClassBody(TuringParser::ClassBodyContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitExportStatement(TuringParser::ExportStatementContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitExportListItem(TuringParser::ExportListItemContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitExportList(TuringParser::ExportListContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitInheritStatement(TuringParser::InheritStatementContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitImplementStatement(TuringParser::ImplementStatementContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitIdOrFileItem(TuringParser::IdOrFileItemContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitHowExport(TuringParser::HowExportContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitParamDeclaration(TuringParser::ParamDeclarationContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitParamDeclarationList(TuringParser::ParamDeclarationListContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitProcBody(TuringParser::ProcBodyContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitStatementOrDeclaration(TuringParser::StatementOrDeclarationContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitPrimaryExpression(TuringParser::PrimaryExpressionContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitExpression(TuringParser::ExpressionContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitExpressionList(TuringParser::ExpressionListContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitDeclaration(TuringParser::DeclarationContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitStatements(TuringParser::StatementsContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitAssignmentOperator(TuringParser::AssignmentOperatorContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitPutStatement(TuringParser::PutStatementContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitPutItem(TuringParser::PutItemContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitPutItemList(TuringParser::PutItemListContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitBeginStatement(TuringParser::BeginStatementContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitResultStatement(TuringParser::ResultStatementContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitNewStatement(TuringParser::NewStatementContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitFreeStatement(TuringParser::FreeStatementContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitForkStatement(TuringParser::ForkStatementContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitTypeDeclaration(TuringParser::TypeDeclarationContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitBasicType(TuringParser::BasicTypeContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitReferenceType(TuringParser::ReferenceTypeContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitTypeSpec(TuringParser::TypeSpecContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitIndexType(TuringParser::IndexTypeContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitIndexTypeList(TuringParser::IndexTypeListContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitStringType(TuringParser::StringTypeContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitRecordType(TuringParser::RecordTypeContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitRecordField(TuringParser::RecordFieldContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitVariableDeclaration(TuringParser::VariableDeclarationContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitVariableIdentifierList(TuringParser::VariableIdentifierListContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitVariableIdentifier(TuringParser::VariableIdentifierContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitArrayDeclaration(TuringParser::ArrayDeclarationContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitIdentifierList(TuringParser::IdentifierListContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitLiteral(TuringParser::LiteralContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitComment(TuringParser::CommentContext *ctx) override {
    return visitChildren(ctx);
  }


};

