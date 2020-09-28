parser grammar MonkeyParser;

options {
  tokenVocab = MonkeyLexer;
}

program: statement*;
statement: LET letStatement | RETURN returnStatement | expressionStatement;
letStatement: IDENT ASSIGN expression ( PyCOMA?);   
returnStatement: expression ( PyCOMA?);                      
expressionStatement: expression ( PyCOMA?);
expression: additionExpression comparison;                        
comparison: ((LSS|GTR|LEQ|GEQ|EQL) additionExpression)*; 
additionExpression: multiplicationExpression additionFactor;            
additionFactor: ((SUMA|RESTA) multiplicationExpression)*;              
multiplicationExpression: elementExpression multiplicationFactor;                 
multiplicationFactor: ((MULT|DIV) elementExpression)*;                        
elementExpression: primitiveExpression (elementAccess | callExpression)?;                        
elementAccess: LBRACK expression RBRACK; 
callExpression: PIZQ expressionList PDER;
primitiveExpression: LITERAL | STRING | IDENT | TRUE | FALSE | PIZQ expression PDER | arrayLiteral | arrayFunctions PIZQ expressionList PDER | functionLiteral | hashLiteral | printExpression | ifExpression;                              
arrayFunctions: LEN | FIRST | LAST | REST | PUSH;
arrayLiteral: LBRACK expressionList RBRACK; 
functionLiteral: FN PIZQ functionParameters PDER blockStatement;
functionParameters: IDENT moreIdentifiers;
moreIdentifiers: (COMMA IDENT)*;
hashLiteral: LBRACE hashContent moreHashContent RBRACE; 
hashContent: expression DOSPUNTOS expression;
moreHashContent: (COMMA hashContent)*;
expressionList: expression moreExpressions?;
moreExpressions: (COMMA expression)*; 
printExpression: PUTS PIZQ expression PDER;   
ifExpression: IF expression blockStatement (ELSE blockStatement?);
blockStatement: LBRACE statement* RBRACE;
