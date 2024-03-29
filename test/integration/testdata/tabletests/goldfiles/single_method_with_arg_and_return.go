// Code generated by middleware generator version "test" DO NOT EDIT.

package interfacesmw

import (
	"context"

	"github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces"
	"github.com/Makpoc/gomigen/types"
)

type SingleMethodWithArgAndReturnMiddleware struct {
	next interfaces.SingleMethodWithArgAndReturn
	hook types.Hook
}

// if this check fails middleware needs to be re-generated.
var _ interfaces.SingleMethodWithArgAndReturn = (*SingleMethodWithArgAndReturnMiddleware)(nil)

func NewSingleMethodWithArgAndReturnMiddleware(
	next interfaces.SingleMethodWithArgAndReturn,
	hook types.Hook,
) *SingleMethodWithArgAndReturnMiddleware {
	return &SingleMethodWithArgAndReturnMiddleware{
		next: next,
		hook: hook,
	}
}

func (mw *SingleMethodWithArgAndReturnMiddleware) Process(arg0 int) int {
	methodInfo := types.MethodInfo{
		Package:   "github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces",
		Interface: "SingleMethodWithArgAndReturn",
		Method:    "Process",
		Params:    []interface{}{arg0},
	}

	ctx := context.Background()

	ctx = mw.hook.OnEntry(ctx, methodInfo)
	res0 := mw.next.Process(arg0)
	mw.hook.OnExit(ctx, methodInfo, nil)
	return res0
}
