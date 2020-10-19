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

func (v *contextualVisitor) VisitFunctionParametersTree(ctx *parser.FunctionParametersTreeContext) interface{} {
	token := ctx.IDENTIFIER(0).GetSymbol()
	attr := identification.NewAttribute(token, nil)
	v.identificationTable.Enter(ctx.IDENTIFIER(0).GetText(), attr)

	if ctx.IDENTIFIER(0) != nil {
		ctx.IDENTIFIER(0).GetText()
	}

	totalBranches := len(ctx.AllIDENTIFIER())
	index := 1
	for index < totalBranches {
		ctx.COMMA(index - 1).GetText()

		token := ctx.IDENTIFIER(index).GetSymbol()
		attr := identification.NewAttribute(token, nil)
		v.identificationTable.Enter(ctx.IDENTIFIER(index).GetText(), attr)

		if ctx.IDENTIFIER(index) != nil {
			ctx.IDENTIFIER(index).GetText()
		}

		index++
	}

	return nil
}
