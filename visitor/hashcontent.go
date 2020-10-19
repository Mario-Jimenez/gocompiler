package visitor

import (
	"strings"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	hashContent : expression COLON expression # hashPairTree;
*/

func (v *monkeyVisitor) VisitHashPairTree(ctx *parser.HashPairTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	if ctx.Expression(0) != nil {
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
	} else {
		hasError = true
		children = append(children, childNode("ExpressionTree", errorColor))
	}

	if ctx.COLON() != nil {
		if !strings.Contains(ctx.COLON().GetText(), "<missing") {
			children = append(children, childNode(ctx.COLON().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("COLON", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("COLON", errorColor))
	}

	if ctx.Expression(1) != nil {
		node := v.Visit(ctx.Expression(1))
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

	return nodeResponse("HashPairTree", hasError, children)
}
