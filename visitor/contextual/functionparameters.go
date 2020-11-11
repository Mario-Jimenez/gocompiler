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
	token := v.declaration.getToken()
	// check if we're working on a declaration
	if token != nil {
		// store function declaration in identification table
		// before entering the body of the function
		// this allows recursion to work
		functionData := identification.NewFunctionData(len(ctx.AllIDENTIFIER()))
		attr := identification.NewAttribute(identification.FUNCTION, token, functionData)
		v.table.Enter(token.GetText(), attr)
		// avoid creating a function declaration in a nested function expression that is not a declaration
		v.declaration.setToken(nil)
	}

	v.table.OpenScope()

	// store function declaration parameters
	// and check that the identifiers are not repeated
	parameters := map[string]interface{}{}
	parameters[ctx.IDENTIFIER(0).GetText()] = nil

	token = ctx.IDENTIFIER(0).GetSymbol()
	attr := identification.NewAttribute(identification.IDENTIFIER, token, nil)
	v.table.Enter(ctx.IDENTIFIER(0).GetText(), attr)

	totalBranches := len(ctx.AllIDENTIFIER())

	// handle function inside array
	v.array.setParameters(totalBranches)
	v.array.activate()

	index := 1
	for index < totalBranches {
		token := ctx.IDENTIFIER(index).GetSymbol()

		if _, ok := parameters[ctx.IDENTIFIER(index).GetText()]; ok {
			// error: parameter already exists
			newError := fmt.Sprintf("line %d:%d parameter '%s' was already declared in function", token.GetLine(), token.GetColumn(), token.GetText())
			v.table.AddError(newError, token.GetLine())

			index++
			continue
		}

		parameters[ctx.IDENTIFIER(0).GetText()] = nil

		attr := identification.NewAttribute(identification.IDENTIFIER, token, nil)
		v.table.Enter(ctx.IDENTIFIER(index).GetText(), attr)

		index++
	}

	return nil
}
