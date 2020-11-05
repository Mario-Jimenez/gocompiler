package identification

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ExpressionType int

const (
	NEUTRAL ExpressionType = iota
	IDENTIFIER
	FUNCTION
	HASH
	ARRAY
)

type attribute struct {
	expression ExpressionType
	token      antlr.Token
	visited    bool
	data       interface{}
}

func NewAttribute(expression ExpressionType, token antlr.Token, data interface{}) *attribute {
	return &attribute{
		expression: expression,
		token:      token,
		data:       data,
	}
}

func (a *attribute) getToken() antlr.Token {
	return a.token
}

func (a *attribute) markVisited() {
	a.visited = true
}

func (a *attribute) wasVisited() bool {
	return a.visited
}

func (a *attribute) GetType() ExpressionType {
	return a.expression
}

func (a *attribute) GetData() interface{} {
	return a.data
}
