//go:build go1.27

package main

import (
	"bytes"
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"errors"
)

func main() {
	m := json.MarshalFunc[*string](func(s *string) ([]byte, error) { return []byte("\"custom:" + *s + "\""), nil })
	source := "x"
	b, err := json.Marshal(&source, json.WithMarshalers(m))
	if err != nil || string(b) != `"custom:x"` {
		panic("MarshalFunc")
	}
	mt := json.MarshalToFunc[*string](func(enc *jsontext.Encoder, s *string) error { return enc.WriteToken(jsontext.String("to:" + *s)) })
	b, err = json.Marshal(&source, json.WithMarshalers(mt))
	if err != nil || string(b) != `"to:x"` {
		panic("MarshalToFunc")
	}
	u := json.UnmarshalFunc[*string](func(data []byte, p *string) error {
		if string(data) != `"x"` {
			return errors.New("input")
		}
		*p = "decoded"
		return nil
	})
	var s string
	if err = json.Unmarshal([]byte(`"x"`), &s, json.WithUnmarshalers(u)); err != nil || s != "decoded" {
		panic("UnmarshalFunc")
	}
	uf := json.UnmarshalFromFunc[*string](func(dec *jsontext.Decoder, p *string) error {
		value, err := dec.ReadValue()
		if err != nil || string(value) != `"x"` {
			return errors.New("token")
		}
		*p = "from"
		return nil
	})
	dec := jsontext.NewDecoder(bytes.NewReader([]byte(`"x"`)))
	if err = json.UnmarshalDecode(dec, &s, json.WithUnmarshalers(uf)); err != nil || s != "from" {
		panic("UnmarshalFromFunc")
	}
}
