package example

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Makpoc/gomigen/types"
)

type ctxStartTimeKey struct{}

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
	l.logger.Printf("Entered %s at %s",
		formatMethodInfo(info), startTime,
	)
	return context.WithValue(ctx, ctxStartTimeKey{}, startTime)
}

func (l *LogHook) OnExit(ctx context.Context, info types.MethodInfo, err error) {
	startTime := ctx.Value(ctxStartTimeKey{}).(time.Time)
	result := "ok"
	if err != nil {
		result = fmt.Sprintf("error: %v", err)
	}
	l.logger.Printf("Exited %s after spending %s in it. Result: %s",
		formatMethodInfo(info), time.Since(startTime).String(), result,
	)
}

func formatMethodInfo(info types.MethodInfo) string {
	return fmt.Sprintf("%s/%s.%s", info.Package, info.Interface, info.Method)
}
