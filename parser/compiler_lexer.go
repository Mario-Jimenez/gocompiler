// Code generated from ./parser/Compiler.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 34, 191,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 31, 7, 31, 169, 10, 31, 12, 31, 14, 31, 172,
	11, 31, 3, 32, 3, 32, 7, 32, 176, 10, 32, 12, 32, 14, 32, 179, 11, 32,
	3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 6, 35, 186, 10, 35, 13, 35, 14, 35,
	187, 3, 35, 3, 35, 2, 2, 36, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51,
	27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 2, 67, 2, 69, 34,
	3, 2, 4, 4, 2, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 192, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2,
	2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3,
	2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49,
	3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2,
	57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2,
	2, 69, 3, 2, 2, 2, 3, 71, 3, 2, 2, 2, 5, 73, 3, 2, 2, 2, 7, 75, 3, 2, 2,
	2, 9, 77, 3, 2, 2, 2, 11, 79, 3, 2, 2, 2, 13, 81, 3, 2, 2, 2, 15, 83, 3,
	2, 2, 2, 17, 85, 3, 2, 2, 2, 19, 87, 3, 2, 2, 2, 21, 90, 3, 2, 2, 2, 23,
	92, 3, 2, 2, 2, 25, 95, 3, 2, 2, 2, 27, 98, 3, 2, 2, 2, 29, 100, 3, 2,
	2, 2, 31, 102, 3, 2, 2, 2, 33, 104, 3, 2, 2, 2, 35, 106, 3, 2, 2, 2, 37,
	108, 3, 2, 2, 2, 39, 110, 3, 2, 2, 2, 41, 113, 3, 2, 2, 2, 43, 118, 3,
	2, 2, 2, 45, 122, 3, 2, 2, 2, 47, 129, 3, 2, 2, 2, 49, 134, 3, 2, 2, 2,
	51, 140, 3, 2, 2, 2, 53, 144, 3, 2, 2, 2, 55, 150, 3, 2, 2, 2, 57, 155,
	3, 2, 2, 2, 59, 160, 3, 2, 2, 2, 61, 165, 3, 2, 2, 2, 63, 173, 3, 2, 2,
	2, 65, 180, 3, 2, 2, 2, 67, 182, 3, 2, 2, 2, 69, 185, 3, 2, 2, 2, 71, 72,
	7, 61, 2, 2, 72, 4, 3, 2, 2, 2, 73, 74, 7, 60, 2, 2, 74, 6, 3, 2, 2, 2,
	75, 76, 7, 42, 2, 2, 76, 8, 3, 2, 2, 2, 77, 78, 7, 43, 2, 2, 78, 10, 3,
	2, 2, 2, 79, 80, 7, 93, 2, 2, 80, 12, 3, 2, 2, 2, 81, 82, 7, 95, 2, 2,
	82, 14, 3, 2, 2, 2, 83, 84, 7, 125, 2, 2, 84, 16, 3, 2, 2, 2, 85, 86, 7,
	127, 2, 2, 86, 18, 3, 2, 2, 2, 87, 88, 7, 63, 2, 2, 88, 89, 7, 63, 2, 2,
	89, 20, 3, 2, 2, 2, 90, 91, 7, 63, 2, 2, 91, 22, 3, 2, 2, 2, 92, 93, 7,
	62, 2, 2, 93, 94, 7, 63, 2, 2, 94, 24, 3, 2, 2, 2, 95, 96, 7, 64, 2, 2,
	96, 97, 7, 63, 2, 2, 97, 26, 3, 2, 2, 2, 98, 99, 7, 62, 2, 2, 99, 28, 3,
	2, 2, 2, 100, 101, 7, 64, 2, 2, 101, 30, 3, 2, 2, 2, 102, 103, 7, 45, 2,
	2, 103, 32, 3, 2, 2, 2, 104, 105, 7, 47, 2, 2, 105, 34, 3, 2, 2, 2, 106,
	107, 7, 44, 2, 2, 107, 36, 3, 2, 2, 2, 108, 109, 7, 49, 2, 2, 109, 38,
	3, 2, 2, 2, 110, 111, 7, 107, 2, 2, 111, 112, 7, 104, 2, 2, 112, 40, 3,
	2, 2, 2, 113, 114, 7, 103, 2, 2, 114, 115, 7, 110, 2, 2, 115, 116, 7, 117,
	2, 2, 116, 117, 7, 103, 2, 2, 117, 42, 3, 2, 2, 2, 118, 119, 7, 110, 2,
	2, 119, 120, 7, 103, 2, 2, 120, 121, 7, 118, 2, 2, 121, 44, 3, 2, 2, 2,
	122, 123, 7, 116, 2, 2, 123, 124, 7, 103, 2, 2, 124, 125, 7, 118, 2, 2,
	125, 126, 7, 119, 2, 2, 126, 127, 7, 116, 2, 2, 127, 128, 7, 112, 2, 2,
	128, 46, 3, 2, 2, 2, 129, 130, 7, 118, 2, 2, 130, 131, 7, 116, 2, 2, 131,
	132, 7, 119, 2, 2, 132, 133, 7, 103, 2, 2, 133, 48, 3, 2, 2, 2, 134, 135,
	7, 104, 2, 2, 135, 136, 7, 99, 2, 2, 136, 137, 7, 110, 2, 2, 137, 138,
	7, 117, 2, 2, 138, 139, 7, 103, 2, 2, 139, 50, 3, 2, 2, 2, 140, 141, 7,
	110, 2, 2, 141, 142, 7, 103, 2, 2, 142, 143, 7, 112, 2, 2, 143, 52, 3,
	2, 2, 2, 144, 145, 7, 104, 2, 2, 145, 146, 7, 107, 2, 2, 146, 147, 7, 116,
	2, 2, 147, 148, 7, 117, 2, 2, 148, 149, 7, 118, 2, 2, 149, 54, 3, 2, 2,
	2, 150, 151, 7, 110, 2, 2, 151, 152, 7, 99, 2, 2, 152, 153, 7, 117, 2,
	2, 153, 154, 7, 118, 2, 2, 154, 56, 3, 2, 2, 2, 155, 156, 7, 116, 2, 2,
	156, 157, 7, 103, 2, 2, 157, 158, 7, 117, 2, 2, 158, 159, 7, 118, 2, 2,
	159, 58, 3, 2, 2, 2, 160, 161, 7, 114, 2, 2, 161, 162, 7, 119, 2, 2, 162,
	163, 7, 117, 2, 2, 163, 164, 7, 106, 2, 2, 164, 60, 3, 2, 2, 2, 165, 170,
	5, 65, 33, 2, 166, 169, 5, 65, 33, 2, 167, 169, 5, 67, 34, 2, 168, 166,
	3, 2, 2, 2, 168, 167, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2,
	2, 2, 170, 171, 3, 2, 2, 2, 171, 62, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2,
	173, 177, 5, 67, 34, 2, 174, 176, 5, 67, 34, 2, 175, 174, 3, 2, 2, 2, 176,
	179, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 64, 3,
	2, 2, 2, 179, 177, 3, 2, 2, 2, 180, 181, 9, 2, 2, 2, 181, 66, 3, 2, 2,
	2, 182, 183, 4, 50, 59, 2, 183, 68, 3, 2, 2, 2, 184, 186, 9, 3, 2, 2, 185,
	184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188,
	3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 8, 35, 2, 2, 190, 70, 3, 2,
	2, 2, 7, 2, 168, 170, 177, 187, 3, 8, 2, 2,
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
	"", "';'", "':'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'=='", "'='",
	"'<='", "'>='", "'<'", "'>'", "'+'", "'-'", "'*'", "'/'", "'if'", "'else'",
	"'let'", "'return'", "'true'", "'false'", "'len'", "'first'", "'last'",
	"'rest'", "'push'",
}

