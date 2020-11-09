package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	expression:
		additionExpression (comparison additionExpression)* # expressionTree
		;
*/

func (v *visitor) VisitExpressionTree(ctx *parser.ExpressionTreeContext) interface{} {
	v.Visit(ctx.AdditionExpression(0))

	totalBranches := len(ctx.AllAdditionExpression())
	index := 1
	for index < totalBranches {
		v.hash.setType(HCOMPLEX)

		v.Visit(ctx.Comparison(index - 1))

		v.Visit(ctx.AdditionExpression(index))

		index++
	}

	return nil
}
