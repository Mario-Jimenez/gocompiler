package codegenerator

/*
	stores array info
*/
type arrayHelper struct {
	global          bool
	name            string
	indexes         []int
	functionIndexes []int
}

// constructor
func newArrayHelper(global bool, name string) *arrayHelper {
	return &arrayHelper{
		global: global,
		name:   name,
	}
}

// stores instruction index
func (a *arrayHelper) addIndex(index int) {
	a.indexes = append(a.indexes, index)
}

// stores function instruction index
func (a *arrayHelper) addFunctionIndex(index int) {
	a.functionIndexes = append(a.functionIndexes, index)
}

// check if element in array is a function
func (a *arrayHelper) isFunction(index int) bool {
	for _, val := range a.functionIndexes {
		if val == index {
			return true
		}
	}

	return false
}
