package graph

import (
	"strings"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	statement:
		LET IDENTIFIER ASSIGN expression SEMI?	# letStatementTree
		| RETURN expression SEMI?				# returnStatementTree
		| expression SEMI?						# expressionStatementTree
		;
*/

func (v *visitor) VisitLetStatementTree(ctx *parser.LetStatementTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	children = append(children, childNode(ctx.LET().GetText(), terminalColor))

	if ctx.IDENTIFIER() != nil {
		if !strings.Contains(ctx.IDENTIFIER().GetText(), "<missing") {
			children = append(children, childNode(ctx.IDENTIFIER().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("IDENTIFIER", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("IDENTIFIER", errorColor))
	}

	if ctx.ASSIGN() != nil {
		if !strings.Contains(ctx.ASSIGN().GetText(), "<missing") {
			children = append(children, childNode(ctx.ASSIGN().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("ASSIGN", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("ASSIGN", errorColor))
	}

	if ctx.Expression() != nil {
		node := v.Visit(ctx.Expression())
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
		}
	} else {
		hasError = true
		children = append(children, childNode("ExpressionTree", errorColor))
	}

	if ctx.SEMI() != nil {
		children = append(children, childNode(ctx.SEMI().GetText(), terminalColor))
	}

	return nodeResponse("LetStatementTree", hasError, children)
}

func (v *visitor) VisitReturnStatementTree(ctx *parser.ReturnStatementTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	children = append(children, childNode(ctx.RETURN().GetText(), terminalColor))

	if ctx.Expression() != nil {
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
	} else {
		hasError = true
		children = append(children, childNode("ExpressionTree", errorColor))
	}

	if ctx.SEMI() != nil {
		children = append(children, childNode(ctx.SEMI().GetText(), terminalColor))
	}

	return nodeResponse("ReturnStatementTree", hasError, children)
}

func (v *visitor) VisitExpressionStatementTree(ctx *parser.ExpressionStatementTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

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

	if ctx.SEMI() != nil {
		children = append(children, childNode(ctx.SEMI().GetText(), terminalColor))
	}

	return nodeResponse("ExpressionStatementTree", hasError, children)
}
