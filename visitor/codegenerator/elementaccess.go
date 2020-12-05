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
	v.Visit(ctx.Expression())

	return nil
}
