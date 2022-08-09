package model

import (
	"fmt"
	"strings"
)

const (
	typeError = "error"
)

// Returns is the list of method's return parameters.
type Returns Params

// LastVarTypeIsError reports if the last return parameter is of type "error".
func (rs Returns) LastVarTypeIsError() bool {
	if len(rs) == 0 {
		return false
	}
	lastParam := rs[len(rs)-1]
	return lastParam.Type == typeError
}

// ErrorVarName returns the name of the (last) error parameter of a method or an empty string if
// it's not of type "error".
func (rs Returns) ErrorVarName() string {
	if len(rs) == 0 {
		return ""
	}

	lastVar := rs[len(rs)-1]
	if lastVar.Type != "error" {
		return ""
	}

	return lastVar.generatedName
}

// ForMethodSignature builds a string that can be used to instantiate the method return parameters
// part of a method declaration template.
//
// The result removes the names of named return parameters but keeps them in comments
// next to the type.
func (rs Returns) ForMethodSignature() string {
	if len(rs) == 0 {
		return ""
	}

	var returns []string
	for _, r := range rs {
		res := r.Type
		if r.Name != "" {
			// gofmt moves the comments before the commas, which looks a bit weird:
			//   (/* val */ interface{} /* err */, error)
			// ----------------------------------^-------
			// this is why we put [type] /*varName*/ in that order here.
			res = fmt.Sprintf("%s /* %s */", res, r.Name)
		}

		returns = append(returns, res)
	}

	return strings.Join(returns, ", ")
}

// ReturnVarNames builds a comma-separated list of all return parameters variable names.
func (rs Returns) ReturnVarNames() string {
	return Params(rs).varNames(true)
}
