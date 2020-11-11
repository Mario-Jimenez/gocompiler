package identification

// array's function element information
type arrayElement struct {
	index      int
	parameters int
}

/*
	handles information for an array declaration
	stores information for function elements only
*/
type ArrayData struct {
	elements []*arrayElement
}

// constructor
func NewArrayData() *ArrayData {
	return &ArrayData{}
}

// new array element
func (a *ArrayData) AddElement(index int, parameters int) {
	a.elements = append(a.elements, &arrayElement{
		index:      index,
		parameters: parameters,
	})
}

// search element in array
func (a *ArrayData) FindElement(index int) *arrayElement {
	for _, e := range a.elements {
		if e.index == index {
			return e
		}
	}

	return nil
}

// getter
func (a *arrayElement) GetParameters() int {
	return a.parameters
}
