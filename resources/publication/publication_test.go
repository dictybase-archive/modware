package publication

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"testing"

	mwtest "github.com/dictyBase/modware/modwaretest"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"github.com/stretchr/testify/assert"

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
	pubMockRow := sqlmock.NewRows(pubColumns)
	pubMockRow.FromCSVString(strings.Join(pubTestData[0], ","))
	mock.ExpectQuery("SELECT (.+) FROM pub JOIN (.+)").
		WillReturnRows(pubMockRow)
	propMockRow := sqlmock.NewRows([]string{"value", "term"})
	for k, v := range propTestData[0] {
		propMockRow.AddRow(v, k)
	}
	mock.ExpectQuery("SELECT (.+) FROM pubprop JOIN (.+) JOIN (.+) JOIN (.+)").
		WillReturnRows(propMockRow)

	// create the app instance with mock db
	pubResource := &Publication{Dbh: GetMockedDb(db), PathPrefix: mwtest.PathPrefix}
	cont := mwtest.NewHTTPExpectBuilder(t, mwtest.APIServer(), pubResource).
		Get(fmt.Sprintf("/publications/%s", mwtest.PubID)).
		AddRouterParam("id", mwtest.PubID).
		Expect().
		Status(http.StatusOK).
		JSON()
	assert := assert.New(t)
	if assert.True(cont.Exists("links", "self"), "should have link member") {
		value, _ := cont.Path("links.self").Data().(string)
		assert.Equal(
			value, fmt.Sprintf(
				"%s/publications/%s",
				mwtest.APIServer(),
				mwtest.PubID,
			),
			"should match the top level link",
		)
	}
	if assert.True(cont.Exists("data", "type"), "should have type member") {
		value, _ := cont.Path("data.type").Data().(string)
		assert.Equal(value, "publications", "should match the type value")
	}
	if assert.True(cont.Exists("data", "id"), "should have ID member") {
		value, _ := cont.Path("data.id").Data().(string)
		assert.Equal(value, "99", "should match the id value")
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error %s\n", err)
	}
}
