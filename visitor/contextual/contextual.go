package contextual

import "github.com/Mario-Jimenez/gocompiler/identification"

// visitor implementation of MonkeyParserVisitor interface
// contextual analysis
// methods are located in smaller files with rules names
type visitor struct {
	generalTable  *identification.GeneralTable
	functionTable *identification.FunctionTable

	declaration *declarationHelper
	identifier  *identifierHelper
}

// NewVisitor instance
func NewVisitor(generalTable *identification.GeneralTable,
	functionTable *identification.FunctionTable) *visitor {
	return &visitor{
		generalTable:  generalTable,
		functionTable: functionTable,
		declaration:   newDeclarationHelper(),
		identifier:    newIdentifierHelper(),
	}
}
