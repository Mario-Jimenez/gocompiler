package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	hashContent : expression COLON expression # hashPairTree;
*/

func (v *visitor) VisitHashPairTree(ctx *parser.HashPairTreeContext) interface{} {
	if ctx.Expression(0) != nil {
		v.Visit(ctx.Expression(0))
	}

	if ctx.Expression(1) != nil {
		v.Visit(ctx.Expression(1))
	}

	return nil
}