var lexerSymbolicNames = []string{
	"", "PyCOMA", "DOSPUNTOS", "PIZQ", "PDER", "LBRACK", "RBRACK", "LBRACE",
	"RBRACE", "EQL", "ASSIGN", "LEQ", "GEQ", "LSS", "GTR", "SUMA", "RESTA",
	"MULT", "DIV", "IF", "ELSE", "LET", "RETURN", "TRUE", "FALSE", "LEN", "FIRST",
	"LAST", "REST", "PUSH", "IDENT", "LITERAL", "WS",
}

var lexerRuleNames = []string{
	"PyCOMA", "DOSPUNTOS", "PIZQ", "PDER", "LBRACK", "RBRACK", "LBRACE", "RBRACE",
	"EQL", "ASSIGN", "LEQ", "GEQ", "LSS", "GTR", "SUMA", "RESTA", "MULT", "DIV",
	"IF", "ELSE", "LET", "RETURN", "TRUE", "FALSE", "LEN", "FIRST", "LAST",
	"REST", "PUSH", "IDENT", "LITERAL", "LETRA", "DIGITO", "WS",
}

type Compiler struct {
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

func NewCompiler(input antlr.CharStream) *Compiler {

	l := new(Compiler)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Compiler.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Compiler tokens.
const (
	CompilerPyCOMA    = 1
	CompilerDOSPUNTOS = 2
	CompilerPIZQ      = 3
	CompilerPDER      = 4
	CompilerLBRACK    = 5
	CompilerRBRACK    = 6
	CompilerLBRACE    = 7
	CompilerRBRACE    = 8
	CompilerEQL       = 9
	CompilerASSIGN    = 10
	CompilerLEQ       = 11
	CompilerGEQ       = 12
	CompilerLSS       = 13
	CompilerGTR       = 14
	CompilerSUMA      = 15
	CompilerRESTA     = 16
	CompilerMULT      = 17
	CompilerDIV       = 18
	CompilerIF        = 19
	CompilerELSE      = 20
	CompilerLET       = 21
	CompilerRETURN    = 22
	CompilerTRUE      = 23
	CompilerFALSE     = 24
	CompilerLEN       = 25
	CompilerFIRST     = 26
	CompilerLAST      = 27
	CompilerREST      = 28
	CompilerPUSH      = 29
	CompilerIDENT     = 30
	CompilerLITERAL   = 31
	CompilerWS        = 32
)
