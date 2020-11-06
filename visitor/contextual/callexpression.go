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
	if v.identifier.getType() != IIDENTIFIER {
		token := ctx.L_PAREN().GetSymbol()
		newError := fmt.Sprintf("line %d:%d invalid function call", token.GetLine(), token.GetColumn())
		v.table.AddError(newError, token.GetLine())

		return nil
	}

	v.identifier.setType(ICALL)

	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
	}

	attr := v.table.Retrieve(v.identifier.getToken())

	if attr != nil {
		token := v.identifier.getToken()
		if attr.GetType() != identification.FUNCTION {
			newError := fmt.Sprintf("line %d:%d identifier '%s' is not a function", token.GetLine(), token.GetColumn(), token.GetText())
			v.table.AddError(newError, token.GetLine())
		} else {
			functionData := attr.GetData().(*identification.FunctionData)
			if functionData.GetParameters() != v.identifier.getParameters() {
				newError := fmt.Sprintf("line %d:%d mismatched parameters for function '%s'", token.GetLine(), token.GetColumn(), token.GetText())
				v.table.AddError(newError, token.GetLine())
			}
		}
	}

	return nil
}
