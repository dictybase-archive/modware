package render

import (
	"encoding/json"
	"net/http"

	"github.com/manyminds/api2go/jsonapi"
)

func renderJSONAPI(w http.ResponseWriter, status int, data *jsonapi.Document) error {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
