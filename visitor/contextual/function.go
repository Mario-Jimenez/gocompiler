package contextual

import (
	"github.com/Mario-Jimenez/gocompiler/parser"
)

//function locals[declaration: *FunctionTreeContext]: FUNC # functionNode;

func (v *contextualVisitor) VisitFunctionNode(ctx *parser.FunctionNodeContext) interface{} {
	/*attr := v.identificationTable.Retrieve(ctx.FUNC().GetSymbol())
	if attr != nil {
		ctx.SetDeclaration(attr.Declaration())
	}
	*/
	return nil
}
