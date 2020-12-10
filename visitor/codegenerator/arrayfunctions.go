package codegenerator

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
	v.addInstruction("LOAD_GLOBAL", "len")
	v.addInstruction("CALL_FUNCTION", "1")

	return nil
}

func (v *visitor) VisitArrayFirst(ctx *parser.ArrayFirstContext) interface{} {
	v.addInstruction("LOAD_GLOBAL", "first")
	v.addInstruction("CALL_FUNCTION", "1")

	return nil
}

func (v *visitor) VisitArrayLast(ctx *parser.ArrayLastContext) interface{} {
	v.addInstruction("LOAD_GLOBAL", "last")
	v.addInstruction("CALL_FUNCTION", "1")

	return nil
}

func (v *visitor) VisitArrayRest(ctx *parser.ArrayRestContext) interface{} {
	v.addInstruction("LOAD_GLOBAL", "rest")
	v.addInstruction("CALL_FUNCTION", "1")

	return nil
}

func (v *visitor) VisitArrayPush(ctx *parser.ArrayPushContext) interface{} {
	v.addInstruction("LOAD_GLOBAL", "push")
	v.addInstruction("CALL_FUNCTION", "2")

	return nil
}
