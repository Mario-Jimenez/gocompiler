package identification

import "github.com/antlr/antlr4/runtime/Go/antlr"

type functionAttribute struct {
	token      antlr.Token
	parameters int
}

func NewFunctionAttribute(token antlr.Token, parameters int) *functionAttribute {
	return &functionAttribute{
		token:      token,
		parameters: parameters,
	}
}

func (a *functionAttribute) getToken() antlr.Token {
	return a.token
}

func (a *functionAttribute) getParameters() int {
	return a.parameters
}
