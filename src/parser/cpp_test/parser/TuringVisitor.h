
// Generated from Turing.g4 by ANTLR 4.7.2

#pragma once


#include "antlr4-runtime.h"
#include "TuringParser.h"



/**
 * This class defines an abstract visitor for a parse tree
 * produced by TuringParser.
 */
class  TuringVisitor : public antlr4::tree::AbstractParseTreeVisitor {
public:

  /**
   * Visit parse trees produced by TuringParser.
   */
    virtual antlrcpp::Any visitProgram(TuringParser::ProgramContext *context) = 0;

    virtual antlrcpp::Any visitTopLevel(TuringParser::TopLevelContext *context) = 0;

    virtual antlrcpp::Any visitProcHeader(TuringParser::ProcHeaderContext *context) = 0;

    virtual antlrcpp::Any visitFuncHeader(TuringParser::FuncHeaderContext *context) = 0;

    virtual antlrcpp::Any visitSubprogramHeader(TuringParser::SubprogramHeaderContext *context) = 0;

    virtual antlrcpp::Any visitSubprogramDeclaration(TuringParser::SubprogramDeclarationContext *context) = 0;

    virtual antlrcpp::Any visitSubprogramPrefix(TuringParser::SubprogramPrefixContext *context) = 0;

    virtual antlrcpp::Any visitClassDeclaration(TuringParser::ClassDeclarationContext *context) = 0;

    virtual antlrcpp::Any visitClassBody(TuringParser::ClassBodyContext *context) = 0;

    virtual antlrcpp::Any visitExportStatement(TuringParser::ExportStatementContext *context) = 0;

    virtual antlrcpp::Any visitExportListItem(TuringParser::ExportListItemContext *context) = 0;

    virtual antlrcpp::Any visitExportList(TuringParser::ExportListContext *context) = 0;

    virtual antlrcpp::Any visitInheritStatement(TuringParser::InheritStatementContext *context) = 0;

    virtual antlrcpp::Any visitImplementStatement(TuringParser::ImplementStatementContext *context) = 0;

    virtual antlrcpp::Any visitIdOrFileItem(TuringParser::IdOrFileItemContext *context) = 0;

    virtual antlrcpp::Any visitHowExport(TuringParser::HowExportContext *context) = 0;

    virtual antlrcpp::Any visitParamDeclaration(TuringParser::ParamDeclarationContext *context) = 0;

    virtual antlrcpp::Any visitParamDeclarationList(TuringParser::ParamDeclarationListContext *context) = 0;

    virtual antlrcpp::Any visitProcBody(TuringParser::ProcBodyContext *context) = 0;

    virtual antlrcpp::Any visitStatementOrDeclaration(TuringParser::StatementOrDeclarationContext *context) = 0;

    virtual antlrcpp::Any visitPrimaryExpression(TuringParser::PrimaryExpressionContext *context) = 0;

    virtual antlrcpp::Any visitArgumentList(TuringParser::ArgumentListContext *context) = 0;

    virtual antlrcpp::Any visitExponentialExpression(TuringParser::ExponentialExpressionContext *context) = 0;

    virtual antlrcpp::Any visitPointerExpression(TuringParser::PointerExpressionContext *context) = 0;

    virtual antlrcpp::Any visitPostfixExpression(TuringParser::PostfixExpressionContext *context) = 0;

    virtual antlrcpp::Any visitPrefixExpression(TuringParser::PrefixExpressionContext *context) = 0;

    virtual antlrcpp::Any visitMultiplicativeExpression(TuringParser::MultiplicativeExpressionContext *context) = 0;

    virtual antlrcpp::Any visitMultiplicativeOperator(TuringParser::MultiplicativeOperatorContext *context) = 0;

    virtual antlrcpp::Any visitAdditiveExpression(TuringParser::AdditiveExpressionContext *context) = 0;

    virtual antlrcpp::Any visitAdditiveOperator(TuringParser::AdditiveOperatorContext *context) = 0;

    virtual antlrcpp::Any visitComparativeExpression(TuringParser::ComparativeExpressionContext *context) = 0;

    virtual antlrcpp::Any visitNotExpression(TuringParser::NotExpressionContext *context) = 0;

    virtual antlrcpp::Any visitAndExpression(TuringParser::AndExpressionContext *context) = 0;

    virtual antlrcpp::Any visitOrExpression(TuringParser::OrExpressionContext *context) = 0;

    virtual antlrcpp::Any visitImpliesExpression(TuringParser::ImpliesExpressionContext *context) = 0;

    virtual antlrcpp::Any visitExpression(TuringParser::ExpressionContext *context) = 0;

    virtual antlrcpp::Any visitExpressionList(TuringParser::ExpressionListContext *context) = 0;

    virtual antlrcpp::Any visitDeclaration(TuringParser::DeclarationContext *context) = 0;

    virtual antlrcpp::Any visitStatements(TuringParser::StatementsContext *context) = 0;

    virtual antlrcpp::Any visitAssignmentStatement(TuringParser::AssignmentStatementContext *context) = 0;

    virtual antlrcpp::Any visitPutStatement(TuringParser::PutStatementContext *context) = 0;

    virtual antlrcpp::Any visitPutItem(TuringParser::PutItemContext *context) = 0;

    virtual antlrcpp::Any visitPutItemList(TuringParser::PutItemListContext *context) = 0;

    virtual antlrcpp::Any visitBeginStatement(TuringParser::BeginStatementContext *context) = 0;

    virtual antlrcpp::Any visitResultStatement(TuringParser::ResultStatementContext *context) = 0;

    virtual antlrcpp::Any visitNewStatement(TuringParser::NewStatementContext *context) = 0;

    virtual antlrcpp::Any visitFreeStatement(TuringParser::FreeStatementContext *context) = 0;

    virtual antlrcpp::Any visitForkStatement(TuringParser::ForkStatementContext *context) = 0;

    virtual antlrcpp::Any visitTypeDeclaration(TuringParser::TypeDeclarationContext *context) = 0;

    virtual antlrcpp::Any visitBasicType(TuringParser::BasicTypeContext *context) = 0;

    virtual antlrcpp::Any visitReferenceType(TuringParser::ReferenceTypeContext *context) = 0;

    virtual antlrcpp::Any visitTypeSpec(TuringParser::TypeSpecContext *context) = 0;

    virtual antlrcpp::Any visitIndexType(TuringParser::IndexTypeContext *context) = 0;

    virtual antlrcpp::Any visitIndexTypeList(TuringParser::IndexTypeListContext *context) = 0;

    virtual antlrcpp::Any visitStringType(TuringParser::StringTypeContext *context) = 0;

    virtual antlrcpp::Any visitRecordType(TuringParser::RecordTypeContext *context) = 0;

    virtual antlrcpp::Any visitRecordField(TuringParser::RecordFieldContext *context) = 0;

    virtual antlrcpp::Any visitVariableDeclaration(TuringParser::VariableDeclarationContext *context) = 0;

    virtual antlrcpp::Any visitVariableIdentifierList(TuringParser::VariableIdentifierListContext *context) = 0;

    virtual antlrcpp::Any visitVariableIdentifier(TuringParser::VariableIdentifierContext *context) = 0;

    virtual antlrcpp::Any visitArrayDeclaration(TuringParser::ArrayDeclarationContext *context) = 0;

    virtual antlrcpp::Any visitIdentifierList(TuringParser::IdentifierListContext *context) = 0;

    virtual antlrcpp::Any visitLiteral(TuringParser::LiteralContext *context) = 0;

    virtual antlrcpp::Any visitComment(TuringParser::CommentContext *context) = 0;


};

