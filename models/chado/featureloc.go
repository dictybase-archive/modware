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
	"gopkg.in/nullbio/null.v5"
)

// Featureloc is an object representing the database table.
type Featureloc struct {
	FeaturelocID  int         `boil:"featureloc_id" json:"featureloc_id" toml:"featureloc_id" yaml:"featureloc_id"`
	FeatureID     int         `boil:"feature_id" json:"feature_id" toml:"feature_id" yaml:"feature_id"`
	SrcfeatureID  null.Int    `boil:"srcfeature_id" json:"srcfeature_id,omitempty" toml:"srcfeature_id" yaml:"srcfeature_id,omitempty"`
	Fmin          null.Int    `boil:"fmin" json:"fmin,omitempty" toml:"fmin" yaml:"fmin,omitempty"`
	IsFminPartial bool        `boil:"is_fmin_partial" json:"is_fmin_partial" toml:"is_fmin_partial" yaml:"is_fmin_partial"`
	Fmax          null.Int    `boil:"fmax" json:"fmax,omitempty" toml:"fmax" yaml:"fmax,omitempty"`
	IsFmaxPartial bool        `boil:"is_fmax_partial" json:"is_fmax_partial" toml:"is_fmax_partial" yaml:"is_fmax_partial"`
	Strand        null.Int16  `boil:"strand" json:"strand,omitempty" toml:"strand" yaml:"strand,omitempty"`
	Phase         null.Int    `boil:"phase" json:"phase,omitempty" toml:"phase" yaml:"phase,omitempty"`
	ResidueInfo   null.String `boil:"residue_info" json:"residue_info,omitempty" toml:"residue_info" yaml:"residue_info,omitempty"`
	Locgroup      int         `boil:"locgroup" json:"locgroup" toml:"locgroup" yaml:"locgroup"`
	Rank          int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *featurelocR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featurelocL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featurelocR is where relationships are stored.
type featurelocR struct {
	Feature       *Feature
	Srcfeature    *Feature
	FeaturelocPub *FeaturelocPub
}

// featurelocL is where Load methods for each relationship are stored.
type featurelocL struct{}

var (
	featurelocColumns               = []string{"featureloc_id", "feature_id", "srcfeature_id", "fmin", "is_fmin_partial", "fmax", "is_fmax_partial", "strand", "phase", "residue_info", "locgroup", "rank"}
	featurelocColumnsWithoutDefault = []string{"feature_id", "srcfeature_id", "fmin", "fmax", "strand", "phase", "residue_info"}
	featurelocColumnsWithDefault    = []string{"featureloc_id", "is_fmin_partial", "is_fmax_partial", "locgroup", "rank"}
	featurelocPrimaryKeyColumns     = []string{"featureloc_id"}
)

type (
	// FeaturelocSlice is an alias for a slice of pointers to Featureloc.
	// This should generally be used opposed to []Featureloc.
	FeaturelocSlice []*Featureloc
	// FeaturelocHook is the signature for custom Featureloc hook methods
	FeaturelocHook func(boil.Executor, *Featureloc) error

	featurelocQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featurelocType                 = reflect.TypeOf(&Featureloc{})
	featurelocMapping              = queries.MakeStructMapping(featurelocType)
	featurelocPrimaryKeyMapping, _ = queries.BindMapping(featurelocType, featurelocMapping, featurelocPrimaryKeyColumns)
	featurelocInsertCacheMut       sync.RWMutex
	featurelocInsertCache          = make(map[string]insertCache)
	featurelocUpdateCacheMut       sync.RWMutex
	featurelocUpdateCache          = make(map[string]updateCache)
	featurelocUpsertCacheMut       sync.RWMutex
	featurelocUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featurelocBeforeInsertHooks []FeaturelocHook
var featurelocBeforeUpdateHooks []FeaturelocHook
var featurelocBeforeDeleteHooks []FeaturelocHook
var featurelocBeforeUpsertHooks []FeaturelocHook

var featurelocAfterInsertHooks []FeaturelocHook
var featurelocAfterSelectHooks []FeaturelocHook
var featurelocAfterUpdateHooks []FeaturelocHook
var featurelocAfterDeleteHooks []FeaturelocHook
var featurelocAfterUpsertHooks []FeaturelocHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Featureloc) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Featureloc) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Featureloc) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Featureloc) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Featureloc) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Featureloc) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Featureloc) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Featureloc) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Featureloc) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurelocAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeaturelocHook registers your hook function for all future operations.
func AddFeaturelocHook(hookPoint boil.HookPoint, featurelocHook FeaturelocHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featurelocBeforeInsertHooks = append(featurelocBeforeInsertHooks, featurelocHook)
	case boil.BeforeUpdateHook:
		featurelocBeforeUpdateHooks = append(featurelocBeforeUpdateHooks, featurelocHook)
	case boil.BeforeDeleteHook:
		featurelocBeforeDeleteHooks = append(featurelocBeforeDeleteHooks, featurelocHook)
	case boil.BeforeUpsertHook:
		featurelocBeforeUpsertHooks = append(featurelocBeforeUpsertHooks, featurelocHook)
	case boil.AfterInsertHook:
		featurelocAfterInsertHooks = append(featurelocAfterInsertHooks, featurelocHook)
	case boil.AfterSelectHook:
		featurelocAfterSelectHooks = append(featurelocAfterSelectHooks, featurelocHook)
	case boil.AfterUpdateHook:
		featurelocAfterUpdateHooks = append(featurelocAfterUpdateHooks, featurelocHook)
	case boil.AfterDeleteHook:
		featurelocAfterDeleteHooks = append(featurelocAfterDeleteHooks, featurelocHook)
	case boil.AfterUpsertHook:
		featurelocAfterUpsertHooks = append(featurelocAfterUpsertHooks, featurelocHook)
	}
}

