package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	blockStatement : L_CURLY statement* R_CURLY # blockTree;
*/

func (v *contextualVisitor) VisitBlockTree(ctx *parser.BlockTreeContext) interface{} {
	v.identificationTable.OpenScope()

	for _, branch := range ctx.AllStatement() {
		v.Visit(branch)
		// node :=
		// node.(*visitResponse)
	}

	v.identificationTable.CloseScope()

	return nil
}
