// Code generated by middleware generator version "test" DO NOT EDIT.

package interfacesmw

import (
	"context"

	"github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces"
	"github.com/Makpoc/gomigen/types"
)

type MapInArgAndReturnMiddleware struct {
	next interfaces.MapInArgAndReturn
	hook types.Hook
}

// if this check fails middleware needs to be re-generated.
var _ interfaces.MapInArgAndReturn = (*MapInArgAndReturnMiddleware)(nil)

func NewMapInArgAndReturnMiddleware(
	next interfaces.MapInArgAndReturn,
	hook types.Hook,
) *MapInArgAndReturnMiddleware {
	return &MapInArgAndReturnMiddleware{
		next: next,
		hook: hook,
	}
}

func (mw *MapInArgAndReturnMiddleware) Process(arg0 map[string]int) map[interface{}]*interface{} {
	methodInfo := types.MethodInfo{
		Package:   "github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces",
		Interface: "MapInArgAndReturn",
		Method:    "Process",
		Params:    []interface{}{arg0},
	}

	ctx := context.Background()

	ctx = mw.hook.OnEntry(ctx, methodInfo)
	res0 := mw.next.Process(arg0)
	mw.hook.OnExit(ctx, methodInfo, nil)
	return res0
}
