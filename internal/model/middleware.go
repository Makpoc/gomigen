package model

// Middleware is the top-level model that wraps all the parsed elements of an Interface. It and
// its child models all contain methods, that can be used to instantiate a go source template.
type Middleware struct {
	// Imports contain all imports needed by the middleware package.
	// Within one middleware all imports are uniquely identified by their aliases.
	Imports *Imports
	// InterfaceName is the name of the interface we are generating a middleware for.
	InterfaceName string
	// OriginalPackage describes the package of the interface we are generating middleware for.
	OriginalPackage Package
	// GenPackage holds the name of the package to be set for all middleware files.
	GenPackage string
	// Methods holds the discovered methodset of the interface we are generating a middleware for.
	Methods []Method

	middlewareTypesPackageAlias string
}

// NewMiddleware creates a new Middleware instance with the mandatory-needed imports already added.
func NewMiddleware(forInterfaceName string) *Middleware {
	mw := &Middleware{
		InterfaceName: forInterfaceName,
		Imports:       NewImports(),
		Methods:       make([]Method, 0),
	}
	mw.Imports.Add(Import{
		Alias:   "context",
		Package: "context",
	})
	// needed to resolve Hooks and MethodInfo
	mw.middlewareTypesPackageAlias = mw.Imports.Add(Import{
		Alias:   "types",
		Package: "github.com/Makpoc/gomigen/types",
	})
	return mw
}

// GetMiddlewareTypesPackageAlias returns the alias of the types package in this library.
func (mw *Middleware) GetMiddlewareTypesPackageAlias() string {
	return mw.middlewareTypesPackageAlias
}

// AddMethod adds a new interface method.
func (mw *Middleware) AddMethod(m Method) {
	mw.Methods = append(mw.Methods, m)
}

// MiddlewareName builds the name of the generated middleware based on the original interface name.
func (mw *Middleware) MiddlewareName() string {
	return mw.InterfaceName + "Middleware"
}

// UniqueImportPaths returns the generated import statements with custom aliases where needed.
func (mw *Middleware) UniqueImportPaths() []Import {
	return mw.Imports.ImportStatements()
}
