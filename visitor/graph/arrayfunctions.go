package graph

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

func (v *visitor) VisitArrayLen(ctx *parser.ArrayLenContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.LEN().GetText(), terminalColor),
		failed: false,
	}
}

func (v *visitor) VisitArrayFirst(ctx *parser.ArrayFirstContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.FIRST().GetText(), terminalColor),
		failed: false,
	}
}

func (v *visitor) VisitArrayLast(ctx *parser.ArrayLastContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.LAST().GetText(), terminalColor),
		failed: false,
	}
}

func (v *visitor) VisitArrayRest(ctx *parser.ArrayRestContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.REST().GetText(), terminalColor),
		failed: false,
	}
}

func (v *visitor) VisitArrayPush(ctx *parser.ArrayPushContext) interface{} {
	return &visitResponse{
		node:   childNode(ctx.PUSH().GetText(), terminalColor),
		failed: false,
	}
}
