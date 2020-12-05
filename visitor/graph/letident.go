package graph

import (
	"strings"

	"github.com/Mario-Jimenez/gocompiler/parser"
)

func (v *visitor) VisitLetIdentifier(ctx *parser.LetIdentifierContext) interface{} {
	hasError := false
	children := []map[string]interface{}{}

	if ctx.IDENTIFIER() != nil {
		if !strings.Contains(ctx.IDENTIFIER().GetText(), "<missing") {
			children = append(children, childNode(ctx.IDENTIFIER().GetText(), terminalColor))
		} else {
			hasError = true
			children = append(children, childNode("IDENTIFIER", errorColor))
		}
	} else {
		hasError = true
		children = append(children, childNode("IDENTIFIER", errorColor))
	}

	return nodeResponse("LetIdentifier", hasError, children)
}
