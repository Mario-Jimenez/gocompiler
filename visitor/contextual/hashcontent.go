package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	hashContent : expression COLON expression # hashPairTree;
*/

func (v *visitor) VisitHashPairTree(ctx *parser.HashPairTreeContext) interface{} {
	v.hash.newHash()

	v.Visit(ctx.Expression(0))

	/*if !v.hash.isValid() {
		// error: first expression must be integer or string
		token := ctx.COLON().GetSymbol()
		newError := fmt.Sprintf("line %d:%d hash keys must be INTEGER or STRING", token.GetLine(), token.GetColumn())
		v.generalTable.AddError(newError, token.GetLine())
	}*/

	v.hash.closeHash()

	v.Visit(ctx.Expression(1))

	return nil
}
