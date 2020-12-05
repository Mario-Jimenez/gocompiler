package codegenerator

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	multiplicationFactor:
		MULT	# multiplicationOperator
		| DIV	# divisionOperator
		;
*/

func (v *visitor) VisitMultiplicationOperator(ctx *parser.MultiplicationOperatorContext) interface{} {
	v.addInstruction("BINARY_MULTIPLY", "*")

	return nil
}

func (v *visitor) VisitDivisionOperator(ctx *parser.DivisionOperatorContext) interface{} {
	v.addInstruction("BINARY_DIVIDE", "/")

	return nil
}
