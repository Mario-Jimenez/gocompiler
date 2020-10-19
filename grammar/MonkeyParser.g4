parser grammar MonkeyParser;

options {
	tokenVocab = MonkeyLexer;
}

program : statement* EOF # programTree;
statement:
	LET IDENTIFIER ASSIGN expression SEMI?	# letStatementTree
	| RETURN expression SEMI?				# returnStatementTree
	| expression SEMI?						# expressionStatementTree
	;
expression:
	additionExpression (comparison additionExpression)* # expressionTree
	;
comparison:
	LESS				# lessComparison
	| GREATER			# greaterComparison
	| LESS_OR_EQUALS	# lessOrEqualsComparison
	| GREATER_OR_EQUALS	# greaterOrEqualsComparison
	| EQUALS			# equalsComparison
	;
additionExpression:
	multiplicationExpression (
		additionFactor multiplicationExpression
	)* # additionTree
	;
additionFactor : PLUS # plusOperator | MINUS # minusOperator;
multiplicationExpression:
	elementExpression (multiplicationFactor elementExpression)* # multiplicationTree
	;
multiplicationFactor:
	MULT	# multiplicationOperator
	| DIV	# divisionOperator
	;
elementExpression:
	primitiveExpression (elementAccess | callExpression)? # elementTree
	;
elementAccess:
	L_BRACKET expression R_BRACKET # elementAccessTree
	;
callExpression:
	L_PAREN expressionList? R_PAREN # functionCallTree
	;
primitiveExpression:
	INTEGER														# integer
	| STRING													# string
	| identifier												# identifierTree
	| TRUE														# true
	| FALSE														# false
	| L_PAREN expression R_PAREN								# groupedExpressionTree
	| L_BRACKET expressionList? R_BRACKET						# arrayTree
	| arrayFunctions L_PAREN expressionList? R_PAREN			# arrayFunctionTree
	| FUNC L_PAREN functionParameters R_PAREN blockStatement	# functionTree
	| L_CURLY hashContent (COMMA hashContent)* R_CURLY			# hashObjectTree
	| PUTS L_PAREN expression R_PAREN							# printTree
	| IF expression blockStatement (ELSE blockStatement)?		# conditionalTree
	;
arrayFunctions:
	LEN		# arrayLen
	| FIRST	# arrayFirst
	| LAST	# arrayLast
	| REST	# arrayRest
	| PUSH	# arrayPush
	;
functionParameters:
	IDENTIFIER (COMMA IDENTIFIER)* # functionParametersTree
	;
hashContent : expression COLON expression # hashPairTree;
expressionList:
	expression (COMMA expression)* # expressionListTree
	;
blockStatement : L_CURLY statement* R_CURLY # blockTree;
identifier
	locals[declaration:*LetStatementTreeContext]:
	IDENTIFIER # identifierNode
	;
