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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 216,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 7, 2, 42, 10, 2, 12, 2, 14, 2, 45,
	11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 54, 10, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 59, 10, 3, 3, 3, 3, 3, 5, 3, 63, 10, 3, 5, 3, 65, 10,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 71, 10, 4, 12, 4, 14, 4, 74, 11, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 81, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7,
	6, 87, 10, 6, 12, 6, 14, 6, 90, 11, 6, 3, 7, 3, 7, 5, 7, 94, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 7, 8, 100, 10, 8, 12, 8, 14, 8, 103, 11, 8, 3, 9,
	3, 9, 5, 9, 107, 10, 9, 3, 10, 3, 10, 3, 10, 5, 10, 112, 10, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 5, 12, 120, 10, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	5, 13, 135, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 141, 10, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 7, 13, 155, 10, 13, 12, 13, 14, 13, 158, 11, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13,
	172, 10, 13, 5, 13, 174, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5,
	14, 181, 10, 14, 3, 15, 3, 15, 3, 15, 7, 15, 186, 10, 15, 12, 15, 14, 15,
	189, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 7, 17, 198,
	10, 17, 12, 17, 14, 17, 201, 11, 17, 3, 18, 3, 18, 7, 18, 205, 10, 18,
	12, 18, 14, 18, 208, 11, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 2, 2, 21, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 2, 2, 2, 236, 2, 43, 3, 2, 2, 2, 4, 64, 3, 2, 2, 2, 6,
	66, 3, 2, 2, 2, 8, 80, 3, 2, 2, 2, 10, 82, 3, 2, 2, 2, 12, 93, 3, 2, 2,
	2, 14, 95, 3, 2, 2, 2, 16, 106, 3, 2, 2, 2, 18, 108, 3, 2, 2, 2, 20, 113,
	3, 2, 2, 2, 22, 117, 3, 2, 2, 2, 24, 173, 3, 2, 2, 2, 26, 180, 3, 2, 2,
	2, 28, 182, 3, 2, 2, 2, 30, 190, 3, 2, 2, 2, 32, 194, 3, 2, 2, 2, 34, 202,
	3, 2, 2, 2, 36, 211, 3, 2, 2, 2, 38, 213, 3, 2, 2, 2, 40, 42, 5, 4, 3,
	2, 41, 40, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44,
	3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 47, 7, 2, 2, 3,
	47, 3, 3, 2, 2, 2, 48, 49, 7, 24, 2, 2, 49, 50, 7, 35, 2, 2, 50, 51, 7,
	13, 2, 2, 51, 53, 5, 6, 4, 2, 52, 54, 7, 3, 2, 2, 53, 52, 3, 2, 2, 2, 53,
	54, 3, 2, 2, 2, 54, 65, 3, 2, 2, 2, 55, 56, 7, 25, 2, 2, 56, 58, 5, 6,
	4, 2, 57, 59, 7, 3, 2, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 65,
	3, 2, 2, 2, 60, 62, 5, 6, 4, 2, 61, 63, 7, 3, 2, 2, 62, 61, 3, 2, 2, 2,
	62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 48, 3, 2, 2, 2, 64, 55, 3,
	2, 2, 2, 64, 60, 3, 2, 2, 2, 65, 5, 3, 2, 2, 2, 66, 72, 5, 10, 6, 2, 67,
	68, 5, 8, 5, 2, 68, 69, 5, 10, 6, 2, 69, 71, 3, 2, 2, 2, 70, 67, 3, 2,
	2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 7,
	3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 81, 7, 16, 2, 2, 76, 81, 7, 17, 2,
	2, 77, 81, 7, 14, 2, 2, 78, 81, 7, 15, 2, 2, 79, 81, 7, 12, 2, 2, 80, 75,
	3, 2, 2, 2, 80, 76, 3, 2, 2, 2, 80, 77, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2,
	80, 79, 3, 2, 2, 2, 81, 9, 3, 2, 2, 2, 82, 88, 5, 14, 8, 2, 83, 84, 5,
	12, 7, 2, 84, 85, 5, 14, 8, 2, 85, 87, 3, 2, 2, 2, 86, 83, 3, 2, 2, 2,
	87, 90, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 11, 3,
	2, 2, 2, 90, 88, 3, 2, 2, 2, 91, 94, 7, 18, 2, 2, 92, 94, 7, 19, 2, 2,
	93, 91, 3, 2, 2, 2, 93, 92, 3, 2, 2, 2, 94, 13, 3, 2, 2, 2, 95, 101, 5,
	18, 10, 2, 96, 97, 5, 16, 9, 2, 97, 98, 5, 18, 10, 2, 98, 100, 3, 2, 2,
	2, 99, 96, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101,
	102, 3, 2, 2, 2, 102, 15, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104, 107, 7,
	20, 2, 2, 105, 107, 7, 21, 2, 2, 106, 104, 3, 2, 2, 2, 106, 105, 3, 2,
	2, 2, 107, 17, 3, 2, 2, 2, 108, 111, 5, 24, 13, 2, 109, 112, 5, 20, 11,
	2, 110, 112, 5, 22, 12, 2, 111, 109, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2,
	111, 112, 3, 2, 2, 2, 112, 19, 3, 2, 2, 2, 113, 114, 7, 8, 2, 2, 114, 115,
	5, 6, 4, 2, 115, 116, 7, 9, 2, 2, 116, 21, 3, 2, 2, 2, 117, 119, 7, 6,
	2, 2, 118, 120, 5, 32, 17, 2, 119, 118, 3, 2, 2, 2, 119, 120, 3, 2, 2,
	2, 120, 121, 3, 2, 2, 2, 121, 122, 7, 7, 2, 2, 122, 23, 3, 2, 2, 2, 123,
	174, 7, 36, 2, 2, 124, 174, 7, 37, 2, 2, 125, 174, 5, 36, 19, 2, 126, 174,
	7, 26, 2, 2, 127, 174, 7, 27, 2, 2, 128, 129, 7, 6, 2, 2, 129, 130, 5,
	6, 4, 2, 130, 131, 7, 7, 2, 2, 131, 174, 3, 2, 2, 2, 132, 134, 7, 8, 2,
	2, 133, 135, 5, 32, 17, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2,
	135, 136, 3, 2, 2, 2, 136, 174, 7, 9, 2, 2, 137, 138, 5, 26, 14, 2, 138,
	140, 7, 6, 2, 2, 139, 141, 5, 32, 17, 2, 140, 139, 3, 2, 2, 2, 140, 141,
	3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 7, 7, 2, 2, 143, 174, 3, 2,
	2, 2, 144, 145, 5, 38, 20, 2, 145, 146, 7, 6, 2, 2, 146, 147, 5, 28, 15,
	2, 147, 148, 7, 7, 2, 2, 148, 149, 5, 34, 18, 2, 149, 174, 3, 2, 2, 2,
	150, 151, 7, 10, 2, 2, 151, 156, 5, 30, 16, 2, 152, 153, 7, 5, 2, 2, 153,
	155, 5, 30, 16, 2, 154, 152, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154,
	3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 159, 3, 2, 2, 2, 158, 156, 3, 2,
	2, 2, 159, 160, 7, 11, 2, 2, 160, 174, 3, 2, 2, 2, 161, 162, 7, 28, 2,
	2, 162, 163, 7, 6, 2, 2, 163, 164, 5, 6, 4, 2, 164, 165, 7, 7, 2, 2, 165,
	174, 3, 2, 2, 2, 166, 167, 7, 22, 2, 2, 167, 168, 5, 6, 4, 2, 168, 171,
	5, 34, 18, 2, 169, 170, 7, 23, 2, 2, 170, 172, 5, 34, 18, 2, 171, 169,
	3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 174, 3, 2, 2, 2, 173, 123, 3, 2,
	2, 2, 173, 124, 3, 2, 2, 2, 173, 125, 3, 2, 2, 2, 173, 126, 3, 2, 2, 2,
	173, 127, 3, 2, 2, 2, 173, 128, 3, 2, 2, 2, 173, 132, 3, 2, 2, 2, 173,
	137, 3, 2, 2, 2, 173, 144, 3, 2, 2, 2, 173, 150, 3, 2, 2, 2, 173, 161,
	3, 2, 2, 2, 173, 166, 3, 2, 2, 2, 174, 25, 3, 2, 2, 2, 175, 181, 7, 30,
	2, 2, 176, 181, 7, 31, 2, 2, 177, 181, 7, 32, 2, 2, 178, 181, 7, 33, 2,
	2, 179, 181, 7, 34, 2, 2, 180, 175, 3, 2, 2, 2, 180, 176, 3, 2, 2, 2, 180,
	177, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 27, 3,
	2, 2, 2, 182, 187, 7, 35, 2, 2, 183, 184, 7, 5, 2, 2, 184, 186, 7, 35,
	2, 2, 185, 183, 3, 2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2,
	187, 188, 3, 2, 2, 2, 188, 29, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 191,
	5, 6, 4, 2, 191, 192, 7, 4, 2, 2, 192, 193, 5, 6, 4, 2, 193, 31, 3, 2,
	2, 2, 194, 199, 5, 6, 4, 2, 195, 196, 7, 5, 2, 2, 196, 198, 5, 6, 4, 2,
	197, 195, 3, 2, 2, 2, 198, 201, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199,
	200, 3, 2, 2, 2, 200, 33, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202, 206, 7,
	10, 2, 2, 203, 205, 5, 4, 3, 2, 204, 203, 3, 2, 2, 2, 205, 208, 3, 2, 2,
	2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 209, 3, 2, 2, 2, 208,
	206, 3, 2, 2, 2, 209, 210, 7, 11, 2, 2, 210, 35, 3, 2, 2, 2, 211, 212,
	7, 35, 2, 2, 212, 37, 3, 2, 2, 2, 213, 214, 7, 29, 2, 2, 214, 39, 3, 2,
	2, 2, 24, 43, 53, 58, 62, 64, 72, 80, 88, 93, 101, 106, 111, 119, 134,
	140, 156, 171, 173, 180, 187, 199, 206,
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
	"program", "statement", "expression", "comparison", "additionExpression",
	"additionFactor", "multiplicationExpression", "multiplicationFactor", "elementExpression",
	"elementAccess", "callExpression", "primitiveExpression", "arrayFunctions",
	"functionParameters", "hashContent", "expressionList", "blockStatement",
	"identifier", "function",
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
	MonkeyParserRULE_expression               = 2
	MonkeyParserRULE_comparison               = 3
	MonkeyParserRULE_additionExpression       = 4
	MonkeyParserRULE_additionFactor           = 5
	MonkeyParserRULE_multiplicationExpression = 6
	MonkeyParserRULE_multiplicationFactor     = 7
	MonkeyParserRULE_elementExpression        = 8
	MonkeyParserRULE_elementAccess            = 9
	MonkeyParserRULE_callExpression           = 10
	MonkeyParserRULE_primitiveExpression      = 11
	MonkeyParserRULE_arrayFunctions           = 12
	MonkeyParserRULE_functionParameters       = 13
	MonkeyParserRULE_hashContent              = 14
	MonkeyParserRULE_expressionList           = 15
	MonkeyParserRULE_blockStatement           = 16
	MonkeyParserRULE_identifier               = 17
	MonkeyParserRULE_function                 = 18
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

