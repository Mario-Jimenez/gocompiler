package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	identifier	: IDENTIFIER # identifierNode;
*/

func (v *contextualVisitor) VisitIdentifierNode(ctx *parser.IdentifierNodeContext) interface{} {
	attr := v.identificationTable.Retrieve(ctx.IDENTIFIER().GetSymbol())
	if attr != nil {
		ctx.SetDeclaration(attr.Declaration())
	}

	v.isIdentifier = true
	v.identifierToken = ctx.IDENTIFIER().GetSymbol()

	return nil
}
