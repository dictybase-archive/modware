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

// FeaturePub is an object representing the database table.
type FeaturePub struct {
	FeaturePubID int `boil:"feature_pub_id" json:"feature_pub_id" toml:"feature_pub_id" yaml:"feature_pub_id"`
	FeatureID    int `boil:"feature_id" json:"feature_id" toml:"feature_id" yaml:"feature_id"`
	PubID        int `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`

	R *featurePubR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featurePubL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featurePubR is where relationships are stored.
type featurePubR struct {
	Pub            *Pub
	Feature        *Feature
	FeaturePubprop *FeaturePubprop
}

// featurePubL is where Load methods for each relationship are stored.
type featurePubL struct{}

var (
	featurePubColumns               = []string{"feature_pub_id", "feature_id", "pub_id"}
	featurePubColumnsWithoutDefault = []string{"feature_id", "pub_id"}
	featurePubColumnsWithDefault    = []string{"feature_pub_id"}
	featurePubPrimaryKeyColumns     = []string{"feature_pub_id"}
)

type (
	// FeaturePubSlice is an alias for a slice of pointers to FeaturePub.
	// This should generally be used opposed to []FeaturePub.
	FeaturePubSlice []*FeaturePub
	// FeaturePubHook is the signature for custom FeaturePub hook methods
	FeaturePubHook func(boil.Executor, *FeaturePub) error

	featurePubQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featurePubType                 = reflect.TypeOf(&FeaturePub{})
	featurePubMapping              = queries.MakeStructMapping(featurePubType)
	featurePubPrimaryKeyMapping, _ = queries.BindMapping(featurePubType, featurePubMapping, featurePubPrimaryKeyColumns)
	featurePubInsertCacheMut       sync.RWMutex
	featurePubInsertCache          = make(map[string]insertCache)
	featurePubUpdateCacheMut       sync.RWMutex
	featurePubUpdateCache          = make(map[string]updateCache)
	featurePubUpsertCacheMut       sync.RWMutex
	featurePubUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featurePubBeforeInsertHooks []FeaturePubHook
var featurePubBeforeUpdateHooks []FeaturePubHook
var featurePubBeforeDeleteHooks []FeaturePubHook
var featurePubBeforeUpsertHooks []FeaturePubHook

var featurePubAfterInsertHooks []FeaturePubHook
var featurePubAfterSelectHooks []FeaturePubHook
var featurePubAfterUpdateHooks []FeaturePubHook
var featurePubAfterDeleteHooks []FeaturePubHook
var featurePubAfterUpsertHooks []FeaturePubHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeaturePub) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeaturePub) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeaturePub) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeaturePub) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeaturePub) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeaturePub) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeaturePub) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeaturePub) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeaturePub) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePubAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeaturePubHook registers your hook function for all future operations.
func AddFeaturePubHook(hookPoint boil.HookPoint, featurePubHook FeaturePubHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featurePubBeforeInsertHooks = append(featurePubBeforeInsertHooks, featurePubHook)
	case boil.BeforeUpdateHook:
		featurePubBeforeUpdateHooks = append(featurePubBeforeUpdateHooks, featurePubHook)
	case boil.BeforeDeleteHook:
		featurePubBeforeDeleteHooks = append(featurePubBeforeDeleteHooks, featurePubHook)
	case boil.BeforeUpsertHook:
		featurePubBeforeUpsertHooks = append(featurePubBeforeUpsertHooks, featurePubHook)
	case boil.AfterInsertHook:
		featurePubAfterInsertHooks = append(featurePubAfterInsertHooks, featurePubHook)
	case boil.AfterSelectHook:
		featurePubAfterSelectHooks = append(featurePubAfterSelectHooks, featurePubHook)
	case boil.AfterUpdateHook:
		featurePubAfterUpdateHooks = append(featurePubAfterUpdateHooks, featurePubHook)
	case boil.AfterDeleteHook:
		featurePubAfterDeleteHooks = append(featurePubAfterDeleteHooks, featurePubHook)
	case boil.AfterUpsertHook:
		featurePubAfterUpsertHooks = append(featurePubAfterUpsertHooks, featurePubHook)
	}
}

