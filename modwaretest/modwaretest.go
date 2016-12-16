// Package modwaretest provides common constants and functions for unit testing
package modwaretest

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	APIHost    = "https://api.dictybase.org"
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

// MatchJSON compares actual and expected json
func MatchJSON(actual []byte, data interface{}) error {
	expected, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if bytes.Compare(IndentJSON(actual), IndentJSON(expected)) != 0 {
		return fmt.Errorf("actual %s and expected json %s are different", string(actual), string(expected))
	}
	return nil
}
