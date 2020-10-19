package contextual

import "github.com/Mario-Jimenez/gocompiler/identification"

// contextualVisitor implementation of MonkeyParserVisitor interface
// methods are located in smaller files with rules names
type contextualVisitor struct {
	identificationTable *identification.Table
}

// NewContextualVisitor instance
func NewContextualVisitor(table *identification.Table) *contextualVisitor {
	return &contextualVisitor{
		identificationTable: table,
	}
}