func (s *ProgramContext) CopyFrom(ctx *ProgramContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ProgramTreeContext struct {
	*ProgramContext
}

func NewProgramTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProgramTreeContext {
	var p = new(ProgramTreeContext)

	p.ProgramContext = NewEmptyProgramContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ProgramContext))

	return p
}

func (s *ProgramTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramTreeContext) EOF() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEOF, 0)
}

func (s *ProgramTreeContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramTreeContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitProgramTree(s)

	default:
		return t.VisitChildren(s)
	}
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

	localctx = NewProgramTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(MonkeyParserL_PAREN-4))|(1<<(MonkeyParserL_BRACKET-4))|(1<<(MonkeyParserL_CURLY-4))|(1<<(MonkeyParserIF-4))|(1<<(MonkeyParserLET-4))|(1<<(MonkeyParserRETURN-4))|(1<<(MonkeyParserTRUE-4))|(1<<(MonkeyParserFALSE-4))|(1<<(MonkeyParserPUTS-4))|(1<<(MonkeyParserFUNC-4))|(1<<(MonkeyParserLEN-4))|(1<<(MonkeyParserFIRST-4))|(1<<(MonkeyParserLAST-4))|(1<<(MonkeyParserREST-4))|(1<<(MonkeyParserPUSH-4))|(1<<(MonkeyParserIDENTIFIER-4))|(1<<(MonkeyParserINTEGER-4))|(1<<(MonkeyParserSTRING-4)))) != 0 {
		{
			p.SetState(38)
			p.Statement()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
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

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LetStatementTreeContext struct {
	*StatementContext
}

func NewLetStatementTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetStatementTreeContext {
	var p = new(LetStatementTreeContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *LetStatementTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetStatementTreeContext) LET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLET, 0)
}

func (s *LetStatementTreeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENTIFIER, 0)
}

