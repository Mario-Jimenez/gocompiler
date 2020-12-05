package identification

// Declaration stored info
type Declaration struct {
	level      int
	expression ExpressionType
	data       interface{}
}

func (d Declaration) Level() interface{} {
	return d.level
}

func (d Declaration) Expression() interface{} {
	return d.expression
}

func (d Declaration) Data() interface{} {
	return d.data
}
