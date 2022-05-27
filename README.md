# Go Middleware Generator

![build](https://github.com/Makpoc/gomigen/workflows/Go/badge.svg)
![lint](https://github.com/Makpoc/gomigen/workflows/golangci-lint/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/Makpoc/gomigen.svg)](https://pkg.go.dev/github.com/Makpoc/gomigen)

This tool can generate a generic middleware, that injects hooks around
the invocations of given interface' methods.

This tool is heavily inspired by [gentools](https://github.com/Bo0mer/gentools) and
[counterfeiter](https://github.com/maxbrunsfeld/counterfeiter/)

# Description

The generated middleware wraps around another implementation of the same interface
and calls specific hooks at specific points of the execution.

Given the following interface and `go:generate` command

```go
//go:generate go run github.com/Makpoc/gomigen/cmd/gomigen . Interface
type Interface interface {
    ProcessOne(string) (string, error)
    ProcessTwo(context.Context, int) bool
}
```

running `go:generate` will produce the following middleware:

### Struct

The generated struct implements the Interface

```go
type InterfaceMiddleware struct {
	next interfaces.Interface
	hook types.Hook
}
```

### Constructor

The middleware constructor wraps another implementation of that interface
and also accepts a hook implementation, that will be called when a method of
this middleware is called.

```go
func NewInterfaceMiddleware(next interfaces.Interface, hook types.Hook) *InterfaceMiddleware {
	return &InterfaceMiddleware{
		next: next,
		hook: hook,
	}
}
```

### Methods

Each generated method calls the `OnEntry` hook method, then calls the wrapped Interface
implementation and finally calls the `OnExit` hook method.

Depending on the wrapped method's signature the generated middleware method's implementation varies:

```go
func (mw *InterfaceMiddleware) ProcessOne(arg0 string) (string, error) {
	methodInfo := types.MethodInfo{
		Package:   "<full-package-path>",
		Interface: "Interface",
		Method:    "ProcessTwo",
	}

	// ProcessOne doesn't accept context so one is created
	ctx := context.Background()

	ctx = mw.hook.OnEntry(ctx, methodInfo)
	res0, res1 := mw.next.ProcessOne(arg0)
	// ProcessOne returns an error as last return parm and it's passed to the OnExit hook
	mw.hook.OnExit(ctx, methodInfo, res1)
	return res0, res1
}

func (mw *TwoMethodsWithArgAndReturnMiddleware) ProcessTwo(arg0 context.Context, arg1 int) bool {
	methodInfo := types.MethodInfo{
		Package:   "<full-package-path>",
		Interface: "Interface",
		Method:    "ProcessTwo",
	}

	// ProcessTwo accepts a context and it is used
	ctx := arg0

	ctx = mw.hook.OnEntry(ctx, methodInfo)
	res0 := mw.next.ProcessTwo(ctx, arg1)
	// ProcessTwo doesn't return an error type so OnExit always receives nil
	mw.hook.OnExit(ctx, methodInfo, nil)
	return res0
}
```

All hook methods accept a `Context` they can extract values from as well as `MethodInfo`, containing
information about the method being invoked.

If the Interface method accepts a `context.Context` as its first parameter it will be passed
to all hooks and the wrapped method. If the first parameter is not a context, a
`context.Background()` context will be instantiated and passed to all hook methods instead.

#### OnEntry

```go
OnEntry(context.Context, MethodInfo) context.Context
```

This hook method is called just before the wrapped method is invoked.

It returns a `context` that is passed to the `OnExit` hook methods as well as to the
wrapped method (if it accepts a context as its first argument). Hook implementations can
use this to store and communicate values within a method call context (e.g. span
start/end, processing duration etc.)

#### OnExit

```go
mw.hook.OnExit(context.Context, MethodInfo, error)
```

This hook method is called right after the wrapped method returns. If the wrapped
method's last return parameter is of type `error` it is passed to the OnExit
method. If not - the error parameter is set to `nil`.

# Usage

## Generating the Middleware

### Command line

```sh
go run github.com/Makpoc/gomigen/cmd/gomigen . InterfaceName
```

The destination directory where the generated package with the middleware will be saved
can be provided as an optional parameter.

```sh
go run github.com/Makpoc/gomigen/cmd/gomigen -out ../mw . InterfaceName
```

The generated file will import this module's `types` packages so make sure to add that
package to the `tools.go` file (see below).

### Go Generate

Simply inject a go:generate comment like the following:

```go
//go:generate go run github.com/Makpoc/gomigen/cmd/gomigen . InterfaceName
```

and then run `go generate` in your project's root folder.

Don't forget to underscore-import the package to the `tools.go` to lock its version and ensure
it's available when `go generate` is called.

`tools/tools.go`:

```go
import (
	// contains the executable code that does the middleware generation.
    _ "github.com/Makpoc/gomigen/cmd/gomigen"
	// contains the types, needed by the generated middleware for compilation.
    _ "github.com/Makpoc/gomigen/types"
)
```

# TODOs

* Support extracting a context from `http.Request` parameter
* Support the model counterfeiter uses with a
custom `//<tool>:generate` declaration for faster execution.
* Create a package with hooks for error logging, monitoring and tracing.
* Optionally expose an argument to generate a template that doesn't need to import the `types` package
  * The hooks interface can be inlined
  * MethodInfo can be turned into a `map[string]string`
