// Code generated from c:\Users\Emmanuel\source\CompiPR1\monkeycompiler\parser\MonkeyParser.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 220,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 3, 2, 7, 2, 60, 10, 2, 12, 2, 14, 2, 63, 11, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 5, 3, 70, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 76, 10,
	4, 3, 5, 3, 5, 5, 5, 80, 10, 5, 3, 6, 3, 6, 5, 6, 84, 10, 6, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 7, 8, 91, 10, 8, 12, 8, 14, 8, 94, 11, 8, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 7, 10, 101, 10, 10, 12, 10, 14, 10, 104, 11, 10, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 7, 12, 111, 10, 12, 12, 12, 14, 12, 114,
	11, 12, 3, 13, 3, 13, 3, 13, 5, 13, 119, 10, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 5, 16, 148, 10, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	21, 3, 21, 7, 21, 167, 10, 21, 12, 21, 14, 21, 170, 11, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 7, 24, 183,
	10, 24, 12, 24, 14, 24, 186, 11, 24, 3, 25, 3, 25, 5, 25, 190, 10, 25,
	3, 26, 3, 26, 7, 26, 194, 10, 26, 12, 26, 14, 26, 197, 11, 26, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 209,
	10, 28, 3, 29, 3, 29, 7, 29, 213, 10, 29, 12, 29, 14, 29, 216, 11, 29,
	3, 29, 3, 29, 3, 29, 2, 2, 30, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 2,
	6, 4, 2, 11, 11, 13, 16, 3, 2, 19, 20, 3, 2, 21, 22, 3, 2, 31, 35, 2, 219,
	2, 61, 3, 2, 2, 2, 4, 69, 3, 2, 2, 2, 6, 71, 3, 2, 2, 2, 8, 77, 3, 2, 2,
	2, 10, 81, 3, 2, 2, 2, 12, 85, 3, 2, 2, 2, 14, 92, 3, 2, 2, 2, 16, 95,
	3, 2, 2, 2, 18, 102, 3, 2, 2, 2, 20, 105, 3, 2, 2, 2, 22, 112, 3, 2, 2,
	2, 24, 115, 3, 2, 2, 2, 26, 120, 3, 2, 2, 2, 28, 124, 3, 2, 2, 2, 30, 147,
	3, 2, 2, 2, 32, 149, 3, 2, 2, 2, 34, 151, 3, 2, 2, 2, 36, 155, 3, 2, 2,
	2, 38, 161, 3, 2, 2, 2, 40, 168, 3, 2, 2, 2, 42, 171, 3, 2, 2, 2, 44, 176,
	3, 2, 2, 2, 46, 184, 3, 2, 2, 2, 48, 187, 3, 2, 2, 2, 50, 195, 3, 2, 2,
	2, 52, 198, 3, 2, 2, 2, 54, 203, 3, 2, 2, 2, 56, 210, 3, 2, 2, 2, 58, 60,
	5, 4, 3, 2, 59, 58, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2,
	61, 62, 3, 2, 2, 2, 62, 3, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65, 7, 25,
	2, 2, 65, 70, 5, 6, 4, 2, 66, 67, 7, 26, 2, 2, 67, 70, 5, 8, 5, 2, 68,
	70, 5, 10, 6, 2, 69, 64, 3, 2, 2, 2, 69, 66, 3, 2, 2, 2, 69, 68, 3, 2,
	2, 2, 70, 5, 3, 2, 2, 2, 71, 72, 7, 36, 2, 2, 72, 73, 7, 12, 2, 2, 73,
	75, 5, 12, 7, 2, 74, 76, 7, 3, 2, 2, 75, 74, 3, 2, 2, 2, 75, 76, 3, 2,
	2, 2, 76, 7, 3, 2, 2, 2, 77, 79, 5, 12, 7, 2, 78, 80, 7, 3, 2, 2, 79, 78,
	3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 9, 3, 2, 2, 2, 81, 83, 5, 12, 7, 2,
	82, 84, 7, 3, 2, 2, 83, 82, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 11, 3,
	2, 2, 2, 85, 86, 5, 16, 9, 2, 86, 87, 5, 14, 8, 2, 87, 13, 3, 2, 2, 2,
	88, 89, 9, 2, 2, 2, 89, 91, 5, 16, 9, 2, 90, 88, 3, 2, 2, 2, 91, 94, 3,
	2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 15, 3, 2, 2, 2, 94,
	92, 3, 2, 2, 2, 95, 96, 5, 20, 11, 2, 96, 97, 5, 18, 10, 2, 97, 17, 3,
	2, 2, 2, 98, 99, 9, 3, 2, 2, 99, 101, 5, 20, 11, 2, 100, 98, 3, 2, 2, 2,
	101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103,
	19, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 106, 5, 24, 13, 2, 106, 107,
	5, 22, 12, 2, 107, 21, 3, 2, 2, 2, 108, 109, 9, 4, 2, 2, 109, 111, 5, 24,
	13, 2, 110, 108, 3, 2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2,
	112, 113, 3, 2, 2, 2, 113, 23, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 118,
	5, 30, 16, 2, 116, 119, 5, 26, 14, 2, 117, 119, 5, 28, 15, 2, 118, 116,
	3, 2, 2, 2, 118, 117, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 25, 3, 2,
	2, 2, 120, 121, 7, 7, 2, 2, 121, 122, 5, 12, 7, 2, 122, 123, 7, 8, 2, 2,
	123, 27, 3, 2, 2, 2, 124, 125, 7, 5, 2, 2, 125, 126, 5, 48, 25, 2, 126,
	127, 7, 6, 2, 2, 127, 29, 3, 2, 2, 2, 128, 148, 7, 37, 2, 2, 129, 148,
	7, 38, 2, 2, 130, 148, 7, 36, 2, 2, 131, 148, 7, 27, 2, 2, 132, 148, 7,
	28, 2, 2, 133, 134, 7, 5, 2, 2, 134, 135, 5, 12, 7, 2, 135, 136, 7, 6,
	2, 2, 136, 148, 3, 2, 2, 2, 137, 148, 5, 34, 18, 2, 138, 139, 5, 32, 17,
	2, 139, 140, 7, 5, 2, 2, 140, 141, 5, 48, 25, 2, 141, 142, 7, 6, 2, 2,
	142, 148, 3, 2, 2, 2, 143, 148, 5, 36, 19, 2, 144, 148, 5, 42, 22, 2, 145,
	148, 5, 52, 27, 2, 146, 148, 5, 54, 28, 2, 147, 128, 3, 2, 2, 2, 147, 129,
	3, 2, 2, 2, 147, 130, 3, 2, 2, 2, 147, 131, 3, 2, 2, 2, 147, 132, 3, 2,
	2, 2, 147, 133, 3, 2, 2, 2, 147, 137, 3, 2, 2, 2, 147, 138, 3, 2, 2, 2,
	147, 143, 3, 2, 2, 2, 147, 144, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147,
	146, 3, 2, 2, 2, 148, 31, 3, 2, 2, 2, 149, 150, 9, 5, 2, 2, 150, 33, 3,
	2, 2, 2, 151, 152, 7, 7, 2, 2, 152, 153, 5, 48, 25, 2, 153, 154, 7, 8,
	2, 2, 154, 35, 3, 2, 2, 2, 155, 156, 7, 30, 2, 2, 156, 157, 7, 5, 2, 2,
	157, 158, 5, 38, 20, 2, 158, 159, 7, 6, 2, 2, 159, 160, 5, 56, 29, 2, 160,
	37, 3, 2, 2, 2, 161, 162, 7, 36, 2, 2, 162, 163, 5, 40, 21, 2, 163, 39,
	3, 2, 2, 2, 164, 165, 7, 17, 2, 2, 165, 167, 7, 36, 2, 2, 166, 164, 3,
	2, 2, 2, 167, 170, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2,
	2, 169, 41, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 171, 172, 7, 9, 2, 2, 172,
	173, 5, 44, 23, 2, 173, 174, 5, 46, 24, 2, 174, 175, 7, 10, 2, 2, 175,
	43, 3, 2, 2, 2, 176, 177, 5, 12, 7, 2, 177, 178, 7, 4, 2, 2, 178, 179,
	5, 12, 7, 2, 179, 45, 3, 2, 2, 2, 180, 181, 7, 17, 2, 2, 181, 183, 5, 44,
	23, 2, 182, 180, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2,
	184, 185, 3, 2, 2, 2, 185, 47, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 187, 189,
	5, 12, 7, 2, 188, 190, 5, 50, 26, 2, 189, 188, 3, 2, 2, 2, 189, 190, 3,
	2, 2, 2, 190, 49, 3, 2, 2, 2, 191, 192, 7, 17, 2, 2, 192, 194, 5, 12, 7,
	2, 193, 191, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 195,
	196, 3, 2, 2, 2, 196, 51, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 198, 199, 7,
	29, 2, 2, 199, 200, 7, 5, 2, 2, 200, 201, 5, 12, 7, 2, 201, 202, 7, 6,
	2, 2, 202, 53, 3, 2, 2, 2, 203, 204, 7, 23, 2, 2, 204, 205, 5, 12, 7, 2,
	205, 206, 5, 56, 29, 2, 206, 208, 7, 24, 2, 2, 207, 209, 5, 56, 29, 2,
	208, 207, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 55, 3, 2, 2, 2, 210, 214,
	7, 9, 2, 2, 211, 213, 5, 4, 3, 2, 212, 211, 3, 2, 2, 2, 213, 216, 3, 2,
	2, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 217, 3, 2, 2, 2,
	216, 214, 3, 2, 2, 2, 217, 218, 7, 10, 2, 2, 218, 57, 3, 2, 2, 2, 18, 61,
	69, 75, 79, 83, 92, 102, 112, 118, 147, 168, 184, 189, 195, 208, 214,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "':'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'=='", "'='",
	"'<='", "'>='", "'<'", "'>'", "','", "'_'", "'+'", "'-'", "'*'", "'/'",
	"'if'", "'else'", "'let'", "'return'", "'true'", "'false'", "'puts'", "'fn'",
	"'len'", "'first'", "'last'", "'rest'", "'push'",
}
var symbolicNames = []string{
	"", "PyCOMA", "DOSPUNTOS", "PIZQ", "PDER", "LBRACK", "RBRACK", "LBRACE",
	"RBRACE", "EQL", "ASSIGN", "LEQ", "GEQ", "LSS", "GTR", "COMMA", "GUIBAJO",
	"SUMA", "RESTA", "MULT", "DIV", "IF", "ELSE", "LET", "RETURN", "TRUE",
	"FALSE", "PUTS", "FN", "LEN", "FIRST", "LAST", "REST", "PUSH", "IDENT",
	"LITERAL", "STRING", "WS", "COMMENT", "TERMINATOR", "LINE_COMMENT",
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
	MonkeyParserEOF          = antlr.TokenEOF
	MonkeyParserPyCOMA       = 1
	MonkeyParserDOSPUNTOS    = 2
	MonkeyParserPIZQ         = 3
	MonkeyParserPDER         = 4
	MonkeyParserLBRACK       = 5
	MonkeyParserRBRACK       = 6
	MonkeyParserLBRACE       = 7
	MonkeyParserRBRACE       = 8
	MonkeyParserEQL          = 9
	MonkeyParserASSIGN       = 10
	MonkeyParserLEQ          = 11
	MonkeyParserGEQ          = 12
	MonkeyParserLSS          = 13
	MonkeyParserGTR          = 14
	MonkeyParserCOMMA        = 15
	MonkeyParserGUIBAJO      = 16
	MonkeyParserSUMA         = 17
	MonkeyParserRESTA        = 18
	MonkeyParserMULT         = 19
	MonkeyParserDIV          = 20
	MonkeyParserIF           = 21
	MonkeyParserELSE         = 22
	MonkeyParserLET          = 23
	MonkeyParserRETURN       = 24
	MonkeyParserTRUE         = 25
	MonkeyParserFALSE        = 26
	MonkeyParserPUTS         = 27
	MonkeyParserFN           = 28
	MonkeyParserLEN          = 29
	MonkeyParserFIRST        = 30
	MonkeyParserLAST         = 31
	MonkeyParserREST         = 32
	MonkeyParserPUSH         = 33
	MonkeyParserIDENT        = 34
	MonkeyParserLITERAL      = 35
	MonkeyParserSTRING       = 36
	MonkeyParserWS           = 37
	MonkeyParserCOMMENT      = 38
	MonkeyParserTERMINATOR   = 39
	MonkeyParserLINE_COMMENT = 40
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

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MonkeyParserPIZQ)|(1<<MonkeyParserLBRACK)|(1<<MonkeyParserLBRACE)|(1<<MonkeyParserIF)|(1<<MonkeyParserLET)|(1<<MonkeyParserRETURN)|(1<<MonkeyParserTRUE)|(1<<MonkeyParserFALSE)|(1<<MonkeyParserPUTS)|(1<<MonkeyParserFN)|(1<<MonkeyParserLEN)|(1<<MonkeyParserFIRST)|(1<<MonkeyParserLAST))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(MonkeyParserREST-32))|(1<<(MonkeyParserPUSH-32))|(1<<(MonkeyParserIDENT-32))|(1<<(MonkeyParserLITERAL-32))|(1<<(MonkeyParserSTRING-32)))) != 0) {
		{
			p.SetState(56)
			p.Statement()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(MonkeyParserLET)
		}
		{
			p.SetState(63)
			p.LetStatement()
		}

	case MonkeyParserRETURN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Match(MonkeyParserRETURN)
		}
		{
			p.SetState(65)
			p.ReturnStatement()
		}

	case MonkeyParserPIZQ, MonkeyParserLBRACK, MonkeyParserLBRACE, MonkeyParserIF, MonkeyParserTRUE, MonkeyParserFALSE, MonkeyParserPUTS, MonkeyParserFN, MonkeyParserLEN, MonkeyParserFIRST, MonkeyParserLAST, MonkeyParserREST, MonkeyParserPUSH, MonkeyParserIDENT, MonkeyParserLITERAL, MonkeyParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
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

