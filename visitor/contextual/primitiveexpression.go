package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	primitiveExpression:
		INTEGER														# integer
		| STRING													# string
		| IDENTIFIER												# identifier
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

func (v *visitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
	return nil
}

func (v *visitor) VisitString(ctx *parser.StringContext) interface{} {
	return nil
}

func (v *visitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	v.identifier.newIdentifier()
	v.identifier.markIdentifier()
	v.identifier.setToken(ctx.IDENTIFIER().GetSymbol())

	return nil
}

func (v *visitor) VisitTrue(ctx *parser.TrueContext) interface{} {
	return nil
}

func (v *visitor) VisitFalse(ctx *parser.FalseContext) interface{} {
	return nil
}

func (v *visitor) VisitGroupedExpressionTree(ctx *parser.GroupedExpressionTreeContext) interface{} {
	v.Visit(ctx.Expression())

	return nil
}

func (v *visitor) VisitArrayTree(ctx *parser.ArrayTreeContext) interface{} {
	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
	}

	return nil
}

func (v *visitor) VisitArrayFunctionTree(ctx *parser.ArrayFunctionTreeContext) interface{} {
	v.Visit(ctx.ArrayFunctions())

	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
	}

	return nil
}

func (v *visitor) VisitFunctionTree(ctx *parser.FunctionTreeContext) interface{} {
	v.generalTable.OpenScope()

	v.declaration.markFunction()

	if ctx.FunctionParameters() != nil {
		v.Visit(ctx.FunctionParameters())
	}

	if ctx.BlockStatement() != nil {
		v.Visit(ctx.BlockStatement())
	}

	v.generalTable.CloseScope()

	return nil
}

func (v *visitor) VisitHashObjectTree(ctx *parser.HashObjectTreeContext) interface{} {
	v.Visit(ctx.HashContent(0))

	totalBranches := len(ctx.AllHashContent())
	index := 1
	for index < totalBranches {
		v.Visit(ctx.HashContent(index))

		index++
	}

	return nil
}

func (v *visitor) VisitPrintTree(ctx *parser.PrintTreeContext) interface{} {
	if ctx.Expression() != nil {
		v.Visit(ctx.Expression())
	}

	return nil
}

func (v *visitor) VisitConditionalTree(ctx *parser.ConditionalTreeContext) interface{} {
	v.Visit(ctx.Expression())

	v.Visit(ctx.BlockStatement(0))

	if ctx.ELSE() != nil {
		v.Visit(ctx.BlockStatement(1))
	}

	return nil
}
