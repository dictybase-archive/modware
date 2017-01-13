package validate

import (
	"fmt"

	"github.com/dictyBase/apihelpers/aphcollection"
	"github.com/dictyBase/modware/models/jsonapi"
)

//Relationships matches if the relationships are implemented in the
//given JSONAPI implementing type
func Relationships(data interface{}, rels []string) error {
	self, ok := data.(jsonapi.MarshalSelfRelations)
	if ok {
		for _, rel := range self.GetSelfLinksInfo() {
			if !aphcollection.Contains(rels, rel.Name) {
				return fmt.Errorf("cannot match %s self relationship", rel.Name)
			}
		}
	}
	related, ok := data.(jsonapi.MarshalRelatedRelations)
	if ok {
		for _, rel := range related.GetRelatedLinksInfo() {
			if !aphcollection.Contains(rels, rel.Name) {
				return fmt.Errorf("cannot match %s self relationship", rel.Name)
			}
		}
	}
	return nil
}
