package graph

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	additionFactor : PLUS # plusOperator | MINUS # minusOperator;
*/

func (v *visitor) VisitPlusOperator(ctx *parser.PlusOperatorContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.PLUS().GetText(), terminalColor),
		failed: false,
	}
}

func (v *visitor) VisitMinusOperator(ctx *parser.MinusOperatorContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.MINUS().GetText(), terminalColor),
		failed: false,
	}
}
