package publication

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/Jeffail/gabs"
	mwtest "github.com/dictyBase/apihelpers/aphtest"
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
	if assert.True(cont.Exists("data", "attributes", "source")) {
		value, _ := cont.Path("data.attributes.source").Data().(string)
		assert.Equal(value, "pubmed", "should match the source value")
	}
	if assert.True(cont.Exists("data", "relationships", "authors", "links", "related")) {
		value, _ := cont.Path("data.relationships.authors.links.related").Data().(string)
		assert.Equal(
			value,
			fmt.Sprintf("%s/%s", mwtest.APIServer(), "publications/99/authors"),
			"should match the related links of authors relationships",
		)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error %s\n", err)
	}
}

func TestGetWithInclude(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error %s during stub database connection\n", err)
	}
	defer db.Close()

	authorMockRow := sqlmock.NewRows(authorColumns)
	for _, d := range authorData {
		authorMockRow.AddRow(d[0], d[1], d[2], d[3])
	}
	mock.ExpectQuery("SELECT (.+) FROM pubauthor JOIN (.+)").WillReturnRows(authorMockRow)

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
		AddIncludes("authors").
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
	if assert.True(cont.Exists("data", "attributes", "source")) {
		value, _ := cont.Path("data.attributes.source").Data().(string)
		assert.Equal(value, "pubmed", "should match the source value")
	}
	if assert.True(cont.Exists("data", "relationships", "authors", "links", "related")) {
		value, _ := cont.Path("data.relationships.authors.links.related").Data().(string)
		assert.Equal(
			value,
			fmt.Sprintf("%s/%s", mwtest.APIServer(), "publications/99/authors"),
			"should match the related links of authors relationships",
		)
	}
	includes, _ := cont.S("included").Children()
	assert.Equal(len(includes), 2, "should have two includes")
	for _, m := range includes {
		for _, f := range []string{"type", "id", "attributes", "relationships"} {
			assert.True(m.Exists(f), fmt.Sprintf("include resource should have %s field", f))
		}
		tval, _ := m.Path("type").Data().(string)
		assert.Equal(tval, "authors", "should be authors type")
		assert.True(
			m.Exists("relationships", "publications", "links", "related"),
			"include resource should have publications link field",
		)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error %s\n", err)
	}
}

func TestGetWithSparseField(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error %s during stub database connection\n", err)
	}
	defer db.Close()

	pubMockRow := sqlmock.NewRows(selectPubCols)
	pubMockRow.FromCSVString(strings.Join(selectpubTestData[0], ","))
	mock.ExpectQuery("SELECT (.+) FROM pub JOIN (.+)").
		WillReturnRows(pubMockRow)

	propMockRow := sqlmock.NewRows([]string{"value", "term"})
	for k, v := range selectpropTestData[0] {
		propMockRow.AddRow(v, k)
	}
	mock.ExpectQuery("SELECT (.+) FROM pubprop JOIN (.+) JOIN (.+) JOIN (.+)").
		WillReturnRows(propMockRow)

	// create the app instance with mock db
	pubResource := &Publication{Dbh: GetMockedDb(db), PathPrefix: mwtest.PathPrefix}
	spf := []string{
		"journal",
		"issue",
		"pages",
		"source",
		"year",
		"doi",
		"month",
	}
	cont := mwtest.NewHTTPExpectBuilder(t, mwtest.APIServer(), pubResource).
		Get(fmt.Sprintf("/publications/%s", mwtest.PubID)).
		AddRouterParam("id", mwtest.PubID).
		AddFieldSets("publications", false, spf...).
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
	for _, f := range spf {
		assert.True(
			cont.Exists(
				"data",
				"attributes",
				f,
			), fmt.Sprintf("should have %s fields", f),
		)
	}
	for _, f := range []string{"title", "abstract", "status"} {
		assert.False(cont.Exists("data", "attributes", f), fmt.Sprintf("should not have %s field", f))

	}
	if assert.True(cont.Exists("data", "relationships", "authors", "links", "related")) {
		value, _ := cont.Path("data.relationships.authors.links.related").Data().(string)
		assert.Equal(
			value,
			fmt.Sprintf("%s/%s", mwtest.APIServer(), "publications/99/authors"),
			"should match the related links of authors relationships",
		)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error %s\n", err)
	}
}

