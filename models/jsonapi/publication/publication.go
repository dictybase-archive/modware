package publication

import jsapi "github.com/dictybase/modware/models/jsonapi"

type Author struct {
	ID        string `json:"-"`
	LastName  string `json:"last_name"`
	GivenName string `json:"given_name"`
	Rank      int    `json:"rank"`
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
