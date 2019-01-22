// Code generated from Turing.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Turing

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTuringVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTuringVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitTopLevel(ctx *TopLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitProcHeader(ctx *ProcHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitFuncHeader(ctx *FuncHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitSubprogramHeader(ctx *SubprogramHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitSubprogramDeclaration(ctx *SubprogramDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitSubprogramPrefix(ctx *SubprogramPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitClassBody(ctx *ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitExportListItem(ctx *ExportListItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitExportList(ctx *ExportListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitInheritStatement(ctx *InheritStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitImplementStatement(ctx *ImplementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitIdOrFileItem(ctx *IdOrFileItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitHowExport(ctx *HowExportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitParamDeclaration(ctx *ParamDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitParamDeclarationList(ctx *ParamDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitProcBody(ctx *ProcBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitStatementOrDeclaration(ctx *StatementOrDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitPutStatement(ctx *PutStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitPutItem(ctx *PutItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitPutItemList(ctx *PutItemListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitBeginStatement(ctx *BeginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitResultStatement(ctx *ResultStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitNewStatement(ctx *NewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitFreeStatement(ctx *FreeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitForkStatement(ctx *ForkStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitBasicType(ctx *BasicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitReferenceType(ctx *ReferenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitTypeSpec(ctx *TypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitIndexTypeList(ctx *IndexTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitRecordType(ctx *RecordTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitRecordField(ctx *RecordFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitVariableIdentifierList(ctx *VariableIdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitVariableIdentifier(ctx *VariableIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitArrayDeclaration(ctx *ArrayDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}
