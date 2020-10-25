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

func (v *contextualVisitor) VisitLetStatementTree(ctx *parser.LetStatementTreeContext) interface{} {
	token := ctx.IDENTIFIER().GetSymbol()
	if v.isFunction {
		attr := identification.NewAttributeFunction(token, v.parametersCount)
		v.functionTable.Enter(ctx.IDENTIFIER().GetText(), attr)
		v.isFunction = false
	} else {
		attr := identification.NewAttribute(token, ctx)
		v.identificationTable.Enter(ctx.IDENTIFIER().GetText(), attr)
	}

	ctx.LET().GetText()

	if ctx.IDENTIFIER() != nil {
		ctx.IDENTIFIER().GetText()
	}

	if ctx.ASSIGN() != nil {
		ctx.ASSIGN().GetText()
	}

	if ctx.Expression() != nil {
		v.Visit(ctx.Expression())
		// node :=
		// node.(*visitResponse)
	}

	if ctx.SEMI() != nil {
		ctx.SEMI().GetText()
	}

	return nil
}

func (v *contextualVisitor) VisitReturnStatementTree(ctx *parser.ReturnStatementTreeContext) interface{} {
	ctx.RETURN().GetText()

	if ctx.Expression() != nil {
		v.Visit(ctx.Expression())
		// node :=
		// node.(*visitResponse)
	}

	if ctx.SEMI() != nil {
		ctx.SEMI().GetText()
	}

	return nil
}

func (v *contextualVisitor) VisitExpressionStatementTree(ctx *parser.ExpressionStatementTreeContext) interface{} {
	v.Visit(ctx.Expression())
	// node :=
	// node.(*visitResponse)

	if ctx.SEMI() != nil {
		ctx.SEMI().GetText()
	}

	return nil
}
