package visitor

/*
	parser visitor generates data structure to work with react's react-d3-tree component
	in front end
	https://github.com/bkrem/react-d3-tree
*/

// nodes colors
const rootColor string = "#0091ea"
const successColor string = "#00c853"
const errorColor string = "#d50000"
const terminalColor string = "#ffd600"

// parentNode data
func parentNode(name, color string, children []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"name": name,
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": color,
			},
		},
		"children": children,
	}
}

// childNode data
func childNode(name, color string) map[string]interface{} {
	return map[string]interface{}{
		"name": name,
		"nodeSvgShape": map[string]interface{}{
			"shape": "circle",
			"shapeProps": map[string]interface{}{
				"r":    10,
				"fill": color,
			},
		},
	}
}

// nodeResponse handles node's success or error response
func nodeResponse(name string, hasError bool, children []map[string]interface{}) *visitResponse {
	if hasError {
		return &visitResponse{
			node:   parentNode(name, errorColor, children),
			failed: true,
		}
	}

	return &visitResponse{
		node:   parentNode(name, successColor, children),
		failed: false,
	}
}
