package identification

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// GeneralTable for declarations
type GeneralTable struct {
	table  map[int]map[string]*generalAttribute
	level  int
	errors *errorsHandler
}

func NewGeneralTable(handler *errorsHandler) *GeneralTable {
	return &GeneralTable{
		table:  map[int]map[string]*generalAttribute{},
		errors: handler,
	}
}

// TODO: parameters shouldn't be repeated
func (t *GeneralTable) Enter(id string, attr *generalAttribute, parameter bool) {
	if parameter {
		attr.setVisited()
	}

	// overwrite if it already exists, token data changes
	t.table[t.level][id] = attr
}

func (t *GeneralTable) Validate(token antlr.Token) {
	id := token.GetText()
	tempLevel := t.level
	for tempLevel > 0 {
		if attr, ok := t.table[tempLevel][id]; ok {
			attr.setVisited()
			return
		}
		tempLevel--
	}

	newError := fmt.Sprintf("line %d:%d declaration for '%s' not found", token.GetLine(), token.GetColumn(), token.GetText())
	t.errors.Add(newError, token.GetLine())
}

func (t *GeneralTable) OpenScope() {
	t.level++
	t.table[t.level] = map[string]*generalAttribute{}
}

func (t *GeneralTable) CloseScope() {
	for _, declaration := range t.table[t.level] {
		if !declaration.wasVisited() {
			token := declaration.getToken()
			newError := fmt.Sprintf("line %d:%d identifier '%s' was declared but not used", token.GetLine(), token.GetColumn(), token.GetText())
			t.errors.Add(newError, token.GetLine())
		}
	}
	delete(t.table, t.level)
	t.level--
}

func (t *GeneralTable) AddError(newError string, newLine int) {
	t.errors.Add(newError, newLine)
}