func (s *LetStatementContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, 0)
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

func (s *LetStatementContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPyCOMA, 0)
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
		p.SetState(69)
		p.Match(MonkeyParserIDENT)
	}
	{
		p.SetState(70)
		p.Match(MonkeyParserASSIGN)
	}
	{
		p.SetState(71)
		p.Expression()
	}

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MonkeyParserPyCOMA {
		{
			p.SetState(72)
			p.Match(MonkeyParserPyCOMA)
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

func (s *ReturnStatementContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPyCOMA, 0)
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
		p.SetState(75)
		p.Expression()
	}

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MonkeyParserPyCOMA {
		{
			p.SetState(76)
			p.Match(MonkeyParserPyCOMA)
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

func (s *ExpressionStatementContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPyCOMA, 0)
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
		p.SetState(79)
		p.Expression()
	}

	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MonkeyParserPyCOMA {
		{
			p.SetState(80)
			p.Match(MonkeyParserPyCOMA)
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
		p.SetState(83)
		p.AdditionExpression()
	}
	{
		p.SetState(84)
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

func (s *ComparisonContext) AllLSS() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserLSS)
}

func (s *ComparisonContext) LSS(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserLSS, i)
}

func (s *ComparisonContext) AllGTR() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserGTR)
}

func (s *ComparisonContext) GTR(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserGTR, i)
}

