package packagecollision

import (
	afoo "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/a/foo"
	bfoo "github.com/Makpoc/gomigen/test/integration/testdata/packagecollision/b/foo"
)

type Compare interface {
	Equals(afoo.Foo, bfoo.Foo) bool
}
