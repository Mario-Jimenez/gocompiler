package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	additionExpression:
		multiplicationExpression (
			additionFactor multiplicationExpression
		)* # additionTree
		;
*/

func (v *contextualVisitor) VisitAdditionTree(ctx *parser.AdditionTreeContext) interface{} {
	v.Visit(ctx.MultiplicationExpression(0))
	// node :=
	// node.(*visitResponse)

	totalBranches := len(ctx.AllMultiplicationExpression())
	index := 1
	for index < totalBranches {
		v.Visit(ctx.AdditionFactor(index - 1))
		// node =
		// node.(*visitResponse)

		v.Visit(ctx.MultiplicationExpression(index))
		// node =
		// node.(*visitResponse)

		index++
	}

	return nil
}
