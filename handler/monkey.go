package handler

import (
	"net/http"

	"github.com/Mario-Jimenez/gocompiler/contextual"
	"github.com/Mario-Jimenez/gocompiler/errors"
	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
	"github.com/Mario-Jimenez/gocompiler/visitor"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gin-gonic/gin"
)

// Monkey contains compiler endpoints functions
type Monkey struct{}

// Compile request
type Compile struct {
	Program string `json:"program"`
}

// Compile incoming program
func (*Monkey) Compile(c *gin.Context) {
	// parse incoming request
	var req Compile
	if err := c.ShouldBindJSON(&req); err != nil {
		// response with error, bad request, missing program parameter
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	// parsing
	parseErrors, parseLines, parseTree := parsing(req.Program)

	// response
	c.JSON(200, gin.H{
		"errors": parseErrors,
		"lines":  parseLines,
		"tree":   parseTree,
	})
}

// parsing program
func parsing(program string) ([]string, []int, interface{}) {
	// compiler input
	input := antlr.NewInputStream(program)

	parserErrors := errors.NewParserErrorListener()

	// compiler lexer
	lexer := parser.NewMonkeyLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(parserErrors)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// compiler parser
	parser := parser.NewMonkeyParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrors)

	// initial rule to start parsing process
	tree := parser.Program()

	// parser tree visitor
	visitor := visitor.NewMonkeyVisitor()
	parseTree := visitor.Visit(tree)

	if parserErrors.Errors() == nil {
		// contextual analysis visitor
		table := identification.NewTable()
		tableFunction := identification.NewTableFunction()
		ctxVisitor := contextual.NewContextualVisitor(table, tableFunction)
		ctxVisitor.Visit(tree)

		if table.Errors() == nil {
			return []string{}, []int{}, parseTree
		}

		return table.Errors(), table.Lines(), parseTree
	}

	return parserErrors.Errors(), parserErrors.Lines(), parseTree
}
