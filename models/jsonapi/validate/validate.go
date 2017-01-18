package validate

import (
	"fmt"

	"github.com/dictyBase/apihelpers/aphcollection"
	"github.com/dictyBase/modware/models/jsonapi"
)

//Relationships matches if the relationships rels are implemented in the
//given JSONAPI implementing data
func Relationships(data interface{}, rels []string) error {
	hasSelf, hasRelated := false
	self, ok := data.(jsonapi.MarshalSelfRelations)
	if ok {
		for _, rel := range self.GetSelfLinksInfo() {
			if !aphcollection.Contains(rels, rel.Name) {
				return fmt.Errorf("cannot match %s to self relationship", rel.Name)
			}
		}
		hasSelf = true
	}
	related, ok := data.(jsonapi.MarshalRelatedRelations)
	if ok {
		for _, rel := range related.GetRelatedLinksInfo() {
			if !aphcollection.Contains(rels, rel.Name) {
				return fmt.Errorf("cannot match %s to self relationship", rel.Name)
			}
		}
		hasRelated = true
	}
	if !hasSelf && !hasRelated {
		return fmt.Errorf("no self or related relationship defined")
	}
	return nil
}
