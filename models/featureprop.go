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

// Featureprop is an object representing the database table.
type Featureprop struct {
	FeaturepropID int         `boil:"featureprop_id" json:"featureprop_id" toml:"featureprop_id" yaml:"featureprop_id"`
	FeatureID     int         `boil:"feature_id" json:"feature_id" toml:"feature_id" yaml:"feature_id"`
	TypeID        int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value         null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank          int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *featurepropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featurepropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featurepropR is where relationships are stored.
type featurepropR struct {
	Type           *Cvterm
	Feature        *Feature
	FeaturepropPub *FeaturepropPub
}

// featurepropL is where Load methods for each relationship are stored.
type featurepropL struct{}

var (
	featurepropColumns               = []string{"featureprop_id", "feature_id", "type_id", "value", "rank"}
	featurepropColumnsWithoutDefault = []string{"feature_id", "type_id", "value"}
	featurepropColumnsWithDefault    = []string{"featureprop_id", "rank"}
	featurepropPrimaryKeyColumns     = []string{"featureprop_id"}
)

type (
	// FeaturepropSlice is an alias for a slice of pointers to Featureprop.
	// This should generally be used opposed to []Featureprop.
	FeaturepropSlice []*Featureprop
	// FeaturepropHook is the signature for custom Featureprop hook methods
	FeaturepropHook func(boil.Executor, *Featureprop) error

	featurepropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featurepropType                 = reflect.TypeOf(&Featureprop{})
	featurepropMapping              = queries.MakeStructMapping(featurepropType)
	featurepropPrimaryKeyMapping, _ = queries.BindMapping(featurepropType, featurepropMapping, featurepropPrimaryKeyColumns)
	featurepropInsertCacheMut       sync.RWMutex
	featurepropInsertCache          = make(map[string]insertCache)
	featurepropUpdateCacheMut       sync.RWMutex
	featurepropUpdateCache          = make(map[string]updateCache)
	featurepropUpsertCacheMut       sync.RWMutex
	featurepropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featurepropBeforeInsertHooks []FeaturepropHook
var featurepropBeforeUpdateHooks []FeaturepropHook
var featurepropBeforeDeleteHooks []FeaturepropHook
var featurepropBeforeUpsertHooks []FeaturepropHook

var featurepropAfterInsertHooks []FeaturepropHook
var featurepropAfterSelectHooks []FeaturepropHook
var featurepropAfterUpdateHooks []FeaturepropHook
var featurepropAfterDeleteHooks []FeaturepropHook
var featurepropAfterUpsertHooks []FeaturepropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Featureprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Featureprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Featureprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Featureprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Featureprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Featureprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Featureprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Featureprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Featureprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurepropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeaturepropHook registers your hook function for all future operations.
func AddFeaturepropHook(hookPoint boil.HookPoint, featurepropHook FeaturepropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featurepropBeforeInsertHooks = append(featurepropBeforeInsertHooks, featurepropHook)
	case boil.BeforeUpdateHook:
		featurepropBeforeUpdateHooks = append(featurepropBeforeUpdateHooks, featurepropHook)
	case boil.BeforeDeleteHook:
		featurepropBeforeDeleteHooks = append(featurepropBeforeDeleteHooks, featurepropHook)
	case boil.BeforeUpsertHook:
		featurepropBeforeUpsertHooks = append(featurepropBeforeUpsertHooks, featurepropHook)
	case boil.AfterInsertHook:
		featurepropAfterInsertHooks = append(featurepropAfterInsertHooks, featurepropHook)
	case boil.AfterSelectHook:
		featurepropAfterSelectHooks = append(featurepropAfterSelectHooks, featurepropHook)
	case boil.AfterUpdateHook:
		featurepropAfterUpdateHooks = append(featurepropAfterUpdateHooks, featurepropHook)
	case boil.AfterDeleteHook:
		featurepropAfterDeleteHooks = append(featurepropAfterDeleteHooks, featurepropHook)
	case boil.AfterUpsertHook:
		featurepropAfterUpsertHooks = append(featurepropAfterUpsertHooks, featurepropHook)
	}
}

// OneP returns a single featureprop record from the query, and panics on error.
func (q featurepropQuery) OneP() *Featureprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureprop record from the query.
func (q featurepropQuery) One() (*Featureprop, error) {
	o := &Featureprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for featureprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Featureprop records from the query, and panics on error.
func (q featurepropQuery) AllP() FeaturepropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Featureprop records from the query.
func (q featurepropQuery) All() (FeaturepropSlice, error) {
	var o FeaturepropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Featureprop slice")
	}

	if len(featurepropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Featureprop records in the query, and panics on error.
func (q featurepropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Featureprop records in the query.
func (q featurepropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count featureprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featurepropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featurepropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if featureprop exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *Featureprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Featureprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// FeatureG pointed to by the foreign key.
func (o *Featureprop) FeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Feature(boil.GetDB(), mods...)
}

// Feature pointed to by the foreign key.
func (o *Featureprop) Feature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// FeaturepropPubG pointed to by the foreign key.
func (o *Featureprop) FeaturepropPubG(mods ...qm.QueryMod) featurepropPubQuery {
	return o.FeaturepropPub(boil.GetDB(), mods...)
}

// FeaturepropPub pointed to by the foreign key.
func (o *Featureprop) FeaturepropPub(exec boil.Executor, mods ...qm.QueryMod) featurepropPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("featureprop_id=$1", o.FeaturepropID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturepropPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"featureprop_pub\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurepropL) LoadType(e boil.Executor, singular bool, maybeFeatureprop interface{}) error {
	var slice []*Featureprop
	var object *Featureprop

	count := 1
	if singular {
		object = maybeFeatureprop.(*Featureprop)
	} else {
		slice = *maybeFeatureprop.(*FeaturepropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurepropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &featurepropR{}
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

	if len(featurepropAfterSelectHooks) != 0 {
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

// LoadFeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurepropL) LoadFeature(e boil.Executor, singular bool, maybeFeatureprop interface{}) error {
	var slice []*Featureprop
	var object *Featureprop

	count := 1
	if singular {
		object = maybeFeatureprop.(*Featureprop)
	} else {
		slice = *maybeFeatureprop.(*FeaturepropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurepropR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featurepropR{}
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

	if len(featurepropAfterSelectHooks) != 0 {
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

// LoadFeaturepropPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurepropL) LoadFeaturepropPub(e boil.Executor, singular bool, maybeFeatureprop interface{}) error {
	var slice []*Featureprop
	var object *Featureprop

	count := 1
	if singular {
		object = maybeFeatureprop.(*Featureprop)
	} else {
		slice = *maybeFeatureprop.(*FeaturepropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurepropR{}
		args[0] = object.FeaturepropID
	} else {
		for i, obj := range slice {
			obj.R = &featurepropR{}
			args[i] = obj.FeaturepropID
		}
	}

	query := fmt.Sprintf(
		"select * from \"featureprop_pub\" where \"featureprop_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeaturepropPub")
	}
	defer results.Close()

	var resultSlice []*FeaturepropPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeaturepropPub")
	}

	if len(featurepropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeaturepropPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeaturepropID == foreign.FeaturepropID {
				local.R.FeaturepropPub = foreign
				break
			}
		}
	}

	return nil
}

// SetType of the featureprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeFeatureprop.
func (o *Featureprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"featureprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurepropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.FeaturepropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &featurepropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeFeatureprop: o,
		}
	} else {
		related.R.TypeFeatureprop = o
	}

	return nil
}

// SetFeature of the featureprop to the related item.
// Sets o.R.Feature to related.
// Adds o to related.R.Featureprop.
func (o *Featureprop) SetFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"featureprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurepropPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeaturepropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureID = related.FeatureID

	if o.R == nil {
		o.R = &featurepropR{
			Feature: related,
		}
	} else {
		o.R.Feature = related
	}

	if related.R == nil {
		related.R = &featureR{
			Featureprop: o,
		}
	} else {
		related.R.Featureprop = o
	}

	return nil
}

// SetFeaturepropPub of the featureprop to the related item.
// Sets o.R.FeaturepropPub to related.
// Adds o to related.R.Featureprop.
func (o *Featureprop) SetFeaturepropPub(exec boil.Executor, insert bool, related *FeaturepropPub) error {
	var err error

	if insert {
		related.FeaturepropID = o.FeaturepropID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"featureprop_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"featureprop_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurepropPubPrimaryKeyColumns),
		)
		values := []interface{}{o.FeaturepropID, related.FeaturepropPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeaturepropID = o.FeaturepropID

	}

	if o.R == nil {
		o.R = &featurepropR{
			FeaturepropPub: related,
		}
	} else {
		o.R.FeaturepropPub = related
	}

	if related.R == nil {
		related.R = &featurepropPubR{
			Featureprop: o,
		}
	} else {
		related.R.Featureprop = o
	}
	return nil
}

// FeaturepropsG retrieves all records.
func FeaturepropsG(mods ...qm.QueryMod) featurepropQuery {
	return Featureprops(boil.GetDB(), mods...)
}

// Featureprops retrieves all the records using an executor.
func Featureprops(exec boil.Executor, mods ...qm.QueryMod) featurepropQuery {
	mods = append(mods, qm.From("\"featureprop\""))
	return featurepropQuery{NewQuery(exec, mods...)}
}

// FindFeaturepropG retrieves a single record by ID.
func FindFeaturepropG(featurepropID int, selectCols ...string) (*Featureprop, error) {
	return FindFeatureprop(boil.GetDB(), featurepropID, selectCols...)
}

// FindFeaturepropGP retrieves a single record by ID, and panics on error.
func FindFeaturepropGP(featurepropID int, selectCols ...string) *Featureprop {
	retobj, err := FindFeatureprop(boil.GetDB(), featurepropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureprop(exec boil.Executor, featurepropID int, selectCols ...string) (*Featureprop, error) {
	featurepropObj := &Featureprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"featureprop\" where \"featureprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featurepropID)

	err := q.Bind(featurepropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from featureprop")
	}

	return featurepropObj, nil
}

// FindFeaturepropP retrieves a single record by ID with an executor, and panics on error.
func FindFeaturepropP(exec boil.Executor, featurepropID int, selectCols ...string) *Featureprop {
	retobj, err := FindFeatureprop(exec, featurepropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Featureprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Featureprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Featureprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Featureprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no featureprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurepropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featurepropInsertCacheMut.RLock()
	cache, cached := featurepropInsertCache[key]
	featurepropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featurepropColumns,
			featurepropColumnsWithDefault,
			featurepropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featurepropType, featurepropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featurepropType, featurepropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"featureprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into featureprop")
	}

	if !cached {
		featurepropInsertCacheMut.Lock()
		featurepropInsertCache[key] = cache
		featurepropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Featureprop record. See Update for
// whitelist behavior description.
func (o *Featureprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Featureprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Featureprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Featureprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Featureprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Featureprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Featureprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featurepropUpdateCacheMut.RLock()
	cache, cached := featurepropUpdateCache[key]
	featurepropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featurepropColumns, featurepropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update featureprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"featureprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featurepropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featurepropType, featurepropMapping, append(wl, featurepropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update featureprop row")
	}

	if !cached {
		featurepropUpdateCacheMut.Lock()
		featurepropUpdateCache[key] = cache
		featurepropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featurepropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featurepropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for featureprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeaturepropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeaturepropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeaturepropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeaturepropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"featureprop\" SET %s WHERE (\"featureprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurepropPrimaryKeyColumns), len(colNames)+1, len(featurepropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in featureprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Featureprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Featureprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Featureprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Featureprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no featureprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurepropColumnsWithDefault, o)

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

	featurepropUpsertCacheMut.RLock()
	cache, cached := featurepropUpsertCache[key]
	featurepropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featurepropColumns,
			featurepropColumnsWithDefault,
			featurepropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featurepropColumns,
			featurepropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert featureprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featurepropPrimaryKeyColumns))
			copy(conflict, featurepropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"featureprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featurepropType, featurepropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featurepropType, featurepropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for featureprop")
	}

	if !cached {
		featurepropUpsertCacheMut.Lock()
		featurepropUpsertCache[key] = cache
		featurepropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Featureprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Featureprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Featureprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Featureprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no Featureprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Featureprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Featureprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Featureprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Featureprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Featureprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featurepropPrimaryKeyMapping)
	sql := "DELETE FROM \"featureprop\" WHERE \"featureprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from featureprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featurepropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featurepropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no featurepropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from featureprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeaturepropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeaturepropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Featureprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeaturepropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeaturepropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Featureprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featurepropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"featureprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurepropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurepropPrimaryKeyColumns), 1, len(featurepropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from featureprop slice")
	}

	if len(featurepropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Featureprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Featureprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Featureprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no Featureprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Featureprop) Reload(exec boil.Executor) error {
	ret, err := FindFeatureprop(exec, o.FeaturepropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturepropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturepropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturepropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FeaturepropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturepropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureprops := FeaturepropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"featureprop\".* FROM \"featureprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurepropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featurepropPrimaryKeyColumns), 1, len(featurepropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeaturepropSlice")
	}

	*o = featureprops

	return nil
}

// FeaturepropExists checks if the Featureprop row exists.
func FeaturepropExists(exec boil.Executor, featurepropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"featureprop\" where \"featureprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featurepropID)
	}

	row := exec.QueryRow(sql, featurepropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if featureprop exists")
	}

	return exists, nil
}

// FeaturepropExistsG checks if the Featureprop row exists.
func FeaturepropExistsG(featurepropID int) (bool, error) {
	return FeaturepropExists(boil.GetDB(), featurepropID)
}

// FeaturepropExistsGP checks if the Featureprop row exists. Panics on error.
func FeaturepropExistsGP(featurepropID int) bool {
	e, err := FeaturepropExists(boil.GetDB(), featurepropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeaturepropExistsP checks if the Featureprop row exists. Panics on error.
func FeaturepropExistsP(exec boil.Executor, featurepropID int) bool {
	e, err := FeaturepropExists(exec, featurepropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
