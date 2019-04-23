// Code generated from tlang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // tlang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetlangListener is a complete listener for a parse tree produced by tlangParser.
type BasetlangListener struct{}

var _ tlangListener = &BasetlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasetlangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasetlangListener) ExitProgram(ctx *ProgramContext) {}

// EnterTopLevel is called when production topLevel is entered.
func (s *BasetlangListener) EnterTopLevel(ctx *TopLevelContext) {}

// ExitTopLevel is called when production topLevel is exited.
func (s *BasetlangListener) ExitTopLevel(ctx *TopLevelContext) {}

// EnterProcHeader is called when production procHeader is entered.
func (s *BasetlangListener) EnterProcHeader(ctx *ProcHeaderContext) {}

// ExitProcHeader is called when production procHeader is exited.
func (s *BasetlangListener) ExitProcHeader(ctx *ProcHeaderContext) {}

// EnterFuncHeader is called when production funcHeader is entered.
func (s *BasetlangListener) EnterFuncHeader(ctx *FuncHeaderContext) {}

// ExitFuncHeader is called when production funcHeader is exited.
func (s *BasetlangListener) ExitFuncHeader(ctx *FuncHeaderContext) {}

// EnterSubprogramHeader is called when production subprogramHeader is entered.
func (s *BasetlangListener) EnterSubprogramHeader(ctx *SubprogramHeaderContext) {}

// ExitSubprogramHeader is called when production subprogramHeader is exited.
func (s *BasetlangListener) ExitSubprogramHeader(ctx *SubprogramHeaderContext) {}

// EnterSubprogramDeclaration is called when production subprogramDeclaration is entered.
func (s *BasetlangListener) EnterSubprogramDeclaration(ctx *SubprogramDeclarationContext) {}

// ExitSubprogramDeclaration is called when production subprogramDeclaration is exited.
func (s *BasetlangListener) ExitSubprogramDeclaration(ctx *SubprogramDeclarationContext) {}

// EnterSubprogramPrefix is called when production subprogramPrefix is entered.
func (s *BasetlangListener) EnterSubprogramPrefix(ctx *SubprogramPrefixContext) {}

