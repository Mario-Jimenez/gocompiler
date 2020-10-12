package visitor

import (
	"strings"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	elementAccess:
		L_BRACKET expression R_BRACKET # elementAccessTree
		;
*/

func (v *monkeyVisitor) VisitElementAccessTree(ctx *parser.ElementAccessTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	children = append(children, childNode(ctx.L_BRACKET().GetText(), terminalColor))

	node := v.Visit(ctx.Expression())
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
		children = append(children, childNode("ExpressionTree", errorColor))
	}

	if ctx.R_BRACKET() != nil {
		if !strings.Contains(ctx.R_BRACKET().GetText(), "<missing") {
			children = append(children, childNode(ctx.R_BRACKET().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("R_BRACKET", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("R_BRACKET", errorColor))
	}

	return nodeResponse("ElementAccessTree", hasError, children)
}
