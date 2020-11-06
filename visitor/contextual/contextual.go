package contextual

import "github.com/Mario-Jimenez/gocompiler/identification"

// visitor implementation of MonkeyParserVisitor interface
// contextual analysis
// methods are located in smaller files with rules names
type visitor struct {
	table *identification.Table

	declaration *declarationHelper
	identifier  *identifierHelper
	hash        *hashHelper
}

// NewVisitor instance
func NewVisitor(table *identification.Table) *visitor {
	return &visitor{
		table:       table,
		declaration: newDeclarationHelper(),
		identifier:  newIdentifierHelper(),
		hash:        newHashHelper(),
	}
}
