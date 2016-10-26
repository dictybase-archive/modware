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
)

// EnvironmentCvterm is an object representing the database table.
type EnvironmentCvterm struct {
	EnvironmentCvtermID int `boil:"environment_cvterm_id" json:"environment_cvterm_id" toml:"environment_cvterm_id" yaml:"environment_cvterm_id"`
	EnvironmentID       int `boil:"environment_id" json:"environment_id" toml:"environment_id" yaml:"environment_id"`
	CvtermID            int `boil:"cvterm_id" json:"cvterm_id" toml:"cvterm_id" yaml:"cvterm_id"`

	R *environmentCvtermR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L environmentCvtermL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// environmentCvtermR is where relationships are stored.
type environmentCvtermR struct {
	Environment *Environment
	Cvterm      *Cvterm
}

// environmentCvtermL is where Load methods for each relationship are stored.
type environmentCvtermL struct{}

var (
	environmentCvtermColumns               = []string{"environment_cvterm_id", "environment_id", "cvterm_id"}
	environmentCvtermColumnsWithoutDefault = []string{"environment_id", "cvterm_id"}
	environmentCvtermColumnsWithDefault    = []string{"environment_cvterm_id"}
	environmentCvtermPrimaryKeyColumns     = []string{"environment_cvterm_id"}
)

type (
	// EnvironmentCvtermSlice is an alias for a slice of pointers to EnvironmentCvterm.
	// This should generally be used opposed to []EnvironmentCvterm.
	EnvironmentCvtermSlice []*EnvironmentCvterm
	// EnvironmentCvtermHook is the signature for custom EnvironmentCvterm hook methods
	EnvironmentCvtermHook func(boil.Executor, *EnvironmentCvterm) error

	environmentCvtermQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	environmentCvtermType                 = reflect.TypeOf(&EnvironmentCvterm{})
	environmentCvtermMapping              = queries.MakeStructMapping(environmentCvtermType)
	environmentCvtermPrimaryKeyMapping, _ = queries.BindMapping(environmentCvtermType, environmentCvtermMapping, environmentCvtermPrimaryKeyColumns)
	environmentCvtermInsertCacheMut       sync.RWMutex
	environmentCvtermInsertCache          = make(map[string]insertCache)
	environmentCvtermUpdateCacheMut       sync.RWMutex
	environmentCvtermUpdateCache          = make(map[string]updateCache)
	environmentCvtermUpsertCacheMut       sync.RWMutex
	environmentCvtermUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var environmentCvtermBeforeInsertHooks []EnvironmentCvtermHook
var environmentCvtermBeforeUpdateHooks []EnvironmentCvtermHook
var environmentCvtermBeforeDeleteHooks []EnvironmentCvtermHook
var environmentCvtermBeforeUpsertHooks []EnvironmentCvtermHook

var environmentCvtermAfterInsertHooks []EnvironmentCvtermHook
var environmentCvtermAfterSelectHooks []EnvironmentCvtermHook
var environmentCvtermAfterUpdateHooks []EnvironmentCvtermHook
var environmentCvtermAfterDeleteHooks []EnvironmentCvtermHook
var environmentCvtermAfterUpsertHooks []EnvironmentCvtermHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EnvironmentCvterm) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentCvtermBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EnvironmentCvterm) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentCvtermBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EnvironmentCvterm) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentCvtermBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EnvironmentCvterm) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentCvtermBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EnvironmentCvterm) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentCvtermAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EnvironmentCvterm) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentCvtermAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EnvironmentCvterm) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentCvtermAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EnvironmentCvterm) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentCvtermAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EnvironmentCvterm) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentCvtermAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEnvironmentCvtermHook registers your hook function for all future operations.
func AddEnvironmentCvtermHook(hookPoint boil.HookPoint, environmentCvtermHook EnvironmentCvtermHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		environmentCvtermBeforeInsertHooks = append(environmentCvtermBeforeInsertHooks, environmentCvtermHook)
	case boil.BeforeUpdateHook:
		environmentCvtermBeforeUpdateHooks = append(environmentCvtermBeforeUpdateHooks, environmentCvtermHook)
	case boil.BeforeDeleteHook:
		environmentCvtermBeforeDeleteHooks = append(environmentCvtermBeforeDeleteHooks, environmentCvtermHook)
	case boil.BeforeUpsertHook:
		environmentCvtermBeforeUpsertHooks = append(environmentCvtermBeforeUpsertHooks, environmentCvtermHook)
	case boil.AfterInsertHook:
		environmentCvtermAfterInsertHooks = append(environmentCvtermAfterInsertHooks, environmentCvtermHook)
	case boil.AfterSelectHook:
		environmentCvtermAfterSelectHooks = append(environmentCvtermAfterSelectHooks, environmentCvtermHook)
	case boil.AfterUpdateHook:
		environmentCvtermAfterUpdateHooks = append(environmentCvtermAfterUpdateHooks, environmentCvtermHook)
	case boil.AfterDeleteHook:
		environmentCvtermAfterDeleteHooks = append(environmentCvtermAfterDeleteHooks, environmentCvtermHook)
	case boil.AfterUpsertHook:
		environmentCvtermAfterUpsertHooks = append(environmentCvtermAfterUpsertHooks, environmentCvtermHook)
	}
}

