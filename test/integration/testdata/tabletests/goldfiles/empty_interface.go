// Code generated by middleware generator version "test" DO NOT EDIT.

package interfacesmw

import (
	"github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces"
	"github.com/Makpoc/gomigen/types"
)

type EmptyInterfaceMiddleware struct {
	next interfaces.EmptyInterface
	hook types.Hook
}

// if this check fails middleware needs to be re-generated.
var _ interfaces.EmptyInterface = (*EmptyInterfaceMiddleware)(nil)

func NewEmptyInterfaceMiddleware(
	next interfaces.EmptyInterface,
	hook types.Hook,
) *EmptyInterfaceMiddleware {
	return &EmptyInterfaceMiddleware{
		next: next,
		hook: hook,
	}
}
