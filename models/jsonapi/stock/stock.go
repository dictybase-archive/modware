package stock

import (
	"fmt"

	jsapi "github.com/dictybase/modware/models/jsonapi"
	"github.com/manyminds/api2go/jsonapi"
)

type StockOrder struct {
	ID              string `json:"-"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Courier         string `json:"courier"`
	CourierAccount  string `json:"courier_account"`
	Comments        string `"json:"comments"`
	Payment         string `"json"payment"`
	PurchaseOrderId int64  `"json:"purchase_order_num"`
	Status          string `"json:"status"`
	IsInclude       bool   `json:"-"`
}

func (so *StockOrder) GetID() string {
	return so.ID
}

func (so *StockOrder) SetID(id string) error {
	so.ID = id
	return nil
}

func (so *StockOrder) GetName() string {
	return "orders"
}

func (so *StockOrder) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		jsonapi.Reference{Type: "stocks", Name: "stocks", IsNotLoaded: so.IsInclude},
		jsonapi.Reference{Type: "users", Name: "consumers", IsNotLoaded: so.IsInclude},
		jsonapi.Reference{Type: "users", Name: "payer", IsNotLoaded: so.IsInclude},
		jsonapi.Reference{Type: "users", Name: "purchaser", IsNotLoaded: so.IsInclude},
	}
}

type Phenotype struct {
	ID          string       `json:"-"`
	Name        string       `json:"name"`
	Observation string       `json:"observation"`
	Attribute   string       `json:"phen_attribute"`
	Value       string       `json:"value"`
	Cvalue      string       `json:"cvalue"`
	Evidence    string       `json:"evidence"`
	Props       []*PhenoProp `json:"-"`
	PropIDs     []string     `json:"-"`
}

func (ph *Phenotype) GetID() string {
	return ph.ID
}

func (ph *Phenotype) SetID(id string) error {
	ph.ID = id
	return nil
}

func (ph *Phenotype) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "properties"},
	}
}

func (ph *Phenotype) GetReferencedStructs() (result []jsonapi.MarshalIdentifier) {
	if len(ph.Props) > 0 {
		for _, p := range ph.Props {
			result = append(result, p)
		}
	}
	return result
}

func (ph *Phenotype) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "properties" {
		ph.PropIDs = IDs
		return nil
	}
	return fmt.Errorf("%s No such has many relationships", name)
}

type PhenoProp struct {
	ID     string `json:"-"`
	Value  string `json:"value"`
	Cv     string `json:"-"`
	TermID string `json:"-"`
}

func (prop *PhenoProp) GetID() string {
	return prop.ID
}

func (prop *PhenoProp) SetID(id string) error {
	prop.ID = id
	return nil
}

func (prop *PhenoProp) GetName() string {
	return "chadoprops"
}

func (prop *PhenoProp) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{
			Name:           "proptype",
			SuffixFragment: fmt.Sprintf("cvs/%s/cvterms/%s", prop.Cv, prop.TermID),
		},
	}
}

type Genotype struct {
	ID          string       `json:"-"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Props       []*PhenoProp `json:"-"`
	PropIDs     []string     `json:"-"`
}

func (g *Genotype) GetID() string {
	return g.ID
}

func (g *Genotype) SetID(id string) error {
	g.ID = id
	return nil
}

func (g *Genotype) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{Name: "properties"},
	}
}

func (g *Genotype) GetReferencedStructs() (result []jsonapi.MarshalIdentifier) {
	if len(g.Props) > 0 {
		for _, p := range g.Props {
			result = append(result, p)
		}
	}
	return result
}

func (g *Genotype) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "properties" {
		g.PropIDs = IDs
		return nil
	}
	return fmt.Errorf("%s No such has many relationships", name)
}

type GenoProp struct {
	ID     string `json:"-"`
	Value  string `json:"value"`
	Cv     string `json:"-"`
	TermID string `json:"-"`
}

func (prop *GenoProp) GetID() string {
	return prop.ID
}

func (prop *GenoProp) SetID(id string) error {
	prop.ID = id
	return nil
}

func (prop *GenoProp) GetName() string {
	return "chadoprops"
}

func (prop *GenoProp) GetRelatedLinksInfo() []jsapi.RelationShipLink {
	return []jsapi.RelationShipLink{
		jsapi.RelationShipLink{
			Name:           "proptype",
			SuffixFragment: fmt.Sprintf("cvs/%s/cvterms/%s", prop.Cv, prop.TermID),
		},
	}
}

type Stock struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
