// Code generated from ./grammar/MonkeyParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MonkeyParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 223,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 3, 2, 7, 2, 60, 10, 2, 12, 2, 14, 2, 63, 11, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 72, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 78, 10, 4, 3, 5, 3, 5, 5, 5, 82, 10, 5, 3, 6, 3, 6, 5, 6, 86, 10,
	6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 7, 8, 93, 10, 8, 12, 8, 14, 8, 96, 11,
	8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 7, 10, 103, 10, 10, 12, 10, 14, 10,
	106, 11, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 7, 12, 113, 10, 12, 12,
	12, 14, 12, 116, 11, 12, 3, 13, 3, 13, 3, 13, 5, 13, 121, 10, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 150, 10, 16, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 7, 21, 169, 10, 21, 12, 21, 14, 21, 172, 11,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24,
	3, 24, 7, 24, 185, 10, 24, 12, 24, 14, 24, 188, 11, 24, 3, 25, 3, 25, 3,
	25, 5, 25, 193, 10, 25, 3, 26, 3, 26, 7, 26, 197, 10, 26, 12, 26, 14, 26,
	200, 11, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 5, 28, 212, 10, 28, 3, 29, 3, 29, 7, 29, 216, 10, 29, 12, 29,
	14, 29, 219, 11, 29, 3, 29, 3, 29, 3, 29, 2, 2, 30, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 52, 54, 56, 2, 6, 4, 2, 12, 12, 14, 17, 3, 2, 18, 19, 3, 2, 20, 21,
	3, 2, 30, 34, 2, 222, 2, 61, 3, 2, 2, 2, 4, 71, 3, 2, 2, 2, 6, 73, 3, 2,
	2, 2, 8, 79, 3, 2, 2, 2, 10, 83, 3, 2, 2, 2, 12, 87, 3, 2, 2, 2, 14, 94,
	3, 2, 2, 2, 16, 97, 3, 2, 2, 2, 18, 104, 3, 2, 2, 2, 20, 107, 3, 2, 2,
	2, 22, 114, 3, 2, 2, 2, 24, 117, 3, 2, 2, 2, 26, 122, 3, 2, 2, 2, 28, 126,
	3, 2, 2, 2, 30, 149, 3, 2, 2, 2, 32, 151, 3, 2, 2, 2, 34, 153, 3, 2, 2,
	2, 36, 157, 3, 2, 2, 2, 38, 163, 3, 2, 2, 2, 40, 170, 3, 2, 2, 2, 42, 173,
	3, 2, 2, 2, 44, 178, 3, 2, 2, 2, 46, 186, 3, 2, 2, 2, 48, 192, 3, 2, 2,
	2, 50, 198, 3, 2, 2, 2, 52, 201, 3, 2, 2, 2, 54, 206, 3, 2, 2, 2, 56, 213,
	3, 2, 2, 2, 58, 60, 5, 4, 3, 2, 59, 58, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2,
	61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 61, 3,
	2, 2, 2, 64, 65, 7, 2, 2, 3, 65, 3, 3, 2, 2, 2, 66, 67, 7, 24, 2, 2, 67,
	72, 5, 6, 4, 2, 68, 69, 7, 25, 2, 2, 69, 72, 5, 8, 5, 2, 70, 72, 5, 10,
	6, 2, 71, 66, 3, 2, 2, 2, 71, 68, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 5,
	3, 2, 2, 2, 73, 74, 7, 35, 2, 2, 74, 75, 7, 13, 2, 2, 75, 77, 5, 12, 7,
	2, 76, 78, 7, 3, 2, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 7, 3,
	2, 2, 2, 79, 81, 5, 12, 7, 2, 80, 82, 7, 3, 2, 2, 81, 80, 3, 2, 2, 2, 81,
	82, 3, 2, 2, 2, 82, 9, 3, 2, 2, 2, 83, 85, 5, 12, 7, 2, 84, 86, 7, 3, 2,
	2, 85, 84, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 11, 3, 2, 2, 2, 87, 88,
	5, 16, 9, 2, 88, 89, 5, 14, 8, 2, 89, 13, 3, 2, 2, 2, 90, 91, 9, 2, 2,
	2, 91, 93, 5, 16, 9, 2, 92, 90, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92,
	3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 15, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2,
	97, 98, 5, 20, 11, 2, 98, 99, 5, 18, 10, 2, 99, 17, 3, 2, 2, 2, 100, 101,
	9, 3, 2, 2, 101, 103, 5, 20, 11, 2, 102, 100, 3, 2, 2, 2, 103, 106, 3,
	2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 19, 3, 2, 2,
	2, 106, 104, 3, 2, 2, 2, 107, 108, 5, 24, 13, 2, 108, 109, 5, 22, 12, 2,
	109, 21, 3, 2, 2, 2, 110, 111, 9, 4, 2, 2, 111, 113, 5, 24, 13, 2, 112,
	110, 3, 2, 2, 2, 113, 116, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 115,
	3, 2, 2, 2, 115, 23, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 117, 120, 5, 30,
	16, 2, 118, 121, 5, 26, 14, 2, 119, 121, 5, 28, 15, 2, 120, 118, 3, 2,
	2, 2, 120, 119, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 25, 3, 2, 2, 2,
	122, 123, 7, 8, 2, 2, 123, 124, 5, 12, 7, 2, 124, 125, 7, 9, 2, 2, 125,
	27, 3, 2, 2, 2, 126, 127, 7, 6, 2, 2, 127, 128, 5, 48, 25, 2, 128, 129,
	7, 7, 2, 2, 129, 29, 3, 2, 2, 2, 130, 150, 7, 36, 2, 2, 131, 150, 7, 37,
	2, 2, 132, 150, 7, 35, 2, 2, 133, 150, 7, 26, 2, 2, 134, 150, 7, 27, 2,
	2, 135, 136, 7, 6, 2, 2, 136, 137, 5, 12, 7, 2, 137, 138, 7, 7, 2, 2, 138,
	150, 3, 2, 2, 2, 139, 150, 5, 34, 18, 2, 140, 141, 5, 32, 17, 2, 141, 142,
	7, 6, 2, 2, 142, 143, 5, 48, 25, 2, 143, 144, 7, 7, 2, 2, 144, 150, 3,
	2, 2, 2, 145, 150, 5, 36, 19, 2, 146, 150, 5, 42, 22, 2, 147, 150, 5, 52,
	27, 2, 148, 150, 5, 54, 28, 2, 149, 130, 3, 2, 2, 2, 149, 131, 3, 2, 2,
	2, 149, 132, 3, 2, 2, 2, 149, 133, 3, 2, 2, 2, 149, 134, 3, 2, 2, 2, 149,
	135, 3, 2, 2, 2, 149, 139, 3, 2, 2, 2, 149, 140, 3, 2, 2, 2, 149, 145,
	3, 2, 2, 2, 149, 146, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 148, 3, 2,
	2, 2, 150, 31, 3, 2, 2, 2, 151, 152, 9, 5, 2, 2, 152, 33, 3, 2, 2, 2, 153,
	154, 7, 8, 2, 2, 154, 155, 5, 48, 25, 2, 155, 156, 7, 9, 2, 2, 156, 35,
	3, 2, 2, 2, 157, 158, 7, 29, 2, 2, 158, 159, 7, 6, 2, 2, 159, 160, 5, 38,
	20, 2, 160, 161, 7, 7, 2, 2, 161, 162, 5, 56, 29, 2, 162, 37, 3, 2, 2,
	2, 163, 164, 7, 35, 2, 2, 164, 165, 5, 40, 21, 2, 165, 39, 3, 2, 2, 2,
	166, 167, 7, 5, 2, 2, 167, 169, 7, 35, 2, 2, 168, 166, 3, 2, 2, 2, 169,
	172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 41, 3,
	2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 174, 7, 10, 2, 2, 174, 175, 5, 44,
	23, 2, 175, 176, 5, 46, 24, 2, 176, 177, 7, 11, 2, 2, 177, 43, 3, 2, 2,
	2, 178, 179, 5, 12, 7, 2, 179, 180, 7, 4, 2, 2, 180, 181, 5, 12, 7, 2,
	181, 45, 3, 2, 2, 2, 182, 183, 7, 5, 2, 2, 183, 185, 5, 44, 23, 2, 184,
	182, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186, 187,
	3, 2, 2, 2, 187, 47, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189, 190, 5, 12,
	7, 2, 190, 191, 5, 50, 26, 2, 191, 193, 3, 2, 2, 2, 192, 189, 3, 2, 2,
	2, 192, 193, 3, 2, 2, 2, 193, 49, 3, 2, 2, 2, 194, 195, 7, 5, 2, 2, 195,
	197, 5, 12, 7, 2, 196, 194, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2, 198, 196,
	3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 51, 3, 2, 2, 2, 200, 198, 3, 2,
	2, 2, 201, 202, 7, 28, 2, 2, 202, 203, 7, 6, 2, 2, 203, 204, 5, 12, 7,
	2, 204, 205, 7, 7, 2, 2, 205, 53, 3, 2, 2, 2, 206, 207, 7, 22, 2, 2, 207,
	208, 5, 12, 7, 2, 208, 211, 5, 56, 29, 2, 209, 210, 7, 23, 2, 2, 210, 212,
	5, 56, 29, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 55, 3, 2,
	2, 2, 213, 217, 7, 10, 2, 2, 214, 216, 5, 4, 3, 2, 215, 214, 3, 2, 2, 2,
	216, 219, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218,
	220, 3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 220, 221, 7, 11, 2, 2, 221, 57,
	3, 2, 2, 2, 18, 61, 71, 77, 81, 85, 94, 104, 114, 120, 149, 170, 186, 192,
	198, 211, 217,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "':'", "','", "'('", "')'", "'['", "']'", "'{'", "'}'", "'=='",
	"'='", "'<='", "'>='", "'<'", "'>'", "'+'", "'-'", "'*'", "'/'", "'if'",
	"'else'", "'let'", "'return'", "'true'", "'false'", "'puts'", "'fn'", "'len'",
	"'first'", "'last'", "'rest'", "'push'",
}
var symbolicNames = []string{
	"", "SEMI", "COLON", "COMMA", "L_PAREN", "R_PAREN", "L_BRACKET", "R_BRACKET",
	"L_CURLY", "R_CURLY", "EQUALS", "ASSIGN", "LESS_OR_EQUALS", "GREATER_OR_EQUALS",
	"LESS", "GREATER", "PLUS", "MINUS", "MULT", "DIV", "IF", "ELSE", "LET",
	"RETURN", "TRUE", "FALSE", "PUTS", "FUNC", "LEN", "FIRST", "LAST", "REST",
	"PUSH", "IDENTIFIER", "INTEGER", "STRING", "WS", "COMMENT", "TERMINATOR",
	"LINE_COMMENT",
}

