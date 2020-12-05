package codegenerator

import (
	"fmt"

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
	v.addInstruction("LOAD_CONST", ctx.INTEGER().GetText())

	return nil
}

func (v *visitor) VisitString(ctx *parser.StringContext) interface{} {
	v.addInstruction("LOAD_CONST", ctx.STRING().GetText())

	return nil
}

func (v *visitor) VisitIdentifierTree(ctx *parser.IdentifierTreeContext) interface{} {
	token := v.Visit(ctx.Identifier()).(antlr.Token)
	declarationContext := ctx.Identifier().GetDeclaration()
	if declarationContext != nil {
		// let declaration
		declaration := declarationContext.LetIdent().GetDeclaration().(identification.Declaration)
		// TODO: check other types of identifiers
		if declaration.Expression() == identification.IDENTIFIER {
			if declaration.Level() == 1 {
				v.addInstruction("LOAD_GLOBAL", token.GetText())
			} else {
				v.addInstruction("LOAD_FAST", token.GetText())
			}
		} else if declaration.Expression() == identification.FUNCTION {
			var call *callHelper
			function := declaration.Data().(*identification.FunctionData)
			if declaration.Level() == 1 {
				call = newCallHelper(true, token.GetText(), function.GetParameters(), function.HasReturn())
			} else {
				call = newCallHelper(false, token.GetText(), function.GetParameters(), function.HasReturn())
			}
			v.call = call
		}
	} else {
		// function parameters
		v.addInstruction("LOAD_FAST", token.GetText())
	}

	return nil
}

func (v *visitor) VisitTrue(ctx *parser.TrueContext) interface{} {
	v.addInstruction("LOAD_CONST", "true")

	return nil
}

func (v *visitor) VisitFalse(ctx *parser.FalseContext) interface{} {
	v.addInstruction("LOAD_CONST", "false")

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

	v.Visit(ctx.FunctionParameters())
	v.Visit(ctx.BlockStatement())

	v.addInstruction("RETURN", "")

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
	v.Visit(ctx.Expression())

	v.addInstruction("LOAD_GLOBAL", "write")
	v.addInstruction("CALL_FUNCTION", "1")

	return nil
}

func (v *visitor) VisitConditionalTree(ctx *parser.ConditionalTreeContext) interface{} {
	v.Visit(ctx.Expression())

	conditionIndex := v.instructionIndex
	v.addInstruction("JUMP_IF_FALSE", "")

	v.Visit(ctx.BlockStatement(0))

	trueIndex := v.instructionIndex
	v.addInstruction("JUMP_ABSOLUTE", "")

	// backpatching
	v.updateInstruction(conditionIndex, fmt.Sprintf("%d", v.instructionIndex))

	if ctx.ELSE() != nil {
		v.Visit(ctx.BlockStatement(1))
	}

	// backpatching
	v.updateInstruction(trueIndex, fmt.Sprintf("%d", v.instructionIndex))

	return nil
}
