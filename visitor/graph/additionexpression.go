package graph

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	additionExpression:
		multiplicationExpression (
			additionFactor multiplicationExpression
		)* # additionTree
		;
*/

func (v *visitor) VisitAdditionTree(ctx *parser.AdditionTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	node := v.Visit(ctx.MultiplicationExpression(0))
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
		children = append(children, childNode("MultiplicationTree", errorColor))
	}

	totalBranches := len(ctx.AllMultiplicationExpression())
	index := 1
	for index < totalBranches {
		node = v.Visit(ctx.AdditionFactor(index - 1))
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
		}

		node = v.Visit(ctx.MultiplicationExpression(index))
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("MultiplicationTree", errorColor))
		}

		index++
	}

	return nodeResponse("AdditionTree", hasError, children)
}