var ruleNames = []string{
	"program", "statement", "letStatement", "returnStatement", "expressionStatement",
	"expression", "comparison", "additionExpression", "additionFactor", "multiplicationExpression",
	"multiplicationFactor", "elementExpression", "elementAccess", "callExpression",
	"primitiveExpression", "arrayFunctions", "arrayLiteral", "functionLiteral",
	"functionParameters", "moreIdentifiers", "hashLiteral", "hashContent",
	"moreHashContent", "expressionList", "moreExpressions", "printExpression",
	"ifExpression", "blockStatement",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MonkeyParser struct {
	*antlr.BaseParser
}

func NewMonkeyParser(input antlr.TokenStream) *MonkeyParser {
	this := new(MonkeyParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "MonkeyParser.g4"

	return this
}

// MonkeyParser tokens.
const (
	MonkeyParserEOF               = antlr.TokenEOF
	MonkeyParserSEMI              = 1
	MonkeyParserCOLON             = 2
	MonkeyParserCOMMA             = 3
	MonkeyParserL_PAREN           = 4
	MonkeyParserR_PAREN           = 5
	MonkeyParserL_BRACKET         = 6
	MonkeyParserR_BRACKET         = 7
	MonkeyParserL_CURLY           = 8
	MonkeyParserR_CURLY           = 9
	MonkeyParserEQUALS            = 10
	MonkeyParserASSIGN            = 11
	MonkeyParserLESS_OR_EQUALS    = 12
	MonkeyParserGREATER_OR_EQUALS = 13
	MonkeyParserLESS              = 14
	MonkeyParserGREATER           = 15
	MonkeyParserPLUS              = 16
	MonkeyParserMINUS             = 17
	MonkeyParserMULT              = 18
	MonkeyParserDIV               = 19
	MonkeyParserIF                = 20
	MonkeyParserELSE              = 21
	MonkeyParserLET               = 22
	MonkeyParserRETURN            = 23
	MonkeyParserTRUE              = 24
	MonkeyParserFALSE             = 25
	MonkeyParserPUTS              = 26
	MonkeyParserFUNC              = 27
	MonkeyParserLEN               = 28
	MonkeyParserFIRST             = 29
	MonkeyParserLAST              = 30
	MonkeyParserREST              = 31
	MonkeyParserPUSH              = 32
	MonkeyParserIDENTIFIER        = 33
	MonkeyParserINTEGER           = 34
	MonkeyParserSTRING            = 35
	MonkeyParserWS                = 36
	MonkeyParserCOMMENT           = 37
	MonkeyParserTERMINATOR        = 38
	MonkeyParserLINE_COMMENT      = 39
)

// MonkeyParser rules.
const (
	MonkeyParserRULE_program                  = 0
	MonkeyParserRULE_statement                = 1
	MonkeyParserRULE_letStatement             = 2
	MonkeyParserRULE_returnStatement          = 3
	MonkeyParserRULE_expressionStatement      = 4
	MonkeyParserRULE_expression               = 5
	MonkeyParserRULE_comparison               = 6
	MonkeyParserRULE_additionExpression       = 7
	MonkeyParserRULE_additionFactor           = 8
	MonkeyParserRULE_multiplicationExpression = 9
	MonkeyParserRULE_multiplicationFactor     = 10
	MonkeyParserRULE_elementExpression        = 11
	MonkeyParserRULE_elementAccess            = 12
	MonkeyParserRULE_callExpression           = 13
	MonkeyParserRULE_primitiveExpression      = 14
	MonkeyParserRULE_arrayFunctions           = 15
	MonkeyParserRULE_arrayLiteral             = 16
	MonkeyParserRULE_functionLiteral          = 17
	MonkeyParserRULE_functionParameters       = 18
	MonkeyParserRULE_moreIdentifiers          = 19
	MonkeyParserRULE_hashLiteral              = 20
	MonkeyParserRULE_hashContent              = 21
	MonkeyParserRULE_moreHashContent          = 22
	MonkeyParserRULE_expressionList           = 23
	MonkeyParserRULE_moreExpressions          = 24
	MonkeyParserRULE_printExpression          = 25
	MonkeyParserRULE_ifExpression             = 26
	MonkeyParserRULE_blockStatement           = 27
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MonkeyParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(MonkeyParserL_PAREN-4))|(1<<(MonkeyParserL_BRACKET-4))|(1<<(MonkeyParserL_CURLY-4))|(1<<(MonkeyParserIF-4))|(1<<(MonkeyParserLET-4))|(1<<(MonkeyParserRETURN-4))|(1<<(MonkeyParserTRUE-4))|(1<<(MonkeyParserFALSE-4))|(1<<(MonkeyParserPUTS-4))|(1<<(MonkeyParserFUNC-4))|(1<<(MonkeyParserLEN-4))|(1<<(MonkeyParserFIRST-4))|(1<<(MonkeyParserLAST-4))|(1<<(MonkeyParserREST-4))|(1<<(MonkeyParserPUSH-4))|(1<<(MonkeyParserIDENTIFIER-4))|(1<<(MonkeyParserINTEGER-4))|(1<<(MonkeyParserSTRING-4)))) != 0 {
		{
			p.SetState(56)
			p.Statement()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(MonkeyParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) LET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLET, 0)
}

