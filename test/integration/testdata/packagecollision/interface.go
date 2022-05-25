package packagecollision

import (
	afoo "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/a/foo"
	bfoo "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/b/foo"
	. "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/dotimported/foo"
	// do not alias this import when editing. goimports will try to do it automatically
	// so edit this file in an enditor without goimports before committing.
	"github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/renamedpackage/bar"
)

type Compare interface {
	Equals(afoo.Foo, bfoo.Foo, Foo, foo.Foo) bool
}
