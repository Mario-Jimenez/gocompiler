package graph

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	expressionList:
		expression (COMMA expression)* # expressionListTree
		;
*/

func (v *visitor) VisitExpressionListTree(ctx *parser.ExpressionListTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	node := v.Visit(ctx.Expression(0))
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
		children = append(children, childNode("ExpressionTree", errorColor))
	}

	totalBranches := len(ctx.AllExpression())
	index := 1
	for index < totalBranches {
		children = append(children, childNode(ctx.COMMA(index-1).GetText(), terminalColor))

		if ctx.Expression(index) != nil {
			node := v.Visit(ctx.Expression(index))
			if n, ok := node.(*visitResponse); ok {
				if n.hasFailed() {
					hasError = true
				}
				children = append(children, n.info())
			} else {
				hasError = true
				children = append(children, childNode("ExpressionTree", errorColor))
			}
		} else {
			hasError = true
			children = append(children, childNode("ExpressionTree", errorColor))
		}

		index++
	}

	return nodeResponse("ExpressionListTree", hasError, children)
}
