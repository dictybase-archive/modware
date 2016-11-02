// Package jsonapi provides additional interfaces and wrapper functions for original
// jsonapi package("github.com/manyminds/api2go/jsonapi") to create and customize
// the self and related relationship links
package jsonapi

import (
	"fmt"
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
	jst.Links = &jsonapi.Links{Self: generateSelfLink(jst, ep)}

	// create relationship links
	baselink := generateBaseLink(ep)
	relationships := make(map[string]jsonapi.Relationship)
	self, ok := data.(MarshalSelfRelations)
	if ok {
		for _, rel := range self.GetSelfLinksInfo() {
			links := &jsonapi.Links{}
			if len(rel.SuffixFragment) > 0 {
				links.Self = fmt.Sprintf("%s/%s", baselink, strings.Trim(rel.SuffixFragment, "/"))
			} else {
				links.Self = fmt.Sprintf("%s/%s/%s/relationships/%s",
					baselink,
					jst.Data.DataObject.Type,
					jst.Data.DataObject.ID,
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
					jst.Data.DataObject.Type,
					jst.Data.DataObject.ID, rel.Name,
				)
			}
			if _, ok := relationships[rel.Name]; ok {
				relationships[rel.Name].Links.Related = rlink
			} else {
				relationships[rel.Name] = jsonapi.Relationship{Links: &jsonapi.Links{Related: rlink}}
			}
		}
	}
	if len(relationships) > 0 {
		jst.Data.DataObject.Relationships = relationships
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

func generateSelfLink(jst *jsonapi.Document, ep jsonapi.ServerInformation) string {
	return fmt.Sprintf(
		"%s/%s/%s",
		generateBaseLink(ep),
		jst.Data.DataObject.Type,
		jst.Data.DataObject.ID,
	)
}
