package codegenerator

import "github.com/antlr/antlr4/runtime/Go/antlr"

/*
	stores hash info
*/
type hashHelper struct {
	global  bool
	name    string
	indexes []int
	keys    []antlr.Token
}

// constructor
func newHashHelper(global bool, name string) *hashHelper {
	return &hashHelper{
		global: global,
		name:   name,
	}
}

// add instruction index
func (h *hashHelper) addIndex(index int) {
	h.indexes = append(h.indexes, index)
}
