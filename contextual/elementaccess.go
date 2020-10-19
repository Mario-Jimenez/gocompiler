package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	elementAccess:
		L_BRACKET expression R_BRACKET # elementAccessTree
		;
*/

func (v *contextualVisitor) VisitElementAccessTree(ctx *parser.ElementAccessTreeContext) interface{} {
	ctx.L_BRACKET().GetText()

	v.Visit(ctx.Expression())
	// node :=
	// node.(*visitResponse)

	if ctx.R_BRACKET() != nil {
		ctx.R_BRACKET().GetText()
	}

	return nil
}
