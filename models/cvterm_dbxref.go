package models

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
)

// CvtermDbxref is an object representing the database table.
type CvtermDbxref struct {
	CvtermDbxrefID  int `boil:"cvterm_dbxref_id" json:"cvterm_dbxref_id" toml:"cvterm_dbxref_id" yaml:"cvterm_dbxref_id"`
	CvtermID        int `boil:"cvterm_id" json:"cvterm_id" toml:"cvterm_id" yaml:"cvterm_id"`
	DbxrefID        int `boil:"dbxref_id" json:"dbxref_id" toml:"dbxref_id" yaml:"dbxref_id"`
	IsForDefinition int `boil:"is_for_definition" json:"is_for_definition" toml:"is_for_definition" yaml:"is_for_definition"`

	R *cvtermDbxrefR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cvtermDbxrefL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// cvtermDbxrefR is where relationships are stored.
type cvtermDbxrefR struct {
	Cvterm *Cvterm
	Dbxref *Dbxref
}

// cvtermDbxrefL is where Load methods for each relationship are stored.
type cvtermDbxrefL struct{}

var (
	cvtermDbxrefColumns               = []string{"cvterm_dbxref_id", "cvterm_id", "dbxref_id", "is_for_definition"}
	cvtermDbxrefColumnsWithoutDefault = []string{"cvterm_id", "dbxref_id"}
	cvtermDbxrefColumnsWithDefault    = []string{"cvterm_dbxref_id", "is_for_definition"}
	cvtermDbxrefPrimaryKeyColumns     = []string{"cvterm_dbxref_id"}
)

type (
	// CvtermDbxrefSlice is an alias for a slice of pointers to CvtermDbxref.
	// This should generally be used opposed to []CvtermDbxref.
	CvtermDbxrefSlice []*CvtermDbxref
	// CvtermDbxrefHook is the signature for custom CvtermDbxref hook methods
	CvtermDbxrefHook func(boil.Executor, *CvtermDbxref) error

	cvtermDbxrefQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cvtermDbxrefType                 = reflect.TypeOf(&CvtermDbxref{})
	cvtermDbxrefMapping              = queries.MakeStructMapping(cvtermDbxrefType)
	cvtermDbxrefPrimaryKeyMapping, _ = queries.BindMapping(cvtermDbxrefType, cvtermDbxrefMapping, cvtermDbxrefPrimaryKeyColumns)
	cvtermDbxrefInsertCacheMut       sync.RWMutex
	cvtermDbxrefInsertCache          = make(map[string]insertCache)
	cvtermDbxrefUpdateCacheMut       sync.RWMutex
	cvtermDbxrefUpdateCache          = make(map[string]updateCache)
	cvtermDbxrefUpsertCacheMut       sync.RWMutex
	cvtermDbxrefUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var cvtermDbxrefBeforeInsertHooks []CvtermDbxrefHook
var cvtermDbxrefBeforeUpdateHooks []CvtermDbxrefHook
var cvtermDbxrefBeforeDeleteHooks []CvtermDbxrefHook
var cvtermDbxrefBeforeUpsertHooks []CvtermDbxrefHook

var cvtermDbxrefAfterInsertHooks []CvtermDbxrefHook
var cvtermDbxrefAfterSelectHooks []CvtermDbxrefHook
var cvtermDbxrefAfterUpdateHooks []CvtermDbxrefHook
var cvtermDbxrefAfterDeleteHooks []CvtermDbxrefHook
var cvtermDbxrefAfterUpsertHooks []CvtermDbxrefHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CvtermDbxref) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermDbxrefBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CvtermDbxref) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermDbxrefBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CvtermDbxref) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermDbxrefBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CvtermDbxref) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermDbxrefBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CvtermDbxref) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermDbxrefAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CvtermDbxref) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermDbxrefAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CvtermDbxref) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermDbxrefAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CvtermDbxref) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermDbxrefAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CvtermDbxref) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermDbxrefAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCvtermDbxrefHook registers your hook function for all future operations.
func AddCvtermDbxrefHook(hookPoint boil.HookPoint, cvtermDbxrefHook CvtermDbxrefHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cvtermDbxrefBeforeInsertHooks = append(cvtermDbxrefBeforeInsertHooks, cvtermDbxrefHook)
	case boil.BeforeUpdateHook:
		cvtermDbxrefBeforeUpdateHooks = append(cvtermDbxrefBeforeUpdateHooks, cvtermDbxrefHook)
	case boil.BeforeDeleteHook:
		cvtermDbxrefBeforeDeleteHooks = append(cvtermDbxrefBeforeDeleteHooks, cvtermDbxrefHook)
	case boil.BeforeUpsertHook:
		cvtermDbxrefBeforeUpsertHooks = append(cvtermDbxrefBeforeUpsertHooks, cvtermDbxrefHook)
	case boil.AfterInsertHook:
		cvtermDbxrefAfterInsertHooks = append(cvtermDbxrefAfterInsertHooks, cvtermDbxrefHook)
	case boil.AfterSelectHook:
		cvtermDbxrefAfterSelectHooks = append(cvtermDbxrefAfterSelectHooks, cvtermDbxrefHook)
	case boil.AfterUpdateHook:
		cvtermDbxrefAfterUpdateHooks = append(cvtermDbxrefAfterUpdateHooks, cvtermDbxrefHook)
	case boil.AfterDeleteHook:
		cvtermDbxrefAfterDeleteHooks = append(cvtermDbxrefAfterDeleteHooks, cvtermDbxrefHook)
	case boil.AfterUpsertHook:
		cvtermDbxrefAfterUpsertHooks = append(cvtermDbxrefAfterUpsertHooks, cvtermDbxrefHook)
	}
}

