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

// CV is an object representing the database table.
type CV struct {
	CVID       int         `boil:"cv_id" json:"cv_id" toml:"cv_id" yaml:"cv_id"`
	Name       string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Definition null.String `boil:"definition" json:"definition,omitempty" toml:"definition" yaml:"definition,omitempty"`

	R *cvR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cvL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// cvR is where relationships are stored.
type cvR struct {
	Cvterm      *Cvterm
	Cvprop      *Cvprop
	Cvtermpaths CvtermpathSlice
}

// cvL is where Load methods for each relationship are stored.
type cvL struct{}

var (
	cvColumns               = []string{"cv_id", "name", "definition"}
	cvColumnsWithoutDefault = []string{"name", "definition"}
	cvColumnsWithDefault    = []string{"cv_id"}
	cvPrimaryKeyColumns     = []string{"cv_id"}
)

type (
	// CVSlice is an alias for a slice of pointers to CV.
	// This should generally be used opposed to []CV.
	CVSlice []*CV
	// CVHook is the signature for custom CV hook methods
	CVHook func(boil.Executor, *CV) error

	cvQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cvType                 = reflect.TypeOf(&CV{})
	cvMapping              = queries.MakeStructMapping(cvType)
	cvPrimaryKeyMapping, _ = queries.BindMapping(cvType, cvMapping, cvPrimaryKeyColumns)
	cvInsertCacheMut       sync.RWMutex
	cvInsertCache          = make(map[string]insertCache)
	cvUpdateCacheMut       sync.RWMutex
	cvUpdateCache          = make(map[string]updateCache)
	cvUpsertCacheMut       sync.RWMutex
	cvUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var cvBeforeInsertHooks []CVHook
var cvBeforeUpdateHooks []CVHook
var cvBeforeDeleteHooks []CVHook
var cvBeforeUpsertHooks []CVHook

var cvAfterInsertHooks []CVHook
var cvAfterSelectHooks []CVHook
var cvAfterUpdateHooks []CVHook
var cvAfterDeleteHooks []CVHook
var cvAfterUpsertHooks []CVHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CV) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CV) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CV) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CV) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CV) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CV) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range cvAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CV) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CV) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CV) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCVHook registers your hook function for all future operations.
func AddCVHook(hookPoint boil.HookPoint, cvHook CVHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cvBeforeInsertHooks = append(cvBeforeInsertHooks, cvHook)
	case boil.BeforeUpdateHook:
		cvBeforeUpdateHooks = append(cvBeforeUpdateHooks, cvHook)
	case boil.BeforeDeleteHook:
		cvBeforeDeleteHooks = append(cvBeforeDeleteHooks, cvHook)
	case boil.BeforeUpsertHook:
		cvBeforeUpsertHooks = append(cvBeforeUpsertHooks, cvHook)
	case boil.AfterInsertHook:
		cvAfterInsertHooks = append(cvAfterInsertHooks, cvHook)
	case boil.AfterSelectHook:
		cvAfterSelectHooks = append(cvAfterSelectHooks, cvHook)
	case boil.AfterUpdateHook:
		cvAfterUpdateHooks = append(cvAfterUpdateHooks, cvHook)
	case boil.AfterDeleteHook:
		cvAfterDeleteHooks = append(cvAfterDeleteHooks, cvHook)
	case boil.AfterUpsertHook:
		cvAfterUpsertHooks = append(cvAfterUpsertHooks, cvHook)
	}
}

// OneP returns a single cv record from the query, and panics on error.
func (q cvQuery) OneP() *CV {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single cv record from the query.
func (q cvQuery) One() (*CV, error) {
	o := &CV{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for cv")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all CV records from the query, and panics on error.
func (q cvQuery) AllP() CVSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all CV records from the query.
func (q cvQuery) All() (CVSlice, error) {
	var o CVSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to CV slice")
	}

	if len(cvAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all CV records in the query, and panics on error.
func (q cvQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all CV records in the query.
func (q cvQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count cv rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q cvQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q cvQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if cv exists")
	}

	return count > 0, nil
}

// CvtermG pointed to by the foreign key.
func (o *CV) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *CV) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cv_id=$1", o.CVID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// CvpropG pointed to by the foreign key.
func (o *CV) CvpropG(mods ...qm.QueryMod) cvpropQuery {
	return o.Cvprop(boil.GetDB(), mods...)
}

// Cvprop pointed to by the foreign key.
func (o *CV) Cvprop(exec boil.Executor, mods ...qm.QueryMod) cvpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cv_id=$1", o.CVID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvprop\"")

	return query
}

// CvtermpathsG retrieves all the cvtermpath's cvtermpath.
func (o *CV) CvtermpathsG(mods ...qm.QueryMod) cvtermpathQuery {
	return o.Cvtermpaths(boil.GetDB(), mods...)
}

// Cvtermpaths retrieves all the cvtermpath's cvtermpath with an executor.
func (o *CV) Cvtermpaths(exec boil.Executor, mods ...qm.QueryMod) cvtermpathQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"cv_id\"=$1", o.CVID),
	)

	query := Cvtermpaths(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvtermpath\" as \"a\"")
	return query
}

// LoadCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvL) LoadCvterm(e boil.Executor, singular bool, maybeCV interface{}) error {
	var slice []*CV
	var object *CV

	count := 1
	if singular {
		object = maybeCV.(*CV)
	} else {
		slice = *maybeCV.(*CVSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvR{}
		args[0] = object.CVID
	} else {
		for i, obj := range slice {
			obj.R = &cvR{}
			args[i] = obj.CVID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm\" where \"cv_id\" in (%s)",
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

	if len(cvAfterSelectHooks) != 0 {
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
			if local.CVID == foreign.CVID {
				local.R.Cvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadCvprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvL) LoadCvprop(e boil.Executor, singular bool, maybeCV interface{}) error {
	var slice []*CV
	var object *CV

	count := 1
	if singular {
		object = maybeCV.(*CV)
	} else {
		slice = *maybeCV.(*CVSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvR{}
		args[0] = object.CVID
	} else {
		for i, obj := range slice {
			obj.R = &cvR{}
			args[i] = obj.CVID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvprop\" where \"cv_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvprop")
	}
	defer results.Close()

	var resultSlice []*Cvprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvprop")
	}

	if len(cvAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Cvprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CVID == foreign.CVID {
				local.R.Cvprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadCvtermpaths allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvL) LoadCvtermpaths(e boil.Executor, singular bool, maybeCV interface{}) error {
	var slice []*CV
	var object *CV

	count := 1
	if singular {
		object = maybeCV.(*CV)
	} else {
		slice = *maybeCV.(*CVSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvR{}
		args[0] = object.CVID
	} else {
		for i, obj := range slice {
			obj.R = &cvR{}
			args[i] = obj.CVID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvtermpath\" where \"cv_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load cvtermpath")
	}
	defer results.Close()

	var resultSlice []*Cvtermpath
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice cvtermpath")
	}

	if len(cvtermpathAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Cvtermpaths = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CVID == foreign.CVID {
				local.R.Cvtermpaths = append(local.R.Cvtermpaths, foreign)
				break
			}
		}
	}

	return nil
}

// SetCvterm of the cv to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.CV.
func (o *CV) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error

	if insert {
		related.CVID = o.CVID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cv_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.CVID, related.CvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CVID = o.CVID

	}

	if o.R == nil {
		o.R = &cvR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			CV: o,
		}
	} else {
		related.R.CV = o
	}
	return nil
}

// SetCvprop of the cv to the related item.
// Sets o.R.Cvprop to related.
// Adds o to related.R.CV.
func (o *CV) SetCvprop(exec boil.Executor, insert bool, related *Cvprop) error {
	var err error

	if insert {
		related.CVID = o.CVID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cv_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CVID, related.CvpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CVID = o.CVID

	}

	if o.R == nil {
		o.R = &cvR{
			Cvprop: related,
		}
	} else {
		o.R.Cvprop = related
	}

	if related.R == nil {
		related.R = &cvpropR{
			CV: o,
		}
	} else {
		related.R.CV = o
	}
	return nil
}

// AddCvtermpaths adds the given related objects to the existing relationships
// of the cv, optionally inserting them as new records.
// Appends related to o.R.Cvtermpaths.
// Sets related.R.CV appropriately.
func (o *CV) AddCvtermpaths(exec boil.Executor, insert bool, related ...*Cvtermpath) error {
	var err error
	for _, rel := range related {
		rel.CVID = o.CVID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "cv_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvR{
			Cvtermpaths: related,
		}
	} else {
		o.R.Cvtermpaths = append(o.R.Cvtermpaths, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &cvtermpathR{
				CV: o,
			}
		} else {
			rel.R.CV = o
		}
	}
	return nil
}

// CVSG retrieves all records.
func CVSG(mods ...qm.QueryMod) cvQuery {
	return CVS(boil.GetDB(), mods...)
}

// CVS retrieves all the records using an executor.
func CVS(exec boil.Executor, mods ...qm.QueryMod) cvQuery {
	mods = append(mods, qm.From("\"cv\""))
	return cvQuery{NewQuery(exec, mods...)}
}

// FindCVG retrieves a single record by ID.
func FindCVG(cvID int, selectCols ...string) (*CV, error) {
	return FindCV(boil.GetDB(), cvID, selectCols...)
}

// FindCVGP retrieves a single record by ID, and panics on error.
func FindCVGP(cvID int, selectCols ...string) *CV {
	retobj, err := FindCV(boil.GetDB(), cvID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCV retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCV(exec boil.Executor, cvID int, selectCols ...string) (*CV, error) {
	cvObj := &CV{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cv\" where \"cv_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, cvID)

	err := q.Bind(cvObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from cv")
	}

	return cvObj, nil
}

// FindCVP retrieves a single record by ID with an executor, and panics on error.
func FindCVP(exec boil.Executor, cvID int, selectCols ...string) *CV {
	retobj, err := FindCV(exec, cvID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CV) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *CV) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *CV) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *CV) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no cv provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	cvInsertCacheMut.RLock()
	cache, cached := cvInsertCache[key]
	cvInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			cvColumns,
			cvColumnsWithDefault,
			cvColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(cvType, cvMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cvType, cvMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"cv\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into cv")
	}

	if !cached {
		cvInsertCacheMut.Lock()
		cvInsertCache[key] = cache
		cvInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single CV record. See Update for
// whitelist behavior description.
func (o *CV) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single CV record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *CV) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the CV, and panics on error.
// See Update for whitelist behavior description.
func (o *CV) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the CV.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *CV) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	cvUpdateCacheMut.RLock()
	cache, cached := cvUpdateCache[key]
	cvUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(cvColumns, cvPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update cv, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cv\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cvPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cvType, cvMapping, append(wl, cvPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update cv row")
	}

	if !cached {
		cvUpdateCacheMut.Lock()
		cvUpdateCache[key] = cache
		cvUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q cvQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q cvQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for cv")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CVSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CVSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CVSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CVSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"cv\" SET %s WHERE (\"cv_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvPrimaryKeyColumns), len(colNames)+1, len(cvPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in cv slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CV) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *CV) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *CV) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *CV) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no cv provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvColumnsWithDefault, o)

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

	cvUpsertCacheMut.RLock()
	cache, cached := cvUpsertCache[key]
	cvUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			cvColumns,
			cvColumnsWithDefault,
			cvColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			cvColumns,
			cvPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert cv, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cvPrimaryKeyColumns))
			copy(conflict, cvPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"cv\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(cvType, cvMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cvType, cvMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for cv")
	}

	if !cached {
		cvUpsertCacheMut.Lock()
		cvUpsertCache[key] = cache
		cvUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single CV record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CV) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single CV record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CV) DeleteG() error {
	if o == nil {
		return errors.New("chado: no CV provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single CV record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CV) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single CV record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CV) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no CV provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cvPrimaryKeyMapping)
	sql := "DELETE FROM \"cv\" WHERE \"cv_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from cv")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q cvQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q cvQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no cvQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from cv")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CVSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CVSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no CV slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CVSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CVSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no CV slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(cvBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"cv\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvPrimaryKeyColumns), 1, len(cvPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from cv slice")
	}

	if len(cvAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *CV) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *CV) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CV) ReloadG() error {
	if o == nil {
		return errors.New("chado: no CV provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CV) Reload(exec boil.Executor) error {
	ret, err := FindCV(exec, o.CVID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CVSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CVSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CVSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty CVSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CVSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	cvs := CVSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"cv\".* FROM \"cv\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(cvPrimaryKeyColumns), 1, len(cvPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&cvs)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in CVSlice")
	}

	*o = cvs

	return nil
}

// CVExists checks if the CV row exists.
func CVExists(exec boil.Executor, cvID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"cv\" where \"cv_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, cvID)
	}

	row := exec.QueryRow(sql, cvID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if cv exists")
	}

	return exists, nil
}

// CVExistsG checks if the CV row exists.
func CVExistsG(cvID int) (bool, error) {
	return CVExists(boil.GetDB(), cvID)
}

// CVExistsGP checks if the CV row exists. Panics on error.
func CVExistsGP(cvID int) bool {
	e, err := CVExists(boil.GetDB(), cvID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CVExistsP checks if the CV row exists. Panics on error.
func CVExistsP(exec boil.Executor, cvID int) bool {
	e, err := CVExists(exec, cvID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