func (s *ComparisonContext) AllLEQ() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserLEQ)
}

func (s *ComparisonContext) LEQ(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserLEQ, i)
}

func (s *ComparisonContext) AllGEQ() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserGEQ)
}

func (s *ComparisonContext) GEQ(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserGEQ, i)
}

func (s *ComparisonContext) AllEQL() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserEQL)
}

func (s *ComparisonContext) EQL(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserEQL, i)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MonkeyParserEQL)|(1<<MonkeyParserLEQ)|(1<<MonkeyParserGEQ)|(1<<MonkeyParserLSS)|(1<<MonkeyParserGTR))) != 0 {
		{
			p.SetState(86)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MonkeyParserEQL)|(1<<MonkeyParserLEQ)|(1<<MonkeyParserGEQ)|(1<<MonkeyParserLSS)|(1<<MonkeyParserGTR))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(87)
			p.AdditionExpression()
		}

		p.SetState(92)
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
		p.SetState(93)
		p.MultiplicationExpression()
	}
	{
		p.SetState(94)
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

func (s *AdditionFactorContext) AllSUMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserSUMA)
}

func (s *AdditionFactorContext) SUMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserSUMA, i)
}

func (s *AdditionFactorContext) AllRESTA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserRESTA)
}

func (s *AdditionFactorContext) RESTA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserRESTA, i)
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserSUMA || _la == MonkeyParserRESTA {
		{
			p.SetState(96)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserSUMA || _la == MonkeyParserRESTA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(97)
			p.MultiplicationExpression()
		}

		p.SetState(102)
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
		p.SetState(103)
		p.ElementExpression()
	}
	{
		p.SetState(104)
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
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserMULT || _la == MonkeyParserDIV {
		{
			p.SetState(106)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserMULT || _la == MonkeyParserDIV) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(107)
			p.ElementExpression()
		}

		p.SetState(112)
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
		p.SetState(113)
		p.PrimitiveExpression()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(114)
			p.ElementAccess()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(115)
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

func (s *ElementAccessContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACK, 0)
}

