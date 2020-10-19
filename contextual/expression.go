package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	expression:
		additionExpression (comparison additionExpression)* # expressionTree
		;
*/

func (v *contextualVisitor) VisitExpressionTree(ctx *parser.ExpressionTreeContext) interface{} {
	v.Visit(ctx.AdditionExpression(0))
	// node :=
	// node.(*visitResponse)

	totalBranches := len(ctx.AllAdditionExpression())
	index := 1
	for index < totalBranches {
		v.Visit(ctx.Comparison(index - 1))
		// node =
		// node.(*visitResponse)

		v.Visit(ctx.AdditionExpression(index))
		// node =
		// node.(*visitResponse)

		index++
	}

	return nil
}
