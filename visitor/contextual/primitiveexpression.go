package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
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
	// when working with a hash expression and the hash key type is still unknown
	// set hash key as integer
	if v.hash.getType() == HUNKNOWN {
		v.hash.setToken(ctx.INTEGER().GetSymbol())
		v.hash.setType(HINTEGER)
	}

	return nil
}

func (v *visitor) VisitString(ctx *parser.StringContext) interface{} {
	// when working with a hash expression and the hash key type is still unknown
	// set hash key as string
	if v.hash.getType() == HUNKNOWN {
		v.hash.setToken(ctx.STRING().GetSymbol())
		v.hash.setType(HSTRING)
	}

	return nil
}

func (v *visitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	// mark expression as an identifier
	v.identifier.setType(IIDENTIFIER)
	v.identifier.setToken(ctx.IDENTIFIER().GetSymbol())

	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	return nil
}

func (v *visitor) VisitTrue(ctx *parser.TrueContext) interface{} {
	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	return nil
}

func (v *visitor) VisitFalse(ctx *parser.FalseContext) interface{} {
	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	return nil
}

func (v *visitor) VisitGroupedExpressionTree(ctx *parser.GroupedExpressionTreeContext) interface{} {
	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	v.Visit(ctx.Expression())

	return nil
}

func (v *visitor) VisitArrayTree(ctx *parser.ArrayTreeContext) interface{} {
	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	// mark declaration as array
	v.declaration.setType(DARRAY)

	v.array.newArray()

	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
	}

	// check that we're not working with a nested array
	// and store array in identification table
	if !v.array.isNestedArray() {
		token := v.declaration.getToken()
		// check if we're working on a declaration
		if token != nil {
			arrayData := identification.NewArrayData()
			for _, k := range v.declaration.getArrayElements() {
				arrayData.AddElement(k.index, k.parameters)
			}
			attr := identification.NewAttribute(identification.ARRAY, token, arrayData)
			v.table.Enter(token.GetText(), attr)
		}
	}

	v.array.closeArray()

	return nil
}

func (v *visitor) VisitArrayFunctionTree(ctx *parser.ArrayFunctionTreeContext) interface{} {
	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	// v.Visit(ctx.ArrayFunctions())

	if ctx.ExpressionList() != nil {
		v.Visit(ctx.ExpressionList())
	}

	return nil
}

func (v *visitor) VisitFunctionTree(ctx *parser.FunctionTreeContext) interface{} {
	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	// handle nested functions
	// we create a backup of the current declaration
	// (hash or array declarations that contain a function expression in them)
	// because the function tree will replace these values
	var tempType DeclarationType
	var tempToken antlr.Token
	var restore bool
	if v.declaration.getType() != DNEUTRAL {
		restore = true
		tempType = v.declaration.getType()
		tempToken = v.declaration.getToken()
	}

	// mark declaration as function
	v.declaration.setType(DFUNCTION)

	v.Visit(ctx.FunctionParameters())

	v.Visit(ctx.BlockStatement())

	v.table.CloseScope()

	// we reload the backup values of the previous declaration
	if restore {
		v.declaration.setType(tempType)
		v.declaration.setToken(tempToken)
	}

	return nil
}

func (v *visitor) VisitHashObjectTree(ctx *parser.HashObjectTreeContext) interface{} {
	// mark declaration as hash
	v.declaration.setType(DHASH)

	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	v.Visit(ctx.HashContent(0))

	totalBranches := len(ctx.AllHashContent())
	index := 1
	for index < totalBranches {
		v.Visit(ctx.HashContent(index))

		index++
	}

	// check that we're not working with a nested hash
	// and store hash in identification table
	if !v.declaration.isNestedHash() {
		token := v.declaration.getToken()
		// check if we're working on a declaration
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
	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	v.Visit(ctx.Expression())

	return nil
}

func (v *visitor) VisitConditionalTree(ctx *parser.ConditionalTreeContext) interface{} {
	// when working with a hash expression, tell hash key is not integer or string
	v.hash.setType(HCOMPLEX)

	v.Visit(ctx.Expression())

	v.Visit(ctx.BlockStatement(0))

	if ctx.ELSE() != nil {
		v.Visit(ctx.BlockStatement(1))
	}

	return nil
}
