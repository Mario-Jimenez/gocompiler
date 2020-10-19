package visitor

// monkeyVisitor implementation of MonkeyParserVisitor interface
// methods are located in smaller files with rules names
type monkeyVisitor struct{}

// NewMonkeyVisitor instance
func NewMonkeyVisitor() *monkeyVisitor {
	return new(monkeyVisitor)
}
