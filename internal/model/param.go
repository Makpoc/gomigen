package model

import (
	"strings"
)

// Params is a list of params
type Params []Param

// Param holds the parameter name and its type ([package name if any].[type]).
type Param struct {
	Name          string
	generatedName string
	Type          string
	IsVariadic    bool
}

func (pl Params) varNames() string {
	varNames := make([]string, 0, len(pl))
	var varIdx int
	for _, p := range pl {
		name := p.generatedName
		if p.IsVariadic {
			name = name + "..."
		}
		varNames = append(varNames, name)
		varIdx++
	}
	return strings.Join(varNames, ", ")
}
