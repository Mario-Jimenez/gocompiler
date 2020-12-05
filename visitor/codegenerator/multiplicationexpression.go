package codegenerator

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
		v.Visit(ctx.ElementExpression(index))

		v.Visit(ctx.MultiplicationFactor(index - 1))

		index++
	}

	return nil
}
