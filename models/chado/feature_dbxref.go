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

// FeatureDbxref is an object representing the database table.
type FeatureDbxref struct {
	FeatureDbxrefID int  `boil:"feature_dbxref_id" json:"feature_dbxref_id" toml:"feature_dbxref_id" yaml:"feature_dbxref_id"`
	FeatureID       int  `boil:"feature_id" json:"feature_id" toml:"feature_id" yaml:"feature_id"`
	DbxrefID        int  `boil:"dbxref_id" json:"dbxref_id" toml:"dbxref_id" yaml:"dbxref_id"`
	IsCurrent       bool `boil:"is_current" json:"is_current" toml:"is_current" yaml:"is_current"`

	R *featureDbxrefR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureDbxrefL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureDbxrefR is where relationships are stored.
type featureDbxrefR struct {
	Dbxref  *Dbxref
	Feature *Feature
}

// featureDbxrefL is where Load methods for each relationship are stored.
type featureDbxrefL struct{}

var (
	featureDbxrefColumns               = []string{"feature_dbxref_id", "feature_id", "dbxref_id", "is_current"}
	featureDbxrefColumnsWithoutDefault = []string{"feature_id", "dbxref_id"}
	featureDbxrefColumnsWithDefault    = []string{"feature_dbxref_id", "is_current"}
	featureDbxrefPrimaryKeyColumns     = []string{"feature_dbxref_id"}
)

type (
	// FeatureDbxrefSlice is an alias for a slice of pointers to FeatureDbxref.
	// This should generally be used opposed to []FeatureDbxref.
	FeatureDbxrefSlice []*FeatureDbxref
	// FeatureDbxrefHook is the signature for custom FeatureDbxref hook methods
	FeatureDbxrefHook func(boil.Executor, *FeatureDbxref) error

	featureDbxrefQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureDbxrefType                 = reflect.TypeOf(&FeatureDbxref{})
	featureDbxrefMapping              = queries.MakeStructMapping(featureDbxrefType)
	featureDbxrefPrimaryKeyMapping, _ = queries.BindMapping(featureDbxrefType, featureDbxrefMapping, featureDbxrefPrimaryKeyColumns)
	featureDbxrefInsertCacheMut       sync.RWMutex
	featureDbxrefInsertCache          = make(map[string]insertCache)
	featureDbxrefUpdateCacheMut       sync.RWMutex
	featureDbxrefUpdateCache          = make(map[string]updateCache)
	featureDbxrefUpsertCacheMut       sync.RWMutex
	featureDbxrefUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureDbxrefBeforeInsertHooks []FeatureDbxrefHook
var featureDbxrefBeforeUpdateHooks []FeatureDbxrefHook
var featureDbxrefBeforeDeleteHooks []FeatureDbxrefHook
var featureDbxrefBeforeUpsertHooks []FeatureDbxrefHook

var featureDbxrefAfterInsertHooks []FeatureDbxrefHook
var featureDbxrefAfterSelectHooks []FeatureDbxrefHook
var featureDbxrefAfterUpdateHooks []FeatureDbxrefHook
var featureDbxrefAfterDeleteHooks []FeatureDbxrefHook
var featureDbxrefAfterUpsertHooks []FeatureDbxrefHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeatureDbxref) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureDbxrefBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeatureDbxref) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureDbxrefBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeatureDbxref) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureDbxrefBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeatureDbxref) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureDbxrefBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeatureDbxref) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureDbxrefAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeatureDbxref) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureDbxrefAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeatureDbxref) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureDbxrefAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeatureDbxref) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureDbxrefAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeatureDbxref) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureDbxrefAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureDbxrefHook registers your hook function for all future operations.
func AddFeatureDbxrefHook(hookPoint boil.HookPoint, featureDbxrefHook FeatureDbxrefHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureDbxrefBeforeInsertHooks = append(featureDbxrefBeforeInsertHooks, featureDbxrefHook)
	case boil.BeforeUpdateHook:
		featureDbxrefBeforeUpdateHooks = append(featureDbxrefBeforeUpdateHooks, featureDbxrefHook)
	case boil.BeforeDeleteHook:
		featureDbxrefBeforeDeleteHooks = append(featureDbxrefBeforeDeleteHooks, featureDbxrefHook)
	case boil.BeforeUpsertHook:
		featureDbxrefBeforeUpsertHooks = append(featureDbxrefBeforeUpsertHooks, featureDbxrefHook)
	case boil.AfterInsertHook:
		featureDbxrefAfterInsertHooks = append(featureDbxrefAfterInsertHooks, featureDbxrefHook)
	case boil.AfterSelectHook:
		featureDbxrefAfterSelectHooks = append(featureDbxrefAfterSelectHooks, featureDbxrefHook)
	case boil.AfterUpdateHook:
		featureDbxrefAfterUpdateHooks = append(featureDbxrefAfterUpdateHooks, featureDbxrefHook)
	case boil.AfterDeleteHook:
		featureDbxrefAfterDeleteHooks = append(featureDbxrefAfterDeleteHooks, featureDbxrefHook)
	case boil.AfterUpsertHook:
		featureDbxrefAfterUpsertHooks = append(featureDbxrefAfterUpsertHooks, featureDbxrefHook)
	}
}

