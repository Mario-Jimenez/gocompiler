package graph

import "github.com/Mario-Jimenez/gocompiler/parser"

// visitor implementation of MonkeyParserVisitor interface
// it generates a tree structure to be displayed as a tree graph
// methods are located in smaller files with rules names
type visitor struct{}

// NewVisitor instance
func NewVisitor() parser.MonkeyParserVisitor {
	return new(visitor)
}
