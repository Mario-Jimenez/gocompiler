package contextual

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	additionFactor : PLUS # plusOperator | MINUS # minusOperator;
*/

func (v *visitor) VisitPlusOperator(ctx *parser.PlusOperatorContext) interface{} {
	return nil
}

func (v *visitor) VisitMinusOperator(ctx *parser.MinusOperatorContext) interface{} {
	return nil
}
