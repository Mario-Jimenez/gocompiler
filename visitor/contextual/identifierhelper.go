package contextual

import "github.com/antlr/antlr4/runtime/Go/antlr"

type identifier struct {
	identifier bool
	token      antlr.Token
	call       bool
	parameters int
}

type identifierHelper struct {
	levels      int
	identifiers []identifier
}

func newIdentifierHelper() *identifierHelper {
	return &identifierHelper{
		levels:      -1,
		identifiers: []identifier{},
	}
}

func (h *identifierHelper) newIdentifier() {
	h.levels++
	h.identifiers = append(h.identifiers, identifier{})
}

func (h *identifierHelper) markIdentifier() {
	h.identifiers[h.levels].identifier = true
}

func (h *identifierHelper) isIdentifier() bool {
	return h.identifiers[h.levels].identifier
}

func (h *identifierHelper) setToken(token antlr.Token) {
	h.identifiers[h.levels].token = token
}

func (h *identifierHelper) getToken() antlr.Token {
	return h.identifiers[h.levels].token
}

func (h *identifierHelper) markCall() {
	h.identifiers[h.levels].call = true
}

func (h *identifierHelper) isCall() bool {
	return h.identifiers[h.levels].call
}

func (h *identifierHelper) setParameters(parameters int) {
	h.identifiers[h.levels].parameters = parameters
}

func (h *identifierHelper) getParameters() int {
	return h.identifiers[h.levels].parameters
}

func (h *identifierHelper) closeIdentifier() {
	h.identifiers = h.identifiers[:h.levels]
	h.levels--
}
