package render

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONAPIError(t *testing.T) {
	jsonBlob :=
		[]byte(`{
					"errors": [
						{
							"status": "400",
							"title": "json api test",
							"detail": "json api fake test error",
							"meta": {
							"creator": "modware api"
						}
					}
				]
			}
	`)

	w := httptest.NewRecorder()
	detailErr := errors.New("json api fake test error")
	err := JSONAPIError(w, http.StatusBadRequest, detailErr, "json api test")
	if err != nil {
		t.Fatalf("unexpected rendering error %s\n", err)
	}
	if w.Header().Get("Content-Type") != "application/vnd.api+json" {
		t.Fatalf("no jsonapi response header %s", w.Header().Get("Content-Type"))
	}
	if w.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code %d\n", w.Code)
	}
	expJson := IndentJSON(jsonBlob)
	matchJson := IndentJSON(w.Body.Bytes())
	if bytes.Compare(expJson, matchJson) != 0 {
		t.Fatalf("expected \n%s jsonapi error response does not match with \n%s\n", string(expJson), string(matchJson))
	}
}

func IndentJSON(b []byte) []byte {
	var out bytes.Buffer
	json.Indent(&out, b, "", " ")
	return bytes.TrimSpace(out.Bytes())
}
