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
	"gopkg.in/nullbio/null.v5"
)

// Analysi is an object representing the database table.
type Analysi struct {
	AnalysisID     int         `boil:"analysis_id" json:"analysis_id" toml:"analysis_id" yaml:"analysis_id"`
	Name           null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Description    null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Program        string      `boil:"program" json:"program" toml:"program" yaml:"program"`
	Programversion string      `boil:"programversion" json:"programversion" toml:"programversion" yaml:"programversion"`
	Algorithm      null.String `boil:"algorithm" json:"algorithm,omitempty" toml:"algorithm" yaml:"algorithm,omitempty"`
	Sourcename     null.String `boil:"sourcename" json:"sourcename,omitempty" toml:"sourcename" yaml:"sourcename,omitempty"`
	Sourceversion  null.String `boil:"sourceversion" json:"sourceversion,omitempty" toml:"sourceversion" yaml:"sourceversion,omitempty"`
	Sourceuri      null.String `boil:"sourceuri" json:"sourceuri,omitempty" toml:"sourceuri" yaml:"sourceuri,omitempty"`
	Timeexecuted   time.Time   `boil:"timeexecuted" json:"timeexecuted" toml:"timeexecuted" yaml:"timeexecuted"`

	R *analysiR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L analysiL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// analysiR is where relationships are stored.
type analysiR struct {
	Analysisfeature *Analysisfeature
	Analysisprop    *Analysisprop
}

// analysiL is where Load methods for each relationship are stored.
type analysiL struct{}

var (
	analysiColumns               = []string{"analysis_id", "name", "description", "program", "programversion", "algorithm", "sourcename", "sourceversion", "sourceuri", "timeexecuted"}
	analysiColumnsWithoutDefault = []string{"name", "description", "program", "programversion", "algorithm", "sourcename", "sourceversion", "sourceuri"}
	analysiColumnsWithDefault    = []string{"analysis_id", "timeexecuted"}
	analysiPrimaryKeyColumns     = []string{"analysis_id"}
)

type (
	// AnalysiSlice is an alias for a slice of pointers to Analysi.
	// This should generally be used opposed to []Analysi.
	AnalysiSlice []*Analysi
	// AnalysiHook is the signature for custom Analysi hook methods
	AnalysiHook func(boil.Executor, *Analysi) error

	analysiQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	analysiType                 = reflect.TypeOf(&Analysi{})
	analysiMapping              = queries.MakeStructMapping(analysiType)
	analysiPrimaryKeyMapping, _ = queries.BindMapping(analysiType, analysiMapping, analysiPrimaryKeyColumns)
	analysiInsertCacheMut       sync.RWMutex
	analysiInsertCache          = make(map[string]insertCache)
	analysiUpdateCacheMut       sync.RWMutex
	analysiUpdateCache          = make(map[string]updateCache)
	analysiUpsertCacheMut       sync.RWMutex
	analysiUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var analysiBeforeInsertHooks []AnalysiHook
var analysiBeforeUpdateHooks []AnalysiHook
var analysiBeforeDeleteHooks []AnalysiHook
var analysiBeforeUpsertHooks []AnalysiHook

var analysiAfterInsertHooks []AnalysiHook
var analysiAfterSelectHooks []AnalysiHook
var analysiAfterUpdateHooks []AnalysiHook
var analysiAfterDeleteHooks []AnalysiHook
var analysiAfterUpsertHooks []AnalysiHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Analysi) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysiBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Analysi) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range analysiBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Analysi) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range analysiBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Analysi) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysiBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Analysi) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysiAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Analysi) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range analysiAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Analysi) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range analysiAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Analysi) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range analysiAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Analysi) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysiAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnalysiHook registers your hook function for all future operations.
func AddAnalysiHook(hookPoint boil.HookPoint, analysiHook AnalysiHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		analysiBeforeInsertHooks = append(analysiBeforeInsertHooks, analysiHook)
	case boil.BeforeUpdateHook:
		analysiBeforeUpdateHooks = append(analysiBeforeUpdateHooks, analysiHook)
	case boil.BeforeDeleteHook:
		analysiBeforeDeleteHooks = append(analysiBeforeDeleteHooks, analysiHook)
	case boil.BeforeUpsertHook:
		analysiBeforeUpsertHooks = append(analysiBeforeUpsertHooks, analysiHook)
	case boil.AfterInsertHook:
		analysiAfterInsertHooks = append(analysiAfterInsertHooks, analysiHook)
	case boil.AfterSelectHook:
		analysiAfterSelectHooks = append(analysiAfterSelectHooks, analysiHook)
	case boil.AfterUpdateHook:
		analysiAfterUpdateHooks = append(analysiAfterUpdateHooks, analysiHook)
	case boil.AfterDeleteHook:
		analysiAfterDeleteHooks = append(analysiAfterDeleteHooks, analysiHook)
	case boil.AfterUpsertHook:
		analysiAfterUpsertHooks = append(analysiAfterUpsertHooks, analysiHook)
	}
}

