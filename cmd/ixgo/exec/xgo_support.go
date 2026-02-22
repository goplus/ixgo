//go:build linknamefix
// +build linknamefix

package exec

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"io"
	"net/http"
	"time"
	_ "unsafe"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/xgobuild"
	_ "github.com/goplus/ixgo/xgobuild/pkg"
)

type operandMode byte
type builtinId int

type operand struct {
	mode operandMode
	expr ast.Expr
	typ  types.Type
	val  constant.Value
	id   builtinId
}

type positioner interface {
	Pos() token.Pos
}

//go:linkname checker_infer go/types.(*Checker).infer
func checker_infer(check *types.Checker, posn positioner, tparams []*types.TypeParam, targs []types.Type, params *types.Tuple, args []*operand) (result []types.Type)

//go:linkname serveContent net/http.serveContent
func serveContent(w http.ResponseWriter, r *http.Request, name string, modtime time.Time, sizeFunc func() (int64, error), content io.ReadSeeker)

//go:linkname toHTTPError net/http.toHTTPError
func toHTTPError(err error) (msg string, httpStatus int)

func init() {
	// support github.com/goplus/gogen.checker_infer
	ixgo.RegisterExternal("go/types.(*Checker).infer", checker_infer)
	ixgo.RegisterExternal("net/http.serveContent", serveContent)
	ixgo.RegisterExternal("net/http.toHTTPError", toHTTPError)
}
