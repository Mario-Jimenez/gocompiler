package codegenerator

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	hashContent : expression COLON expression # hashPairTree;
*/

func (v *visitor) VisitHashPairTree(ctx *parser.HashPairTreeContext) interface{} {
	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	return nil
}
