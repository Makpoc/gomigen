package model

import (
	"fmt"
	"strings"
)

// Arguments is the list of method arguments.
type Arguments Params

// ForMethodSignature builds a string that can be used to instantiate the method arguments part
// of a method declaration template.
func (as Arguments) ForMethodSignature() string {
	if len(as) == 0 {
		return ""
	}

	var argIdx int
	var pairs []string
	for _, a := range as {
		arg := a.Type
		if a.IsVariadic {
			// variadic argument has type []T. Here we turn it into ...T
			arg = "..." + strings.TrimPrefix(arg, "[]")
		}
		name := a.generatedName
		if a.Name != "" {
			name = fmt.Sprintf("%s /* %s */", name, a.Name)
		}
		if name != "" {
			arg = name + " " + arg
		}

		pairs = append(pairs, arg)

		argIdx++
	}

	return strings.Join(pairs, ", ")
}

// ContextVarName returns the name of the variable, that is of type context.Context, but only if
// it was the first variable to be passed to the method. It returns empty string in any other case.
func (as Arguments) ContextVarName() string {
	if len(as) == 0 {
		return ""
	}

	if as.firstVarTypeIsContext() {
		return as[0].generatedName
	}
	return ""
}

// ForMethodInvocationWithoutContext builds a string that holds a comma-separated list of
// all method argument names without the first context parameter. If the first parameter was not a
// context it is included in that list too.
//
// This method can be used when passing the arguments down to another method.
func (as Arguments) ForMethodInvocationWithoutContext() string {
	if len(as) == 0 {
		return ""
	}
	if as.firstVarTypeIsContext() {
		return Params(as[1:]).varNames()
	}
	return Params(as).varNames()
}

func (as Arguments) firstVarTypeIsContext() bool {
	if len(as) == 0 {
		return false
	}
	return as[0].Type == "context.Context"
}
