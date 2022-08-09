package types

// MethodInfo holds information about the method that is calling a hook.
type MethodInfo struct {
	Package   string
	Interface string
	Method    string
	Params    []interface{}
}