func TestGetWithRelatedSparseField(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error %s during stub database connection\n", err)
	}
	defer db.Close()

	pubMockRow := sqlmock.NewRows(selectPubCols)
	pubMockRow.FromCSVString(strings.Join(selectpubTestData[0], ","))
	mock.ExpectQuery("SELECT (.+) FROM pub JOIN (.+)").
		WillReturnRows(pubMockRow)

	propMockRow := sqlmock.NewRows([]string{"value", "term"})
	for k, v := range selectpropTestData[0] {
		propMockRow.AddRow(v, k)
	}
	mock.ExpectQuery("SELECT (.+) FROM pubprop JOIN (.+) JOIN (.+) JOIN (.+)").
		WillReturnRows(propMockRow)

	authorMockRow := sqlmock.NewRows(selectauthorColumns)
	for _, d := range selectauthorData {
		authorMockRow.AddRow(d[0], d[1])
	}
	mock.ExpectQuery("SELECT (.+) FROM pubauthor JOIN (.+)").WillReturnRows(authorMockRow)

	// create the app instance with mock db
	pubResource := &Publication{Dbh: GetMockedDb(db), PathPrefix: mwtest.PathPrefix}
	spf := []string{
		"journal",
		"issue",
		"pages",
		"source",
		"year",
		"doi",
		"month",
	}
	cont := mwtest.NewHTTPExpectBuilder(t, mwtest.APIServer(), pubResource).
		Get(fmt.Sprintf("/publications/%s", mwtest.PubID)).
		AddRouterParam("id", mwtest.PubID).
		AddIncludes("authors").
		AddFieldSets("publications", false, spf...).
		AddFieldSets("authors", true, "rank").
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
	for _, f := range spf {
		assert.True(
			cont.Exists(
				"data",
				"attributes",
				f,
			), fmt.Sprintf("should have %s fields", f),
		)
	}
	for _, f := range []string{"title", "abstract", "status"} {
		assert.False(cont.Exists("data", "attributes", f), fmt.Sprintf("should not have %s field", f))

	}
	if assert.True(cont.Exists("data", "relationships", "authors", "links", "related")) {
		value, _ := cont.Path("data.relationships.authors.links.related").Data().(string)
		assert.Equal(
			value,
			fmt.Sprintf("%s/%s", mwtest.APIServer(), "publications/99/authors"),
			"should match the related links of authors relationships",
		)
	}
	includes, _ := cont.S("included").Children()
	assert.Equal(len(includes), 2, "should have two includes")
	for _, m := range includes {
		for _, f := range []string{"type", "id", "attributes", "relationships"} {
			assert.True(m.Exists(f), fmt.Sprintf("include resource should have %s field", f))
		}
		tval, _ := m.Path("type").Data().(string)
		assert.Equal(tval, "authors", "should be authors type")
		assert.True(
			m.Exists("relationships", "publications", "links", "related"),
			"include resource should have publications link field",
		)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error %s\n", err)
	}
}

// TestGetAll runs /publications?page[number]=2&page[size]=3
func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error %s during stub database connection\n", err)
	}
	defer db.Close()

	// mock the sql backend with test data
	countMockRow := sqlmock.NewRows([]string{"records"})
	countMockRow.AddRow("7")
	mock.ExpectQuery("SELECT (.+) FROM pub").
		WillReturnRows(countMockRow)

	pubMockRow := sqlmock.NewRows(pubColumns)
	for _, v := range getPubTestDataRows() {
		pubMockRow.FromCSVString(strings.Join(v, ","))
	}
	mock.ExpectQuery("SELECT (.+) FROM pub JOIN (.+) LIMIT 3 OFFSET 3").
		WillReturnRows(pubMockRow)

	for _, r := range getPropsTestDataRows() {
		propMockRow := sqlmock.NewRows([]string{"value", "term"})
		for k, v := range r {
			propMockRow.AddRow(v, k)
		}
		mock.ExpectQuery("SELECT (.+) FROM pubprop JOIN (.+) JOIN (.+) JOIN (.+)").
			WillReturnRows(propMockRow)
	}

	// create the app instance with mock db
	pubResource := &Publication{Dbh: GetMockedDb(db), PathPrefix: mwtest.PathPrefix}
	// run the http request
	pageNum := 2
	pageSize := 3
	cont := mwtest.NewHTTPExpectBuilder(t, mwtest.APIServer(), pubResource).
		GetAll("/publications").
		AddPagination(pageNum, pageSize).
		Expect().
		Status(http.StatusOK).
		JSON()

	assert := assert.New(t)
	// tests the members
	members, _ := cont.S("data").Children()
	assert.Equal(len(members), 3, "should have 3 members")
	for i, v := range []string{"10", "11", "12"} {
		testMembers(assert, members[i], v)
	}

	// test the meta section
	if assert.True(cont.Exists("meta", "pagination")) {
		num, _ := cont.Path("meta.pagination.number").Data().(float64)
		assert.Equal(pageNum, int(num), "should match the current page number")
		size, _ := cont.Path("meta.pagination.size").Data().(float64)
		assert.Equal(pageSize, int(size), "should match the page size")
		last, _ := cont.Path("meta.pagination.total").Data().(float64)
		assert.Equal(3, int(last), "should match the last page")
		rec, _ := cont.Path("meta.pagination.records").Data().(float64)
		assert.Equal(7, int(rec), "should match the total records")
	}
	// test the pagination links
	if assert.True(cont.Exists("links")) {
		fmap := map[string]int{
			"first": 1,
			"last":  3,
			"next":  3,
			"prev":  1,
			"self":  2,
		}
		lnk := cont.Path("links")
		for k, v := range fmap {
			testPageLink(assert, lnk, k, v, 3)
		}
	}
	//t.Log(string(mwtest.IndentJSON(cont.Bytes())))

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error %s\n", err)
	}
}

