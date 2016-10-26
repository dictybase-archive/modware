package chado

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v5"
)

// Cvtermpath is an object representing the database table.
type Cvtermpath struct {
	CvtermpathID int      `boil:"cvtermpath_id" json:"cvtermpath_id" toml:"cvtermpath_id" yaml:"cvtermpath_id"`
	TypeID       null.Int `boil:"type_id" json:"type_id,omitempty" toml:"type_id" yaml:"type_id,omitempty"`
	SubjectID    int      `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	ObjectID     int      `boil:"object_id" json:"object_id" toml:"object_id" yaml:"object_id"`
	CVID         int      `boil:"cv_id" json:"cv_id" toml:"cv_id" yaml:"cv_id"`
	Pathdistance null.Int `boil:"pathdistance" json:"pathdistance,omitempty" toml:"pathdistance" yaml:"pathdistance,omitempty"`

	R *cvtermpathR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cvtermpathL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// cvtermpathR is where relationships are stored.
type cvtermpathR struct {
	CV      *CV
	Object  *Cvterm
	Subject *Cvterm
	Type    *Cvterm
}

// cvtermpathL is where Load methods for each relationship are stored.
type cvtermpathL struct{}

var (
	cvtermpathColumns               = []string{"cvtermpath_id", "type_id", "subject_id", "object_id", "cv_id", "pathdistance"}
	cvtermpathColumnsWithoutDefault = []string{"type_id", "subject_id", "object_id", "cv_id", "pathdistance"}
	cvtermpathColumnsWithDefault    = []string{"cvtermpath_id"}
	cvtermpathPrimaryKeyColumns     = []string{"cvtermpath_id"}
)

type (
	// CvtermpathSlice is an alias for a slice of pointers to Cvtermpath.
	// This should generally be used opposed to []Cvtermpath.
	CvtermpathSlice []*Cvtermpath
	// CvtermpathHook is the signature for custom Cvtermpath hook methods
	CvtermpathHook func(boil.Executor, *Cvtermpath) error

	cvtermpathQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cvtermpathType                 = reflect.TypeOf(&Cvtermpath{})
	cvtermpathMapping              = queries.MakeStructMapping(cvtermpathType)
	cvtermpathPrimaryKeyMapping, _ = queries.BindMapping(cvtermpathType, cvtermpathMapping, cvtermpathPrimaryKeyColumns)
	cvtermpathInsertCacheMut       sync.RWMutex
	cvtermpathInsertCache          = make(map[string]insertCache)
	cvtermpathUpdateCacheMut       sync.RWMutex
	cvtermpathUpdateCache          = make(map[string]updateCache)
	cvtermpathUpsertCacheMut       sync.RWMutex
	cvtermpathUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var cvtermpathBeforeInsertHooks []CvtermpathHook
var cvtermpathBeforeUpdateHooks []CvtermpathHook
var cvtermpathBeforeDeleteHooks []CvtermpathHook
var cvtermpathBeforeUpsertHooks []CvtermpathHook

var cvtermpathAfterInsertHooks []CvtermpathHook
var cvtermpathAfterSelectHooks []CvtermpathHook
var cvtermpathAfterUpdateHooks []CvtermpathHook
var cvtermpathAfterDeleteHooks []CvtermpathHook
var cvtermpathAfterUpsertHooks []CvtermpathHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Cvtermpath) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpathBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Cvtermpath) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpathBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Cvtermpath) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpathBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Cvtermpath) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpathBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Cvtermpath) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpathAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Cvtermpath) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpathAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Cvtermpath) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpathAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Cvtermpath) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpathAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Cvtermpath) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpathAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCvtermpathHook registers your hook function for all future operations.
func AddCvtermpathHook(hookPoint boil.HookPoint, cvtermpathHook CvtermpathHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cvtermpathBeforeInsertHooks = append(cvtermpathBeforeInsertHooks, cvtermpathHook)
	case boil.BeforeUpdateHook:
		cvtermpathBeforeUpdateHooks = append(cvtermpathBeforeUpdateHooks, cvtermpathHook)
	case boil.BeforeDeleteHook:
		cvtermpathBeforeDeleteHooks = append(cvtermpathBeforeDeleteHooks, cvtermpathHook)
	case boil.BeforeUpsertHook:
		cvtermpathBeforeUpsertHooks = append(cvtermpathBeforeUpsertHooks, cvtermpathHook)
	case boil.AfterInsertHook:
		cvtermpathAfterInsertHooks = append(cvtermpathAfterInsertHooks, cvtermpathHook)
	case boil.AfterSelectHook:
		cvtermpathAfterSelectHooks = append(cvtermpathAfterSelectHooks, cvtermpathHook)
	case boil.AfterUpdateHook:
		cvtermpathAfterUpdateHooks = append(cvtermpathAfterUpdateHooks, cvtermpathHook)
	case boil.AfterDeleteHook:
		cvtermpathAfterDeleteHooks = append(cvtermpathAfterDeleteHooks, cvtermpathHook)
	case boil.AfterUpsertHook:
		cvtermpathAfterUpsertHooks = append(cvtermpathAfterUpsertHooks, cvtermpathHook)
	}
}

// OneP returns a single cvtermpath record from the query, and panics on error.
func (q cvtermpathQuery) OneP() *Cvtermpath {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single cvtermpath record from the query.
func (q cvtermpathQuery) One() (*Cvtermpath, error) {
	o := &Cvtermpath{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for cvtermpath")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Cvtermpath records from the query, and panics on error.
func (q cvtermpathQuery) AllP() CvtermpathSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Cvtermpath records from the query.
func (q cvtermpathQuery) All() (CvtermpathSlice, error) {
	var o CvtermpathSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Cvtermpath slice")
	}

	if len(cvtermpathAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Cvtermpath records in the query, and panics on error.
func (q cvtermpathQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Cvtermpath records in the query.
func (q cvtermpathQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count cvtermpath rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q cvtermpathQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q cvtermpathQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if cvtermpath exists")
	}

	return count > 0, nil
}

// CVG pointed to by the foreign key.
func (o *Cvtermpath) CVG(mods ...qm.QueryMod) cvQuery {
	return o.CV(boil.GetDB(), mods...)
}

// CV pointed to by the foreign key.
func (o *Cvtermpath) CV(exec boil.Executor, mods ...qm.QueryMod) cvQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cv_id=$1", o.CVID),
	}

	queryMods = append(queryMods, mods...)

	query := CVS(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cv\"")

	return query
}

// ObjectG pointed to by the foreign key.
func (o *Cvtermpath) ObjectG(mods ...qm.QueryMod) cvtermQuery {
	return o.Object(boil.GetDB(), mods...)
}

// Object pointed to by the foreign key.
func (o *Cvtermpath) Object(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.ObjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// SubjectG pointed to by the foreign key.
func (o *Cvtermpath) SubjectG(mods ...qm.QueryMod) cvtermQuery {
	return o.Subject(boil.GetDB(), mods...)
}

// Subject pointed to by the foreign key.
func (o *Cvtermpath) Subject(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.SubjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Cvtermpath) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Cvtermpath) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadCV allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermpathL) LoadCV(e boil.Executor, singular bool, maybeCvtermpath interface{}) error {
	var slice []*Cvtermpath
	var object *Cvtermpath

	count := 1
	if singular {
		object = maybeCvtermpath.(*Cvtermpath)
	} else {
		slice = *maybeCvtermpath.(*CvtermpathSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermpathR{}
		args[0] = object.CVID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermpathR{}
			args[i] = obj.CVID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cv\" where \"cv_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CV")
	}
	defer results.Close()

	var resultSlice []*CV
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CV")
	}

	if len(cvtermpathAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.CV = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CVID == foreign.CVID {
				local.R.CV = foreign
				break
			}
		}
	}

	return nil
}

// LoadObject allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermpathL) LoadObject(e boil.Executor, singular bool, maybeCvtermpath interface{}) error {
	var slice []*Cvtermpath
	var object *Cvtermpath

	count := 1
	if singular {
		object = maybeCvtermpath.(*Cvtermpath)
	} else {
		slice = *maybeCvtermpath.(*CvtermpathSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermpathR{}
		args[0] = object.ObjectID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermpathR{}
			args[i] = obj.ObjectID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvterm")
	}
	defer results.Close()

	var resultSlice []*Cvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvterm")
	}

	if len(cvtermpathAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Object = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ObjectID == foreign.CvtermID {
				local.R.Object = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubject allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermpathL) LoadSubject(e boil.Executor, singular bool, maybeCvtermpath interface{}) error {
	var slice []*Cvtermpath
	var object *Cvtermpath

	count := 1
	if singular {
		object = maybeCvtermpath.(*Cvtermpath)
	} else {
		slice = *maybeCvtermpath.(*CvtermpathSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermpathR{}
		args[0] = object.SubjectID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermpathR{}
			args[i] = obj.SubjectID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvterm")
	}
	defer results.Close()

	var resultSlice []*Cvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvterm")
	}

	if len(cvtermpathAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Subject = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.SubjectID == foreign.CvtermID {
				local.R.Subject = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermpathL) LoadType(e boil.Executor, singular bool, maybeCvtermpath interface{}) error {
	var slice []*Cvtermpath
	var object *Cvtermpath

	count := 1
	if singular {
		object = maybeCvtermpath.(*Cvtermpath)
	} else {
		slice = *maybeCvtermpath.(*CvtermpathSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermpathR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermpathR{}
			args[i] = obj.TypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvterm")
	}
	defer results.Close()

	var resultSlice []*Cvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvterm")
	}

	if len(cvtermpathAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Type = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.TypeID.Int == foreign.CvtermID {
				local.R.Type = foreign
				break
			}
		}
	}

	return nil
}

// SetCV of the cvtermpath to the related item.
// Sets o.R.CV to related.
// Adds o to related.R.Cvtermpaths.
func (o *Cvtermpath) SetCV(exec boil.Executor, insert bool, related *CV) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvtermpath\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cv_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermpathPrimaryKeyColumns),
	)
	values := []interface{}{related.CVID, o.CvtermpathID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CVID = related.CVID

	if o.R == nil {
		o.R = &cvtermpathR{
			CV: related,
		}
	} else {
		o.R.CV = related
	}

	if related.R == nil {
		related.R = &cvR{
			Cvtermpaths: CvtermpathSlice{o},
		}
	} else {
		related.R.Cvtermpaths = append(related.R.Cvtermpaths, o)
	}

	return nil
}

// SetObject of the cvtermpath to the related item.
// Sets o.R.Object to related.
// Adds o to related.R.ObjectCvtermpath.
func (o *Cvtermpath) SetObject(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvtermpath\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermpathPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermpathID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ObjectID = related.CvtermID

	if o.R == nil {
		o.R = &cvtermpathR{
			Object: related,
		}
	} else {
		o.R.Object = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			ObjectCvtermpath: o,
		}
	} else {
		related.R.ObjectCvtermpath = o
	}

	return nil
}

// SetSubject of the cvtermpath to the related item.
// Sets o.R.Subject to related.
// Adds o to related.R.SubjectCvtermpath.
func (o *Cvtermpath) SetSubject(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvtermpath\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermpathPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermpathID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SubjectID = related.CvtermID

	if o.R == nil {
		o.R = &cvtermpathR{
			Subject: related,
		}
	} else {
		o.R.Subject = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			SubjectCvtermpath: o,
		}
	} else {
		related.R.SubjectCvtermpath = o
	}

	return nil
}

// SetType of the cvtermpath to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeCvtermpath.
func (o *Cvtermpath) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvtermpath\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermpathPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermpathID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID.Int = related.CvtermID
	o.TypeID.Valid = true

	if o.R == nil {
		o.R = &cvtermpathR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeCvtermpath: o,
		}
	} else {
		related.R.TypeCvtermpath = o
	}

	return nil
}

// RemoveType relationship.
// Sets o.R.Type to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Cvtermpath) RemoveType(exec boil.Executor, related *Cvterm) error {
	var err error

	o.TypeID.Valid = false
	if err = o.Update(exec, "type_id"); err != nil {
		o.TypeID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Type = nil
	if related == nil || related.R == nil {
		return nil
	}

	related.R.TypeCvtermpath = nil
	return nil
}

// CvtermpathsG retrieves all records.
func CvtermpathsG(mods ...qm.QueryMod) cvtermpathQuery {
	return Cvtermpaths(boil.GetDB(), mods...)
}

// Cvtermpaths retrieves all the records using an executor.
func Cvtermpaths(exec boil.Executor, mods ...qm.QueryMod) cvtermpathQuery {
	mods = append(mods, qm.From("\"cvtermpath\""))
	return cvtermpathQuery{NewQuery(exec, mods...)}
}

// FindCvtermpathG retrieves a single record by ID.
func FindCvtermpathG(cvtermpathID int, selectCols ...string) (*Cvtermpath, error) {
	return FindCvtermpath(boil.GetDB(), cvtermpathID, selectCols...)
}

// FindCvtermpathGP retrieves a single record by ID, and panics on error.
func FindCvtermpathGP(cvtermpathID int, selectCols ...string) *Cvtermpath {
	retobj, err := FindCvtermpath(boil.GetDB(), cvtermpathID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCvtermpath retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCvtermpath(exec boil.Executor, cvtermpathID int, selectCols ...string) (*Cvtermpath, error) {
	cvtermpathObj := &Cvtermpath{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cvtermpath\" where \"cvtermpath_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, cvtermpathID)

	err := q.Bind(cvtermpathObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from cvtermpath")
	}

	return cvtermpathObj, nil
}

// FindCvtermpathP retrieves a single record by ID with an executor, and panics on error.
func FindCvtermpathP(exec boil.Executor, cvtermpathID int, selectCols ...string) *Cvtermpath {
	retobj, err := FindCvtermpath(exec, cvtermpathID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Cvtermpath) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Cvtermpath) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Cvtermpath) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Cvtermpath) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no cvtermpath provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermpathColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	cvtermpathInsertCacheMut.RLock()
	cache, cached := cvtermpathInsertCache[key]
	cvtermpathInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			cvtermpathColumns,
			cvtermpathColumnsWithDefault,
			cvtermpathColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(cvtermpathType, cvtermpathMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cvtermpathType, cvtermpathMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"cvtermpath\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "chado: unable to insert into cvtermpath")
	}

	if !cached {
		cvtermpathInsertCacheMut.Lock()
		cvtermpathInsertCache[key] = cache
		cvtermpathInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Cvtermpath record. See Update for
// whitelist behavior description.
func (o *Cvtermpath) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Cvtermpath record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Cvtermpath) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Cvtermpath, and panics on error.
// See Update for whitelist behavior description.
func (o *Cvtermpath) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Cvtermpath.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Cvtermpath) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	cvtermpathUpdateCacheMut.RLock()
	cache, cached := cvtermpathUpdateCache[key]
	cvtermpathUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(cvtermpathColumns, cvtermpathPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update cvtermpath, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cvtermpath\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cvtermpathPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cvtermpathType, cvtermpathMapping, append(wl, cvtermpathPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update cvtermpath row")
	}

	if !cached {
		cvtermpathUpdateCacheMut.Lock()
		cvtermpathUpdateCache[key] = cache
		cvtermpathUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q cvtermpathQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q cvtermpathQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for cvtermpath")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CvtermpathSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CvtermpathSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CvtermpathSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CvtermpathSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("chado: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermpathPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"cvtermpath\" SET %s WHERE (\"cvtermpath_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermpathPrimaryKeyColumns), len(colNames)+1, len(cvtermpathPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in cvtermpath slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Cvtermpath) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Cvtermpath) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Cvtermpath) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Cvtermpath) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no cvtermpath provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermpathColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	cvtermpathUpsertCacheMut.RLock()
	cache, cached := cvtermpathUpsertCache[key]
	cvtermpathUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			cvtermpathColumns,
			cvtermpathColumnsWithDefault,
			cvtermpathColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			cvtermpathColumns,
			cvtermpathPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert cvtermpath, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cvtermpathPrimaryKeyColumns))
			copy(conflict, cvtermpathPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"cvtermpath\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(cvtermpathType, cvtermpathMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cvtermpathType, cvtermpathMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "chado: unable to upsert for cvtermpath")
	}

	if !cached {
		cvtermpathUpsertCacheMut.Lock()
		cvtermpathUpsertCache[key] = cache
		cvtermpathUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Cvtermpath record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cvtermpath) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Cvtermpath record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Cvtermpath) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Cvtermpath provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Cvtermpath record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cvtermpath) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Cvtermpath record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Cvtermpath) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Cvtermpath provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cvtermpathPrimaryKeyMapping)
	sql := "DELETE FROM \"cvtermpath\" WHERE \"cvtermpath_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from cvtermpath")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q cvtermpathQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q cvtermpathQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no cvtermpathQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from cvtermpath")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CvtermpathSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CvtermpathSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Cvtermpath slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CvtermpathSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CvtermpathSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Cvtermpath slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(cvtermpathBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermpathPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"cvtermpath\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermpathPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermpathPrimaryKeyColumns), 1, len(cvtermpathPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from cvtermpath slice")
	}

	if len(cvtermpathAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Cvtermpath) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Cvtermpath) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Cvtermpath) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Cvtermpath provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Cvtermpath) Reload(exec boil.Executor) error {
	ret, err := FindCvtermpath(exec, o.CvtermpathID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermpathSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermpathSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermpathSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty CvtermpathSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermpathSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	cvtermpaths := CvtermpathSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermpathPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"cvtermpath\".* FROM \"cvtermpath\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermpathPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(cvtermpathPrimaryKeyColumns), 1, len(cvtermpathPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&cvtermpaths)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in CvtermpathSlice")
	}

	*o = cvtermpaths

	return nil
}

// CvtermpathExists checks if the Cvtermpath row exists.
func CvtermpathExists(exec boil.Executor, cvtermpathID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"cvtermpath\" where \"cvtermpath_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, cvtermpathID)
	}

	row := exec.QueryRow(sql, cvtermpathID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if cvtermpath exists")
	}

	return exists, nil
}

// CvtermpathExistsG checks if the Cvtermpath row exists.
func CvtermpathExistsG(cvtermpathID int) (bool, error) {
	return CvtermpathExists(boil.GetDB(), cvtermpathID)
}

// CvtermpathExistsGP checks if the Cvtermpath row exists. Panics on error.
func CvtermpathExistsGP(cvtermpathID int) bool {
	e, err := CvtermpathExists(boil.GetDB(), cvtermpathID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CvtermpathExistsP checks if the Cvtermpath row exists. Panics on error.
func CvtermpathExistsP(exec boil.Executor, cvtermpathID int) bool {
	e, err := CvtermpathExists(exec, cvtermpathID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
