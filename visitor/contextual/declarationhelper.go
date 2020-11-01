package contextual

type declaration struct {
	function   bool
	parameters int
}

type declarationHelper struct {
	levels       int
	declarations []declaration
}

func newDeclarationHelper() *declarationHelper {
	return &declarationHelper{
		levels:       -1,
		declarations: []declaration{},
	}
}

func (h *declarationHelper) newDeclaration() {
	h.levels++
	h.declarations = append(h.declarations, declaration{})
}

func (h *declarationHelper) markFunction() {
	h.declarations[h.levels].function = true
}

func (h *declarationHelper) isFunction() bool {
	return h.declarations[h.levels].function
}

func (h *declarationHelper) setParameters(parameters int) {
	h.declarations[h.levels].parameters = parameters
}

func (h *declarationHelper) getParameters() int {
	return h.declarations[h.levels].parameters
}

func (h *declarationHelper) closeDeclaration() {
	h.declarations = h.declarations[:h.levels]
	h.levels--
}
