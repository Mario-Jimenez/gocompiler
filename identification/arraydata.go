package identification

// array's function element information
type arrayElement struct {
	parameters int
	hasReturn  bool
}

/*
	handles information for an array declaration
	stores information for function elements only
*/
type ArrayData struct {
	elements map[int]*arrayElement
}

// constructor
func NewArrayData() *ArrayData {
	return &ArrayData{
		elements: map[int]*arrayElement{},
	}
}

// new array element
func (a *ArrayData) AddElement(index int, parameters int, hasReturn bool) {
	a.elements[index] = &arrayElement{
		parameters: parameters,
		hasReturn:  hasReturn,
	}
}

// search element in array
func (a *ArrayData) FindElement(index int) *arrayElement {
	if element, ok := a.elements[index]; ok {
		return element
	}

	return nil
}

// getter
func (a *arrayElement) GetParameters() int {
	return a.parameters
}

// getter
func (a *arrayElement) HasReturn() bool {
	return a.hasReturn
}
