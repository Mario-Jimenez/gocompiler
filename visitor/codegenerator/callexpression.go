package codegenerator

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	callExpression:
		L_PAREN expressionList? R_PAREN # functionCallTree
		;
*/

func (v *visitor) VisitFunctionCallTree(ctx *parser.FunctionCallTreeContext) interface{} {
	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
	}

	return nil
}