func (s *StatementContext) LetStatement() ILetStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILetStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILetStatementContext)
}

func (s *StatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRETURN, 0)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MonkeyParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(MonkeyParserLET)
		}
		{
			p.SetState(65)
			p.LetStatement()
		}

	case MonkeyParserRETURN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(MonkeyParserRETURN)
		}
		{
			p.SetState(67)
			p.ReturnStatement()
		}

	case MonkeyParserL_PAREN, MonkeyParserL_BRACKET, MonkeyParserL_CURLY, MonkeyParserIF, MonkeyParserTRUE, MonkeyParserFALSE, MonkeyParserPUTS, MonkeyParserFUNC, MonkeyParserLEN, MonkeyParserFIRST, MonkeyParserLAST, MonkeyParserREST, MonkeyParserPUSH, MonkeyParserIDENTIFIER, MonkeyParserINTEGER, MonkeyParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.ExpressionStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILetStatementContext is an interface to support dynamic dispatch.
type ILetStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLetStatementContext differentiates from other interfaces.
	IsLetStatementContext()
}

type LetStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetStatementContext() *LetStatementContext {
	var p = new(LetStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_letStatement
	return p
}

func (*LetStatementContext) IsLetStatementContext() {}

func NewLetStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetStatementContext {
	var p = new(LetStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_letStatement

	return p
}

func (s *LetStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LetStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENTIFIER, 0)
}

func (s *LetStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserASSIGN, 0)
}

