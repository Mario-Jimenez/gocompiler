package codegenerator

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	identifier
		locals[declaration:*LetStatementTreeContext]:
		IDENTIFIER # identifierNode
		;
*/

func (v *visitor) VisitIdentifierNode(ctx *parser.IdentifierNodeContext) interface{} {
	return ctx.IDENTIFIER().GetSymbol()
}
