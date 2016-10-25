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

// FeaturelocPub is an object representing the database table.
type FeaturelocPub struct {
	FeaturelocPubID int `boil:"featureloc_pub_id" json:"featureloc_pub_id" toml:"featureloc_pub_id" yaml:"featureloc_pub_id"`
	FeaturelocID    int `boil:"featureloc_id" json:"featureloc_id" toml:"featureloc_id" yaml:"featureloc_id"`
	PubID           int `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`

	R *featurelocPubR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featurelocPubL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featurelocPubR is where relationships are stored.
type featurelocPubR struct {
	Pub        *Pub
	Featureloc *Featureloc
}

// featurelocPubL is where Load methods for each relationship are stored.
type featurelocPubL struct{}

var (
	featurelocPubColumns               = []string{"featureloc_pub_id", "featureloc_id", "pub_id"}
	featurelocPubColumnsWithoutDefault = []string{"featureloc_id", "pub_id"}
	featurelocPubColumnsWithDefault    = []string{"featureloc_pub_id"}
	featurelocPubPrimaryKeyColumns     = []string{"featureloc_pub_id"}
)

type (
	// FeaturelocPubSlice is an alias for a slice of pointers to FeaturelocPub.
	// This should generally be used opposed to []FeaturelocPub.
	FeaturelocPubSlice []*FeaturelocPub
	// FeaturelocPubHook is the signature for custom FeaturelocPub hook methods
	FeaturelocPubHook func(boil.Executor, *FeaturelocPub) error

	featurelocPubQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featurelocPubType                 = reflect.TypeOf(&FeaturelocPub{})
	featurelocPubMapping              = queries.MakeStructMapping(featurelocPubType)
	featurelocPubPrimaryKeyMapping, _ = queries.BindMapping(featurelocPubType, featurelocPubMapping, featurelocPubPrimaryKeyColumns)
	featurelocPubInsertCacheMut       sync.RWMutex
	featurelocPubInsertCache          = make(map[string]insertCache)
	featurelocPubUpdateCacheMut       sync.RWMutex
	featurelocPubUpdateCache          = make(map[string]updateCache)
	featurelocPubUpsertCacheMut       sync.RWMutex
	featurelocPubUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featurelocPubBeforeInsertHooks []FeaturelocPubHook
var featurelocPubBeforeUpdateHooks []FeaturelocPubHook
var featurelocPubBeforeDeleteHooks []FeaturelocPubHook
var featurelocPubBeforeUpsertHooks []FeaturelocPubHook

var featurelocPubAfterInsertHooks []FeaturelocPubHook
var featurelocPubAfterSelectHooks []FeaturelocPubHook
var featurelocPubAfterUpdateHooks []FeaturelocPubHook
var featurelocPubAfterDeleteHooks []FeaturelocPubHook
var featurelocPubAfterUpsertHooks []FeaturelocPubHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeaturelocPub) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocPubBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeaturelocPub) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocPubBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeaturelocPub) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocPubBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeaturelocPub) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocPubBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeaturelocPub) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocPubAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeaturelocPub) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocPubAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeaturelocPub) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocPubAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeaturelocPub) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocPubAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeaturelocPub) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocPubAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeaturelocPubHook registers your hook function for all future operations.
func AddFeaturelocPubHook(hookPoint boil.HookPoint, featurelocPubHook FeaturelocPubHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featurelocPubBeforeInsertHooks = append(featurelocPubBeforeInsertHooks, featurelocPubHook)
	case boil.BeforeUpdateHook:
		featurelocPubBeforeUpdateHooks = append(featurelocPubBeforeUpdateHooks, featurelocPubHook)
	case boil.BeforeDeleteHook:
		featurelocPubBeforeDeleteHooks = append(featurelocPubBeforeDeleteHooks, featurelocPubHook)
	case boil.BeforeUpsertHook:
		featurelocPubBeforeUpsertHooks = append(featurelocPubBeforeUpsertHooks, featurelocPubHook)
	case boil.AfterInsertHook:
		featurelocPubAfterInsertHooks = append(featurelocPubAfterInsertHooks, featurelocPubHook)
	case boil.AfterSelectHook:
		featurelocPubAfterSelectHooks = append(featurelocPubAfterSelectHooks, featurelocPubHook)
	case boil.AfterUpdateHook:
		featurelocPubAfterUpdateHooks = append(featurelocPubAfterUpdateHooks, featurelocPubHook)
	case boil.AfterDeleteHook:
		featurelocPubAfterDeleteHooks = append(featurelocPubAfterDeleteHooks, featurelocPubHook)
	case boil.AfterUpsertHook:
		featurelocPubAfterUpsertHooks = append(featurelocPubAfterUpsertHooks, featurelocPubHook)
	}
}