func (s *LetStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LetStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSEMI, 0)
}

func (s *LetStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) LetStatement() (localctx ILetStatementContext) {
	localctx = NewLetStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MonkeyParserRULE_letStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(MonkeyParserIDENTIFIER)
	}
	{
		p.SetState(72)
		p.Match(MonkeyParserASSIGN)
	}
	{
		p.SetState(73)
		p.Expression()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MonkeyParserSEMI {
		{
			p.SetState(74)
			p.Match(MonkeyParserSEMI)
		}

	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSEMI, 0)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MonkeyParserRULE_returnStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Expression()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MonkeyParserSEMI {
		{
			p.SetState(78)
			p.Match(MonkeyParserSEMI)
		}

	}

	return localctx
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_expressionStatement
	return p
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSEMI, 0)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MonkeyParserRULE_expressionStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Expression()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MonkeyParserSEMI {
		{
			p.SetState(82)
			p.Match(MonkeyParserSEMI)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AdditionExpression() IAdditionExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditionExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditionExpressionContext)
}

func (s *ExpressionContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MonkeyParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.AdditionExpression()
	}
	{
		p.SetState(86)
		p.Comparison()
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) AllAdditionExpression() []IAdditionExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdditionExpressionContext)(nil)).Elem())
	var tst = make([]IAdditionExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdditionExpressionContext)
		}
	}

	return tst
}

func (s *ComparisonContext) AdditionExpression(i int) IAdditionExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditionExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdditionExpressionContext)
}

func (s *ComparisonContext) AllLESS() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserLESS)
}

func (s *ComparisonContext) LESS(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserLESS, i)
}

func (s *ComparisonContext) AllGREATER() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserGREATER)
}

func (s *ComparisonContext) GREATER(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserGREATER, i)
}

func (s *ComparisonContext) AllLESS_OR_EQUALS() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserLESS_OR_EQUALS)
}

func (s *ComparisonContext) LESS_OR_EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserLESS_OR_EQUALS, i)
}

func (s *ComparisonContext) AllGREATER_OR_EQUALS() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserGREATER_OR_EQUALS)
}

func (s *ComparisonContext) GREATER_OR_EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserGREATER_OR_EQUALS, i)
}

func (s *ComparisonContext) AllEQUALS() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserEQUALS)
}

func (s *ComparisonContext) EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserEQUALS, i)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MonkeyParserRULE_comparison)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MonkeyParserEQUALS)|(1<<MonkeyParserLESS_OR_EQUALS)|(1<<MonkeyParserGREATER_OR_EQUALS)|(1<<MonkeyParserLESS)|(1<<MonkeyParserGREATER))) != 0 {
		{
			p.SetState(88)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MonkeyParserEQUALS)|(1<<MonkeyParserLESS_OR_EQUALS)|(1<<MonkeyParserGREATER_OR_EQUALS)|(1<<MonkeyParserLESS)|(1<<MonkeyParserGREATER))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(89)
			p.AdditionExpression()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAdditionExpressionContext is an interface to support dynamic dispatch.
type IAdditionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditionExpressionContext differentiates from other interfaces.
	IsAdditionExpressionContext()
}

type AdditionExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditionExpressionContext() *AdditionExpressionContext {
	var p = new(AdditionExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_additionExpression
	return p
}

func (*AdditionExpressionContext) IsAdditionExpressionContext() {}

func NewAdditionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditionExpressionContext {
	var p = new(AdditionExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_additionExpression

	return p
}

func (s *AdditionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditionExpressionContext) MultiplicationExpression() IMultiplicationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicationExpressionContext)
}

func (s *AdditionExpressionContext) AdditionFactor() IAdditionFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditionFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditionFactorContext)
}

func (s *AdditionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) AdditionExpression() (localctx IAdditionExpressionContext) {
	localctx = NewAdditionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MonkeyParserRULE_additionExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.MultiplicationExpression()
	}
	{
		p.SetState(96)
		p.AdditionFactor()
	}

	return localctx
}

// IAdditionFactorContext is an interface to support dynamic dispatch.
type IAdditionFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditionFactorContext differentiates from other interfaces.
	IsAdditionFactorContext()
}

type AdditionFactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditionFactorContext() *AdditionFactorContext {
	var p = new(AdditionFactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_additionFactor
	return p
}

func (*AdditionFactorContext) IsAdditionFactorContext() {}

func NewAdditionFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditionFactorContext {
	var p = new(AdditionFactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_additionFactor

	return p
}

func (s *AdditionFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditionFactorContext) AllMultiplicationExpression() []IMultiplicationExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplicationExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplicationExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplicationExpressionContext)
		}
	}

	return tst
}

func (s *AdditionFactorContext) MultiplicationExpression(i int) IMultiplicationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicationExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplicationExpressionContext)
}

func (s *AdditionFactorContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserPLUS)
}

func (s *AdditionFactorContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserPLUS, i)
}

func (s *AdditionFactorContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserMINUS)
}

func (s *AdditionFactorContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserMINUS, i)
}

func (s *AdditionFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) AdditionFactor() (localctx IAdditionFactorContext) {
	localctx = NewAdditionFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MonkeyParserRULE_additionFactor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserPLUS || _la == MonkeyParserMINUS {
		{
			p.SetState(98)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserPLUS || _la == MonkeyParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(99)
			p.MultiplicationExpression()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplicationExpressionContext is an interface to support dynamic dispatch.
type IMultiplicationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicationExpressionContext differentiates from other interfaces.
	IsMultiplicationExpressionContext()
}

type MultiplicationExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicationExpressionContext() *MultiplicationExpressionContext {
	var p = new(MultiplicationExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_multiplicationExpression
	return p
}

func (*MultiplicationExpressionContext) IsMultiplicationExpressionContext() {}

func NewMultiplicationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicationExpressionContext {
	var p = new(MultiplicationExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_multiplicationExpression

	return p
}

func (s *MultiplicationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicationExpressionContext) ElementExpression() IElementExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementExpressionContext)
}

func (s *MultiplicationExpressionContext) MultiplicationFactor() IMultiplicationFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicationFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicationFactorContext)
}

