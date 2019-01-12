grammar Turing;

/* primary rules */
program
	/* imports */
	: topLevel+
	;

topLevel
	: statementOrDeclaration
	| subprogramDeclaration
	;

procHeader
	: (PROCEDURE | PROCESS) IDENTIFIER (L_PAREN paramDeclarationList R_PAREN)?
	;

funcHeader
	: FUNCTION IDENTIFIER (L_PAREN paramDeclarationList R_PAREN)? (IDENTIFIER)? COLON typeSpec
	;

subprogramDeclaration
	: subprogramHeader procBody END IDENTIFIER
	;

subprogramHeader
	: funcHeader
	| procHeader
	;

paramDeclaration
	: (VAR)? identifierList COLON typeSpec
	| subprogramHeader
	;

paramDeclarationList
	: paramDeclaration
	| paramDeclarationList COMMA paramDeclaration
	;

procBody
	: statementOrDeclaration+
	;

statementOrDeclaration
	: statement
	| declaration
	;

primaryExpression
	: literal
	| IDENTIFIER
	| L_PAREN expression R_PAREN
	;

argumentList
	: expression
	| argumentList ',' argumentList
	;

exponentialExpression
	: primaryExpression EXP primaryExpression
	;

pointerExpression
	: CARET primaryExpression
	;

postfixExpression
	: primaryExpression
	| pointerExpression
	| exponentialExpression
	| postfixExpression L_PAREN argumentList? R_PAREN
	| postfixExpression DOT IDENTIFIER
	| postfixExpression DEREFERENCE IDENTIFIER
	;

prefixExpression
	: postfixExpression
	| PLUS prefixExpression
	| MINUS prefixExpression
	| CHEAT prefixExpression
	;

multiplicativeExpression
	: postfixExpression
	| prefixExpression
	| multiplicativeExpression (MULTIPLY 
	| DIVIDE 
	| DIV 
	| MOD 
	| REM 
	| SHR 
	| SHL) (postfixExpression | prefixExpression)
	;

additiveExpression
	: multiplicativeExpression
	| additiveExpression (PLUS | MINUS | XOR) multiplicativeExpression
	;

comparativeExpression
	: additiveExpression
	| comparativeExpression (LESSTHAN 
	| GREATERTHAN
	| EQUAL
	| LESSTHANEQUAL
	| GREATERTHANEQUAL
	| NOTEQUAL
	| IN
	| NOT IN) additiveExpression
	;

notExpression
	: comparativeExpression
	| NOT notExpression
	;

andExpression
	: notExpression
	| andExpression AND notExpression
	;

orExpression
	: andExpression
	| orExpression OR andExpression
	;

impliesExpression
	: orExpression
	| impliesExpression IMPLIES orExpression
	;

expression
	: impliesExpression
	;

expressionList
	: expression
	| expressionList COMMA expression
	;

declaration
	: typeDeclaration
	| variableDeclaration
	| arrayDeclaration
	;

statement
	: expression
	| assignmentStatement
	// file statements
	| putStatement
	| EXIT
	| beginStatement
	| RETURN
	| resultStatement
	| newStatement
	| freeStatement
	| forkStatement
	;

assignmentStatement
	: postfixExpression (ASSIGNMENT
	| PLUSEQUALS
	| MINUSEQUALS
	| MULTIPLYEQUALS
	| DIVIDEEQUALS
	| DIVEQUALS
	| SHLEQUALS
	| SHREQUALS) expression
	;

putStatement
	: PUT putItemList
	;

putItem
	: statement
	;

putItemList
	: putItem
	| putItemList COMMA putItem
	;

beginStatement
	: BEGIN statementOrDeclaration END
	;

resultStatement
	: RESULT expression
	;

newStatement
	: NEW (IDENTIFIER COMMA)? IDENTIFIER
	;

freeStatement
	: FREE (IDENTIFIER COMMA)? IDENTIFIER
	;

forkStatement
	: FORK expression
	;

typeDeclaration
	: TYPE IDENTIFIER COLON typeSpec
	;

typeSpec
	: INT
	| REAL
	| BOOLEAN
	| NAT
	| INTN
	| NATN
	| REALN
	| CHAR
	| indexType
	| stringType
	| recordType
	;

indexType
	: integerLiteral RANGE integerLiteral
	;

indexTypeList
	: indexType
	| indexTypeList COMMA indexType
	;