func (s *LetStatementTreeContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserASSIGN, 0)
}

func (s *LetStatementTreeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LetStatementTreeContext) SEMI() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSEMI, 0)
}

func (s *LetStatementTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitLetStatementTree(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionStatementTreeContext struct {
	*StatementContext
}

func NewExpressionStatementTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionStatementTreeContext {
	var p = new(ExpressionStatementTreeContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ExpressionStatementTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementTreeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementTreeContext) SEMI() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSEMI, 0)
}

func (s *ExpressionStatementTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitExpressionStatementTree(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementTreeContext struct {
	*StatementContext
}

func NewReturnStatementTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementTreeContext {
	var p = new(ReturnStatementTreeContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnStatementTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementTreeContext) RETURN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRETURN, 0)
}

func (s *ReturnStatementTreeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementTreeContext) SEMI() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSEMI, 0)
}

func (s *ReturnStatementTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitReturnStatementTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MonkeyParserRULE_statement)
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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserLET:
		localctx = NewLetStatementTreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(MonkeyParserLET)
		}
		{
			p.SetState(47)
			p.Match(MonkeyParserIDENTIFIER)
		}
		{
			p.SetState(48)
			p.Match(MonkeyParserASSIGN)
		}
		{
			p.SetState(49)
			p.Expression()
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MonkeyParserSEMI {
			{
				p.SetState(50)
				p.Match(MonkeyParserSEMI)
			}

		}

	case MonkeyParserRETURN:
		localctx = NewReturnStatementTreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(MonkeyParserRETURN)
		}
		{
			p.SetState(54)
			p.Expression()
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MonkeyParserSEMI {
			{
				p.SetState(55)
				p.Match(MonkeyParserSEMI)
			}

		}

	case MonkeyParserL_PAREN, MonkeyParserL_BRACKET, MonkeyParserL_CURLY, MonkeyParserIF, MonkeyParserTRUE, MonkeyParserFALSE, MonkeyParserPUTS, MonkeyParserFUNC, MonkeyParserLEN, MonkeyParserFIRST, MonkeyParserLAST, MonkeyParserREST, MonkeyParserPUSH, MonkeyParserIDENTIFIER, MonkeyParserINTEGER, MonkeyParserSTRING:
		localctx = NewExpressionStatementTreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.Expression()
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MonkeyParserSEMI {
			{
				p.SetState(59)
				p.Match(MonkeyParserSEMI)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionTreeContext struct {
	*ExpressionContext
}

func NewExpressionTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionTreeContext {
	var p = new(ExpressionTreeContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionTreeContext) AllAdditionExpression() []IAdditionExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdditionExpressionContext)(nil)).Elem())
	var tst = make([]IAdditionExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdditionExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionTreeContext) AdditionExpression(i int) IAdditionExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditionExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdditionExpressionContext)
}

func (s *ExpressionTreeContext) AllComparison() []IComparisonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComparisonContext)(nil)).Elem())
	var tst = make([]IComparisonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComparisonContext)
		}
	}

	return tst
}

func (s *ExpressionTreeContext) Comparison(i int) IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *ExpressionTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitExpressionTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MonkeyParserRULE_expression)
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

	localctx = NewExpressionTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.AdditionExpression()
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MonkeyParserEQUALS)|(1<<MonkeyParserLESS_OR_EQUALS)|(1<<MonkeyParserGREATER_OR_EQUALS)|(1<<MonkeyParserLESS)|(1<<MonkeyParserGREATER))) != 0 {
		{
			p.SetState(65)
			p.Comparison()
		}
		{
			p.SetState(66)
			p.AdditionExpression()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *ComparisonContext) CopyFrom(ctx *ComparisonContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GreaterComparisonContext struct {
	*ComparisonContext
}

func NewGreaterComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterComparisonContext {
	var p = new(GreaterComparisonContext)

	p.ComparisonContext = NewEmptyComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonContext))

	return p
}

func (s *GreaterComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterComparisonContext) GREATER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserGREATER, 0)
}

func (s *GreaterComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitGreaterComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type GreaterOrEqualsComparisonContext struct {
	*ComparisonContext
}

func NewGreaterOrEqualsComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterOrEqualsComparisonContext {
	var p = new(GreaterOrEqualsComparisonContext)

	p.ComparisonContext = NewEmptyComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonContext))

	return p
}

func (s *GreaterOrEqualsComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterOrEqualsComparisonContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserGREATER_OR_EQUALS, 0)
}

