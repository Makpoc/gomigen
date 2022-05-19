# Go Middleware Generator

This tool can generate a generic middleware, that injects hooks around
the invocations of given interface' methods.

# Description

The generated middleware wraps around another implementation of the same interface
and calls specific hooks at specific points in the execution.

Given the following interface and `go:generate` command
```go
//go:generate go run github.com/Makpoc/gomigen/cmd/gomigen . Interface
type Interface interface {
    ProcessOne(string) (string, error)
    ProcessTwo(context.Context, int) bool
}
```

running `go:generate` will produce the following Middleware:

* Struct:

The generated struct implements the Interface

```go
type InterfaceMiddleware struct {
	next interfaces.Interface
	hook types.Hook
}
```

* Constructor

The middleware constructor wraps another implementation of that interface
and also accepts a hook implementation, that will be called when a method of
this middleware is called.

```go
func NewInterfaceMiddleware(
	next interfaces.Interface,
	hook types.Hook,
) *InterfaceMiddleware {
	return &InterfaceMiddleware{
		next: next,
		hook: hook,
	}
}
```

* Methods

Each generated method calls the OnEntry hook, then calls the wrapped Interface
implementation and finally calls the OnExit hook.

Depending on that method arguments and return parameters the hooks and the
wrapped method receive different values. See the inlined comments for notable
differences.

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

All hooks accept a `Context` they can extract values from as well as `MethodInfo`, containing
information about the method being invoked. Some hooks accept additional parameters also
return a value. These are covered below.

If the Interface method accepts a `context.Context` as its first parameter it will be passed
to all hooks and the wrapped method. If the first parameter is not a context, a
`context.Background()` context will be instantiated and passed to all hooks instead.

### OnEntry

```go
OnEntry(context.Context, MethodInfo) context.Context
```

This hook is called just before the wrapped method is called.

It returns a `context` that is passed to the rest of the hooks as well as to the
wrapped method. Hook implementations can use this to store and communicate values
within a method call context (e.g. span start/end, processing duration etc.)

### OnExit

```go
mw.hook.OnExit(context.Context, MethodInfo, error)
```

This hook is called right after the wrapped method returns. If the wrapped
method's last return parameter was of type `error` it is passed to the OnExit
method. If not - OnExit gets nil.

# Usage

## Generating the Middleware

### Command line

```sh
go run github.com/Makpoc/gomigen/cmd/gomigen . InterfaceName
```

The destination directory where the generated package with the middleware will be saved
can be provided as an optional parameter.

```sh
go run github.com/Makpoc/gomigen/cmd/gomigen -o ../mw . InterfaceName
```

The generated file will import this module's `types` packages so make sure to add that
package to the `tools.go` file (see below).

### Go Generate

Simply inject a go comment like you would with counterfeiter or gentools

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
* Support the model counterfeiter uses with a custom `//<tool>:generate` declaration for faster execution.
* Create a package with hooks for error logging, monitoring and tracing.
* Optionally expose an option to generate a template that doesn't need to import the `types` package
  * The hooks interface can be inlined
  * MethodInfo can be turned into a `map[string]string`