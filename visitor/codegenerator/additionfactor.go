package codegenerator

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	additionFactor : PLUS # plusOperator | MINUS # minusOperator;
*/

func (v *visitor) VisitPlusOperator(ctx *parser.PlusOperatorContext) interface{} {
	v.addInstruction("BINARY_ADD", "+")

	return nil
}

func (v *visitor) VisitMinusOperator(ctx *parser.MinusOperatorContext) interface{} {
	v.addInstruction("BINARY_SUBSTRACT", "-")

	return nil
}
