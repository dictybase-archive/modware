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

// APIErrorOptions specifies the optional parameters for displaying jsonapi formatted http errors
type APIErrorOptions struct {
	Title     string
	Detail    string
	Pointer   string
	Parameter string
}

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

func JSONAPIError(w http.ResponseWriter, status int, opt *APIErrorOptions) error {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)
	jsnErr := api2go.Error{
		Status: strconv.Itoa(status),
		Title:  opt.Title,
		Detail: opt.Detail,
		Meta: map[string]interface{}{
			"creator": "modware api",
		},
	}
	if len(opt.Pointer) > 0 {
		jsnErr.Source.Pointer = opt.Pointer
	}
	if len(opt.Parameter) > 0 {
		jsnErr.Source.Parameter = opt.Parameter
	}
	return json.NewEncoder(w).Encode(api2go.HTTPError{Errors: []api2go.Error{jsnErr}})
}

func DatabaseError(w http.ResponseWriter, err error) {
	if err == dbr.ErrNotFound {
		jerr := JSONAPIError(
			w,
			http.StatusNotFound,
			&APIErrorOptions{
				Detail: err.Error(),
				Title:  resources.ErrNotExist.Error(),
			},
		)
		if jerr != nil {
			http.Error(w, jerr.Error(), http.StatusInternalServerError)
		}
	} else { // possible database query error
		jerr := JSONAPIError(
			w,
			http.StatusInternalServerError,
			&APIErrorOptions{
				Detail: err.Error(),
				Title:  resources.ErrDatabaseQuery.Error(),
			},
		)
		if jerr != nil {
			http.Error(w, jerr.Error(), http.StatusInternalServerError)
		}
	}
}

func StructMarshallingError(w http.ResponseWriter, err error) {
	jerr := JSONAPIError(
		w,
		http.StatusInternalServerError,
		&APIErrorOptions{
			Detail: err.Error(),
			Title:  resources.ErrStructMarshal.Error(),
		},
	)
	if jerr != nil {
		http.Error(w, jerr.Error(), http.StatusInternalServerError)
	}
}

func JSONEncodingError(w http.ResponseWriter, err error) {
	jerr := JSONAPIError(
		w,
		http.StatusInternalServerError,
		&APIErrorOptions{
			Detail: err.Error(),
			Title:  resources.ErrJSONEncoding.Error(),
		},
	)
	if jerr != nil {
		http.Error(w, jerr.Error(), http.StatusInternalServerError)
	}
}

func IncludeParamError(w http.ResponseWriter, detail string) {
	jerr := JSONAPIError(
		w,
		http.StatusBadRequest,
		&APIErrorOptions{
			Detail:    detail,
			Title:     resources.ErrIncludeParam.Error(),
			Parameter: "include",
		},
	)
	if jerr != nil {
		http.Error(w, jerr.Error(), http.StatusInternalServerError)
	}
}
