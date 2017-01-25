package validate

import "fmt"
import jsapi "github.com/dictyBase/apihelpers/aphjsonapi"

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
		jsapi.RelationShipLink{Name: "users", Type: "users"},
		jsapi.RelationShipLink{Name: "permissions", Type: "users"},
	}
}

func (r *Role) ValidateSelfLinks() error {
	return nil
}

func (r *Role) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{
			Name:           "users",
			SuffixFragment: fmt.Sprintf("%s/%s/%s", "roles", r.GetID(), "consumers"),
			Type:           "users",
		},
		jsapi.RelationShipLink{Name: "permissions", Type: "permissions"},
	}
}

func (r *Role) ValidateRelatedLinks() error {
	return nil
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
		jsapi.RelationShipLink{Name: "roles", Type: "roles"},
	}
}

func (u *User) ValidateSelfLinks() error {
	return nil
}

func (u *User) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "roles", Type: "roles"},
	}
}

func (u *User) GetAttributeFields(name string) []string {
	var attr []string
	if name == "roles" {
		attr = append(attr, jsapi.GetAttributeFields(&Role{})...)
	}
	return attr
}

func (u *User) ValidateRelatedLinks() error {
	return nil
}
