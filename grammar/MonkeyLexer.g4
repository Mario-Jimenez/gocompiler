lexer grammar MonkeyLexer;

//SIMBOLOS
SEMI: ';';
COLON: ':';
COMMA: ',';
L_PAREN: '(';
R_PAREN: ')';
L_BRACKET: '[';
R_BRACKET: ']';
L_CURLY: '{';
R_CURLY: '}';

EQUALS: '==';
ASSIGN: '=';
LESS_OR_EQUALS: '<=';
GREATER_OR_EQUALS: '>=';
LESS: '<';
GREATER: '>';

//OPERADORES
PLUS: '+';
MINUS: '-';
MULT: '*';
DIV: '/';

//PALABRAS RESERVADAS
IF: 'if';
ELSE: 'else';
LET: 'let';
RETURN: 'return';
TRUE: 'true';
FALSE: 'false';
PUTS: 'puts';
FUNC: 'fn';

//FUNCIONES DE ARRAY
LEN: 'len';
FIRST: 'first';
LAST: 'last';
REST: 'rest';
PUSH: 'push';

//OTROS TOKENS
IDENTIFIER: LETTER (LETTER | DIGIT)*;
INTEGER: DIGIT DIGIT*;
STRING: '"' ~ ["\r\n]* '"';

fragment LETTER: [a-zA-Z_];
fragment DIGIT: [0-9];

// Hidden tokens
WS: [ \t]+ -> channel(HIDDEN);
COMMENT: '/*' .*? '*/' -> channel(HIDDEN);
TERMINATOR: [\r\n]+ -> channel(HIDDEN);
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
