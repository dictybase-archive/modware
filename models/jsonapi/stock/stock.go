package stock

import (
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
	ID          string `json:"-"`
	Name        string `json:"name"`
	Observation string `json:"observation"`
	Attribute   string `json:"phen_attribute"`
	Value       string `json:"value"`
	Cvalue      string `json:"cvalue"`
	Evidence    string `json:"evidence"`
}

func (ph *Phenotype) GetID() string {
	return ph.ID
}

func (ph *Phenotype) SetID(id string) error {
	ph.ID = id
	return nil
}
