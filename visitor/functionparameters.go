package visitor

import (
	"strings"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	functionParameters:
		IDENTIFIER (COMMA IDENTIFIER)* # functionParametersTree
		;
*/

func (v *monkeyVisitor) VisitFunctionParametersTree(ctx *parser.FunctionParametersTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	if ctx.IDENTIFIER(0) != nil {
		if !strings.Contains(ctx.IDENTIFIER(0).GetText(), "<missing") {
			children = append(children, childNode(ctx.IDENTIFIER(0).GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("IDENTIFIER", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("IDENTIFIER", errorColor))
	}

	totalBranches := len(ctx.AllIDENTIFIER())
	index := 1
	for index < totalBranches {
		children = append(children, childNode(ctx.COMMA(index-1).GetText(), terminalColor))

		if ctx.IDENTIFIER(index) != nil {
			if !strings.Contains(ctx.IDENTIFIER(index).GetText(), "<missing") {
				children = append(children, childNode(ctx.IDENTIFIER(index).GetText(), terminalColor))
			} else {
				hasError = true
				children = append(children, childNode("IDENTIFIER", errorColor))
			}
		} else {
			hasError = true
			children = append(children, childNode("IDENTIFIER", errorColor))
		}

		index++
	}

	return nodeResponse("FunctionParametersTree", hasError, children)
}