// TestGetAll runs /publications?filter[title]=pand
func TestGetAllWithFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error %s during stub database connection\n", err)
	}
	defer db.Close()

	// mock the sql backend with test data
	countMockRow := sqlmock.NewRows([]string{"records"})
	countMockRow.AddRow("3")
	mock.ExpectQuery(`SELECT count\(pub_id\) AS records FROM pub WHERE \(\(pub.title ILIKE '%Exp%'\) AND.*\)`).
		WillReturnRows(countMockRow)

	pubMockRow := sqlmock.NewRows(pubColumns)
	for _, v := range getPubTestDataRows() {
		pubMockRow.FromCSVString(strings.Join(v, ","))
	}
	mock.ExpectQuery(`SELECT (.+) FROM pub JOIN (.+) WHERE \(\(pub.title ILIKE '%Exp%'\) AND.*\) LIMIT 3 OFFSET 0`).
		WillReturnRows(pubMockRow)

	for _, r := range getPropsTestDataRows() {
		propMockRow := sqlmock.NewRows([]string{"value", "term"})
		for k, v := range r {
			propMockRow.AddRow(v, k)
		}
		mock.ExpectQuery("SELECT (.+) FROM pubprop JOIN (.+) JOIN (.+) JOIN (.+)").
			WillReturnRows(propMockRow)
	}

	// create the app instance with mock db
	pubResource := &Publication{Dbh: GetMockedDb(db), PathPrefix: mwtest.PathPrefix}
	// run the http request
	cont := mwtest.NewHTTPExpectBuilder(t, mwtest.APIServer(), pubResource).
		GetAll("/publications").
		AddFilter("title", "Exp").
		Expect().
		Status(http.StatusOK).
		JSON()

	assert := assert.New(t)
	// tests the members
	members, _ := cont.S("data").Children()
	assert.Equal(len(members), 3, "should have 3 members")
	for i, v := range []string{"10", "11", "12"} {
		testMembers(assert, members[i], v)
	}

	//// test the meta section
	if assert.True(cont.Exists("meta", "pagination")) {
		num, _ := cont.Path("meta.pagination.number").Data().(float64)
		assert.Equal(1, int(num), "should match the current page number")
		size, _ := cont.Path("meta.pagination.size").Data().(float64)
		assert.Equal(3, int(size), "should match the page size")
		last, _ := cont.Path("meta.pagination.total").Data().(float64)
		assert.Equal(1, int(last), "should match the last page")
		rec, _ := cont.Path("meta.pagination.records").Data().(float64)
		assert.Equal(3, int(rec), "should match the total records")
	}
	//// test the pagination links
	if assert.True(cont.Exists("links")) {
		fmap := map[string]int{
			"first": 1,
			"last":  1,
			"self":  1,
		}
		lnk := cont.Path("links")
		for k, v := range fmap {
			testPageLink(assert, lnk, k, v, 3)
		}
	}
	//t.Log(string(mwtest.IndentJSON(cont.Bytes())))

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error %s\n", err)
	}
}

func testPageLink(assert *assert.Assertions, cont *gabs.Container, field string, pageNum, pageSize int) {
	if assert.True(cont.Exists(field), fmt.Sprintf("should have %s page field", field)) {
		value, _ := cont.Path(field).Data().(string)
		assert.Equal(
			value,
			fmt.Sprintf("%s/publications?page[number]=%d&page[size]=%d", mwtest.APIServer(), pageNum, pageSize),
			fmt.Sprintf("should match url %s page field", field),
		)
	}
}

func testMembers(assert *assert.Assertions, cont *gabs.Container, id string) {
	if assert.True(cont.Exists("type"), "should have type member") {
		value, _ := cont.Path("type").Data().(string)
		assert.Equal(value, "publications", "should match the type value")
	}
	if assert.True(cont.Exists("id"), "should have ID member") {
		value, _ := cont.Path("id").Data().(string)
		assert.Equal(value, id, "should match the id value")
	}
	if assert.True(cont.Exists("attributes", "source")) {
		value, _ := cont.Path("attributes.source").Data().(string)
		assert.Equal(value, "pubmed", "should match the source value")
	}
	if assert.True(cont.Exists("relationships", "authors", "links", "related")) {
		value, _ := cont.Path("relationships.authors.links.related").Data().(string)
		assert.Equal(
			value,
			fmt.Sprintf("%s/publications/%s/authors", mwtest.APIServer(), id),
			"should match the related links of authors relationships",
		)
	}
}
