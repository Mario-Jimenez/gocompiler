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

func (v *visitor) VisitLessComparison(ctx *parser.LessComparisonContext) interface{} {
	return nil
}

func (v *visitor) VisitGreaterComparison(ctx *parser.GreaterComparisonContext) interface{} {
	return nil
}

func (v *visitor) VisitLessOrEqualsComparison(ctx *parser.LessOrEqualsComparisonContext) interface{} {
	return nil
}

func (v *visitor) VisitGreaterOrEqualsComparison(ctx *parser.GreaterOrEqualsComparisonContext) interface{} {
	return nil
}

func (v *visitor) VisitEqualsComparison(ctx *parser.EqualsComparisonContext) interface{} {
	return nil
}
