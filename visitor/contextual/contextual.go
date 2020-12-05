package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/Mario-Jimenez/gocompiler/parser"
)

// visitor implementation of MonkeyParserVisitor interface
// fot the contextual analysis
// methods are located in smaller files with rules names
type visitor struct {
	table *identification.Table

	declaration *declarationHelper
	identifier  *identifierHelper
	hash        *hashHelper
	array       *arrayHelper
}

// NewVisitor instance
func NewVisitor(table *identification.Table) parser.MonkeyParserVisitor {
	return &visitor{
		table:       table,
		declaration: newDeclarationHelper(),
		identifier:  newIdentifierHelper(),
		hash:        newHashHelper(),
		array:       newArrayHelper(),
	}
}
