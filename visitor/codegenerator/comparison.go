package codegenerator

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
	v.addInstruction("COMPARE_OP", "<")

	return nil
}

func (v *visitor) VisitGreaterComparison(ctx *parser.GreaterComparisonContext) interface{} {
	v.addInstruction("COMPARE_OP", ">")

	return nil
}

func (v *visitor) VisitLessOrEqualsComparison(ctx *parser.LessOrEqualsComparisonContext) interface{} {
	v.addInstruction("COMPARE_OP", "<=")

	return nil
}

func (v *visitor) VisitGreaterOrEqualsComparison(ctx *parser.GreaterOrEqualsComparisonContext) interface{} {
	v.addInstruction("COMPARE_OP", ">=")

	return nil
}

func (v *visitor) VisitEqualsComparison(ctx *parser.EqualsComparisonContext) interface{} {
	v.addInstruction("COMPARE_OP", "==")

	return nil
}
