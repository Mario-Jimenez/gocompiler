// Code generated from ./grammar/MonkeyParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MonkeyParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by MonkeyParser.
type MonkeyParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MonkeyParser#programTree.
	VisitProgramTree(ctx *ProgramTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#letStatementTree.
	VisitLetStatementTree(ctx *LetStatementTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#returnStatementTree.
	VisitReturnStatementTree(ctx *ReturnStatementTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expressionStatementTree.
	VisitExpressionStatementTree(ctx *ExpressionStatementTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expressionTree.
	VisitExpressionTree(ctx *ExpressionTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#lessComparison.
	VisitLessComparison(ctx *LessComparisonContext) interface{}

	// Visit a parse tree produced by MonkeyParser#greaterComparison.
	VisitGreaterComparison(ctx *GreaterComparisonContext) interface{}

	// Visit a parse tree produced by MonkeyParser#lessOrEqualsComparison.
	VisitLessOrEqualsComparison(ctx *LessOrEqualsComparisonContext) interface{}

	// Visit a parse tree produced by MonkeyParser#greaterOrEqualsComparison.
	VisitGreaterOrEqualsComparison(ctx *GreaterOrEqualsComparisonContext) interface{}

	// Visit a parse tree produced by MonkeyParser#equalsComparison.
	VisitEqualsComparison(ctx *EqualsComparisonContext) interface{}

	// Visit a parse tree produced by MonkeyParser#additionTree.
	VisitAdditionTree(ctx *AdditionTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#plusOperator.
	VisitPlusOperator(ctx *PlusOperatorContext) interface{}

	// Visit a parse tree produced by MonkeyParser#minusOperator.
	VisitMinusOperator(ctx *MinusOperatorContext) interface{}

	// Visit a parse tree produced by MonkeyParser#multiplicationTree.
	VisitMultiplicationTree(ctx *MultiplicationTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#multiplicationOperator.
	VisitMultiplicationOperator(ctx *MultiplicationOperatorContext) interface{}

	// Visit a parse tree produced by MonkeyParser#divisionOperator.
	VisitDivisionOperator(ctx *DivisionOperatorContext) interface{}

	// Visit a parse tree produced by MonkeyParser#elementTree.
	VisitElementTree(ctx *ElementTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#elementAccessTree.
	VisitElementAccessTree(ctx *ElementAccessTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#functionCallTree.
	VisitFunctionCallTree(ctx *FunctionCallTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by MonkeyParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by MonkeyParser#identifierTree.
	VisitIdentifierTree(ctx *IdentifierTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#true.
	VisitTrue(ctx *TrueContext) interface{}

	// Visit a parse tree produced by MonkeyParser#false.
	VisitFalse(ctx *FalseContext) interface{}

	// Visit a parse tree produced by MonkeyParser#groupedExpressionTree.
	VisitGroupedExpressionTree(ctx *GroupedExpressionTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#arrayTree.
	VisitArrayTree(ctx *ArrayTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#arrayFunctionTree.
	VisitArrayFunctionTree(ctx *ArrayFunctionTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#functionTree.
	VisitFunctionTree(ctx *FunctionTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#hashObjectTree.
	VisitHashObjectTree(ctx *HashObjectTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#printTree.
	VisitPrintTree(ctx *PrintTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#conditionalTree.
	VisitConditionalTree(ctx *ConditionalTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#arrayLen.
	VisitArrayLen(ctx *ArrayLenContext) interface{}

	// Visit a parse tree produced by MonkeyParser#arrayFirst.
	VisitArrayFirst(ctx *ArrayFirstContext) interface{}

	// Visit a parse tree produced by MonkeyParser#arrayLast.
	VisitArrayLast(ctx *ArrayLastContext) interface{}

	// Visit a parse tree produced by MonkeyParser#arrayRest.
	VisitArrayRest(ctx *ArrayRestContext) interface{}

	// Visit a parse tree produced by MonkeyParser#arrayPush.
	VisitArrayPush(ctx *ArrayPushContext) interface{}

	// Visit a parse tree produced by MonkeyParser#functionParametersTree.
	VisitFunctionParametersTree(ctx *FunctionParametersTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#hashPairTree.
	VisitHashPairTree(ctx *HashPairTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expressionListTree.
	VisitExpressionListTree(ctx *ExpressionListTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#blockTree.
	VisitBlockTree(ctx *BlockTreeContext) interface{}

	// Visit a parse tree produced by MonkeyParser#identifierNode.
	VisitIdentifierNode(ctx *IdentifierNodeContext) interface{}
}
