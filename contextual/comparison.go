package contextual

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	comparison:
		LESS				# lessComparison
		| GREATER			# greaterComparison
		| LESS_OR_EQUALS	# lessOrEqualsComparison
		| GREATER_OR_EQUALS	# greaterOrEqualsComparison
		| EQUALS			# equalsComparison
		;
*/

func (v *contextualVisitor) VisitLessComparison(ctx *parser.LessComparisonContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitGreaterComparison(ctx *parser.GreaterComparisonContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitLessOrEqualsComparison(ctx *parser.LessOrEqualsComparisonContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitGreaterOrEqualsComparison(ctx *parser.GreaterOrEqualsComparisonContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitEqualsComparison(ctx *parser.EqualsComparisonContext) interface{} {
	return nil
}
