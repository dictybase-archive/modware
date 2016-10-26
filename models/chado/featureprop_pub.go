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

// FeaturepropPub is an object representing the database table.
type FeaturepropPub struct {
	FeaturepropPubID int `boil:"featureprop_pub_id" json:"featureprop_pub_id" toml:"featureprop_pub_id" yaml:"featureprop_pub_id"`
	FeaturepropID    int `boil:"featureprop_id" json:"featureprop_id" toml:"featureprop_id" yaml:"featureprop_id"`
	PubID            int `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`

	R *featurepropPubR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featurepropPubL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featurepropPubR is where relationships are stored.
type featurepropPubR struct {
	Featureprop *Featureprop
	Pub         *Pub
}

// featurepropPubL is where Load methods for each relationship are stored.
type featurepropPubL struct{}

var (
	featurepropPubColumns               = []string{"featureprop_pub_id", "featureprop_id", "pub_id"}
	featurepropPubColumnsWithoutDefault = []string{"featureprop_id", "pub_id"}
	featurepropPubColumnsWithDefault    = []string{"featureprop_pub_id"}
	featurepropPubPrimaryKeyColumns     = []string{"featureprop_pub_id"}
)

type (
	// FeaturepropPubSlice is an alias for a slice of pointers to FeaturepropPub.
	// This should generally be used opposed to []FeaturepropPub.
	FeaturepropPubSlice []*FeaturepropPub
	// FeaturepropPubHook is the signature for custom FeaturepropPub hook methods
	FeaturepropPubHook func(boil.Executor, *FeaturepropPub) error

	featurepropPubQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featurepropPubType                 = reflect.TypeOf(&FeaturepropPub{})
	featurepropPubMapping              = queries.MakeStructMapping(featurepropPubType)
	featurepropPubPrimaryKeyMapping, _ = queries.BindMapping(featurepropPubType, featurepropPubMapping, featurepropPubPrimaryKeyColumns)
	featurepropPubInsertCacheMut       sync.RWMutex
	featurepropPubInsertCache          = make(map[string]insertCache)
	featurepropPubUpdateCacheMut       sync.RWMutex
	featurepropPubUpdateCache          = make(map[string]updateCache)
	featurepropPubUpsertCacheMut       sync.RWMutex
	featurepropPubUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featurepropPubBeforeInsertHooks []FeaturepropPubHook
var featurepropPubBeforeUpdateHooks []FeaturepropPubHook
var featurepropPubBeforeDeleteHooks []FeaturepropPubHook
var featurepropPubBeforeUpsertHooks []FeaturepropPubHook

var featurepropPubAfterInsertHooks []FeaturepropPubHook
var featurepropPubAfterSelectHooks []FeaturepropPubHook
var featurepropPubAfterUpdateHooks []FeaturepropPubHook
var featurepropPubAfterDeleteHooks []FeaturepropPubHook
var featurepropPubAfterUpsertHooks []FeaturepropPubHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeaturepropPub) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropPubBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeaturepropPub) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropPubBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeaturepropPub) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropPubBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeaturepropPub) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropPubBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeaturepropPub) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropPubAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeaturepropPub) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropPubAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeaturepropPub) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropPubAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeaturepropPub) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropPubAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeaturepropPub) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropPubAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeaturepropPubHook registers your hook function for all future operations.
func AddFeaturepropPubHook(hookPoint boil.HookPoint, featurepropPubHook FeaturepropPubHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featurepropPubBeforeInsertHooks = append(featurepropPubBeforeInsertHooks, featurepropPubHook)
	case boil.BeforeUpdateHook:
		featurepropPubBeforeUpdateHooks = append(featurepropPubBeforeUpdateHooks, featurepropPubHook)
	case boil.BeforeDeleteHook:
		featurepropPubBeforeDeleteHooks = append(featurepropPubBeforeDeleteHooks, featurepropPubHook)
	case boil.BeforeUpsertHook:
		featurepropPubBeforeUpsertHooks = append(featurepropPubBeforeUpsertHooks, featurepropPubHook)
	case boil.AfterInsertHook:
		featurepropPubAfterInsertHooks = append(featurepropPubAfterInsertHooks, featurepropPubHook)
	case boil.AfterSelectHook:
		featurepropPubAfterSelectHooks = append(featurepropPubAfterSelectHooks, featurepropPubHook)
	case boil.AfterUpdateHook:
		featurepropPubAfterUpdateHooks = append(featurepropPubAfterUpdateHooks, featurepropPubHook)
	case boil.AfterDeleteHook:
		featurepropPubAfterDeleteHooks = append(featurepropPubAfterDeleteHooks, featurepropPubHook)
	case boil.AfterUpsertHook:
		featurepropPubAfterUpsertHooks = append(featurepropPubAfterUpsertHooks, featurepropPubHook)
	}
}

