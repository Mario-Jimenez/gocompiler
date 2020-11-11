package handler

import (
	"net/http"

	"github.com/Mario-Jimenez/gocompiler/errors"
	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
	"github.com/Mario-Jimenez/gocompiler/visitor/contextual"
	"github.com/Mario-Jimenez/gocompiler/visitor/graph"
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

	// tree graph visitor
	graphVisitor := graph.NewVisitor()
	treeGraph := graphVisitor.Visit(tree)

	if parserErrors.Errors() == nil {
		// errors handler for contextual visitor
		contextualErrors := identification.NewErrorsHandler()
		// identification table
		table := identification.NewTable(contextualErrors)
		// contextual analysis visitor
		contextualVisitor := contextual.NewVisitor(table)
		// start of contextual visitor
		contextualVisitor.Visit(tree)

		if contextualErrors.Errors() == nil {
			return []string{}, []int{}, treeGraph
		}

		return contextualErrors.Errors(), contextualErrors.Lines(), treeGraph
	}

	return parserErrors.Errors(), parserErrors.Lines(), treeGraph
}
