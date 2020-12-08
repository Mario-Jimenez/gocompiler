package codegenerator

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	expressionList:
		expression (COMMA expression)* # expressionListTree
		;
*/

func (v *visitor) VisitExpressionListTree(ctx *parser.ExpressionListTreeContext) interface{} {
	if v.array != nil {
		v.array.addIndex(v.instructionIndex)
	}

	v.Visit(ctx.Expression(0))

	if v.array != nil {
		if !v.array.isFunction(0) {
			v.addInstruction("RETURN_ACCESS", "")
		}
	}

	totalBranches := len(ctx.AllExpression())
	index := 1
	for index < totalBranches {
		if v.array != nil {
			v.array.addIndex(v.instructionIndex)
		}

		v.Visit(ctx.Expression(index))

		if v.array != nil {
			if !v.array.isFunction(index) {
				v.addInstruction("RETURN_ACCESS", "")
			}
		}

		index++
	}

	return nil
}
