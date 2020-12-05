package graph

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	identifier
		locals[declaration:*LetStatementTreeContext]:
		IDENTIFIER # identifierNode
		;
*/

func (v *visitor) VisitIdentifierNode(ctx *parser.IdentifierNodeContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.IDENTIFIER().GetText(), terminalColor),
		failed: false,
	}
}
