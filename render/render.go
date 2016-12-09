package render

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dictyBase/modware/resources"

	"github.com/gocraft/dbr"
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

func DatabaseError(w http.ResponseWriter, err error) {
	if err == dbr.ErrNotFound {
		jerr := JSONAPIError(w, http.StatusNotFound, err, resources.ErrNotExist.Error())
		if jerr != nil {
			http.Error(w, jerr.Error(), http.StatusInternalServerError)
		}
	} else { // possible database query error
		jerr := JSONAPIError(w, http.StatusInternalServerError, err, resources.ErrDatabaseQuery.Error())
		if jerr != nil {
			http.Error(w, jerr.Error(), http.StatusInternalServerError)
		}
	}
}

func StructMarshallingError(w http.ResponseWriter, err error) {
	jerr := JSONAPIError(w, http.StatusInternalServerError, err, resources.ErrStructMarshal.Error())
	if jerr != nil {
		http.Error(w, jerr.Error(), http.StatusInternalServerError)
	}
}

func JSONEncodingError(w http.ResponseWriter, err error) {
	jerr := JSONAPIError(w, http.StatusInternalServerError, err, resources.ErrJSONEncoding.Error())
	if jerr != nil {
		http.Error(w, jerr.Error(), http.StatusInternalServerError)
	}
}
