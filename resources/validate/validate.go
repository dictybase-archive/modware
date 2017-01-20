package validate

import (
	"fmt"
	"net/http"

	"github.com/dictyBase/apihelpers/aphcollection"
	"github.com/dictyBase/apihelpers/apherror"
	jsapi "github.com/dictyBase/apihelpers/aphjsonapi"
	jsvalidate "github.com/dictyBase/apihelpers/aphjsonapi/validate"
	"github.com/dictyBase/go-middlewares/middlewares/query"
)

//Param validates include, fields and filter query params
func Params(r *http.Request, data interface{}) (*query.Params, bool, error) {
	p, ok := r.Context().Value(query.ContextKeyQueryParams).(*query.Params)
	if !ok {
		return p, ok, nil
	}
	allRels := jsvalidate.GetAllRelationships(data)
	if p.HasIncludes {
		if len(allRels) == 0 {
			return p, ok, apherror.ErrIncludeParam.New("No relationship defined")
		}
		if err := IncludeParam(p, allRels); err != nil {
			return p, ok, err
		}
	}
}

//IncludeParam validates the include param of JSONAPI specifications
func IncludeParam(p *query.Params, rs []jsapi.RelationShipLink) error {
	err := jsvalidate.HasRelationships(p.Includes, rs)
	if err != nil {
		return apherror.ErrIncludeParam.New(err.Error())
	}
}

//FieldsParam validates the fields(sparse fieldsets) param of JSONAPI specifications
func FieldsParam(p *query.Params, rs []jsapi.RelationShipLink, data interface{}) error {
	tn := jsapi.GetTypeName(data)
	attrs := jsapi.AttributeNames(data)
	for ftype, f := range p.SparseFields {
		if name == tn {
			for _, field := range f.GetAll() {
				if !aphcollection.Contains(attrs, field) {
					return apherror.ErrQueryParam.New(
						fmt.Sprintf("field value %s is not an attribute in resource %s", field, ftype),
					)
				}
			}
			continue
		}
		if len(rs) == 0 {
			return apherror.ErrQueryParam.New(
				fmt.Sprintf("no relationship defined, cannot process resource %s in sparse field", ftype),
			)
		}
		rname, err := jsvalidate.RelationshipResourceType(ftype, rs)
		if err != nil {
			return apherror.ErrQueryParam.New(err.Error())
		}
		// Now the relationship name should be in include param
		if !aphcollection.Contains(p.Includes, rname) {
			return apherror.ErrQueryParam.New("resource %s of sparse field %s is not in include param", rname, ftype)
		}
		// Check the attribute fields of relationship resource
		atype, ok := data.(jsapi.RelationshipAttributes)
		if !ok {
			apherror.ErrQueryParam.New(
				fmt.Sprintf("RelationshipAttribute interface not defined for relationship resource %s", ftype),
			)
		}
		rattrs := atype.GetAttributeFields(ftype)
		for _, field := range f.GetAll() {
			if !aphcollection.Contains(rattrs, field) {
				return apherror.ErrQueryParam.New(
					fmt.Sprintf("field value %s is not an attribute in relationship resource %s", field, ftype),
				)
			}
		}
		// tag the relationship resource type in sparse fieldset
		f.Relationship = true
	}
	return nil
}