func (s *GreaterOrEqualsComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitGreaterOrEqualsComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessOrEqualsComparisonContext struct {
	*ComparisonContext
}

func NewLessOrEqualsComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessOrEqualsComparisonContext {
	var p = new(LessOrEqualsComparisonContext)

	p.ComparisonContext = NewEmptyComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonContext))

	return p
}

func (s *LessOrEqualsComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessOrEqualsComparisonContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLESS_OR_EQUALS, 0)
}

func (s *LessOrEqualsComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitLessOrEqualsComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessComparisonContext struct {
	*ComparisonContext
}

func NewLessComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessComparisonContext {
	var p = new(LessComparisonContext)

	p.ComparisonContext = NewEmptyComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonContext))

	return p
}

func (s *LessComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessComparisonContext) LESS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLESS, 0)
}

func (s *LessComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitLessComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualsComparisonContext struct {
	*ComparisonContext
}

func NewEqualsComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualsComparisonContext {
	var p = new(EqualsComparisonContext)

	p.ComparisonContext = NewEmptyComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonContext))

	return p
}

func (s *EqualsComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualsComparisonContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEQUALS, 0)
}

func (s *EqualsComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitEqualsComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MonkeyParserRULE_comparison)

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

	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserLESS:
		localctx = NewLessComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Match(MonkeyParserLESS)
		}

	case MonkeyParserGREATER:
		localctx = NewGreaterComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Match(MonkeyParserGREATER)
		}

	case MonkeyParserLESS_OR_EQUALS:
		localctx = NewLessOrEqualsComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.Match(MonkeyParserLESS_OR_EQUALS)
		}

	case MonkeyParserGREATER_OR_EQUALS:
		localctx = NewGreaterOrEqualsComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.Match(MonkeyParserGREATER_OR_EQUALS)
		}

	case MonkeyParserEQUALS:
		localctx = NewEqualsComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(77)
			p.Match(MonkeyParserEQUALS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *AdditionExpressionContext) CopyFrom(ctx *AdditionExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AdditionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AdditionTreeContext struct {
	*AdditionExpressionContext
}

func NewAdditionTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditionTreeContext {
	var p = new(AdditionTreeContext)

	p.AdditionExpressionContext = NewEmptyAdditionExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AdditionExpressionContext))

	return p
}

func (s *AdditionTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionTreeContext) AllMultiplicationExpression() []IMultiplicationExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplicationExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplicationExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplicationExpressionContext)
		}
	}

	return tst
}

func (s *AdditionTreeContext) MultiplicationExpression(i int) IMultiplicationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicationExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplicationExpressionContext)
}

func (s *AdditionTreeContext) AllAdditionFactor() []IAdditionFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdditionFactorContext)(nil)).Elem())
	var tst = make([]IAdditionFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdditionFactorContext)
		}
	}

	return tst
}

func (s *AdditionTreeContext) AdditionFactor(i int) IAdditionFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditionFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdditionFactorContext)
}

func (s *AdditionTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitAdditionTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) AdditionExpression() (localctx IAdditionExpressionContext) {
	localctx = NewAdditionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MonkeyParserRULE_additionExpression)
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

	localctx = NewAdditionTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.MultiplicationExpression()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserPLUS || _la == MonkeyParserMINUS {
		{
			p.SetState(81)
			p.AdditionFactor()
		}
		{
			p.SetState(82)
			p.MultiplicationExpression()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *AdditionFactorContext) CopyFrom(ctx *AdditionFactorContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AdditionFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PlusOperatorContext struct {
	*AdditionFactorContext
}

func NewPlusOperatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusOperatorContext {
	var p = new(PlusOperatorContext)

	p.AdditionFactorContext = NewEmptyAdditionFactorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AdditionFactorContext))

	return p
}

func (s *PlusOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPLUS, 0)
}

func (s *PlusOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitPlusOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

type MinusOperatorContext struct {
	*AdditionFactorContext
}

func NewMinusOperatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusOperatorContext {
	var p = new(MinusOperatorContext)

	p.AdditionFactorContext = NewEmptyAdditionFactorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AdditionFactorContext))

	return p
}

func (s *MinusOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserMINUS, 0)
}

func (s *MinusOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitMinusOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) AdditionFactor() (localctx IAdditionFactorContext) {
	localctx = NewAdditionFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MonkeyParserRULE_additionFactor)

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

	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserPLUS:
		localctx = NewPlusOperatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Match(MonkeyParserPLUS)
		}

	case MonkeyParserMINUS:
		localctx = NewMinusOperatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(MonkeyParserMINUS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *MultiplicationExpressionContext) CopyFrom(ctx *MultiplicationExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MultiplicationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultiplicationTreeContext struct {
	*MultiplicationExpressionContext
}

func NewMultiplicationTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationTreeContext {
	var p = new(MultiplicationTreeContext)

	p.MultiplicationExpressionContext = NewEmptyMultiplicationExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultiplicationExpressionContext))

	return p
}

func (s *MultiplicationTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationTreeContext) AllElementExpression() []IElementExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementExpressionContext)(nil)).Elem())
	var tst = make([]IElementExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicationTreeContext) ElementExpression(i int) IElementExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementExpressionContext)
}

func (s *MultiplicationTreeContext) AllMultiplicationFactor() []IMultiplicationFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplicationFactorContext)(nil)).Elem())
	var tst = make([]IMultiplicationFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplicationFactorContext)
		}
	}

	return tst
}

func (s *MultiplicationTreeContext) MultiplicationFactor(i int) IMultiplicationFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicationFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplicationFactorContext)
}

func (s *MultiplicationTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitMultiplicationTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) MultiplicationExpression() (localctx IMultiplicationExpressionContext) {
	localctx = NewMultiplicationExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MonkeyParserRULE_multiplicationExpression)
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

	localctx = NewMultiplicationTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.ElementExpression()
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserMULT || _la == MonkeyParserDIV {
		{
			p.SetState(94)
			p.MultiplicationFactor()
		}
		{
			p.SetState(95)
			p.ElementExpression()
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *MultiplicationFactorContext) CopyFrom(ctx *MultiplicationFactorContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MultiplicationFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DivisionOperatorContext struct {
	*MultiplicationFactorContext
}

func NewDivisionOperatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivisionOperatorContext {
	var p = new(DivisionOperatorContext)

	p.MultiplicationFactorContext = NewEmptyMultiplicationFactorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultiplicationFactorContext))

	return p
}