// OneP returns a single featureDbxref record from the query, and panics on error.
func (q featureDbxrefQuery) OneP() *FeatureDbxref {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureDbxref record from the query.
func (q featureDbxrefQuery) One() (*FeatureDbxref, error) {
	o := &FeatureDbxref{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for feature_dbxref")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeatureDbxref records from the query, and panics on error.
func (q featureDbxrefQuery) AllP() FeatureDbxrefSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeatureDbxref records from the query.
func (q featureDbxrefQuery) All() (FeatureDbxrefSlice, error) {
	var o FeatureDbxrefSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to FeatureDbxref slice")
	}

	if len(featureDbxrefAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeatureDbxref records in the query, and panics on error.
func (q featureDbxrefQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeatureDbxref records in the query.
func (q featureDbxrefQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count feature_dbxref rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureDbxrefQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureDbxrefQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if feature_dbxref exists")
	}

	return count > 0, nil
}

// DbxrefG pointed to by the foreign key.
func (o *FeatureDbxref) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *FeatureDbxref) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// FeatureG pointed to by the foreign key.
func (o *FeatureDbxref) FeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Feature(boil.GetDB(), mods...)
}

// Feature pointed to by the foreign key.
func (o *FeatureDbxref) Feature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureDbxrefL) LoadDbxref(e boil.Executor, singular bool, maybeFeatureDbxref interface{}) error {
	var slice []*FeatureDbxref
	var object *FeatureDbxref

	count := 1
	if singular {
		object = maybeFeatureDbxref.(*FeatureDbxref)
	} else {
		slice = *maybeFeatureDbxref.(*FeatureDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureDbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &featureDbxrefR{}
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

	if len(featureDbxrefAfterSelectHooks) != 0 {
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

// LoadFeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureDbxrefL) LoadFeature(e boil.Executor, singular bool, maybeFeatureDbxref interface{}) error {
	var slice []*FeatureDbxref
	var object *FeatureDbxref

	count := 1
	if singular {
		object = maybeFeatureDbxref.(*FeatureDbxref)
	} else {
		slice = *maybeFeatureDbxref.(*FeatureDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureDbxrefR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureDbxrefR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Feature")
	}
	defer results.Close()

	var resultSlice []*Feature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Feature")
	}

	if len(featureDbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Feature = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.Feature = foreign
				break
			}
		}
	}

	return nil
}

// SetDbxref of the feature_dbxref to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.FeatureDbxref.
func (o *FeatureDbxref) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.DbxrefID, o.FeatureDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DbxrefID = related.DbxrefID

	if o.R == nil {
		o.R = &featureDbxrefR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			FeatureDbxref: o,
		}
	} else {
		related.R.FeatureDbxref = o
	}

	return nil
}

// SetFeature of the feature_dbxref to the related item.
// Sets o.R.Feature to related.
// Adds o to related.R.FeatureDbxref.
func (o *FeatureDbxref) SetFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeatureDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureID = related.FeatureID

	if o.R == nil {
		o.R = &featureDbxrefR{
			Feature: related,
		}
	} else {
		o.R.Feature = related
	}

	if related.R == nil {
		related.R = &featureR{
			FeatureDbxref: o,
		}
	} else {
		related.R.FeatureDbxref = o
	}

	return nil
}

// FeatureDbxrefsG retrieves all records.
func FeatureDbxrefsG(mods ...qm.QueryMod) featureDbxrefQuery {
	return FeatureDbxrefs(boil.GetDB(), mods...)
}

// FeatureDbxrefs retrieves all the records using an executor.
func FeatureDbxrefs(exec boil.Executor, mods ...qm.QueryMod) featureDbxrefQuery {
	mods = append(mods, qm.From("\"feature_dbxref\""))
	return featureDbxrefQuery{NewQuery(exec, mods...)}
}

// FindFeatureDbxrefG retrieves a single record by ID.
func FindFeatureDbxrefG(featureDbxrefID int, selectCols ...string) (*FeatureDbxref, error) {
	return FindFeatureDbxref(boil.GetDB(), featureDbxrefID, selectCols...)
}

