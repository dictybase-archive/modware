package publication

import (
	"net/http"

	"github.com/dictyBase/apihelpers/apherror"
	"github.com/dictyBase/apihelpers/aphrender"
	"github.com/dictyBase/go-middlewares/middlewares/pagination"
	"github.com/dictyBase/go-middlewares/middlewares/query"
	"github.com/dictyBase/go-middlewares/middlewares/router"
	"github.com/dictyBase/modware/models/jsonapi/publication"
	"github.com/dictyBase/modware/resources"
	"github.com/dictyBase/modware/resources/validate"
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
	//Validates all params
	p, ok, err := validate.Params(r, &publication.Publication{})
	if err != nil {
		apherror.JSONAPIError(w, err)
		return
	}
	id := router.Params(r).ByName("id")
	sess := pub.GetDbh().NewSession(nil)
	pubStr := new(publication.Publication)
	authors := make([]*publication.Author, 0)
	if ok {
		switch {
		case includeAuthors(p):
			authors, err = pub.getAuthors(sess, id)
			if err != nil {
				apherror.DatabaseError(w, err)
				return
			}
		case p.HasSparseFields:
			if _, ok := p.SparseFields["publications"]; ok {
				pubStr, err = pub.getSelectedRows(p.SparseFields["publications"], sess, id)
				if err != nil {
					apherror.DatabaseError(w, err)
					return
				}
			}
			if _, ok := p.SparseFields["authors"]; ok {
				authors, err = pub.getSelectedAuthors(p.SparseFields["authors"], sess, id)
				if err != nil {
					apherror.DatabaseError(w, err)
					return
				}
			}
		}
	}
	if len(pubStr.ID) == 0 {
		pubStr, err = pub.getRows(sess, id)
		if err != nil {
			apherror.DatabaseError(w, err)
			return
		}
	}
	if len(authors) > 0 {
		pubStr.Authors = authors
	}
	aphrender.Resource(pubStr, resources.GetAPIServerInfo(r, pub.PathPrefix), w)
}

