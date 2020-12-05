package contextual

import "github.com/Mario-Jimenez/gocompiler/parser"

func (v *visitor) VisitLetIdentifier(ctx *parser.LetIdentifierContext) interface{} {
	return ctx.IDENTIFIER().GetSymbol()
}
