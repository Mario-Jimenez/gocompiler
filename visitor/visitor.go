package visitor

import "github.com/antlr/antlr4/runtime/Go/antlr"

/*
	implementation of antlr's ParseTreeVisitor interface
	not implemented in Go by default
	was copied from the Java implementation of the same interface
	https://github.com/antlr/antlr4/blob/master/runtime/Java/src/org/antlr/v4/runtime/tree/AbstractParseTreeVisitor.java
*/

func (v *monkeyVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *monkeyVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var result interface{}
	for _, child := range node.GetChildren() {
		result = child.(antlr.ParseTree).Accept(v)
	}
	return result
}

func (v *monkeyVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}
func (v *monkeyVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}
