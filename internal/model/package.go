package model

// Package holds a package name and its full import path.
//
// Note that Name here is NOT the alias, but the name it's imported by default. (e.g. module
// may have package path github.com/yetanotherlogger/log/v2 but the imported package is
// named "log").
type Package struct {
	Name string
	Path string
}
