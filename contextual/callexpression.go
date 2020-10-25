package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	callExpression:
		L_PAREN expressionList? R_PAREN # functionCallTree
		;
*/

func (v *contextualVisitor) VisitFunctionCallTree(ctx *parser.FunctionCallTreeContext) interface{} {
	ctx.L_PAREN().GetText()

	v.isCall = true

	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
		// node :=
		// node.(*visitResponse)
	}

	if ctx.R_PAREN() != nil {
		ctx.R_PAREN().GetText()
	}

	return nil
}