func (pub *Publication) GetAll(w http.ResponseWriter, r *http.Request) {
	//Validates all params
	p, ok, err := validate.Params(r, &publication.Publication{})
	if err != nil {
		apherror.JSONAPIError(w, err)
		return
	}
	pageProps := resources.GetPaginationProp(r)
	// Get the total no of publications
	sess := pub.GetDbh().NewSession(nil)
	pubSlice := make([]*publication.Publication, 0)
	// Without any query params
	if !ok {
		count, err := pub.getCount(sess)
		if err != nil {
			apherror.DatabaseError(w, err)
			return
		}
		SetPagination(count, pageProps)
		pubSlice, err = pub.getAllRows(sess, pageProps)
		if err != nil {
			apherror.DatabaseError(w, err)
			return
		}
		aphrender.ResourceCollection(pubSlice, resources.GetAPIServerInfo(r, pub.PathPrefix), w, pageProps)
		return
	}
	switch { // Various combinations of query parameters
	case p.HasFilters && p.HasSparseFields:
		if _, ok := p.SparseFields["publications"]; ok {
			count, err := pub.getFilteredCount(p, sess)
			if err != nil {
				apherror.DatabaseError(w, err)
				return
			}
			SetPagination(count, pageProps)
			pubSlice, err = pub.getAllSelectFilteredRows(
				p,
				p.SparseFields["publications"],
				sess,
				pageProps,
			)
			if err != nil {
				apherror.DatabaseError(w, err)
				return
			}
		}
		if _, ok := p.SparseFields["authors"]; ok {
			if len(pubSlice) == 0 {
				count, err := pub.getCount(sess)
				if err != nil {
					apherror.DatabaseError(w, err)
					return
				}
				SetPagination(count, pageProps)
				pubSlice, err = pub.getAllRows(sess, pageProps)
				if err != nil {
					apherror.DatabaseError(w, err)
					return
				}
			}
			for _, ps := range pubSlice {
				authors, err := pub.getSelectedAuthors(p.SparseFields["authors"], sess, ps.ID)
				if err != nil {
					apherror.DatabaseError(w, err)
					return
				}
				if len(authors) > 0 {
					ps.Authors = authors
				}
			}
		} else {
			if p.HasIncludes { // only author includes
				for _, ps := range pubSlice {
					authors, err := pub.getAuthors(sess, ps.ID)
					if err != nil {
						apherror.DatabaseError(w, err)
						return
					}
					if len(authors) > 0 {
						ps.Authors = authors
					}
				}
			}
		}
	case p.HasFilters: // /publications?filter[title]=something
		count, err := pub.getFilteredCount(p, sess)
		if err != nil {
			apherror.DatabaseError(w, err)
			return
		}
		SetPagination(count, pageProps)
		pubSlice, err = pub.getAllFilteredRows(p, sess, pageProps)
		if err != nil {
			apherror.DatabaseError(w, err)
			return
		}
		if p.HasIncludes { // /publications?filter[title]=something&include=authors
			for _, ps := range pubSlice {
				authors, err := pub.getAuthors(sess, ps.ID)
				if err != nil {
					apherror.DatabaseError(w, err)
					return
				}
				if len(authors) > 0 {
					ps.Authors = authors
				}
			}
		}
	case p.HasSparseFields: // /publications?fields[publications]=title,doi,year&fields[authors]=rank&include=authors
		count, err := pub.getCount(sess)
		if err != nil {
			apherror.DatabaseError(w, err)
			return
		}
		SetPagination(count, pageProps)
		if _, ok := p.SparseFields["publications"]; ok {
			pubSlice, err = pub.getAllSelectedRows(p.SparseFields["publications"], sess, pageProps)
			if err != nil {
				apherror.DatabaseError(w, err)
				return
			}
		}
		if _, ok := p.SparseFields["authors"]; ok {
			if len(pubSlice) == 0 {
				count, err := pub.getCount(sess)
				if err != nil {
					apherror.DatabaseError(w, err)
					return
				}
				SetPagination(count, pageProps)
				pubSlice, err = pub.getAllRows(sess, pageProps)
				if err != nil {
					apherror.DatabaseError(w, err)
					return
				}
			}
			for _, ps := range pubSlice {
				authors, err := pub.getSelectedAuthors(p.SparseFields["authors"], sess, ps.ID)
				if err != nil {
					apherror.DatabaseError(w, err)
					return
				}
				if len(authors) > 0 {
					ps.Authors = authors
				}
			}
		} else {
			if p.HasIncludes { // /publications?fields[publications]=title,doi,year&include=authors
				for _, ps := range pubSlice {
					authors, err := pub.getAuthors(sess, ps.ID)
					if err != nil {
						apherror.DatabaseError(w, err)
						return
					}
					if len(authors) > 0 {
						ps.Authors = authors
					}
				}
			}
		}
	default:
		count, err := pub.getCount(sess)
		if err != nil {
			apherror.DatabaseError(w, err)
			return
		}
		SetPagination(count, pageProps)
		pubSlice, err = pub.getAllRows(sess, pageProps)
		if err != nil {
			apherror.DatabaseError(w, err)
			return
		}
		for _, ps := range pubSlice {
			authors, err := pub.getAuthors(sess, ps.ID)
			if err != nil {
				apherror.DatabaseError(w, err)
				return
			}
			if len(authors) > 0 {
				ps.Authors = authors
			}
		}
	}
	aphrender.ResourceCollection(pubSlice, resources.GetAPIServerInfo(r, pub.PathPrefix), w, pageProps)
}

func (pub *Publication) Create(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) Update(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) Delete(w http.ResponseWriter, r *http.Request) {
}

func (pub *Publication) getRows(sess *dbr.Session, id string) (*publication.Publication, error) {
	row := new(pubData)
	pubStr := new(publication.Publication)
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
		).LoadStruct(row)
	if err != nil {
		return pubStr, err
	}
	props, err := pub.getProps(sess, id)
	if err != nil {
		return pubStr, err
	}
	for _, p := range props {
		v := p.Value
		switch p.Term {
		case "doi":
			pubStr.Doi = v
		case "status":
			pubStr.Status = v
		case "month":
			pubStr.Month = v
		case "issn":
			pubStr.Issn = v
		case "abstract":
			pubStr.Abstract = v
		}
	}
	pubStr.ID = row.PubId
	pubStr.Title = row.Title
	pubStr.Journal = row.Journal
	pubStr.Year = row.Year
	pubStr.Volume = row.Volume.String
	pubStr.Pages = row.Pages.String
	pubStr.PubType = row.PubType
	pubStr.Source = row.Source
	pubStr.Issue = row.Issue.String
	return pubStr, nil
}

func (pub *Publication) getCount(sess *dbr.Session) (int, error) {
	var count int
	err := sess.Select("count(pub_id) as records").From("pub").LoadValue(&count)
	return count, err
}

