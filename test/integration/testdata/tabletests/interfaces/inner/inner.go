package inner

import (
	"context"
)

type PackageInterface interface {
	ProcessInner(ctx context.Context) error
}

type GenericInterface[T any] interface {
	AMethod(t T) error
}
