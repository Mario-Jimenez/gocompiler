package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	hashContent : expression COLON expression # hashPairTree;
*/

func (v *contextualVisitor) VisitHashPairTree(ctx *parser.HashPairTreeContext) interface{} {
	if ctx.Expression(0) != nil {
		v.Visit(ctx.Expression(0))
		// node :=
		// node.(*visitResponse)
	}

	if ctx.COLON() != nil {
		ctx.COLON().GetText()
	}

	if ctx.Expression(1) != nil {
		v.Visit(ctx.Expression(1))
		// node :=
		// node.(*visitResponse)
	}

	return nil
}
