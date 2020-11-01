package graph

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
	return &visitResponse{
		node:   childNode(ctx.LESS().GetText(), terminalColor),
		failed: false,
	}
}

func (v *visitor) VisitGreaterComparison(ctx *parser.GreaterComparisonContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.GREATER().GetText(), terminalColor),
		failed: false,
	}
}

func (v *visitor) VisitLessOrEqualsComparison(ctx *parser.LessOrEqualsComparisonContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.LESS_OR_EQUALS().GetText(), terminalColor),
		failed: false,
	}
}

func (v *visitor) VisitGreaterOrEqualsComparison(ctx *parser.GreaterOrEqualsComparisonContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.GREATER_OR_EQUALS().GetText(), terminalColor),
		failed: false,
	}
}

func (v *visitor) VisitEqualsComparison(ctx *parser.EqualsComparisonContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.EQUALS().GetText(), terminalColor),
		failed: false,
	}
}
