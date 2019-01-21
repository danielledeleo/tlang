// Code generated from Turing.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Turing

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by TuringParser.
type TuringVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TuringParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by TuringParser#topLevel.
	VisitTopLevel(ctx *TopLevelContext) interface{}

	// Visit a parse tree produced by TuringParser#procHeader.
	VisitProcHeader(ctx *ProcHeaderContext) interface{}

	// Visit a parse tree produced by TuringParser#funcHeader.
	VisitFuncHeader(ctx *FuncHeaderContext) interface{}

	// Visit a parse tree produced by TuringParser#subprogramHeader.
	VisitSubprogramHeader(ctx *SubprogramHeaderContext) interface{}

	// Visit a parse tree produced by TuringParser#subprogramDeclaration.
	VisitSubprogramDeclaration(ctx *SubprogramDeclarationContext) interface{}

	// Visit a parse tree produced by TuringParser#subprogramPrefix.
	VisitSubprogramPrefix(ctx *SubprogramPrefixContext) interface{}

	// Visit a parse tree produced by TuringParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by TuringParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by TuringParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by TuringParser#exportListItem.
	VisitExportListItem(ctx *ExportListItemContext) interface{}

	// Visit a parse tree produced by TuringParser#exportList.
	VisitExportList(ctx *ExportListContext) interface{}

	// Visit a parse tree produced by TuringParser#inheritStatement.
	VisitInheritStatement(ctx *InheritStatementContext) interface{}

	// Visit a parse tree produced by TuringParser#implementStatement.
	VisitImplementStatement(ctx *ImplementStatementContext) interface{}

	// Visit a parse tree produced by TuringParser#idOrFileItem.
	VisitIdOrFileItem(ctx *IdOrFileItemContext) interface{}

	// Visit a parse tree produced by TuringParser#howExport.
	VisitHowExport(ctx *HowExportContext) interface{}

	// Visit a parse tree produced by TuringParser#paramDeclaration.
	VisitParamDeclaration(ctx *ParamDeclarationContext) interface{}

	// Visit a parse tree produced by TuringParser#paramDeclarationList.
	VisitParamDeclarationList(ctx *ParamDeclarationListContext) interface{}

	// Visit a parse tree produced by TuringParser#procBody.
	VisitProcBody(ctx *ProcBodyContext) interface{}

	// Visit a parse tree produced by TuringParser#statementOrDeclaration.
	VisitStatementOrDeclaration(ctx *StatementOrDeclarationContext) interface{}

	// Visit a parse tree produced by TuringParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by TuringParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by TuringParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by TuringParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by TuringParser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by TuringParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by TuringParser#putStatement.
	VisitPutStatement(ctx *PutStatementContext) interface{}

	// Visit a parse tree produced by TuringParser#putItem.
	VisitPutItem(ctx *PutItemContext) interface{}

	// Visit a parse tree produced by TuringParser#putItemList.
	VisitPutItemList(ctx *PutItemListContext) interface{}

	// Visit a parse tree produced by TuringParser#beginStatement.
	VisitBeginStatement(ctx *BeginStatementContext) interface{}

	// Visit a parse tree produced by TuringParser#resultStatement.
	VisitResultStatement(ctx *ResultStatementContext) interface{}

	// Visit a parse tree produced by TuringParser#newStatement.
	VisitNewStatement(ctx *NewStatementContext) interface{}

	// Visit a parse tree produced by TuringParser#freeStatement.
	VisitFreeStatement(ctx *FreeStatementContext) interface{}

	// Visit a parse tree produced by TuringParser#forkStatement.
	VisitForkStatement(ctx *ForkStatementContext) interface{}

	// Visit a parse tree produced by TuringParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by TuringParser#basicType.
	VisitBasicType(ctx *BasicTypeContext) interface{}

	// Visit a parse tree produced by TuringParser#referenceType.
	VisitReferenceType(ctx *ReferenceTypeContext) interface{}

	// Visit a parse tree produced by TuringParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by TuringParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) interface{}

	// Visit a parse tree produced by TuringParser#indexTypeList.
	VisitIndexTypeList(ctx *IndexTypeListContext) interface{}

	// Visit a parse tree produced by TuringParser#stringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by TuringParser#recordType.
	VisitRecordType(ctx *RecordTypeContext) interface{}

	// Visit a parse tree produced by TuringParser#recordField.
	VisitRecordField(ctx *RecordFieldContext) interface{}

	// Visit a parse tree produced by TuringParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by TuringParser#variableIdentifierList.
	VisitVariableIdentifierList(ctx *VariableIdentifierListContext) interface{}

	// Visit a parse tree produced by TuringParser#variableIdentifier.
	VisitVariableIdentifier(ctx *VariableIdentifierContext) interface{}

	// Visit a parse tree produced by TuringParser#arrayDeclaration.
	VisitArrayDeclaration(ctx *ArrayDeclarationContext) interface{}

	// Visit a parse tree produced by TuringParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by TuringParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by TuringParser#comment.
	VisitComment(ctx *CommentContext) interface{}
}
