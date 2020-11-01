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
	v.Visit(ctx.PrimitiveExpression())

	if ctx.ElementAccess() != nil {
		v.Visit(ctx.ElementAccess())
	} else if ctx.CallExpression() != nil {
		v.Visit(ctx.CallExpression())
	}

	if v.identifier.isIdentifier() {
		if v.identifier.isCall() {
			v.functionTable.Validate(v.identifier.getToken(), v.identifier.getParameters())
			v.identifier.reset()
			return nil
		}
		v.generalTable.Validate(v.identifier.getToken())
		v.identifier.reset()
	}

	return nil
}
