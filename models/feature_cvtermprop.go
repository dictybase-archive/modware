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

// FeatureCvtermprop is an object representing the database table.
type FeatureCvtermprop struct {
	FeatureCvtermpropID int         `boil:"feature_cvtermprop_id" json:"feature_cvtermprop_id" toml:"feature_cvtermprop_id" yaml:"feature_cvtermprop_id"`
	FeatureCvtermID     int         `boil:"feature_cvterm_id" json:"feature_cvterm_id" toml:"feature_cvterm_id" yaml:"feature_cvterm_id"`
	TypeID              int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value               null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank                int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *featureCvtermpropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureCvtermpropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureCvtermpropR is where relationships are stored.
type featureCvtermpropR struct {
	FeatureCvterm *FeatureCvterm
	Type          *Cvterm
}

// featureCvtermpropL is where Load methods for each relationship are stored.
type featureCvtermpropL struct{}

var (
	featureCvtermpropColumns               = []string{"feature_cvtermprop_id", "feature_cvterm_id", "type_id", "value", "rank"}
	featureCvtermpropColumnsWithoutDefault = []string{"feature_cvterm_id", "type_id", "value"}
	featureCvtermpropColumnsWithDefault    = []string{"feature_cvtermprop_id", "rank"}
	featureCvtermpropPrimaryKeyColumns     = []string{"feature_cvtermprop_id"}
)

type (
	// FeatureCvtermpropSlice is an alias for a slice of pointers to FeatureCvtermprop.
	// This should generally be used opposed to []FeatureCvtermprop.
	FeatureCvtermpropSlice []*FeatureCvtermprop
	// FeatureCvtermpropHook is the signature for custom FeatureCvtermprop hook methods
	FeatureCvtermpropHook func(boil.Executor, *FeatureCvtermprop) error

	featureCvtermpropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureCvtermpropType                 = reflect.TypeOf(&FeatureCvtermprop{})
	featureCvtermpropMapping              = queries.MakeStructMapping(featureCvtermpropType)
	featureCvtermpropPrimaryKeyMapping, _ = queries.BindMapping(featureCvtermpropType, featureCvtermpropMapping, featureCvtermpropPrimaryKeyColumns)
	featureCvtermpropInsertCacheMut       sync.RWMutex
	featureCvtermpropInsertCache          = make(map[string]insertCache)
	featureCvtermpropUpdateCacheMut       sync.RWMutex
	featureCvtermpropUpdateCache          = make(map[string]updateCache)
	featureCvtermpropUpsertCacheMut       sync.RWMutex
	featureCvtermpropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureCvtermpropBeforeInsertHooks []FeatureCvtermpropHook
var featureCvtermpropBeforeUpdateHooks []FeatureCvtermpropHook
var featureCvtermpropBeforeDeleteHooks []FeatureCvtermpropHook
var featureCvtermpropBeforeUpsertHooks []FeatureCvtermpropHook

var featureCvtermpropAfterInsertHooks []FeatureCvtermpropHook
var featureCvtermpropAfterSelectHooks []FeatureCvtermpropHook
var featureCvtermpropAfterUpdateHooks []FeatureCvtermpropHook
var featureCvtermpropAfterDeleteHooks []FeatureCvtermpropHook
var featureCvtermpropAfterUpsertHooks []FeatureCvtermpropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeatureCvtermprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermpropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeatureCvtermprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermpropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeatureCvtermprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermpropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeatureCvtermprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermpropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeatureCvtermprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermpropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeatureCvtermprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermpropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeatureCvtermprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermpropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeatureCvtermprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermpropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeatureCvtermprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermpropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureCvtermpropHook registers your hook function for all future operations.
func AddFeatureCvtermpropHook(hookPoint boil.HookPoint, featureCvtermpropHook FeatureCvtermpropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureCvtermpropBeforeInsertHooks = append(featureCvtermpropBeforeInsertHooks, featureCvtermpropHook)
	case boil.BeforeUpdateHook:
		featureCvtermpropBeforeUpdateHooks = append(featureCvtermpropBeforeUpdateHooks, featureCvtermpropHook)
	case boil.BeforeDeleteHook:
		featureCvtermpropBeforeDeleteHooks = append(featureCvtermpropBeforeDeleteHooks, featureCvtermpropHook)
	case boil.BeforeUpsertHook:
		featureCvtermpropBeforeUpsertHooks = append(featureCvtermpropBeforeUpsertHooks, featureCvtermpropHook)
	case boil.AfterInsertHook:
		featureCvtermpropAfterInsertHooks = append(featureCvtermpropAfterInsertHooks, featureCvtermpropHook)
	case boil.AfterSelectHook:
		featureCvtermpropAfterSelectHooks = append(featureCvtermpropAfterSelectHooks, featureCvtermpropHook)
	case boil.AfterUpdateHook:
		featureCvtermpropAfterUpdateHooks = append(featureCvtermpropAfterUpdateHooks, featureCvtermpropHook)
	case boil.AfterDeleteHook:
		featureCvtermpropAfterDeleteHooks = append(featureCvtermpropAfterDeleteHooks, featureCvtermpropHook)
	case boil.AfterUpsertHook:
		featureCvtermpropAfterUpsertHooks = append(featureCvtermpropAfterUpsertHooks, featureCvtermpropHook)
	}
}

// OneP returns a single featureCvtermprop record from the query, and panics on error.
func (q featureCvtermpropQuery) OneP() *FeatureCvtermprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureCvtermprop record from the query.
func (q featureCvtermpropQuery) One() (*FeatureCvtermprop, error) {
	o := &FeatureCvtermprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for feature_cvtermprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeatureCvtermprop records from the query, and panics on error.
func (q featureCvtermpropQuery) AllP() FeatureCvtermpropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeatureCvtermprop records from the query.
func (q featureCvtermpropQuery) All() (FeatureCvtermpropSlice, error) {
	var o FeatureCvtermpropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FeatureCvtermprop slice")
	}

	if len(featureCvtermpropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeatureCvtermprop records in the query, and panics on error.
func (q featureCvtermpropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeatureCvtermprop records in the query.
func (q featureCvtermpropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count feature_cvtermprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureCvtermpropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureCvtermpropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if feature_cvtermprop exists")
	}

	return count > 0, nil
}

// FeatureCvtermG pointed to by the foreign key.
func (o *FeatureCvtermprop) FeatureCvtermG(mods ...qm.QueryMod) featureCvtermQuery {
	return o.FeatureCvterm(boil.GetDB(), mods...)
}

// FeatureCvterm pointed to by the foreign key.
func (o *FeatureCvtermprop) FeatureCvterm(exec boil.Executor, mods ...qm.QueryMod) featureCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_cvterm_id=$1", o.FeatureCvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvterm\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *FeatureCvtermprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *FeatureCvtermprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadFeatureCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermpropL) LoadFeatureCvterm(e boil.Executor, singular bool, maybeFeatureCvtermprop interface{}) error {
	var slice []*FeatureCvtermprop
	var object *FeatureCvtermprop

	count := 1
	if singular {
		object = maybeFeatureCvtermprop.(*FeatureCvtermprop)
	} else {
		slice = *maybeFeatureCvtermprop.(*FeatureCvtermpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermpropR{}
		args[0] = object.FeatureCvtermID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermpropR{}
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

	if len(featureCvtermpropAfterSelectHooks) != 0 {
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

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermpropL) LoadType(e boil.Executor, singular bool, maybeFeatureCvtermprop interface{}) error {
	var slice []*FeatureCvtermprop
	var object *FeatureCvtermprop

	count := 1
	if singular {
		object = maybeFeatureCvtermprop.(*FeatureCvtermprop)
	} else {
		slice = *maybeFeatureCvtermprop.(*FeatureCvtermpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermpropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermpropR{}
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

	if len(featureCvtermpropAfterSelectHooks) != 0 {
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

// SetFeatureCvterm of the feature_cvtermprop to the related item.
// Sets o.R.FeatureCvterm to related.
// Adds o to related.R.FeatureCvtermprop.
func (o *FeatureCvtermprop) SetFeatureCvterm(exec boil.Executor, insert bool, related *FeatureCvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_cvtermprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureCvtermpropPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureCvtermID, o.FeatureCvtermpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureCvtermID = related.FeatureCvtermID

	if o.R == nil {
		o.R = &featureCvtermpropR{
			FeatureCvterm: related,
		}
	} else {
		o.R.FeatureCvterm = related
	}

	if related.R == nil {
		related.R = &featureCvtermR{
			FeatureCvtermprop: o,
		}
	} else {
		related.R.FeatureCvtermprop = o
	}

	return nil
}

// SetType of the feature_cvtermprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeFeatureCvtermprop.
func (o *FeatureCvtermprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_cvtermprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureCvtermpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.FeatureCvtermpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &featureCvtermpropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeFeatureCvtermprop: o,
		}
	} else {
		related.R.TypeFeatureCvtermprop = o
	}

	return nil
}

// FeatureCvtermpropsG retrieves all records.
func FeatureCvtermpropsG(mods ...qm.QueryMod) featureCvtermpropQuery {
	return FeatureCvtermprops(boil.GetDB(), mods...)
}

// FeatureCvtermprops retrieves all the records using an executor.
func FeatureCvtermprops(exec boil.Executor, mods ...qm.QueryMod) featureCvtermpropQuery {
	mods = append(mods, qm.From("\"feature_cvtermprop\""))
	return featureCvtermpropQuery{NewQuery(exec, mods...)}
}

// FindFeatureCvtermpropG retrieves a single record by ID.
func FindFeatureCvtermpropG(featureCvtermpropID int, selectCols ...string) (*FeatureCvtermprop, error) {
	return FindFeatureCvtermprop(boil.GetDB(), featureCvtermpropID, selectCols...)
}

// FindFeatureCvtermpropGP retrieves a single record by ID, and panics on error.
func FindFeatureCvtermpropGP(featureCvtermpropID int, selectCols ...string) *FeatureCvtermprop {
	retobj, err := FindFeatureCvtermprop(boil.GetDB(), featureCvtermpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureCvtermprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureCvtermprop(exec boil.Executor, featureCvtermpropID int, selectCols ...string) (*FeatureCvtermprop, error) {
	featureCvtermpropObj := &FeatureCvtermprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_cvtermprop\" where \"feature_cvtermprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureCvtermpropID)

	err := q.Bind(featureCvtermpropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from feature_cvtermprop")
	}

	return featureCvtermpropObj, nil
}

// FindFeatureCvtermpropP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureCvtermpropP(exec boil.Executor, featureCvtermpropID int, selectCols ...string) *FeatureCvtermprop {
	retobj, err := FindFeatureCvtermprop(exec, featureCvtermpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeatureCvtermprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeatureCvtermprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeatureCvtermprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeatureCvtermprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_cvtermprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureCvtermpropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureCvtermpropInsertCacheMut.RLock()
	cache, cached := featureCvtermpropInsertCache[key]
	featureCvtermpropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureCvtermpropColumns,
			featureCvtermpropColumnsWithDefault,
			featureCvtermpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureCvtermpropType, featureCvtermpropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureCvtermpropType, featureCvtermpropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_cvtermprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into feature_cvtermprop")
	}

	if !cached {
		featureCvtermpropInsertCacheMut.Lock()
		featureCvtermpropInsertCache[key] = cache
		featureCvtermpropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeatureCvtermprop record. See Update for
// whitelist behavior description.
func (o *FeatureCvtermprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeatureCvtermprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeatureCvtermprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeatureCvtermprop, and panics on error.
// See Update for whitelist behavior description.
func (o *FeatureCvtermprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeatureCvtermprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeatureCvtermprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureCvtermpropUpdateCacheMut.RLock()
	cache, cached := featureCvtermpropUpdateCache[key]
	featureCvtermpropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureCvtermpropColumns, featureCvtermpropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update feature_cvtermprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_cvtermprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featureCvtermpropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureCvtermpropType, featureCvtermpropMapping, append(wl, featureCvtermpropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update feature_cvtermprop row")
	}

	if !cached {
		featureCvtermpropUpdateCacheMut.Lock()
		featureCvtermpropUpdateCache[key] = cache
		featureCvtermpropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureCvtermpropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureCvtermpropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for feature_cvtermprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureCvtermpropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureCvtermpropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureCvtermpropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureCvtermpropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_cvtermprop\" SET %s WHERE (\"feature_cvtermprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureCvtermpropPrimaryKeyColumns), len(colNames)+1, len(featureCvtermpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in featureCvtermprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeatureCvtermprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeatureCvtermprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeatureCvtermprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeatureCvtermprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_cvtermprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureCvtermpropColumnsWithDefault, o)

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

	featureCvtermpropUpsertCacheMut.RLock()
	cache, cached := featureCvtermpropUpsertCache[key]
	featureCvtermpropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureCvtermpropColumns,
			featureCvtermpropColumnsWithDefault,
			featureCvtermpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureCvtermpropColumns,
			featureCvtermpropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert feature_cvtermprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featureCvtermpropPrimaryKeyColumns))
			copy(conflict, featureCvtermpropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_cvtermprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureCvtermpropType, featureCvtermpropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureCvtermpropType, featureCvtermpropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for feature_cvtermprop")
	}

	if !cached {
		featureCvtermpropUpsertCacheMut.Lock()
		featureCvtermpropUpsertCache[key] = cache
		featureCvtermpropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeatureCvtermprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureCvtermprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeatureCvtermprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeatureCvtermprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no FeatureCvtermprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeatureCvtermprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureCvtermprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeatureCvtermprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeatureCvtermprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeatureCvtermprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureCvtermpropPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_cvtermprop\" WHERE \"feature_cvtermprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from feature_cvtermprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureCvtermpropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureCvtermpropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no featureCvtermpropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from feature_cvtermprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureCvtermpropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureCvtermpropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no FeatureCvtermprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureCvtermpropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureCvtermpropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeatureCvtermprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureCvtermpropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_cvtermprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureCvtermpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureCvtermpropPrimaryKeyColumns), 1, len(featureCvtermpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from featureCvtermprop slice")
	}

	if len(featureCvtermpropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeatureCvtermprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeatureCvtermprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeatureCvtermprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no FeatureCvtermprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeatureCvtermprop) Reload(exec boil.Executor) error {
	ret, err := FindFeatureCvtermprop(exec, o.FeatureCvtermpropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureCvtermpropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureCvtermpropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureCvtermpropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FeatureCvtermpropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureCvtermpropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureCvtermprops := FeatureCvtermpropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_cvtermprop\".* FROM \"feature_cvtermprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureCvtermpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featureCvtermpropPrimaryKeyColumns), 1, len(featureCvtermpropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureCvtermprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeatureCvtermpropSlice")
	}

	*o = featureCvtermprops

	return nil
}

// FeatureCvtermpropExists checks if the FeatureCvtermprop row exists.
func FeatureCvtermpropExists(exec boil.Executor, featureCvtermpropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_cvtermprop\" where \"feature_cvtermprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureCvtermpropID)
	}

	row := exec.QueryRow(sql, featureCvtermpropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if feature_cvtermprop exists")
	}

	return exists, nil
}

// FeatureCvtermpropExistsG checks if the FeatureCvtermprop row exists.
func FeatureCvtermpropExistsG(featureCvtermpropID int) (bool, error) {
	return FeatureCvtermpropExists(boil.GetDB(), featureCvtermpropID)
}

// FeatureCvtermpropExistsGP checks if the FeatureCvtermprop row exists. Panics on error.
func FeatureCvtermpropExistsGP(featureCvtermpropID int) bool {
	e, err := FeatureCvtermpropExists(boil.GetDB(), featureCvtermpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureCvtermpropExistsP checks if the FeatureCvtermprop row exists. Panics on error.
func FeatureCvtermpropExistsP(exec boil.Executor, featureCvtermpropID int) bool {
	e, err := FeatureCvtermpropExists(exec, featureCvtermpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
