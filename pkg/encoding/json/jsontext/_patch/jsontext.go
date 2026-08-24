package jsontext

import q "encoding/json/jsontext"

// AppendFormat formats a JSON value and appends it to dst.
func AppendFormat[Bytes ~[]byte | ~string](dst []byte, src Bytes, opts ...q.Options) ([]byte, error) {
	return appendFormatBytes(dst, []byte(src), opts...)
}

// AppendQuote appends a quoted JSON string to dst.
func AppendQuote[Bytes ~[]byte | ~string](dst []byte, src Bytes) ([]byte, error) {
	return appendQuoteBytes(dst, []byte(src))
}

// AppendUnquote appends the decoded JSON string to dst.
func AppendUnquote[Bytes ~[]byte | ~string](dst []byte, src Bytes) ([]byte, error) {
	return appendUnquoteBytes(dst, []byte(src))
}

func appendFormatBytes(dst, src []byte, opts ...q.Options) ([]byte, error)
func appendQuoteBytes(dst, src []byte) ([]byte, error)
func appendUnquoteBytes(dst, src []byte) ([]byte, error)
