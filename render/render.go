package render

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/jsonapi"
)

func JSONAPI(w http.ResponseWriter, status int, data *jsonapi.Document) error {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func JSONAPIError(w http.ResponseWriter, status int, err error, msg string) error {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)
	jsnErr := api2go.Error{
		Status: strconv.Itoa(status),
		Title:  msg,
		Detail: err.Error(),
		Meta: map[string]interface{}{
			"creator": "modware api",
		},
	}
	return json.NewEncoder(w).Encode(api2go.HTTPError{Errors: []api2go.Error{jsnErr}})
}
