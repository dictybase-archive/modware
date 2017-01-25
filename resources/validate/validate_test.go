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

func TestIncludeParams(t *testing.T) {
	q := &query.Params{
		Includes:    []string{"roles"},
		HasIncludes: true,
	}
	u := getPrimaryResource()
	allRels := getRelationships(u)
	err := IncludeParam(q, allRels)
	if err != nil {
		t.Errorf("unexpected error in matching include params %s", err)
	}
}

func TestInvalidIncludeParam(t *testing.T) {
	q := &query.Params{
		Includes:    []string{"authors"},
		HasIncludes: true,
	}
	u := getPrimaryResource()
	allRels := getRelationships(u)
	err := IncludeParam(q, allRels)
	if !apherror.ErrIncludeParam.Contains(err) {
		t.Fatalf("actual error type does not match with expected type %s", err)
	}
}

func TestPrimaryResourceSpField(t *testing.T) {
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

func TestPrimaryResourceInvalidSpField(t *testing.T) {
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

func TestRelationshipResourceSpField(t *testing.T) {
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

func TestEmptyRelationshipResourceSpField(t *testing.T) {
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

func TestAllResourcesSpField(t *testing.T) {
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

func TestRelResourceInvalidSpField(t *testing.T) {
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

func TestRelResourceWrongIncludeSpField(t *testing.T) {
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

func TestRelResourceWrongSpField(t *testing.T) {
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

func TestFilterParams(t *testing.T) {
	q := &query.Params{
		Filters: map[string]string{
			"name": "bhola",
		},
		HasFilters: true,
	}
	u := getPrimaryResource()
	err := FilterParam(q, u)
	if err != nil {
		t.Errorf("unexpected error %s in validating filtering field", err)
	}
}

func TestInvalidFilterParams(t *testing.T) {
	q := &query.Params{
		Filters: map[string]string{
			"name":  "bhola",
			"email": "abc",
		},
		HasFilters: true,
	}
	u := getPrimaryResource()
	err := FilterParam(q, u)
	if !apherror.ErrFilterParam.Contains(err) {
		t.Fatalf("actual error type does not match with expected type %s", err)
	}
	m, rerr := regexp.MatchString("given filter param email", err.Error())
	if !m {
		t.Fatalf("actual error message %s does not match with expected message %s", rerr, err)
	}
}

func TestNoFilterParams(t *testing.T) {
	q := &query.Params{
		Filters: map[string]string{
			"role": "bhetki",
		},
		HasFilters: true,
	}
	err := FilterParam(q, &Role{})
	if !apherror.ErrFilterParam.Contains(err) {
		t.Fatalf("actual error type does not match with expected type %s", err)
	}
	m, rerr := regexp.MatchString("no filter attributes", err.Error())
	if !m {
		t.Fatalf("actual error message %s does not match with expected message %s", rerr, err)
	}
}
