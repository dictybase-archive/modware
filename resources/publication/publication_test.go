package publication

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dictyBase/go-middlewares/middlewares/router"
	"github.com/dictyBase/modware/modwaretest"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go/jsonapi"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func GetMockedDb(dbh *sql.DB) *dbr.Connection {
	return &dbr.Connection{
		DB:            dbh,
		Dialect:       dialect.PostgreSQL,
		EventReceiver: &dbr.NullEventReceiver{},
	}
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error %s during stub database connection\n", err)
	}
	defer db.Close()
	pubMockRow := sqlmock.NewRows([]string{
		"title",
		"volume",
		"series_name",
		"issue",
		"pages",
		"uniquename",
		"name",
		"pubplace",
		"pyear",
	})
	pubMockRow.AddRow(
		"dictyBase 2015: Expanding data and annotations in a new software environment",
		"12",
		"Genesis",
		"8",
		"765-80",
		"99",
		"journal_article",
		"pubmed",
		"2015",
	)
	mock.ExpectQuery("SELECT (.+) FROM pub JOIN (.+)").
		WillReturnRows(pubMockRow)

	propMockRow := sqlmock.NewRows([]string{"value", "term"})
	propMockRow.AddRow("10.1002/dvg.22867", "doi").
		AddRow("This is an abstract", "abstract").
		AddRow("ppublish", "status").
		AddRow("june", "month").
		AddRow("1526-968X", "issn")
	mock.ExpectQuery("SELECT (.+) FROM pubprop JOIN (.+) JOIN (.+) JOIN (.+)").
		WillReturnRows(propMockRow)

	//create mock http response and request
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", fmt.Sprintf("%s/%s", modwaretest.APIServer(), modwaretest.PubId), nil)
	// mock the httprouter.Params in the context
	params := make(httprouter.Params, 1)
	params[0].Key = "id"
	params[0].Value = modwaretest.PubId
	ctx := context.WithValue(context.Background(), router.ContextKeyParams, params)
	rctx := r.WithContext(ctx)

	// create the app instance with mock db
	pubResource := &Publication{Dbh: GetMockedDb(db), PathPrefix: "1.0"}
	// call the method
	pubResource.Get(w, rctx)

	if w.Code != http.StatusOK {
		t.Fatalf("unexpected http response %s with http code %d\n", w.Body.String(), w.Code)
	}
	authorRel := jsonapi.Relationship{
		Links: &jsonapi.Links{
			Self:    fmt.Sprintf("%s/publications/%s/relationships/authors", modwaretest.APIServer(), modwaretest.PubId),
			Related: fmt.Sprintf("%s/publications/%s/authors", modwaretest.APIServer(), modwaretest.PubId),
		},
	}
	pubjStruct := &jsonapi.Document{
		Links: &jsonapi.Links{Self: fmt.Sprintf("%s/%s/%s", modwaretest.APIServer(), "publications", modwaretest.PubId)},
		Data: &jsonapi.DataContainer{
			DataObject: &jsonapi.Data{
				Type: "publications",
				ID:   modwaretest.PubId,
				Attributes: []byte(`
					{
						"doi": "10.1002/dvg.22867",
						"title": "dictyBase 2015: Expanding data and annotations in a new software environment",
						"abstract": "This is an abstract",
						"journal": "Genesis",
						"year": "2015",
						"volume": "12",
						"pages":"765-80",
						"month": "june",
						"issn": "1526-968X",
						"issue":"8",
						"source": "pubmed",
						"pub_type": "journal_article",
						"status": "ppublish"
					}
				`),
				Relationships: map[string]jsonapi.Relationship{"authors": authorRel},
			},
		},
	}
	if err := modwaretest.MatchJSON(w.Body.Bytes(), pubjStruct); err != nil {
		t.Error(err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error %s\n", err)
	}

}
