package contextual

/*
	store data for function elements in array
*/
type arrayElement struct {
	index      int
	parameters int
	hasReturn  bool
	active     bool
}

/*
	stores array element data, for function elements only
	can have nested arrays, but only stores data for the main array
*/
type arrayHelper struct {
	levels int
	arrays []arrayElement
}

// constructor
func newArrayHelper() *arrayHelper {
	return &arrayHelper{
		levels: -1,
		arrays: []arrayElement{},
	}
}

// creates new array helper (can be nested)
func (h *arrayHelper) newArray() {
	h.levels++
	h.arrays = append(h.arrays, arrayElement{})
}

// setter
func (h *arrayHelper) setParameters(parameters int) {
	// validates that we're working with an array expression
	if h.levels > -1 {
		h.arrays[h.levels].parameters = parameters
	}
}

// getter
func (h *arrayHelper) getParameters() int {
	return h.arrays[h.levels].parameters
}

// setter
func (h *arrayHelper) setReturn() {
	// validates that we're working with an array expression
	if h.levels > -1 {
		h.arrays[h.levels].hasReturn = true
	}
}

// getter
func (h *arrayHelper) getReturn() bool {
	return h.arrays[h.levels].hasReturn
}

// allow storage of function element
func (h *arrayHelper) activate() {
	// validates that we're working with an array expression
	if h.levels > -1 {
		h.arrays[h.levels].active = true
	}
}

// checks if we're working with a function element
func (h *arrayHelper) isActive() bool {
	return h.arrays[h.levels].active
}

// checks if it is a nested array
func (h *arrayHelper) isNestedArray() bool {
	if h.levels == 0 {
		return false
	}

	return true
}

// clears element data
func (h *arrayHelper) reset() {
	h.arrays[h.levels].active = false
}

// close array helper
func (h *arrayHelper) closeArray() {
	h.arrays = h.arrays[:h.levels]
	h.levels--
}
