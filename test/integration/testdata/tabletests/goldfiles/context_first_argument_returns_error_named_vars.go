// Code generated by middleware generator version "test" DO NOT EDIT.

package interfacesmw

import (
	"context"

	"github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces"
	"github.com/Makpoc/gomigen/types"
)

type ContextFirstArgumentReturnsErrorNamedVarsMiddleware struct {
	next interfaces.ContextFirstArgumentReturnsErrorNamedVars
	hook types.Hook
}

// if this check fails middleware needs to be re-generated.
var _ interfaces.ContextFirstArgumentReturnsErrorNamedVars = (*ContextFirstArgumentReturnsErrorNamedVarsMiddleware)(nil)

func NewContextFirstArgumentReturnsErrorNamedVarsMiddleware(
	next interfaces.ContextFirstArgumentReturnsErrorNamedVars,
	hook types.Hook,
) *ContextFirstArgumentReturnsErrorNamedVarsMiddleware {
	return &ContextFirstArgumentReturnsErrorNamedVarsMiddleware{
		next: next,
		hook: hook,
	}
}

func (mw *ContextFirstArgumentReturnsErrorNamedVarsMiddleware) Process(arg0 /* argCtx */ context.Context, arg1 /* argInt */ int) (int /* returnInt */, error /* returnErr */) {
	methodInfo := types.MethodInfo{
		Package:   "github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces",
		Interface: "ContextFirstArgumentReturnsErrorNamedVars",
		Method:    "Process",
		Params:    []interface{}{arg0, arg1},
	}

	ctx := arg0

	ctx = mw.hook.OnEntry(ctx, methodInfo)
	res0, res1 := mw.next.Process(ctx, arg1)
	mw.hook.OnExit(ctx, methodInfo, res1)
	return res0, res1
}