func (s *MultiplicationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) MultiplicationExpression() (localctx IMultiplicationExpressionContext) {
	localctx = NewMultiplicationExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MonkeyParserRULE_multiplicationExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.ElementExpression()
	}
	{
		p.SetState(106)
		p.MultiplicationFactor()
	}

	return localctx
}

// IMultiplicationFactorContext is an interface to support dynamic dispatch.
type IMultiplicationFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicationFactorContext differentiates from other interfaces.
	IsMultiplicationFactorContext()
}

type MultiplicationFactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicationFactorContext() *MultiplicationFactorContext {
	var p = new(MultiplicationFactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_multiplicationFactor
	return p
}

func (*MultiplicationFactorContext) IsMultiplicationFactorContext() {}

func NewMultiplicationFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicationFactorContext {
	var p = new(MultiplicationFactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_multiplicationFactor

	return p
}

func (s *MultiplicationFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicationFactorContext) AllElementExpression() []IElementExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementExpressionContext)(nil)).Elem())
	var tst = make([]IElementExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicationFactorContext) ElementExpression(i int) IElementExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementExpressionContext)
}

func (s *MultiplicationFactorContext) AllMULT() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserMULT)
}

func (s *MultiplicationFactorContext) MULT(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserMULT, i)
}

func (s *MultiplicationFactorContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserDIV)
}

func (s *MultiplicationFactorContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserDIV, i)
}

func (s *MultiplicationFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) MultiplicationFactor() (localctx IMultiplicationFactorContext) {
	localctx = NewMultiplicationFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MonkeyParserRULE_multiplicationFactor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserMULT || _la == MonkeyParserDIV {
		{
			p.SetState(108)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserMULT || _la == MonkeyParserDIV) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(109)
			p.ElementExpression()
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IElementExpressionContext is an interface to support dynamic dispatch.
type IElementExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementExpressionContext differentiates from other interfaces.
	IsElementExpressionContext()
}

type ElementExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementExpressionContext() *ElementExpressionContext {
	var p = new(ElementExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_elementExpression
	return p
}

func (*ElementExpressionContext) IsElementExpressionContext() {}

func NewElementExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementExpressionContext {
	var p = new(ElementExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_elementExpression

	return p
}

func (s *ElementExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementExpressionContext) PrimitiveExpression() IPrimitiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveExpressionContext)
}

func (s *ElementExpressionContext) ElementAccess() IElementAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementAccessContext)
}

func (s *ElementExpressionContext) CallExpression() ICallExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallExpressionContext)
}

func (s *ElementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) ElementExpression() (localctx IElementExpressionContext) {
	localctx = NewElementExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MonkeyParserRULE_elementExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.PrimitiveExpression()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(116)
			p.ElementAccess()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(117)
			p.CallExpression()
		}

	}

	return localctx
}

// IElementAccessContext is an interface to support dynamic dispatch.
type IElementAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementAccessContext differentiates from other interfaces.
	IsElementAccessContext()
}

type ElementAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementAccessContext() *ElementAccessContext {
	var p = new(ElementAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_elementAccess
	return p
}

func (*ElementAccessContext) IsElementAccessContext() {}

func NewElementAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementAccessContext {
	var p = new(ElementAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_elementAccess

	return p
}

func (s *ElementAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementAccessContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_BRACKET, 0)
}

func (s *ElementAccessContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElementAccessContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_BRACKET, 0)
}

func (s *ElementAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) ElementAccess() (localctx IElementAccessContext) {
	localctx = NewElementAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MonkeyParserRULE_elementAccess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(MonkeyParserL_BRACKET)
	}
	{
		p.SetState(121)
		p.Expression()
	}
	{
		p.SetState(122)
		p.Match(MonkeyParserR_BRACKET)
	}

	return localctx
}

// ICallExpressionContext is an interface to support dynamic dispatch.
type ICallExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallExpressionContext differentiates from other interfaces.
	IsCallExpressionContext()
}

type CallExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallExpressionContext() *CallExpressionContext {
	var p = new(CallExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_callExpression
	return p
}

func (*CallExpressionContext) IsCallExpressionContext() {}

func NewCallExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallExpressionContext {
	var p = new(CallExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_callExpression

	return p
}

func (s *CallExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_PAREN, 0)
}

func (s *CallExpressionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *CallExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_PAREN, 0)
}

func (s *CallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) CallExpression() (localctx ICallExpressionContext) {
	localctx = NewCallExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MonkeyParserRULE_callExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(MonkeyParserL_PAREN)
	}
	{
		p.SetState(125)
		p.ExpressionList()
	}
	{
		p.SetState(126)
		p.Match(MonkeyParserR_PAREN)
	}

	return localctx
}

// IPrimitiveExpressionContext is an interface to support dynamic dispatch.
type IPrimitiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveExpressionContext differentiates from other interfaces.
	IsPrimitiveExpressionContext()
}

type PrimitiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveExpressionContext() *PrimitiveExpressionContext {
	var p = new(PrimitiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_primitiveExpression
	return p
}

func (*PrimitiveExpressionContext) IsPrimitiveExpressionContext() {}

func NewPrimitiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveExpressionContext {
	var p = new(PrimitiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_primitiveExpression

	return p
}

func (s *PrimitiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveExpressionContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserINTEGER, 0)
}

func (s *PrimitiveExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSTRING, 0)
}

