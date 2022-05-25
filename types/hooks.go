package types

import "context"

// Hook represents the interface that can be passed to a generated middleware.
type Hook interface {
	// OnEntry is invoked immediately after entering the function.
	//
	// The returned context is passed to the other hook methods. If the
	// function's first argument is of type context.Context it is passed
	// to the OnEntry hook and is reassigned to the returned context.
	// As such the OnEntry hook must at least return the received context
	// even if it does not use it for anything else.
	OnEntry(context.Context, MethodInfo) context.Context
	// OnExit is invoked just before the middleware function returns.
	//
	// If the function returns an error as its last return parameter that
	// error will be passed to the hook.
	//
	// If the function's last return parameter is not of type error or
	// the returned error is nil the hook will receive nil.
	OnExit(context.Context, MethodInfo, error)
}
