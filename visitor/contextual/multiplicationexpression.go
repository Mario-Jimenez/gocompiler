package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	multiplicationExpression:
		elementExpression (multiplicationFactor elementExpression)* # multiplicationTree
		;
*/

func (v *visitor) VisitMultiplicationTree(ctx *parser.MultiplicationTreeContext) interface{} {
	v.Visit(ctx.ElementExpression(0))

	totalBranches := len(ctx.AllElementExpression())
	index := 1
	for index < totalBranches {
		v.hash.setType(HCOMPLEX)

		v.Visit(ctx.MultiplicationFactor(index - 1))

		v.Visit(ctx.ElementExpression(index))

		index++
	}

	return nil
}
