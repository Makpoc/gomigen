package parser

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
	"golang.org/x/tools/imports"

	"github.com/Makpoc/gomigen/internal/model"
)

// Parse loads the package under dir and looks for an interface with the given name.
// If one is found Parse parses its components to build a model.Middleware.
func Parse(dir, interfaceName string) (*model.Middleware, error) {
	targetPackage, targetInterface, err := loadInterface(dir, interfaceName)
	if err != nil {
		return nil, err
	}

	middleware := model.NewMiddleware(interfaceName)

	originalInterfacePkgAlias := middleware.Imports.Add(model.Import{
		Alias:   targetPackage.Name,
		Package: targetPackage.PkgPath,
	})
	middleware.OriginalPackage = model.Package{
		Name: originalInterfacePkgAlias,
		Path: targetPackage.PkgPath,
	}

	middleware.GenPackage = targetPackage.Name + "mw"

	mSet := typeutil.IntuitiveMethodSet(targetInterface, nil)
	for _, m := range mSet {
		err = processMethod(middleware, m)
		if err != nil {
			return nil, fmt.Errorf("failed to parse method: %w", err)
		}
	}
	return middleware, nil
}

func loadInterface(dir string, interfaceName string) (pkg *packages.Package, iface *types.Named, err error) {
	p, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedImports | packages.NeedTypes,
	}, dir)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load packages in directory %q: %w", dir, err)
	}
	if len(p) == 0 {
		return nil, nil, fmt.Errorf("no go packages found in directory %q", dir)
	}

	targetPackage := p[0]
	pkgScope := targetPackage.Types.Scope()
	targetIFaceObj := pkgScope.Lookup(interfaceName)
	if targetIFaceObj == nil {
		return nil, nil, fmt.Errorf("no interface named %s in package %s", interfaceName, targetPackage.PkgPath)
	}
	targetIFace, ok := (targetIFaceObj.Type()).(*types.Named)
	if !ok {
		return nil, nil, fmt.Errorf("%s is not an interface declaration", interfaceName)
	}

	return targetPackage, targetIFace, nil
}

func processMethod(middleware *model.Middleware, method *types.Selection) error {
	if method == nil || method.Obj() == nil || method.Type() == nil {
		return fmt.Errorf("unexpected method definition: %+v", method)
	}

	f, ok := method.Obj().(*types.Func)
	if !ok {
		// for go 1.17 interface this should not happen.
		// TODO test with go 1.18 and interface with generic constraints.
		return nil
	}
	m := model.Method{
		Name: f.Name(),
	}

	sig, ok := method.Type().(*types.Signature)
	if !ok {
		return fmt.Errorf("unexpected method type: %+v", method.Type())
	}

	err := processMethodArguments(middleware, &m, sig)
	if err != nil {
		return fmt.Errorf("failed to parse arguments for method %q: %w", m.Name, err)
	}

	err = processMethodReturns(middleware, &m, sig)
	if err != nil {
		return fmt.Errorf("failed to parse return parameters for method %q: %w", m.Name, err)
	}

	middleware.AddMethod(m)

	return nil
}

func processMethodArguments(middleware *model.Middleware, method *model.Method, sig *types.Signature) error {
	args := sig.Params()
	for i := 0; i < args.Len(); i++ {
		arg := args.At(i)

		methodArg, err := parseParameter(middleware, arg)
		if err != nil {
			return fmt.Errorf("failed to parse argument %q: %w", arg.String(), err)
		}

		if i == args.Len()-1 {
			methodArg.IsVariadic = sig.Variadic()
		}

		method.AddArgument(methodArg)
	}
	return nil
}

func processMethodReturns(middleware *model.Middleware, method *model.Method, sig *types.Signature) error {
	methodReturns := sig.Results()
	for i := 0; i < methodReturns.Len(); i++ {
		ret := methodReturns.At(i)

		param, err := parseParameter(middleware, ret)
		if err != nil {
			return fmt.Errorf("failed to parse return parameter %q: %w", ret.String(), err)
		}

		method.AddReturn(param)
	}
	return nil
}

func parseParameter(middleware *model.Middleware, v *types.Var) (model.Param, error) {
	param := model.Param{
		Name: v.Name(),
		Type: types.TypeString(v.Type(), captureImports(middleware.Imports)),
	}
	return param, nil
}

// captureImports builds a function, suitable for use as a types.TypeString Qualifier. It resolves the package names,
// and based on the package properties builds and adds a new Import to the provided Imports object.
func captureImports(dst *model.Imports) func(p *types.Package) string {
	return func(p *types.Package) string {
		alias := dst.Add(model.Import{
			Alias:   p.Name(),
			Package: imports.VendorlessPath(p.Path()),
		})
		return alias
	}
}
