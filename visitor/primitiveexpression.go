package visitor

import (
	"strings"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

/*
	primitiveExpression:
		INTEGER														# integer
		| STRING													# string
		| identifier												# identifierTree
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

func (v *monkeyVisitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.INTEGER().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitString(ctx *parser.StringContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.STRING().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitIdentifierTree(ctx *parser.IdentifierTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	node := v.Visit(ctx.Identifier())
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
	}

	return nodeResponse("IdentifierTree", hasError, children)
}

func (v *monkeyVisitor) VisitTrue(ctx *parser.TrueContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.TRUE().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitFalse(ctx *parser.FalseContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.FALSE().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitGroupedExpressionTree(ctx *parser.GroupedExpressionTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	children = append(children, childNode(ctx.L_PAREN().GetText(), terminalColor))

	node := v.Visit(ctx.Expression())
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
		children = append(children, childNode("ExpressionTree", errorColor))
	}

	if ctx.R_PAREN() != nil {
		if !strings.Contains(ctx.R_PAREN().GetText(), "<missing") {
			children = append(children, childNode(ctx.R_PAREN().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("R_PAREN", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("R_PAREN", errorColor))
	}

	return nodeResponse("GroupedExpressionTree", hasError, children)
}

func (v *monkeyVisitor) VisitArrayTree(ctx *parser.ArrayTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	children = append(children, childNode(ctx.L_BRACKET().GetText(), terminalColor))

	if ctx.ExpressionList() != nil {
		node := v.Visit(ctx.ExpressionList())
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("ExpressionListTree", errorColor))
		}
	}

	if ctx.R_BRACKET() != nil {
		if !strings.Contains(ctx.R_BRACKET().GetText(), "<missing") {
			children = append(children, childNode(ctx.R_BRACKET().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("R_BRACKET", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("R_BRACKET", errorColor))
	}

	return nodeResponse("ArrayTree", hasError, children)
}

func (v *monkeyVisitor) VisitArrayFunctionTree(ctx *parser.ArrayFunctionTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	node := v.Visit(ctx.ArrayFunctions())
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
	}

	if ctx.L_PAREN() != nil {
		if !strings.Contains(ctx.L_PAREN().GetText(), "<missing") {
			children = append(children, childNode(ctx.L_PAREN().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("L_PAREN", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("L_PAREN", errorColor))
	}

	if ctx.ExpressionList() != nil {
		node = v.Visit(ctx.ExpressionList())
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("ExpressionListTree", errorColor))
		}
	}

	if ctx.R_PAREN() != nil {
		if !strings.Contains(ctx.R_PAREN().GetText(), "<missing") {
			children = append(children, childNode(ctx.R_PAREN().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("R_PAREN", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("R_PAREN", errorColor))
	}

	return nodeResponse("ArrayFunctionTree", hasError, children)
}

func (v *monkeyVisitor) VisitFunctionATree(ctx *parser.FunctionATreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	//TODO: REVISAR
	v.Visit(ctx.Function())

	if ctx.L_PAREN() != nil {
		if !strings.Contains(ctx.L_PAREN().GetText(), "<missing") {
			children = append(children, childNode(ctx.L_PAREN().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("L_PAREN", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("L_PAREN", errorColor))
	}

	if ctx.FunctionParameters() != nil {
		node := v.Visit(ctx.FunctionParameters())
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("FunctionParametersTree", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("FunctionParametersTree", errorColor))
	}

	if ctx.R_PAREN() != nil {
		if !strings.Contains(ctx.R_PAREN().GetText(), "<missing") {
			children = append(children, childNode(ctx.R_PAREN().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("R_PAREN", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("R_PAREN", errorColor))
	}

	if ctx.BlockStatement() != nil {
		node := v.Visit(ctx.BlockStatement())
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("BlockTree", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("BlockTree", errorColor))
	}

	return nodeResponse("FunctionTree", hasError, children)
}

func (v *monkeyVisitor) VisitHashObjectTree(ctx *parser.HashObjectTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	children = append(children, childNode(ctx.L_CURLY().GetText(), terminalColor))

	node := v.Visit(ctx.HashContent(0))
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
		children = append(children, childNode("HashPairTree", errorColor))
	}

	totalBranches := len(ctx.AllHashContent())
	index := 1
	for index < totalBranches {
		children = append(children, childNode(ctx.COMMA(index-1).GetText(), terminalColor))

		node = v.Visit(ctx.HashContent(index))
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("HashPairTree", errorColor))
		}

		index++
	}

	if ctx.R_CURLY() != nil {
		if !strings.Contains(ctx.R_CURLY().GetText(), "<missing") {
			children = append(children, childNode(ctx.R_CURLY().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("R_CURLY", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("R_CURLY", errorColor))
	}

	return nodeResponse("HashObjectTree", hasError, children)
}

func (v *monkeyVisitor) VisitPrintTree(ctx *parser.PrintTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	children = append(children, childNode(ctx.PUTS().GetText(), terminalColor))

	if ctx.L_PAREN() != nil {
		if !strings.Contains(ctx.L_PAREN().GetText(), "<missing") {
			children = append(children, childNode(ctx.L_PAREN().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("L_PAREN", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("L_PAREN", errorColor))
	}

	if ctx.Expression() != nil {
		node := v.Visit(ctx.Expression())
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
		}
	} else {
		hasError = true
		children = append(children, childNode("ExpressionTree", errorColor))
	}

	if ctx.R_PAREN() != nil {
		if !strings.Contains(ctx.R_PAREN().GetText(), "<missing") {
			children = append(children, childNode(ctx.R_PAREN().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("R_PAREN", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("R_PAREN", errorColor))
	}

	return nodeResponse("PrintTree", hasError, children)
}

func (v *monkeyVisitor) VisitConditionalTree(ctx *parser.ConditionalTreeContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	children = append(children, childNode(ctx.IF().GetText(), terminalColor))

	node := v.Visit(ctx.Expression())
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
		children = append(children, childNode("ExpressionTree", errorColor))
	}

	node = v.Visit(ctx.BlockStatement(0))
	if n, ok := node.(*visitResponse); ok {
		if n.hasFailed() {
			hasError = true
		}
		children = append(children, n.info())
	} else {
		hasError = true
		children = append(children, childNode("BlockTree", errorColor))
	}

	if ctx.ELSE() != nil {
		children = append(children, childNode(ctx.ELSE().GetText(), terminalColor))

		node = v.Visit(ctx.BlockStatement(1))
		if n, ok := node.(*visitResponse); ok {
			if n.hasFailed() {
				hasError = true
			}
			children = append(children, n.info())
		} else {
			hasError = true
			children = append(children, childNode("BlockTree", errorColor))
		}
	}

	return nodeResponse("ConditionalTree", hasError, children)
}
