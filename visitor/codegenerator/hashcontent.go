package codegenerator

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	hashContent : expression COLON expression # hashPairTree;
*/

func (v *visitor) VisitHashPairTree(ctx *parser.HashPairTreeContext) interface{} {
	v.hash.addIndex(v.instructionIndex)

	v.Visit(ctx.Expression(1))

	v.addInstruction("RETURN_ACCESS", "")

	return nil
}