func (s *PrimitiveExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENTIFIER, 0)
}

func (s *PrimitiveExpressionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserTRUE, 0)
}

func (s *PrimitiveExpressionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFALSE, 0)
}

func (s *PrimitiveExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_PAREN, 0)
}

func (s *PrimitiveExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimitiveExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_PAREN, 0)
}

func (s *PrimitiveExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *PrimitiveExpressionContext) ArrayFunctions() IArrayFunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayFunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayFunctionsContext)
}

func (s *PrimitiveExpressionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *PrimitiveExpressionContext) FunctionLiteral() IFunctionLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionLiteralContext)
}

func (s *PrimitiveExpressionContext) HashLiteral() IHashLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHashLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHashLiteralContext)
}

func (s *PrimitiveExpressionContext) PrintExpression() IPrintExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintExpressionContext)
}

func (s *PrimitiveExpressionContext) IfExpression() IIfExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExpressionContext)
}

func (s *PrimitiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) PrimitiveExpression() (localctx IPrimitiveExpressionContext) {
	localctx = NewPrimitiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MonkeyParserRULE_primitiveExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(MonkeyParserINTEGER)
		}

	case MonkeyParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(MonkeyParserSTRING)
		}

	case MonkeyParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Match(MonkeyParserIDENTIFIER)
		}

	case MonkeyParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.Match(MonkeyParserTRUE)
		}

	case MonkeyParserFALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
			p.Match(MonkeyParserFALSE)
		}

	case MonkeyParserL_PAREN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(133)
			p.Match(MonkeyParserL_PAREN)
		}
		{
			p.SetState(134)
			p.Expression()
		}
		{
			p.SetState(135)
			p.Match(MonkeyParserR_PAREN)
		}

	case MonkeyParserL_BRACKET:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(137)
			p.ArrayLiteral()
		}

	case MonkeyParserLEN, MonkeyParserFIRST, MonkeyParserLAST, MonkeyParserREST, MonkeyParserPUSH:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(138)
			p.ArrayFunctions()
		}
		{
			p.SetState(139)
			p.Match(MonkeyParserL_PAREN)
		}
		{
			p.SetState(140)
			p.ExpressionList()
		}
		{
			p.SetState(141)
			p.Match(MonkeyParserR_PAREN)
		}

	case MonkeyParserFUNC:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(143)
			p.FunctionLiteral()
		}

	case MonkeyParserL_CURLY:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(144)
			p.HashLiteral()
		}

	case MonkeyParserPUTS:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(145)
			p.PrintExpression()
		}

	case MonkeyParserIF:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(146)
			p.IfExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayFunctionsContext is an interface to support dynamic dispatch.
type IArrayFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayFunctionsContext differentiates from other interfaces.
	IsArrayFunctionsContext()
}

type ArrayFunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayFunctionsContext() *ArrayFunctionsContext {
	var p = new(ArrayFunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_arrayFunctions
	return p
}

func (*ArrayFunctionsContext) IsArrayFunctionsContext() {}

func NewArrayFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayFunctionsContext {
	var p = new(ArrayFunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_arrayFunctions

	return p
}

func (s *ArrayFunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayFunctionsContext) LEN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLEN, 0)
}

func (s *ArrayFunctionsContext) FIRST() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFIRST, 0)
}

func (s *ArrayFunctionsContext) LAST() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLAST, 0)
}

func (s *ArrayFunctionsContext) REST() antlr.TerminalNode {
	return s.GetToken(MonkeyParserREST, 0)
}

func (s *ArrayFunctionsContext) PUSH() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPUSH, 0)
}

func (s *ArrayFunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayFunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) ArrayFunctions() (localctx IArrayFunctionsContext) {
	localctx = NewArrayFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MonkeyParserRULE_arrayFunctions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-28)&-(0x1f+1)) == 0 && ((1<<uint((_la-28)))&((1<<(MonkeyParserLEN-28))|(1<<(MonkeyParserFIRST-28))|(1<<(MonkeyParserLAST-28))|(1<<(MonkeyParserREST-28))|(1<<(MonkeyParserPUSH-28)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_BRACKET, 0)
}

func (s *ArrayLiteralContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArrayLiteralContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_BRACKET, 0)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MonkeyParserRULE_arrayLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(MonkeyParserL_BRACKET)
	}
	{
		p.SetState(152)
		p.ExpressionList()
	}
	{
		p.SetState(153)
		p.Match(MonkeyParserR_BRACKET)
	}

	return localctx
}

// IFunctionLiteralContext is an interface to support dynamic dispatch.
type IFunctionLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionLiteralContext differentiates from other interfaces.
	IsFunctionLiteralContext()
}

type FunctionLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionLiteralContext() *FunctionLiteralContext {
	var p = new(FunctionLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_functionLiteral
	return p
}

func (*FunctionLiteralContext) IsFunctionLiteralContext() {}

func NewFunctionLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionLiteralContext {
	var p = new(FunctionLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_functionLiteral

	return p
}

func (s *FunctionLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionLiteralContext) FUNC() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFUNC, 0)
}

func (s *FunctionLiteralContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_PAREN, 0)
}

func (s *FunctionLiteralContext) FunctionParameters() IFunctionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *FunctionLiteralContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_PAREN, 0)
}

func (s *FunctionLiteralContext) BlockStatement() IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *FunctionLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) FunctionLiteral() (localctx IFunctionLiteralContext) {
	localctx = NewFunctionLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MonkeyParserRULE_functionLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(MonkeyParserFUNC)
	}
	{
		p.SetState(156)
		p.Match(MonkeyParserL_PAREN)
	}
	{
		p.SetState(157)
		p.FunctionParameters()
	}
	{
		p.SetState(158)
		p.Match(MonkeyParserR_PAREN)
	}
	{
		p.SetState(159)
		p.BlockStatement()
	}

	return localctx
}

