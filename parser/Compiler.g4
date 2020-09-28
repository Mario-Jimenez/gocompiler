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
PUTS        :'puts'; //
FN          :'fn'; //

//FUNCIONES DE ARRAY
LEN         :'len';
FIRST       :'first';
LAST        :'last';
REST        :'rest';
PUSH        :'push';

//OTROS TOKENS
IDENT   : LETRA (LETRA|DIGITO)* ;
LITERAL : DIGITO DIGITO* ;
STRING  : '"' ~ ["\r\n]* '"';

fragment LETRA  : 'a'..'z' | 'A'..'Z' ;
fragment DIGITO : '0'..'9' ;

// Hidden tokens
WS                     : [ \t]+             -> channel(HIDDEN);
COMMENT                : '/*' .*? '*/'      -> channel(HIDDEN);
TERMINATOR             : [\r\n]+            -> channel(HIDDEN);
LINE_COMMENT           : '//' ~[\r\n]*      -> channel(HIDDEN);
