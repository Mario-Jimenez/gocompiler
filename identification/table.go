package identification

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Table for identification
type Table struct {
	table  map[int]map[string]*attribute
	level  int
	errors *errorsHandler
}

func NewTable(handler *errorsHandler) *Table {
	return &Table{
		table:  map[int]map[string]*attribute{},
		errors: handler,
	}
}

func (t *Table) Enter(id string, attr *attribute) {
	// overwrite if it already exists, token data changes
	t.table[t.level][id] = attr
}

func (t *Table) Retrieve(token antlr.Token) *attribute {
	id := token.GetText()
	tempLevel := t.level
	for tempLevel > 0 {
		if attr, ok := t.table[tempLevel][id]; ok {
			attr.markVisited()
			return attr
		}
		tempLevel--
	}

	newError := fmt.Sprintf("line %d:%d declaration for '%s' not found", token.GetLine(), token.GetColumn(), token.GetText())
	t.errors.Add(newError, token.GetLine())

	return nil
}

func (t *Table) OpenScope() {
	t.level++
	t.table[t.level] = map[string]*attribute{}
}

func (t *Table) CloseScope() {
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

func (t *Table) AddError(newError string, newLine int) {
	t.errors.Add(newError, newLine)
}
