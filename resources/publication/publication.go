package publication

import (
	"net/http"

	"github.com/dictyBase/apihelpers/apherror"
	"github.com/dictyBase/go-middlewares/middlewares/query"
	"github.com/dictyBase/go-middlewares/middlewares/router"
	jsapi "github.com/dictyBase/modware/models/jsonapi"
	"github.com/dictyBase/modware/models/jsonapi/publication"
	"github.com/dictyBase/modware/render"
	"github.com/dictyBase/modware/resources"
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
	//Validates include params
	p, ok := r.Context().Value(query.ContextKeyQueryParams).(*query.Params)
	if ok && p.HasIncludes {
		err := jsapi.ValidateRelationships(&publication.Publication{}, p.Includes)
		if err != nil {
			apherror.JSONAPIError(
				w,
				apherror.ErrIncludeParam.New(err.Error()),
			)
			return
		}
	}
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
		apherror.DatabaseError(w, err)
		return
	}
	props, err := pub.getProps(sess, id)
	if err != nil {
		apherror.DatabaseError(w, err)
		return
	}
	var authors []*publication.Author
	if ok && p.HasIncludes {
		// Now include author
		authors, err = pub.getAuthors(sess, id)
		if err != nil {
			apherror.DatabaseError(w, err)
			return
		}
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
	if len(authors) > 0 {
		pubj.Authors = authors
	}
	render.Resource(pubj, resources.GetAPIServerInfo(r, pub.PathPrefix), w)
}

func (pub *Publication) GetAll(w http.ResponseWriter, r *http.Request) {
	pageProps := resources.GetPaginationProp(r)
	// Get the total no of publications
	sess := pub.GetDbh().NewSession(nil)
	var count int
	err := sess.Select("count(pub_id)").From("pub").LoadValue(&count)
	if err != nil {
		apherror.DatabaseError(w, err)
		return
	}
	pageProps.Records = count
	// Now get a list of publications within that page range
	var pubRows []*pubData
	err = sess.Select(
		"pub.title",
		"pub.volume",
		"pub.series_name",
		"pub.issue",
		"pub.pages",
		"pub.uniquename",
		"cvterm.name",
		"pub.pubplace",
		"pub.pyear",
	).From("pub").
		Join("cvterm", dbr.Eq("pub.type_id", "cvterm.cvterm_id")).
		Where(dbr.Eq("pub.is_obsolete", "0")).
		Paginate(uint64(pageProps.Current), uint64(pageProps.Entries)).
		LoadStruct(&pubRows)
	if err != nil {
		apherror.DatabaseError(w, err)
		return
	}
	var pubSlice []*publication.Publication
	for _, pb := range pubRows {
		props, err := pub.getProps(sess, pb.PubId)
		if err != nil {
			apherror.DatabaseError(w, err)
			return
		}
		pubj := &publication.Publication{
			ID:      pb.PubId,
			Title:   pb.Title,
			Journal: pb.Journal,
			Year:    pb.Year,
			Volume:  pb.Volume.String,
			Pages:   pb.Pages.String,
			PubType: pb.PubType,
			Source:  pb.Source,
			Issue:   pb.Issue.String,
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
		pubSlice = append(pubSlice, pubj)
	}
	render.ResourceCollection(pubSlice, resources.GetAPIServerInfo(r, pub.PathPrefix), w, pageProps)
}

func (pub *Publication) Create(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) Update(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) Delete(w http.ResponseWriter, r *http.Request) {
}

func (pub *Publication) getProps(sess *dbr.Session, id string) ([]*pubProp, error) {
	var props []*pubProp
	_, err := sess.Select(
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
	return props, err
}

func (pub *Publication) getAuthors(sess *dbr.Session, id string) ([]*publication.Author, error) {
	var authors []*publication.Author
	_, err := sess.Select(
		"pubauthor.pubauthor_id",
		"pubauthor.rank",
		"pubauthor.surname",
		"pubauthor.givennames",
	).From("pubauthor").
		Join("pub", dbr.Eq("pubauthor.pub_id", "pub.pub_id")).
		Where(
			dbr.And(
				dbr.Eq("pub.uniquename", id),
				dbr.Eq("pub.is_obsolete", "0"),
			),
		).LoadStructs(&authors)
	if err != nil {
		return authors, err
	}
	return authors, nil
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
