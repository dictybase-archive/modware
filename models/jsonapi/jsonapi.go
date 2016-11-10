// Package jsonapi provides additional interfaces and wrapper functions for original
// jsonapi package("github.com/manyminds/api2go/jsonapi") to create and customize
// the self and related relationship links
package jsonapi

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/manyminds/api2go/jsonapi"
)

// Container type for having information about relationship links
type RelationShipLink struct {
	Name string
	// To override the default links, it will be appended to
	// the base url.
	SuffixFragment string
}

// Interface to implement for creating self relationship links
type MarshalSelfRelations interface {
	GetSelfLinksInfo() []RelationShipLink
}

// Interface to implement for creating related relationship links
type MarshalRelatedRelations interface {
	GetRelatedLinksInfo() []RelationShipLink
}

func MarshalToStructWrapper(data interface{}, ep jsonapi.ServerInformation) (*jsonapi.Document, error) {
	jst, err := jsonapi.MarshalToStruct(data, ep)
	if err != nil {
		return jst, err
	}
	if len(jst.Data.DataArray) > 0 { //array resource objects
		// picking first element both from the generated and given typed structures
		elem := jst.Data.DataArray[0]
		value := reflect.ValueOf(data).Index(0)
		// link for the array resource itself
		jst.Links = &jsonapi.Links{Self: generateMultiResourceLink(&elem, ep)}
		// link for individual resource
		slink := &jsonapi.Links{Self: generateSingleResourceLink(&elem, ep)}
		// Add it to every member
		for i, _ := range jst.Data.DataArray {
			jst.Data.DataArray[i].Links = slink
		}
		// Add relationships to every member
		relationships := generateRelationshipLinks(value, &elem, ep)
		if len(relationships) > 0 {
			for i, _ := range jst.Data.DataArray {
				jst.Data.DataArray[i].Relationships = relationships
			}
		}
	} else {
		jst.Links = &jsonapi.Links{Self: generateSingleResourceLink(jst.Data.DataObject, ep)}
		relationships := generateRelationshipLinks(data, jst.Data.DataObject, ep)
		if len(relationships) > 0 {
			jst.Data.DataObject.Relationships = relationships
		}
	}
	return jst, nil
}

func generateBaseLink(ep jsonapi.ServerInformation) string {
	return fmt.Sprintf(
		"%s/%s",
		strings.Trim(ep.GetBaseURL(), "/"),
		strings.Trim(ep.GetPrefix(), "/"),
	)
}

func generateSingleResourceLink(jdata *jsonapi.Data, ep jsonapi.ServerInformation) string {
	return fmt.Sprintf(
		"%s/%s/%s",
		generateBaseLink(ep),
		jdata.Type,
		jdata.ID,
	)
}

func generateMultiResourceLink(jdata *jsonapi.Data, ep jsonapi.ServerInformation) string {
	return fmt.Sprintf(
		"%s/%s",
		generateBaseLink(ep),
		jdata.Type,
	)
}

func generateRelationshipLinks(data interface{}, jdata *jsonapi.Data, ep jsonapi.ServerInformation) map[string]jsonapi.Relationship {
	relationships := make(map[string]jsonapi.Relationship)
	baselink := generateBaseLink(ep)
	self, ok := data.(MarshalSelfRelations)
	if ok {
		for _, rel := range self.GetSelfLinksInfo() {
			links := &jsonapi.Links{}
			if len(rel.SuffixFragment) > 0 {
				links.Self = fmt.Sprintf("%s/%s", baselink, strings.Trim(rel.SuffixFragment, "/"))
			} else {
				links.Self = fmt.Sprintf("%s/%s/%s/relationships/%s",
					baselink,
					jdata.Type,
					jdata.ID,
					rel.Name,
				)
			}
			relationships[rel.Name] = jsonapi.Relationship{Links: links}
		}
	}
	related, ok := data.(MarshalRelatedRelations)
	if ok {
		for _, rel := range related.GetRelatedLinksInfo() {
			var rlink string
			if len(rel.SuffixFragment) > 0 {
				rlink = fmt.Sprintf("%s/%s", baselink, strings.Trim(rel.SuffixFragment, "/"))
			} else {
				rlink = fmt.Sprintf("%s/%s/%s/%s",
					baselink,
					jdata.Type,
					jdata.ID, rel.Name,
				)
			}
			if _, ok := relationships[rel.Name]; ok {
				relationships[rel.Name].Links.Related = rlink
			} else {
				relationships[rel.Name] = jsonapi.Relationship{Links: &jsonapi.Links{Related: rlink}}
			}
		}
	}
	return relationships
}
