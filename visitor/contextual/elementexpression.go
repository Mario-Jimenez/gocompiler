package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	elementExpression:
		primitiveExpression (elementAccess | callExpression)? # elementTree
		;
*/

func (v *visitor) VisitElementTree(ctx *parser.ElementTreeContext) interface{} {
	v.identifier.newIdentifier()

	v.Visit(ctx.PrimitiveExpression())

	if ctx.ElementAccess() != nil {
		v.Visit(ctx.ElementAccess())
	} else if ctx.CallExpression() != nil {
		v.Visit(ctx.CallExpression())
	}

	if v.identifier.getType() == IIDENTIFIER {
		v.table.Retrieve(v.identifier.getToken())
	}

	v.identifier.closeIdentifier()

	return nil
}
