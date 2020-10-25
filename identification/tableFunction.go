package identification

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// TableFunction of declarations
type TableFunction struct {
	table      map[int]map[string]*attributeFunction
	level      int
	errorList  []string
	errorlines []int
}

func NewTableFunction() *TableFunction {
	return &TableFunction{
		table: map[int]map[string]*attributeFunction{},
	}
}

// c.currentToken.GetLine(), c.currentToken.GetTokenSource().GetCharPositionInLine()
func (t *TableFunction) Enter(id string, attr *attributeFunction) {
	if _, ok := t.table[t.level][id]; ok {
		return
	}

	t.table[t.level][id] = attr
}

func (t *TableFunction) Retrieve(token antlr.Token, parametersCount int) *attributeFunction {
	id := token.GetText()
	tempLevel := t.level
	for tempLevel > 0 {
		if attr, ok := t.table[tempLevel][id]; ok {
			if attr.parameters == parametersCount {
				attr.setVisited()
				return attr
			}
		}
		tempLevel--
	}

	newError := fmt.Sprintf("line %d:%d identifier '%s' not found", token.GetLine(), token.GetColumn(), token.GetText())
	t.addError(newError)
	t.addLine(token.GetLine())

	return nil
}

func (t *TableFunction) OpenScope() {
	t.level++
	t.table[t.level] = map[string]*attributeFunction{}
}
func (t *TableFunction) CloseScope() {
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

func (t *TableFunction) addError(newError string) {
	t.errorList = append(t.errorList, newError)
}

func (t *TableFunction) addLine(newLine int) {
	t.errorlines = append(t.errorlines, newLine)
}

// Errors
func (t *TableFunction) Errors() []string {
	return t.errorList
}

// Lines
func (t *TableFunction) Lines() []int {
	return t.errorlines
}
