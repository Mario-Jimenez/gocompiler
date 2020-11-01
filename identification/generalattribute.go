package identification

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type generalAttribute struct {
	token   antlr.Token
	visited bool
}

func NewGeneralAttribute(token antlr.Token) *generalAttribute {
	return &generalAttribute{
		token: token,
	}
}

func (a *generalAttribute) setVisited() {
	a.visited = true
}

func (a *generalAttribute) wasVisited() bool {
	return a.visited
}

func (a *generalAttribute) getToken() antlr.Token {
	return a.token
}