func (s *ElementAccessContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElementAccessContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACK, 0)
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
		p.SetState(118)
		p.Match(MonkeyParserLBRACK)
	}
	{
		p.SetState(119)
		p.Expression()
	}
	{
		p.SetState(120)
		p.Match(MonkeyParserRBRACK)
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

func (s *CallExpressionContext) PIZQ() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPIZQ, 0)
}

func (s *CallExpressionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *CallExpressionContext) PDER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPDER, 0)
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
		p.SetState(122)
		p.Match(MonkeyParserPIZQ)
	}
	{
		p.SetState(123)
		p.ExpressionList()
	}
	{
		p.SetState(124)
		p.Match(MonkeyParserPDER)
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

func (s *PrimitiveExpressionContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLITERAL, 0)
}

func (s *PrimitiveExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSTRING, 0)
}

func (s *PrimitiveExpressionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, 0)
}

func (s *PrimitiveExpressionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserTRUE, 0)
}

func (s *PrimitiveExpressionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFALSE, 0)
}

func (s *PrimitiveExpressionContext) PIZQ() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPIZQ, 0)
}

func (s *PrimitiveExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimitiveExpressionContext) PDER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPDER, 0)
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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserLITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Match(MonkeyParserLITERAL)
		}

	case MonkeyParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Match(MonkeyParserSTRING)
		}

	case MonkeyParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.Match(MonkeyParserIDENT)
		}

	case MonkeyParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(129)
			p.Match(MonkeyParserTRUE)
		}

	case MonkeyParserFALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(130)
			p.Match(MonkeyParserFALSE)
		}

	case MonkeyParserPIZQ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(131)
			p.Match(MonkeyParserPIZQ)
		}
		{
			p.SetState(132)
			p.Expression()
		}
		{
			p.SetState(133)
			p.Match(MonkeyParserPDER)
		}

	case MonkeyParserLBRACK:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(135)
			p.ArrayLiteral()
		}

	case MonkeyParserLEN, MonkeyParserFIRST, MonkeyParserLAST, MonkeyParserREST, MonkeyParserPUSH:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(136)
			p.ArrayFunctions()
		}
		{
			p.SetState(137)
			p.Match(MonkeyParserPIZQ)
		}
		{
			p.SetState(138)
			p.ExpressionList()
		}
		{
			p.SetState(139)
			p.Match(MonkeyParserPDER)
		}

	case MonkeyParserFN:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(141)
			p.FunctionLiteral()
		}

	case MonkeyParserLBRACE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(142)
			p.HashLiteral()
		}

	case MonkeyParserPUTS:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(143)
			p.PrintExpression()
		}

	case MonkeyParserIF:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(144)
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
		p.SetState(147)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(MonkeyParserLEN-29))|(1<<(MonkeyParserFIRST-29))|(1<<(MonkeyParserLAST-29))|(1<<(MonkeyParserREST-29))|(1<<(MonkeyParserPUSH-29)))) != 0) {
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

func (s *ArrayLiteralContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACK, 0)
}

