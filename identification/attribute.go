package identification

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type attribute struct {
	token              antlr.Token
	declarationContext *parser.LetStatementTreeContext
	visited            bool
}

func NewAttribute(token antlr.Token, ctx *parser.LetStatementTreeContext) *attribute {
	return &attribute{
		token:              token,
		declarationContext: ctx,
	}
}

func (a *attribute) Declaration() *parser.LetStatementTreeContext {
	return a.declarationContext
}

func (a *attribute) setVisited() {
	a.visited = true
}

func (a *attribute) wasVisited() bool {
	return a.visited
}

func (a *attribute) getToken() antlr.Token {
	return a.token
}