// OneP returns a single environmentCvterm record from the query, and panics on error.
func (q environmentCvtermQuery) OneP() *EnvironmentCvterm {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single environmentCvterm record from the query.
func (q environmentCvtermQuery) One() (*EnvironmentCvterm, error) {
	o := &EnvironmentCvterm{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for environment_cvterm")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all EnvironmentCvterm records from the query, and panics on error.
func (q environmentCvtermQuery) AllP() EnvironmentCvtermSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all EnvironmentCvterm records from the query.
func (q environmentCvtermQuery) All() (EnvironmentCvtermSlice, error) {
	var o EnvironmentCvtermSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to EnvironmentCvterm slice")
	}

	if len(environmentCvtermAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all EnvironmentCvterm records in the query, and panics on error.
func (q environmentCvtermQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all EnvironmentCvterm records in the query.
func (q environmentCvtermQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count environment_cvterm rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q environmentCvtermQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q environmentCvtermQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if environment_cvterm exists")
	}

	return count > 0, nil
}

// EnvironmentG pointed to by the foreign key.
func (o *EnvironmentCvterm) EnvironmentG(mods ...qm.QueryMod) environmentQuery {
	return o.Environment(boil.GetDB(), mods...)
}

// Environment pointed to by the foreign key.
func (o *EnvironmentCvterm) Environment(exec boil.Executor, mods ...qm.QueryMod) environmentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("environment_id=$1", o.EnvironmentID),
	}

	queryMods = append(queryMods, mods...)

	query := Environments(exec, queryMods...)
	queries.SetFrom(query.Query, "\"environment\"")

	return query
}

// CvtermG pointed to by the foreign key.
func (o *EnvironmentCvterm) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *EnvironmentCvterm) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadEnvironment allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (environmentCvtermL) LoadEnvironment(e boil.Executor, singular bool, maybeEnvironmentCvterm interface{}) error {
	var slice []*EnvironmentCvterm
	var object *EnvironmentCvterm

	count := 1
	if singular {
		object = maybeEnvironmentCvterm.(*EnvironmentCvterm)
	} else {
		slice = *maybeEnvironmentCvterm.(*EnvironmentCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &environmentCvtermR{}
		args[0] = object.EnvironmentID
	} else {
		for i, obj := range slice {
			obj.R = &environmentCvtermR{}
			args[i] = obj.EnvironmentID
		}
	}

	query := fmt.Sprintf(
		"select * from \"environment\" where \"environment_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Environment")
	}
	defer results.Close()

	var resultSlice []*Environment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Environment")
	}

	if len(environmentCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Environment = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.EnvironmentID == foreign.EnvironmentID {
				local.R.Environment = foreign
				break
			}
		}
	}

	return nil
}

// LoadCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (environmentCvtermL) LoadCvterm(e boil.Executor, singular bool, maybeEnvironmentCvterm interface{}) error {
	var slice []*EnvironmentCvterm
	var object *EnvironmentCvterm

	count := 1
	if singular {
		object = maybeEnvironmentCvterm.(*EnvironmentCvterm)
	} else {
		slice = *maybeEnvironmentCvterm.(*EnvironmentCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &environmentCvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &environmentCvtermR{}
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

	if len(environmentCvtermAfterSelectHooks) != 0 {
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

// SetEnvironment of the environment_cvterm to the related item.
// Sets o.R.Environment to related.
// Adds o to related.R.EnvironmentCvterm.
func (o *EnvironmentCvterm) SetEnvironment(exec boil.Executor, insert bool, related *Environment) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"environment_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"environment_id"}),
		strmangle.WhereClause("\"", "\"", 2, environmentCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.EnvironmentID, o.EnvironmentCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EnvironmentID = related.EnvironmentID

	if o.R == nil {
		o.R = &environmentCvtermR{
			Environment: related,
		}
	} else {
		o.R.Environment = related
	}

	if related.R == nil {
		related.R = &environmentR{
			EnvironmentCvterm: o,
		}
	} else {
		related.R.EnvironmentCvterm = o
	}

	return nil
}

// SetCvterm of the environment_cvterm to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.EnvironmentCvterm.
func (o *EnvironmentCvterm) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"environment_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, environmentCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.EnvironmentCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CvtermID = related.CvtermID

	if o.R == nil {
		o.R = &environmentCvtermR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			EnvironmentCvterm: o,
		}
	} else {
		related.R.EnvironmentCvterm = o
	}

	return nil
}

// EnvironmentCvtermsG retrieves all records.
func EnvironmentCvtermsG(mods ...qm.QueryMod) environmentCvtermQuery {
	return EnvironmentCvterms(boil.GetDB(), mods...)
}

// EnvironmentCvterms retrieves all the records using an executor.
func EnvironmentCvterms(exec boil.Executor, mods ...qm.QueryMod) environmentCvtermQuery {
	mods = append(mods, qm.From("\"environment_cvterm\""))
	return environmentCvtermQuery{NewQuery(exec, mods...)}
}

// FindEnvironmentCvtermG retrieves a single record by ID.
func FindEnvironmentCvtermG(environmentCvtermID int, selectCols ...string) (*EnvironmentCvterm, error) {
	return FindEnvironmentCvterm(boil.GetDB(), environmentCvtermID, selectCols...)
}

// FindEnvironmentCvtermGP retrieves a single record by ID, and panics on error.
func FindEnvironmentCvtermGP(environmentCvtermID int, selectCols ...string) *EnvironmentCvterm {
	retobj, err := FindEnvironmentCvterm(boil.GetDB(), environmentCvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindEnvironmentCvterm retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEnvironmentCvterm(exec boil.Executor, environmentCvtermID int, selectCols ...string) (*EnvironmentCvterm, error) {
	environmentCvtermObj := &EnvironmentCvterm{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"environment_cvterm\" where \"environment_cvterm_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, environmentCvtermID)

	err := q.Bind(environmentCvtermObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from environment_cvterm")
	}

	return environmentCvtermObj, nil
}

// FindEnvironmentCvtermP retrieves a single record by ID with an executor, and panics on error.
func FindEnvironmentCvtermP(exec boil.Executor, environmentCvtermID int, selectCols ...string) *EnvironmentCvterm {
	retobj, err := FindEnvironmentCvterm(exec, environmentCvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *EnvironmentCvterm) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *EnvironmentCvterm) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *EnvironmentCvterm) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *EnvironmentCvterm) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no environment_cvterm provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(environmentCvtermColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	environmentCvtermInsertCacheMut.RLock()
	cache, cached := environmentCvtermInsertCache[key]
	environmentCvtermInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			environmentCvtermColumns,
			environmentCvtermColumnsWithDefault,
			environmentCvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(environmentCvtermType, environmentCvtermMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(environmentCvtermType, environmentCvtermMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"environment_cvterm\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into environment_cvterm")
	}

	if !cached {
		environmentCvtermInsertCacheMut.Lock()
		environmentCvtermInsertCache[key] = cache
		environmentCvtermInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single EnvironmentCvterm record. See Update for
// whitelist behavior description.
func (o *EnvironmentCvterm) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single EnvironmentCvterm record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *EnvironmentCvterm) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the EnvironmentCvterm, and panics on error.
// See Update for whitelist behavior description.
func (o *EnvironmentCvterm) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the EnvironmentCvterm.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *EnvironmentCvterm) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	environmentCvtermUpdateCacheMut.RLock()
	cache, cached := environmentCvtermUpdateCache[key]
	environmentCvtermUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(environmentCvtermColumns, environmentCvtermPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update environment_cvterm, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"environment_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, environmentCvtermPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(environmentCvtermType, environmentCvtermMapping, append(wl, environmentCvtermPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update environment_cvterm row")
	}

	if !cached {
		environmentCvtermUpdateCacheMut.Lock()
		environmentCvtermUpdateCache[key] = cache
		environmentCvtermUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q environmentCvtermQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q environmentCvtermQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for environment_cvterm")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o EnvironmentCvtermSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o EnvironmentCvtermSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o EnvironmentCvtermSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EnvironmentCvtermSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), environmentCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"environment_cvterm\" SET %s WHERE (\"environment_cvterm_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(environmentCvtermPrimaryKeyColumns), len(colNames)+1, len(environmentCvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in environmentCvterm slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *EnvironmentCvterm) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *EnvironmentCvterm) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *EnvironmentCvterm) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *EnvironmentCvterm) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no environment_cvterm provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(environmentCvtermColumnsWithDefault, o)

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

	environmentCvtermUpsertCacheMut.RLock()
	cache, cached := environmentCvtermUpsertCache[key]
	environmentCvtermUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			environmentCvtermColumns,
			environmentCvtermColumnsWithDefault,
			environmentCvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			environmentCvtermColumns,
			environmentCvtermPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert environment_cvterm, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(environmentCvtermPrimaryKeyColumns))
			copy(conflict, environmentCvtermPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"environment_cvterm\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(environmentCvtermType, environmentCvtermMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(environmentCvtermType, environmentCvtermMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for environment_cvterm")
	}

	if !cached {
		environmentCvtermUpsertCacheMut.Lock()
		environmentCvtermUpsertCache[key] = cache
		environmentCvtermUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single EnvironmentCvterm record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *EnvironmentCvterm) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single EnvironmentCvterm record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *EnvironmentCvterm) DeleteG() error {
	if o == nil {
		return errors.New("chado: no EnvironmentCvterm provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single EnvironmentCvterm record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *EnvironmentCvterm) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single EnvironmentCvterm record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EnvironmentCvterm) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no EnvironmentCvterm provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), environmentCvtermPrimaryKeyMapping)
	sql := "DELETE FROM \"environment_cvterm\" WHERE \"environment_cvterm_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from environment_cvterm")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q environmentCvtermQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q environmentCvtermQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no environmentCvtermQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from environment_cvterm")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o EnvironmentCvtermSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o EnvironmentCvtermSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no EnvironmentCvterm slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o EnvironmentCvtermSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EnvironmentCvtermSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no EnvironmentCvterm slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(environmentCvtermBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), environmentCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"environment_cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, environmentCvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(environmentCvtermPrimaryKeyColumns), 1, len(environmentCvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from environmentCvterm slice")
	}

	if len(environmentCvtermAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *EnvironmentCvterm) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *EnvironmentCvterm) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *EnvironmentCvterm) ReloadG() error {
	if o == nil {
		return errors.New("chado: no EnvironmentCvterm provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EnvironmentCvterm) Reload(exec boil.Executor) error {
	ret, err := FindEnvironmentCvterm(exec, o.EnvironmentCvtermID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *EnvironmentCvtermSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *EnvironmentCvtermSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EnvironmentCvtermSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty EnvironmentCvtermSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EnvironmentCvtermSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	environmentCvterms := EnvironmentCvtermSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), environmentCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"environment_cvterm\".* FROM \"environment_cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, environmentCvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(environmentCvtermPrimaryKeyColumns), 1, len(environmentCvtermPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&environmentCvterms)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in EnvironmentCvtermSlice")
	}

	*o = environmentCvterms

	return nil
}

// EnvironmentCvtermExists checks if the EnvironmentCvterm row exists.
func EnvironmentCvtermExists(exec boil.Executor, environmentCvtermID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"environment_cvterm\" where \"environment_cvterm_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, environmentCvtermID)
	}

	row := exec.QueryRow(sql, environmentCvtermID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if environment_cvterm exists")
	}

	return exists, nil
}

// EnvironmentCvtermExistsG checks if the EnvironmentCvterm row exists.
func EnvironmentCvtermExistsG(environmentCvtermID int) (bool, error) {
	return EnvironmentCvtermExists(boil.GetDB(), environmentCvtermID)
}

// EnvironmentCvtermExistsGP checks if the EnvironmentCvterm row exists. Panics on error.
func EnvironmentCvtermExistsGP(environmentCvtermID int) bool {
	e, err := EnvironmentCvtermExists(boil.GetDB(), environmentCvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// EnvironmentCvtermExistsP checks if the EnvironmentCvterm row exists. Panics on error.
func EnvironmentCvtermExistsP(exec boil.Executor, environmentCvtermID int) bool {
	e, err := EnvironmentCvtermExists(exec, environmentCvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
