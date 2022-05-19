package generator

import (
	_ "embed"
)

//go:embed template.tmpl
var sourceTemplate string