// OneP returns a single analysi record from the query, and panics on error.
func (q analysiQuery) OneP() *Analysi {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single analysi record from the query.
func (q analysiQuery) One() (*Analysi, error) {
	o := &Analysi{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for analysis")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Analysi records from the query, and panics on error.
func (q analysiQuery) AllP() AnalysiSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Analysi records from the query.
func (q analysiQuery) All() (AnalysiSlice, error) {
	var o AnalysiSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Analysi slice")
	}

	if len(analysiAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Analysi records in the query, and panics on error.
func (q analysiQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Analysi records in the query.
func (q analysiQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count analysis rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q analysiQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q analysiQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if analysis exists")
	}

	return count > 0, nil
}

// AnalysisfeatureG pointed to by the foreign key.
func (o *Analysi) AnalysisfeatureG(mods ...qm.QueryMod) analysisfeatureQuery {
	return o.Analysisfeature(boil.GetDB(), mods...)
}

// Analysisfeature pointed to by the foreign key.
func (o *Analysi) Analysisfeature(exec boil.Executor, mods ...qm.QueryMod) analysisfeatureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("analysis_id=$1", o.AnalysisID),
	}

	queryMods = append(queryMods, mods...)

	query := Analysisfeatures(exec, queryMods...)
	queries.SetFrom(query.Query, "\"analysisfeature\"")

	return query
}

// AnalysispropG pointed to by the foreign key.
func (o *Analysi) AnalysispropG(mods ...qm.QueryMod) analysispropQuery {
	return o.Analysisprop(boil.GetDB(), mods...)
}

// Analysisprop pointed to by the foreign key.
func (o *Analysi) Analysisprop(exec boil.Executor, mods ...qm.QueryMod) analysispropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("analysis_id=$1", o.AnalysisID),
	}

	queryMods = append(queryMods, mods...)

	query := Analysisprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"analysisprop\"")

	return query
}

// LoadAnalysisfeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (analysiL) LoadAnalysisfeature(e boil.Executor, singular bool, maybeAnalysi interface{}) error {
	var slice []*Analysi
	var object *Analysi

	count := 1
	if singular {
		object = maybeAnalysi.(*Analysi)
	} else {
		slice = *maybeAnalysi.(*AnalysiSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &analysiR{}
		args[0] = object.AnalysisID
	} else {
		for i, obj := range slice {
			obj.R = &analysiR{}
			args[i] = obj.AnalysisID
		}
	}

	query := fmt.Sprintf(
		"select * from \"analysisfeature\" where \"analysis_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Analysisfeature")
	}
	defer results.Close()

	var resultSlice []*Analysisfeature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Analysisfeature")
	}

	if len(analysiAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Analysisfeature = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AnalysisID == foreign.AnalysisID {
				local.R.Analysisfeature = foreign
				break
			}
		}
	}

	return nil
}

// LoadAnalysisprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (analysiL) LoadAnalysisprop(e boil.Executor, singular bool, maybeAnalysi interface{}) error {
	var slice []*Analysi
	var object *Analysi

	count := 1
	if singular {
		object = maybeAnalysi.(*Analysi)
	} else {
		slice = *maybeAnalysi.(*AnalysiSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &analysiR{}
		args[0] = object.AnalysisID
	} else {
		for i, obj := range slice {
			obj.R = &analysiR{}
			args[i] = obj.AnalysisID
		}
	}

	query := fmt.Sprintf(
		"select * from \"analysisprop\" where \"analysis_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Analysisprop")
	}
	defer results.Close()

	var resultSlice []*Analysisprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Analysisprop")
	}

	if len(analysiAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Analysisprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AnalysisID == foreign.AnalysisID {
				local.R.Analysisprop = foreign
				break
			}
		}
	}

	return nil
}

// SetAnalysisfeature of the analysi to the related item.
// Sets o.R.Analysisfeature to related.
// Adds o to related.R.Analysi.
func (o *Analysi) SetAnalysisfeature(exec boil.Executor, insert bool, related *Analysisfeature) error {
	var err error

	if insert {
		related.AnalysisID = o.AnalysisID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"analysisfeature\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"analysis_id"}),
			strmangle.WhereClause("\"", "\"", 2, analysisfeaturePrimaryKeyColumns),
		)
		values := []interface{}{o.AnalysisID, related.AnalysisfeatureID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.AnalysisID = o.AnalysisID

	}

	if o.R == nil {
		o.R = &analysiR{
			Analysisfeature: related,
		}
	} else {
		o.R.Analysisfeature = related
	}

	if related.R == nil {
		related.R = &analysisfeatureR{
			Analysi: o,
		}
	} else {
		related.R.Analysi = o
	}
	return nil
}

// SetAnalysisprop of the analysi to the related item.
// Sets o.R.Analysisprop to related.
// Adds o to related.R.Analysi.
func (o *Analysi) SetAnalysisprop(exec boil.Executor, insert bool, related *Analysisprop) error {
	var err error

	if insert {
		related.AnalysisID = o.AnalysisID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"analysisprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"analysis_id"}),
			strmangle.WhereClause("\"", "\"", 2, analysispropPrimaryKeyColumns),
		)
		values := []interface{}{o.AnalysisID, related.AnalysispropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.AnalysisID = o.AnalysisID

	}

	if o.R == nil {
		o.R = &analysiR{
			Analysisprop: related,
		}
	} else {
		o.R.Analysisprop = related
	}

	if related.R == nil {
		related.R = &analysispropR{
			Analysi: o,
		}
	} else {
		related.R.Analysi = o
	}
	return nil
}

// AnalysesG retrieves all records.
func AnalysesG(mods ...qm.QueryMod) analysiQuery {
	return Analyses(boil.GetDB(), mods...)
}

// Analyses retrieves all the records using an executor.
func Analyses(exec boil.Executor, mods ...qm.QueryMod) analysiQuery {
	mods = append(mods, qm.From("\"analysis\""))
	return analysiQuery{NewQuery(exec, mods...)}
}

// FindAnalysiG retrieves a single record by ID.
func FindAnalysiG(analysisID int, selectCols ...string) (*Analysi, error) {
	return FindAnalysi(boil.GetDB(), analysisID, selectCols...)
}

