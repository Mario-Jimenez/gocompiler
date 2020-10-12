package visitor

import (
	"strings"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	program : statement* EOF # programTree;
*/

func (v *monkeyVisitor) VisitProgramTree(ctx *parser.ProgramTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	for _, branch := range ctx.AllStatement() {
		node := v.Visit(branch)
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
		}
	}

	if ctx.EOF() != nil {
		if !strings.Contains(ctx.EOF().GetText(), "<missing") {
			children = append(children, childNode(ctx.EOF().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("EOF", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("EOF", errorColor))
	}

	if hasError {
		return parentNode("ProgramTree", errorColor, children)
	}

	return parentNode("ProgramTree", rootColor, children)
}
