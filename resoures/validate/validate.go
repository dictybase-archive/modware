package validate

import (
	"net/http"

	"github.com/dictyBase/apihelpers/apherror"
	"github.com/dictyBase/go-middlewares/middlewares/query"
	jsvalidate "github.com/dictyBase/modware/models/jsonapi/validate"
)

//Param validates include, fields and filter query params
func Param(r *http.Request, data interface{}) (*query.Params, bool, error) {
	p, ok := r.Context().Value(query.ContextKeyQueryParams).(*query.Params)
	if !ok {
		return p, ok, nil
	}
	if err := IncludeParam(p, r, data); err != nil {
		return p, ok, err
	}
}

//IncludeParam validates the include param of JSONAPI specifications
func IncludeParam(p *query.Params, r *http.Request, data interface{}) error {
	if p.HasIncludes {
		err := jsvalidate.Relationships(data, p.Includes)
		if err != nil {
			return apherror.ErrIncludeParam.New(err.Error())
		}
	}
	return nil
}
