// Code generated from tlang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // tlang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by tlangParser.
type tlangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tlangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by tlangParser#topLevel.
	VisitTopLevel(ctx *TopLevelContext) interface{}

	// Visit a parse tree produced by tlangParser#procHeader.
	VisitProcHeader(ctx *ProcHeaderContext) interface{}

	// Visit a parse tree produced by tlangParser#funcHeader.
	VisitFuncHeader(ctx *FuncHeaderContext) interface{}

	// Visit a parse tree produced by tlangParser#subprogramHeader.
	VisitSubprogramHeader(ctx *SubprogramHeaderContext) interface{}

	// Visit a parse tree produced by tlangParser#subprogramDeclaration.
	VisitSubprogramDeclaration(ctx *SubprogramDeclarationContext) interface{}

	// Visit a parse tree produced by tlangParser#subprogramPrefix.
	VisitSubprogramPrefix(ctx *SubprogramPrefixContext) interface{}

	// Visit a parse tree produced by tlangParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by tlangParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by tlangParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by tlangParser#exportListItem.
	VisitExportListItem(ctx *ExportListItemContext) interface{}

	// Visit a parse tree produced by tlangParser#exportList.
	VisitExportList(ctx *ExportListContext) interface{}

	// Visit a parse tree produced by tlangParser#inheritStatement.
	VisitInheritStatement(ctx *InheritStatementContext) interface{}

	// Visit a parse tree produced by tlangParser#implementStatement.
	VisitImplementStatement(ctx *ImplementStatementContext) interface{}

	// Visit a parse tree produced by tlangParser#idOrFileItem.
	VisitIdOrFileItem(ctx *IdOrFileItemContext) interface{}

	// Visit a parse tree produced by tlangParser#howExport.
	VisitHowExport(ctx *HowExportContext) interface{}

	// Visit a parse tree produced by tlangParser#paramDeclaration.
	VisitParamDeclaration(ctx *ParamDeclarationContext) interface{}

	// Visit a parse tree produced by tlangParser#paramDeclarationList.
	VisitParamDeclarationList(ctx *ParamDeclarationListContext) interface{}

	// Visit a parse tree produced by tlangParser#procBody.
	VisitProcBody(ctx *ProcBodyContext) interface{}

	// Visit a parse tree produced by tlangParser#statementOrDeclaration.
	VisitStatementOrDeclaration(ctx *StatementOrDeclarationContext) interface{}

	// Visit a parse tree produced by tlangParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by tlangParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by tlangParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by tlangParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by tlangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by tlangParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by tlangParser#putStatement.
	VisitPutStatement(ctx *PutStatementContext) interface{}

	// Visit a parse tree produced by tlangParser#putItem.
	VisitPutItem(ctx *PutItemContext) interface{}

	// Visit a parse tree produced by tlangParser#putItemList.
	VisitPutItemList(ctx *PutItemListContext) interface{}

	// Visit a parse tree produced by tlangParser#beginStatement.
	VisitBeginStatement(ctx *BeginStatementContext) interface{}

	// Visit a parse tree produced by tlangParser#resultStatement.
	VisitResultStatement(ctx *ResultStatementContext) interface{}

	// Visit a parse tree produced by tlangParser#newStatement.
	VisitNewStatement(ctx *NewStatementContext) interface{}

	// Visit a parse tree produced by tlangParser#freeStatement.
	VisitFreeStatement(ctx *FreeStatementContext) interface{}

	// Visit a parse tree produced by tlangParser#forkStatement.
	VisitForkStatement(ctx *ForkStatementContext) interface{}

	// Visit a parse tree produced by tlangParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by tlangParser#basicType.
	VisitBasicType(ctx *BasicTypeContext) interface{}

	// Visit a parse tree produced by tlangParser#referenceType.
	VisitReferenceType(ctx *ReferenceTypeContext) interface{}

	// Visit a parse tree produced by tlangParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by tlangParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) interface{}

	// Visit a parse tree produced by tlangParser#indexTypeList.
	VisitIndexTypeList(ctx *IndexTypeListContext) interface{}

	// Visit a parse tree produced by tlangParser#stringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by tlangParser#recordType.
	VisitRecordType(ctx *RecordTypeContext) interface{}

	// Visit a parse tree produced by tlangParser#recordField.
	VisitRecordField(ctx *RecordFieldContext) interface{}

	// Visit a parse tree produced by tlangParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by tlangParser#variableIdentifierList.
	VisitVariableIdentifierList(ctx *VariableIdentifierListContext) interface{}

	// Visit a parse tree produced by tlangParser#variableIdentifier.
	VisitVariableIdentifier(ctx *VariableIdentifierContext) interface{}

	// Visit a parse tree produced by tlangParser#arrayDeclaration.
	VisitArrayDeclaration(ctx *ArrayDeclarationContext) interface{}

	// Visit a parse tree produced by tlangParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by tlangParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by tlangParser#integer_literal.
	VisitInteger_literal(ctx *Integer_literalContext) interface{}

	// Visit a parse tree produced by tlangParser#comment.
	VisitComment(ctx *CommentContext) interface{}
}
