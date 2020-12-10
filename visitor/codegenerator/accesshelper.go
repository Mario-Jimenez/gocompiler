package codegenerator

/*
	stores array info to access an element
*/
type accessHelper struct {
	global          bool
	name            string
	index           int
	functionIndexes []int
	accessElement   bool
}

// constructor
func newAccessHelper(global bool, name string, functionIndexes []int) *accessHelper {
	return &accessHelper{
		global:          global,
		name:            name,
		index:           -1,
		functionIndexes: functionIndexes,
	}
}

func (a *accessHelper) isFunction() bool {
	for _, val := range a.functionIndexes {
		if val == a.index {
			return true
		}
	}

	return false
}
