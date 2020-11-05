package contextual

import (
	"fmt"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	elementAccess:
		L_BRACKET expression R_BRACKET # elementAccessTree
		;
*/

func (v *visitor) VisitElementAccessTree(ctx *parser.ElementAccessTreeContext) interface{} {
	if v.identifier.getType() != IIDENTIFIER {
		token := ctx.L_BRACKET().GetSymbol()
		newError := fmt.Sprintf("line %d:%d invalid element access", token.GetLine(), token.GetColumn())
		v.table.AddError(newError, token.GetLine())

		return nil
	}

	v.Visit(ctx.Expression())

	return nil
}