// IFunctionParametersContext is an interface to support dynamic dispatch.
type IFunctionParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionParametersContext differentiates from other interfaces.
	IsFunctionParametersContext()
}

type FunctionParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParametersContext() *FunctionParametersContext {
	var p = new(FunctionParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_functionParameters
	return p
}

func (*FunctionParametersContext) IsFunctionParametersContext() {}

func NewFunctionParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParametersContext {
	var p = new(FunctionParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_functionParameters

	return p
}

func (s *FunctionParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParametersContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENTIFIER, 0)
}

func (s *FunctionParametersContext) MoreIdentifiers() IMoreIdentifiersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoreIdentifiersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMoreIdentifiersContext)
}

func (s *FunctionParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) FunctionParameters() (localctx IFunctionParametersContext) {
	localctx = NewFunctionParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MonkeyParserRULE_functionParameters)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(MonkeyParserIDENTIFIER)
	}
	{
		p.SetState(162)
		p.MoreIdentifiers()
	}

	return localctx
}

// IMoreIdentifiersContext is an interface to support dynamic dispatch.
type IMoreIdentifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoreIdentifiersContext differentiates from other interfaces.
	IsMoreIdentifiersContext()
}

type MoreIdentifiersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoreIdentifiersContext() *MoreIdentifiersContext {
	var p = new(MoreIdentifiersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_moreIdentifiers
	return p
}

func (*MoreIdentifiersContext) IsMoreIdentifiersContext() {}

func NewMoreIdentifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoreIdentifiersContext {
	var p = new(MoreIdentifiersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_moreIdentifiers

	return p
}

func (s *MoreIdentifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *MoreIdentifiersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *MoreIdentifiersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *MoreIdentifiersContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserIDENTIFIER)
}

func (s *MoreIdentifiersContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENTIFIER, i)
}

func (s *MoreIdentifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoreIdentifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) MoreIdentifiers() (localctx IMoreIdentifiersContext) {
	localctx = NewMoreIdentifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MonkeyParserRULE_moreIdentifiers)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserCOMMA {
		{
			p.SetState(164)
			p.Match(MonkeyParserCOMMA)
		}
		{
			p.SetState(165)
			p.Match(MonkeyParserIDENTIFIER)
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHashLiteralContext is an interface to support dynamic dispatch.
type IHashLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHashLiteralContext differentiates from other interfaces.
	IsHashLiteralContext()
}

type HashLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHashLiteralContext() *HashLiteralContext {
	var p = new(HashLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_hashLiteral
	return p
}

func (*HashLiteralContext) IsHashLiteralContext() {}

func NewHashLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HashLiteralContext {
	var p = new(HashLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_hashLiteral

	return p
}

func (s *HashLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *HashLiteralContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_CURLY, 0)
}

func (s *HashLiteralContext) HashContent() IHashContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHashContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHashContentContext)
}

func (s *HashLiteralContext) MoreHashContent() IMoreHashContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoreHashContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMoreHashContentContext)
}

func (s *HashLiteralContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_CURLY, 0)
}

func (s *HashLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) HashLiteral() (localctx IHashLiteralContext) {
	localctx = NewHashLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MonkeyParserRULE_hashLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(MonkeyParserL_CURLY)
	}
	{
		p.SetState(172)
		p.HashContent()
	}
	{
		p.SetState(173)
		p.MoreHashContent()
	}
	{
		p.SetState(174)
		p.Match(MonkeyParserR_CURLY)
	}

	return localctx
}

// IHashContentContext is an interface to support dynamic dispatch.
type IHashContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHashContentContext differentiates from other interfaces.
	IsHashContentContext()
}

type HashContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHashContentContext() *HashContentContext {
	var p = new(HashContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_hashContent
	return p
}

func (*HashContentContext) IsHashContentContext() {}

func NewHashContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HashContentContext {
	var p = new(HashContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_hashContent

	return p
}

func (s *HashContentContext) GetParser() antlr.Parser { return s.parser }

func (s *HashContentContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *HashContentContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HashContentContext) COLON() antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOLON, 0)
}

func (s *HashContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) HashContent() (localctx IHashContentContext) {
	localctx = NewHashContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MonkeyParserRULE_hashContent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Expression()
	}
	{
		p.SetState(177)
		p.Match(MonkeyParserCOLON)
	}
	{
		p.SetState(178)
		p.Expression()
	}

	return localctx
}

// IMoreHashContentContext is an interface to support dynamic dispatch.
type IMoreHashContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoreHashContentContext differentiates from other interfaces.
	IsMoreHashContentContext()
}

type MoreHashContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoreHashContentContext() *MoreHashContentContext {
	var p = new(MoreHashContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_moreHashContent
	return p
}

func (*MoreHashContentContext) IsMoreHashContentContext() {}

func NewMoreHashContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoreHashContentContext {
	var p = new(MoreHashContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_moreHashContent

	return p
}

func (s *MoreHashContentContext) GetParser() antlr.Parser { return s.parser }

func (s *MoreHashContentContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *MoreHashContentContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *MoreHashContentContext) AllHashContent() []IHashContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHashContentContext)(nil)).Elem())
	var tst = make([]IHashContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHashContentContext)
		}
	}

	return tst
}

func (s *MoreHashContentContext) HashContent(i int) IHashContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHashContentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHashContentContext)
}

