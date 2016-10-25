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

// FeatureCvtermPub is an object representing the database table.
type FeatureCvtermPub struct {
	FeatureCvtermPubID int `boil:"feature_cvterm_pub_id" json:"feature_cvterm_pub_id" toml:"feature_cvterm_pub_id" yaml:"feature_cvterm_pub_id"`
	FeatureCvtermID    int `boil:"feature_cvterm_id" json:"feature_cvterm_id" toml:"feature_cvterm_id" yaml:"feature_cvterm_id"`
	PubID              int `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`

	R *featureCvtermPubR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureCvtermPubL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureCvtermPubR is where relationships are stored.
type featureCvtermPubR struct {
	FeatureCvterm *FeatureCvterm
	Pub           *Pub
}

// featureCvtermPubL is where Load methods for each relationship are stored.
type featureCvtermPubL struct{}

var (
	featureCvtermPubColumns               = []string{"feature_cvterm_pub_id", "feature_cvterm_id", "pub_id"}
	featureCvtermPubColumnsWithoutDefault = []string{"feature_cvterm_id", "pub_id"}
	featureCvtermPubColumnsWithDefault    = []string{"feature_cvterm_pub_id"}
	featureCvtermPubPrimaryKeyColumns     = []string{"feature_cvterm_pub_id"}
)

type (
	// FeatureCvtermPubSlice is an alias for a slice of pointers to FeatureCvtermPub.
	// This should generally be used opposed to []FeatureCvtermPub.
	FeatureCvtermPubSlice []*FeatureCvtermPub
	// FeatureCvtermPubHook is the signature for custom FeatureCvtermPub hook methods
	FeatureCvtermPubHook func(boil.Executor, *FeatureCvtermPub) error

	featureCvtermPubQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureCvtermPubType                 = reflect.TypeOf(&FeatureCvtermPub{})
	featureCvtermPubMapping              = queries.MakeStructMapping(featureCvtermPubType)
	featureCvtermPubPrimaryKeyMapping, _ = queries.BindMapping(featureCvtermPubType, featureCvtermPubMapping, featureCvtermPubPrimaryKeyColumns)
	featureCvtermPubInsertCacheMut       sync.RWMutex
	featureCvtermPubInsertCache          = make(map[string]insertCache)
	featureCvtermPubUpdateCacheMut       sync.RWMutex
	featureCvtermPubUpdateCache          = make(map[string]updateCache)
	featureCvtermPubUpsertCacheMut       sync.RWMutex
	featureCvtermPubUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureCvtermPubBeforeInsertHooks []FeatureCvtermPubHook
var featureCvtermPubBeforeUpdateHooks []FeatureCvtermPubHook
var featureCvtermPubBeforeDeleteHooks []FeatureCvtermPubHook
var featureCvtermPubBeforeUpsertHooks []FeatureCvtermPubHook

var featureCvtermPubAfterInsertHooks []FeatureCvtermPubHook
var featureCvtermPubAfterSelectHooks []FeatureCvtermPubHook
var featureCvtermPubAfterUpdateHooks []FeatureCvtermPubHook
var featureCvtermPubAfterDeleteHooks []FeatureCvtermPubHook
var featureCvtermPubAfterUpsertHooks []FeatureCvtermPubHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeatureCvtermPub) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermPubBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeatureCvtermPub) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermPubBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeatureCvtermPub) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermPubBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeatureCvtermPub) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermPubBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeatureCvtermPub) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermPubAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeatureCvtermPub) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermPubAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeatureCvtermPub) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermPubAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeatureCvtermPub) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermPubAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeatureCvtermPub) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermPubAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureCvtermPubHook registers your hook function for all future operations.
func AddFeatureCvtermPubHook(hookPoint boil.HookPoint, featureCvtermPubHook FeatureCvtermPubHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureCvtermPubBeforeInsertHooks = append(featureCvtermPubBeforeInsertHooks, featureCvtermPubHook)
	case boil.BeforeUpdateHook:
		featureCvtermPubBeforeUpdateHooks = append(featureCvtermPubBeforeUpdateHooks, featureCvtermPubHook)
	case boil.BeforeDeleteHook:
		featureCvtermPubBeforeDeleteHooks = append(featureCvtermPubBeforeDeleteHooks, featureCvtermPubHook)
	case boil.BeforeUpsertHook:
		featureCvtermPubBeforeUpsertHooks = append(featureCvtermPubBeforeUpsertHooks, featureCvtermPubHook)
	case boil.AfterInsertHook:
		featureCvtermPubAfterInsertHooks = append(featureCvtermPubAfterInsertHooks, featureCvtermPubHook)
	case boil.AfterSelectHook:
		featureCvtermPubAfterSelectHooks = append(featureCvtermPubAfterSelectHooks, featureCvtermPubHook)
	case boil.AfterUpdateHook:
		featureCvtermPubAfterUpdateHooks = append(featureCvtermPubAfterUpdateHooks, featureCvtermPubHook)
	case boil.AfterDeleteHook:
		featureCvtermPubAfterDeleteHooks = append(featureCvtermPubAfterDeleteHooks, featureCvtermPubHook)
	case boil.AfterUpsertHook:
		featureCvtermPubAfterUpsertHooks = append(featureCvtermPubAfterUpsertHooks, featureCvtermPubHook)
	}
}

