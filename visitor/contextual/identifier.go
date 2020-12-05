package contextual

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	identifier
		locals[declaration:*LetStatementTreeContext]:
		IDENTIFIER # identifierNode
		;
*/

func (v *visitor) VisitIdentifierNode(ctx *parser.IdentifierNodeContext) interface{} {
	v.identifier.setToken(ctx.IDENTIFIER().GetSymbol())

	attr := v.table.Retrieve(ctx.IDENTIFIER().GetSymbol())
	if attr != nil {
		ctx.SetDeclaration(attr.GetDeclaration())
	}

	return nil
}
