package contextual

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type DeclarationType int

const (
	DNEUTRAL DeclarationType = iota
	DFUNCTION
	DHASH
	DARRAY
)

type declaration struct {
	token       antlr.Token
	declaration DeclarationType
	parameters  int
	hashKeys    []hashKey
	nestedHash  int
}

type declarationHelper struct {
	levels       int
	declarations []declaration
}

func newDeclarationHelper() *declarationHelper {
	return &declarationHelper{
		levels:       -1,
		declarations: []declaration{},
	}
}

func (h *declarationHelper) newDeclaration() {
	h.levels++
	h.declarations = append(h.declarations, declaration{})
}

func (h *declarationHelper) setToken(token antlr.Token) {
	h.declarations[h.levels].token = token
}

func (h *declarationHelper) getToken() antlr.Token {
	if h.levels > -1 {
		return h.declarations[h.levels].token
	}

	return nil
}

func (h *declarationHelper) setType(declaration DeclarationType) {
	if h.levels > -1 {
		h.declarations[h.levels].declaration = declaration
	}
}

func (h *declarationHelper) getType() DeclarationType {
	if h.levels > -1 {
		return h.declarations[h.levels].declaration
	}

	return DNEUTRAL
}

func (h *declarationHelper) setParameters(parameters int) {
	h.declarations[h.levels].parameters = parameters
}

func (h *declarationHelper) getParameters() int {
	return h.declarations[h.levels].parameters
}

func (h *declarationHelper) incNestedHash() {
	if h.levels > -1 {
		h.declarations[h.levels].nestedHash++
	}
}

func (h *declarationHelper) decNestedHash() {
	if h.levels > -1 {
		h.declarations[h.levels].nestedHash--
	}
}

func (h *declarationHelper) isNestedHash() bool {
	if h.levels > -1 {
		if h.declarations[h.levels].nestedHash == 0 {
			return false
		}
	}

	return true
}

func (h *declarationHelper) addHashKey(token antlr.Token, key HashType) bool {
	for _, k := range h.declarations[h.levels].hashKeys {
		if k.token.GetText() == token.GetText() {
			// error: key already exists
			return false
		}
	}

	h.declarations[h.levels].hashKeys = append(h.declarations[h.levels].hashKeys, hashKey{
		token: token,
		key:   key,
	})

	return true
}

func (h *declarationHelper) getHashKeys() []hashKey {
	return h.declarations[h.levels].hashKeys
}

func (h *declarationHelper) closeDeclaration() {
	h.declarations = h.declarations[:h.levels]
	h.levels--
}
