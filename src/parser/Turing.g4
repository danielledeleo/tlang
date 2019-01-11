grammar Turing;

/* primary rules */
program
	: typeDeclaration
	| variableDeclaration
	;

typeSpec
	:	INT
	|	REAL
	|	BOOLEAN
	| 	NAT
	|	INTN
	|	NATN
	|	REALN
	|	CHAR
	|	STRING
	
	|	STRING L_PAREN integerLiteral R_PAREN
	;

indexType
	:	integerLiteral RANGE integerLiteral
	;

indexTypeList
	:	indexType
	|	indexTypeList COMMA indexType
	;
	
typeDeclaration
	:	TYPE IDENTIFIER COLON typeSpec
	;
	
variableDeclaration
	:	VAR IDENTIFIER COLON typeSpec
	|	VAR IDENTIFIER ASSIGNMENT initializingValue
	|	VAR IDENTIFIER COLON typeSpec ASSIGNMENT initializingValue
	;

initializingValue
	:
	;

arrayDeclaration
	:	ARRAY indexType OF typeSpec
	|	ARRAY indexType COMMA
	;
/* literals */
integerLiteral
	:	DIGIT+
	;

/* keywords */
END		:	'end';
OF		:	'of';
TO		:	'to';
TYPE	:	'type';
VAR		:	'var';

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

/* symbols */
CARET	:	'^';	
COLON	:	':';
L_PAREN	:	'(';
R_PAREN	:	')';
RANGE	:	'..';
COMMA	:	',';
ASSIGNMENT	:	':=';


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