// OneP returns a single featurePub record from the query, and panics on error.
func (q featurePubQuery) OneP() *FeaturePub {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featurePub record from the query.
func (q featurePubQuery) One() (*FeaturePub, error) {
	o := &FeaturePub{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for feature_pub")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeaturePub records from the query, and panics on error.
func (q featurePubQuery) AllP() FeaturePubSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeaturePub records from the query.
func (q featurePubQuery) All() (FeaturePubSlice, error) {
	var o FeaturePubSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to FeaturePub slice")
	}

	if len(featurePubAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeaturePub records in the query, and panics on error.
func (q featurePubQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeaturePub records in the query.
func (q featurePubQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count feature_pub rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featurePubQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featurePubQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if feature_pub exists")
	}

	return count > 0, nil
}

// PubG pointed to by the foreign key.
func (o *FeaturePub) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *FeaturePub) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// FeatureG pointed to by the foreign key.
func (o *FeaturePub) FeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Feature(boil.GetDB(), mods...)
}

// Feature pointed to by the foreign key.
func (o *FeaturePub) Feature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// FeaturePubpropG pointed to by the foreign key.
func (o *FeaturePub) FeaturePubpropG(mods ...qm.QueryMod) featurePubpropQuery {
	return o.FeaturePubprop(boil.GetDB(), mods...)
}

// FeaturePubprop pointed to by the foreign key.
func (o *FeaturePub) FeaturePubprop(exec boil.Executor, mods ...qm.QueryMod) featurePubpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_pub_id=$1", o.FeaturePubID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturePubprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_pubprop\"")

	return query
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurePubL) LoadPub(e boil.Executor, singular bool, maybeFeaturePub interface{}) error {
	var slice []*FeaturePub
	var object *FeaturePub

	count := 1
	if singular {
		object = maybeFeaturePub.(*FeaturePub)
	} else {
		slice = *maybeFeaturePub.(*FeaturePubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurePubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &featurePubR{}
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

	if len(featurePubAfterSelectHooks) != 0 {
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

// LoadFeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurePubL) LoadFeature(e boil.Executor, singular bool, maybeFeaturePub interface{}) error {
	var slice []*FeaturePub
	var object *FeaturePub

	count := 1
	if singular {
		object = maybeFeaturePub.(*FeaturePub)
	} else {
		slice = *maybeFeaturePub.(*FeaturePubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurePubR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featurePubR{}
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

	if len(featurePubAfterSelectHooks) != 0 {
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

// LoadFeaturePubprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurePubL) LoadFeaturePubprop(e boil.Executor, singular bool, maybeFeaturePub interface{}) error {
	var slice []*FeaturePub
	var object *FeaturePub

	count := 1
	if singular {
		object = maybeFeaturePub.(*FeaturePub)
	} else {
		slice = *maybeFeaturePub.(*FeaturePubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurePubR{}
		args[0] = object.FeaturePubID
	} else {
		for i, obj := range slice {
			obj.R = &featurePubR{}
			args[i] = obj.FeaturePubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_pubprop\" where \"feature_pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeaturePubprop")
	}
	defer results.Close()

	var resultSlice []*FeaturePubprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeaturePubprop")
	}

	if len(featurePubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeaturePubprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeaturePubID == foreign.FeaturePubID {
				local.R.FeaturePubprop = foreign
				break
			}
		}
	}

	return nil
}

// SetPub of the feature_pub to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.FeaturePub.
func (o *FeaturePub) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurePubPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.FeaturePubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &featurePubR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			FeaturePub: o,
		}
	} else {
		related.R.FeaturePub = o
	}

	return nil
}

// SetFeature of the feature_pub to the related item.
// Sets o.R.Feature to related.
// Adds o to related.R.FeaturePub.
func (o *FeaturePub) SetFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurePubPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeaturePubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureID = related.FeatureID

	if o.R == nil {
		o.R = &featurePubR{
			Feature: related,
		}
	} else {
		o.R.Feature = related
	}

	if related.R == nil {
		related.R = &featureR{
			FeaturePub: o,
		}
	} else {
		related.R.FeaturePub = o
	}

	return nil
}

// SetFeaturePubprop of the feature_pub to the related item.
// Sets o.R.FeaturePubprop to related.
// Adds o to related.R.FeaturePub.
func (o *FeaturePub) SetFeaturePubprop(exec boil.Executor, insert bool, related *FeaturePubprop) error {
	var err error

	if insert {
		related.FeaturePubID = o.FeaturePubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_pubprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurePubpropPrimaryKeyColumns),
		)
		values := []interface{}{o.FeaturePubID, related.FeaturePubpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeaturePubID = o.FeaturePubID

	}

	if o.R == nil {
		o.R = &featurePubR{
			FeaturePubprop: related,
		}
	} else {
		o.R.FeaturePubprop = related
	}

	if related.R == nil {
		related.R = &featurePubpropR{
			FeaturePub: o,
		}
	} else {
		related.R.FeaturePub = o
	}
	return nil
}

// FeaturePubsG retrieves all records.
func FeaturePubsG(mods ...qm.QueryMod) featurePubQuery {
	return FeaturePubs(boil.GetDB(), mods...)
}

// FeaturePubs retrieves all the records using an executor.
func FeaturePubs(exec boil.Executor, mods ...qm.QueryMod) featurePubQuery {
	mods = append(mods, qm.From("\"feature_pub\""))
	return featurePubQuery{NewQuery(exec, mods...)}
}

// FindFeaturePubG retrieves a single record by ID.
func FindFeaturePubG(featurePubID int, selectCols ...string) (*FeaturePub, error) {
	return FindFeaturePub(boil.GetDB(), featurePubID, selectCols...)
}

