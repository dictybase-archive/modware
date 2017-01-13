package render

import (
	"encoding/json"
	"net/http"

	"github.com/dictyBase/apihelpers/apherror"
	"github.com/dictyBase/go-middlewares/middlewares/pagination"

	jsapi "github.com/dictyBase/modware/models/jsonapi"
	"github.com/manyminds/api2go/jsonapi"
)

// Resource generates jsonapi response for resource
func Resource(data interface{}, srv jsonapi.ServerInformation, w http.ResponseWriter) {
	doc, err := jsapi.MarshalToStructWrapper(data, srv)
	if err != nil {
		apherror.JSONAPIError(w, apherror.ErrStructMarshal.New(err.Error()))
	}
	if err := JSONAPI(w, http.StatusOK, doc); err != nil {
		apherror.JSONAPIError(w, apherror.ErrJSONEncoding.New(err.Error()))
	}
}

// ResourceCollection generates jsonapi response for resource collection
func ResourceCollection(data interface{}, srv jsonapi.ServerInformation, w http.ResponseWriter, props *pagination.Props) {
	doc, err := jsapi.MarshalWithPagination(data, srv, props)
	if err != nil {
		apherror.JSONAPIError(w, apherror.ErrStructMarshal.New(err.Error()))
	}
	if err := JSONAPI(w, http.StatusOK, doc); err != nil {
		apherror.JSONAPIError(w, apherror.ErrJSONEncoding.New(err.Error()))
	}
}

func JSONAPI(w http.ResponseWriter, status int, data *jsonapi.Document) error {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(data)
}
