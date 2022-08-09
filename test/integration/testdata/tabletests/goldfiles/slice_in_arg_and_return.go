// Code generated by middleware generator version "test" DO NOT EDIT.

package interfacesmw

import (
	"context"

	"github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces"
	"github.com/Makpoc/gomigen/types"
)

type SliceInArgAndReturnMiddleware struct {
	next interfaces.SliceInArgAndReturn
	hook types.Hook
}

// if this check fails middleware needs to be re-generated.
var _ interfaces.SliceInArgAndReturn = (*SliceInArgAndReturnMiddleware)(nil)

func NewSliceInArgAndReturnMiddleware(
	next interfaces.SliceInArgAndReturn,
	hook types.Hook,
) *SliceInArgAndReturnMiddleware {
	return &SliceInArgAndReturnMiddleware{
		next: next,
		hook: hook,
	}
}

func (mw *SliceInArgAndReturnMiddleware) Process(arg0 []int) []*interface{} {
	methodInfo := types.MethodInfo{
		Package:   "github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces",
		Interface: "SliceInArgAndReturn",
		Method:    "Process",
		Params:    []interface{}{arg0},
	}

	ctx := context.Background()

	ctx = mw.hook.OnEntry(ctx, methodInfo)
	res0 := mw.next.Process(arg0)
	mw.hook.OnExit(ctx, methodInfo, nil)
	return res0
}
