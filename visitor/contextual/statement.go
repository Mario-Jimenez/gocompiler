package contextual

import (
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

	v.Visit(ctx.Expression())

	token := ctx.IDENTIFIER().GetSymbol()

	if v.declaration.isFunction() {
		attr := identification.NewFunctionAttribute(token, v.declaration.getParameters())
		v.functionTable.Enter(ctx.IDENTIFIER().GetText(), attr)
		v.declaration.closeDeclaration()

		return nil
	}

	attr := identification.NewGeneralAttribute(token)
	v.generalTable.Enter(ctx.IDENTIFIER().GetText(), attr, false)

	v.declaration.closeDeclaration()

	return nil
}

func (v *visitor) VisitReturnStatementTree(ctx *parser.ReturnStatementTreeContext) interface{} {
	if !v.declaration.isFunction() {
		// TODO: error: return outside function definition
	}

	v.Visit(ctx.Expression())

	return nil
}

func (v *visitor) VisitExpressionStatementTree(ctx *parser.ExpressionStatementTreeContext) interface{} {
	v.Visit(ctx.Expression())

	return nil
}
