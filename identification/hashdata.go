package identification

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// handles information for a hash declaration
type HashData struct {
	keys []antlr.Token
}

// constructor
func NewHashData() *HashData {
	return &HashData{}
}

// new hash key
func (h *HashData) AddKey(token antlr.Token) {
	h.keys = append(h.keys, token)
}

// search for a key
func (h *HashData) FindKey(key antlr.Token) bool {
	for _, k := range h.keys {
		if k.GetText() == key.GetText() {
			return true
		}
	}

	return false
}

// getter
func (h *HashData) GetKeys() []antlr.Token {
	return h.keys
}
