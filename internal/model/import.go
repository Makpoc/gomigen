package model

import (
	"fmt"
	"math"
)

// Import is an import declaration with an optional Alias and the package path.
type Import struct {
	Alias          string
	aliasGenerated bool
	Package        string
}

// Imports is a set of Import.
type Imports struct {
	byAlias   map[string]Import
	byPackage map[string]Import
}

func newImports() *Imports {
	return &Imports{
		byAlias:   make(map[string]Import),
		byPackage: make(map[string]Import),
	}
}

// Add adds an Import to the Imports list. It will automatically discover name collisions and
// assign a custom alias to the new import if such collisions are found. The returned value
// is the final alias of this package.
//
// Importing the same package multiple times will result in a single Import instance in the list.
func (il *Imports) Add(newImport Import) string {
	if imprt, seen := il.byPackage[newImport.Package]; seen {
		return imprt.Alias
	}

	imprt := Import{
		Alias:   newImport.Alias,
		Package: newImport.Package,
	}
	if _, seen := il.byAlias[newImport.Alias]; seen {
		// alias is already taken. We need a new one
		imprt.aliasGenerated = true
		imprt.Alias = il.findFreeAlias(newImport)
	}

	il.byAlias[imprt.Alias] = imprt
	il.byPackage[imprt.Package] = imprt

	return imprt.Alias
}

// ImportStatements builds a list of imports that can be used when instantiating the import
// section of a go source templates. If an import was renamed due to collision the Import alias
// will be set. If it has the original name resolved form the package path the Alias will be empty.
func (il *Imports) ImportStatements() []Import {
	result := make([]Import, 0, len(il.byAlias))
	for alias, imprt := range il.byAlias {
		resultImprt := Import{
			Alias:   "",
			Package: imprt.Package,
		}
		if imprt.aliasGenerated {
			resultImprt.Alias = alias
		}
		result = append(result, resultImprt)
	}
	return result
}

func (il *Imports) findFreeAlias(imprt Import) string {
	for i := 0; i < math.MaxInt; i++ {
		alias := fmt.Sprintf("%s%d", imprt.Alias, i)
		if _, seen := il.byAlias[alias]; !seen {
			return alias
		}
	}
	panic(fmt.Sprintf("no free aliases available for package %q", imprt.Package))
}
