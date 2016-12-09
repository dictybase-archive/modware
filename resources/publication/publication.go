package publication

import (
	"net/http"

	"github.com/dictyBase/go-middlewares/middlewares/router"
	jsapi "github.com/dictyBase/modware/models/jsonapi"
	"github.com/dictyBase/modware/models/jsonapi/publication"
	"github.com/dictyBase/modware/render"
	"github.com/gocraft/dbr"
)

type pubData struct {
	PubId   string `db:"uniquename"`
	Title   string
	Volume  dbr.NullString
	Journal string `db:"series_name"`
	Issue   dbr.NullString
	Year    string `db:"pyear"`
	Pages   dbr.NullString
	PubType string `db:"name"`
	Source  string `db:"pubplace"`
}

type pubProp struct {
	Value string
	Term  string `dbr:"name"`
}

type Publication struct {
	Dbh        *dbr.Connection
	PathPrefix string
}

func (pub *Publication) GetDbh() *dbr.Connection {
	return pub.Dbh
}

func (pub *Publication) Get(w http.ResponseWriter, r *http.Request) {
	id := router.Params(r).ByName("id")
	sess := pub.GetDbh().NewSession(nil)
	var pubRow pubData
	err := sess.Select(
		"pub.title",
		"pub.volume",
		"pub.series_name",
		"pub.issue",
		"pub.pages",
		"pub.uniquename",
		"cvterm.name",
		"pub.pubplace",
		"pub.pyear",
	).From("pub").Join("cvterm", dbr.Eq("pub.type_id", "cvterm.cvterm_id")).
		Where(
			dbr.And(
				dbr.Eq("pub.uniquename", id),
				dbr.Eq("pub.is_obsolete", "0"),
			),
		).LoadStruct(&pubRow)
	if err != nil {
		render.DatabaseError(w, err)
		return
	}

	var props []pubProp
	_, err = sess.Select(
		"pubprop.value",
		"cvterm.name as term",
	).From("pubprop").
		Join("pub", dbr.Eq("pubprop.pub_id", "pub.pub_id")).
		Join("cvterm", dbr.Eq("pubprop.type_id", "cvterm.type_id")).
		Join("cv", dbr.Eq("cvterm.cv_id", "cv.cv_id")).
		Where(
			dbr.And(
				dbr.Eq("pub.uniquename", id),
				dbr.Eq("pub.is_obsolete", "0"),
				dbr.Eq("cv.name", "pub_type"),
			),
		).LoadStructs(&props)
	if err != nil {
		render.DatabaseError(w, err)
		return
	}

	// publication type struct that will be converted to json
	pubj := &publication.Publication{
		ID:      pubRow.PubId,
		Title:   pubRow.Title,
		Journal: pubRow.Journal,
		Year:    pubRow.Year,
		Volume:  pubRow.Volume.String,
		Pages:   pubRow.Pages.String,
		PubType: pubRow.PubType,
		Source:  pubRow.Source,
		Issue:   pubRow.Issue.String,
	}
	for _, p := range props {
		v := p.Value
		switch p.Term {
		case "doi":
			pubj.Doi = v
		case "status":
			pubj.Status = v
		case "month":
			pubj.Month = v
		case "issn":
			pubj.Issn = v
		case "abstract":
			pubj.Abstract = v
		}
	}
	pubJsapi, err := jsapi.MarshalToStructWrapper(pubj, resoures.GetApiServerInfo(r, pub.PathPrefix))
	if err != nil {
		render.StructMarshallingError(w, err)
	}
	if err := render.JSONAPI(w, http.StatusOK, pubJsapi); err != nil {
		render.JSONEncodingError(w, err)
	}
}

func (pub *Publication) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) Create(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) Update(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) Delete(w http.ResponseWriter, r *http.Request) {
}

type Author struct {
	Dbh        *dbr.Connection
	PathPrefix string
}

func (a *Author) GetDbh() *dbr.Connection {
	return a.Dbh
}

func (a *Author) Get(w http.ResponseWriter, r *http.Request) {

}

func (a *Author) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (a *Author) Create(w http.ResponseWriter, r *http.Request) {

}

func (a *Author) Update(w http.ResponseWriter, r *http.Request) {

}

func (a *Author) Delete(w http.ResponseWriter, r *http.Request) {

}
