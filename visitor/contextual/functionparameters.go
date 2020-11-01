package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	functionParameters:
		IDENTIFIER (COMMA IDENTIFIER)* # functionParametersTree
		;
*/

func (v *visitor) VisitFunctionParametersTree(ctx *parser.FunctionParametersTreeContext) interface{} {
	token := ctx.IDENTIFIER(0).GetSymbol()
	attr := identification.NewGeneralAttribute(token)
	v.generalTable.Enter(ctx.IDENTIFIER(0).GetText(), attr, true)

	totalBranches := len(ctx.AllIDENTIFIER())
	v.declaration.setParameters(totalBranches)

	index := 1
	for index < totalBranches {

		token := ctx.IDENTIFIER(index).GetSymbol()
		attr := identification.NewGeneralAttribute(token)
		v.generalTable.Enter(ctx.IDENTIFIER(index).GetText(), attr, true)

		index++
	}

	return nil
}
