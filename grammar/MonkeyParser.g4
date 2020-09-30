parser grammar MonkeyParser;

options {
	tokenVocab = MonkeyLexer;
}

program: statement* EOF;
statement:
	LET letStatement
	| RETURN returnStatement
	| expressionStatement;
letStatement: IDENTIFIER ASSIGN expression SEMI?;
returnStatement: expression SEMI?;
expressionStatement: expression SEMI?;
expression: additionExpression comparison;
comparison: (
		(
			LESS
			| GREATER
			| LESS_OR_EQUALS
			| GREATER_OR_EQUALS
			| EQUALS
		) additionExpression
	)*;
additionExpression: multiplicationExpression additionFactor;
additionFactor: ((PLUS | MINUS) multiplicationExpression)*;
multiplicationExpression:
	elementExpression multiplicationFactor;
multiplicationFactor: ((MULT | DIV) elementExpression)*;
elementExpression:
	primitiveExpression (elementAccess | callExpression)?;
elementAccess: L_BRACKET expression R_BRACKET;
callExpression: L_PAREN expressionList R_PAREN;
primitiveExpression:
	INTEGER
	| STRING
	| IDENTIFIER
	| TRUE
	| FALSE
	| L_PAREN expression R_PAREN
	| arrayLiteral
	| arrayFunctions L_PAREN expressionList R_PAREN
	| functionLiteral
	| hashLiteral
	| printExpression
	| ifExpression;
arrayFunctions: LEN | FIRST | LAST | REST | PUSH;
arrayLiteral: L_BRACKET expressionList R_BRACKET;
functionLiteral:
	FUNC L_PAREN functionParameters R_PAREN blockStatement;
functionParameters: IDENTIFIER moreIdentifiers;
moreIdentifiers: (COMMA IDENTIFIER)*;
hashLiteral: L_CURLY hashContent moreHashContent R_CURLY;
hashContent: expression COLON expression;
moreHashContent: (COMMA hashContent)*;
expressionList: (expression moreExpressions)?;
moreExpressions: (COMMA expression)*;
printExpression: PUTS L_PAREN expression R_PAREN;
ifExpression:
	IF expression blockStatement (ELSE blockStatement)?;
blockStatement: L_CURLY statement* R_CURLY;
