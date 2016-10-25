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

// FeaturePubprop is an object representing the database table.
type FeaturePubprop struct {
	FeaturePubpropID int         `boil:"feature_pubprop_id" json:"feature_pubprop_id" toml:"feature_pubprop_id" yaml:"feature_pubprop_id"`
	FeaturePubID     int         `boil:"feature_pub_id" json:"feature_pub_id" toml:"feature_pub_id" yaml:"feature_pub_id"`
	TypeID           int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value            null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank             int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *featurePubpropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featurePubpropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featurePubpropR is where relationships are stored.
type featurePubpropR struct {
	FeaturePub *FeaturePub
	Type       *Cvterm
}

// featurePubpropL is where Load methods for each relationship are stored.
type featurePubpropL struct{}

var (
	featurePubpropColumns               = []string{"feature_pubprop_id", "feature_pub_id", "type_id", "value", "rank"}
	featurePubpropColumnsWithoutDefault = []string{"feature_pub_id", "type_id", "value"}
	featurePubpropColumnsWithDefault    = []string{"feature_pubprop_id", "rank"}
	featurePubpropPrimaryKeyColumns     = []string{"feature_pubprop_id"}
)

type (
	// FeaturePubpropSlice is an alias for a slice of pointers to FeaturePubprop.
	// This should generally be used opposed to []FeaturePubprop.
	FeaturePubpropSlice []*FeaturePubprop
	// FeaturePubpropHook is the signature for custom FeaturePubprop hook methods
	FeaturePubpropHook func(boil.Executor, *FeaturePubprop) error

	featurePubpropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featurePubpropType                 = reflect.TypeOf(&FeaturePubprop{})
	featurePubpropMapping              = queries.MakeStructMapping(featurePubpropType)
	featurePubpropPrimaryKeyMapping, _ = queries.BindMapping(featurePubpropType, featurePubpropMapping, featurePubpropPrimaryKeyColumns)
	featurePubpropInsertCacheMut       sync.RWMutex
	featurePubpropInsertCache          = make(map[string]insertCache)
	featurePubpropUpdateCacheMut       sync.RWMutex
	featurePubpropUpdateCache          = make(map[string]updateCache)
	featurePubpropUpsertCacheMut       sync.RWMutex
	featurePubpropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featurePubpropBeforeInsertHooks []FeaturePubpropHook
var featurePubpropBeforeUpdateHooks []FeaturePubpropHook
var featurePubpropBeforeDeleteHooks []FeaturePubpropHook
var featurePubpropBeforeUpsertHooks []FeaturePubpropHook

var featurePubpropAfterInsertHooks []FeaturePubpropHook
var featurePubpropAfterSelectHooks []FeaturePubpropHook
var featurePubpropAfterUpdateHooks []FeaturePubpropHook
var featurePubpropAfterDeleteHooks []FeaturePubpropHook
var featurePubpropAfterUpsertHooks []FeaturePubpropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeaturePubprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubpropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeaturePubprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubpropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeaturePubprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubpropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeaturePubprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubpropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeaturePubprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubpropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeaturePubprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubpropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeaturePubprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubpropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeaturePubprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubpropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeaturePubprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubpropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeaturePubpropHook registers your hook function for all future operations.
func AddFeaturePubpropHook(hookPoint boil.HookPoint, featurePubpropHook FeaturePubpropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featurePubpropBeforeInsertHooks = append(featurePubpropBeforeInsertHooks, featurePubpropHook)
	case boil.BeforeUpdateHook:
		featurePubpropBeforeUpdateHooks = append(featurePubpropBeforeUpdateHooks, featurePubpropHook)
	case boil.BeforeDeleteHook:
		featurePubpropBeforeDeleteHooks = append(featurePubpropBeforeDeleteHooks, featurePubpropHook)
	case boil.BeforeUpsertHook:
		featurePubpropBeforeUpsertHooks = append(featurePubpropBeforeUpsertHooks, featurePubpropHook)
	case boil.AfterInsertHook:
		featurePubpropAfterInsertHooks = append(featurePubpropAfterInsertHooks, featurePubpropHook)
	case boil.AfterSelectHook:
		featurePubpropAfterSelectHooks = append(featurePubpropAfterSelectHooks, featurePubpropHook)
	case boil.AfterUpdateHook:
		featurePubpropAfterUpdateHooks = append(featurePubpropAfterUpdateHooks, featurePubpropHook)
	case boil.AfterDeleteHook:
		featurePubpropAfterDeleteHooks = append(featurePubpropAfterDeleteHooks, featurePubpropHook)
	case boil.AfterUpsertHook:
		featurePubpropAfterUpsertHooks = append(featurePubpropAfterUpsertHooks, featurePubpropHook)
	}
}

// OneP returns a single featurePubprop record from the query, and panics on error.
func (q featurePubpropQuery) OneP() *FeaturePubprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featurePubprop record from the query.
func (q featurePubpropQuery) One() (*FeaturePubprop, error) {
	o := &FeaturePubprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for feature_pubprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeaturePubprop records from the query, and panics on error.
func (q featurePubpropQuery) AllP() FeaturePubpropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeaturePubprop records from the query.
func (q featurePubpropQuery) All() (FeaturePubpropSlice, error) {
	var o FeaturePubpropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FeaturePubprop slice")
	}

	if len(featurePubpropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeaturePubprop records in the query, and panics on error.
func (q featurePubpropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeaturePubprop records in the query.
func (q featurePubpropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count feature_pubprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featurePubpropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featurePubpropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if feature_pubprop exists")
	}

	return count > 0, nil
}

// FeaturePubG pointed to by the foreign key.
func (o *FeaturePubprop) FeaturePubG(mods ...qm.QueryMod) featurePubQuery {
	return o.FeaturePub(boil.GetDB(), mods...)
}

// FeaturePub pointed to by the foreign key.
func (o *FeaturePubprop) FeaturePub(exec boil.Executor, mods ...qm.QueryMod) featurePubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_pub_id=$1", o.FeaturePubID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturePubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_pub\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *FeaturePubprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *FeaturePubprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadFeaturePub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurePubpropL) LoadFeaturePub(e boil.Executor, singular bool, maybeFeaturePubprop interface{}) error {
	var slice []*FeaturePubprop
	var object *FeaturePubprop

	count := 1
	if singular {
		object = maybeFeaturePubprop.(*FeaturePubprop)
	} else {
		slice = *maybeFeaturePubprop.(*FeaturePubpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurePubpropR{}
		args[0] = object.FeaturePubID
	} else {
		for i, obj := range slice {
			obj.R = &featurePubpropR{}
			args[i] = obj.FeaturePubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_pub\" where \"feature_pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeaturePub")
	}
	defer results.Close()

	var resultSlice []*FeaturePub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeaturePub")
	}

	if len(featurePubpropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeaturePub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeaturePubID == foreign.FeaturePubID {
				local.R.FeaturePub = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurePubpropL) LoadType(e boil.Executor, singular bool, maybeFeaturePubprop interface{}) error {
	var slice []*FeaturePubprop
	var object *FeaturePubprop

	count := 1
	if singular {
		object = maybeFeaturePubprop.(*FeaturePubprop)
	} else {
		slice = *maybeFeaturePubprop.(*FeaturePubpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurePubpropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &featurePubpropR{}
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

	if len(featurePubpropAfterSelectHooks) != 0 {
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
			if local.TypeID == foreign.CvtermID {
				local.R.Type = foreign
				break
			}
		}
	}

	return nil
}

// SetFeaturePub of the feature_pubprop to the related item.
// Sets o.R.FeaturePub to related.
// Adds o to related.R.FeaturePubprop.
func (o *FeaturePubprop) SetFeaturePub(exec boil.Executor, insert bool, related *FeaturePub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_pubprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurePubpropPrimaryKeyColumns),
	)
	values := []interface{}{related.FeaturePubID, o.FeaturePubpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeaturePubID = related.FeaturePubID

	if o.R == nil {
		o.R = &featurePubpropR{
			FeaturePub: related,
		}
	} else {
		o.R.FeaturePub = related
	}

	if related.R == nil {
		related.R = &featurePubR{
			FeaturePubprop: o,
		}
	} else {
		related.R.FeaturePubprop = o
	}

	return nil
}

// SetType of the feature_pubprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeFeaturePubprop.
func (o *FeaturePubprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_pubprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurePubpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.FeaturePubpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &featurePubpropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeFeaturePubprop: o,
		}
	} else {
		related.R.TypeFeaturePubprop = o
	}

	return nil
}

// FeaturePubpropsG retrieves all records.
func FeaturePubpropsG(mods ...qm.QueryMod) featurePubpropQuery {
	return FeaturePubprops(boil.GetDB(), mods...)
}

// FeaturePubprops retrieves all the records using an executor.
func FeaturePubprops(exec boil.Executor, mods ...qm.QueryMod) featurePubpropQuery {
	mods = append(mods, qm.From("\"feature_pubprop\""))
	return featurePubpropQuery{NewQuery(exec, mods...)}
}

// FindFeaturePubpropG retrieves a single record by ID.
func FindFeaturePubpropG(featurePubpropID int, selectCols ...string) (*FeaturePubprop, error) {
	return FindFeaturePubprop(boil.GetDB(), featurePubpropID, selectCols...)
}

// FindFeaturePubpropGP retrieves a single record by ID, and panics on error.
func FindFeaturePubpropGP(featurePubpropID int, selectCols ...string) *FeaturePubprop {
	retobj, err := FindFeaturePubprop(boil.GetDB(), featurePubpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeaturePubprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeaturePubprop(exec boil.Executor, featurePubpropID int, selectCols ...string) (*FeaturePubprop, error) {
	featurePubpropObj := &FeaturePubprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_pubprop\" where \"feature_pubprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featurePubpropID)

	err := q.Bind(featurePubpropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from feature_pubprop")
	}

	return featurePubpropObj, nil
}

// FindFeaturePubpropP retrieves a single record by ID with an executor, and panics on error.
func FindFeaturePubpropP(exec boil.Executor, featurePubpropID int, selectCols ...string) *FeaturePubprop {
	retobj, err := FindFeaturePubprop(exec, featurePubpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeaturePubprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeaturePubprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeaturePubprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeaturePubprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_pubprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurePubpropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featurePubpropInsertCacheMut.RLock()
	cache, cached := featurePubpropInsertCache[key]
	featurePubpropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featurePubpropColumns,
			featurePubpropColumnsWithDefault,
			featurePubpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featurePubpropType, featurePubpropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featurePubpropType, featurePubpropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_pubprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into feature_pubprop")
	}

	if !cached {
		featurePubpropInsertCacheMut.Lock()
		featurePubpropInsertCache[key] = cache
		featurePubpropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeaturePubprop record. See Update for
// whitelist behavior description.
func (o *FeaturePubprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeaturePubprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeaturePubprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeaturePubprop, and panics on error.
// See Update for whitelist behavior description.
func (o *FeaturePubprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeaturePubprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeaturePubprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featurePubpropUpdateCacheMut.RLock()
	cache, cached := featurePubpropUpdateCache[key]
	featurePubpropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featurePubpropColumns, featurePubpropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update feature_pubprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_pubprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featurePubpropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featurePubpropType, featurePubpropMapping, append(wl, featurePubpropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update feature_pubprop row")
	}

	if !cached {
		featurePubpropUpdateCacheMut.Lock()
		featurePubpropUpdateCache[key] = cache
		featurePubpropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featurePubpropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featurePubpropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for feature_pubprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeaturePubpropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeaturePubpropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeaturePubpropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeaturePubpropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePubpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_pubprop\" SET %s WHERE (\"feature_pubprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurePubpropPrimaryKeyColumns), len(colNames)+1, len(featurePubpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in featurePubprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeaturePubprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeaturePubprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeaturePubprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeaturePubprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_pubprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurePubpropColumnsWithDefault, o)

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

	featurePubpropUpsertCacheMut.RLock()
	cache, cached := featurePubpropUpsertCache[key]
	featurePubpropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featurePubpropColumns,
			featurePubpropColumnsWithDefault,
			featurePubpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featurePubpropColumns,
			featurePubpropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert feature_pubprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featurePubpropPrimaryKeyColumns))
			copy(conflict, featurePubpropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_pubprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featurePubpropType, featurePubpropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featurePubpropType, featurePubpropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for feature_pubprop")
	}

	if !cached {
		featurePubpropUpsertCacheMut.Lock()
		featurePubpropUpsertCache[key] = cache
		featurePubpropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeaturePubprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeaturePubprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeaturePubprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeaturePubprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no FeaturePubprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeaturePubprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeaturePubprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeaturePubprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeaturePubprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeaturePubprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featurePubpropPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_pubprop\" WHERE \"feature_pubprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from feature_pubprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featurePubpropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featurePubpropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no featurePubpropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from feature_pubprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeaturePubpropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeaturePubpropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no FeaturePubprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeaturePubpropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeaturePubpropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeaturePubprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featurePubpropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePubpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_pubprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurePubpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurePubpropPrimaryKeyColumns), 1, len(featurePubpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from featurePubprop slice")
	}

	if len(featurePubpropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeaturePubprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeaturePubprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeaturePubprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no FeaturePubprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeaturePubprop) Reload(exec boil.Executor) error {
	ret, err := FindFeaturePubprop(exec, o.FeaturePubpropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturePubpropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturePubpropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturePubpropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FeaturePubpropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturePubpropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featurePubprops := FeaturePubpropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePubpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_pubprop\".* FROM \"feature_pubprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurePubpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featurePubpropPrimaryKeyColumns), 1, len(featurePubpropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featurePubprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeaturePubpropSlice")
	}

	*o = featurePubprops

	return nil
}

// FeaturePubpropExists checks if the FeaturePubprop row exists.
func FeaturePubpropExists(exec boil.Executor, featurePubpropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_pubprop\" where \"feature_pubprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featurePubpropID)
	}

	row := exec.QueryRow(sql, featurePubpropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if feature_pubprop exists")
	}

	return exists, nil
}

// FeaturePubpropExistsG checks if the FeaturePubprop row exists.
func FeaturePubpropExistsG(featurePubpropID int) (bool, error) {
	return FeaturePubpropExists(boil.GetDB(), featurePubpropID)
}

// FeaturePubpropExistsGP checks if the FeaturePubprop row exists. Panics on error.
func FeaturePubpropExistsGP(featurePubpropID int) bool {
	e, err := FeaturePubpropExists(boil.GetDB(), featurePubpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeaturePubpropExistsP checks if the FeaturePubprop row exists. Panics on error.
func FeaturePubpropExistsP(exec boil.Executor, featurePubpropID int) bool {
	e, err := FeaturePubpropExists(exec, featurePubpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
