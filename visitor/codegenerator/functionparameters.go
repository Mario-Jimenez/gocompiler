package codegenerator

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	functionParameters:
		IDENTIFIER (COMMA IDENTIFIER)* # functionParametersTree
		;
*/

func (v *visitor) VisitFunctionParametersTree(ctx *parser.FunctionParametersTreeContext) interface{} {
	for _, id := range ctx.AllIDENTIFIER() {
		v.addInstruction("PUSH_LOCAL", id.GetText())
	}

	return nil
}
