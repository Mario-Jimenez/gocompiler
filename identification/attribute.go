package identification

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// expressions types
type ExpressionType int

const (
	NEUTRAL ExpressionType = iota
	IDENTIFIER
	FUNCTION
	HASH
	ARRAY
)

/*
	indentification table attribute
	data stores information for a especific type of expression
		- identifier (no data)
		- function (functiondata.go)
		- hash (hashdata.go)
		- array (arraydata.go)
*/
type attribute struct {
	expression         ExpressionType
	token              antlr.Token
	declarationContext *parser.LetStatementTreeContext
	visited            bool
	data               interface{}
}

// constructor
func NewAttribute(expression ExpressionType, token antlr.Token, ctx *parser.LetStatementTreeContext, data interface{}) *attribute {
	return &attribute{
		expression:         expression,
		token:              token,
		declarationContext: ctx,
		data:               data,
	}
}

// getter
func (a *attribute) getToken() antlr.Token {
	return a.token
}

// a declaration was used
func (a *attribute) markVisited() {
	a.visited = true
}

// check if a declaration was used
func (a *attribute) wasVisited() bool {
	return a.visited
}

// getter
func (a *attribute) GetType() ExpressionType {
	return a.expression
}

// getter
func (a *attribute) GetDeclaration() *parser.LetStatementTreeContext {
	return a.declarationContext
}

// getter
func (a *attribute) GetData() interface{} {
	return a.data
}
