package contextual

import "github.com/antlr/antlr4/runtime/Go/antlr"

type HashType int

const (
	HUNKNOWN HashType = iota
	HINTEGER
	HSTRING
	HCOMPLEX
)

type hashKey struct {
	token antlr.Token
	key   HashType
}

type hashHelper struct {
	levels int
	hashs  []hashKey
}

func newHashHelper() *hashHelper {
	return &hashHelper{
		levels: -1,
		hashs:  []hashKey{},
	}
}

func (h *hashHelper) newHash() {
	h.levels++
	h.hashs = append(h.hashs, hashKey{})
}

func (h *hashHelper) setToken(token antlr.Token) {
	if h.levels > -1 {
		h.hashs[h.levels].token = token
	}
}

func (h *hashHelper) getToken() antlr.Token {
	return h.hashs[h.levels].token
}

func (h *hashHelper) setType(key HashType) {
	if h.levels > -1 {
		h.hashs[h.levels].key = key
	}
}

func (h *hashHelper) getType() HashType {
	if h.levels > -1 {
		return h.hashs[h.levels].key
	}

	return HCOMPLEX
}

func (h *hashHelper) closeHash() {
	h.hashs = h.hashs[:h.levels]
	h.levels--
}
