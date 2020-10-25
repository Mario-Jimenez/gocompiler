package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/identification"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// contextualVisitor implementation of MonkeyParserVisitor interface
// methods are located in smaller files with rules names
type contextualVisitor struct {
	identificationTable *identification.Table
	functionTable       *identification.TableFunction

	//Declaracion de identificadores
	isFunction      bool
	parametersCount int

	//Uso de indentificadores
	isIdentifier    bool
	identifierToken antlr.Token
	isCall          bool
	callParameters  int
}

// NewContextualVisitor instance
func NewContextualVisitor(table *identification.Table, tableFunction *identification.TableFunction) *contextualVisitor {
	return &contextualVisitor{
		identificationTable: table,
		functionTable:       tableFunction,
	}
}