// OneP returns a single featureloc record from the query, and panics on error.
func (q featurelocQuery) OneP() *Featureloc {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureloc record from the query.
func (q featurelocQuery) One() (*Featureloc, error) {
	o := &Featureloc{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for featureloc")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Featureloc records from the query, and panics on error.
func (q featurelocQuery) AllP() FeaturelocSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Featureloc records from the query.
func (q featurelocQuery) All() (FeaturelocSlice, error) {
	var o FeaturelocSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Featureloc slice")
	}

	if len(featurelocAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Featureloc records in the query, and panics on error.
func (q featurelocQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Featureloc records in the query.
func (q featurelocQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count featureloc rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featurelocQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featurelocQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if featureloc exists")
	}

	return count > 0, nil
}

// FeatureG pointed to by the foreign key.
func (o *Featureloc) FeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Feature(boil.GetDB(), mods...)
}

// Feature pointed to by the foreign key.
func (o *Featureloc) Feature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// SrcfeatureG pointed to by the foreign key.
func (o *Featureloc) SrcfeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Srcfeature(boil.GetDB(), mods...)
}

// Srcfeature pointed to by the foreign key.
func (o *Featureloc) Srcfeature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.SrcfeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// FeaturelocPubG pointed to by the foreign key.
func (o *Featureloc) FeaturelocPubG(mods ...qm.QueryMod) featurelocPubQuery {
	return o.FeaturelocPub(boil.GetDB(), mods...)
}

// FeaturelocPub pointed to by the foreign key.
func (o *Featureloc) FeaturelocPub(exec boil.Executor, mods ...qm.QueryMod) featurelocPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("featureloc_id=$1", o.FeaturelocID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturelocPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"featureloc_pub\"")

	return query
}

// LoadFeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurelocL) LoadFeature(e boil.Executor, singular bool, maybeFeatureloc interface{}) error {
	var slice []*Featureloc
	var object *Featureloc

	count := 1
	if singular {
		object = maybeFeatureloc.(*Featureloc)
	} else {
		slice = *maybeFeatureloc.(*FeaturelocSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurelocR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featurelocR{}
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

	if len(featurelocAfterSelectHooks) != 0 {
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

// LoadSrcfeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurelocL) LoadSrcfeature(e boil.Executor, singular bool, maybeFeatureloc interface{}) error {
	var slice []*Featureloc
	var object *Featureloc

	count := 1
	if singular {
		object = maybeFeatureloc.(*Featureloc)
	} else {
		slice = *maybeFeatureloc.(*FeaturelocSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurelocR{}
		args[0] = object.SrcfeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featurelocR{}
			args[i] = obj.SrcfeatureID
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

	if len(featurelocAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Srcfeature = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.SrcfeatureID.Int == foreign.FeatureID {
				local.R.Srcfeature = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeaturelocPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurelocL) LoadFeaturelocPub(e boil.Executor, singular bool, maybeFeatureloc interface{}) error {
	var slice []*Featureloc
	var object *Featureloc

	count := 1
	if singular {
		object = maybeFeatureloc.(*Featureloc)
	} else {
		slice = *maybeFeatureloc.(*FeaturelocSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurelocR{}
		args[0] = object.FeaturelocID
	} else {
		for i, obj := range slice {
			obj.R = &featurelocR{}
			args[i] = obj.FeaturelocID
		}
	}

	query := fmt.Sprintf(
		"select * from \"featureloc_pub\" where \"featureloc_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeaturelocPub")
	}
	defer results.Close()

	var resultSlice []*FeaturelocPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeaturelocPub")
	}

	if len(featurelocAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeaturelocPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeaturelocID == foreign.FeaturelocID {
				local.R.FeaturelocPub = foreign
				break
			}
		}
	}

	return nil
}

// SetFeature of the featureloc to the related item.
// Sets o.R.Feature to related.
// Adds o to related.R.Featureloc.
func (o *Featureloc) SetFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"featureloc\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurelocPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeaturelocID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureID = related.FeatureID

	if o.R == nil {
		o.R = &featurelocR{
			Feature: related,
		}
	} else {
		o.R.Feature = related
	}

	if related.R == nil {
		related.R = &featureR{
			Featureloc: o,
		}
	} else {
		related.R.Featureloc = o
	}

	return nil
}

// SetSrcfeature of the featureloc to the related item.
// Sets o.R.Srcfeature to related.
// Adds o to related.R.SrcfeatureFeaturelocs.
func (o *Featureloc) SetSrcfeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"featureloc\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"srcfeature_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurelocPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeaturelocID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SrcfeatureID.Int = related.FeatureID
	o.SrcfeatureID.Valid = true

	if o.R == nil {
		o.R = &featurelocR{
			Srcfeature: related,
		}
	} else {
		o.R.Srcfeature = related
	}

	if related.R == nil {
		related.R = &featureR{
			SrcfeatureFeaturelocs: FeaturelocSlice{o},
		}
	} else {
		related.R.SrcfeatureFeaturelocs = append(related.R.SrcfeatureFeaturelocs, o)
	}

	return nil
}

// RemoveSrcfeature relationship.
// Sets o.R.Srcfeature to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Featureloc) RemoveSrcfeature(exec boil.Executor, related *Feature) error {
	var err error

	o.SrcfeatureID.Valid = false
	if err = o.Update(exec, "srcfeature_id"); err != nil {
		o.SrcfeatureID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Srcfeature = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.SrcfeatureFeaturelocs {
		if o.SrcfeatureID.Int != ri.SrcfeatureID.Int {
			continue
		}

		ln := len(related.R.SrcfeatureFeaturelocs)
		if ln > 1 && i < ln-1 {
			related.R.SrcfeatureFeaturelocs[i] = related.R.SrcfeatureFeaturelocs[ln-1]
		}
		related.R.SrcfeatureFeaturelocs = related.R.SrcfeatureFeaturelocs[:ln-1]
		break
	}
	return nil
}

// SetFeaturelocPub of the featureloc to the related item.
// Sets o.R.FeaturelocPub to related.
// Adds o to related.R.Featureloc.
func (o *Featureloc) SetFeaturelocPub(exec boil.Executor, insert bool, related *FeaturelocPub) error {
	var err error

	if insert {
		related.FeaturelocID = o.FeaturelocID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"featureloc_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"featureloc_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurelocPubPrimaryKeyColumns),
		)
		values := []interface{}{o.FeaturelocID, related.FeaturelocPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeaturelocID = o.FeaturelocID

	}

	if o.R == nil {
		o.R = &featurelocR{
			FeaturelocPub: related,
		}
	} else {
		o.R.FeaturelocPub = related
	}

	if related.R == nil {
		related.R = &featurelocPubR{
			Featureloc: o,
		}
	} else {
		related.R.Featureloc = o
	}
	return nil
}

// FeaturelocsG retrieves all records.
func FeaturelocsG(mods ...qm.QueryMod) featurelocQuery {
	return Featurelocs(boil.GetDB(), mods...)
}

// Featurelocs retrieves all the records using an executor.
func Featurelocs(exec boil.Executor, mods ...qm.QueryMod) featurelocQuery {
	mods = append(mods, qm.From("\"featureloc\""))
	return featurelocQuery{NewQuery(exec, mods...)}
}

// FindFeaturelocG retrieves a single record by ID.
func FindFeaturelocG(featurelocID int, selectCols ...string) (*Featureloc, error) {
	return FindFeatureloc(boil.GetDB(), featurelocID, selectCols...)
}

// FindFeaturelocGP retrieves a single record by ID, and panics on error.
func FindFeaturelocGP(featurelocID int, selectCols ...string) *Featureloc {
	retobj, err := FindFeatureloc(boil.GetDB(), featurelocID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureloc retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureloc(exec boil.Executor, featurelocID int, selectCols ...string) (*Featureloc, error) {
	featurelocObj := &Featureloc{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"featureloc\" where \"featureloc_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featurelocID)

	err := q.Bind(featurelocObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from featureloc")
	}

	return featurelocObj, nil
}

// FindFeaturelocP retrieves a single record by ID with an executor, and panics on error.
func FindFeaturelocP(exec boil.Executor, featurelocID int, selectCols ...string) *Featureloc {
	retobj, err := FindFeatureloc(exec, featurelocID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Featureloc) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Featureloc) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Featureloc) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Featureloc) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no featureloc provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurelocColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featurelocInsertCacheMut.RLock()
	cache, cached := featurelocInsertCache[key]
	featurelocInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featurelocColumns,
			featurelocColumnsWithDefault,
			featurelocColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featurelocType, featurelocMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featurelocType, featurelocMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"featureloc\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into featureloc")
	}

	if !cached {
		featurelocInsertCacheMut.Lock()
		featurelocInsertCache[key] = cache
		featurelocInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Featureloc record. See Update for
// whitelist behavior description.
func (o *Featureloc) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Featureloc record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Featureloc) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Featureloc, and panics on error.
// See Update for whitelist behavior description.
func (o *Featureloc) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Featureloc.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Featureloc) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featurelocUpdateCacheMut.RLock()
	cache, cached := featurelocUpdateCache[key]
	featurelocUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featurelocColumns, featurelocPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update featureloc, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"featureloc\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featurelocPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featurelocType, featurelocMapping, append(wl, featurelocPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update featureloc row")
	}

	if !cached {
		featurelocUpdateCacheMut.Lock()
		featurelocUpdateCache[key] = cache
		featurelocUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featurelocQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featurelocQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for featureloc")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeaturelocSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeaturelocSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeaturelocSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeaturelocSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurelocPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"featureloc\" SET %s WHERE (\"featureloc_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurelocPrimaryKeyColumns), len(colNames)+1, len(featurelocPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in featureloc slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Featureloc) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Featureloc) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Featureloc) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Featureloc) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no featureloc provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurelocColumnsWithDefault, o)

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

	featurelocUpsertCacheMut.RLock()
	cache, cached := featurelocUpsertCache[key]
	featurelocUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featurelocColumns,
			featurelocColumnsWithDefault,
			featurelocColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featurelocColumns,
			featurelocPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert featureloc, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featurelocPrimaryKeyColumns))
			copy(conflict, featurelocPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"featureloc\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featurelocType, featurelocMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featurelocType, featurelocMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for featureloc")
	}

	if !cached {
		featurelocUpsertCacheMut.Lock()
		featurelocUpsertCache[key] = cache
		featurelocUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Featureloc record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Featureloc) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Featureloc record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Featureloc) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Featureloc provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Featureloc record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Featureloc) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Featureloc record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Featureloc) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Featureloc provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featurelocPrimaryKeyMapping)
	sql := "DELETE FROM \"featureloc\" WHERE \"featureloc_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from featureloc")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featurelocQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featurelocQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no featurelocQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featureloc")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeaturelocSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeaturelocSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Featureloc slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeaturelocSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeaturelocSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Featureloc slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featurelocBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurelocPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"featureloc\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurelocPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurelocPrimaryKeyColumns), 1, len(featurelocPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featureloc slice")
	}

	if len(featurelocAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Featureloc) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Featureloc) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Featureloc) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Featureloc provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Featureloc) Reload(exec boil.Executor) error {
	ret, err := FindFeatureloc(exec, o.FeaturelocID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturelocSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturelocSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturelocSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty FeaturelocSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturelocSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featurelocs := FeaturelocSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurelocPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"featureloc\".* FROM \"featureloc\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurelocPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featurelocPrimaryKeyColumns), 1, len(featurelocPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featurelocs)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in FeaturelocSlice")
	}

	*o = featurelocs

	return nil
}

// FeaturelocExists checks if the Featureloc row exists.
func FeaturelocExists(exec boil.Executor, featurelocID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"featureloc\" where \"featureloc_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featurelocID)
	}

	row := exec.QueryRow(sql, featurelocID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if featureloc exists")
	}

	return exists, nil
}

// FeaturelocExistsG checks if the Featureloc row exists.
func FeaturelocExistsG(featurelocID int) (bool, error) {
	return FeaturelocExists(boil.GetDB(), featurelocID)
}

// FeaturelocExistsGP checks if the Featureloc row exists. Panics on error.
func FeaturelocExistsGP(featurelocID int) bool {
	e, err := FeaturelocExists(boil.GetDB(), featurelocID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeaturelocExistsP checks if the Featureloc row exists. Panics on error.
func FeaturelocExistsP(exec boil.Executor, featurelocID int) bool {
	e, err := FeaturelocExists(exec, featurelocID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