// OneP returns a single cvtermDbxref record from the query, and panics on error.
func (q cvtermDbxrefQuery) OneP() *CvtermDbxref {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single cvtermDbxref record from the query.
func (q cvtermDbxrefQuery) One() (*CvtermDbxref, error) {
	o := &CvtermDbxref{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cvterm_dbxref")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all CvtermDbxref records from the query, and panics on error.
func (q cvtermDbxrefQuery) AllP() CvtermDbxrefSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all CvtermDbxref records from the query.
func (q cvtermDbxrefQuery) All() (CvtermDbxrefSlice, error) {
	var o CvtermDbxrefSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CvtermDbxref slice")
	}

	if len(cvtermDbxrefAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all CvtermDbxref records in the query, and panics on error.
func (q cvtermDbxrefQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all CvtermDbxref records in the query.
func (q cvtermDbxrefQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cvterm_dbxref rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q cvtermDbxrefQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q cvtermDbxrefQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cvterm_dbxref exists")
	}

	return count > 0, nil
}

// CvtermG pointed to by the foreign key.
func (o *CvtermDbxref) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *CvtermDbxref) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// DbxrefG pointed to by the foreign key.
func (o *CvtermDbxref) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *CvtermDbxref) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// LoadCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermDbxrefL) LoadCvterm(e boil.Executor, singular bool, maybeCvtermDbxref interface{}) error {
	var slice []*CvtermDbxref
	var object *CvtermDbxref

	count := 1
	if singular {
		object = maybeCvtermDbxref.(*CvtermDbxref)
	} else {
		slice = *maybeCvtermDbxref.(*CvtermDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermDbxrefR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermDbxrefR{}
			args[i] = obj.CvtermID
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

	if len(cvtermDbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Cvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvtermID {
				local.R.Cvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermDbxrefL) LoadDbxref(e boil.Executor, singular bool, maybeCvtermDbxref interface{}) error {
	var slice []*CvtermDbxref
	var object *CvtermDbxref

	count := 1
	if singular {
		object = maybeCvtermDbxref.(*CvtermDbxref)
	} else {
		slice = *maybeCvtermDbxref.(*CvtermDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermDbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermDbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"dbxref\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Dbxref")
	}
	defer results.Close()

	var resultSlice []*Dbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Dbxref")
	}

	if len(cvtermDbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Dbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID == foreign.DbxrefID {
				local.R.Dbxref = foreign
				break
			}
		}
	}

	return nil
}

// SetCvterm of the cvterm_dbxref to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.CvtermDbxref.
func (o *CvtermDbxref) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvterm_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CvtermID = related.CvtermID

	if o.R == nil {
		o.R = &cvtermDbxrefR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			CvtermDbxref: o,
		}
	} else {
		related.R.CvtermDbxref = o
	}

	return nil
}

// SetDbxref of the cvterm_dbxref to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.CvtermDbxref.
func (o *CvtermDbxref) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvterm_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.DbxrefID, o.CvtermDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DbxrefID = related.DbxrefID

	if o.R == nil {
		o.R = &cvtermDbxrefR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			CvtermDbxref: o,
		}
	} else {
		related.R.CvtermDbxref = o
	}

	return nil
}

