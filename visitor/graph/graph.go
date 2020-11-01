package graph

// visitor implementation of MonkeyParserVisitor interface
// it generates a tree structure to be displayed as a tree graph
// methods are located in smaller files with rules names
type visitor struct{}

// NewVisitor instance
func NewVisitor() *visitor {
	return new(visitor)
}
