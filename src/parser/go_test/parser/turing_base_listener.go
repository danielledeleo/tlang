// Code generated from Turing.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Turing

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTuringListener is a complete listener for a parse tree produced by TuringParser.
type BaseTuringListener struct{}

var _ TuringListener = &BaseTuringListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTuringListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTuringListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTuringListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTuringListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseTuringListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseTuringListener) ExitProgram(ctx *ProgramContext) {}

// EnterTopLevel is called when production topLevel is entered.
func (s *BaseTuringListener) EnterTopLevel(ctx *TopLevelContext) {}

// ExitTopLevel is called when production topLevel is exited.
func (s *BaseTuringListener) ExitTopLevel(ctx *TopLevelContext) {}

// EnterProcHeader is called when production procHeader is entered.
func (s *BaseTuringListener) EnterProcHeader(ctx *ProcHeaderContext) {}

// ExitProcHeader is called when production procHeader is exited.
func (s *BaseTuringListener) ExitProcHeader(ctx *ProcHeaderContext) {}

// EnterFuncHeader is called when production funcHeader is entered.
func (s *BaseTuringListener) EnterFuncHeader(ctx *FuncHeaderContext) {}

// ExitFuncHeader is called when production funcHeader is exited.
func (s *BaseTuringListener) ExitFuncHeader(ctx *FuncHeaderContext) {}

// EnterSubprogramHeader is called when production subprogramHeader is entered.
func (s *BaseTuringListener) EnterSubprogramHeader(ctx *SubprogramHeaderContext) {}

// ExitSubprogramHeader is called when production subprogramHeader is exited.
func (s *BaseTuringListener) ExitSubprogramHeader(ctx *SubprogramHeaderContext) {}

// EnterSubprogramDeclaration is called when production subprogramDeclaration is entered.
func (s *BaseTuringListener) EnterSubprogramDeclaration(ctx *SubprogramDeclarationContext) {}

// ExitSubprogramDeclaration is called when production subprogramDeclaration is exited.
func (s *BaseTuringListener) ExitSubprogramDeclaration(ctx *SubprogramDeclarationContext) {}

// EnterSubprogramPrefix is called when production subprogramPrefix is entered.
func (s *BaseTuringListener) EnterSubprogramPrefix(ctx *SubprogramPrefixContext) {}

