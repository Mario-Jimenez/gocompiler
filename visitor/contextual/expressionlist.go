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
	// if we're working with an identifier that is a function call
	// we store the total parameters passed to that function
	if v.identifier.getType() == ICALL {
		v.identifier.setParameters(len(ctx.AllExpression()))
	}

	v.Visit(ctx.Expression(0))

	// handle function inside array
	if !v.array.isNestedArray() && v.array.isActive() {
		v.declaration.addArrayElement(0, v.array.getParameters())
		v.array.reset()
	}

	totalBranches := len(ctx.AllExpression())
	index := 1
	for index < totalBranches {
		v.Visit(ctx.Expression(index))

		// handle function inside array
		if !v.array.isNestedArray() && v.array.isActive() {
			v.declaration.addArrayElement(index, v.array.getParameters())
			v.array.reset()
		}

		index++
	}

	return nil
}
