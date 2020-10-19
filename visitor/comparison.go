package visitor

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

func (v *monkeyVisitor) VisitLessComparison(ctx *parser.LessComparisonContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.LESS().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitGreaterComparison(ctx *parser.GreaterComparisonContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.GREATER().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitLessOrEqualsComparison(ctx *parser.LessOrEqualsComparisonContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.LESS_OR_EQUALS().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitGreaterOrEqualsComparison(ctx *parser.GreaterOrEqualsComparisonContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.GREATER_OR_EQUALS().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitEqualsComparison(ctx *parser.EqualsComparisonContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.EQUALS().GetText(), terminalColor),
		failed: false,
	}
}
