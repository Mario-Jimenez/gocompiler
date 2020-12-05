package codegenerator

import (
	"fmt"

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

	lines := len(v.code)
	index := 0
	for index < lines {
		code := v.code[index]
		if code.parameter == "" {
			fmt.Printf("%d %s\n", index, code.name)
		} else {
			fmt.Printf("%d %s %s\n", index, code.name, code.parameter)
		}
		index++
	}

	return nil
}
