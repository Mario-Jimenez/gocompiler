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

	v.Visit(ctx.Expression())

	if declaration.Expression() == identification.IDENTIFIER {
		if declaration.Level() == 1 {
			v.addInstruction("STORE_GLOBAL", token.GetText())
		} else {
			v.addInstruction("STORE_FAST", token.GetText())
		}
	}

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
