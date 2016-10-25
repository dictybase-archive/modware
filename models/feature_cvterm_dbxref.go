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

// FeatureCvtermDbxref is an object representing the database table.
type FeatureCvtermDbxref struct {
	FeatureCvtermDbxrefID int `boil:"feature_cvterm_dbxref_id" json:"feature_cvterm_dbxref_id" toml:"feature_cvterm_dbxref_id" yaml:"feature_cvterm_dbxref_id"`
	FeatureCvtermID       int `boil:"feature_cvterm_id" json:"feature_cvterm_id" toml:"feature_cvterm_id" yaml:"feature_cvterm_id"`
	DbxrefID              int `boil:"dbxref_id" json:"dbxref_id" toml:"dbxref_id" yaml:"dbxref_id"`

	R *featureCvtermDbxrefR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureCvtermDbxrefL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureCvtermDbxrefR is where relationships are stored.
type featureCvtermDbxrefR struct {
	FeatureCvterm *FeatureCvterm
	Dbxref        *Dbxref
}

// featureCvtermDbxrefL is where Load methods for each relationship are stored.
type featureCvtermDbxrefL struct{}

var (
	featureCvtermDbxrefColumns               = []string{"feature_cvterm_dbxref_id", "feature_cvterm_id", "dbxref_id"}
	featureCvtermDbxrefColumnsWithoutDefault = []string{"feature_cvterm_id", "dbxref_id"}
	featureCvtermDbxrefColumnsWithDefault    = []string{"feature_cvterm_dbxref_id"}
	featureCvtermDbxrefPrimaryKeyColumns     = []string{"feature_cvterm_dbxref_id"}
)

type (
	// FeatureCvtermDbxrefSlice is an alias for a slice of pointers to FeatureCvtermDbxref.
	// This should generally be used opposed to []FeatureCvtermDbxref.
	FeatureCvtermDbxrefSlice []*FeatureCvtermDbxref
	// FeatureCvtermDbxrefHook is the signature for custom FeatureCvtermDbxref hook methods
	FeatureCvtermDbxrefHook func(boil.Executor, *FeatureCvtermDbxref) error

	featureCvtermDbxrefQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureCvtermDbxrefType                 = reflect.TypeOf(&FeatureCvtermDbxref{})
	featureCvtermDbxrefMapping              = queries.MakeStructMapping(featureCvtermDbxrefType)
	featureCvtermDbxrefPrimaryKeyMapping, _ = queries.BindMapping(featureCvtermDbxrefType, featureCvtermDbxrefMapping, featureCvtermDbxrefPrimaryKeyColumns)
	featureCvtermDbxrefInsertCacheMut       sync.RWMutex
	featureCvtermDbxrefInsertCache          = make(map[string]insertCache)
	featureCvtermDbxrefUpdateCacheMut       sync.RWMutex
	featureCvtermDbxrefUpdateCache          = make(map[string]updateCache)
	featureCvtermDbxrefUpsertCacheMut       sync.RWMutex
	featureCvtermDbxrefUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureCvtermDbxrefBeforeInsertHooks []FeatureCvtermDbxrefHook
var featureCvtermDbxrefBeforeUpdateHooks []FeatureCvtermDbxrefHook
var featureCvtermDbxrefBeforeDeleteHooks []FeatureCvtermDbxrefHook
var featureCvtermDbxrefBeforeUpsertHooks []FeatureCvtermDbxrefHook

var featureCvtermDbxrefAfterInsertHooks []FeatureCvtermDbxrefHook
var featureCvtermDbxrefAfterSelectHooks []FeatureCvtermDbxrefHook
var featureCvtermDbxrefAfterUpdateHooks []FeatureCvtermDbxrefHook
var featureCvtermDbxrefAfterDeleteHooks []FeatureCvtermDbxrefHook
var featureCvtermDbxrefAfterUpsertHooks []FeatureCvtermDbxrefHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeatureCvtermDbxref) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermDbxrefBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeatureCvtermDbxref) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermDbxrefBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeatureCvtermDbxref) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermDbxrefBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeatureCvtermDbxref) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermDbxrefBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeatureCvtermDbxref) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermDbxrefAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeatureCvtermDbxref) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermDbxrefAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeatureCvtermDbxref) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermDbxrefAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeatureCvtermDbxref) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermDbxrefAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeatureCvtermDbxref) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermDbxrefAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureCvtermDbxrefHook registers your hook function for all future operations.
func AddFeatureCvtermDbxrefHook(hookPoint boil.HookPoint, featureCvtermDbxrefHook FeatureCvtermDbxrefHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureCvtermDbxrefBeforeInsertHooks = append(featureCvtermDbxrefBeforeInsertHooks, featureCvtermDbxrefHook)
	case boil.BeforeUpdateHook:
		featureCvtermDbxrefBeforeUpdateHooks = append(featureCvtermDbxrefBeforeUpdateHooks, featureCvtermDbxrefHook)
	case boil.BeforeDeleteHook:
		featureCvtermDbxrefBeforeDeleteHooks = append(featureCvtermDbxrefBeforeDeleteHooks, featureCvtermDbxrefHook)
	case boil.BeforeUpsertHook:
		featureCvtermDbxrefBeforeUpsertHooks = append(featureCvtermDbxrefBeforeUpsertHooks, featureCvtermDbxrefHook)
	case boil.AfterInsertHook:
		featureCvtermDbxrefAfterInsertHooks = append(featureCvtermDbxrefAfterInsertHooks, featureCvtermDbxrefHook)
	case boil.AfterSelectHook:
		featureCvtermDbxrefAfterSelectHooks = append(featureCvtermDbxrefAfterSelectHooks, featureCvtermDbxrefHook)
	case boil.AfterUpdateHook:
		featureCvtermDbxrefAfterUpdateHooks = append(featureCvtermDbxrefAfterUpdateHooks, featureCvtermDbxrefHook)
	case boil.AfterDeleteHook:
		featureCvtermDbxrefAfterDeleteHooks = append(featureCvtermDbxrefAfterDeleteHooks, featureCvtermDbxrefHook)
	case boil.AfterUpsertHook:
		featureCvtermDbxrefAfterUpsertHooks = append(featureCvtermDbxrefAfterUpsertHooks, featureCvtermDbxrefHook)
	}
}

