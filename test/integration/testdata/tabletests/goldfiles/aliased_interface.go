// Code generated by middleware generator version "test" DO NOT EDIT.

package interfacesmw

import (
	"context"

	"github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces"
	"github.com/Makpoc/gomigen/types"
)

type AliasedInterfaceMiddleware struct {
	next interfaces.AliasedInterface
	hook types.Hook
}

// if this check fails middleware needs to be re-generated.
var _ interfaces.AliasedInterface = (*AliasedInterfaceMiddleware)(nil)

func NewAliasedInterfaceMiddleware(
	next interfaces.AliasedInterface,
	hook types.Hook,
) *AliasedInterfaceMiddleware {
	return &AliasedInterfaceMiddleware{
		next: next,
		hook: hook,
	}
}

func (mw *AliasedInterfaceMiddleware) AMethod(arg0 /* t */ string) error {
	methodInfo := types.MethodInfo{
		Package:   "github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces",
		Interface: "AliasedInterface",
		Method:    "AMethod",
		Params:    []interface{}{arg0},
	}

	ctx := context.Background()

	ctx = mw.hook.OnEntry(ctx, methodInfo)
	res0 := mw.next.AMethod(arg0)
	mw.hook.OnExit(ctx, methodInfo, res0)
	return res0
}