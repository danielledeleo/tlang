// Code generated from tlang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // tlang

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasetlangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetlangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitTopLevel(ctx *TopLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitProcHeader(ctx *ProcHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitFuncHeader(ctx *FuncHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitSubprogramHeader(ctx *SubprogramHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitSubprogramDeclaration(ctx *SubprogramDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitSubprogramPrefix(ctx *SubprogramPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitClassBody(ctx *ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitExportListItem(ctx *ExportListItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitExportList(ctx *ExportListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitInheritStatement(ctx *InheritStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitImplementStatement(ctx *ImplementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitIdOrFileItem(ctx *IdOrFileItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitHowExport(ctx *HowExportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitParamDeclaration(ctx *ParamDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitParamDeclarationList(ctx *ParamDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitProcBody(ctx *ProcBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitStatementOrDeclaration(ctx *StatementOrDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitPutStatement(ctx *PutStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitPutItem(ctx *PutItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitPutItemList(ctx *PutItemListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitBeginStatement(ctx *BeginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitResultStatement(ctx *ResultStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitNewStatement(ctx *NewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitFreeStatement(ctx *FreeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitForkStatement(ctx *ForkStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitBasicType(ctx *BasicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitReferenceType(ctx *ReferenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitTypeSpec(ctx *TypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitIndexTypeList(ctx *IndexTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitRecordType(ctx *RecordTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitRecordField(ctx *RecordFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitVariableIdentifierList(ctx *VariableIdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitVariableIdentifier(ctx *VariableIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitArrayDeclaration(ctx *ArrayDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitInteger_literal(ctx *Integer_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetlangVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}
