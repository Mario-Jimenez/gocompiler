package contextual

import (
	"fmt"

	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	statement:
		LET IDENTIFIER ASSIGN expression SEMI?	# letStatementTree
		| RETURN expression SEMI?				# returnStatementTree
		| expression SEMI?						# expressionStatementTree
		;
*/

func (v *visitor) VisitLetStatementTree(ctx *parser.LetStatementTreeContext) interface{} {
	v.declaration.newDeclaration()

	token := ctx.IDENTIFIER().GetSymbol()
	v.declaration.setToken(token)

	v.Visit(ctx.Expression())

	if v.declaration.getType() == DNEUTRAL {
		attr := identification.NewAttribute(identification.IDENTIFIER, token, nil)
		v.table.Enter(ctx.IDENTIFIER().GetText(), attr)
	}

	v.declaration.closeDeclaration()

	return nil
}

func (v *visitor) VisitReturnStatementTree(ctx *parser.ReturnStatementTreeContext) interface{} {
	if v.declaration.getType() != DFUNCTION {
		// error: return outside function declaration
		token := ctx.RETURN().GetSymbol()
		newError := fmt.Sprintf("line %d:%d return outside function declaration", token.GetLine(), token.GetColumn())
		v.table.AddError(newError, token.GetLine())
	}

	v.Visit(ctx.Expression())

	return nil
}

func (v *visitor) VisitExpressionStatementTree(ctx *parser.ExpressionStatementTreeContext) interface{} {
	v.Visit(ctx.Expression())

	return nil
}
