// Package modwaretest provides common constants and functions for unit testing
package modwaretest

import (
	"bytes"
	"encoding/json"
)

const (
	TestAPIHost    = "https://dictybase.org"
	TestPathPrefix = "1.0"
	TestPubId      = "99"
)

// IndentJSON uniformly indent the json byte
func IndentJSON(b []byte) []byte {
	var out bytes.Buffer
	json.Indent(&out, b, "", " ")
	return bytes.TrimSpace(out.Bytes())
}
