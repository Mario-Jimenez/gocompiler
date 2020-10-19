package contextual

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

func (v *contextualVisitor) VisitArrayLen(ctx *parser.ArrayLenContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitArrayFirst(ctx *parser.ArrayFirstContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitArrayLast(ctx *parser.ArrayLastContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitArrayRest(ctx *parser.ArrayRestContext) interface{} {
	return nil
}

func (v *contextualVisitor) VisitArrayPush(ctx *parser.ArrayPushContext) interface{} {
	return nil
}
