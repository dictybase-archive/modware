package auth

import (
	"fmt"

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

func (p Permission) SetID(id string) error {
	p.ID = id
	return nil
}

type Role struct {
	ID            string       `json:"-"`
	Role          string       `json:"role"`
	Description   string       `json:"description"`
	Permissions   []Permission `json:"-"`
	Users         []User       `json:"-"`
	IsInclude     bool         `json:"-"`
	UserIDs       []string     `json:"-"`
	PermissionIDs []string     `json:"-"`
}

func (r Role) GetID() string {
	return r.ID
}

func (r Role) SetID(id string) error {
	r.ID = id
	return nil
}

func (r Role) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		jsonapi.Reference{Type: "permissions", Name: "permissions", IsNotLoaded: r.IsInclude},
		jsonapi.Reference{Type: "users", Name: "users", IsNotLoaded: r.IsInclude},
	}
}

func (r Role) GetReferencedStructs() (result []jsonapi.MarshalIdentifier) {
	if r.IsInclude {
		for _, u := range r.Users {
			result = append(result, u)
		}
		for _, p := range r.Permissions {
			result = append(result, p)
		}
	}
	return result
}

func (r Role) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "permissions" {
		r.PermissionIDs = IDs
		return nil
	}
	if name == "users" {
		r.UserIDs = IDs
		return nil
	}
	return fmt.Errorf("%s No such has many relationships", name)
}

type User struct {
	ID            string   `json:"-"`
	Organization  string   `"json:"organization`
	GroupName     string   `json:"group_name"`
	FirstAddress  string   `json:"first_address"`
	SecondAddress string   `json:"second_address"`
	City          string   `json:"city"`
	State         string   `json:"state"`
	Zipcode       string   `json:"zipcode"`
	Country       string   `json:"country"`
	Phone         string   `json:"phone"`
	Roles         []Role   `json:"-"`
	RoleIDs       []string `json:"-"`
	IsInclude     bool     `json:"-"`
}

func (u User) GetID() string {
	return u.ID
}

func (u User) SetID(id string) error {
	u.ID = id
	return nil
}

func (u User) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		jsonapi.Reference{Type: "roles", Name: "roles", IsNotLoaded: u.IsInclude},
	}
}

func (u User) GetReferencedStructs() (result []jsonapi.MarshalIdentifier) {
	if u.IsInclude {
		for _, r := range u.Roles {
			result = append(result, r)
		}
	}
	return result
}

func (u User) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "roles" {
		u.RoleIDs = IDs
		return nil
	}
	return fmt.Errorf("%s No such has many relationships", name)
}
