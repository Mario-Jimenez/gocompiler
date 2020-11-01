package contextual

import "github.com/antlr/antlr4/runtime/Go/antlr"

type identifierHelper struct {
	identifier bool
	token      antlr.Token
	call       bool
	parameters int
}

func newIdentifierHelper() *identifierHelper {
	return &identifierHelper{}
}

func (h *identifierHelper) markIdentifier() {
	h.identifier = true
}

func (h *identifierHelper) isIdentifier() bool {
	return h.identifier
}

func (h *identifierHelper) setToken(token antlr.Token) {
	h.token = token
}

func (h *identifierHelper) getToken() antlr.Token {
	return h.token
}

func (h *identifierHelper) markCall() {
	h.call = true
}

func (h *identifierHelper) isCall() bool {
	return h.call
}

func (h *identifierHelper) setParameters(parameters int) {
	h.parameters = parameters
}

func (h *identifierHelper) getParameters() int {
	return h.parameters
}

func (h *identifierHelper) reset() {
	h.identifier = false
	h.token = nil
	h.call = false
	h.parameters = 0
}
