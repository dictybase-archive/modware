package jsonapi

import (
	"fmt"
	"reflect"
	"testing"

	jsapi "github.com/manyminds/api2go/jsonapi"
)

const (
	base   = "https://api.dictybase.org"
	prefix = "1.0"
)

func GetApiServerInfo() *ApiServerInfo {
	return &ApiServerInfo{base, prefix}
}

func TestLinkGeneration(t *testing.T) {
	srvinfo := GetApiServerInfo()
	exlink := fmt.Sprintf("%s/%s", base, prefix)
	bslink := generateBaseLink(srvinfo)
	if bslink != exlink {
		t.Fatalf("expected base link %s does not match with generated link %s", exlink, bslink)
	}

	id := "14"
	jtype := "books"
	exrlink := fmt.Sprintf("%s/%s/%s", exlink, jtype, id)
	rlink := generateSingleResourceLink(&jsapi.Data{Type: jtype, ID: id}, srvinfo)
	if exrlink != rlink {
		t.Fatalf("expected single resource link %s does not match generated link %s", exrlink, rlink)
	}
}

func TestSelfLink(t *testing.T) {
	perm := &Permission{
		ID:          "10",
		Permission:  "gene curation",
		Description: "Authority to edit gene information",
	}
	srvinfo := GetApiServerInfo()
	pstruct, err := MarshalToStructWrapper(perm, srvinfo)
	if err != nil {
		t.Errorf("error in marshaling to structure %s\n", err)
	}
	exstruct := &jsapi.Document{
		Links: &jsapi.Links{Self: fmt.Sprintf("%s/%s/%s/%s", srvinfo.GetBaseURL(), srvinfo.GetPrefix(), "permissions", "10")},
		Data: &jsapi.DataContainer{
			DataObject: &jsapi.Data{
				Type:       "permissions",
				ID:         "10",
				Attributes: []byte(`{"permission":"gene curation","description":"Authority to edit gene information"}`),
			},
		},
	}
	if !reflect.DeepEqual(pstruct, exstruct) {
		t.Fatal("expected and generated jsonapi structure did not match")
	}
}

func TestRelationshipLink(t *testing.T) {

}
