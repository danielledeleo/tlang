// Code generated from tlang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // tlang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// tlangListener is a complete listener for a parse tree produced by tlangParser.
type tlangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTopLevel is called when entering the topLevel production.
	EnterTopLevel(c *TopLevelContext)

	// EnterProcHeader is called when entering the procHeader production.
	EnterProcHeader(c *ProcHeaderContext)

	// EnterFuncHeader is called when entering the funcHeader production.
	EnterFuncHeader(c *FuncHeaderContext)

	// EnterSubprogramHeader is called when entering the subprogramHeader production.
	EnterSubprogramHeader(c *SubprogramHeaderContext)

	// EnterSubprogramDeclaration is called when entering the subprogramDeclaration production.
	EnterSubprogramDeclaration(c *SubprogramDeclarationContext)

	// EnterSubprogramPrefix is called when entering the subprogramPrefix production.
	EnterSubprogramPrefix(c *SubprogramPrefixContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterExportStatement is called when entering the exportStatement production.
	EnterExportStatement(c *ExportStatementContext)

	// EnterExportListItem is called when entering the exportListItem production.
	EnterExportListItem(c *ExportListItemContext)

	// EnterExportList is called when entering the exportList production.
	EnterExportList(c *ExportListContext)

	// EnterInheritStatement is called when entering the inheritStatement production.
	EnterInheritStatement(c *InheritStatementContext)

	// EnterImplementStatement is called when entering the implementStatement production.
	EnterImplementStatement(c *ImplementStatementContext)

	// EnterIdOrFileItem is called when entering the idOrFileItem production.
	EnterIdOrFileItem(c *IdOrFileItemContext)

	// EnterHowExport is called when entering the howExport production.
	EnterHowExport(c *HowExportContext)

	// EnterParamDeclaration is called when entering the paramDeclaration production.
	EnterParamDeclaration(c *ParamDeclarationContext)

	// EnterParamDeclarationList is called when entering the paramDeclarationList production.
	EnterParamDeclarationList(c *ParamDeclarationListContext)

	// EnterProcBody is called when entering the procBody production.
	EnterProcBody(c *ProcBodyContext)

	// EnterStatementOrDeclaration is called when entering the statementOrDeclaration production.
	EnterStatementOrDeclaration(c *StatementOrDeclarationContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterPutStatement is called when entering the putStatement production.
	EnterPutStatement(c *PutStatementContext)

	// EnterPutItem is called when entering the putItem production.
	EnterPutItem(c *PutItemContext)

	// EnterPutItemList is called when entering the putItemList production.
	EnterPutItemList(c *PutItemListContext)

	// EnterBeginStatement is called when entering the beginStatement production.
	EnterBeginStatement(c *BeginStatementContext)

	// EnterResultStatement is called when entering the resultStatement production.
	EnterResultStatement(c *ResultStatementContext)

	// EnterNewStatement is called when entering the newStatement production.
	EnterNewStatement(c *NewStatementContext)

	// EnterFreeStatement is called when entering the freeStatement production.
	EnterFreeStatement(c *FreeStatementContext)

	// EnterForkStatement is called when entering the forkStatement production.
	EnterForkStatement(c *ForkStatementContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterBasicType is called when entering the basicType production.
	EnterBasicType(c *BasicTypeContext)

	// EnterReferenceType is called when entering the referenceType production.
	EnterReferenceType(c *ReferenceTypeContext)

	// EnterTypeSpec is called when entering the typeSpec production.
	EnterTypeSpec(c *TypeSpecContext)

	// EnterIndexType is called when entering the indexType production.
	EnterIndexType(c *IndexTypeContext)

	// EnterIndexTypeList is called when entering the indexTypeList production.
	EnterIndexTypeList(c *IndexTypeListContext)

	// EnterStringType is called when entering the stringType production.
	EnterStringType(c *StringTypeContext)

	// EnterRecordType is called when entering the recordType production.
	EnterRecordType(c *RecordTypeContext)

	// EnterRecordField is called when entering the recordField production.
	EnterRecordField(c *RecordFieldContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterVariableIdentifierList is called when entering the variableIdentifierList production.
	EnterVariableIdentifierList(c *VariableIdentifierListContext)

	// EnterVariableIdentifier is called when entering the variableIdentifier production.
	EnterVariableIdentifier(c *VariableIdentifierContext)

	// EnterArrayDeclaration is called when entering the arrayDeclaration production.
	EnterArrayDeclaration(c *ArrayDeclarationContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterInteger_literal is called when entering the integer_literal production.
	EnterInteger_literal(c *Integer_literalContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTopLevel is called when exiting the topLevel production.
	ExitTopLevel(c *TopLevelContext)

	// ExitProcHeader is called when exiting the procHeader production.
	ExitProcHeader(c *ProcHeaderContext)

	// ExitFuncHeader is called when exiting the funcHeader production.
	ExitFuncHeader(c *FuncHeaderContext)

	// ExitSubprogramHeader is called when exiting the subprogramHeader production.
	ExitSubprogramHeader(c *SubprogramHeaderContext)

	// ExitSubprogramDeclaration is called when exiting the subprogramDeclaration production.
	ExitSubprogramDeclaration(c *SubprogramDeclarationContext)

	// ExitSubprogramPrefix is called when exiting the subprogramPrefix production.
	ExitSubprogramPrefix(c *SubprogramPrefixContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitExportStatement is called when exiting the exportStatement production.
	ExitExportStatement(c *ExportStatementContext)

	// ExitExportListItem is called when exiting the exportListItem production.
	ExitExportListItem(c *ExportListItemContext)

	// ExitExportList is called when exiting the exportList production.
	ExitExportList(c *ExportListContext)

	// ExitInheritStatement is called when exiting the inheritStatement production.
	ExitInheritStatement(c *InheritStatementContext)

	// ExitImplementStatement is called when exiting the implementStatement production.
	ExitImplementStatement(c *ImplementStatementContext)

	// ExitIdOrFileItem is called when exiting the idOrFileItem production.
	ExitIdOrFileItem(c *IdOrFileItemContext)

	// ExitHowExport is called when exiting the howExport production.
	ExitHowExport(c *HowExportContext)

	// ExitParamDeclaration is called when exiting the paramDeclaration production.
	ExitParamDeclaration(c *ParamDeclarationContext)

	// ExitParamDeclarationList is called when exiting the paramDeclarationList production.
	ExitParamDeclarationList(c *ParamDeclarationListContext)

	// ExitProcBody is called when exiting the procBody production.
	ExitProcBody(c *ProcBodyContext)

	// ExitStatementOrDeclaration is called when exiting the statementOrDeclaration production.
	ExitStatementOrDeclaration(c *StatementOrDeclarationContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitPutStatement is called when exiting the putStatement production.
	ExitPutStatement(c *PutStatementContext)

	// ExitPutItem is called when exiting the putItem production.
	ExitPutItem(c *PutItemContext)

	// ExitPutItemList is called when exiting the putItemList production.
	ExitPutItemList(c *PutItemListContext)

	// ExitBeginStatement is called when exiting the beginStatement production.
	ExitBeginStatement(c *BeginStatementContext)

	// ExitResultStatement is called when exiting the resultStatement production.
	ExitResultStatement(c *ResultStatementContext)

	// ExitNewStatement is called when exiting the newStatement production.
	ExitNewStatement(c *NewStatementContext)

	// ExitFreeStatement is called when exiting the freeStatement production.
	ExitFreeStatement(c *FreeStatementContext)

	// ExitForkStatement is called when exiting the forkStatement production.
	ExitForkStatement(c *ForkStatementContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitBasicType is called when exiting the basicType production.
	ExitBasicType(c *BasicTypeContext)

	// ExitReferenceType is called when exiting the referenceType production.
	ExitReferenceType(c *ReferenceTypeContext)

	// ExitTypeSpec is called when exiting the typeSpec production.
	ExitTypeSpec(c *TypeSpecContext)

	// ExitIndexType is called when exiting the indexType production.
	ExitIndexType(c *IndexTypeContext)

	// ExitIndexTypeList is called when exiting the indexTypeList production.
	ExitIndexTypeList(c *IndexTypeListContext)

	// ExitStringType is called when exiting the stringType production.
	ExitStringType(c *StringTypeContext)

	// ExitRecordType is called when exiting the recordType production.
	ExitRecordType(c *RecordTypeContext)

	// ExitRecordField is called when exiting the recordField production.
	ExitRecordField(c *RecordFieldContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitVariableIdentifierList is called when exiting the variableIdentifierList production.
	ExitVariableIdentifierList(c *VariableIdentifierListContext)

	// ExitVariableIdentifier is called when exiting the variableIdentifier production.
	ExitVariableIdentifier(c *VariableIdentifierContext)

	// ExitArrayDeclaration is called when exiting the arrayDeclaration production.
	ExitArrayDeclaration(c *ArrayDeclarationContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitInteger_literal is called when exiting the integer_literal production.
	ExitInteger_literal(c *Integer_literalContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
