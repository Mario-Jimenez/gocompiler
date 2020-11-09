package contextual

import (
	"fmt"

	"github.com/Mario-Jimenez/gocompiler/identification"
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
		attr := v.table.Retrieve(v.identifier.getToken())
		if attr != nil && attr.GetType() != identification.IDENTIFIER {
			token := v.identifier.getToken()
			newError := fmt.Sprintf("line %d:%d invalid usage for identifier '%s'", token.GetLine(), token.GetColumn(), token.GetText())
			v.table.AddError(newError, token.GetLine())
		}
	}

	v.identifier.closeIdentifier()

	return nil
}
