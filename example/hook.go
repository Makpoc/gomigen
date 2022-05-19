package example

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Makpoc/gomigen/types"
)

type ctxKey string

const (
	startTimeKey ctxKey = "startTime"
)

// LogHook is an example of a hook that logs messages on entry and exit and the time interval
// between they are called.
type LogHook struct {
	logger *log.Logger
}

var _ types.Hook = (*LogHook)(nil)

func New(log *log.Logger) *LogHook {
	return &LogHook{
		logger: log,
	}
}

func (l *LogHook) OnEntry(ctx context.Context, info types.MethodInfo) context.Context {
	startTime := time.Now()
	ctx = context.WithValue(ctx, startTimeKey, startTime)
	l.logger.Printf("Entered %s at %s",
		formatMethodInfo(info), startTime,
	)
	return ctx
}

func (l *LogHook) OnExit(ctx context.Context, info types.MethodInfo, err error) {
	duration := "unknown"
	startTime := ctx.Value(startTimeKey)
	if startTime != nil {
		duration = time.Since(startTime.(time.Time)).String()
	}

	l.logger.Printf("Exited %s after spending %s in it",
		formatMethodInfo(info), duration,
	)
}

func formatMethodInfo(info types.MethodInfo) string {
	return fmt.Sprintf("%s/%s.%s", info.Package, info.Interface, info.Method)
}
