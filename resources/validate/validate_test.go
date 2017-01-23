package validate

import (
	"testing"

	jsapi "github.com/dictyBase/apihelpers/aphjsonapi"
	"github.com/dictyBase/go-middlewares/middlewares/query"
)

func getParamsWithSparseF() *query.Params {
	sf := &query.Fields{}
	sf.Append([]string{"name", "email"}...)
	return &query.Params{
		SparseFields: map[string]*query.Fields{
			"users": sf,
		},
	}
}

func getPrimaryResource() (*User, []jsapi.RelationShipLink) {
	u := &User{
		ID:    "32",
		Name:  "Tucker",
		Email: "tucker@jumbo.com",
	}
	allRels := jsapi.GetAllRelationships(u)
	return u, allRels
}

func TestPrimaryResource(t *testing.T) {
	p := getParamsWithSparseF()
	u, allRels := getPrimaryResource()
	err := FieldsParam(p, allRels, u)
	if err != nil {
		t.Errorf("error in validating sparse fieldset with primary resource type %s", err)
	}
	if p.SparseFields["users"].Relationship {
		t.Errorf("field should not be relationship resource type")
	}
}
