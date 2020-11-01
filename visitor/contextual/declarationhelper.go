package contextual

type declarationHelper struct {
	function   bool
	parameters int
}

func newDeclarationHelper() *declarationHelper {
	return &declarationHelper{}
}

func (h *declarationHelper) markFunction() {
	h.function = true
}

func (h *declarationHelper) isFunction() bool {
	return h.function
}

func (h *declarationHelper) setParameters(parameters int) {
	h.parameters = parameters
}

func (h *declarationHelper) getParameters() int {
	return h.parameters
}

func (h *declarationHelper) reset() {
	h.function = false
	h.parameters = 0
}
