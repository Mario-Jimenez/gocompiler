package contextual

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// identifiers types
type IdentifierType int

const (
	INEUTRAL IdentifierType = iota
	IIDENTIFIER
	ICALL
	IHASH
	IARRAY
)

/*
	store data for identifiers
*/
type identifier struct {
	token      antlr.Token
	identifier IdentifierType
	parameters int
}

/*
	stores an identifier's data
*/
type identifierHelper struct {
	levels      int
	identifiers []identifier
}

// constructor
func newIdentifierHelper() *identifierHelper {
	return &identifierHelper{
		levels:      -1,
		identifiers: []identifier{},
	}
}

// creates new identifier helper (can be nested)
func (h *identifierHelper) newIdentifier() {
	h.levels++
	h.identifiers = append(h.identifiers, identifier{})
}

// setter
func (h *identifierHelper) setToken(token antlr.Token) {
	h.identifiers[h.levels].token = token
}

// getter
func (h *identifierHelper) getToken() antlr.Token {
	return h.identifiers[h.levels].token
}

// setter
func (h *identifierHelper) setType(identifier IdentifierType) {
	h.identifiers[h.levels].identifier = identifier
}

// getter
func (h *identifierHelper) getType() IdentifierType {
	return h.identifiers[h.levels].identifier
}

// setter
func (h *identifierHelper) setParameters(parameters int) {
	h.identifiers[h.levels].parameters = parameters
}

// getter
func (h *identifierHelper) getParameters() int {
	return h.identifiers[h.levels].parameters
}

// close identifier helper
func (h *identifierHelper) closeIdentifier() {
	h.identifiers = h.identifiers[:h.levels]
	h.levels--
}
