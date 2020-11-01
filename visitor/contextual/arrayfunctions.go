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

func (v *visitor) VisitArrayLen(ctx *parser.ArrayLenContext) interface{} {
	return nil
}

func (v *visitor) VisitArrayFirst(ctx *parser.ArrayFirstContext) interface{} {
	return nil
}

func (v *visitor) VisitArrayLast(ctx *parser.ArrayLastContext) interface{} {
	return nil
}

func (v *visitor) VisitArrayRest(ctx *parser.ArrayRestContext) interface{} {
	return nil
}

func (v *visitor) VisitArrayPush(ctx *parser.ArrayPushContext) interface{} {
	return nil
}
