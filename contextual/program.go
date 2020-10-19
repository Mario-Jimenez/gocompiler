package contextual

import (
	"fmt"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	program : statement* EOF # programTree;
*/

func (v *contextualVisitor) VisitProgramTree(ctx *parser.ProgramTreeContext) interface{} {
	v.identificationTable.OpenScope()

	for _, branch := range ctx.AllStatement() {
		v.Visit(branch)
		// node :=
		// node.(*visitResponse)
	}

	fmt.Printf("%+v\n", v.identificationTable)

	v.identificationTable.CloseScope()

	fmt.Printf("%+v\n", v.identificationTable)

	return nil
}
