package contextual

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	multiplicationFactor:
		MULT	# multiplicationOperator
		| DIV	# divisionOperator
		;
*/

func (v *contextualVisitor) VisitMultiplicationOperator(ctx *parser.MultiplicationOperatorContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitDivisionOperator(ctx *parser.DivisionOperatorContext) interface{} {
	return nil
}
