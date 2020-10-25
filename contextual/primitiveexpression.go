package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	primitiveExpression:
		INTEGER														# integer
		| STRING													# string
		| identifier												# identifierTree
		| TRUE														# true
		| FALSE														# false
		| L_PAREN expression R_PAREN								# groupedExpressionTree
		| L_BRACKET expressionList? R_BRACKET						# arrayTree
		| arrayFunctions L_PAREN expressionList? R_PAREN			# arrayFunctionTree
		| FUNC L_PAREN functionParameters R_PAREN blockStatement	# functionTree
		| L_CURLY hashContent (COMMA hashContent)* R_CURLY			# hashObjectTree
		| PUTS L_PAREN expression R_PAREN							# printTree
		| IF expression blockStatement (ELSE blockStatement)?		# conditionalTree
		;
*/

func (v *contextualVisitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitString(ctx *parser.StringContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitIdentifierTree(ctx *parser.IdentifierTreeContext) interface{} {
	v.Visit(ctx.Identifier())
	// node :=
	// node.(*visitResponse)

	return nil
}

func (v *contextualVisitor) VisitTrue(ctx *parser.TrueContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitFalse(ctx *parser.FalseContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitGroupedExpressionTree(ctx *parser.GroupedExpressionTreeContext) interface{} {
	ctx.L_PAREN().GetText()

	v.Visit(ctx.Expression())
	// node :=
	// node.(*visitResponse)

	if ctx.R_PAREN() != nil {
		ctx.R_PAREN().GetText()
	}

	return nil
}

func (v *contextualVisitor) VisitArrayTree(ctx *parser.ArrayTreeContext) interface{} {
	ctx.L_BRACKET().GetText()

	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
		// node :=
		// node.(*visitResponse)
	}

	if ctx.R_BRACKET() != nil {
		ctx.R_BRACKET().GetText()
	}

	return nil
}

func (v *contextualVisitor) VisitArrayFunctionTree(ctx *parser.ArrayFunctionTreeContext) interface{} {
	v.Visit(ctx.ArrayFunctions())
	// node :=
	// node.(*visitResponse)

	if ctx.L_PAREN() != nil {
		ctx.L_PAREN().GetText()
	}

	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
		// node =
		// node.(*visitResponse)
	}

	if ctx.R_PAREN() != nil {
		ctx.R_PAREN().GetText()
	}

	return nil
}

func (v *contextualVisitor) VisitFunctionATree(ctx *parser.FunctionATreeContext) interface{} {
	v.identificationTable.OpenScope()

	//TODO: REVISAR
	v.Visit(ctx.Function())

	if ctx.L_PAREN() != nil {
		ctx.L_PAREN().GetText()
	}

	if ctx.FunctionParameters() != nil {
		v.Visit(ctx.FunctionParameters())
		// node :=
		// node.(*visitResponse)
	}

	if ctx.R_PAREN() != nil {
		ctx.R_PAREN().GetText()
	}

	v.isFunction = true

	if ctx.BlockStatement() != nil {
		v.Visit(ctx.BlockStatement())
		// node :=
		// node.(*visitResponse)
	}

	v.identificationTable.CloseScope()

	return nil
}

func (v *contextualVisitor) VisitHashObjectTree(ctx *parser.HashObjectTreeContext) interface{} {
	ctx.L_CURLY().GetText()

	v.Visit(ctx.HashContent(0))
	// node :=
	// node.(*visitResponse)

	totalBranches := len(ctx.AllHashContent())
	index := 1
	for index < totalBranches {
		ctx.COMMA(index - 1).GetText()

		v.Visit(ctx.HashContent(index))
		// node =
		// node.(*visitResponse)

		index++
	}

	if ctx.R_CURLY() != nil {
		ctx.R_CURLY().GetText()
	}

	return nil
}

func (v *contextualVisitor) VisitPrintTree(ctx *parser.PrintTreeContext) interface{} {
	ctx.PUTS().GetText()

	if ctx.L_PAREN() != nil {
		ctx.L_PAREN().GetText()
	}

	if ctx.Expression() != nil {
		v.Visit(ctx.Expression())
		// node :=
		// node.(*visitResponse)
	}

	if ctx.R_PAREN() != nil {
		ctx.R_PAREN().GetText()
	}

	return nil
}

func (v *contextualVisitor) VisitConditionalTree(ctx *parser.ConditionalTreeContext) interface{} {
	ctx.IF().GetText()

	v.Visit(ctx.Expression())
	// node :=
	// node.(*visitResponse)

	v.Visit(ctx.BlockStatement(0))
	// node =
	// node.(*visitResponse)

	if ctx.ELSE() != nil {
		ctx.ELSE().GetText()

		v.Visit(ctx.BlockStatement(1))
		// node =
		// node.(*visitResponse)
	}

	return nil
}
