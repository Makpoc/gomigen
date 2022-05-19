package interfaces

import (
	"context"
	"io"

	"github.com/Makpoc/gomigen/test/integration/testdata/tabletests/interfaces/inner"
)

type SingleMethodNoArgsNoReturns interface {
	Process()
}
type SingleMethodWithArgNoReturns interface {
	Process(int)
}
type SingleMethodNoArgWithReturn interface {
	Process() int
}
type SingleMethodWithArgAndReturn interface {
	Process(int) int
}

type ContextFirstArgument interface {
	Process(context.Context, int)
}
type ReturnsSingleError interface {
	Process(int) error
}
type ReturnsMultipleValuesWithError interface {
	Process(int) (int, error)
}
type ReturnsMultipleValuesNoError interface {
	Process(int) (int, error)
}

type ContextFirstArgumentReturnsError interface {
	Process(context.Context, int) (int, error)
}

type ContextFirstArgumentReturnsErrorNamedVars interface {
	Process(argCtx context.Context, argInt int) (returnInt int, returnErr error)
}

type TwoMethodsOneWithContextAndError interface {
	ProcessOne(context.Context, string) (string, error)
	ProcessTwo(argInt int) (returnInt int)
}

type MapInArgAndReturn interface {
	Process(map[string]int) map[interface{}]*interface{}
}
type SliceInArgAndReturn interface {
	Process([]int) []*interface{}
}

type VariadicArgument interface {
	Process(string, ...int)
}

type EmbeddedInterface interface {
	ContextFirstArgumentReturnsError
}

type CustomInterfaceInArgAndReturn interface {
	Process(ContextFirstArgumentReturnsError) ReturnsMultipleValuesNoError
}

type EmptyInterface interface{}

type InnerPackageReference interface {
	Process(inner.PackageInterface) inner.PackageInterface
}

type EmbedStandardInterface interface {
	io.Writer
}

type EmbedCustomInterface interface {
	InnerPackageReference
}

type ComposeMultipleInterface interface {
	io.Writer
	inner.PackageInterface
}

type ExtendsAnotherInterface interface {
	io.Writer
	Close() error
}
