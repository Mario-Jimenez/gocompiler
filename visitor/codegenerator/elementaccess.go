package codegenerator

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	elementAccess:
		L_BRACKET expression R_BRACKET # elementAccessTree
		;
*/

func (v *visitor) VisitElementAccessTree(ctx *parser.ElementAccessTreeContext) interface{} {
	v.access.accessElement = true

	v.Visit(ctx.Expression())

	if v.access.global {
		if v.access.isFunction() {
			v.addInstruction("ACCESS_FUNCTION_GLOBAL", v.access.name)
			v.addInstruction("RETURN", "")
		} else {
			v.addInstruction("ACCESS_ELEMENT_GLOBAL", v.access.name)
		}
	} else {
		if v.access.isFunction() {
			v.addInstruction("ACCESS_FUNCTION_LOCAL", v.access.name)
			v.addInstruction("RETURN", "")
		} else {
			v.addInstruction("ACCESS_ELEMENT_LOCAL", v.access.name)
		}
	}

	return nil
}