func (s *DivisionOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivisionOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(MonkeyParserDIV, 0)
}

func (s *DivisionOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitDivisionOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicationOperatorContext struct {
	*MultiplicationFactorContext
}

func NewMultiplicationOperatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationOperatorContext {
	var p = new(MultiplicationOperatorContext)

	p.MultiplicationFactorContext = NewEmptyMultiplicationFactorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultiplicationFactorContext))

	return p
}

func (s *MultiplicationOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationOperatorContext) MULT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserMULT, 0)
}

func (s *MultiplicationOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitMultiplicationOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) MultiplicationFactor() (localctx IMultiplicationFactorContext) {
	localctx = NewMultiplicationFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MonkeyParserRULE_multiplicationFactor)

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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserMULT:
		localctx = NewMultiplicationOperatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Match(MonkeyParserMULT)
		}

	case MonkeyParserDIV:
		localctx = NewDivisionOperatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(MonkeyParserDIV)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *ElementExpressionContext) CopyFrom(ctx *ElementExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ElementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElementTreeContext struct {
	*ElementExpressionContext
}

func NewElementTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElementTreeContext {
	var p = new(ElementTreeContext)

	p.ElementExpressionContext = NewEmptyElementExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ElementExpressionContext))

	return p
}

func (s *ElementTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementTreeContext) PrimitiveExpression() IPrimitiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveExpressionContext)
}

func (s *ElementTreeContext) ElementAccess() IElementAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementAccessContext)
}

func (s *ElementTreeContext) CallExpression() ICallExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallExpressionContext)
}

func (s *ElementTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitElementTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) ElementExpression() (localctx IElementExpressionContext) {
	localctx = NewElementExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MonkeyParserRULE_elementExpression)

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

	localctx = NewElementTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.PrimitiveExpression()
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(107)
			p.ElementAccess()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(108)
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

func (s *ElementAccessContext) CopyFrom(ctx *ElementAccessContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ElementAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElementAccessTreeContext struct {
	*ElementAccessContext
}

func NewElementAccessTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElementAccessTreeContext {
	var p = new(ElementAccessTreeContext)

	p.ElementAccessContext = NewEmptyElementAccessContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ElementAccessContext))

	return p
}

func (s *ElementAccessTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementAccessTreeContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_BRACKET, 0)
}

func (s *ElementAccessTreeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElementAccessTreeContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_BRACKET, 0)
}

func (s *ElementAccessTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitElementAccessTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) ElementAccess() (localctx IElementAccessContext) {
	localctx = NewElementAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MonkeyParserRULE_elementAccess)

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

	localctx = NewElementAccessTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(MonkeyParserL_BRACKET)
	}
	{
		p.SetState(112)
		p.Expression()
	}
	{
		p.SetState(113)
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

func (s *CallExpressionContext) CopyFrom(ctx *CallExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionCallTreeContext struct {
	*CallExpressionContext
}

func NewFunctionCallTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallTreeContext {
	var p = new(FunctionCallTreeContext)

	p.CallExpressionContext = NewEmptyCallExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CallExpressionContext))

	return p
}

func (s *FunctionCallTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallTreeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_PAREN, 0)
}

func (s *FunctionCallTreeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_PAREN, 0)
}

func (s *FunctionCallTreeContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionCallTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitFunctionCallTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) CallExpression() (localctx ICallExpressionContext) {
	localctx = NewCallExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MonkeyParserRULE_callExpression)
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

	localctx = NewFunctionCallTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(MonkeyParserL_PAREN)
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(MonkeyParserL_PAREN-4))|(1<<(MonkeyParserL_BRACKET-4))|(1<<(MonkeyParserL_CURLY-4))|(1<<(MonkeyParserIF-4))|(1<<(MonkeyParserTRUE-4))|(1<<(MonkeyParserFALSE-4))|(1<<(MonkeyParserPUTS-4))|(1<<(MonkeyParserFUNC-4))|(1<<(MonkeyParserLEN-4))|(1<<(MonkeyParserFIRST-4))|(1<<(MonkeyParserLAST-4))|(1<<(MonkeyParserREST-4))|(1<<(MonkeyParserPUSH-4))|(1<<(MonkeyParserIDENTIFIER-4))|(1<<(MonkeyParserINTEGER-4))|(1<<(MonkeyParserSTRING-4)))) != 0 {
		{
			p.SetState(116)
			p.ExpressionList()
		}

	}
	{
		p.SetState(119)
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

func (s *PrimitiveExpressionContext) CopyFrom(ctx *PrimitiveExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimitiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringContext struct {
	*PrimitiveExpressionContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSTRING, 0)
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionATreeContext struct {
	*PrimitiveExpressionContext
}

func NewFunctionATreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionATreeContext {
	var p = new(FunctionATreeContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *FunctionATreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionATreeContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionATreeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_PAREN, 0)
}

func (s *FunctionATreeContext) FunctionParameters() IFunctionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *FunctionATreeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_PAREN, 0)
}

func (s *FunctionATreeContext) BlockStatement() IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *FunctionATreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitFunctionATree(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayTreeContext struct {
	*PrimitiveExpressionContext
}

func NewArrayTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayTreeContext {
	var p = new(ArrayTreeContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *ArrayTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTreeContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_BRACKET, 0)
}

func (s *ArrayTreeContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_BRACKET, 0)
}

func (s *ArrayTreeContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArrayTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitArrayTree(s)

	default:
		return t.VisitChildren(s)
	}
}

