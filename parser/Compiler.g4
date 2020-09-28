lexer grammar Compiler;

//SIMBOLOS
PyCOMA      : ';' ; //
DOSPUNTOS   : ':' ; //
PIZQ        : '(' ; //
PDER        : ')' ; //
LBRACK      : '[' ; //
RBRACK      : ']' ; //
LBRACE      : '{' ; //
RBRACE      : '}' ; //
EQL         : '=='; //
ASSIGN      : '=';  //
LEQ         : '<='; //
GEQ         : '>='; //
LSS         : '<'; //
GTR         : '>'; //

//OPERADORES
SUMA        : '+' ; //
RESTA       : '-' ; //
MULT        : '*' ; //
DIV         : '/' ; //

//PALABRAS RESERVADAS
IF          : 'if' ; //
ELSE        : 'else' ; //
LET         : 'let' ; //
RETURN      :'return'; //
TRUE        :'true'; //
FALSE       :'false'; //

//FUNCIONES DE ARRAY
LEN         :'len';
FIRST       :'first';
LAST        :'last';
REST        :'rest';
PUSH        :'push';

//OTROS TOKENS
IDENT   : LETRA (LETRA|DIGITO)* ;
LITERAL : DIGITO DIGITO* ;

//fragment EXPR : ('0'+ '1')* '01';
fragment LETRA  : 'a'..'z' | 'A'..'Z' ;
fragment DIGITO : '0'..'9' ;

//PARA SALTAR CIERTAS EXPRESIONES --> COMO EL ESPACIO
WS : [ \r\t\n]+ -> skip ;