func (s *ArrayLiteralContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArrayLiteralContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACK, 0)
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
		p.SetState(149)
		p.Match(MonkeyParserLBRACK)
	}
	{
		p.SetState(150)
		p.ExpressionList()
	}
	{
		p.SetState(151)
		p.Match(MonkeyParserRBRACK)
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

func (s *FunctionLiteralContext) FN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFN, 0)
}

func (s *FunctionLiteralContext) PIZQ() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPIZQ, 0)
}

func (s *FunctionLiteralContext) FunctionParameters() IFunctionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *FunctionLiteralContext) PDER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPDER, 0)
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
		p.SetState(153)
		p.Match(MonkeyParserFN)
	}
	{
		p.SetState(154)
		p.Match(MonkeyParserPIZQ)
	}
	{
		p.SetState(155)
		p.FunctionParameters()
	}
	{
		p.SetState(156)
		p.Match(MonkeyParserPDER)
	}
	{
		p.SetState(157)
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

func (s *FunctionParametersContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, 0)
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
		p.SetState(159)
		p.Match(MonkeyParserIDENT)
	}
	{
		p.SetState(160)
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

func (s *MoreIdentifiersContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserIDENT)
}

func (s *MoreIdentifiersContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, i)
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
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserCOMMA {
		{
			p.SetState(162)
			p.Match(MonkeyParserCOMMA)
		}
		{
			p.SetState(163)
			p.Match(MonkeyParserIDENT)
		}

		p.SetState(168)
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

func (s *HashLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACE, 0)
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

func (s *HashLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACE, 0)
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
		p.SetState(169)
		p.Match(MonkeyParserLBRACE)
	}
	{
		p.SetState(170)
		p.HashContent()
	}
	{
		p.SetState(171)
		p.MoreHashContent()
	}
	{
		p.SetState(172)
		p.Match(MonkeyParserRBRACE)
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

func (s *HashContentContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserDOSPUNTOS, 0)
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
		p.SetState(174)
		p.Expression()
	}
	{
		p.SetState(175)
		p.Match(MonkeyParserDOSPUNTOS)
	}
	{
		p.SetState(176)
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
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserCOMMA {
		{
			p.SetState(178)
			p.Match(MonkeyParserCOMMA)
		}
		{
			p.SetState(179)
			p.HashContent()
		}

		p.SetState(184)
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
		p.SetState(185)
		p.Expression()
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(186)
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
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserCOMMA {
		{
			p.SetState(189)
			p.Match(MonkeyParserCOMMA)
		}
		{
			p.SetState(190)
			p.Expression()
		}

		p.SetState(195)
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

func (s *PrintExpressionContext) PIZQ() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPIZQ, 0)
}

func (s *PrintExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintExpressionContext) PDER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPDER, 0)
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
		p.SetState(196)
		p.Match(MonkeyParserPUTS)
	}
	{
		p.SetState(197)
		p.Match(MonkeyParserPIZQ)
	}
	{
		p.SetState(198)
		p.Expression()
	}
	{
		p.SetState(199)
		p.Match(MonkeyParserPDER)
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
		p.SetState(201)
		p.Match(MonkeyParserIF)
	}
	{
		p.SetState(202)
		p.Expression()
	}
	{
		p.SetState(203)
		p.BlockStatement()
	}

	{
		p.SetState(204)
		p.Match(MonkeyParserELSE)
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(205)
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

func (s *BlockStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACE, 0)
}

func (s *BlockStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACE, 0)
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
		p.SetState(208)
		p.Match(MonkeyParserLBRACE)
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MonkeyParserPIZQ)|(1<<MonkeyParserLBRACK)|(1<<MonkeyParserLBRACE)|(1<<MonkeyParserIF)|(1<<MonkeyParserLET)|(1<<MonkeyParserRETURN)|(1<<MonkeyParserTRUE)|(1<<MonkeyParserFALSE)|(1<<MonkeyParserPUTS)|(1<<MonkeyParserFN)|(1<<MonkeyParserLEN)|(1<<MonkeyParserFIRST)|(1<<MonkeyParserLAST))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(MonkeyParserREST-32))|(1<<(MonkeyParserPUSH-32))|(1<<(MonkeyParserIDENT-32))|(1<<(MonkeyParserLITERAL-32))|(1<<(MonkeyParserSTRING-32)))) != 0) {
		{
			p.SetState(209)
			p.Statement()
		}

		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(215)
		p.Match(MonkeyParserRBRACE)
	}

	return localctx
}
