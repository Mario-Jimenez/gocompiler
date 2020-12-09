package codegenerator

import (
	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

/*
	statement:
		LET IDENTIFIER ASSIGN expression SEMI?	# letStatementTree
		| RETURN expression SEMI?				# returnStatementTree
		| expression SEMI?						# expressionStatementTree
		;
*/

func (v *visitor) VisitLetStatementTree(ctx *parser.LetStatementTreeContext) interface{} {
	token := v.Visit(ctx.LetIdent()).(antlr.Token)
	declaration := ctx.LetIdent().GetDeclaration().(identification.Declaration)
	if declaration.Expression() == identification.IDENTIFIER {
		if declaration.Level() == 1 {
			v.addInstruction("PUSH_GLOBAL", token.GetText())
		} else {
			v.addInstruction("PUSH_LOCAL", token.GetText())
		}
	}

	if declaration.Expression() == identification.FUNCTION {
		if declaration.Level() == 1 {
			v.addInstruction("DEF_GLOBAL", token.GetText())
		} else {
			v.addInstruction("DEF_LOCAL", token.GetText())
		}
	}

	if declaration.Expression() == identification.ARRAY {
		if declaration.Level() == 1 {
			v.addInstruction("PUSH_GLOBAL", token.GetText())
			v.array = newArrayHelper(true, token.GetText())
		} else {
			v.addInstruction("PUSH_LOCAL", token.GetText())
			v.array = newArrayHelper(false, token.GetText())
		}
		array := declaration.Data().(*identification.ArrayData)
		v.array.functionIndexes = array.GetFunctionIndexes()
	}

	if declaration.Expression() == identification.HASH {
		if declaration.Level() == 1 {
			v.addInstruction("PUSH_GLOBAL", token.GetText())
			v.hash = newHashHelper(true, token.GetText())
		} else {
			v.addInstruction("PUSH_LOCAL", token.GetText())
			v.hash = newHashHelper(false, token.GetText())
		}
		hash := declaration.Data().(*identification.HashData)
		v.hash.keys = hash.GetKeys()
	}

	v.Visit(ctx.Expression())

	if declaration.Expression() == identification.IDENTIFIER {
		if declaration.Level() == 1 {
			v.addInstruction("STORE_GLOBAL", token.GetText())
		} else {
			v.addInstruction("STORE_FAST", token.GetText())
		}
	}

	v.array = nil
	v.hash = nil

	return nil
}

func (v *visitor) VisitReturnStatementTree(ctx *parser.ReturnStatementTreeContext) interface{} {
	v.Visit(ctx.Expression())
	v.addInstruction("RETURN_VALUE", "")

	return nil
}

func (v *visitor) VisitExpressionStatementTree(ctx *parser.ExpressionStatementTreeContext) interface{} {
	v.Visit(ctx.Expression())

	return nil
}
