// Package modwaretest provides common constants and functions for unit testing
package modwaretest

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	APIHost    = "https://dictybase.org"
	PathPrefix = "1.0"
	PubId      = "99"
)

// IndentJSON uniformly indent the json byte
func IndentJSON(b []byte) []byte {
	var out bytes.Buffer
	json.Indent(&out, b, "", " ")
	return bytes.TrimSpace(out.Bytes())
}

// APIServer returns a server URL
func APIServer() string {
	return fmt.Sprintf("%s/%s", APIHost, PathPrefix)
}
