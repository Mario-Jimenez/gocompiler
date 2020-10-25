package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	elementExpression:
		primitiveExpression (elementAccess | callExpression)? # elementTree
		;
*/

func (v *contextualVisitor) VisitElementTree(ctx *parser.ElementTreeContext) interface{} {
	v.Visit(ctx.PrimitiveExpression())
	// node :=
	// node.(*visitResponse)

	if ctx.ElementAccess() != nil {
		v.Visit(ctx.ElementAccess())
		// node :=
		// node.(*visitResponse)
	} else if ctx.CallExpression() != nil {
		v.Visit(ctx.CallExpression())
		// node :=
		// node.(*visitResponse)
	}

	if v.isIdentifier {
		if v.isCall {
			v.functionTable.Retrieve(v.identifierToken, v.callParameters)
		} else {
			v.identificationTable.Retrieve(v.identifierToken)
		}
	}

	return nil
}