// OneP returns a single featurelocPub record from the query, and panics on error.
func (q featurelocPubQuery) OneP() *FeaturelocPub {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featurelocPub record from the query.
func (q featurelocPubQuery) One() (*FeaturelocPub, error) {
	o := &FeaturelocPub{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for featureloc_pub")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeaturelocPub records from the query, and panics on error.
func (q featurelocPubQuery) AllP() FeaturelocPubSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeaturelocPub records from the query.
func (q featurelocPubQuery) All() (FeaturelocPubSlice, error) {
	var o FeaturelocPubSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FeaturelocPub slice")
	}

	if len(featurelocPubAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeaturelocPub records in the query, and panics on error.
func (q featurelocPubQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeaturelocPub records in the query.
func (q featurelocPubQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count featureloc_pub rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featurelocPubQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featurelocPubQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if featureloc_pub exists")
	}

	return count > 0, nil
}

// PubG pointed to by the foreign key.
func (o *FeaturelocPub) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *FeaturelocPub) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// FeaturelocG pointed to by the foreign key.
func (o *FeaturelocPub) FeaturelocG(mods ...qm.QueryMod) featurelocQuery {
	return o.Featureloc(boil.GetDB(), mods...)
}

// Featureloc pointed to by the foreign key.
func (o *FeaturelocPub) Featureloc(exec boil.Executor, mods ...qm.QueryMod) featurelocQuery {
	queryMods := []qm.QueryMod{
		qm.Where("featureloc_id=$1", o.FeaturelocID),
	}

	queryMods = append(queryMods, mods...)

	query := Featurelocs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"featureloc\"")

	return query
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurelocPubL) LoadPub(e boil.Executor, singular bool, maybeFeaturelocPub interface{}) error {
	var slice []*FeaturelocPub
	var object *FeaturelocPub

	count := 1
	if singular {
		object = maybeFeaturelocPub.(*FeaturelocPub)
	} else {
		slice = *maybeFeaturelocPub.(*FeaturelocPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurelocPubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &featurelocPubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Pub")
	}
	defer results.Close()

	var resultSlice []*Pub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Pub")
	}

	if len(featurelocPubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Pub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.Pub = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureloc allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurelocPubL) LoadFeatureloc(e boil.Executor, singular bool, maybeFeaturelocPub interface{}) error {
	var slice []*FeaturelocPub
	var object *FeaturelocPub

	count := 1
	if singular {
		object = maybeFeaturelocPub.(*FeaturelocPub)
	} else {
		slice = *maybeFeaturelocPub.(*FeaturelocPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurelocPubR{}
		args[0] = object.FeaturelocID
	} else {
		for i, obj := range slice {
			obj.R = &featurelocPubR{}
			args[i] = obj.FeaturelocID
		}
	}

	query := fmt.Sprintf(
		"select * from \"featureloc\" where \"featureloc_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Featureloc")
	}
	defer results.Close()

	var resultSlice []*Featureloc
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Featureloc")
	}

	if len(featurelocPubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Featureloc = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeaturelocID == foreign.FeaturelocID {
				local.R.Featureloc = foreign
				break
			}
		}
	}

	return nil
}

// SetPub of the featureloc_pub to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.FeaturelocPub.
func (o *FeaturelocPub) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"featureloc_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurelocPubPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.FeaturelocPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &featurelocPubR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			FeaturelocPub: o,
		}
	} else {
		related.R.FeaturelocPub = o
	}

	return nil
}

