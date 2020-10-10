package visitor

import (
	"strings"

	"github.com/Mario-Jimenez/gocompiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const rootColor string = "#0091ea"
const successColor string = "#00c853"
const errorColor string = "#d50000"
const terminalColor string = "#ffd600"

type MonkeyVisitor struct{}

func NewMonkeyVisitor() *MonkeyVisitor {
	return new(MonkeyVisitor)
}

func (v *MonkeyVisitor) parserError(name string, children []map[string]interface{}) interface{} {
	return map[string]interface{}{
		"name": name,
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": errorColor,
			},
		},
		"children": children,
	}
}

// Start ParseTreeVisitor implementation

func (v *MonkeyVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *MonkeyVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var result interface{}
	for _, child := range node.GetChildren() {
		result = child.(antlr.ParseTree).Accept(v)
	}
	return result
}

func (v *MonkeyVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}
func (v *MonkeyVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

// End ParseTreeVisitor implementation

func (v *MonkeyVisitor) VisitProgramTree(ctx *parser.ProgramTreeContext) interface{} {
	children := []map[string]interface{}{}

	for _, branch := range ctx.AllStatement() {
		node := v.Visit(branch)
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}
	}

	children = append(children, map[string]interface{}{
		"name": ctx.EOF().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	root := []map[string]interface{}{
		{
			"name": "ProgramTree",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": rootColor,
				},
			},
			"children": children,
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitLetStatementTree(ctx *parser.LetStatementTreeContext) interface{} {
	children := []map[string]interface{}{}

	children = append(children, map[string]interface{}{
		"name": ctx.LET().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	if ctx.IDENTIFIER() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.IDENTIFIER().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "IDENTIFIER",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("LetStatementTree", children)
	}

	if ctx.ASSIGN() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.ASSIGN().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "ASSIGN",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("LetStatementTree", children)
	}

	node := v.Visit(ctx.Expression())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.SEMI() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.SEMI().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	}

	root := map[string]interface{}{
		"name": "LetStatementTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitReturnStatementTree(ctx *parser.ReturnStatementTreeContext) interface{} {
	children := []map[string]interface{}{}

	children = append(children, map[string]interface{}{
		"name": ctx.RETURN().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	node := v.Visit(ctx.Expression())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.SEMI() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.SEMI().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	}

	root := map[string]interface{}{
		"name": "ReturnStatementTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitExpressionStatementTree(ctx *parser.ExpressionStatementTreeContext) interface{} {
	children := []map[string]interface{}{}

	node := v.Visit(ctx.Expression())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.SEMI() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.SEMI().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	}

	root := map[string]interface{}{
		"name": "ExpressionStatementTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitExpressionTree(ctx *parser.ExpressionTreeContext) interface{} {
	children := []map[string]interface{}{}

	node := v.Visit(ctx.AdditionExpression(0))
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	totalBranches := len(ctx.AllAdditionExpression())
	index := 1
	for index < totalBranches {
		node = v.Visit(ctx.Comparison(index - 1))
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}

		node = v.Visit(ctx.AdditionExpression(index))
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}

		index++
	}

	root := map[string]interface{}{
		"name": "ExpressionTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitLessComparison(ctx *parser.LessComparisonContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.LESS().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitGreaterComparison(ctx *parser.GreaterComparisonContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.GREATER().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitLessOrEqualsComparison(ctx *parser.LessOrEqualsComparisonContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.LESS_OR_EQUALS().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitGreaterOrEqualsComparison(ctx *parser.GreaterOrEqualsComparisonContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.GREATER_OR_EQUALS().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitEqualsComparison(ctx *parser.EqualsComparisonContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.EQUALS().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitAdditionTree(ctx *parser.AdditionTreeContext) interface{} {
	children := []map[string]interface{}{}

	node := v.Visit(ctx.MultiplicationExpression(0))
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	totalBranches := len(ctx.AllMultiplicationExpression())
	index := 1
	for index < totalBranches {
		node = v.Visit(ctx.AdditionFactor(index - 1))
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}

		node = v.Visit(ctx.MultiplicationExpression(index))
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}

		index++
	}

	root := map[string]interface{}{
		"name": "AdditionTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitPlusOperator(ctx *parser.PlusOperatorContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.PLUS().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitMinusOperator(ctx *parser.MinusOperatorContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.MINUS().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitMultiplicationTree(ctx *parser.MultiplicationTreeContext) interface{} {
	children := []map[string]interface{}{}

	node := v.Visit(ctx.ElementExpression(0))
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	totalBranches := len(ctx.AllElementExpression())
	index := 1
	for index < totalBranches {
		node = v.Visit(ctx.MultiplicationFactor(index - 1))
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}

		node = v.Visit(ctx.ElementExpression(index))
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}

		index++
	}

	root := map[string]interface{}{
		"name": "MultiplicationTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitMultiplicationOperator(ctx *parser.MultiplicationOperatorContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.MULT().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitDivisionOperator(ctx *parser.DivisionOperatorContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.DIV().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitElementTree(ctx *parser.ElementTreeContext) interface{} {
	children := []map[string]interface{}{}

	node := v.Visit(ctx.PrimitiveExpression())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.ElementAccess() != nil {
		node := v.Visit(ctx.ElementAccess())
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}
	} else if ctx.CallExpression() != nil {
		node := v.Visit(ctx.CallExpression())
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}
	}

	root := map[string]interface{}{
		"name": "ElementTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitElementAccessTree(ctx *parser.ElementAccessTreeContext) interface{} {
	children := []map[string]interface{}{}

	children = append(children, map[string]interface{}{
		"name": ctx.L_BRACKET().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	node := v.Visit(ctx.Expression())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.R_BRACKET() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.R_BRACKET().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "R_BRACKET",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("ElementAccessTree", children)
	}

	root := map[string]interface{}{
		"name": "ElementAccessTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitFunctionCallTree(ctx *parser.FunctionCallTreeContext) interface{} {
	children := []map[string]interface{}{}

	children = append(children, map[string]interface{}{
		"name": ctx.L_PAREN().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	node := v.Visit(ctx.ExpressionList())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.R_PAREN() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.R_PAREN().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "R_PAREN",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("FunctionCallTree", children)
	}

	root := map[string]interface{}{
		"name": "FunctionCallTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.INTEGER().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitString(ctx *parser.StringContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.STRING().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.IDENTIFIER().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitTrue(ctx *parser.TrueContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.TRUE().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitFalse(ctx *parser.FalseContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.FALSE().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitGroupedExpressionTree(ctx *parser.GroupedExpressionTreeContext) interface{} {
	children := []map[string]interface{}{}

	children = append(children, map[string]interface{}{
		"name": ctx.L_PAREN().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	node := v.Visit(ctx.Expression())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.R_PAREN() != nil {
		if !strings.Contains(ctx.R_PAREN().GetText(), "missing") {
			children = append(children, map[string]interface{}{
				"name": ctx.R_PAREN().GetText(),
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": terminalColor,
					},
				},
			})
		} else {
			children = append(children, map[string]interface{}{
				"name": "R_PAREN",
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": errorColor,
					},
				},
			})
			return v.parserError("GroupedExpressionTree", children)
		}
	} else {
		children = append(children, map[string]interface{}{
			"name": "R_PAREN",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("GroupedExpressionTree", children)
	}

	root := map[string]interface{}{
		"name": "GroupedExpressionTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitArrayTree(ctx *parser.ArrayTreeContext) interface{} {
	children := []map[string]interface{}{}

	children = append(children, map[string]interface{}{
		"name": ctx.L_BRACKET().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	node := v.Visit(ctx.ExpressionList())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.R_BRACKET() != nil {
		if !strings.Contains(ctx.R_BRACKET().GetText(), "missing") {
			children = append(children, map[string]interface{}{
				"name": ctx.R_BRACKET().GetText(),
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": terminalColor,
					},
				},
			})
		} else {
			children = append(children, map[string]interface{}{
				"name": "R_BRACKET",
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": errorColor,
					},
				},
			})
			return v.parserError("ArrayTree", children)
		}
	} else {
		children = append(children, map[string]interface{}{
			"name": "R_BRACKET",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("ArrayTree", children)
	}

	root := map[string]interface{}{
		"name": "ArrayTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitArrayFunctionTree(ctx *parser.ArrayFunctionTreeContext) interface{} {
	children := []map[string]interface{}{}

	node := v.Visit(ctx.ArrayFunctions())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.L_PAREN() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.L_PAREN().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "L_PAREN",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("ArrayFunctionTree", children)
	}

	node = v.Visit(ctx.ExpressionList())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.R_PAREN() != nil {
		if !strings.Contains(ctx.R_PAREN().GetText(), "missing") {
			children = append(children, map[string]interface{}{
				"name": ctx.R_PAREN().GetText(),
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": terminalColor,
					},
				},
			})
		} else {
			children = append(children, map[string]interface{}{
				"name": "R_PAREN",
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": errorColor,
					},
				},
			})
			return v.parserError("ArrayFunctionTree", children)
		}
	} else {
		children = append(children, map[string]interface{}{
			"name": "R_PAREN",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("ArrayFunctionTree", children)
	}

	root := map[string]interface{}{
		"name": "ArrayFunctionTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitFunctionTree(ctx *parser.FunctionTreeContext) interface{} {
	children := []map[string]interface{}{}

	children = append(children, map[string]interface{}{
		"name": ctx.FUNC().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	if ctx.L_PAREN() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.L_PAREN().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "L_PAREN",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("FunctionTree", children)
	}

	node := v.Visit(ctx.FunctionParameters())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.R_PAREN() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.R_PAREN().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "R_PAREN",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("FunctionTree", children)
	}

	node = v.Visit(ctx.BlockStatement())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	root := map[string]interface{}{
		"name": "FunctionTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitHashObjectTree(ctx *parser.HashObjectTreeContext) interface{} {
	children := []map[string]interface{}{}

	children = append(children, map[string]interface{}{
		"name": ctx.L_CURLY().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	node := v.Visit(ctx.HashContent(0))
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	totalBranches := len(ctx.AllHashContent())
	index := 1
	for index < totalBranches {
		children = append(children, map[string]interface{}{
			"name": ctx.COMMA(index - 1).GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})

		node = v.Visit(ctx.HashContent(index))
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}

		index++
	}

	if ctx.R_CURLY() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.R_CURLY().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "R_CURLY",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("HashObjectTree", children)
	}

	root := map[string]interface{}{
		"name": "HashObjectTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitPrintTree(ctx *parser.PrintTreeContext) interface{} {
	children := []map[string]interface{}{}

	children = append(children, map[string]interface{}{
		"name": ctx.PUTS().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	if ctx.L_PAREN() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.L_PAREN().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "L_PAREN",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("PrintTree", children)
	}

	node := v.Visit(ctx.Expression())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.R_PAREN() != nil {
		if !strings.Contains(ctx.R_PAREN().GetText(), "missing") {
			children = append(children, map[string]interface{}{
				"name": ctx.R_PAREN().GetText(),
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": terminalColor,
					},
				},
			})
		} else {
			children = append(children, map[string]interface{}{
				"name": "R_PAREN",
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": errorColor,
					},
				},
			})
			return v.parserError("PrintTree", children)
		}
	} else {
		children = append(children, map[string]interface{}{
			"name": "R_PAREN",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("PrintTree", children)
	}

	root := map[string]interface{}{
		"name": "PrintTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitConditionalTree(ctx *parser.ConditionalTreeContext) interface{} {
	children := []map[string]interface{}{}

	children = append(children, map[string]interface{}{
		"name": ctx.IF().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	})

	node := v.Visit(ctx.Expression())
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	node = v.Visit(ctx.BlockStatement(0))
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.ELSE() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.ELSE().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})

		node = v.Visit(ctx.BlockStatement(1))
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}
	}

	root := map[string]interface{}{
		"name": "ConditionalTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitArrayLen(ctx *parser.ArrayLenContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.LEN().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitArrayFirst(ctx *parser.ArrayFirstContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.FIRST().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitArrayLast(ctx *parser.ArrayLastContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.LAST().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitArrayRest(ctx *parser.ArrayRestContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.REST().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitArrayPush(ctx *parser.ArrayPushContext) interface{} {
	root := map[string]interface{}{
		"name": ctx.PUSH().GetText(),
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": terminalColor,
			},
		},
	}

	return root
}

func (v *MonkeyVisitor) VisitFunctionParametersTree(ctx *parser.FunctionParametersTreeContext) interface{} {
	children := []map[string]interface{}{}

	if ctx.IDENTIFIER(0) != nil {
		if !strings.Contains(ctx.IDENTIFIER(0).GetText(), "<missing") {
			children = append(children, map[string]interface{}{
				"name": ctx.IDENTIFIER(0).GetText(),
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": terminalColor,
					},
				},
			})
		} else {
			children = append(children, map[string]interface{}{
				"name": "IDENTIFIER",
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": errorColor,
					},
				},
			})
			return v.parserError("FunctionParametersTree", children)
		}
	} else {
		children = append(children, map[string]interface{}{
			"name": "IDENTIFIER",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("FunctionParametersTree", children)
	}

	totalBranches := len(ctx.AllIDENTIFIER())
	index := 1
	for index < totalBranches {
		children = append(children, map[string]interface{}{
			"name": ctx.COMMA(index - 1).GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})

		if ctx.IDENTIFIER(index) != nil {
			if !strings.Contains(ctx.IDENTIFIER(index).GetText(), "<missing") {
				children = append(children, map[string]interface{}{
					"name": ctx.IDENTIFIER(index).GetText(),
					"nodeSvgShape": map[string]interface{}{
						"shape": "circle",
						"shapeProps": map[string]interface{}{
							"r":    10,
							"fill": terminalColor,
						},
					},
				})
			} else {
				children = append(children, map[string]interface{}{
					"name": "IDENTIFIER",
					"nodeSvgShape": map[string]interface{}{
						"shape": "circle",
						"shapeProps": map[string]interface{}{
							"r":    10,
							"fill": errorColor,
						},
					},
				})
				return v.parserError("FunctionParametersTree", children)
			}
		} else {
			children = append(children, map[string]interface{}{
				"name": "IDENTIFIER",
				"nodeSvgShape": map[string]interface{}{
					"shape": "circle",
					"shapeProps": map[string]interface{}{
						"r":    10,
						"fill": errorColor,
					},
				},
			})
			return v.parserError("FunctionParametersTree", children)
		}

		index++
	}

	root := map[string]interface{}{
		"name": "FunctionParametersTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitHashPairTree(ctx *parser.HashPairTreeContext) interface{} {
	children := []map[string]interface{}{}

	node := v.Visit(ctx.Expression(0))
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	if ctx.COLON() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.COLON().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "COLON",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("HashPairTree", children)
	}

	node = v.Visit(ctx.Expression(1))
	if n, ok := node.(map[string]interface{}); ok {
		children = append(children, n)
	}

	root := map[string]interface{}{
		"name": "HashPairTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitExpressionListTree(ctx *parser.ExpressionListTreeContext) interface{} {
	children := []map[string]interface{}{}

	if ctx.Expression(0) != nil {
		node := v.Visit(ctx.Expression(0))
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}
	} else {
		return nil
	}

	totalBranches := len(ctx.AllExpression())
	index := 1
	for index < totalBranches {
		children = append(children, map[string]interface{}{
			"name": ctx.COMMA(index - 1).GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})

		node := v.Visit(ctx.Expression(index))
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}

		index++
	}

	root := map[string]interface{}{
		"name": "ExpressionListTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}

func (v *MonkeyVisitor) VisitBlockTree(ctx *parser.BlockTreeContext) interface{} {
	children := []map[string]interface{}{}

	if ctx.L_CURLY() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.L_CURLY().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "L_CURLY",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("BlockTree", children)
	}

	for _, branch := range ctx.AllStatement() {
		node := v.Visit(branch)
		if n, ok := node.(map[string]interface{}); ok {
			children = append(children, n)
		}
	}

	if ctx.R_CURLY() != nil {
		children = append(children, map[string]interface{}{
			"name": ctx.R_CURLY().GetText(),
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": terminalColor,
				},
			},
		})
	} else {
		children = append(children, map[string]interface{}{
			"name": "R_CURLY",
			"nodeSvgShape": map[string]interface{}{
				"shape": "circle",
				"shapeProps": map[string]interface{}{
					"r":    10,
					"fill": errorColor,
				},
			},
		})
		return v.parserError("BlockTree", children)
	}

	root := map[string]interface{}{
		"name": "BlockTree",
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": successColor,
			},
		},
		"children": children,
	}

	return root
}
