package codegenerator

/*
	stores function call info
*/
type callHelper struct {
	global     bool
	name       string
	parameters int
	hasReturn  bool
}

// constructor
func newCallHelper(global bool, name string, parameters int, hasReturn bool) *callHelper {
	return &callHelper{
		global:     global,
		name:       name,
		parameters: parameters,
		hasReturn:  hasReturn,
	}
}
