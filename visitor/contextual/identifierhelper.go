package contextual

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type IdentifierType int

const (
	INEUTRAL IdentifierType = iota
	IIDENTIFIER
	ICALL
	IHASH
	IARRAY
)

type identifier struct {
	token      antlr.Token
	identifier IdentifierType
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

func (h *identifierHelper) setToken(token antlr.Token) {
	h.identifiers[h.levels].token = token
}

func (h *identifierHelper) getToken() antlr.Token {
	return h.identifiers[h.levels].token
}

func (h *identifierHelper) setType(identifier IdentifierType) {
	h.identifiers[h.levels].identifier = identifier
}

func (h *identifierHelper) getType() IdentifierType {
	return h.identifiers[h.levels].identifier
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
