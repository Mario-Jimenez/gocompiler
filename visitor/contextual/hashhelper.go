package contextual

import "github.com/antlr/antlr4/runtime/Go/antlr"

// hash key types
type HashType int

const (
	HUNKNOWN HashType = iota
	HINTEGER
	HSTRING
	HCOMPLEX
)

/*
	store hash key data
*/
type hashKey struct {
	token antlr.Token
	key   HashType
}

/*
	store hash key data
	can have nested hashs, but only stores data for the main hash
*/
type hashHelper struct {
	levels int
	hashs  []hashKey
}

// constructor
func newHashHelper() *hashHelper {
	return &hashHelper{
		levels: -1,
		hashs:  []hashKey{},
	}
}

// creates new hash helper (can be nested)
func (h *hashHelper) newHash() {
	h.levels++
	h.hashs = append(h.hashs, hashKey{})
}

// setter
func (h *hashHelper) setToken(token antlr.Token) {
	// validates that we're working with a hash expression
	if h.levels > -1 {
		h.hashs[h.levels].token = token
	}
}

// getter
func (h *hashHelper) getToken() antlr.Token {
	return h.hashs[h.levels].token
}

// setter
func (h *hashHelper) setType(key HashType) {
	// validates that we're working with a hash expression
	if h.levels > -1 {
		h.hashs[h.levels].key = key
	}
}

// getter
func (h *hashHelper) getType() HashType {
	// validates that we're working with a hash expression
	if h.levels > -1 {
		return h.hashs[h.levels].key
	}

	return HCOMPLEX
}

// close hash helper
func (h *hashHelper) closeHash() {
	h.hashs = h.hashs[:h.levels]
	h.levels--
}