// OneP returns a single featurepropPub record from the query, and panics on error.
func (q featurepropPubQuery) OneP() *FeaturepropPub {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featurepropPub record from the query.
func (q featurepropPubQuery) One() (*FeaturepropPub, error) {
	o := &FeaturepropPub{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for featureprop_pub")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeaturepropPub records from the query, and panics on error.
func (q featurepropPubQuery) AllP() FeaturepropPubSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeaturepropPub records from the query.
func (q featurepropPubQuery) All() (FeaturepropPubSlice, error) {
	var o FeaturepropPubSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to FeaturepropPub slice")
	}

	if len(featurepropPubAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeaturepropPub records in the query, and panics on error.
func (q featurepropPubQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeaturepropPub records in the query.
func (q featurepropPubQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count featureprop_pub rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featurepropPubQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featurepropPubQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if featureprop_pub exists")
	}

	return count > 0, nil
}

// FeaturepropG pointed to by the foreign key.
func (o *FeaturepropPub) FeaturepropG(mods ...qm.QueryMod) featurepropQuery {
	return o.Featureprop(boil.GetDB(), mods...)
}

// Featureprop pointed to by the foreign key.
func (o *FeaturepropPub) Featureprop(exec boil.Executor, mods ...qm.QueryMod) featurepropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("featureprop_id=$1", o.FeaturepropID),
	}

	queryMods = append(queryMods, mods...)

	query := Featureprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"featureprop\"")

	return query
}

// PubG pointed to by the foreign key.
func (o *FeaturepropPub) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *FeaturepropPub) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// LoadFeatureprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurepropPubL) LoadFeatureprop(e boil.Executor, singular bool, maybeFeaturepropPub interface{}) error {
	var slice []*FeaturepropPub
	var object *FeaturepropPub

	count := 1
	if singular {
		object = maybeFeaturepropPub.(*FeaturepropPub)
	} else {
		slice = *maybeFeaturepropPub.(*FeaturepropPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurepropPubR{}
		args[0] = object.FeaturepropID
	} else {
		for i, obj := range slice {
			obj.R = &featurepropPubR{}
			args[i] = obj.FeaturepropID
		}
	}

	query := fmt.Sprintf(
		"select * from \"featureprop\" where \"featureprop_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Featureprop")
	}
	defer results.Close()

	var resultSlice []*Featureprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Featureprop")
	}

	if len(featurepropPubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Featureprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeaturepropID == foreign.FeaturepropID {
				local.R.Featureprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurepropPubL) LoadPub(e boil.Executor, singular bool, maybeFeaturepropPub interface{}) error {
	var slice []*FeaturepropPub
	var object *FeaturepropPub

	count := 1
	if singular {
		object = maybeFeaturepropPub.(*FeaturepropPub)
	} else {
		slice = *maybeFeaturepropPub.(*FeaturepropPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurepropPubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &featurepropPubR{}
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

	if len(featurepropPubAfterSelectHooks) != 0 {
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

// SetFeatureprop of the featureprop_pub to the related item.
// Sets o.R.Featureprop to related.
// Adds o to related.R.FeaturepropPub.
func (o *FeaturepropPub) SetFeatureprop(exec boil.Executor, insert bool, related *Featureprop) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"featureprop_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"featureprop_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurepropPubPrimaryKeyColumns),
	)
	values := []interface{}{related.FeaturepropID, o.FeaturepropPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeaturepropID = related.FeaturepropID

	if o.R == nil {
		o.R = &featurepropPubR{
			Featureprop: related,
		}
	} else {
		o.R.Featureprop = related
	}

	if related.R == nil {
		related.R = &featurepropR{
			FeaturepropPub: o,
		}
	} else {
		related.R.FeaturepropPub = o
	}

	return nil
}

// SetPub of the featureprop_pub to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.FeaturepropPub.
func (o *FeaturepropPub) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"featureprop_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurepropPubPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.FeaturepropPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &featurepropPubR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			FeaturepropPub: o,
		}
	} else {
		related.R.FeaturepropPub = o
	}

	return nil
}

