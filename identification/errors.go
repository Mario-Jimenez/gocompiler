package identification

type errorsHandler struct {
	errors []string
	lines  []int
}

func NewErrorsHandler() *errorsHandler {
	return &errorsHandler{
		errors: []string{},
		lines:  []int{},
	}
}

func (e *errorsHandler) Add(newError string, newLine int) {
	e.errors = append(e.errors, newError)
	e.lines = append(e.lines, newLine)
}

// Errors
func (e *errorsHandler) Errors() []string {
	return e.errors
}

// Lines
func (e *errorsHandler) Lines() []int {
	return e.lines
}
