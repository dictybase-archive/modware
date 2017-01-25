package validate

import (
	"regexp"
	"testing"

	"github.com/dictyBase/apihelpers/apherror"
	jsapi "github.com/dictyBase/apihelpers/aphjsonapi"
	"github.com/dictyBase/go-middlewares/middlewares/query"
)

func getParamsWithRelationshipSparseF() *query.Params {
	sf := &query.Fields{}
	sf.Append([]string{"role", "description"}...)
	return &query.Params{
		SparseFields: map[string]*query.Fields{
			"roles": sf,
		},
		Includes:        []string{"roles"},
		HasSparseFields: true,
		HasIncludes:     true,
	}
}

func getWrongFieldValWithRelSparseF() *query.Params {
	sf := &query.Fields{}
	sf.Append([]string{"name", "email"}...)
	rf := &query.Fields{}
	rf.Append([]string{"role", "authors"}...)
	return &query.Params{
		SparseFields: map[string]*query.Fields{
			"roles": rf,
			"users": sf,
		},
		Includes:        []string{"roles"},
		HasSparseFields: true,
		HasIncludes:     true,
	}
}

func getWrongResourceWithRelSparseF() *query.Params {
	sf := &query.Fields{}
	sf.Append([]string{"name", "email"}...)
	rf := &query.Fields{}
	rf.Append([]string{"role", "description"}...)
	return &query.Params{
		SparseFields: map[string]*query.Fields{
			"authors": rf,
			"users":   sf,
		},
		Includes:        []string{"authors"},
		HasSparseFields: true,
		HasIncludes:     true,
	}
}

func getAllParamsWithSparseF() *query.Params {
	sf := &query.Fields{}
	sf.Append([]string{"name", "email"}...)
	rf := &query.Fields{}
	rf.Append([]string{"role", "description"}...)
	return &query.Params{
		SparseFields: map[string]*query.Fields{
			"roles": rf,
			"users": sf,
		},
		Includes:        []string{"roles"},
		HasSparseFields: true,
		HasIncludes:     true,
	}
}

func getParamsWithSparseF() *query.Params {
	sf := &query.Fields{}
	sf.Append([]string{"name", "email"}...)
	return &query.Params{
		SparseFields: map[string]*query.Fields{
			"users": sf,
		},
		HasSparseFields: true,
	}
}

func getWrongParamsWithSparseF() *query.Params {
	sf := &query.Fields{}
	sf.Append([]string{"name", "address"}...)
	return &query.Params{
		SparseFields: map[string]*query.Fields{
			"users": sf,
		},
		HasSparseFields: true,
	}
}

func getPrimaryResource() *User {
	return &User{
		ID:    "32",
		Name:  "Tucker",
		Email: "tucker@jumbo.com",
	}
}

func getRelationships(data interface{}) []jsapi.RelationShipLink {
	return jsapi.GetAllRelationships(data)
}

func TestPrimaryResource(t *testing.T) {
	p := getParamsWithSparseF()
	u := getPrimaryResource()
	allRels := getRelationships(u)
	err := FieldsParam(p, allRels, u)
	if err != nil {
		t.Errorf("error in validating sparse fieldset with primary resource type %s", err)
	}
	if p.SparseFields["users"].Relationship {
		t.Errorf("field should not be relationship resource type")
	}
}

func TestPrimaryResourceInvalidField(t *testing.T) {
	p := getWrongParamsWithSparseF()
	u := getPrimaryResource()
	allRels := getRelationships(u)
	err := FieldsParam(p, allRels, u)
	if !apherror.ErrSparseFieldSets.Contains(err) {
		t.Fatalf("actual error type does not match with expected type %s", err)
	}
	m, rerr := regexp.MatchString("address", err.Error())
	if !m {
		t.Fatalf("actual error message %s does not match with expected message %s", rerr, err)
	}
}

func TestRelationshipResource(t *testing.T) {
	p := getParamsWithRelationshipSparseF()
	u := getPrimaryResource()
	allRels := getRelationships(u)
	err := FieldsParam(p, allRels, u)
	if err != nil {
		t.Errorf("error in validating sparse fieldset with relationship resource type %s", err)
	}
	if !p.SparseFields["roles"].Relationship {
		t.Errorf("field %s should be relationship resource type", "roles")
	}
}

func TestEmptyRelationshipResource(t *testing.T) {
	p := getParamsWithRelationshipSparseF()
	u := getPrimaryResource()
	var allRels []jsapi.RelationShipLink
	err := FieldsParam(p, allRels, u)
	if !apherror.ErrSparseFieldSets.Contains(err) {
		t.Fatalf("actual error type does not match with expected type %s", err)
	}
	m, rerr := regexp.MatchString("no relationship defined.*roles", err.Error())
	if !m {
		t.Fatalf("actual error message %s does not match with expected message %s", rerr, err)
	}
}

func TestAllResources(t *testing.T) {
	p := getAllParamsWithSparseF()
	u := getPrimaryResource()
	allRels := getRelationships(u)
	err := FieldsParam(p, allRels, u)
	if err != nil {
		t.Errorf("error in validating sparse fieldset with all resource type %s", err)
	}
	if !p.SparseFields["roles"].Relationship {
		t.Errorf("field %s should be relationship resource type", "roles")
	}
	if p.SparseFields["users"].Relationship {
		t.Errorf("field %s should not be relationship resource type", "users")
	}
}

func TestRelResourceInvalidField(t *testing.T) {
	p := getWrongResourceWithRelSparseF()
	u := getPrimaryResource()
	allRels := getRelationships(u)
	err := FieldsParam(p, allRels, u)
	if !apherror.ErrSparseFieldSets.Contains(err) {
		t.Fatalf("actual error type does not match with expected type %s", err)
	}
	m, rerr := regexp.MatchString("authors resource type", err.Error())
	if !m {
		t.Fatalf("actual error message %s does not match with expected message %s", rerr, err)
	}
}

func TestRelResourceWrongIncludeField(t *testing.T) {
	p := getAllParamsWithSparseF()
	p.Includes = []string{"authors"}
	u := getPrimaryResource()
	allRels := getRelationships(u)
	err := FieldsParam(p, allRels, u)
	if !apherror.ErrSparseFieldSets.Contains(err) {
		t.Fatalf("actual error type does not match with expected type %s", err)
	}
	m, rerr := regexp.MatchString("include param", err.Error())
	if !m {
		t.Fatalf("actual error message %s does not match with expected message %s", rerr, err)
	}
}

func TestRelResourceWrongField(t *testing.T) {
	p := getWrongFieldValWithRelSparseF()
	u := getPrimaryResource()
	allRels := getRelationships(u)
	err := FieldsParam(p, allRels, u)
	if !apherror.ErrSparseFieldSets.Contains(err) {
		t.Fatalf("actual error type does not match with expected type %s", err)
	}
	m, rerr := regexp.MatchString("field value authors", err.Error())
	if !m {
		t.Fatalf("actual error message %s does not match with expected message %s", rerr, err)
	}
}
