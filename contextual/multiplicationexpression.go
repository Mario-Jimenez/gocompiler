package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	multiplicationExpression:
		elementExpression (multiplicationFactor elementExpression)* # multiplicationTree
		;
*/

func (v *contextualVisitor) VisitMultiplicationTree(ctx *parser.MultiplicationTreeContext) interface{} {
	v.Visit(ctx.ElementExpression(0))
	// node :=
	// node.(*visitResponse)

	totalBranches := len(ctx.AllElementExpression())
	index := 1
	for index < totalBranches {
		v.Visit(ctx.MultiplicationFactor(index - 1))
		// node =
		// node.(*visitResponse)

		v.Visit(ctx.ElementExpression(index))
		// node =
		// node.(*visitResponse)

		index++
	}

	return nil
}