stringType
	: STRING (L_PAREN integerLiteral R_PAREN)?
	;

recordType
	: RECORD
	| identifierList COLON typeSpec
	;

variableDeclaration
	: VAR IDENTIFIER COLON typeSpec (ASSIGNMENT expression)?
	;

arrayDeclaration
	: ARRAY indexType OF typeSpec
	| ARRAY indexType COMMA
	;

identifierList
	: IDENTIFIER
	| identifierList COMMA IDENTIFIER
	;

/* literals */
literal
	: integerLiteral
	| stringLiteral
	;

stringLiteral
	: STRING_LITERAL
	;

integerLiteral
	: DIGIT+
	;

/* keywords */
END		:	'end';
OF		:	'of';
TO		:	'to';
TYPE	:	'type';
VAR		:	'var';
PROCEDURE	:	'procedure' | 'proc';
FUNCTION	:	'function' | 'fcn';
PROCESS	:	'process';
FOR		:	'for';
LOOP	:	'loop';
EXIT	:	'exit';
IF		:	'if';
ELSE	:	'else';
ELSIF	:	'elsif';
CASE	:	'case';
ASSERT	:	'assert';
BEGIN	:	'begin';
RETURN	:	'return';
RESULT	:	'result';
NEW		:	'new';
FREE	:	'free';
TAG		:	'tag';
FORK	: 	'fork';
SIGNAL	:	'signal';
WAIT	:	'wait';
PAUSE	:	'pause';
QUIT	:	'quit';
UNCHECKED	:	'unchecked';
CHECKED	:	'checked';

/* built in procedures */
PUT		:	'put';

/* primitive type keywords */
INT		:	'int';
REAL	:	'real';
BOOLEAN	:	'boolean';
STRING	:	'string';
ARRAY	:	'array';
SET		:	'set';
RECORD	:	'record';
UNION	:	'union';
POINTER	:	'pointer';
NAT		:	'nat';
INTN	:	'intn';
NATN	:	'natn';
REALN	:	'realn';
CHAR	:	'char';

/* symbols and operators */
NOT		:	'not';
CARET	:	'^';	
COLON	:	':';
L_PAREN	:	'(';
R_PAREN	:	')';
DOT		:	'.';
RANGE	:	'..';
COMMA	:	',';
CHEAT	:	'#';
DEREFERENCE	:	'->';
ASSIGNMENT	:	':=';
PLUS	:	'+';
MINUS	:	'-';
MULTIPLY	:	'*';
DIVIDE	:	'/';
DIV		: 	'div';
MOD		:	'mod';
REM		:	'rem';
EXP		:	'**';
LESSTHAN	:	'<';
GREATERTHAN	:	'>';
EQUAL		:	'=';
LESSTHANEQUAL	:	'<=';
GREATERTHANEQUAL	:	'>=';
NOTEQUAL	:	'not=';
AND		:	'and';
OR		:	'or';
IMPLIES	:	'=>';
IN		:	'in';
SHR		:	'shr';
SHL		:	'shl';
XOR		:	'xor';
PLUSEQUALS		:	'+=';
MINUSEQUALS		:	'-=';
MULTIPLYEQUALS	:	'*=';
DIVIDEEQUALS	:	'/=';
DIVEQUALS		: 	'div=';
MODEQUALS		:	'mod=';
ANDEQUALS		:	'and=';
OREQUALS		:	'or=';
IMPLIESEQUALS	:	'=>=';
SHREQUALS		:	'shr=';
SHLEQUALS		:	'shl=';
XOREQUALS		:	'xor=';

/* extras */
IDENTIFIER
    :   NON_DIGIT (NON_DIGIT | DIGIT)*
    ;

NON_DIGIT
    :   [a-zA-Z_]
	;

DIGIT
	:	[0-9]
	;
	
WHITESPACE	
	:	[ \t\r\n]+ 
	-> skip;

/* borrow in part from C.g4 */
STRING_LITERAL
    :   '"' STRING_CHAR_SEQUENCE? '"'
    ;

fragment STRING_CHAR_SEQUENCE
	:	STRING_CHAR+
	;

fragment STRING_CHAR
	:   ~["\\\r\n]
    |   ESCAPE_SEQUENCE
    |   '\\\n'   // Added line
    |   '\\\r\n' // Added line
;

fragment ESCAPE_SEQUENCE
	:   '\\' ['"?abfnrtv\\]
    ;