package identification

// handles information for a function declaration
type FunctionData struct {
	parameters int
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
