// Code generated from c:\Users\Emmanuel\source\CompiPR1\monkeycompiler\grammar\MonkeyLexer.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 41, 257,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 7, 34, 193, 10,
	34, 12, 34, 14, 34, 196, 11, 34, 3, 35, 3, 35, 7, 35, 200, 10, 35, 12,
	35, 14, 35, 203, 11, 35, 3, 36, 3, 36, 7, 36, 207, 10, 36, 12, 36, 14,
	36, 210, 11, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 6, 39,
	219, 10, 39, 13, 39, 14, 39, 220, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3,
	40, 3, 40, 7, 40, 230, 10, 40, 12, 40, 14, 40, 233, 11, 40, 3, 40, 3, 40,
	3, 40, 3, 40, 3, 40, 3, 41, 6, 41, 241, 10, 41, 13, 41, 14, 41, 242, 3,
	41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 7, 42, 251, 10, 42, 12, 42, 14,
	42, 254, 11, 42, 3, 42, 3, 42, 3, 231, 2, 43, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16,
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25,
	49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34,
	67, 35, 69, 36, 71, 37, 73, 2, 75, 2, 77, 38, 79, 39, 81, 40, 83, 41, 3,
	2, 7, 5, 2, 12, 12, 15, 15, 36, 36, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2,
	50, 59, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 263, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2,
	2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2,
	2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3,
	2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 77,
	3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 3,
	85, 3, 2, 2, 2, 5, 87, 3, 2, 2, 2, 7, 89, 3, 2, 2, 2, 9, 91, 3, 2, 2, 2,
	11, 93, 3, 2, 2, 2, 13, 95, 3, 2, 2, 2, 15, 97, 3, 2, 2, 2, 17, 99, 3,
	2, 2, 2, 19, 101, 3, 2, 2, 2, 21, 103, 3, 2, 2, 2, 23, 106, 3, 2, 2, 2,
	25, 108, 3, 2, 2, 2, 27, 111, 3, 2, 2, 2, 29, 114, 3, 2, 2, 2, 31, 116,
	3, 2, 2, 2, 33, 118, 3, 2, 2, 2, 35, 120, 3, 2, 2, 2, 37, 122, 3, 2, 2,
	2, 39, 124, 3, 2, 2, 2, 41, 126, 3, 2, 2, 2, 43, 129, 3, 2, 2, 2, 45, 134,
	3, 2, 2, 2, 47, 138, 3, 2, 2, 2, 49, 145, 3, 2, 2, 2, 51, 150, 3, 2, 2,
	2, 53, 156, 3, 2, 2, 2, 55, 161, 3, 2, 2, 2, 57, 164, 3, 2, 2, 2, 59, 168,
	3, 2, 2, 2, 61, 174, 3, 2, 2, 2, 63, 179, 3, 2, 2, 2, 65, 184, 3, 2, 2,
	2, 67, 189, 3, 2, 2, 2, 69, 197, 3, 2, 2, 2, 71, 204, 3, 2, 2, 2, 73, 213,
	3, 2, 2, 2, 75, 215, 3, 2, 2, 2, 77, 218, 3, 2, 2, 2, 79, 224, 3, 2, 2,
	2, 81, 240, 3, 2, 2, 2, 83, 246, 3, 2, 2, 2, 85, 86, 7, 61, 2, 2, 86, 4,
	3, 2, 2, 2, 87, 88, 7, 60, 2, 2, 88, 6, 3, 2, 2, 2, 89, 90, 7, 46, 2, 2,
	90, 8, 3, 2, 2, 2, 91, 92, 7, 42, 2, 2, 92, 10, 3, 2, 2, 2, 93, 94, 7,
	43, 2, 2, 94, 12, 3, 2, 2, 2, 95, 96, 7, 93, 2, 2, 96, 14, 3, 2, 2, 2,
	97, 98, 7, 95, 2, 2, 98, 16, 3, 2, 2, 2, 99, 100, 7, 125, 2, 2, 100, 18,
	3, 2, 2, 2, 101, 102, 7, 127, 2, 2, 102, 20, 3, 2, 2, 2, 103, 104, 7, 63,
	2, 2, 104, 105, 7, 63, 2, 2, 105, 22, 3, 2, 2, 2, 106, 107, 7, 63, 2, 2,
	107, 24, 3, 2, 2, 2, 108, 109, 7, 62, 2, 2, 109, 110, 7, 63, 2, 2, 110,
	26, 3, 2, 2, 2, 111, 112, 7, 64, 2, 2, 112, 113, 7, 63, 2, 2, 113, 28,
	3, 2, 2, 2, 114, 115, 7, 62, 2, 2, 115, 30, 3, 2, 2, 2, 116, 117, 7, 64,
	2, 2, 117, 32, 3, 2, 2, 2, 118, 119, 7, 45, 2, 2, 119, 34, 3, 2, 2, 2,
	120, 121, 7, 47, 2, 2, 121, 36, 3, 2, 2, 2, 122, 123, 7, 44, 2, 2, 123,
	38, 3, 2, 2, 2, 124, 125, 7, 49, 2, 2, 125, 40, 3, 2, 2, 2, 126, 127, 7,
	107, 2, 2, 127, 128, 7, 104, 2, 2, 128, 42, 3, 2, 2, 2, 129, 130, 7, 103,
	2, 2, 130, 131, 7, 110, 2, 2, 131, 132, 7, 117, 2, 2, 132, 133, 7, 103,
	2, 2, 133, 44, 3, 2, 2, 2, 134, 135, 7, 110, 2, 2, 135, 136, 7, 103, 2,
	2, 136, 137, 7, 118, 2, 2, 137, 46, 3, 2, 2, 2, 138, 139, 7, 116, 2, 2,
	139, 140, 7, 103, 2, 2, 140, 141, 7, 118, 2, 2, 141, 142, 7, 119, 2, 2,
	142, 143, 7, 116, 2, 2, 143, 144, 7, 112, 2, 2, 144, 48, 3, 2, 2, 2, 145,
	146, 7, 118, 2, 2, 146, 147, 7, 116, 2, 2, 147, 148, 7, 119, 2, 2, 148,
	149, 7, 103, 2, 2, 149, 50, 3, 2, 2, 2, 150, 151, 7, 104, 2, 2, 151, 152,
	7, 99, 2, 2, 152, 153, 7, 110, 2, 2, 153, 154, 7, 117, 2, 2, 154, 155,
	7, 103, 2, 2, 155, 52, 3, 2, 2, 2, 156, 157, 7, 114, 2, 2, 157, 158, 7,
	119, 2, 2, 158, 159, 7, 118, 2, 2, 159, 160, 7, 117, 2, 2, 160, 54, 3,
	2, 2, 2, 161, 162, 7, 104, 2, 2, 162, 163, 7, 112, 2, 2, 163, 56, 3, 2,
	2, 2, 164, 165, 7, 110, 2, 2, 165, 166, 7, 103, 2, 2, 166, 167, 7, 112,
	2, 2, 167, 58, 3, 2, 2, 2, 168, 169, 7, 104, 2, 2, 169, 170, 7, 107, 2,
	2, 170, 171, 7, 116, 2, 2, 171, 172, 7, 117, 2, 2, 172, 173, 7, 118, 2,
	2, 173, 60, 3, 2, 2, 2, 174, 175, 7, 110, 2, 2, 175, 176, 7, 99, 2, 2,
	176, 177, 7, 117, 2, 2, 177, 178, 7, 118, 2, 2, 178, 62, 3, 2, 2, 2, 179,
	180, 7, 116, 2, 2, 180, 181, 7, 103, 2, 2, 181, 182, 7, 117, 2, 2, 182,
	183, 7, 118, 2, 2, 183, 64, 3, 2, 2, 2, 184, 185, 7, 114, 2, 2, 185, 186,
	7, 119, 2, 2, 186, 187, 7, 117, 2, 2, 187, 188, 7, 106, 2, 2, 188, 66,
	3, 2, 2, 2, 189, 194, 5, 73, 37, 2, 190, 193, 5, 73, 37, 2, 191, 193, 5,
	75, 38, 2, 192, 190, 3, 2, 2, 2, 192, 191, 3, 2, 2, 2, 193, 196, 3, 2,
	2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 68, 3, 2, 2, 2,
	196, 194, 3, 2, 2, 2, 197, 201, 5, 75, 38, 2, 198, 200, 5, 75, 38, 2, 199,
	198, 3, 2, 2, 2, 200, 203, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 201, 202,
	3, 2, 2, 2, 202, 70, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 204, 208, 7, 36,
	2, 2, 205, 207, 10, 2, 2, 2, 206, 205, 3, 2, 2, 2, 207, 210, 3, 2, 2, 2,
	208, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 211, 3, 2, 2, 2, 210,
	208, 3, 2, 2, 2, 211, 212, 7, 36, 2, 2, 212, 72, 3, 2, 2, 2, 213, 214,
	9, 3, 2, 2, 214, 74, 3, 2, 2, 2, 215, 216, 9, 4, 2, 2, 216, 76, 3, 2, 2,
	2, 217, 219, 9, 5, 2, 2, 218, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220,
	218, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 223,
	8, 39, 2, 2, 223, 78, 3, 2, 2, 2, 224, 225, 7, 49, 2, 2, 225, 226, 7, 44,
	2, 2, 226, 231, 3, 2, 2, 2, 227, 230, 5, 79, 40, 2, 228, 230, 11, 2, 2,
	2, 229, 227, 3, 2, 2, 2, 229, 228, 3, 2, 2, 2, 230, 233, 3, 2, 2, 2, 231,
	232, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 232, 234, 3, 2, 2, 2, 233, 231,
	3, 2, 2, 2, 234, 235, 7, 44, 2, 2, 235, 236, 7, 49, 2, 2, 236, 237, 3,
	2, 2, 2, 237, 238, 8, 40, 2, 2, 238, 80, 3, 2, 2, 2, 239, 241, 9, 6, 2,
	2, 240, 239, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 242,
	243, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 245, 8, 41, 2, 2, 245, 82,
	3, 2, 2, 2, 246, 247, 7, 49, 2, 2, 247, 248, 7, 49, 2, 2, 248, 252, 3,
	2, 2, 2, 249, 251, 10, 6, 2, 2, 250, 249, 3, 2, 2, 2, 251, 254, 3, 2, 2,
	2, 252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 255, 3, 2, 2, 2, 254,
	252, 3, 2, 2, 2, 255, 256, 8, 42, 2, 2, 256, 84, 3, 2, 2, 2, 12, 2, 192,
	194, 201, 208, 220, 229, 231, 242, 252, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "':'", "','", "'('", "')'", "'['", "']'", "'{'", "'}'", "'=='",
	"'='", "'<='", "'>='", "'<'", "'>'", "'+'", "'-'", "'*'", "'/'", "'if'",
	"'else'", "'let'", "'return'", "'true'", "'false'", "'puts'", "'fn'", "'len'",
	"'first'", "'last'", "'rest'", "'push'",
}