// ExitSubprogramPrefix is called when production subprogramPrefix is exited.
func (s *BaseTuringListener) ExitSubprogramPrefix(ctx *SubprogramPrefixContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseTuringListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseTuringListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseTuringListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseTuringListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterExportStatement is called when production exportStatement is entered.
func (s *BaseTuringListener) EnterExportStatement(ctx *ExportStatementContext) {}

// ExitExportStatement is called when production exportStatement is exited.
func (s *BaseTuringListener) ExitExportStatement(ctx *ExportStatementContext) {}

// EnterExportListItem is called when production exportListItem is entered.
func (s *BaseTuringListener) EnterExportListItem(ctx *ExportListItemContext) {}

// ExitExportListItem is called when production exportListItem is exited.
func (s *BaseTuringListener) ExitExportListItem(ctx *ExportListItemContext) {}

// EnterExportList is called when production exportList is entered.
func (s *BaseTuringListener) EnterExportList(ctx *ExportListContext) {}

// ExitExportList is called when production exportList is exited.
func (s *BaseTuringListener) ExitExportList(ctx *ExportListContext) {}

// EnterInheritStatement is called when production inheritStatement is entered.
func (s *BaseTuringListener) EnterInheritStatement(ctx *InheritStatementContext) {}

// ExitInheritStatement is called when production inheritStatement is exited.
func (s *BaseTuringListener) ExitInheritStatement(ctx *InheritStatementContext) {}

// EnterImplementStatement is called when production implementStatement is entered.
func (s *BaseTuringListener) EnterImplementStatement(ctx *ImplementStatementContext) {}

// ExitImplementStatement is called when production implementStatement is exited.
func (s *BaseTuringListener) ExitImplementStatement(ctx *ImplementStatementContext) {}

// EnterIdOrFileItem is called when production idOrFileItem is entered.
func (s *BaseTuringListener) EnterIdOrFileItem(ctx *IdOrFileItemContext) {}

// ExitIdOrFileItem is called when production idOrFileItem is exited.
func (s *BaseTuringListener) ExitIdOrFileItem(ctx *IdOrFileItemContext) {}

// EnterHowExport is called when production howExport is entered.
func (s *BaseTuringListener) EnterHowExport(ctx *HowExportContext) {}

// ExitHowExport is called when production howExport is exited.
func (s *BaseTuringListener) ExitHowExport(ctx *HowExportContext) {}

// EnterParamDeclaration is called when production paramDeclaration is entered.
func (s *BaseTuringListener) EnterParamDeclaration(ctx *ParamDeclarationContext) {}

// ExitParamDeclaration is called when production paramDeclaration is exited.
func (s *BaseTuringListener) ExitParamDeclaration(ctx *ParamDeclarationContext) {}

// EnterParamDeclarationList is called when production paramDeclarationList is entered.
func (s *BaseTuringListener) EnterParamDeclarationList(ctx *ParamDeclarationListContext) {}

// ExitParamDeclarationList is called when production paramDeclarationList is exited.
func (s *BaseTuringListener) ExitParamDeclarationList(ctx *ParamDeclarationListContext) {}

// EnterProcBody is called when production procBody is entered.
func (s *BaseTuringListener) EnterProcBody(ctx *ProcBodyContext) {}

// ExitProcBody is called when production procBody is exited.
func (s *BaseTuringListener) ExitProcBody(ctx *ProcBodyContext) {}

// EnterStatementOrDeclaration is called when production statementOrDeclaration is entered.
func (s *BaseTuringListener) EnterStatementOrDeclaration(ctx *StatementOrDeclarationContext) {}

// ExitStatementOrDeclaration is called when production statementOrDeclaration is exited.
func (s *BaseTuringListener) ExitStatementOrDeclaration(ctx *StatementOrDeclarationContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseTuringListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseTuringListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTuringListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTuringListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseTuringListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseTuringListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseTuringListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseTuringListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseTuringListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseTuringListener) ExitStatements(ctx *StatementsContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseTuringListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseTuringListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterPutStatement is called when production putStatement is entered.
func (s *BaseTuringListener) EnterPutStatement(ctx *PutStatementContext) {}

// ExitPutStatement is called when production putStatement is exited.
func (s *BaseTuringListener) ExitPutStatement(ctx *PutStatementContext) {}

// EnterPutItem is called when production putItem is entered.
func (s *BaseTuringListener) EnterPutItem(ctx *PutItemContext) {}

// ExitPutItem is called when production putItem is exited.
func (s *BaseTuringListener) ExitPutItem(ctx *PutItemContext) {}

// EnterPutItemList is called when production putItemList is entered.
func (s *BaseTuringListener) EnterPutItemList(ctx *PutItemListContext) {}

// ExitPutItemList is called when production putItemList is exited.
func (s *BaseTuringListener) ExitPutItemList(ctx *PutItemListContext) {}

// EnterBeginStatement is called when production beginStatement is entered.
func (s *BaseTuringListener) EnterBeginStatement(ctx *BeginStatementContext) {}

// ExitBeginStatement is called when production beginStatement is exited.
func (s *BaseTuringListener) ExitBeginStatement(ctx *BeginStatementContext) {}

// EnterResultStatement is called when production resultStatement is entered.
func (s *BaseTuringListener) EnterResultStatement(ctx *ResultStatementContext) {}

// ExitResultStatement is called when production resultStatement is exited.
func (s *BaseTuringListener) ExitResultStatement(ctx *ResultStatementContext) {}

// EnterNewStatement is called when production newStatement is entered.
func (s *BaseTuringListener) EnterNewStatement(ctx *NewStatementContext) {}

// ExitNewStatement is called when production newStatement is exited.
func (s *BaseTuringListener) ExitNewStatement(ctx *NewStatementContext) {}

// EnterFreeStatement is called when production freeStatement is entered.
func (s *BaseTuringListener) EnterFreeStatement(ctx *FreeStatementContext) {}

// ExitFreeStatement is called when production freeStatement is exited.
func (s *BaseTuringListener) ExitFreeStatement(ctx *FreeStatementContext) {}

// EnterForkStatement is called when production forkStatement is entered.
func (s *BaseTuringListener) EnterForkStatement(ctx *ForkStatementContext) {}

// ExitForkStatement is called when production forkStatement is exited.
func (s *BaseTuringListener) ExitForkStatement(ctx *ForkStatementContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseTuringListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseTuringListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterBasicType is called when production basicType is entered.
func (s *BaseTuringListener) EnterBasicType(ctx *BasicTypeContext) {}

// ExitBasicType is called when production basicType is exited.
func (s *BaseTuringListener) ExitBasicType(ctx *BasicTypeContext) {}

// EnterReferenceType is called when production referenceType is entered.
func (s *BaseTuringListener) EnterReferenceType(ctx *ReferenceTypeContext) {}

// ExitReferenceType is called when production referenceType is exited.
func (s *BaseTuringListener) ExitReferenceType(ctx *ReferenceTypeContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseTuringListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseTuringListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseTuringListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseTuringListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterIndexTypeList is called when production indexTypeList is entered.
func (s *BaseTuringListener) EnterIndexTypeList(ctx *IndexTypeListContext) {}

// ExitIndexTypeList is called when production indexTypeList is exited.
func (s *BaseTuringListener) ExitIndexTypeList(ctx *IndexTypeListContext) {}

// EnterStringType is called when production stringType is entered.
func (s *BaseTuringListener) EnterStringType(ctx *StringTypeContext) {}

// ExitStringType is called when production stringType is exited.
func (s *BaseTuringListener) ExitStringType(ctx *StringTypeContext) {}

// EnterRecordType is called when production recordType is entered.
func (s *BaseTuringListener) EnterRecordType(ctx *RecordTypeContext) {}

// ExitRecordType is called when production recordType is exited.
func (s *BaseTuringListener) ExitRecordType(ctx *RecordTypeContext) {}

// EnterRecordField is called when production recordField is entered.
func (s *BaseTuringListener) EnterRecordField(ctx *RecordFieldContext) {}

// ExitRecordField is called when production recordField is exited.
func (s *BaseTuringListener) ExitRecordField(ctx *RecordFieldContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseTuringListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseTuringListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterVariableIdentifierList is called when production variableIdentifierList is entered.
func (s *BaseTuringListener) EnterVariableIdentifierList(ctx *VariableIdentifierListContext) {}

// ExitVariableIdentifierList is called when production variableIdentifierList is exited.
func (s *BaseTuringListener) ExitVariableIdentifierList(ctx *VariableIdentifierListContext) {}

// EnterVariableIdentifier is called when production variableIdentifier is entered.
func (s *BaseTuringListener) EnterVariableIdentifier(ctx *VariableIdentifierContext) {}

// ExitVariableIdentifier is called when production variableIdentifier is exited.
func (s *BaseTuringListener) ExitVariableIdentifier(ctx *VariableIdentifierContext) {}

// EnterArrayDeclaration is called when production arrayDeclaration is entered.
func (s *BaseTuringListener) EnterArrayDeclaration(ctx *ArrayDeclarationContext) {}

// ExitArrayDeclaration is called when production arrayDeclaration is exited.
func (s *BaseTuringListener) ExitArrayDeclaration(ctx *ArrayDeclarationContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseTuringListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseTuringListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseTuringListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseTuringListener) ExitLiteral(ctx *LiteralContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseTuringListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseTuringListener) ExitComment(ctx *CommentContext) {}
