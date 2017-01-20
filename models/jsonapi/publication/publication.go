package publication

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"

	jsapi "github.com/dictyBase/apihelpers/aphjsonapi"
	"github.com/manyminds/api2go/jsonapi"
)

var validate *validator.Validate

type Author struct {
	ID            string         `json:"-" db:"pubauthor_id"`
	LastName      string         `json:"last_name" db:"surname"`
	GivenNames    string         `json:"given_names" db:"givennames"`
	Rank          int            `json:"rank" db:"rank"`
	Publications  []*Publication `json:"-" db:"-"`
	PublictionIDs []string       `json:"-" db:"-"`
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
			Type: "publications",
		},
	}
}

func (a *Author) ValidateRelatedLinks() error {
	validate = validator.New()
	for _, v := range a.GetRelatedLinksInfo() {
		if err := validate.Struct(v); err != nil {
			return err
		}
	}
	return nil
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
	Journal   string    `json:"journal"`
	Year      string    `json:"year"`
	Volume    string    `json:"volume"`
	Pages     string    `json:"pages"`
	Month     string    `json:"month"`
	Issn      string    `json:"issn"`
	Issue     string    `json:"issue"`
	Source    string    `json:"source"`
	PubType   string    `json:"pub_type"`
	Status    string    `json:"status"`
	Authors   []*Author `json:"-"`
	AuthorIDs []string  `json:"-"`
}

func (pub *Publication) GetID() string {
	return pub.ID
}

func (pub *Publication) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "authors", Type: "authors"},
	}
}

func (pub *Publication) ValidateRelatedLinks() error {
	validate = validator.New()
	for _, v := range pub.GetRelatedLinksInfo() {
		if err := validate.Struct(v); err != nil {
			return err
		}
	}
	return nil
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

func (pub *Publication) GetAttributeFields(name string) []string {
	var attr []string
	if name == "authors" {
		attr = jsapi.GetAttributeFields(&Author{})
	}
	return attr
}
