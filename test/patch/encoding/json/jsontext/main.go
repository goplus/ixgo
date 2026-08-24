//go:build go1.27

package main

import "encoding/json/jsontext"

func main() {
	q, e := jsontext.AppendQuote[string](nil, "x\n")
	if e != nil {
		panic(e)
	}
	u, e := jsontext.AppendUnquote[string](nil, string(q))
	if e != nil || string(u) != "x\n" {
		panic("unquote")
	}
	f, e := jsontext.AppendFormat[string](nil, "{\"a\": 1}")
	if e != nil || string(f) != "{\"a\":1}" {
		panic("format")
	}
}
