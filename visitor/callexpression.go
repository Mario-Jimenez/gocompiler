package visitor

import (
	"strings"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	callExpression:
		L_PAREN expressionList? R_PAREN # functionCallTree
		;
*/

func (v *monkeyVisitor) VisitFunctionCallTree(ctx *parser.FunctionCallTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	children = append(children, childNode(ctx.L_PAREN().GetText(), terminalColor))

	if ctx.ExpressionList() != nil {
		node := v.Visit(ctx.ExpressionList())
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("ExpressionListTree", errorColor))
		}
	}

	if ctx.R_PAREN() != nil {
		if !strings.Contains(ctx.R_PAREN().GetText(), "<missing") {
			children = append(children, childNode(ctx.R_PAREN().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("R_PAREN", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("R_PAREN", errorColor))
	}

	return nodeResponse("FunctionCallTree", hasError, children)
}
