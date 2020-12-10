package codegenerator

import (
	"fmt"

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

	totalBranches := len(ctx.AllExpression())

	var jumpIndex int
	if v.isArrayFunction && totalBranches > 1 {
		jumpIndex = v.instructionIndex
		v.addInstruction("JUMP_ABSOLUTE", "")
	}

	v.Visit(ctx.Expression(0))

	if v.isArrayFunction && totalBranches > 1 {
		v.addInstruction("RETURN_ACCESS", "")

		// backpatching
		v.updateInstruction(jumpIndex, fmt.Sprintf("%d", v.instructionIndex))

		v.addInstruction("LOAD_CONST", fmt.Sprintf("%d", jumpIndex+1))
	}

	if v.array != nil {
		if !v.array.isFunction(0) {
			v.addInstruction("RETURN_ACCESS", "")
		}
	}
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