type HashObjectTreeContext struct {
	*PrimitiveExpressionContext
}

func NewHashObjectTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HashObjectTreeContext {
	var p = new(HashObjectTreeContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *HashObjectTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashObjectTreeContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_CURLY, 0)
}

func (s *HashObjectTreeContext) AllHashContent() []IHashContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHashContentContext)(nil)).Elem())
	var tst = make([]IHashContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHashContentContext)
		}
	}

	return tst
}

func (s *HashObjectTreeContext) HashContent(i int) IHashContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHashContentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHashContentContext)
}

func (s *HashObjectTreeContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_CURLY, 0)
}

func (s *HashObjectTreeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *HashObjectTreeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *HashObjectTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitHashObjectTree(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueContext struct {
	*PrimitiveExpressionContext
}

func NewTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueContext {
	var p = new(TrueContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *TrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserTRUE, 0)
}

func (s *TrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitTrue(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseContext struct {
	*PrimitiveExpressionContext
}

func NewFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseContext {
	var p = new(FalseContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *FalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFALSE, 0)
}

func (s *FalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerContext struct {
	*PrimitiveExpressionContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserINTEGER, 0)
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayFunctionTreeContext struct {
	*PrimitiveExpressionContext
}

func NewArrayFunctionTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayFunctionTreeContext {
	var p = new(ArrayFunctionTreeContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *ArrayFunctionTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayFunctionTreeContext) ArrayFunctions() IArrayFunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayFunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayFunctionsContext)
}

func (s *ArrayFunctionTreeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_PAREN, 0)
}

func (s *ArrayFunctionTreeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_PAREN, 0)
}

func (s *ArrayFunctionTreeContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArrayFunctionTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitArrayFunctionTree(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierTreeContext struct {
	*PrimitiveExpressionContext
}

func NewIdentifierTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierTreeContext {
	var p = new(IdentifierTreeContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *IdentifierTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierTreeContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitIdentifierTree(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupedExpressionTreeContext struct {
	*PrimitiveExpressionContext
}

func NewGroupedExpressionTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupedExpressionTreeContext {
	var p = new(GroupedExpressionTreeContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *GroupedExpressionTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupedExpressionTreeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_PAREN, 0)
}

func (s *GroupedExpressionTreeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GroupedExpressionTreeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_PAREN, 0)
}

func (s *GroupedExpressionTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitGroupedExpressionTree(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintTreeContext struct {
	*PrimitiveExpressionContext
}

func NewPrintTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintTreeContext {
	var p = new(PrintTreeContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *PrintTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintTreeContext) PUTS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPUTS, 0)
}

func (s *PrintTreeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_PAREN, 0)
}

func (s *PrintTreeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintTreeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_PAREN, 0)
}

func (s *PrintTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitPrintTree(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConditionalTreeContext struct {
	*PrimitiveExpressionContext
}

func NewConditionalTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalTreeContext {
	var p = new(ConditionalTreeContext)

	p.PrimitiveExpressionContext = NewEmptyPrimitiveExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimitiveExpressionContext))

	return p
}

func (s *ConditionalTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalTreeContext) IF() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIF, 0)
}

func (s *ConditionalTreeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalTreeContext) AllBlockStatement() []IBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem())
	var tst = make([]IBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockStatementContext)
		}
	}

	return tst
}

func (s *ConditionalTreeContext) BlockStatement(i int) IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *ConditionalTreeContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserELSE, 0)
}

