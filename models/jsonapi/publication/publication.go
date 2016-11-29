package publication

import (
	"fmt"

	jsapi "github.com/dictybase/modware/models/jsonapi"
	"github.com/manyminds/api2go/jsonapi"
)

type Author struct {
	ID            string         `json:"-"`
	LastName      string         `json:"last_name"`
	GivenName     string         `json:"given_name"`
	Rank          int            `json:"rank"`
	Publications  []*Publication `json:"-"`
	PublictionIDs []string       `json:"-"`
}

func (a *Author) GetID() string {
	return a.ID
}

func (a *Author) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *Author) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{
			Name: "publications",
		},
	}
}

func (a *Author) GetSelfLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{
			Name: "publications",
		},
	}
}

func (a *Author) GetReferencedStructs() (result []jsonapi.MarshalIdentifier) {
	if len(a.Publications) > 0 {
		for _, pub := range a.Publications {
			result = append(result, pub)
		}
	}
	return result
}

func (a *Author) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "publications" {
		a.PublictionIDs = IDs
		return nil
	}
	return fmt.Errorf("%s No such has many relationships", name)
}

type Publication struct {
	ID        string    `json:"-"`
	Doi       string    `json:"doi"`
	Title     string    `json:"title"`
	Abstract  string    `json:"abstract"`
	Journal   string    `json:"abstract"`
	Year      string    `json:"year"`
	Volume    string    `json:"volume"`
	Pages     string    `json:"pages"`
	Month     string    `json:"month"`
	Issn      string    `json:"issn"`
	Source    string    `json:"source"`
	PubType   string    `json:"pub_type"`
	Authors   []*Author `json:"-"`
	AuthorIDs []string  `json:"-"`
}

func (pub *Publication) GetID() string {
	return pub.ID
}

func (pub *Publication) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "authors"},
	}
}

func (pub *Publication) GetSelfLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "authors"},
	}
}

func (pub *Publication) SetID(id string) error {
	pub.ID = id
	return nil
}

func (pub *Publication) GetReferencedStructs() (result []jsonapi.MarshalIdentifier) {
	if len(pub.Authors) > 0 {
		for _, a := range pub.Authors {
			result = append(result, a)
		}
	}
	return result
}

func (pub *Publication) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "authors" {
		pub.AuthorIDs = IDs
		return nil
	}
	return fmt.Errorf("%s No such has many relationships", name)
}