package identification

type FunctionData struct {
	parameters int
}

func NewFunctionData(parameters int) *FunctionData {
	return &FunctionData{
		parameters: parameters,
	}
}

func (f *FunctionData) GetParameters() int {
	return f.parameters
}
