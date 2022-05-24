package packagecollision

import (
	afoo "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/a/foo"
	bfoo "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/b/foo"
	. "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/dotimported/foo"
	barfoo "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/renamedpackage/bar"
)

type Compare interface {
	Equals(afoo.Foo, bfoo.Foo, Foo, barfoo.Foo) bool
}