// FindFeaturePubGP retrieves a single record by ID, and panics on error.
func FindFeaturePubGP(featurePubID int, selectCols ...string) *FeaturePub {
	retobj, err := FindFeaturePub(boil.GetDB(), featurePubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeaturePub retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeaturePub(exec boil.Executor, featurePubID int, selectCols ...string) (*FeaturePub, error) {
	featurePubObj := &FeaturePub{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_pub\" where \"feature_pub_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featurePubID)

	err := q.Bind(featurePubObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from feature_pub")
	}

	return featurePubObj, nil
}

// FindFeaturePubP retrieves a single record by ID with an executor, and panics on error.
func FindFeaturePubP(exec boil.Executor, featurePubID int, selectCols ...string) *FeaturePub {
	retobj, err := FindFeaturePub(exec, featurePubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeaturePub) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeaturePub) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeaturePub) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeaturePub) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_pub provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurePubColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featurePubInsertCacheMut.RLock()
	cache, cached := featurePubInsertCache[key]
	featurePubInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featurePubColumns,
			featurePubColumnsWithDefault,
			featurePubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featurePubType, featurePubMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featurePubType, featurePubMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_pub\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into feature_pub")
	}

	if !cached {
		featurePubInsertCacheMut.Lock()
		featurePubInsertCache[key] = cache
		featurePubInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeaturePub record. See Update for
// whitelist behavior description.
func (o *FeaturePub) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeaturePub record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeaturePub) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeaturePub, and panics on error.
// See Update for whitelist behavior description.
func (o *FeaturePub) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeaturePub.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeaturePub) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featurePubUpdateCacheMut.RLock()
	cache, cached := featurePubUpdateCache[key]
	featurePubUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featurePubColumns, featurePubPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update feature_pub, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featurePubPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featurePubType, featurePubMapping, append(wl, featurePubPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update feature_pub row")
	}

	if !cached {
		featurePubUpdateCacheMut.Lock()
		featurePubUpdateCache[key] = cache
		featurePubUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featurePubQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featurePubQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for feature_pub")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeaturePubSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeaturePubSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeaturePubSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeaturePubSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_pub\" SET %s WHERE (\"feature_pub_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurePubPrimaryKeyColumns), len(colNames)+1, len(featurePubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in featurePub slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeaturePub) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeaturePub) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeaturePub) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeaturePub) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_pub provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurePubColumnsWithDefault, o)

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

	featurePubUpsertCacheMut.RLock()
	cache, cached := featurePubUpsertCache[key]
	featurePubUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featurePubColumns,
			featurePubColumnsWithDefault,
			featurePubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featurePubColumns,
			featurePubPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert feature_pub, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featurePubPrimaryKeyColumns))
			copy(conflict, featurePubPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_pub\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featurePubType, featurePubMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featurePubType, featurePubMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for feature_pub")
	}

	if !cached {
		featurePubUpsertCacheMut.Lock()
		featurePubUpsertCache[key] = cache
		featurePubUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeaturePub record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeaturePub) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeaturePub record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeaturePub) DeleteG() error {
	if o == nil {
		return errors.New("chado: no FeaturePub provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeaturePub record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeaturePub) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeaturePub record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeaturePub) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeaturePub provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featurePubPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_pub\" WHERE \"feature_pub_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from feature_pub")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featurePubQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featurePubQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no featurePubQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from feature_pub")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeaturePubSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeaturePubSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no FeaturePub slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeaturePubSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeaturePubSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeaturePub slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featurePubBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurePubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurePubPrimaryKeyColumns), 1, len(featurePubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featurePub slice")
	}

	if len(featurePubAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeaturePub) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeaturePub) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeaturePub) ReloadG() error {
	if o == nil {
		return errors.New("chado: no FeaturePub provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeaturePub) Reload(exec boil.Executor) error {
	ret, err := FindFeaturePub(exec, o.FeaturePubID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturePubSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturePubSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturePubSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty FeaturePubSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturePubSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featurePubs := FeaturePubSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_pub\".* FROM \"feature_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurePubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featurePubPrimaryKeyColumns), 1, len(featurePubPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featurePubs)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in FeaturePubSlice")
	}

	*o = featurePubs

	return nil
}

// FeaturePubExists checks if the FeaturePub row exists.
func FeaturePubExists(exec boil.Executor, featurePubID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_pub\" where \"feature_pub_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featurePubID)
	}

	row := exec.QueryRow(sql, featurePubID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if feature_pub exists")
	}

	return exists, nil
}

// FeaturePubExistsG checks if the FeaturePub row exists.
func FeaturePubExistsG(featurePubID int) (bool, error) {
	return FeaturePubExists(boil.GetDB(), featurePubID)
}

// FeaturePubExistsGP checks if the FeaturePub row exists. Panics on error.
func FeaturePubExistsGP(featurePubID int) bool {
	e, err := FeaturePubExists(boil.GetDB(), featurePubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeaturePubExistsP checks if the FeaturePub row exists. Panics on error.
func FeaturePubExistsP(exec boil.Executor, featurePubID int) bool {
	e, err := FeaturePubExists(exec, featurePubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
