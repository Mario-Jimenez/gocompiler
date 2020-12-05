package codegenerator

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	blockStatement : L_CURLY statement* R_CURLY # blockTree;
*/

func (v *visitor) VisitBlockTree(ctx *parser.BlockTreeContext) interface{} {
	for _, branch := range ctx.AllStatement() {
		v.Visit(branch)
	}

	return nil
}
