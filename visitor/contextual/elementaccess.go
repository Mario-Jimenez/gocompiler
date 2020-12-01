package contextual

import (
	"fmt"
	"strconv"

	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	elementAccess:
		L_BRACKET expression R_BRACKET # elementAccessTree
		;
*/

func (v *visitor) VisitElementAccessTree(ctx *parser.ElementAccessTreeContext) interface{} {
	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	// filter invalid element access expressions, e.g. "string"[...], 4[...], (4+2)[...]
	if v.identifier.getType() != IIDENTIFIER {
		token := ctx.L_BRACKET().GetSymbol()
		newError := fmt.Sprintf("line %d:%d invalid element access", token.GetLine(), token.GetColumn())
		v.table.AddError(newError, token.GetLine())

		return nil
	}

	// re-use hash helper to verify that the expression used for the element access is integer or string
	v.hash.newHash()

	v.Visit(ctx.Expression())

	// search for identifier in the identification table
	// if not found, the Retrieve method will create and store the error
	attr := v.table.Retrieve(v.identifier.getToken())
	if attr != nil {
		token := v.identifier.getToken()
		// identifier must be stored in the identification table as hash or array
		if attr.GetType() == identification.HASH {
			// identifier is a hash
			v.identifier.setType(IHASH)
			// hash key must be integer or string
			if v.hash.getType() == HCOMPLEX || v.hash.getType() == HUNKNOWN {
				token := ctx.L_BRACKET().GetSymbol()
				newError := fmt.Sprintf("line %d:%d hash key must be INTEGER or STRING", token.GetLine(), token.GetColumn())
				v.table.AddError(newError, token.GetLine())
			} else {
				// search for key in hash data stored in the identification table
				hashData := attr.GetData().(*identification.HashData)
				if !hashData.FindKey(v.hash.getToken()) {
					// key doesn't exists
					token := v.hash.getToken()
					newError := fmt.Sprintf("line %d:%d hash key %s doesn't exist", token.GetLine(), token.GetColumn(), token.GetText())
					v.table.AddError(newError, token.GetLine())
				}
			}
		} else if attr.GetType() == identification.ARRAY {
			// identifier is an array
			v.identifier.setType(IARRAY)
			// search array index only when the element access expression is an integer
			if v.hash.getType() == HINTEGER {
				arrayData := attr.GetData().(*identification.ArrayData)
				index, _ := strconv.Atoi(v.hash.getToken().GetText())
				arrayElement := arrayData.FindElement(index)
				if arrayElement != nil {
					// mark declaration as function
					v.declaration.setType(DFUNCTION)
					token := v.declaration.getToken()
					// check if we're working on a declaration
					if token != nil {
						functionData := identification.NewFunctionData(arrayElement.GetParameters())
						if arrayElement.HasReturn() {
							functionData.SetReturn()
						}
						attr := identification.NewAttribute(identification.FUNCTION, token, functionData)
						v.table.Enter(token.GetText(), attr)
					}
				}
			}
		} else {
			newError := fmt.Sprintf("line %d:%d identifier '%s' is not a hash or an array", token.GetLine(), token.GetColumn(), token.GetText())
			v.table.AddError(newError, token.GetLine())
		}
	}

	v.hash.closeHash()

	return nil
}
