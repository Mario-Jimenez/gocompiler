package identification

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Table of declarations
type Table struct {
	table      map[int]map[string]*attribute
	level      int
	errorList  []string
	errorlines []int
}

func NewTable() *Table {
	return &Table{
		table: map[int]map[string]*attribute{},
	}
}

// c.currentToken.GetLine(), c.currentToken.GetTokenSource().GetCharPositionInLine()
func (t *Table) Enter(id string, attr *attribute) {
	if declaration, ok := t.table[t.level][id]; ok {
		declarationToken := declaration.getToken()
		msg := fmt.Sprintf("duplicated identifier '%s' at %d:%d", declarationToken.GetText(), declarationToken.GetLine(), declarationToken.GetColumn())
		newToken := attr.token
		newError := fmt.Sprintf("line %d:%d %s", newToken.GetLine(), newToken.GetColumn(), msg)
		t.addError(newError)
		t.addLine(newToken.GetLine())
		return
	}

	t.table[t.level][id] = attr
}

func (t *Table) Retrieve(token antlr.Token) *attribute {
	id := token.GetText()
	tempLevel := t.level
	for tempLevel > 0 {
		if attr, ok := t.table[tempLevel][id]; ok {
			attr.setVisited()
			return attr
		}
		tempLevel--
	}

	newError := fmt.Sprintf("line %d:%d identifier '%s' not found", token.GetLine(), token.GetColumn(), token.GetText())
	t.addError(newError)
	t.addLine(token.GetLine())

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
			t.addError(newError)
			t.addLine(token.GetLine())
		}
	}
	delete(t.table, t.level)
	t.level--
}

func (t *Table) addError(newError string) {
	t.errorList = append(t.errorList, newError)
}

func (t *Table) addLine(newLine int) {
	t.errorlines = append(t.errorlines, newLine)
}

// Errors
func (t *Table) Errors() []string {
	return t.errorList
}

// Lines
func (t *Table) Lines() []int {
	return t.errorlines
}