func (s *MoreHashContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoreHashContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) MoreHashContent() (localctx IMoreHashContentContext) {
	localctx = NewMoreHashContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MonkeyParserRULE_moreHashContent)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserCOMMA {
		{
			p.SetState(180)
			p.Match(MonkeyParserCOMMA)
		}
		{
			p.SetState(181)
			p.HashContent()
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) MoreExpressions() IMoreExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoreExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMoreExpressionsContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MonkeyParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(MonkeyParserL_PAREN-4))|(1<<(MonkeyParserL_BRACKET-4))|(1<<(MonkeyParserL_CURLY-4))|(1<<(MonkeyParserIF-4))|(1<<(MonkeyParserTRUE-4))|(1<<(MonkeyParserFALSE-4))|(1<<(MonkeyParserPUTS-4))|(1<<(MonkeyParserFUNC-4))|(1<<(MonkeyParserLEN-4))|(1<<(MonkeyParserFIRST-4))|(1<<(MonkeyParserLAST-4))|(1<<(MonkeyParserREST-4))|(1<<(MonkeyParserPUSH-4))|(1<<(MonkeyParserIDENTIFIER-4))|(1<<(MonkeyParserINTEGER-4))|(1<<(MonkeyParserSTRING-4)))) != 0 {
		{
			p.SetState(187)
			p.Expression()
		}
		{
			p.SetState(188)
			p.MoreExpressions()
		}

	}

	return localctx
}

// IMoreExpressionsContext is an interface to support dynamic dispatch.
type IMoreExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoreExpressionsContext differentiates from other interfaces.
	IsMoreExpressionsContext()
}

type MoreExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoreExpressionsContext() *MoreExpressionsContext {
	var p = new(MoreExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_moreExpressions
	return p
}

func (*MoreExpressionsContext) IsMoreExpressionsContext() {}

func NewMoreExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoreExpressionsContext {
	var p = new(MoreExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_moreExpressions

	return p
}

func (s *MoreExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *MoreExpressionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *MoreExpressionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *MoreExpressionsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MoreExpressionsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MoreExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoreExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) MoreExpressions() (localctx IMoreExpressionsContext) {
	localctx = NewMoreExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MonkeyParserRULE_moreExpressions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserCOMMA {
		{
			p.SetState(192)
			p.Match(MonkeyParserCOMMA)
		}
		{
			p.SetState(193)
			p.Expression()
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPrintExpressionContext is an interface to support dynamic dispatch.
type IPrintExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintExpressionContext differentiates from other interfaces.
	IsPrintExpressionContext()
}

type PrintExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintExpressionContext() *PrintExpressionContext {
	var p = new(PrintExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_printExpression
	return p
}

func (*PrintExpressionContext) IsPrintExpressionContext() {}

func NewPrintExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintExpressionContext {
	var p = new(PrintExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_printExpression

	return p
}

func (s *PrintExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintExpressionContext) PUTS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPUTS, 0)
}

func (s *PrintExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_PAREN, 0)
}

func (s *PrintExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_PAREN, 0)
}

func (s *PrintExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) PrintExpression() (localctx IPrintExpressionContext) {
	localctx = NewPrintExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MonkeyParserRULE_printExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(MonkeyParserPUTS)
	}
	{
		p.SetState(200)
		p.Match(MonkeyParserL_PAREN)
	}
	{
		p.SetState(201)
		p.Expression()
	}
	{
		p.SetState(202)
		p.Match(MonkeyParserR_PAREN)
	}

	return localctx
}

// IIfExpressionContext is an interface to support dynamic dispatch.
type IIfExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfExpressionContext differentiates from other interfaces.
	IsIfExpressionContext()
}

type IfExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExpressionContext() *IfExpressionContext {
	var p = new(IfExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_ifExpression
	return p
}

func (*IfExpressionContext) IsIfExpressionContext() {}

func NewIfExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExpressionContext {
	var p = new(IfExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_ifExpression

	return p
}

func (s *IfExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExpressionContext) IF() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIF, 0)
}

func (s *IfExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfExpressionContext) AllBlockStatement() []IBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem())
	var tst = make([]IBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockStatementContext)
		}
	}

	return tst
}

func (s *IfExpressionContext) BlockStatement(i int) IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *IfExpressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserELSE, 0)
}

func (s *IfExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) IfExpression() (localctx IIfExpressionContext) {
	localctx = NewIfExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MonkeyParserRULE_ifExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(MonkeyParserIF)
	}
	{
		p.SetState(205)
		p.Expression()
	}
	{
		p.SetState(206)
		p.BlockStatement()
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MonkeyParserELSE {
		{
			p.SetState(207)
			p.Match(MonkeyParserELSE)
		}
		{
			p.SetState(208)
			p.BlockStatement()
		}

	}

	return localctx
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_blockStatement
	return p
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_CURLY, 0)
}

func (s *BlockStatementContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_CURLY, 0)
}

func (s *BlockStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MonkeyParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MonkeyParserRULE_blockStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(MonkeyParserL_CURLY)
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(MonkeyParserL_PAREN-4))|(1<<(MonkeyParserL_BRACKET-4))|(1<<(MonkeyParserL_CURLY-4))|(1<<(MonkeyParserIF-4))|(1<<(MonkeyParserLET-4))|(1<<(MonkeyParserRETURN-4))|(1<<(MonkeyParserTRUE-4))|(1<<(MonkeyParserFALSE-4))|(1<<(MonkeyParserPUTS-4))|(1<<(MonkeyParserFUNC-4))|(1<<(MonkeyParserLEN-4))|(1<<(MonkeyParserFIRST-4))|(1<<(MonkeyParserLAST-4))|(1<<(MonkeyParserREST-4))|(1<<(MonkeyParserPUSH-4))|(1<<(MonkeyParserIDENTIFIER-4))|(1<<(MonkeyParserINTEGER-4))|(1<<(MonkeyParserSTRING-4)))) != 0 {
		{
			p.SetState(212)
			p.Statement()
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(218)
		p.Match(MonkeyParserR_CURLY)
	}

	return localctx
}