// ExitSubprogramPrefix is called when production subprogramPrefix is exited.
func (s *BasetlangListener) ExitSubprogramPrefix(ctx *SubprogramPrefixContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BasetlangListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BasetlangListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BasetlangListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BasetlangListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterExportStatement is called when production exportStatement is entered.
func (s *BasetlangListener) EnterExportStatement(ctx *ExportStatementContext) {}

// ExitExportStatement is called when production exportStatement is exited.
func (s *BasetlangListener) ExitExportStatement(ctx *ExportStatementContext) {}

// EnterExportListItem is called when production exportListItem is entered.
func (s *BasetlangListener) EnterExportListItem(ctx *ExportListItemContext) {}

// ExitExportListItem is called when production exportListItem is exited.
func (s *BasetlangListener) ExitExportListItem(ctx *ExportListItemContext) {}

// EnterExportList is called when production exportList is entered.
func (s *BasetlangListener) EnterExportList(ctx *ExportListContext) {}

// ExitExportList is called when production exportList is exited.
func (s *BasetlangListener) ExitExportList(ctx *ExportListContext) {}

// EnterInheritStatement is called when production inheritStatement is entered.
func (s *BasetlangListener) EnterInheritStatement(ctx *InheritStatementContext) {}

// ExitInheritStatement is called when production inheritStatement is exited.
func (s *BasetlangListener) ExitInheritStatement(ctx *InheritStatementContext) {}

// EnterImplementStatement is called when production implementStatement is entered.
func (s *BasetlangListener) EnterImplementStatement(ctx *ImplementStatementContext) {}

// ExitImplementStatement is called when production implementStatement is exited.
func (s *BasetlangListener) ExitImplementStatement(ctx *ImplementStatementContext) {}

// EnterIdOrFileItem is called when production idOrFileItem is entered.
func (s *BasetlangListener) EnterIdOrFileItem(ctx *IdOrFileItemContext) {}

// ExitIdOrFileItem is called when production idOrFileItem is exited.
func (s *BasetlangListener) ExitIdOrFileItem(ctx *IdOrFileItemContext) {}

// EnterHowExport is called when production howExport is entered.
func (s *BasetlangListener) EnterHowExport(ctx *HowExportContext) {}

// ExitHowExport is called when production howExport is exited.
func (s *BasetlangListener) ExitHowExport(ctx *HowExportContext) {}

// EnterParamDeclaration is called when production paramDeclaration is entered.
func (s *BasetlangListener) EnterParamDeclaration(ctx *ParamDeclarationContext) {}

// ExitParamDeclaration is called when production paramDeclaration is exited.
func (s *BasetlangListener) ExitParamDeclaration(ctx *ParamDeclarationContext) {}

// EnterParamDeclarationList is called when production paramDeclarationList is entered.
func (s *BasetlangListener) EnterParamDeclarationList(ctx *ParamDeclarationListContext) {}

// ExitParamDeclarationList is called when production paramDeclarationList is exited.
func (s *BasetlangListener) ExitParamDeclarationList(ctx *ParamDeclarationListContext) {}

// EnterProcBody is called when production procBody is entered.
func (s *BasetlangListener) EnterProcBody(ctx *ProcBodyContext) {}

// ExitProcBody is called when production procBody is exited.
func (s *BasetlangListener) ExitProcBody(ctx *ProcBodyContext) {}

// EnterStatementOrDeclaration is called when production statementOrDeclaration is entered.
func (s *BasetlangListener) EnterStatementOrDeclaration(ctx *StatementOrDeclarationContext) {}

// ExitStatementOrDeclaration is called when production statementOrDeclaration is exited.
func (s *BasetlangListener) ExitStatementOrDeclaration(ctx *StatementOrDeclarationContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BasetlangListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BasetlangListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasetlangListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasetlangListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasetlangListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasetlangListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasetlangListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasetlangListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasetlangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasetlangListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BasetlangListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BasetlangListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterPutStatement is called when production putStatement is entered.
func (s *BasetlangListener) EnterPutStatement(ctx *PutStatementContext) {}

// ExitPutStatement is called when production putStatement is exited.
func (s *BasetlangListener) ExitPutStatement(ctx *PutStatementContext) {}

// EnterPutItem is called when production putItem is entered.
func (s *BasetlangListener) EnterPutItem(ctx *PutItemContext) {}

// ExitPutItem is called when production putItem is exited.
func (s *BasetlangListener) ExitPutItem(ctx *PutItemContext) {}

// EnterPutItemList is called when production putItemList is entered.
func (s *BasetlangListener) EnterPutItemList(ctx *PutItemListContext) {}

// ExitPutItemList is called when production putItemList is exited.
func (s *BasetlangListener) ExitPutItemList(ctx *PutItemListContext) {}

// EnterBeginStatement is called when production beginStatement is entered.
func (s *BasetlangListener) EnterBeginStatement(ctx *BeginStatementContext) {}

// ExitBeginStatement is called when production beginStatement is exited.
func (s *BasetlangListener) ExitBeginStatement(ctx *BeginStatementContext) {}

// EnterResultStatement is called when production resultStatement is entered.
func (s *BasetlangListener) EnterResultStatement(ctx *ResultStatementContext) {}

// ExitResultStatement is called when production resultStatement is exited.
func (s *BasetlangListener) ExitResultStatement(ctx *ResultStatementContext) {}

// EnterNewStatement is called when production newStatement is entered.
func (s *BasetlangListener) EnterNewStatement(ctx *NewStatementContext) {}

// ExitNewStatement is called when production newStatement is exited.
func (s *BasetlangListener) ExitNewStatement(ctx *NewStatementContext) {}

// EnterFreeStatement is called when production freeStatement is entered.
func (s *BasetlangListener) EnterFreeStatement(ctx *FreeStatementContext) {}

// ExitFreeStatement is called when production freeStatement is exited.
func (s *BasetlangListener) ExitFreeStatement(ctx *FreeStatementContext) {}

// EnterForkStatement is called when production forkStatement is entered.
func (s *BasetlangListener) EnterForkStatement(ctx *ForkStatementContext) {}

// ExitForkStatement is called when production forkStatement is exited.
func (s *BasetlangListener) ExitForkStatement(ctx *ForkStatementContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BasetlangListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BasetlangListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterBasicType is called when production basicType is entered.
func (s *BasetlangListener) EnterBasicType(ctx *BasicTypeContext) {}

// ExitBasicType is called when production basicType is exited.
func (s *BasetlangListener) ExitBasicType(ctx *BasicTypeContext) {}

// EnterReferenceType is called when production referenceType is entered.
func (s *BasetlangListener) EnterReferenceType(ctx *ReferenceTypeContext) {}

// ExitReferenceType is called when production referenceType is exited.
func (s *BasetlangListener) ExitReferenceType(ctx *ReferenceTypeContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BasetlangListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BasetlangListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BasetlangListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BasetlangListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterIndexTypeList is called when production indexTypeList is entered.
func (s *BasetlangListener) EnterIndexTypeList(ctx *IndexTypeListContext) {}

// ExitIndexTypeList is called when production indexTypeList is exited.
func (s *BasetlangListener) ExitIndexTypeList(ctx *IndexTypeListContext) {}

// EnterStringType is called when production stringType is entered.
func (s *BasetlangListener) EnterStringType(ctx *StringTypeContext) {}

// ExitStringType is called when production stringType is exited.
func (s *BasetlangListener) ExitStringType(ctx *StringTypeContext) {}

// EnterRecordType is called when production recordType is entered.
func (s *BasetlangListener) EnterRecordType(ctx *RecordTypeContext) {}

// ExitRecordType is called when production recordType is exited.
func (s *BasetlangListener) ExitRecordType(ctx *RecordTypeContext) {}

// EnterRecordField is called when production recordField is entered.
func (s *BasetlangListener) EnterRecordField(ctx *RecordFieldContext) {}

// ExitRecordField is called when production recordField is exited.
func (s *BasetlangListener) ExitRecordField(ctx *RecordFieldContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BasetlangListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BasetlangListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterVariableIdentifierList is called when production variableIdentifierList is entered.
func (s *BasetlangListener) EnterVariableIdentifierList(ctx *VariableIdentifierListContext) {}

// ExitVariableIdentifierList is called when production variableIdentifierList is exited.
func (s *BasetlangListener) ExitVariableIdentifierList(ctx *VariableIdentifierListContext) {}

// EnterVariableIdentifier is called when production variableIdentifier is entered.
func (s *BasetlangListener) EnterVariableIdentifier(ctx *VariableIdentifierContext) {}

// ExitVariableIdentifier is called when production variableIdentifier is exited.
func (s *BasetlangListener) ExitVariableIdentifier(ctx *VariableIdentifierContext) {}

// EnterArrayDeclaration is called when production arrayDeclaration is entered.
func (s *BasetlangListener) EnterArrayDeclaration(ctx *ArrayDeclarationContext) {}

// ExitArrayDeclaration is called when production arrayDeclaration is exited.
func (s *BasetlangListener) ExitArrayDeclaration(ctx *ArrayDeclarationContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BasetlangListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BasetlangListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasetlangListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasetlangListener) ExitLiteral(ctx *LiteralContext) {}

// EnterInteger_literal is called when production integer_literal is entered.
func (s *BasetlangListener) EnterInteger_literal(ctx *Integer_literalContext) {}

// ExitInteger_literal is called when production integer_literal is exited.
func (s *BasetlangListener) ExitInteger_literal(ctx *Integer_literalContext) {}

// EnterComment is called when production comment is entered.
func (s *BasetlangListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BasetlangListener) ExitComment(ctx *CommentContext) {}
