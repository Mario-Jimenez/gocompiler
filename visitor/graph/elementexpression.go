package graph

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	elementExpression:
		primitiveExpression (elementAccess | callExpression)? # elementTree
		;
*/

func (v *visitor) VisitElementTree(ctx *parser.ElementTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	node := v.Visit(ctx.PrimitiveExpression())
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
	}

	if ctx.ElementAccess() != nil {
		node := v.Visit(ctx.ElementAccess())
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("ElementAccessTree", errorColor))
		}
	} else if ctx.CallExpression() != nil {
		node := v.Visit(ctx.CallExpression())
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("FunctionCallTree", errorColor))
		}
	}

	return nodeResponse("ElementTree", hasError, children)
}
