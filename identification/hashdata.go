package identification

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type HashData struct {
	keys []antlr.Token
}

func NewHashData() *HashData {
	return &HashData{}
}

func (h *HashData) AddKey(token antlr.Token) {
	h.keys = append(h.keys, token)
}

func (h *HashData) FindKey(key antlr.Token) bool {
	for _, k := range h.keys {
		if k.GetText() == key.GetText() {
			return true
		}
	}

	return false
}
