package publication

import (
	"fmt"

	jsapi "github.com/dictyBase/apihelpers/aphjsonapi"
	"github.com/manyminds/api2go/jsonapi"
)

type Author struct {
	ID            string         `json:"-" db:"pubauthor_id"`
	LastName      string         `json:"last_name,omitempty" db:"surname"`
	GivenNames    string         `json:"given_names,omitempty" db:"givennames"`
	Rank          int            `json:"rank,omitempty" db:"rank"`
	Publications  []*Publication `json:"-" db:"-"`
	PublictionIDs []string       `json:"-" db:"-"`
}

// GetID satisfies jsonapi.MarshalIdentifier interface
func (a *Author) GetID() string {
	return a.ID
}

// SetID satisfies jsonapi.UnMarshalIdentifier interface
func (a *Author) SetID(id string) error {
	a.ID = id
	return nil
}

// GetRelatedLinksInfo satisfies jsapi.MarshalRelatedRelations interface
func (a *Author) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "publications"},
	}
}

// GetReferencedStructs satisfies jsonapi.MarshalIncludedRelations interface
func (a *Author) GetReferencedStructs() (result []jsonapi.MarshalIdentifier) {
	if len(a.Publications) > 0 {
		for _, pub := range a.Publications {
			result = append(result, pub)
		}
	}
	return result
}

// GetReferences satisfies jsonapi.MarshalReferences interface
func (a *Author) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		jsonapi.Reference{Type: "publications", Name: "publications"},
	}
}

// GetReferencedIDs satisfies jsonapi.MarshalLinkedRelations interface
func (a *Author) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID
	for _, p := range a.Publications {
		result = append(result, jsonapi.ReferenceID{Type: "publications", ID: p.ID, Name: "publications"})
	}
	return result
}

// SetToManyReferenceIDs satisfies jsonapi.UnmarshalToManyRelations interface
func (a *Author) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "publications" {
		a.PublictionIDs = IDs
		return nil
	}
	return fmt.Errorf("%s No such has many relationships", name)
}

// GetAttributeFields implements jsapi.RelationshipAttribute interface
func (a *Author) GetAttributeFields(name string) []string {
	var attr []string
	if name == "publications" {
		attr = append(attr, jsapi.GetAttributeFields(&Publication{})...)
	}
	return attr
}

type Publication struct {
	ID        string    `json:"-"`
	Doi       string    `json:"doi"`
	Title     string    `json:"title,omitempty" filter:"-"`
	Abstract  string    `json:"abstract,omitempty"`
	Journal   string    `json:"journal,omitempty"`
	Year      string    `json:"year,omitempty"`
	Volume    string    `json:"volume,omitempty"`
	Pages     string    `json:"pages,omitempty"`
	Month     string    `json:"month,omitempty"`
	Issn      string    `json:"issn,omitempty"`
	Issue     string    `json:"issue,omitempty"`
	Source    string    `json:"source,omitempty"`
	PubType   string    `json:"pub_type,omitempty"`
	Status    string    `json:"status,omitempty"`
	Authors   []*Author `json:"-"`
	AuthorIDs []string  `json:"-"`
}

// GetID satisfies jsonapi.MarshalIdentifier interface
func (pub *Publication) GetID() string {
	return pub.ID
}

// GetRelatedLinksInfo satisfies jsapi.MarshalRelatedRelations interface
func (pub *Publication) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "authors"},
	}
}

// SetID satisfies jsonapi.UnMarshalIdentifier interface
func (pub *Publication) SetID(id string) error {
	pub.ID = id
	return nil
}

// GetReferencedStructs satisfies jsonapi.MarshalIncludedRelations interface
func (pub *Publication) GetReferencedStructs() (result []jsonapi.MarshalIdentifier) {
	if len(pub.Authors) > 0 {
		for _, a := range pub.Authors {
			result = append(result, a)
		}
	}
	return result
}

// GetReferences satisfies jsonapi.MarshalReferences interface
func (pub *Publication) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		jsonapi.Reference{Type: "authors", Name: "authors"},
	}
}

// GetReferencedIDs satisfies jsonapi.MarshalLinkedRelations interface
func (pub *Publication) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID
	for _, a := range pub.Authors {
		result = append(result, jsonapi.ReferenceID{Type: "authors", ID: a.ID, Name: "authors"})
	}
	return result
}

// SetToManyReferenceIDs satisfies jsonapi.UnmarshalToManyRelations interface
func (pub *Publication) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "authors" {
		pub.AuthorIDs = IDs
		return nil
	}
	return fmt.Errorf("%s No such has many relationships", name)
}

// GetAttributeFields implements jsapi.RelationshipAttribute interface
func (pub *Publication) GetAttributeFields(name string) []string {
	var attr []string
	if name == "authors" {
		attr = append(attr, jsapi.GetAttributeFields(&Author{})...)
	}
	return attr
}

// GetMap implements jsapi.AttributeToDbRowMapper interface
func (pub *Publication) GetMap() map[string]string {
	return map[string]string{"title": "pub.title"}
}
