package codegenerator

import (
	"fmt"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	elementExpression:
		primitiveExpression (elementAccess | callExpression)? # elementTree
		;
*/

func (v *visitor) VisitElementTree(ctx *parser.ElementTreeContext) interface{} {
	v.Visit(ctx.PrimitiveExpression())

	var isAccess bool
	if ctx.ElementAccess() != nil {
		isAccess = true
		v.Visit(ctx.ElementAccess())
	} else if ctx.CallExpression() != nil {
		v.Visit(ctx.CallExpression())
		if v.call.global {
			v.addInstruction("LOAD_GLOBAL", v.call.name)
		} else {
			v.addInstruction("LOAD_FAST", v.call.name)
		}
		v.addInstruction("CALL_FUNCTION", fmt.Sprintf("%d", v.call.parameters))
		v.call = nil
	}

	if v.access != nil && !v.access.accessElement {
		if v.access.global {
			v.addInstruction("LOAD_GLOBAL", v.access.name)
		} else {
			v.addInstruction("LOAD_FAST", v.access.name)
		}
		v.access = nil
	}

	if isAccess {
		v.access = nil
	}

	return nil
}