// OneP returns a single featureCvtermDbxref record from the query, and panics on error.
func (q featureCvtermDbxrefQuery) OneP() *FeatureCvtermDbxref {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureCvtermDbxref record from the query.
func (q featureCvtermDbxrefQuery) One() (*FeatureCvtermDbxref, error) {
	o := &FeatureCvtermDbxref{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for feature_cvterm_dbxref")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeatureCvtermDbxref records from the query, and panics on error.
func (q featureCvtermDbxrefQuery) AllP() FeatureCvtermDbxrefSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeatureCvtermDbxref records from the query.
func (q featureCvtermDbxrefQuery) All() (FeatureCvtermDbxrefSlice, error) {
	var o FeatureCvtermDbxrefSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FeatureCvtermDbxref slice")
	}

	if len(featureCvtermDbxrefAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeatureCvtermDbxref records in the query, and panics on error.
func (q featureCvtermDbxrefQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeatureCvtermDbxref records in the query.
func (q featureCvtermDbxrefQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count feature_cvterm_dbxref rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureCvtermDbxrefQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureCvtermDbxrefQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if feature_cvterm_dbxref exists")
	}

	return count > 0, nil
}

// FeatureCvtermG pointed to by the foreign key.
func (o *FeatureCvtermDbxref) FeatureCvtermG(mods ...qm.QueryMod) featureCvtermQuery {
	return o.FeatureCvterm(boil.GetDB(), mods...)
}

// FeatureCvterm pointed to by the foreign key.
func (o *FeatureCvtermDbxref) FeatureCvterm(exec boil.Executor, mods ...qm.QueryMod) featureCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_cvterm_id=$1", o.FeatureCvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvterm\"")

	return query
}

// DbxrefG pointed to by the foreign key.
func (o *FeatureCvtermDbxref) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *FeatureCvtermDbxref) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// LoadFeatureCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermDbxrefL) LoadFeatureCvterm(e boil.Executor, singular bool, maybeFeatureCvtermDbxref interface{}) error {
	var slice []*FeatureCvtermDbxref
	var object *FeatureCvtermDbxref

	count := 1
	if singular {
		object = maybeFeatureCvtermDbxref.(*FeatureCvtermDbxref)
	} else {
		slice = *maybeFeatureCvtermDbxref.(*FeatureCvtermDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermDbxrefR{}
		args[0] = object.FeatureCvtermID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermDbxrefR{}
			args[i] = obj.FeatureCvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_cvterm\" where \"feature_cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureCvterm")
	}
	defer results.Close()

	var resultSlice []*FeatureCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureCvterm")
	}

	if len(featureCvtermDbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureCvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureCvtermID == foreign.FeatureCvtermID {
				local.R.FeatureCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermDbxrefL) LoadDbxref(e boil.Executor, singular bool, maybeFeatureCvtermDbxref interface{}) error {
	var slice []*FeatureCvtermDbxref
	var object *FeatureCvtermDbxref

	count := 1
	if singular {
		object = maybeFeatureCvtermDbxref.(*FeatureCvtermDbxref)
	} else {
		slice = *maybeFeatureCvtermDbxref.(*FeatureCvtermDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermDbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermDbxrefR{}
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

	if len(featureCvtermDbxrefAfterSelectHooks) != 0 {
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

// SetFeatureCvterm of the feature_cvterm_dbxref to the related item.
// Sets o.R.FeatureCvterm to related.
// Adds o to related.R.FeatureCvtermDbxref.
func (o *FeatureCvtermDbxref) SetFeatureCvterm(exec boil.Executor, insert bool, related *FeatureCvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_cvterm_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureCvtermDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureCvtermID, o.FeatureCvtermDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureCvtermID = related.FeatureCvtermID

	if o.R == nil {
		o.R = &featureCvtermDbxrefR{
			FeatureCvterm: related,
		}
	} else {
		o.R.FeatureCvterm = related
	}

	if related.R == nil {
		related.R = &featureCvtermR{
			FeatureCvtermDbxref: o,
		}
	} else {
		related.R.FeatureCvtermDbxref = o
	}

	return nil
}

// SetDbxref of the feature_cvterm_dbxref to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.FeatureCvtermDbxref.
func (o *FeatureCvtermDbxref) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_cvterm_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureCvtermDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.DbxrefID, o.FeatureCvtermDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DbxrefID = related.DbxrefID

	if o.R == nil {
		o.R = &featureCvtermDbxrefR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			FeatureCvtermDbxref: o,
		}
	} else {
		related.R.FeatureCvtermDbxref = o
	}

	return nil
}

// FeatureCvtermDbxrefsG retrieves all records.
func FeatureCvtermDbxrefsG(mods ...qm.QueryMod) featureCvtermDbxrefQuery {
	return FeatureCvtermDbxrefs(boil.GetDB(), mods...)
}

// FeatureCvtermDbxrefs retrieves all the records using an executor.
func FeatureCvtermDbxrefs(exec boil.Executor, mods ...qm.QueryMod) featureCvtermDbxrefQuery {
	mods = append(mods, qm.From("\"feature_cvterm_dbxref\""))
	return featureCvtermDbxrefQuery{NewQuery(exec, mods...)}
}

// FindFeatureCvtermDbxrefG retrieves a single record by ID.
func FindFeatureCvtermDbxrefG(featureCvtermDbxrefID int, selectCols ...string) (*FeatureCvtermDbxref, error) {
	return FindFeatureCvtermDbxref(boil.GetDB(), featureCvtermDbxrefID, selectCols...)
}

// FindFeatureCvtermDbxrefGP retrieves a single record by ID, and panics on error.
func FindFeatureCvtermDbxrefGP(featureCvtermDbxrefID int, selectCols ...string) *FeatureCvtermDbxref {
	retobj, err := FindFeatureCvtermDbxref(boil.GetDB(), featureCvtermDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureCvtermDbxref retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureCvtermDbxref(exec boil.Executor, featureCvtermDbxrefID int, selectCols ...string) (*FeatureCvtermDbxref, error) {
	featureCvtermDbxrefObj := &FeatureCvtermDbxref{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_cvterm_dbxref\" where \"feature_cvterm_dbxref_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureCvtermDbxrefID)

	err := q.Bind(featureCvtermDbxrefObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from feature_cvterm_dbxref")
	}

	return featureCvtermDbxrefObj, nil
}

// FindFeatureCvtermDbxrefP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureCvtermDbxrefP(exec boil.Executor, featureCvtermDbxrefID int, selectCols ...string) *FeatureCvtermDbxref {
	retobj, err := FindFeatureCvtermDbxref(exec, featureCvtermDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeatureCvtermDbxref) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeatureCvtermDbxref) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeatureCvtermDbxref) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeatureCvtermDbxref) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_cvterm_dbxref provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureCvtermDbxrefColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureCvtermDbxrefInsertCacheMut.RLock()
	cache, cached := featureCvtermDbxrefInsertCache[key]
	featureCvtermDbxrefInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureCvtermDbxrefColumns,
			featureCvtermDbxrefColumnsWithDefault,
			featureCvtermDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureCvtermDbxrefType, featureCvtermDbxrefMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureCvtermDbxrefType, featureCvtermDbxrefMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_cvterm_dbxref\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into feature_cvterm_dbxref")
	}

	if !cached {
		featureCvtermDbxrefInsertCacheMut.Lock()
		featureCvtermDbxrefInsertCache[key] = cache
		featureCvtermDbxrefInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeatureCvtermDbxref record. See Update for
// whitelist behavior description.
func (o *FeatureCvtermDbxref) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeatureCvtermDbxref record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeatureCvtermDbxref) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeatureCvtermDbxref, and panics on error.
// See Update for whitelist behavior description.
func (o *FeatureCvtermDbxref) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeatureCvtermDbxref.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeatureCvtermDbxref) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureCvtermDbxrefUpdateCacheMut.RLock()
	cache, cached := featureCvtermDbxrefUpdateCache[key]
	featureCvtermDbxrefUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureCvtermDbxrefColumns, featureCvtermDbxrefPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update feature_cvterm_dbxref, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_cvterm_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featureCvtermDbxrefPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureCvtermDbxrefType, featureCvtermDbxrefMapping, append(wl, featureCvtermDbxrefPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update feature_cvterm_dbxref row")
	}

	if !cached {
		featureCvtermDbxrefUpdateCacheMut.Lock()
		featureCvtermDbxrefUpdateCache[key] = cache
		featureCvtermDbxrefUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureCvtermDbxrefQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureCvtermDbxrefQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for feature_cvterm_dbxref")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureCvtermDbxrefSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureCvtermDbxrefSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureCvtermDbxrefSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureCvtermDbxrefSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_cvterm_dbxref\" SET %s WHERE (\"feature_cvterm_dbxref_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureCvtermDbxrefPrimaryKeyColumns), len(colNames)+1, len(featureCvtermDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in featureCvtermDbxref slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeatureCvtermDbxref) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeatureCvtermDbxref) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeatureCvtermDbxref) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeatureCvtermDbxref) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_cvterm_dbxref provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureCvtermDbxrefColumnsWithDefault, o)

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

	featureCvtermDbxrefUpsertCacheMut.RLock()
	cache, cached := featureCvtermDbxrefUpsertCache[key]
	featureCvtermDbxrefUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureCvtermDbxrefColumns,
			featureCvtermDbxrefColumnsWithDefault,
			featureCvtermDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureCvtermDbxrefColumns,
			featureCvtermDbxrefPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert feature_cvterm_dbxref, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featureCvtermDbxrefPrimaryKeyColumns))
			copy(conflict, featureCvtermDbxrefPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_cvterm_dbxref\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureCvtermDbxrefType, featureCvtermDbxrefMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureCvtermDbxrefType, featureCvtermDbxrefMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for feature_cvterm_dbxref")
	}

	if !cached {
		featureCvtermDbxrefUpsertCacheMut.Lock()
		featureCvtermDbxrefUpsertCache[key] = cache
		featureCvtermDbxrefUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeatureCvtermDbxref record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureCvtermDbxref) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeatureCvtermDbxref record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeatureCvtermDbxref) DeleteG() error {
	if o == nil {
		return errors.New("models: no FeatureCvtermDbxref provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeatureCvtermDbxref record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureCvtermDbxref) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeatureCvtermDbxref record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeatureCvtermDbxref) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeatureCvtermDbxref provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureCvtermDbxrefPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_cvterm_dbxref\" WHERE \"feature_cvterm_dbxref_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from feature_cvterm_dbxref")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureCvtermDbxrefQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureCvtermDbxrefQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no featureCvtermDbxrefQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from feature_cvterm_dbxref")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureCvtermDbxrefSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureCvtermDbxrefSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no FeatureCvtermDbxref slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureCvtermDbxrefSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureCvtermDbxrefSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeatureCvtermDbxref slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureCvtermDbxrefBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_cvterm_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureCvtermDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureCvtermDbxrefPrimaryKeyColumns), 1, len(featureCvtermDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from featureCvtermDbxref slice")
	}

	if len(featureCvtermDbxrefAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeatureCvtermDbxref) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeatureCvtermDbxref) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeatureCvtermDbxref) ReloadG() error {
	if o == nil {
		return errors.New("models: no FeatureCvtermDbxref provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeatureCvtermDbxref) Reload(exec boil.Executor) error {
	ret, err := FindFeatureCvtermDbxref(exec, o.FeatureCvtermDbxrefID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureCvtermDbxrefSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureCvtermDbxrefSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureCvtermDbxrefSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FeatureCvtermDbxrefSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureCvtermDbxrefSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureCvtermDbxrefs := FeatureCvtermDbxrefSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_cvterm_dbxref\".* FROM \"feature_cvterm_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureCvtermDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featureCvtermDbxrefPrimaryKeyColumns), 1, len(featureCvtermDbxrefPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureCvtermDbxrefs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeatureCvtermDbxrefSlice")
	}

	*o = featureCvtermDbxrefs

	return nil
}

// FeatureCvtermDbxrefExists checks if the FeatureCvtermDbxref row exists.
func FeatureCvtermDbxrefExists(exec boil.Executor, featureCvtermDbxrefID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_cvterm_dbxref\" where \"feature_cvterm_dbxref_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureCvtermDbxrefID)
	}

	row := exec.QueryRow(sql, featureCvtermDbxrefID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if feature_cvterm_dbxref exists")
	}

	return exists, nil
}

// FeatureCvtermDbxrefExistsG checks if the FeatureCvtermDbxref row exists.
func FeatureCvtermDbxrefExistsG(featureCvtermDbxrefID int) (bool, error) {
	return FeatureCvtermDbxrefExists(boil.GetDB(), featureCvtermDbxrefID)
}

// FeatureCvtermDbxrefExistsGP checks if the FeatureCvtermDbxref row exists. Panics on error.
func FeatureCvtermDbxrefExistsGP(featureCvtermDbxrefID int) bool {
	e, err := FeatureCvtermDbxrefExists(boil.GetDB(), featureCvtermDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureCvtermDbxrefExistsP checks if the FeatureCvtermDbxref row exists. Panics on error.
func FeatureCvtermDbxrefExistsP(exec boil.Executor, featureCvtermDbxrefID int) bool {
	e, err := FeatureCvtermDbxrefExists(exec, featureCvtermDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
