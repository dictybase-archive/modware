package publication

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dictyBase/go-middlewares/middlewares/router"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"github.com/julienschmidt/httprouter"

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
		"26088819",
		"pub_type",
		"pubmed",
		"2015",
	)
	mock.ExpectQuery("SELECT (.+) FROM pub JOIN (.+)").
		WillReturnRows(pubMockRow)

	propMockRow := sqlmock.NewRows([]string{"value", "term"})
	propMockRow.AddRow("doi", "10.1002/dvg.22867").
		AddRow("abstract", "This is an abstract").
		AddRow("status", "ppublish").
		AddRow("month", "june").
		AddRow("issn", "1526-968X")
	mock.ExpectQuery("SELECT (.+) FROM pubprop JOIN (.+) JOIN (.+) JOIN (.+)").
		WillReturnRows(propMockRow)

	//create mock http response and request
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "https://api.dictybase.org/publications/99", nil)
	// mock the httprouter.Params in the context
	params := make(httprouter.Params, 1)
	params[0].Key = "id"
	params[0].Value = "99"
	ctx := context.WithValue(context.Background(), router.ContextKeyParams, params)
	rctx := r.WithContext(ctx)

	// create the app instance with mock db
	pubResource := &Publication{Dbh: GetMockedDb(db), PathPrefix: "1.0"}
	// call the method
	pubResource.Get(w, rctx)

	if w.Code != http.StatusOK {
		t.Fatalf("unexpected http response %s with http code %d\n", w.Body.String(), w.Code)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error %s\n", err)
	}
}
