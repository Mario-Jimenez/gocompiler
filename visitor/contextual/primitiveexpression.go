package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/identification"
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
	if v.hash.getType() == HUNKNOWN {
		v.hash.setToken(ctx.INTEGER().GetSymbol())
		v.hash.setType(HINTEGER)
	}

	return nil
}

func (v *visitor) VisitString(ctx *parser.StringContext) interface{} {
	if v.hash.getType() == HUNKNOWN {
		v.hash.setToken(ctx.STRING().GetSymbol())
		v.hash.setType(HSTRING)
	}

	return nil
}

func (v *visitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	v.identifier.setType(IIDENTIFIER)
	v.identifier.setToken(ctx.IDENTIFIER().GetSymbol())

	v.hash.setType(HCOMPLEX)

	return nil
}

func (v *visitor) VisitTrue(ctx *parser.TrueContext) interface{} {
	v.hash.setType(HCOMPLEX)

	return nil
}

func (v *visitor) VisitFalse(ctx *parser.FalseContext) interface{} {
	v.hash.setType(HCOMPLEX)

	return nil
}

func (v *visitor) VisitGroupedExpressionTree(ctx *parser.GroupedExpressionTreeContext) interface{} {
	v.hash.setType(HCOMPLEX)

	v.Visit(ctx.Expression())

	return nil
}

func (v *visitor) VisitArrayTree(ctx *parser.ArrayTreeContext) interface{} {
	v.hash.setType(HCOMPLEX)

	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
	}

	return nil
}

func (v *visitor) VisitArrayFunctionTree(ctx *parser.ArrayFunctionTreeContext) interface{} {
	v.hash.setType(HCOMPLEX)

	v.Visit(ctx.ArrayFunctions())

	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
	}

	return nil
}

func (v *visitor) VisitFunctionTree(ctx *parser.FunctionTreeContext) interface{} {
	v.hash.setType(HCOMPLEX)

	v.declaration.setType(DFUNCTION)

	v.Visit(ctx.FunctionParameters())

	v.Visit(ctx.BlockStatement())

	v.table.CloseScope()

	return nil
}

func (v *visitor) VisitHashObjectTree(ctx *parser.HashObjectTreeContext) interface{} {
	v.declaration.setType(DHASH)

	v.hash.setType(HCOMPLEX)

	v.Visit(ctx.HashContent(0))

	totalBranches := len(ctx.AllHashContent())
	index := 1
	for index < totalBranches {
		v.Visit(ctx.HashContent(index))

		index++
	}

	if !v.declaration.isNestedHash() {
		token := v.declaration.getToken()
		if token != nil {
			hashData := identification.NewHashData()
			for _, k := range v.declaration.getHashKeys() {
				hashData.AddKey(k.token)
			}
			attr := identification.NewAttribute(identification.HASH, token, hashData)
			v.table.Enter(token.GetText(), attr)
		}
	}

	return nil
}

func (v *visitor) VisitPrintTree(ctx *parser.PrintTreeContext) interface{} {
	v.hash.setType(HCOMPLEX)

	v.Visit(ctx.Expression())

	return nil
}

func (v *visitor) VisitConditionalTree(ctx *parser.ConditionalTreeContext) interface{} {
	v.hash.setType(HCOMPLEX)

	v.Visit(ctx.Expression())

	v.Visit(ctx.BlockStatement(0))

	if ctx.ELSE() != nil {
		v.Visit(ctx.BlockStatement(1))
	}

	return nil
}
