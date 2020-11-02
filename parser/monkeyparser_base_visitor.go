// Code generated from ./grammar/MonkeyParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MonkeyParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMonkeyParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMonkeyParserVisitor) VisitProgramTree(ctx *ProgramTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitLetStatementTree(ctx *LetStatementTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitReturnStatementTree(ctx *ReturnStatementTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitExpressionStatementTree(ctx *ExpressionStatementTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitExpressionTree(ctx *ExpressionTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitLessComparison(ctx *LessComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitGreaterComparison(ctx *GreaterComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitLessOrEqualsComparison(ctx *LessOrEqualsComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitGreaterOrEqualsComparison(ctx *GreaterOrEqualsComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitEqualsComparison(ctx *EqualsComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitAdditionTree(ctx *AdditionTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitPlusOperator(ctx *PlusOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitMinusOperator(ctx *MinusOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitMultiplicationTree(ctx *MultiplicationTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitMultiplicationOperator(ctx *MultiplicationOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitDivisionOperator(ctx *DivisionOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitElementTree(ctx *ElementTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitElementAccessTree(ctx *ElementAccessTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitFunctionCallTree(ctx *FunctionCallTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitTrue(ctx *TrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitFalse(ctx *FalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitGroupedExpressionTree(ctx *GroupedExpressionTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitArrayTree(ctx *ArrayTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitArrayFunctionTree(ctx *ArrayFunctionTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitFunctionTree(ctx *FunctionTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitHashObjectTree(ctx *HashObjectTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitPrintTree(ctx *PrintTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitConditionalTree(ctx *ConditionalTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitArrayLen(ctx *ArrayLenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitArrayFirst(ctx *ArrayFirstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitArrayLast(ctx *ArrayLastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitArrayRest(ctx *ArrayRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitArrayPush(ctx *ArrayPushContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitFunctionParametersTree(ctx *FunctionParametersTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitHashPairTree(ctx *HashPairTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitExpressionListTree(ctx *ExpressionListTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitBlockTree(ctx *BlockTreeContext) interface{} {
	return v.VisitChildren(ctx)
}
