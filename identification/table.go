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

// constructor
func NewTable(handler *errorsHandler) *Table {
	return &Table{
		table:  map[int]map[string]*attribute{},
		errors: handler,
	}
}

// new declaration
func (t *Table) Enter(id string, attr *attribute) {
	// overwrite if it already exists, declaration data changes
	t.table[t.level][id] = attr
}

// search for declaration's attribute
func (t *Table) Retrieve(token antlr.Token) *attribute {
	id := token.GetText()
	tempLevel := t.level
	for tempLevel > 0 {
		if attr, ok := t.table[tempLevel][id]; ok {
			// mark declaration as used
			attr.markVisited()
			return attr
		}
		tempLevel--
	}

	// declaration was not found
	newError := fmt.Sprintf("line %d:%d declaration for '%s' not found", token.GetLine(), token.GetColumn(), token.GetText())
	t.errors.Add(newError, token.GetLine())

	return nil
}

// create scope
func (t *Table) OpenScope() {
	t.level++
	t.table[t.level] = map[string]*attribute{}
}

// close scope
// validates that all declarations in scope were used
func (t *Table) CloseScope() {
	for key, declaration := range t.table[t.level] {
		if key == "Main" || key == "main" {
			continue
		}
		// declaration was not used
		if !declaration.wasVisited() {
			token := declaration.getToken()
			newError := fmt.Sprintf("line %d:%d identifier '%s' was declared but not used", token.GetLine(), token.GetColumn(), token.GetText())
			t.errors.Add(newError, token.GetLine())
		}
	}

	delete(t.table, t.level)
	t.level--
}

// print table declarations in console (debugging)
func (t *Table) Print() {
	var level int
	for level <= t.level {
		for k, v := range t.table[level] {
			fmt.Printf("level: %d, name: %s, type: %d\n", level, k, v.expression)
		}
		level++
	}
}

// store new error
func (t *Table) AddError(newError string, newLine int) {
	t.errors.Add(newError, newLine)
}

func (t *Table) GetInfo(token antlr.Token) Declaration {
	id := token.GetText()
	attr := t.table[t.level][id]

	return Declaration{
		level:      t.level,
		expression: attr.expression,
		data:       attr.data,
	}
}