// SetFeatureloc of the featureloc_pub to the related item.
// Sets o.R.Featureloc to related.
// Adds o to related.R.FeaturelocPub.
func (o *FeaturelocPub) SetFeatureloc(exec boil.Executor, insert bool, related *Featureloc) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"featureloc_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"featureloc_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurelocPubPrimaryKeyColumns),
	)
	values := []interface{}{related.FeaturelocID, o.FeaturelocPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeaturelocID = related.FeaturelocID

	if o.R == nil {
		o.R = &featurelocPubR{
			Featureloc: related,
		}
	} else {
		o.R.Featureloc = related
	}

	if related.R == nil {
		related.R = &featurelocR{
			FeaturelocPub: o,
		}
	} else {
		related.R.FeaturelocPub = o
	}

	return nil
}

// FeaturelocPubsG retrieves all records.
func FeaturelocPubsG(mods ...qm.QueryMod) featurelocPubQuery {
	return FeaturelocPubs(boil.GetDB(), mods...)
}

// FeaturelocPubs retrieves all the records using an executor.
func FeaturelocPubs(exec boil.Executor, mods ...qm.QueryMod) featurelocPubQuery {
	mods = append(mods, qm.From("\"featureloc_pub\""))
	return featurelocPubQuery{NewQuery(exec, mods...)}
}

// FindFeaturelocPubG retrieves a single record by ID.
func FindFeaturelocPubG(featurelocPubID int, selectCols ...string) (*FeaturelocPub, error) {
	return FindFeaturelocPub(boil.GetDB(), featurelocPubID, selectCols...)
}

