package codegenerator

import "github.com/Mario-Jimenez/gocompiler/parser"

// visitor implementation of MonkeyParserVisitor interface
// to generate code
// methods are located in smaller files with rules names
type visitor struct {
	code             map[int]*instruction
	instructionIndex int

	call *callHelper
}

type instruction struct {
	name      string
	parameter string
}

// NewVisitor instance
func NewVisitor() parser.MonkeyParserVisitor {
	return &visitor{
		code: map[int]*instruction{},
	}
}

// check if a declaration exists in global scope
func (v *visitor) isGlobal(declaration string) bool {
	return false
}

func (v *visitor) addInstruction(name, parameter string) {
	v.code[v.instructionIndex] = &instruction{
		name:      name,
		parameter: parameter,
	}

	v.instructionIndex++
}

func (v *visitor) updateInstruction(index int, parameter string) {
	instruction := v.code[index]
	instruction.parameter = parameter
}