func (s *ConditionalTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitConditionalTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) PrimitiveExpression() (localctx IPrimitiveExpressionContext) {
	localctx = NewPrimitiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MonkeyParserRULE_primitiveExpression)
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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserINTEGER:
		localctx = NewIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Match(MonkeyParserINTEGER)
		}

	case MonkeyParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(MonkeyParserSTRING)
		}

	case MonkeyParserIDENTIFIER:
		localctx = NewIdentifierTreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.Identifier()
		}

	case MonkeyParserTRUE:
		localctx = NewTrueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(124)
			p.Match(MonkeyParserTRUE)
		}

	case MonkeyParserFALSE:
		localctx = NewFalseContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(125)
			p.Match(MonkeyParserFALSE)
		}

	case MonkeyParserL_PAREN:
		localctx = NewGroupedExpressionTreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(126)
			p.Match(MonkeyParserL_PAREN)
		}
		{
			p.SetState(127)
			p.Expression()
		}
		{
			p.SetState(128)
			p.Match(MonkeyParserR_PAREN)
		}

	case MonkeyParserL_BRACKET:
		localctx = NewArrayTreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(130)
			p.Match(MonkeyParserL_BRACKET)
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(MonkeyParserL_PAREN-4))|(1<<(MonkeyParserL_BRACKET-4))|(1<<(MonkeyParserL_CURLY-4))|(1<<(MonkeyParserIF-4))|(1<<(MonkeyParserTRUE-4))|(1<<(MonkeyParserFALSE-4))|(1<<(MonkeyParserPUTS-4))|(1<<(MonkeyParserFUNC-4))|(1<<(MonkeyParserLEN-4))|(1<<(MonkeyParserFIRST-4))|(1<<(MonkeyParserLAST-4))|(1<<(MonkeyParserREST-4))|(1<<(MonkeyParserPUSH-4))|(1<<(MonkeyParserIDENTIFIER-4))|(1<<(MonkeyParserINTEGER-4))|(1<<(MonkeyParserSTRING-4)))) != 0 {
			{
				p.SetState(131)
				p.ExpressionList()
			}

		}
		{
			p.SetState(134)
			p.Match(MonkeyParserR_BRACKET)
		}

	case MonkeyParserLEN, MonkeyParserFIRST, MonkeyParserLAST, MonkeyParserREST, MonkeyParserPUSH:
		localctx = NewArrayFunctionTreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(135)
			p.ArrayFunctions()
		}
		{
			p.SetState(136)
			p.Match(MonkeyParserL_PAREN)
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(MonkeyParserL_PAREN-4))|(1<<(MonkeyParserL_BRACKET-4))|(1<<(MonkeyParserL_CURLY-4))|(1<<(MonkeyParserIF-4))|(1<<(MonkeyParserTRUE-4))|(1<<(MonkeyParserFALSE-4))|(1<<(MonkeyParserPUTS-4))|(1<<(MonkeyParserFUNC-4))|(1<<(MonkeyParserLEN-4))|(1<<(MonkeyParserFIRST-4))|(1<<(MonkeyParserLAST-4))|(1<<(MonkeyParserREST-4))|(1<<(MonkeyParserPUSH-4))|(1<<(MonkeyParserIDENTIFIER-4))|(1<<(MonkeyParserINTEGER-4))|(1<<(MonkeyParserSTRING-4)))) != 0 {
			{
				p.SetState(137)
				p.ExpressionList()
			}

		}
		{
			p.SetState(140)
			p.Match(MonkeyParserR_PAREN)
		}

	case MonkeyParserFUNC:
		localctx = NewFunctionATreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(142)
			p.Function()
		}
		{
			p.SetState(143)
			p.Match(MonkeyParserL_PAREN)
		}
		{
			p.SetState(144)
			p.FunctionParameters()
		}
		{
			p.SetState(145)
			p.Match(MonkeyParserR_PAREN)
		}
		{
			p.SetState(146)
			p.BlockStatement()
		}

	case MonkeyParserL_CURLY:
		localctx = NewHashObjectTreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(148)
			p.Match(MonkeyParserL_CURLY)
		}
		{
			p.SetState(149)
			p.HashContent()
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserCOMMA {
			{
				p.SetState(150)
				p.Match(MonkeyParserCOMMA)
			}
			{
				p.SetState(151)
				p.HashContent()
			}

			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(157)
			p.Match(MonkeyParserR_CURLY)
		}

	case MonkeyParserPUTS:
		localctx = NewPrintTreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(159)
			p.Match(MonkeyParserPUTS)
		}
		{
			p.SetState(160)
			p.Match(MonkeyParserL_PAREN)
		}
		{
			p.SetState(161)
			p.Expression()
		}
		{
			p.SetState(162)
			p.Match(MonkeyParserR_PAREN)
		}

	case MonkeyParserIF:
		localctx = NewConditionalTreeContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(164)
			p.Match(MonkeyParserIF)
		}
		{
			p.SetState(165)
			p.Expression()
		}
		{
			p.SetState(166)
			p.BlockStatement()
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MonkeyParserELSE {
			{
				p.SetState(167)
				p.Match(MonkeyParserELSE)
			}
			{
				p.SetState(168)
				p.BlockStatement()
			}

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

func (s *ArrayFunctionsContext) CopyFrom(ctx *ArrayFunctionsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ArrayFunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayFunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayFirstContext struct {
	*ArrayFunctionsContext
}

func NewArrayFirstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayFirstContext {
	var p = new(ArrayFirstContext)

	p.ArrayFunctionsContext = NewEmptyArrayFunctionsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayFunctionsContext))

	return p
}

func (s *ArrayFirstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayFirstContext) FIRST() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFIRST, 0)
}

func (s *ArrayFirstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitArrayFirst(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayLastContext struct {
	*ArrayFunctionsContext
}

func NewArrayLastContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLastContext {
	var p = new(ArrayLastContext)

	p.ArrayFunctionsContext = NewEmptyArrayFunctionsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayFunctionsContext))

	return p
}

func (s *ArrayLastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLastContext) LAST() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLAST, 0)
}

func (s *ArrayLastContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitArrayLast(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayLenContext struct {
	*ArrayFunctionsContext
}

func NewArrayLenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLenContext {
	var p = new(ArrayLenContext)

	p.ArrayFunctionsContext = NewEmptyArrayFunctionsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayFunctionsContext))

	return p
}

func (s *ArrayLenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLenContext) LEN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLEN, 0)
}

func (s *ArrayLenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitArrayLen(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayPushContext struct {
	*ArrayFunctionsContext
}

func NewArrayPushContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayPushContext {
	var p = new(ArrayPushContext)

	p.ArrayFunctionsContext = NewEmptyArrayFunctionsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayFunctionsContext))

	return p
}

func (s *ArrayPushContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayPushContext) PUSH() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPUSH, 0)
}

func (s *ArrayPushContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitArrayPush(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayRestContext struct {
	*ArrayFunctionsContext
}

func NewArrayRestContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayRestContext {
	var p = new(ArrayRestContext)

	p.ArrayFunctionsContext = NewEmptyArrayFunctionsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayFunctionsContext))

	return p
}

func (s *ArrayRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayRestContext) REST() antlr.TerminalNode {
	return s.GetToken(MonkeyParserREST, 0)
}

func (s *ArrayRestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitArrayRest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) ArrayFunctions() (localctx IArrayFunctionsContext) {
	localctx = NewArrayFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MonkeyParserRULE_arrayFunctions)

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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserLEN:
		localctx = NewArrayLenContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Match(MonkeyParserLEN)
		}

	case MonkeyParserFIRST:
		localctx = NewArrayFirstContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Match(MonkeyParserFIRST)
		}

	case MonkeyParserLAST:
		localctx = NewArrayLastContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.Match(MonkeyParserLAST)
		}

	case MonkeyParserREST:
		localctx = NewArrayRestContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(176)
			p.Match(MonkeyParserREST)
		}

	case MonkeyParserPUSH:
		localctx = NewArrayPushContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(177)
			p.Match(MonkeyParserPUSH)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *FunctionParametersContext) CopyFrom(ctx *FunctionParametersContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionParametersTreeContext struct {
	*FunctionParametersContext
}

func NewFunctionParametersTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionParametersTreeContext {
	var p = new(FunctionParametersTreeContext)

	p.FunctionParametersContext = NewEmptyFunctionParametersContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionParametersContext))

	return p
}

