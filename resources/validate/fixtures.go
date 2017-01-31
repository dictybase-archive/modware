package validate

import (
	"fmt"

	jsapi "github.com/dictyBase/apihelpers/aphjsonapi"
	"github.com/manyminds/api2go/jsonapi"
)

type Permission struct {
	ID          string `json:"-"`
	Permission  string `json:"permission"`
	Description string `json:"description"`
}

func (p Permission) GetID() string {
	return p.ID
}

type Role struct {
	ID          string        `json:"-"`
	Role        string        `json:"role"`
	Description string        `json:"description"`
	Permissions []*Permission `json:"-"`
	Users       []*User       `json:"-"`
}

func (r *Role) GetID() string {
	return r.ID
}

func (r *Role) GetSelfLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "users"},
		jsapi.RelationShipLink{Name: "permissions"},
	}
}

func (r *Role) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{
			Name:           "users",
			SuffixFragment: fmt.Sprintf("%s/%s/%s", "roles", r.GetID(), "consumers"),
		},
		jsapi.RelationShipLink{Name: "permissions"},
	}
}

type User struct {
	ID    string  `json:"-"`
	Name  string  `json:"name" filter:"-"`
	Email string  `json:"email"`
	Roles []*Role `json:"-"`
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetSelfLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "roles"},
	}
}

func (u *User) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "roles"},
	}
}

func (u *User) GetAttributeFields(name string) []string {
	var attr []string
	if name == "roles" {
		attr = append(attr, jsapi.GetAttributeFields(&Role{})...)
	}
	return attr
}

// GetReferences satisfies jsonapi.MarshalReferences interface
func (u *User) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		jsonapi.Reference{Type: "roles", Name: "roles"},
	}
}

// GetReferencedStructs satisfies jsonapi.MarshalIncludedRelations interface
func (u *User) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	var result []jsonapi.MarshalIdentifier
	for _, r := range u.Roles {
		result = append(result, r)
	}
	return result
}

// GetReferencedIDs satisfies jsonapi.MarshalLinkedRelations interface
func (u *User) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID
	for _, r := range u.Roles {
		result = append(result, jsonapi.ReferenceID{Type: "roles", ID: r.ID, Name: "roles"})
	}
	return result
}
