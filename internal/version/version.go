package version

import (
	"runtime/debug"
)

// Version returns the module version or "unknown" if the version cannot be determined.
func Version() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	return bi.Main.Version
}
