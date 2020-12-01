package contextual

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// declarations types
type DeclarationType int

const (
	DNEUTRAL DeclarationType = iota
	DFUNCTION
	DHASH
	DARRAY
)

/*
	store data for different types of declarations
*/
type declaration struct {
	token         antlr.Token
	functionToken antlr.Token
	declaration   DeclarationType
	parameters    int
	hashKeys      []hashKey
	nestedHash    int
	arrayElements []arrayElement
}

/*
	stores a declaration's data
*/
type declarationHelper struct {
	levels       int
	declarations []declaration
}

// constructor
func newDeclarationHelper() *declarationHelper {
	return &declarationHelper{
		levels:       -1,
		declarations: []declaration{},
	}
}

// creates new declaration helper (can be nested)
func (h *declarationHelper) newDeclaration() {
	h.levels++
	h.declarations = append(h.declarations, declaration{})
}

// setter
func (h *declarationHelper) setToken(token antlr.Token) {
	h.declarations[h.levels].token = token
}

// getter
func (h *declarationHelper) getToken() antlr.Token {
	// validates that we're working on a declaration
	if h.levels > -1 {
		return h.declarations[h.levels].token
	}

	return nil
}

// setter
func (h *declarationHelper) setFunctionToken(token antlr.Token) {
	h.declarations[h.levels].functionToken = token
}

// getter
func (h *declarationHelper) getFunctionToken() antlr.Token {
	// validates that we're working on a declaration
	if h.levels > -1 {
		return h.declarations[h.levels].functionToken
	}

	return nil
}

// setter
func (h *declarationHelper) setType(declaration DeclarationType) {
	// validates that we're working on a declaration
	if h.levels > -1 {
		h.declarations[h.levels].declaration = declaration
	}
}

// getter
func (h *declarationHelper) getType() DeclarationType {
	// validates that we're working on a declaration
	if h.levels > -1 {
		return h.declarations[h.levels].declaration
	}

	return DNEUTRAL
}

// setter
func (h *declarationHelper) setParameters(parameters int) {
	h.declarations[h.levels].parameters = parameters
}

// getter
func (h *declarationHelper) getParameters() int {
	return h.declarations[h.levels].parameters
}

// when working with hash declarations, handle nested hashs
func (h *declarationHelper) incNestedHash() {
	// validates that we're working on a declaration
	if h.levels > -1 {
		h.declarations[h.levels].nestedHash++
	}
}

// when working with hash declarations, handle nested hashs
func (h *declarationHelper) decNestedHash() {
	// validates that we're working on a declaration
	if h.levels > -1 {
		h.declarations[h.levels].nestedHash--
	}
}

// when working with hash declarations, checks if it is a nested hash
func (h *declarationHelper) isNestedHash() bool {
	// validates that we're working on a declaration
	if h.levels > -1 {
		if h.declarations[h.levels].nestedHash == 0 {
			return false
		}
	}

	return true
}

// when working with hash declarations, store a new hash key
// validate if hash key is repeated
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

// getter
func (h *declarationHelper) getHashKeys() []hashKey {
	return h.declarations[h.levels].hashKeys
}

// when working with array declarations, store new array function element
func (h *declarationHelper) addArrayElement(index, parameters int, hasReturn bool) {
	// validates that we're working on a declaration
	if h.levels > -1 {
		h.declarations[h.levels].arrayElements = append(h.declarations[h.levels].arrayElements, arrayElement{
			index:      index,
			parameters: parameters,
			hasReturn:  hasReturn,
		})
	}
}

// getter
func (h *declarationHelper) getArrayElements() []arrayElement {
	// validates that we're working on a declaration
	if h.levels > -1 {
		return h.declarations[h.levels].arrayElements
	}

	return []arrayElement{}
}

// close declaration helper
func (h *declarationHelper) closeDeclaration() {
	h.declarations = h.declarations[:h.levels]
	h.levels--
}
