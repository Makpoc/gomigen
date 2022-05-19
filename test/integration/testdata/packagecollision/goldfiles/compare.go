// Code generated by middleware generator version "test" DO NOT EDIT.

package packagecollisionmw

import (
	"context"

	"github.com/Makpoc/gomigen/test/integration/testdata/packagecollision"
	"github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/a/foo"
	foo0 "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/b/foo"
	"github.com/Makpoc/gomigen/types"
)

type CompareMiddleware struct {
	next packagecollision.Compare
	hook types.Hook
}

// if this check fails middleware needs to be re-generated.
var _ packagecollision.Compare = (*CompareMiddleware)(nil)

func NewCompareMiddleware(
	next packagecollision.Compare,
	hook types.Hook,
) *CompareMiddleware {
	return &CompareMiddleware{
		next: next,
		hook: hook,
	}
}

func (mw *CompareMiddleware) Equals(arg0 foo.Foo, arg1 foo0.Foo) bool {
	methodInfo := types.MethodInfo{
		Package:   "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision",
		Interface: "Compare",
		Method:    "Equals",
	}

	ctx := context.Background()

	ctx = mw.hook.OnEntry(ctx, methodInfo)
	res0 := mw.next.Equals(arg0, arg1)
	mw.hook.OnExit(ctx, methodInfo, nil)
	return res0
}
