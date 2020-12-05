package contextual

import (
	"fmt"

	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	callExpression:
		L_PAREN expressionList? R_PAREN # functionCallTree
		;
*/

func (v *visitor) VisitFunctionCallTree(ctx *parser.FunctionCallTreeContext) interface{} {
	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	// filter invalid function calls, e.g. "string"(params...), 4(params...), (4+2)(params...)
	if v.identifier.getType() != IIDENTIFIER {
		token := ctx.L_PAREN().GetSymbol()
		newError := fmt.Sprintf("line %d:%d invalid function call", token.GetLine(), token.GetColumn())
		v.table.AddError(newError, token.GetLine())

		return nil
	}

	// identifier is a function call
	v.identifier.setType(ICALL)

	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
	}

	// search for identifier in the identification table
	// if not found, the Retrieve method will create and store the error
	attr := v.table.Retrieve(v.identifier.getToken())
	if attr != nil {
		token := v.identifier.getToken()
		// identifier must be stored in the identification table as a function
		if attr.GetType() != identification.FUNCTION {
			newError := fmt.Sprintf("line %d:%d identifier '%s' is not a function", token.GetLine(), token.GetColumn(), token.GetText())
			v.table.AddError(newError, token.GetLine())
		} else {
			functionData := attr.GetData().(*identification.FunctionData)
			// total call parameters must match with function declaration total parameters
			if functionData.GetParameters() != v.identifier.getParameters() {
				newError := fmt.Sprintf("line %d:%d mismatched parameters for function '%s'", token.GetLine(), token.GetColumn(), token.GetText())
				v.table.AddError(newError, token.GetLine())
			}
			// validate a function has a return statement when calling it in a declaration
			if v.declaration.getToken() != nil {
				if !functionData.HasReturn() {
					newError := fmt.Sprintf("line %d:%d function '%s' has no return statement", token.GetLine(), token.GetColumn(), token.GetText())
					v.table.AddError(newError, token.GetLine())
				}
			}
		}
	}

	return nil
}
