package generator

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Makpoc/gomigen/internal/model"
	"golang.org/x/tools/imports"
)

// GenerateSource instantiates the middleware template with the provided model and returns a
// formatted go source code content.
func GenerateSource(mw *model.Middleware, moduleVersion string) ([]byte, error) {
	tmpl, err := template.
		New("middleware").
		Funcs(template.FuncMap{
			"version": func() string { return moduleVersion },
		}).
		Parse(sourceTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse source template: %w", err)
	}

	b := &bytes.Buffer{}
	err = tmpl.Execute(b, mw)
	if err != nil {
		return nil, fmt.Errorf("failed to render source template: %w", err)
	}

	formatted, err := imports.Process("", b.Bytes(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to format generated source file: %w", err)
	}

	return formatted, nil
}