var lexerSymbolicNames = []string{
	"", "SEMI", "COLON", "COMMA", "L_PAREN", "R_PAREN", "L_BRACKET", "R_BRACKET",
	"L_CURLY", "R_CURLY", "EQUALS", "ASSIGN", "LESS_OR_EQUALS", "GREATER_OR_EQUALS",
	"LESS", "GREATER", "PLUS", "MINUS", "MULT", "DIV", "IF", "ELSE", "LET",
	"RETURN", "TRUE", "FALSE", "PUTS", "FUNC", "LEN", "FIRST", "LAST", "REST",
	"PUSH", "IDENTIFIER", "INTEGER", "STRING", "WS", "COMMENT", "TERMINATOR",
	"LINE_COMMENT",
}

var lexerRuleNames = []string{
	"SEMI", "COLON", "COMMA", "L_PAREN", "R_PAREN", "L_BRACKET", "R_BRACKET",
	"L_CURLY", "R_CURLY", "EQUALS", "ASSIGN", "LESS_OR_EQUALS", "GREATER_OR_EQUALS",
	"LESS", "GREATER", "PLUS", "MINUS", "MULT", "DIV", "IF", "ELSE", "LET",
	"RETURN", "TRUE", "FALSE", "PUTS", "FUNC", "LEN", "FIRST", "LAST", "REST",
	"PUSH", "IDENTIFIER", "INTEGER", "STRING", "LETTER", "DIGIT", "WS", "COMMENT",
	"TERMINATOR", "LINE_COMMENT",
}

type MonkeyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewMonkeyLexer(input antlr.CharStream) *MonkeyLexer {

	l := new(MonkeyLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "MonkeyLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MonkeyLexer tokens.
const (
	MonkeyLexerSEMI              = 1
	MonkeyLexerCOLON             = 2
	MonkeyLexerCOMMA             = 3
	MonkeyLexerL_PAREN           = 4
	MonkeyLexerR_PAREN           = 5
	MonkeyLexerL_BRACKET         = 6
	MonkeyLexerR_BRACKET         = 7
	MonkeyLexerL_CURLY           = 8
	MonkeyLexerR_CURLY           = 9
	MonkeyLexerEQUALS            = 10
	MonkeyLexerASSIGN            = 11
	MonkeyLexerLESS_OR_EQUALS    = 12
	MonkeyLexerGREATER_OR_EQUALS = 13
	MonkeyLexerLESS              = 14
	MonkeyLexerGREATER           = 15
	MonkeyLexerPLUS              = 16
	MonkeyLexerMINUS             = 17
	MonkeyLexerMULT              = 18
	MonkeyLexerDIV               = 19
	MonkeyLexerIF                = 20
	MonkeyLexerELSE              = 21
	MonkeyLexerLET               = 22
	MonkeyLexerRETURN            = 23
	MonkeyLexerTRUE              = 24
	MonkeyLexerFALSE             = 25
	MonkeyLexerPUTS              = 26
	MonkeyLexerFUNC              = 27
	MonkeyLexerLEN               = 28
	MonkeyLexerFIRST             = 29
	MonkeyLexerLAST              = 30
	MonkeyLexerREST              = 31
	MonkeyLexerPUSH              = 32
	MonkeyLexerIDENTIFIER        = 33
	MonkeyLexerINTEGER           = 34
	MonkeyLexerSTRING            = 35
	MonkeyLexerWS                = 36
	MonkeyLexerCOMMENT           = 37
	MonkeyLexerTERMINATOR        = 38
	MonkeyLexerLINE_COMMENT      = 39
)
