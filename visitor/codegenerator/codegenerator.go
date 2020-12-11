package codegenerator

import (
	"fmt"
)

// visitor implementation of MonkeyParserVisitor interface
// to generate code
// methods are located in smaller files with rules names
type visitor struct {
	code             map[int]*instruction
	instructionIndex int

	call   *callHelper
	array  *arrayHelper
	access *accessHelper
	hash   *hashHelper

	isArrayFunction bool
}

// instruction info
type instruction struct {
	name      string
	parameter string
}

// NewVisitor instance
func NewVisitor() *visitor {
	return &visitor{
		code: map[int]*instruction{},
	}
}

// get code instructions
func (v *visitor) Code() string {
	var instructions string
	lines := len(v.code)
	index := 0
	for index < lines {
		code := v.code[index]
		if code.parameter == "" {
			instructions += fmt.Sprintf("%d %s\n", index, code.name)
		} else {
			instructions += fmt.Sprintf("%d %s %s\n", index, code.name, code.parameter)
		}
		index++
	}

	return instructions
}

// add code instruction
func (v *visitor) addInstruction(name, parameter string) {
	v.code[v.instructionIndex] = &instruction{
		name:      name,
		parameter: parameter,
	}

	v.instructionIndex++
}

// update code instruction
func (v *visitor) updateInstruction(index int, parameter string) {
	instruction := v.code[index]
	instruction.parameter = parameter
}
