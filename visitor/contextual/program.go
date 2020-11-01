package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	program : statement* EOF # programTree;
*/

func (v *visitor) VisitProgramTree(ctx *parser.ProgramTreeContext) interface{} {
	v.generalTable.OpenScope()
	v.functionTable.OpenScope()

	for _, branch := range ctx.AllStatement() {
		v.Visit(branch)
	}

	v.generalTable.CloseScope()
	v.functionTable.CloseScope()

	return nil
}
