package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	program : statement* EOF # programTree;
*/

func (v *visitor) VisitProgramTree(ctx *parser.ProgramTreeContext) interface{} {
	v.table.OpenScope()

	for _, branch := range ctx.AllStatement() {
		v.Visit(branch)
	}

	v.table.CloseScope()

	return nil
}