// FindFeaturelocPubGP retrieves a single record by ID, and panics on error.
func FindFeaturelocPubGP(featurelocPubID int, selectCols ...string) *FeaturelocPub {
	retobj, err := FindFeaturelocPub(boil.GetDB(), featurelocPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeaturelocPub retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeaturelocPub(exec boil.Executor, featurelocPubID int, selectCols ...string) (*FeaturelocPub, error) {
	featurelocPubObj := &FeaturelocPub{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"featureloc_pub\" where \"featureloc_pub_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featurelocPubID)

	err := q.Bind(featurelocPubObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from featureloc_pub")
	}

	return featurelocPubObj, nil
}

// FindFeaturelocPubP retrieves a single record by ID with an executor, and panics on error.
func FindFeaturelocPubP(exec boil.Executor, featurelocPubID int, selectCols ...string) *FeaturelocPub {
	retobj, err := FindFeaturelocPub(exec, featurelocPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeaturelocPub) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeaturelocPub) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeaturelocPub) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeaturelocPub) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no featureloc_pub provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurelocPubColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featurelocPubInsertCacheMut.RLock()
	cache, cached := featurelocPubInsertCache[key]
	featurelocPubInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featurelocPubColumns,
			featurelocPubColumnsWithDefault,
			featurelocPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featurelocPubType, featurelocPubMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featurelocPubType, featurelocPubMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"featureloc_pub\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into featureloc_pub")
	}

	if !cached {
		featurelocPubInsertCacheMut.Lock()
		featurelocPubInsertCache[key] = cache
		featurelocPubInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeaturelocPub record. See Update for
// whitelist behavior description.
func (o *FeaturelocPub) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeaturelocPub record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeaturelocPub) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeaturelocPub, and panics on error.
// See Update for whitelist behavior description.
func (o *FeaturelocPub) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeaturelocPub.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeaturelocPub) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featurelocPubUpdateCacheMut.RLock()
	cache, cached := featurelocPubUpdateCache[key]
	featurelocPubUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featurelocPubColumns, featurelocPubPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update featureloc_pub, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"featureloc_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featurelocPubPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featurelocPubType, featurelocPubMapping, append(wl, featurelocPubPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update featureloc_pub row")
	}

	if !cached {
		featurelocPubUpdateCacheMut.Lock()
		featurelocPubUpdateCache[key] = cache
		featurelocPubUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featurelocPubQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featurelocPubQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for featureloc_pub")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeaturelocPubSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeaturelocPubSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeaturelocPubSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeaturelocPubSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurelocPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"featureloc_pub\" SET %s WHERE (\"featureloc_pub_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurelocPubPrimaryKeyColumns), len(colNames)+1, len(featurelocPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in featurelocPub slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeaturelocPub) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeaturelocPub) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeaturelocPub) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeaturelocPub) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no featureloc_pub provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurelocPubColumnsWithDefault, o)

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

	featurelocPubUpsertCacheMut.RLock()
	cache, cached := featurelocPubUpsertCache[key]
	featurelocPubUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featurelocPubColumns,
			featurelocPubColumnsWithDefault,
			featurelocPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featurelocPubColumns,
			featurelocPubPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert featureloc_pub, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featurelocPubPrimaryKeyColumns))
			copy(conflict, featurelocPubPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"featureloc_pub\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featurelocPubType, featurelocPubMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featurelocPubType, featurelocPubMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for featureloc_pub")
	}

	if !cached {
		featurelocPubUpsertCacheMut.Lock()
		featurelocPubUpsertCache[key] = cache
		featurelocPubUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeaturelocPub record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeaturelocPub) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeaturelocPub record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeaturelocPub) DeleteG() error {
	if o == nil {
		return errors.New("models: no FeaturelocPub provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeaturelocPub record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeaturelocPub) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeaturelocPub record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeaturelocPub) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeaturelocPub provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featurelocPubPrimaryKeyMapping)
	sql := "DELETE FROM \"featureloc_pub\" WHERE \"featureloc_pub_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from featureloc_pub")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featurelocPubQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featurelocPubQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no featurelocPubQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from featureloc_pub")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeaturelocPubSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeaturelocPubSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no FeaturelocPub slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeaturelocPubSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeaturelocPubSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeaturelocPub slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featurelocPubBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurelocPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"featureloc_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurelocPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurelocPubPrimaryKeyColumns), 1, len(featurelocPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from featurelocPub slice")
	}

	if len(featurelocPubAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeaturelocPub) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeaturelocPub) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeaturelocPub) ReloadG() error {
	if o == nil {
		return errors.New("models: no FeaturelocPub provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeaturelocPub) Reload(exec boil.Executor) error {
	ret, err := FindFeaturelocPub(exec, o.FeaturelocPubID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturelocPubSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturelocPubSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturelocPubSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FeaturelocPubSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturelocPubSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featurelocPubs := FeaturelocPubSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurelocPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"featureloc_pub\".* FROM \"featureloc_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurelocPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featurelocPubPrimaryKeyColumns), 1, len(featurelocPubPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featurelocPubs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeaturelocPubSlice")
	}

	*o = featurelocPubs

	return nil
}

// FeaturelocPubExists checks if the FeaturelocPub row exists.
func FeaturelocPubExists(exec boil.Executor, featurelocPubID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"featureloc_pub\" where \"featureloc_pub_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featurelocPubID)
	}

	row := exec.QueryRow(sql, featurelocPubID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if featureloc_pub exists")
	}

	return exists, nil
}

// FeaturelocPubExistsG checks if the FeaturelocPub row exists.
func FeaturelocPubExistsG(featurelocPubID int) (bool, error) {
	return FeaturelocPubExists(boil.GetDB(), featurelocPubID)
}

// FeaturelocPubExistsGP checks if the FeaturelocPub row exists. Panics on error.
func FeaturelocPubExistsGP(featurelocPubID int) bool {
	e, err := FeaturelocPubExists(boil.GetDB(), featurelocPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeaturelocPubExistsP checks if the FeaturelocPub row exists. Panics on error.
func FeaturelocPubExistsP(exec boil.Executor, featurelocPubID int) bool {
	e, err := FeaturelocPubExists(exec, featurelocPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
