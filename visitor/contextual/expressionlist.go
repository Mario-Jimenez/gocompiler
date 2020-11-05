package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	expressionList:
		expression (COMMA expression)* # expressionListTree
		;
*/

func (v *visitor) VisitExpressionListTree(ctx *parser.ExpressionListTreeContext) interface{} {
	if v.identifier.getType() == ICALL {
		v.identifier.setParameters(len(ctx.AllExpression()))
	}

	v.Visit(ctx.Expression(0))

	totalBranches := len(ctx.AllExpression())
	index := 1
	for index < totalBranches {
		if ctx.Expression(index) != nil {
			v.Visit(ctx.Expression(index))
		}

		index++
	}

	return nil
}
