package general

type Dbxref struct {
	ID        string `json:"-"`
	Database  string `json:"database"`
	Accession string `json:"accession"`
}

func (dbxref Dbxref) GetID() string {
	return dbxref.ID
}

func (dbxref Dbxref) SetID(id string) error {
	dbxref.ID = id
	return nil
}
