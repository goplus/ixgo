// export by github.com/goplus/ixgo/cmd/qexp

package time

import (
	_ "github.com/goplus/xgo/tpl/variant/time"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/variant/time", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "time",
			Path: "github.com/goplus/xgo/tpl/variant/time",
			Deps: map[string]string{
				"github.com/goplus/xgo/tpl/variant": "variant",
				"time":                              "time",
			},
		}
	})
}