// OneP returns a single featureCvtermPub record from the query, and panics on error.
func (q featureCvtermPubQuery) OneP() *FeatureCvtermPub {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureCvtermPub record from the query.
func (q featureCvtermPubQuery) One() (*FeatureCvtermPub, error) {
	o := &FeatureCvtermPub{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for feature_cvterm_pub")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeatureCvtermPub records from the query, and panics on error.
func (q featureCvtermPubQuery) AllP() FeatureCvtermPubSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeatureCvtermPub records from the query.
func (q featureCvtermPubQuery) All() (FeatureCvtermPubSlice, error) {
	var o FeatureCvtermPubSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FeatureCvtermPub slice")
	}

	if len(featureCvtermPubAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeatureCvtermPub records in the query, and panics on error.
func (q featureCvtermPubQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeatureCvtermPub records in the query.
func (q featureCvtermPubQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count feature_cvterm_pub rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureCvtermPubQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureCvtermPubQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if feature_cvterm_pub exists")
	}

	return count > 0, nil
}

// FeatureCvtermG pointed to by the foreign key.
func (o *FeatureCvtermPub) FeatureCvtermG(mods ...qm.QueryMod) featureCvtermQuery {
	return o.FeatureCvterm(boil.GetDB(), mods...)
}

// FeatureCvterm pointed to by the foreign key.
func (o *FeatureCvtermPub) FeatureCvterm(exec boil.Executor, mods ...qm.QueryMod) featureCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_cvterm_id=$1", o.FeatureCvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvterm\"")

	return query
}

// PubG pointed to by the foreign key.
func (o *FeatureCvtermPub) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *FeatureCvtermPub) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// LoadFeatureCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermPubL) LoadFeatureCvterm(e boil.Executor, singular bool, maybeFeatureCvtermPub interface{}) error {
	var slice []*FeatureCvtermPub
	var object *FeatureCvtermPub

	count := 1
	if singular {
		object = maybeFeatureCvtermPub.(*FeatureCvtermPub)
	} else {
		slice = *maybeFeatureCvtermPub.(*FeatureCvtermPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermPubR{}
		args[0] = object.FeatureCvtermID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermPubR{}
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

	if len(featureCvtermPubAfterSelectHooks) != 0 {
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

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermPubL) LoadPub(e boil.Executor, singular bool, maybeFeatureCvtermPub interface{}) error {
	var slice []*FeatureCvtermPub
	var object *FeatureCvtermPub

	count := 1
	if singular {
		object = maybeFeatureCvtermPub.(*FeatureCvtermPub)
	} else {
		slice = *maybeFeatureCvtermPub.(*FeatureCvtermPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermPubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermPubR{}
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

	if len(featureCvtermPubAfterSelectHooks) != 0 {
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

// SetFeatureCvterm of the feature_cvterm_pub to the related item.
// Sets o.R.FeatureCvterm to related.
// Adds o to related.R.FeatureCvtermPub.
func (o *FeatureCvtermPub) SetFeatureCvterm(exec boil.Executor, insert bool, related *FeatureCvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_cvterm_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureCvtermPubPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureCvtermID, o.FeatureCvtermPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureCvtermID = related.FeatureCvtermID

	if o.R == nil {
		o.R = &featureCvtermPubR{
			FeatureCvterm: related,
		}
	} else {
		o.R.FeatureCvterm = related
	}

	if related.R == nil {
		related.R = &featureCvtermR{
			FeatureCvtermPub: o,
		}
	} else {
		related.R.FeatureCvtermPub = o
	}

	return nil
}

// SetPub of the feature_cvterm_pub to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.FeatureCvtermPub.
func (o *FeatureCvtermPub) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_cvterm_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureCvtermPubPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.FeatureCvtermPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &featureCvtermPubR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			FeatureCvtermPub: o,
		}
	} else {
		related.R.FeatureCvtermPub = o
	}

	return nil
}

// FeatureCvtermPubsG retrieves all records.
func FeatureCvtermPubsG(mods ...qm.QueryMod) featureCvtermPubQuery {
	return FeatureCvtermPubs(boil.GetDB(), mods...)
}

// FeatureCvtermPubs retrieves all the records using an executor.
func FeatureCvtermPubs(exec boil.Executor, mods ...qm.QueryMod) featureCvtermPubQuery {
	mods = append(mods, qm.From("\"feature_cvterm_pub\""))
	return featureCvtermPubQuery{NewQuery(exec, mods...)}
}

// FindFeatureCvtermPubG retrieves a single record by ID.
func FindFeatureCvtermPubG(featureCvtermPubID int, selectCols ...string) (*FeatureCvtermPub, error) {
	return FindFeatureCvtermPub(boil.GetDB(), featureCvtermPubID, selectCols...)
}

// FindFeatureCvtermPubGP retrieves a single record by ID, and panics on error.
func FindFeatureCvtermPubGP(featureCvtermPubID int, selectCols ...string) *FeatureCvtermPub {
	retobj, err := FindFeatureCvtermPub(boil.GetDB(), featureCvtermPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureCvtermPub retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureCvtermPub(exec boil.Executor, featureCvtermPubID int, selectCols ...string) (*FeatureCvtermPub, error) {
	featureCvtermPubObj := &FeatureCvtermPub{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_cvterm_pub\" where \"feature_cvterm_pub_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureCvtermPubID)

	err := q.Bind(featureCvtermPubObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from feature_cvterm_pub")
	}

	return featureCvtermPubObj, nil
}

// FindFeatureCvtermPubP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureCvtermPubP(exec boil.Executor, featureCvtermPubID int, selectCols ...string) *FeatureCvtermPub {
	retobj, err := FindFeatureCvtermPub(exec, featureCvtermPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeatureCvtermPub) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeatureCvtermPub) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeatureCvtermPub) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeatureCvtermPub) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_cvterm_pub provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureCvtermPubColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureCvtermPubInsertCacheMut.RLock()
	cache, cached := featureCvtermPubInsertCache[key]
	featureCvtermPubInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureCvtermPubColumns,
			featureCvtermPubColumnsWithDefault,
			featureCvtermPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureCvtermPubType, featureCvtermPubMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureCvtermPubType, featureCvtermPubMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_cvterm_pub\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into feature_cvterm_pub")
	}

	if !cached {
		featureCvtermPubInsertCacheMut.Lock()
		featureCvtermPubInsertCache[key] = cache
		featureCvtermPubInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeatureCvtermPub record. See Update for
// whitelist behavior description.
func (o *FeatureCvtermPub) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeatureCvtermPub record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeatureCvtermPub) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeatureCvtermPub, and panics on error.
// See Update for whitelist behavior description.
func (o *FeatureCvtermPub) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeatureCvtermPub.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeatureCvtermPub) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureCvtermPubUpdateCacheMut.RLock()
	cache, cached := featureCvtermPubUpdateCache[key]
	featureCvtermPubUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureCvtermPubColumns, featureCvtermPubPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update feature_cvterm_pub, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_cvterm_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featureCvtermPubPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureCvtermPubType, featureCvtermPubMapping, append(wl, featureCvtermPubPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update feature_cvterm_pub row")
	}

	if !cached {
		featureCvtermPubUpdateCacheMut.Lock()
		featureCvtermPubUpdateCache[key] = cache
		featureCvtermPubUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureCvtermPubQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureCvtermPubQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for feature_cvterm_pub")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureCvtermPubSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureCvtermPubSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureCvtermPubSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureCvtermPubSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_cvterm_pub\" SET %s WHERE (\"feature_cvterm_pub_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureCvtermPubPrimaryKeyColumns), len(colNames)+1, len(featureCvtermPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in featureCvtermPub slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeatureCvtermPub) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeatureCvtermPub) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeatureCvtermPub) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeatureCvtermPub) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_cvterm_pub provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureCvtermPubColumnsWithDefault, o)

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

	featureCvtermPubUpsertCacheMut.RLock()
	cache, cached := featureCvtermPubUpsertCache[key]
	featureCvtermPubUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureCvtermPubColumns,
			featureCvtermPubColumnsWithDefault,
			featureCvtermPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureCvtermPubColumns,
			featureCvtermPubPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert feature_cvterm_pub, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featureCvtermPubPrimaryKeyColumns))
			copy(conflict, featureCvtermPubPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_cvterm_pub\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureCvtermPubType, featureCvtermPubMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureCvtermPubType, featureCvtermPubMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for feature_cvterm_pub")
	}

	if !cached {
		featureCvtermPubUpsertCacheMut.Lock()
		featureCvtermPubUpsertCache[key] = cache
		featureCvtermPubUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeatureCvtermPub record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureCvtermPub) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeatureCvtermPub record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeatureCvtermPub) DeleteG() error {
	if o == nil {
		return errors.New("models: no FeatureCvtermPub provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeatureCvtermPub record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureCvtermPub) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeatureCvtermPub record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeatureCvtermPub) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeatureCvtermPub provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureCvtermPubPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_cvterm_pub\" WHERE \"feature_cvterm_pub_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from feature_cvterm_pub")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureCvtermPubQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureCvtermPubQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no featureCvtermPubQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from feature_cvterm_pub")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureCvtermPubSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureCvtermPubSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no FeatureCvtermPub slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureCvtermPubSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureCvtermPubSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeatureCvtermPub slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureCvtermPubBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_cvterm_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureCvtermPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureCvtermPubPrimaryKeyColumns), 1, len(featureCvtermPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from featureCvtermPub slice")
	}

	if len(featureCvtermPubAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeatureCvtermPub) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeatureCvtermPub) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeatureCvtermPub) ReloadG() error {
	if o == nil {
		return errors.New("models: no FeatureCvtermPub provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeatureCvtermPub) Reload(exec boil.Executor) error {
	ret, err := FindFeatureCvtermPub(exec, o.FeatureCvtermPubID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureCvtermPubSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureCvtermPubSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureCvtermPubSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FeatureCvtermPubSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureCvtermPubSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureCvtermPubs := FeatureCvtermPubSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_cvterm_pub\".* FROM \"feature_cvterm_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureCvtermPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featureCvtermPubPrimaryKeyColumns), 1, len(featureCvtermPubPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureCvtermPubs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeatureCvtermPubSlice")
	}

	*o = featureCvtermPubs

	return nil
}

// FeatureCvtermPubExists checks if the FeatureCvtermPub row exists.
func FeatureCvtermPubExists(exec boil.Executor, featureCvtermPubID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_cvterm_pub\" where \"feature_cvterm_pub_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureCvtermPubID)
	}

	row := exec.QueryRow(sql, featureCvtermPubID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if feature_cvterm_pub exists")
	}

	return exists, nil
}

// FeatureCvtermPubExistsG checks if the FeatureCvtermPub row exists.
func FeatureCvtermPubExistsG(featureCvtermPubID int) (bool, error) {
	return FeatureCvtermPubExists(boil.GetDB(), featureCvtermPubID)
}

// FeatureCvtermPubExistsGP checks if the FeatureCvtermPub row exists. Panics on error.
func FeatureCvtermPubExistsGP(featureCvtermPubID int) bool {
	e, err := FeatureCvtermPubExists(boil.GetDB(), featureCvtermPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureCvtermPubExistsP checks if the FeatureCvtermPub row exists. Panics on error.
func FeatureCvtermPubExistsP(exec boil.Executor, featureCvtermPubID int) bool {
	e, err := FeatureCvtermPubExists(exec, featureCvtermPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