// FindAnalysiGP retrieves a single record by ID, and panics on error.
func FindAnalysiGP(analysisID int, selectCols ...string) *Analysi {
	retobj, err := FindAnalysi(boil.GetDB(), analysisID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAnalysi retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnalysi(exec boil.Executor, analysisID int, selectCols ...string) (*Analysi, error) {
	analysiObj := &Analysi{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"analysis\" where \"analysis_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, analysisID)

	err := q.Bind(analysiObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from analysis")
	}

	return analysiObj, nil
}

// FindAnalysiP retrieves a single record by ID with an executor, and panics on error.
func FindAnalysiP(exec boil.Executor, analysisID int, selectCols ...string) *Analysi {
	retobj, err := FindAnalysi(exec, analysisID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Analysi) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Analysi) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Analysi) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Analysi) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no analysis provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(analysiColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	analysiInsertCacheMut.RLock()
	cache, cached := analysiInsertCache[key]
	analysiInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			analysiColumns,
			analysiColumnsWithDefault,
			analysiColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(analysiType, analysiMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(analysiType, analysiMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"analysis\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into analysis")
	}

	if !cached {
		analysiInsertCacheMut.Lock()
		analysiInsertCache[key] = cache
		analysiInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Analysi record. See Update for
// whitelist behavior description.
func (o *Analysi) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Analysi record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Analysi) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Analysi, and panics on error.
// See Update for whitelist behavior description.
func (o *Analysi) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Analysi.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Analysi) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	analysiUpdateCacheMut.RLock()
	cache, cached := analysiUpdateCache[key]
	analysiUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(analysiColumns, analysiPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update analysis, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"analysis\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, analysiPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(analysiType, analysiMapping, append(wl, analysiPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update analysis row")
	}

	if !cached {
		analysiUpdateCacheMut.Lock()
		analysiUpdateCache[key] = cache
		analysiUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q analysiQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q analysiQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for analysis")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AnalysiSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AnalysiSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AnalysiSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnalysiSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"analysis\" SET %s WHERE (\"analysis_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(analysiPrimaryKeyColumns), len(colNames)+1, len(analysiPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in analysi slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Analysi) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Analysi) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Analysi) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Analysi) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no analysis provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(analysiColumnsWithDefault, o)

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

	analysiUpsertCacheMut.RLock()
	cache, cached := analysiUpsertCache[key]
	analysiUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			analysiColumns,
			analysiColumnsWithDefault,
			analysiColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			analysiColumns,
			analysiPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert analysis, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(analysiPrimaryKeyColumns))
			copy(conflict, analysiPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"analysis\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(analysiType, analysiMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(analysiType, analysiMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for analysis")
	}

	if !cached {
		analysiUpsertCacheMut.Lock()
		analysiUpsertCache[key] = cache
		analysiUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Analysi record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Analysi) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Analysi record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Analysi) DeleteG() error {
	if o == nil {
		return errors.New("models: no Analysi provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Analysi record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Analysi) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Analysi record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Analysi) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Analysi provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), analysiPrimaryKeyMapping)
	sql := "DELETE FROM \"analysis\" WHERE \"analysis_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from analysis")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q analysiQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q analysiQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no analysiQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from analysis")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AnalysiSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AnalysiSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Analysi slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AnalysiSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnalysiSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Analysi slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(analysiBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"analysis\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, analysiPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(analysiPrimaryKeyColumns), 1, len(analysiPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from analysi slice")
	}

	if len(analysiAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Analysi) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Analysi) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Analysi) ReloadG() error {
	if o == nil {
		return errors.New("models: no Analysi provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Analysi) Reload(exec boil.Executor) error {
	ret, err := FindAnalysi(exec, o.AnalysisID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AnalysiSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AnalysiSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnalysiSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AnalysiSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnalysiSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	analyses := AnalysiSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"analysis\".* FROM \"analysis\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, analysiPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(analysiPrimaryKeyColumns), 1, len(analysiPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&analyses)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AnalysiSlice")
	}

	*o = analyses

	return nil
}

// AnalysiExists checks if the Analysi row exists.
func AnalysiExists(exec boil.Executor, analysisID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"analysis\" where \"analysis_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, analysisID)
	}

	row := exec.QueryRow(sql, analysisID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if analysis exists")
	}

	return exists, nil
}

// AnalysiExistsG checks if the Analysi row exists.
func AnalysiExistsG(analysisID int) (bool, error) {
	return AnalysiExists(boil.GetDB(), analysisID)
}

// AnalysiExistsGP checks if the Analysi row exists. Panics on error.
func AnalysiExistsGP(analysisID int) bool {
	e, err := AnalysiExists(boil.GetDB(), analysisID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AnalysiExistsP checks if the Analysi row exists. Panics on error.
func AnalysiExistsP(exec boil.Executor, analysisID int) bool {
	e, err := AnalysiExists(exec, analysisID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