func (pub *Publication) getFilteredCount(p *query.Params, sess *dbr.Session) (int, error) {
	rmap := new(publication.Publication).GetMap()
	builders := append(
		resources.BuildFilterQuery(rmap, p),
		dbr.Eq("pub.is_obsolete", "0"),
	)
	var count int
	selBuilder := sess.Select("count(pub_id) AS records").
		From("pub").
		Where(dbr.And(builders...))
	err := selBuilder.LoadValue(&count)
	return count, err
}

func (pub *Publication) getAllSelectFilteredRows(p *query.Params, f *query.Fields, sess *dbr.Session, pageProps *pagination.Props) ([]*publication.Publication, error) {
	builders := append(
		resources.BuildFilterQuery(new(publication.Publication).GetMap(), p),
		dbr.Eq("pub.is_obsolete", "0"),
	)
	pubRows := make([]*pubData, 0)
	pubSlice := make([]*publication.Publication, 0)
	columns := []string{"pub.uniquename"}
	terms := make([]string, 0)
	for _, n := range f.GetAll() {
		switch n {
		case "title":
			columns = append(columns, "pub.title")
		case "volume":
			columns = append(columns, "pub.volume")
		case "journal":
			columns = append(columns, "pub.series_name")
		case "issue":
			columns = append(columns, "pub.issue")
		case "year":
			columns = append(columns, "pub.pyear")
		case "pages":
			columns = append(columns, "pub.pages")
		case "pub_type":
			columns = append(columns, "cvterm.name")
		case "source":
			columns = append(columns, "pub.pubplace")
		default:
			terms = append(terms, n)
		}
	}
	err := sess.Select(columns...).From("pub").
		Join("cvterm", dbr.Eq("pub.type_id", "cvterm.cvterm_id")).
		Where(dbr.And(builders...)).
		Paginate(uint64(pageProps.Current), uint64(pageProps.Entries)).
		LoadStruct(&pubRows)
	if err != nil {
		return pubSlice, err
	}
	if len(terms) > 0 {
		for _, pb := range pubRows {
			props, err := pub.getSelectedProps(terms, sess, pb.PubId)
			if err != nil {
				return pubSlice, err
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
	}
	return pubSlice, nil
}

func (pub *Publication) getAllFilteredRows(p *query.Params, sess *dbr.Session, pageProps *pagination.Props) ([]*publication.Publication, error) {
	rmap := new(publication.Publication).GetMap()
	builders := resources.BuildFilterQuery(rmap, p)
	builders = append(builders, dbr.Eq("pub.is_obsolete", "0"))
	pubRows := make([]*pubData, 0)
	pubSlice := make([]*publication.Publication, 0)
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
	).From("pub").
		Join("cvterm", dbr.Eq("pub.type_id", "cvterm.cvterm_id")).
		Where(dbr.And(builders...)).
		Paginate(uint64(pageProps.Current), uint64(pageProps.Entries)).
		LoadStruct(&pubRows)
	if err != nil {
		return pubSlice, err
	}
	for _, pb := range pubRows {
		props, err := pub.getProps(sess, pb.PubId)
		if err != nil {
			return pubSlice, err
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
	return pubSlice, nil
}

func (pub *Publication) getAllRows(sess *dbr.Session, pageProps *pagination.Props) ([]*publication.Publication, error) {
	pubRows := make([]*pubData, 0)
	pubSlice := make([]*publication.Publication, 0)
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
	).From("pub").
		Join("cvterm", dbr.Eq("pub.type_id", "cvterm.cvterm_id")).
		Where(dbr.Eq("pub.is_obsolete", "0")).
		Paginate(uint64(pageProps.Current), uint64(pageProps.Entries)).
		LoadStruct(&pubRows)
	if err != nil {
		return pubSlice, err
	}
	for _, pb := range pubRows {
		props, err := pub.getProps(sess, pb.PubId)
		if err != nil {
			return pubSlice, err
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
	return pubSlice, nil
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

func (pub *Publication) getSelectedProps(terms []string, sess *dbr.Session, id string) ([]*pubProp, error) {
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
				dbr.Expr("cvterm.name IN ?", terms),
			),
		).LoadStructs(&props)
	return props, err
}

func (pub *Publication) getAllSelectedRows(f *query.Fields, sess *dbr.Session, pageProps *pagination.Props) ([]*publication.Publication, error) {
	terms := make([]string, 0)
	pubRows := make([]*pubData, 0)
	pubSlice := make([]*publication.Publication, 0)
	columns := []string{"pub.uniquename"}
	for _, n := range f.GetAll() {
		switch n {
		case "title":
			columns = append(columns, "pub.title")
		case "volume":
			columns = append(columns, "pub.volume")
		case "journal":
			columns = append(columns, "pub.series_name")
		case "issue":
			columns = append(columns, "pub.issue")
		case "year":
			columns = append(columns, "pub.pyear")
		case "pages":
			columns = append(columns, "pub.pages")
		case "pub_type":
			columns = append(columns, "cvterm.name")
		case "source":
			columns = append(columns, "pub.pubplace")
		default:
			terms = append(terms, n)
		}
	}
	err := sess.Select(columns...).From("pub").
		Join("cvterm", dbr.Eq("pub.type_id", "cvterm.cvterm_id")).
		Where(dbr.Eq("pub.is_obsolete", "0")).
		Paginate(uint64(pageProps.Current), uint64(pageProps.Entries)).
		LoadStruct(&pubRows)
	if err != nil {
		return pubSlice, err
	}
	if len(terms) > 0 {
		for _, pb := range pubRows {
			props, err := pub.getSelectedProps(terms, sess, pb.PubId)
			if err != nil {
				return pubSlice, err
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
	}
	return pubSlice, nil
}

func (pub *Publication) getSelectedRows(f *query.Fields, sess *dbr.Session, id string) (*publication.Publication, error) {
	terms := make([]string, 0)
	row := new(pubData)
	pubStr := new(publication.Publication)
	columns := []string{"pub.uniquename"}
	for _, n := range f.GetAll() {
		switch n {
		case "title":
			columns = append(columns, "pub.title")
		case "volume":
			columns = append(columns, "pub.volume")
		case "journal":
			columns = append(columns, "pub.series_name")
		case "issue":
			columns = append(columns, "pub.issue")
		case "year":
			columns = append(columns, "pub.pyear")
		case "pages":
			columns = append(columns, "pub.pages")
		case "pub_type":
			columns = append(columns, "cvterm.name")
		case "source":
			columns = append(columns, "pub.pubplace")
		default:
			terms = append(terms, n)
		}
	}
	err := sess.Select(columns...).From("pub").Join("cvterm", dbr.Eq("pub.type_id", "cvterm.cvterm_id")).
		Where(
			dbr.And(
				dbr.Eq("pub.uniquename", id),
				dbr.Eq("pub.is_obsolete", "0"),
			),
		).LoadStruct(row)
	if err != nil {
		return pubStr, err
	}
	if len(terms) > 0 {
		props, err := pub.getSelectedProps(terms, sess, id)
		if err != nil {
			return pubStr, err
		}
		for _, p := range props {
			v := p.Value
			switch p.Term {
			case "doi":
				pubStr.Doi = v
			case "status":
				pubStr.Status = v
			case "month":
				pubStr.Month = v
			case "issn":
				pubStr.Issn = v
			case "abstract":
				pubStr.Abstract = v
			}
		}
	}
	pubStr.ID = row.PubId
	pubStr.Title = row.Title
	pubStr.Journal = row.Journal
	pubStr.Year = row.Year
	pubStr.Volume = row.Volume.String
	pubStr.Pages = row.Pages.String
	pubStr.PubType = row.PubType
	pubStr.Source = row.Source
	pubStr.Issue = row.Issue.String
	return pubStr, nil
}

func (pub *Publication) getSelectedAuthors(f *query.Fields, sess *dbr.Session, id string) ([]*publication.Author, error) {
	authors := make([]*publication.Author, 0)
	columns := []string{"pubauthor.pubauthor_id"}
	for _, c := range f.GetAll() {
		switch c {
		case "last_name":
			columns = append(columns, "pubauthor.surname")
		case "given_names":
			columns = append(columns, "pubauthor.givennames")
		case "rank":
			columns = append(columns, "pubauthor.rank")
		default:
			continue
		}
	}
	_, err := sess.Select(
		columns...,
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

func SetPagination(count int, pageProps *pagination.Props) {
	if pageProps.Entries >= count {
		pageProps.Entries = count
		pageProps.Current = 1
	}
	pageProps.Records = count
}

func includeAuthors(p *query.Params) bool {
	if p.HasIncludes && p.HasSparseFields {
		if _, ok := p.SparseFields["authors"]; ok {
			return false
		}
	}
	return p.HasIncludes
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
