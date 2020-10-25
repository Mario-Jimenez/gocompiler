package identification

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type attributeFunction struct {
	token      antlr.Token
	visited    bool
	parameters int
}

func NewAttributeFunction(token antlr.Token, parameters int) *attributeFunction {
	return &attributeFunction{
		token:      token,
		parameters: parameters,
	}
}

/*func (a *attributeFunction) Declaration() *parser.LetStatementTreeContext {
	return a.declarationContext
}*/

func (a *attributeFunction) setVisited() {
	a.visited = true
}

func (a *attributeFunction) wasVisited() bool {
	return a.visited
}

func (a *attributeFunction) getToken() antlr.Token {
	return a.token
}
