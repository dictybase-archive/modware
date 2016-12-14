package render

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dictyBase/go-middlewares/middlewares/pagination"
	"github.com/dictyBase/modware/resources"

	jsapi "github.com/dictyBase/modware/models/jsonapi"
	"github.com/gocraft/dbr"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/jsonapi"
)

// Resource generates jsonapi response for resource
func Resource(data interface{}, srv jsonapi.ServerInformation, w http.ResponseWriter) {
	doc, err := jsapi.MarshalToStructWrapper(data, srv)
	if err != nil {
		StructMarshallingError(w, err)
	}
	if err := JSONAPI(w, http.StatusOK, doc); err != nil {
		JSONEncodingError(w, err)
	}
}

// ResourceCollection generates jsonapi response for resource collection
func ResourceCollection(data interface{}, srv jsonapi.ServerInformation, w http.ResponseWriter, props *pagination.Props) {
	doc, err := jsapi.MarshalWithPagination(data, srv, props)
	if err != nil {
		StructMarshallingError(w, err)
	}
	if err := JSONAPI(w, http.StatusOK, doc); err != nil {
		JSONEncodingError(w, err)
	}
}

func JSONAPI(w http.ResponseWriter, status int, data *jsonapi.Document) error {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(data)
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