func (s *FunctionParametersTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParametersTreeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserIDENTIFIER)
}

func (s *FunctionParametersTreeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENTIFIER, i)
}

func (s *FunctionParametersTreeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *FunctionParametersTreeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *FunctionParametersTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitFunctionParametersTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) FunctionParameters() (localctx IFunctionParametersContext) {
	localctx = NewFunctionParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MonkeyParserRULE_functionParameters)
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

	localctx = NewFunctionParametersTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(MonkeyParserIDENTIFIER)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserCOMMA {
		{
			p.SetState(181)
			p.Match(MonkeyParserCOMMA)
		}
		{
			p.SetState(182)
			p.Match(MonkeyParserIDENTIFIER)
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *HashContentContext) CopyFrom(ctx *HashContentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *HashContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HashPairTreeContext struct {
	*HashContentContext
}

func NewHashPairTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HashPairTreeContext {
	var p = new(HashPairTreeContext)

	p.HashContentContext = NewEmptyHashContentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*HashContentContext))

	return p
}

func (s *HashPairTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashPairTreeContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *HashPairTreeContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HashPairTreeContext) COLON() antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOLON, 0)
}

func (s *HashPairTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitHashPairTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) HashContent() (localctx IHashContentContext) {
	localctx = NewHashContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MonkeyParserRULE_hashContent)

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

	localctx = NewHashPairTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Expression()
	}
	{
		p.SetState(189)
		p.Match(MonkeyParserCOLON)
	}
	{
		p.SetState(190)
		p.Expression()
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

func (s *ExpressionListContext) CopyFrom(ctx *ExpressionListContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionListTreeContext struct {
	*ExpressionListContext
}

func NewExpressionListTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionListTreeContext {
	var p = new(ExpressionListTreeContext)

	p.ExpressionListContext = NewEmptyExpressionListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionListContext))

	return p
}

func (s *ExpressionListTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListTreeContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListTreeContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListTreeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *ExpressionListTreeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *ExpressionListTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitExpressionListTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MonkeyParserRULE_expressionList)
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

	localctx = NewExpressionListTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Expression()
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MonkeyParserCOMMA {
		{
			p.SetState(193)
			p.Match(MonkeyParserCOMMA)
		}
		{
			p.SetState(194)
			p.Expression()
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *BlockStatementContext) CopyFrom(ctx *BlockStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockTreeContext struct {
	*BlockStatementContext
}

func NewBlockTreeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockTreeContext {
	var p = new(BlockTreeContext)

	p.BlockStatementContext = NewEmptyBlockStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BlockStatementContext))

	return p
}

func (s *BlockTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockTreeContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(MonkeyParserL_CURLY, 0)
}

func (s *BlockTreeContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(MonkeyParserR_CURLY, 0)
}

func (s *BlockTreeContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockTreeContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockTreeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitBlockTree(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MonkeyParserRULE_blockStatement)
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

	localctx = NewBlockTreeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(MonkeyParserL_CURLY)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(MonkeyParserL_PAREN-4))|(1<<(MonkeyParserL_BRACKET-4))|(1<<(MonkeyParserL_CURLY-4))|(1<<(MonkeyParserIF-4))|(1<<(MonkeyParserLET-4))|(1<<(MonkeyParserRETURN-4))|(1<<(MonkeyParserTRUE-4))|(1<<(MonkeyParserFALSE-4))|(1<<(MonkeyParserPUTS-4))|(1<<(MonkeyParserFUNC-4))|(1<<(MonkeyParserLEN-4))|(1<<(MonkeyParserFIRST-4))|(1<<(MonkeyParserLAST-4))|(1<<(MonkeyParserREST-4))|(1<<(MonkeyParserPUSH-4))|(1<<(MonkeyParserIDENTIFIER-4))|(1<<(MonkeyParserINTEGER-4))|(1<<(MonkeyParserSTRING-4)))) != 0 {
		{
			p.SetState(201)
			p.Statement()
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(207)
		p.Match(MonkeyParserR_CURLY)
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDeclaration returns the declaration attribute.
	GetDeclaration() *LetStatementTreeContext

	// SetDeclaration sets the declaration attribute.
	SetDeclaration(*LetStatementTreeContext)

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	declaration *LetStatementTreeContext
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) GetDeclaration() *LetStatementTreeContext { return s.declaration }

func (s *IdentifierContext) SetDeclaration(v *LetStatementTreeContext) { s.declaration = v }

func (s *IdentifierContext) CopyFrom(ctx *IdentifierContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.declaration = ctx.declaration
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierNodeContext struct {
	*IdentifierContext
}

func NewIdentifierNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierNodeContext {
	var p = new(IdentifierNodeContext)

	p.IdentifierContext = NewEmptyIdentifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentifierContext))

	return p
}

func (s *IdentifierNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierNodeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENTIFIER, 0)
}

func (s *IdentifierNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitIdentifierNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MonkeyParserRULE_identifier)

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

	localctx = NewIdentifierNodeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(MonkeyParserIDENTIFIER)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDeclaration returns the declaration attribute.
	GetDeclaration() *FunctionATreeContext

	// SetDeclaration sets the declaration attribute.
	SetDeclaration(*FunctionATreeContext)

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	declaration *FunctionATreeContext
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonkeyParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) GetDeclaration() *FunctionATreeContext { return s.declaration }

func (s *FunctionContext) SetDeclaration(v *FunctionATreeContext) { s.declaration = v }

func (s *FunctionContext) CopyFrom(ctx *FunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.declaration = ctx.declaration
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionNodeContext struct {
	*FunctionContext
}

func NewFunctionNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionNodeContext {
	var p = new(FunctionNodeContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *FunctionNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNodeContext) FUNC() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFUNC, 0)
}

func (s *FunctionNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitFunctionNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MonkeyParserRULE_function)

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

	localctx = NewFunctionNodeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(MonkeyParserFUNC)
	}

	return localctx
}
