package contextual

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	additionFactor : PLUS # plusOperator | MINUS # minusOperator;
*/

func (v *contextualVisitor) VisitPlusOperator(ctx *parser.PlusOperatorContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitMinusOperator(ctx *parser.MinusOperatorContext) interface{} {
	return nil
}
