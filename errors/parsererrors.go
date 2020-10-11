package errors

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ParserErrorListener handles error detection,
// store error message and row number where error occured
type ParserErrorListener struct {
	*antlr.DefaultErrorListener
	errorList  []string
	errorlines []int
}

// NewParserErrorListener instance
func NewParserErrorListener() *ParserErrorListener {
	return new(ParserErrorListener)
}

// SyntaxError info, stores error message and error line
func (p *ParserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	p.errorList = append(p.errorList, fmt.Sprintf("line %d:%d %s", line, column, msg))
	p.errorlines = append(p.errorlines, line)
}

// Errors
func (p *ParserErrorListener) Errors() []string {
	return p.errorList
}

// Lines
func (p *ParserErrorListener) Lines() []int {
	return p.errorlines
}
