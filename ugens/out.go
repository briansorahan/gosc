package ugens

import (
	"fmt"
)

// Out write a signal to a bus
var Out = newUgen("Out", func(node *BaseNode, args ...interface{}) {
	nargs := len(args)
	// parse arguments
	if nargs < 2 {
		panic(fmt.Errorf("Out expects at least 2 arguments, but was given %d", nargs))
	}
	getInputs(node, args...)
})
