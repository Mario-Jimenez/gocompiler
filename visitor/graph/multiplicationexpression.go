package graph

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	multiplicationExpression:
		elementExpression (multiplicationFactor elementExpression)* # multiplicationTree
		;
*/

func (v *visitor) VisitMultiplicationTree(ctx *parser.MultiplicationTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	node := v.Visit(ctx.ElementExpression(0))
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
		children = append(children, childNode("ElementTree", errorColor))
	}

	totalBranches := len(ctx.AllElementExpression())
	index := 1
	for index < totalBranches {
		node = v.Visit(ctx.MultiplicationFactor(index - 1))
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
		}

		node = v.Visit(ctx.ElementExpression(index))
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("ElementTree", errorColor))
		}

		index++
	}

	return nodeResponse("MultiplicationTree", hasError, children)
}
