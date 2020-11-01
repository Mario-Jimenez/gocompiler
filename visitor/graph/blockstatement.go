package graph

import (
	"strings"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	blockStatement : L_CURLY statement* R_CURLY # blockTree;
*/

func (v *visitor) VisitBlockTree(ctx *parser.BlockTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	if ctx.L_CURLY() != nil {
		if !strings.Contains(ctx.L_CURLY().GetText(), "<missing") {
			children = append(children, childNode(ctx.L_CURLY().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("L_CURLY", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("L_CURLY", errorColor))
	}

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

	if ctx.R_CURLY() != nil {
		if !strings.Contains(ctx.L_CURLY().GetText(), "<missing") {
			children = append(children, childNode(ctx.R_CURLY().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("R_CURLY", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("R_CURLY", errorColor))
	}

	return nodeResponse("BlockTree", hasError, children)
}
