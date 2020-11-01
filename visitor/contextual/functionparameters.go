package contextual

import (
	"fmt"

	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	functionParameters:
		IDENTIFIER (COMMA IDENTIFIER)* # functionParametersTree
		;
*/

func (v *visitor) VisitFunctionParametersTree(ctx *parser.FunctionParametersTreeContext) interface{} {
	parameters := map[string]interface{}{}
	parameters[ctx.IDENTIFIER(0).GetText()] = nil

	token := ctx.IDENTIFIER(0).GetSymbol()
	attr := identification.NewGeneralAttribute(token)
	v.generalTable.Enter(ctx.IDENTIFIER(0).GetText(), attr, true)

	totalBranches := len(ctx.AllIDENTIFIER())
	v.declaration.setParameters(totalBranches)

	index := 1
	for index < totalBranches {
		token := ctx.IDENTIFIER(index).GetSymbol()

		if _, ok := parameters[ctx.IDENTIFIER(index).GetText()]; ok {
			// error: parameter already exists
			newError := fmt.Sprintf("line %d:%d parameter '%s' was already declared in function", token.GetLine(), token.GetColumn(), token.GetText())
			v.generalTable.AddError(newError, token.GetLine())

			index++
			continue
		}

		parameters[ctx.IDENTIFIER(0).GetText()] = nil

		attr := identification.NewGeneralAttribute(token)
		v.generalTable.Enter(ctx.IDENTIFIER(index).GetText(), attr, true)

		index++
	}

	return nil
}
