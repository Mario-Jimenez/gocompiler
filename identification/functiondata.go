package identification

// handles information for a function declaration
type FunctionData struct {
	parameters int
	hasReturn  bool
}

// constructor
func NewFunctionData(parameters int) *FunctionData {
	return &FunctionData{
		parameters: parameters,
	}
}

// getter
func (f *FunctionData) GetParameters() int {
	return f.parameters
}

// setter
func (f *FunctionData) SetReturn() {
	f.hasReturn = true
}

// getter
func (f *FunctionData) HasReturn() bool {
	return f.hasReturn
}
