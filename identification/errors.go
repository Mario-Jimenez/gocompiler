package identification

// store errors
type errorsHandler struct {
	errors []string
	lines  []int
}

// constructor
func NewErrorsHandler() *errorsHandler {
	return &errorsHandler{
		errors: []string{},
		lines:  []int{},
	}
}

// store new error
func (e *errorsHandler) Add(newError string, newLine int) {
	e.errors = append(e.errors, newError)
	e.lines = append(e.lines, newLine)
}

// getter
func (e *errorsHandler) Errors() []string {
	return e.errors
}

// getter
func (e *errorsHandler) Lines() []int {
	return e.lines
}
