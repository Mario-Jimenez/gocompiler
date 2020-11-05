package contextual

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	multiplicationFactor:
		MULT	# multiplicationOperator
		| DIV	# divisionOperator
		;
*/

func (v *visitor) VisitMultiplicationOperator(ctx *parser.MultiplicationOperatorContext) interface{} {
	return nil
}

func (v *visitor) VisitDivisionOperator(ctx *parser.DivisionOperatorContext) interface{} {
	return nil
}