// CvtermDbxrefsG retrieves all records.
func CvtermDbxrefsG(mods ...qm.QueryMod) cvtermDbxrefQuery {
	return CvtermDbxrefs(boil.GetDB(), mods...)
}

// CvtermDbxrefs retrieves all the records using an executor.
func CvtermDbxrefs(exec boil.Executor, mods ...qm.QueryMod) cvtermDbxrefQuery {
	mods = append(mods, qm.From("\"cvterm_dbxref\""))
	return cvtermDbxrefQuery{NewQuery(exec, mods...)}
}

// FindCvtermDbxrefG retrieves a single record by ID.
func FindCvtermDbxrefG(cvtermDbxrefID int, selectCols ...string) (*CvtermDbxref, error) {
	return FindCvtermDbxref(boil.GetDB(), cvtermDbxrefID, selectCols...)
}

// FindCvtermDbxrefGP retrieves a single record by ID, and panics on error.
func FindCvtermDbxrefGP(cvtermDbxrefID int, selectCols ...string) *CvtermDbxref {
	retobj, err := FindCvtermDbxref(boil.GetDB(), cvtermDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCvtermDbxref retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCvtermDbxref(exec boil.Executor, cvtermDbxrefID int, selectCols ...string) (*CvtermDbxref, error) {
	cvtermDbxrefObj := &CvtermDbxref{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cvterm_dbxref\" where \"cvterm_dbxref_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, cvtermDbxrefID)

	err := q.Bind(cvtermDbxrefObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cvterm_dbxref")
	}

	return cvtermDbxrefObj, nil
}

// FindCvtermDbxrefP retrieves a single record by ID with an executor, and panics on error.
func FindCvtermDbxrefP(exec boil.Executor, cvtermDbxrefID int, selectCols ...string) *CvtermDbxref {
	retobj, err := FindCvtermDbxref(exec, cvtermDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CvtermDbxref) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *CvtermDbxref) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *CvtermDbxref) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *CvtermDbxref) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no cvterm_dbxref provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermDbxrefColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	cvtermDbxrefInsertCacheMut.RLock()
	cache, cached := cvtermDbxrefInsertCache[key]
	cvtermDbxrefInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			cvtermDbxrefColumns,
			cvtermDbxrefColumnsWithDefault,
			cvtermDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(cvtermDbxrefType, cvtermDbxrefMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cvtermDbxrefType, cvtermDbxrefMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"cvterm_dbxref\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into cvterm_dbxref")
	}

	if !cached {
		cvtermDbxrefInsertCacheMut.Lock()
		cvtermDbxrefInsertCache[key] = cache
		cvtermDbxrefInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single CvtermDbxref record. See Update for
// whitelist behavior description.
func (o *CvtermDbxref) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single CvtermDbxref record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *CvtermDbxref) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the CvtermDbxref, and panics on error.
// See Update for whitelist behavior description.
func (o *CvtermDbxref) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the CvtermDbxref.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *CvtermDbxref) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	cvtermDbxrefUpdateCacheMut.RLock()
	cache, cached := cvtermDbxrefUpdateCache[key]
	cvtermDbxrefUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(cvtermDbxrefColumns, cvtermDbxrefPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update cvterm_dbxref, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cvterm_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cvtermDbxrefPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cvtermDbxrefType, cvtermDbxrefMapping, append(wl, cvtermDbxrefPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update cvterm_dbxref row")
	}

	if !cached {
		cvtermDbxrefUpdateCacheMut.Lock()
		cvtermDbxrefUpdateCache[key] = cache
		cvtermDbxrefUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q cvtermDbxrefQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q cvtermDbxrefQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for cvterm_dbxref")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CvtermDbxrefSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CvtermDbxrefSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CvtermDbxrefSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CvtermDbxrefSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"cvterm_dbxref\" SET %s WHERE (\"cvterm_dbxref_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermDbxrefPrimaryKeyColumns), len(colNames)+1, len(cvtermDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in cvtermDbxref slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CvtermDbxref) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *CvtermDbxref) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *CvtermDbxref) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *CvtermDbxref) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no cvterm_dbxref provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermDbxrefColumnsWithDefault, o)

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

	cvtermDbxrefUpsertCacheMut.RLock()
	cache, cached := cvtermDbxrefUpsertCache[key]
	cvtermDbxrefUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			cvtermDbxrefColumns,
			cvtermDbxrefColumnsWithDefault,
			cvtermDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			cvtermDbxrefColumns,
			cvtermDbxrefPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert cvterm_dbxref, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cvtermDbxrefPrimaryKeyColumns))
			copy(conflict, cvtermDbxrefPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"cvterm_dbxref\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(cvtermDbxrefType, cvtermDbxrefMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cvtermDbxrefType, cvtermDbxrefMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cvterm_dbxref")
	}

	if !cached {
		cvtermDbxrefUpsertCacheMut.Lock()
		cvtermDbxrefUpsertCache[key] = cache
		cvtermDbxrefUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single CvtermDbxref record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CvtermDbxref) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single CvtermDbxref record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CvtermDbxref) DeleteG() error {
	if o == nil {
		return errors.New("models: no CvtermDbxref provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single CvtermDbxref record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CvtermDbxref) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single CvtermDbxref record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CvtermDbxref) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no CvtermDbxref provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cvtermDbxrefPrimaryKeyMapping)
	sql := "DELETE FROM \"cvterm_dbxref\" WHERE \"cvterm_dbxref_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from cvterm_dbxref")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q cvtermDbxrefQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q cvtermDbxrefQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no cvtermDbxrefQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from cvterm_dbxref")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CvtermDbxrefSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CvtermDbxrefSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no CvtermDbxref slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CvtermDbxrefSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CvtermDbxrefSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no CvtermDbxref slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(cvtermDbxrefBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"cvterm_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermDbxrefPrimaryKeyColumns), 1, len(cvtermDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from cvtermDbxref slice")
	}

	if len(cvtermDbxrefAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *CvtermDbxref) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *CvtermDbxref) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CvtermDbxref) ReloadG() error {
	if o == nil {
		return errors.New("models: no CvtermDbxref provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CvtermDbxref) Reload(exec boil.Executor) error {
	ret, err := FindCvtermDbxref(exec, o.CvtermDbxrefID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermDbxrefSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermDbxrefSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermDbxrefSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty CvtermDbxrefSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermDbxrefSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	cvtermDbxrefs := CvtermDbxrefSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"cvterm_dbxref\".* FROM \"cvterm_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(cvtermDbxrefPrimaryKeyColumns), 1, len(cvtermDbxrefPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&cvtermDbxrefs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CvtermDbxrefSlice")
	}

	*o = cvtermDbxrefs

	return nil
}

// CvtermDbxrefExists checks if the CvtermDbxref row exists.
func CvtermDbxrefExists(exec boil.Executor, cvtermDbxrefID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"cvterm_dbxref\" where \"cvterm_dbxref_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, cvtermDbxrefID)
	}

	row := exec.QueryRow(sql, cvtermDbxrefID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cvterm_dbxref exists")
	}

	return exists, nil
}

// CvtermDbxrefExistsG checks if the CvtermDbxref row exists.
func CvtermDbxrefExistsG(cvtermDbxrefID int) (bool, error) {
	return CvtermDbxrefExists(boil.GetDB(), cvtermDbxrefID)
}

// CvtermDbxrefExistsGP checks if the CvtermDbxref row exists. Panics on error.
func CvtermDbxrefExistsGP(cvtermDbxrefID int) bool {
	e, err := CvtermDbxrefExists(boil.GetDB(), cvtermDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CvtermDbxrefExistsP checks if the CvtermDbxref row exists. Panics on error.
func CvtermDbxrefExistsP(exec boil.Executor, cvtermDbxrefID int) bool {
	e, err := CvtermDbxrefExists(exec, cvtermDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
