//go:build go1.27
// +build go1.27

package json

import (
	_ "embed"
	qt "encoding/json/jsontext"
	q "encoding/json/v2"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/encoding/json/jsontext"
)

//go:embed _patch/json.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("encoding/json/v2@patch.getOptionAny", getOptionAny)
	ixgo.RegisterExternal("encoding/json/v2@patch.getOptionBool", getOptionBool)
	ixgo.RegisterExternal("encoding/json/v2@patch.marshalFuncAny", marshalFuncAny)
	ixgo.RegisterExternal("encoding/json/v2@patch.marshalToFuncAny", marshalToFuncAny)
	ixgo.RegisterExternal("encoding/json/v2@patch.unmarshalFuncAny", unmarshalFuncAny)
	ixgo.RegisterExternal("encoding/json/v2@patch.unmarshalFromFuncAny", unmarshalFromFuncAny)
	ixgo.RegisterPatch("encoding/json/v2", patch_data)
}

func getOptionAny(opts q.Options, setter func(any) q.Options) (any, bool) {
	return q.GetOption[any](opts, setter)
}

func getOptionBool(opts q.Options, setter func(bool) q.Options) (bool, bool) {
	return q.GetOption[bool](opts, setter)
}

func marshalFuncAny(fn func(any) ([]byte, error), _ any) *q.Marshalers {
	return q.MarshalFunc[any](fn)
}

func marshalToFuncAny(fn func(*qt.Encoder, any) error, _ any) *q.Marshalers {
	return q.MarshalToFunc[any](fn)
}

func unmarshalFuncAny(fn func([]byte, any) error, _ any) *q.Unmarshalers {
	return q.UnmarshalFunc[any](fn)
}

func unmarshalFromFuncAny(fn func(*qt.Decoder, any) error, _ any) *q.Unmarshalers {
	return q.UnmarshalFromFunc[any](fn)
}
