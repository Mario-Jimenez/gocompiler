package codegenerator

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	program : statement* EOF # programTree;
*/

func (v *visitor) VisitProgramTree(ctx *parser.ProgramTreeContext) interface{} {
	for _, branch := range ctx.AllStatement() {
		v.Visit(branch)
	}

	v.addInstruction("END", "")

	return nil
}