// FeaturepropPubsG retrieves all records.
func FeaturepropPubsG(mods ...qm.QueryMod) featurepropPubQuery {
	return FeaturepropPubs(boil.GetDB(), mods...)
}

// FeaturepropPubs retrieves all the records using an executor.
func FeaturepropPubs(exec boil.Executor, mods ...qm.QueryMod) featurepropPubQuery {
	mods = append(mods, qm.From("\"featureprop_pub\""))
	return featurepropPubQuery{NewQuery(exec, mods...)}
}

// FindFeaturepropPubG retrieves a single record by ID.
func FindFeaturepropPubG(featurepropPubID int, selectCols ...string) (*FeaturepropPub, error) {
	return FindFeaturepropPub(boil.GetDB(), featurepropPubID, selectCols...)
}

// FindFeaturepropPubGP retrieves a single record by ID, and panics on error.
func FindFeaturepropPubGP(featurepropPubID int, selectCols ...string) *FeaturepropPub {
	retobj, err := FindFeaturepropPub(boil.GetDB(), featurepropPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeaturepropPub retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeaturepropPub(exec boil.Executor, featurepropPubID int, selectCols ...string) (*FeaturepropPub, error) {
	featurepropPubObj := &FeaturepropPub{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"featureprop_pub\" where \"featureprop_pub_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featurepropPubID)

	err := q.Bind(featurepropPubObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from featureprop_pub")
	}

	return featurepropPubObj, nil
}

// FindFeaturepropPubP retrieves a single record by ID with an executor, and panics on error.
func FindFeaturepropPubP(exec boil.Executor, featurepropPubID int, selectCols ...string) *FeaturepropPub {
	retobj, err := FindFeaturepropPub(exec, featurepropPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeaturepropPub) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeaturepropPub) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeaturepropPub) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeaturepropPub) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no featureprop_pub provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurepropPubColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featurepropPubInsertCacheMut.RLock()
	cache, cached := featurepropPubInsertCache[key]
	featurepropPubInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featurepropPubColumns,
			featurepropPubColumnsWithDefault,
			featurepropPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featurepropPubType, featurepropPubMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featurepropPubType, featurepropPubMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"featureprop_pub\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into featureprop_pub")
	}

	if !cached {
		featurepropPubInsertCacheMut.Lock()
		featurepropPubInsertCache[key] = cache
		featurepropPubInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeaturepropPub record. See Update for
// whitelist behavior description.
func (o *FeaturepropPub) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeaturepropPub record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeaturepropPub) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeaturepropPub, and panics on error.
// See Update for whitelist behavior description.
func (o *FeaturepropPub) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeaturepropPub.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeaturepropPub) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featurepropPubUpdateCacheMut.RLock()
	cache, cached := featurepropPubUpdateCache[key]
	featurepropPubUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featurepropPubColumns, featurepropPubPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update featureprop_pub, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"featureprop_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featurepropPubPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featurepropPubType, featurepropPubMapping, append(wl, featurepropPubPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update featureprop_pub row")
	}

	if !cached {
		featurepropPubUpdateCacheMut.Lock()
		featurepropPubUpdateCache[key] = cache
		featurepropPubUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featurepropPubQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featurepropPubQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for featureprop_pub")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeaturepropPubSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeaturepropPubSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeaturepropPubSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeaturepropPubSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurepropPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"featureprop_pub\" SET %s WHERE (\"featureprop_pub_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurepropPubPrimaryKeyColumns), len(colNames)+1, len(featurepropPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in featurepropPub slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeaturepropPub) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeaturepropPub) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeaturepropPub) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeaturepropPub) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no featureprop_pub provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurepropPubColumnsWithDefault, o)

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

	featurepropPubUpsertCacheMut.RLock()
	cache, cached := featurepropPubUpsertCache[key]
	featurepropPubUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featurepropPubColumns,
			featurepropPubColumnsWithDefault,
			featurepropPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featurepropPubColumns,
			featurepropPubPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert featureprop_pub, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featurepropPubPrimaryKeyColumns))
			copy(conflict, featurepropPubPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"featureprop_pub\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featurepropPubType, featurepropPubMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featurepropPubType, featurepropPubMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for featureprop_pub")
	}

	if !cached {
		featurepropPubUpsertCacheMut.Lock()
		featurepropPubUpsertCache[key] = cache
		featurepropPubUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeaturepropPub record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeaturepropPub) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeaturepropPub record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeaturepropPub) DeleteG() error {
	if o == nil {
		return errors.New("chado: no FeaturepropPub provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeaturepropPub record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeaturepropPub) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeaturepropPub record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeaturepropPub) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeaturepropPub provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featurepropPubPrimaryKeyMapping)
	sql := "DELETE FROM \"featureprop_pub\" WHERE \"featureprop_pub_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from featureprop_pub")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featurepropPubQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featurepropPubQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no featurepropPubQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featureprop_pub")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeaturepropPubSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeaturepropPubSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no FeaturepropPub slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeaturepropPubSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeaturepropPubSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeaturepropPub slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featurepropPubBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurepropPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"featureprop_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurepropPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurepropPubPrimaryKeyColumns), 1, len(featurepropPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featurepropPub slice")
	}

	if len(featurepropPubAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeaturepropPub) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeaturepropPub) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeaturepropPub) ReloadG() error {
	if o == nil {
		return errors.New("chado: no FeaturepropPub provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeaturepropPub) Reload(exec boil.Executor) error {
	ret, err := FindFeaturepropPub(exec, o.FeaturepropPubID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturepropPubSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturepropPubSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturepropPubSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty FeaturepropPubSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturepropPubSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featurepropPubs := FeaturepropPubSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurepropPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"featureprop_pub\".* FROM \"featureprop_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurepropPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featurepropPubPrimaryKeyColumns), 1, len(featurepropPubPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featurepropPubs)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in FeaturepropPubSlice")
	}

	*o = featurepropPubs

	return nil
}

// FeaturepropPubExists checks if the FeaturepropPub row exists.
func FeaturepropPubExists(exec boil.Executor, featurepropPubID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"featureprop_pub\" where \"featureprop_pub_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featurepropPubID)
	}

	row := exec.QueryRow(sql, featurepropPubID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if featureprop_pub exists")
	}

	return exists, nil
}

// FeaturepropPubExistsG checks if the FeaturepropPub row exists.
func FeaturepropPubExistsG(featurepropPubID int) (bool, error) {
	return FeaturepropPubExists(boil.GetDB(), featurepropPubID)
}

// FeaturepropPubExistsGP checks if the FeaturepropPub row exists. Panics on error.
func FeaturepropPubExistsGP(featurepropPubID int) bool {
	e, err := FeaturepropPubExists(boil.GetDB(), featurepropPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeaturepropPubExistsP checks if the FeaturepropPub row exists. Panics on error.
func FeaturepropPubExistsP(exec boil.Executor, featurepropPubID int) bool {
	e, err := FeaturepropPubExists(exec, featurepropPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
