//go:build go1.27
// +build go1.27

package jsontext

import (
	_ "embed"
	q "encoding/json/jsontext"

	"github.com/goplus/ixgo"
)

//go:embed _patch/jsontext.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("encoding/json/jsontext@patch.appendFormatBytes", appendFormatBytes)
	ixgo.RegisterExternal("encoding/json/jsontext@patch.appendQuoteBytes", appendQuoteBytes)
	ixgo.RegisterExternal("encoding/json/jsontext@patch.appendUnquoteBytes", appendUnquoteBytes)
	ixgo.RegisterPatch("encoding/json/jsontext", patch_data)
}

func appendFormatBytes(dst, src []byte, opts ...q.Options) ([]byte, error) {
	return q.AppendFormat[[]byte](dst, src, opts...)
}

func appendQuoteBytes(dst, src []byte) ([]byte, error) {
	return q.AppendQuote[[]byte](dst, src)
}

func appendUnquoteBytes(dst, src []byte) ([]byte, error) {
	return q.AppendUnquote[[]byte](dst, src)
}
