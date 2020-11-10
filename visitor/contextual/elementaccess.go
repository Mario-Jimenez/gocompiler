package contextual

import (
	"fmt"

	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	elementAccess:
		L_BRACKET expression R_BRACKET # elementAccessTree
		;
*/

func (v *visitor) VisitElementAccessTree(ctx *parser.ElementAccessTreeContext) interface{} {
	v.hash.setType(HCOMPLEX)

	if v.identifier.getType() != IIDENTIFIER {
		token := ctx.L_BRACKET().GetSymbol()
		newError := fmt.Sprintf("line %d:%d invalid element access", token.GetLine(), token.GetColumn())
		v.table.AddError(newError, token.GetLine())

		return nil
	}

	v.hash.newHash()

	v.Visit(ctx.Expression())

	// verify key exists
	attr := v.table.Retrieve(v.identifier.getToken())
	if attr != nil {
		token := v.identifier.getToken()
		if attr.GetType() == identification.HASH {
			v.identifier.setType(IHASH)
			if v.hash.getType() == HCOMPLEX || v.hash.getType() == HUNKNOWN {
				// key must be integer or string
				token := ctx.L_BRACKET().GetSymbol()
				newError := fmt.Sprintf("line %d:%d hash key must be INTEGER or STRING", token.GetLine(), token.GetColumn())
				v.table.AddError(newError, token.GetLine())
			} else {
				hashData := attr.GetData().(*identification.HashData)
				if !hashData.FindKey(v.hash.getToken()) {
					// key doesn't exists
					token := v.hash.getToken()
					newError := fmt.Sprintf("line %d:%d hash key %s doesn't exist", token.GetLine(), token.GetColumn(), token.GetText())
					v.table.AddError(newError, token.GetLine())
				}
			}
		} else if attr.GetType() == identification.ARRAY {
			v.identifier.setType(IARRAY)
		} else {
			newError := fmt.Sprintf("line %d:%d identifier '%s' is not a hash or an array", token.GetLine(), token.GetColumn(), token.GetText())
			v.table.AddError(newError, token.GetLine())
		}
	}

	v.hash.closeHash()

	return nil
}
