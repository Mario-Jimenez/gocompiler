package visitor

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	expression:
		additionExpression (comparison additionExpression)* # expressionTree
		;
*/

func (v *monkeyVisitor) VisitExpressionTree(ctx *parser.ExpressionTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	node := v.Visit(ctx.AdditionExpression(0))
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
		children = append(children, childNode("AdditionTree", errorColor))
	}

	totalBranches := len(ctx.AllAdditionExpression())
	index := 1
	for index < totalBranches {
		node = v.Visit(ctx.Comparison(index - 1))
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
		}

		node = v.Visit(ctx.AdditionExpression(index))
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("AdditionTree", errorColor))
		}

		index++
	}

	return nodeResponse("ExpressionTree", hasError, children)
}
