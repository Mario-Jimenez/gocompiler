package identification

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// FunctionTable for function declarations
type FunctionTable struct {
	table  map[int]map[string]*functionAttribute
	level  int
	errors *errorsHandler
}

func NewFunctionTable(handler *errorsHandler) *FunctionTable {
	return &FunctionTable{
		table:  map[int]map[string]*functionAttribute{},
		errors: handler,
	}
}

func (t *FunctionTable) Enter(id string, attr *functionAttribute) {
	if _, ok := t.table[t.level][id]; ok {
		return
	}

	t.table[t.level][id] = attr
}

func (t *FunctionTable) Validate(token antlr.Token, parameters int) {
	var found bool
	id := token.GetText()
	tempLevel := t.level
	for tempLevel > 0 {
		if attr, ok := t.table[tempLevel][id]; ok {
			found = true
			if attr.getParameters() == parameters {
				return
			}
		}
		tempLevel--
	}

	if found {
		newError := fmt.Sprintf("line %d:%d mismatched number of parameters passed to function '%s'", token.GetLine(), token.GetColumn(), token.GetText())
		t.errors.addError(newError)
		t.errors.addLine(token.GetLine())
		return
	}

	newError := fmt.Sprintf("line %d:%d function declaration for '%s' not found", token.GetLine(), token.GetColumn(), token.GetText())
	t.errors.addError(newError)
	t.errors.addLine(token.GetLine())
}

func (t *FunctionTable) OpenScope() {
	t.level++
	t.table[t.level] = map[string]*functionAttribute{}
}

func (t *FunctionTable) CloseScope() {
	delete(t.table, t.level)
	t.level--
}
