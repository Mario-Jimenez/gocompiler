package visitor

import "github.com/Mario-Jimenez/gocompiler/parser"

/*
	arrayFunctions:
		LEN		# arrayLen
		| FIRST	# arrayFirst
		| LAST	# arrayLast
		| REST	# arrayRest
		| PUSH	# arrayPush
		;
*/

func (v *monkeyVisitor) VisitArrayLen(ctx *parser.ArrayLenContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.LEN().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitArrayFirst(ctx *parser.ArrayFirstContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.FIRST().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitArrayLast(ctx *parser.ArrayLastContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.LAST().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitArrayRest(ctx *parser.ArrayRestContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.REST().GetText(), terminalColor),
		failed: false,
	}
}

func (v *monkeyVisitor) VisitArrayPush(ctx *parser.ArrayPushContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.PUSH().GetText(), terminalColor),
		failed: false,
	}
}
