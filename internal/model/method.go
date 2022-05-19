package model

import (
	"fmt"
)

// Method represents a single method of an interface.
type Method struct {
	Name      string
	Arguments Arguments
	Returns   Returns
}

// AddArgument adds a method argument for this method.
func (m *Method) AddArgument(p Param) {
	p.generatedName = fmt.Sprintf("%s%d", "arg", len(m.Arguments))
	m.Arguments = append(m.Arguments, p)
}

// AddReturn adds a method return parameter for this method.
func (m *Method) AddReturn(p Param) {
	p.generatedName = fmt.Sprintf("%s%d", "res", len(m.Returns))
	m.Returns = append(m.Returns, p)
}
