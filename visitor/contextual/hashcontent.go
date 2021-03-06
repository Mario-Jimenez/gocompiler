package contextual

import (
	"fmt"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	hashContent : expression COLON expression # hashPairTree;
*/

func (v *visitor) VisitHashPairTree(ctx *parser.HashPairTreeContext) interface{} {
	v.hash.newHash()

	v.Visit(ctx.Expression(0))

	// check we're not working with a nested hash
	if !v.declaration.isNestedHash() {
		// hash key must be integer or string
		if v.hash.getType() == HCOMPLEX || v.hash.getType() == HUNKNOWN {
			token := ctx.COLON().GetSymbol()
			newError := fmt.Sprintf("line %d:%d hash key must be INTEGER or STRING", token.GetLine(), token.GetColumn())
			v.table.AddError(newError, token.GetLine())
		} else {
			// check that hash key is not repeated
			if !v.declaration.addHashKey(v.hash.getToken(), v.hash.getType()) {
				token := v.hash.getToken()
				newError := fmt.Sprintf("line %d:%d hash key %s already exists", token.GetLine(), token.GetColumn(), token.GetText())
				v.table.AddError(newError, token.GetLine())
			}
		}
	}

	v.hash.closeHash()

	v.declaration.incNestedHash()

	v.Visit(ctx.Expression(1))

	v.declaration.decNestedHash()

	return nil
}
