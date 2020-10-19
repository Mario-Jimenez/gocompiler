package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	expressionList:
		expression (COMMA expression)* # expressionListTree
		;
*/

func (v *contextualVisitor) VisitExpressionListTree(ctx *parser.ExpressionListTreeContext) interface{} {
	v.Visit(ctx.Expression(0))
	// node :=
	// node.(*visitResponse)

	totalBranches := len(ctx.AllExpression())
	index := 1
	for index < totalBranches {
		ctx.COMMA(index - 1).GetText()

		if ctx.Expression(index) != nil {
			v.Visit(ctx.Expression(index))
			// node :=
			// node.(*visitResponse)
		}

		index++
	}

	return nil
}
