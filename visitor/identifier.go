package visitor

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	identifier	: IDENTIFIER # identifierNode;
*/

func (v *monkeyVisitor) VisitIdentifierNode(ctx *parser.IdentifierNodeContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.IDENTIFIER().GetText(), terminalColor),
		failed: false,
	}
}
