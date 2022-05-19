package inner

import (
	"context"
)

type PackageInterface interface {
	ProcessInner(ctx context.Context) error
}
