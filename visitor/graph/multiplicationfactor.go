package graph

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	multiplicationFactor:
		MULT	# multiplicationOperator
		| DIV	# divisionOperator
		;
*/

func (v *visitor) VisitMultiplicationOperator(ctx *parser.MultiplicationOperatorContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.MULT().GetText(), terminalColor),
		failed: false,
	}
}

func (v *visitor) VisitDivisionOperator(ctx *parser.DivisionOperatorContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.DIV().GetText(), terminalColor),
		failed: false,
	}
}
