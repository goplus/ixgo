//go:build go1.22

package main

import "database/sql"

func main() {
	var n sql.Null[string]
	if e := n.Scan("value"); e != nil || !n.Valid || n.V != "value" {
		panic("Null.Scan")
	}
	if v, e := n.Value(); e != nil || v != "value" {
		panic("Null.Value")
	}
}
