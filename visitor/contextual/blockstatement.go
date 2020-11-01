package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	blockStatement : L_CURLY statement* R_CURLY # blockTree;
*/

func (v *visitor) VisitBlockTree(ctx *parser.BlockTreeContext) interface{} {
	v.generalTable.OpenScope()
	v.functionTable.OpenScope()

	for _, branch := range ctx.AllStatement() {
		v.Visit(branch)
	}

	v.generalTable.CloseScope()
	v.functionTable.CloseScope()

	return nil
}