// FindFeatureDbxrefGP retrieves a single record by ID, and panics on error.
func FindFeatureDbxrefGP(featureDbxrefID int, selectCols ...string) *FeatureDbxref {
	retobj, err := FindFeatureDbxref(boil.GetDB(), featureDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureDbxref retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureDbxref(exec boil.Executor, featureDbxrefID int, selectCols ...string) (*FeatureDbxref, error) {
	featureDbxrefObj := &FeatureDbxref{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_dbxref\" where \"feature_dbxref_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureDbxrefID)

	err := q.Bind(featureDbxrefObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from feature_dbxref")
	}

	return featureDbxrefObj, nil
}

// FindFeatureDbxrefP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureDbxrefP(exec boil.Executor, featureDbxrefID int, selectCols ...string) *FeatureDbxref {
	retobj, err := FindFeatureDbxref(exec, featureDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeatureDbxref) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeatureDbxref) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeatureDbxref) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeatureDbxref) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_dbxref provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureDbxrefColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureDbxrefInsertCacheMut.RLock()
	cache, cached := featureDbxrefInsertCache[key]
	featureDbxrefInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureDbxrefColumns,
			featureDbxrefColumnsWithDefault,
			featureDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureDbxrefType, featureDbxrefMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureDbxrefType, featureDbxrefMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_dbxref\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into feature_dbxref")
	}

	if !cached {
		featureDbxrefInsertCacheMut.Lock()
		featureDbxrefInsertCache[key] = cache
		featureDbxrefInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeatureDbxref record. See Update for
// whitelist behavior description.
func (o *FeatureDbxref) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeatureDbxref record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeatureDbxref) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeatureDbxref, and panics on error.
// See Update for whitelist behavior description.
func (o *FeatureDbxref) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeatureDbxref.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeatureDbxref) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureDbxrefUpdateCacheMut.RLock()
	cache, cached := featureDbxrefUpdateCache[key]
	featureDbxrefUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureDbxrefColumns, featureDbxrefPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update feature_dbxref, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featureDbxrefPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureDbxrefType, featureDbxrefMapping, append(wl, featureDbxrefPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update feature_dbxref row")
	}

	if !cached {
		featureDbxrefUpdateCacheMut.Lock()
		featureDbxrefUpdateCache[key] = cache
		featureDbxrefUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureDbxrefQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureDbxrefQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for feature_dbxref")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureDbxrefSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureDbxrefSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureDbxrefSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureDbxrefSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_dbxref\" SET %s WHERE (\"feature_dbxref_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureDbxrefPrimaryKeyColumns), len(colNames)+1, len(featureDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in featureDbxref slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeatureDbxref) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeatureDbxref) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeatureDbxref) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeatureDbxref) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_dbxref provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureDbxrefColumnsWithDefault, o)

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

	featureDbxrefUpsertCacheMut.RLock()
	cache, cached := featureDbxrefUpsertCache[key]
	featureDbxrefUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureDbxrefColumns,
			featureDbxrefColumnsWithDefault,
			featureDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureDbxrefColumns,
			featureDbxrefPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert feature_dbxref, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featureDbxrefPrimaryKeyColumns))
			copy(conflict, featureDbxrefPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_dbxref\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureDbxrefType, featureDbxrefMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureDbxrefType, featureDbxrefMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for feature_dbxref")
	}

	if !cached {
		featureDbxrefUpsertCacheMut.Lock()
		featureDbxrefUpsertCache[key] = cache
		featureDbxrefUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeatureDbxref record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureDbxref) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeatureDbxref record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeatureDbxref) DeleteG() error {
	if o == nil {
		return errors.New("chado: no FeatureDbxref provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeatureDbxref record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureDbxref) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeatureDbxref record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeatureDbxref) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeatureDbxref provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureDbxrefPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_dbxref\" WHERE \"feature_dbxref_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from feature_dbxref")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureDbxrefQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureDbxrefQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no featureDbxrefQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from feature_dbxref")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureDbxrefSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureDbxrefSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no FeatureDbxref slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureDbxrefSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureDbxrefSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeatureDbxref slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureDbxrefBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureDbxrefPrimaryKeyColumns), 1, len(featureDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featureDbxref slice")
	}

	if len(featureDbxrefAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeatureDbxref) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeatureDbxref) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeatureDbxref) ReloadG() error {
	if o == nil {
		return errors.New("chado: no FeatureDbxref provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeatureDbxref) Reload(exec boil.Executor) error {
	ret, err := FindFeatureDbxref(exec, o.FeatureDbxrefID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureDbxrefSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureDbxrefSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureDbxrefSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty FeatureDbxrefSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureDbxrefSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureDbxrefs := FeatureDbxrefSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_dbxref\".* FROM \"feature_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featureDbxrefPrimaryKeyColumns), 1, len(featureDbxrefPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureDbxrefs)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in FeatureDbxrefSlice")
	}

	*o = featureDbxrefs

	return nil
}

// FeatureDbxrefExists checks if the FeatureDbxref row exists.
func FeatureDbxrefExists(exec boil.Executor, featureDbxrefID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_dbxref\" where \"feature_dbxref_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureDbxrefID)
	}

	row := exec.QueryRow(sql, featureDbxrefID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if feature_dbxref exists")
	}

	return exists, nil
}

// FeatureDbxrefExistsG checks if the FeatureDbxref row exists.
func FeatureDbxrefExistsG(featureDbxrefID int) (bool, error) {
	return FeatureDbxrefExists(boil.GetDB(), featureDbxrefID)
}

// FeatureDbxrefExistsGP checks if the FeatureDbxref row exists. Panics on error.
func FeatureDbxrefExistsGP(featureDbxrefID int) bool {
	e, err := FeatureDbxrefExists(boil.GetDB(), featureDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureDbxrefExistsP checks if the FeatureDbxref row exists. Panics on error.
func FeatureDbxrefExistsP(exec boil.Executor, featureDbxrefID int) bool {
	e, err := FeatureDbxrefExists(exec, featureDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
