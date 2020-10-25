package visitor

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

//function locals[declaration: *FunctionTreeContext]: FUNC # functionNode;

func (v *monkeyVisitor) VisitFunctionNode(ctx *parser.FunctionNodeContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.FUNC().GetText(), terminalColor),
		failed: false,
	}
}
