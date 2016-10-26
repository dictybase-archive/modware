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

// FeatureCvterm is an object representing the database table.
type FeatureCvterm struct {
	FeatureCvtermID int  `boil:"feature_cvterm_id" json:"feature_cvterm_id" toml:"feature_cvterm_id" yaml:"feature_cvterm_id"`
	FeatureID       int  `boil:"feature_id" json:"feature_id" toml:"feature_id" yaml:"feature_id"`
	CvtermID        int  `boil:"cvterm_id" json:"cvterm_id" toml:"cvterm_id" yaml:"cvterm_id"`
	PubID           int  `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`
	IsNot           bool `boil:"is_not" json:"is_not" toml:"is_not" yaml:"is_not"`
	Rank            int  `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *featureCvtermR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureCvtermL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureCvtermR is where relationships are stored.
type featureCvtermR struct {
	Cvterm              *Cvterm
	Pub                 *Pub
	Feature             *Feature
	FeatureCvtermPub    *FeatureCvtermPub
	FeatureCvtermDbxref *FeatureCvtermDbxref
	FeatureCvtermprop   *FeatureCvtermprop
}

// featureCvtermL is where Load methods for each relationship are stored.
type featureCvtermL struct{}

var (
	featureCvtermColumns               = []string{"feature_cvterm_id", "feature_id", "cvterm_id", "pub_id", "is_not", "rank"}
	featureCvtermColumnsWithoutDefault = []string{"feature_id", "cvterm_id", "pub_id"}
	featureCvtermColumnsWithDefault    = []string{"feature_cvterm_id", "is_not", "rank"}
	featureCvtermPrimaryKeyColumns     = []string{"feature_cvterm_id"}
)

type (
	// FeatureCvtermSlice is an alias for a slice of pointers to FeatureCvterm.
	// This should generally be used opposed to []FeatureCvterm.
	FeatureCvtermSlice []*FeatureCvterm
	// FeatureCvtermHook is the signature for custom FeatureCvterm hook methods
	FeatureCvtermHook func(boil.Executor, *FeatureCvterm) error

	featureCvtermQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureCvtermType                 = reflect.TypeOf(&FeatureCvterm{})
	featureCvtermMapping              = queries.MakeStructMapping(featureCvtermType)
	featureCvtermPrimaryKeyMapping, _ = queries.BindMapping(featureCvtermType, featureCvtermMapping, featureCvtermPrimaryKeyColumns)
	featureCvtermInsertCacheMut       sync.RWMutex
	featureCvtermInsertCache          = make(map[string]insertCache)
	featureCvtermUpdateCacheMut       sync.RWMutex
	featureCvtermUpdateCache          = make(map[string]updateCache)
	featureCvtermUpsertCacheMut       sync.RWMutex
	featureCvtermUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureCvtermBeforeInsertHooks []FeatureCvtermHook
var featureCvtermBeforeUpdateHooks []FeatureCvtermHook
var featureCvtermBeforeDeleteHooks []FeatureCvtermHook
var featureCvtermBeforeUpsertHooks []FeatureCvtermHook

var featureCvtermAfterInsertHooks []FeatureCvtermHook
var featureCvtermAfterSelectHooks []FeatureCvtermHook
var featureCvtermAfterUpdateHooks []FeatureCvtermHook
var featureCvtermAfterDeleteHooks []FeatureCvtermHook
var featureCvtermAfterUpsertHooks []FeatureCvtermHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeatureCvterm) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeatureCvterm) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeatureCvterm) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeatureCvterm) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeatureCvterm) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeatureCvterm) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeatureCvterm) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeatureCvterm) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeatureCvterm) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureCvtermAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureCvtermHook registers your hook function for all future operations.
func AddFeatureCvtermHook(hookPoint boil.HookPoint, featureCvtermHook FeatureCvtermHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureCvtermBeforeInsertHooks = append(featureCvtermBeforeInsertHooks, featureCvtermHook)
	case boil.BeforeUpdateHook:
		featureCvtermBeforeUpdateHooks = append(featureCvtermBeforeUpdateHooks, featureCvtermHook)
	case boil.BeforeDeleteHook:
		featureCvtermBeforeDeleteHooks = append(featureCvtermBeforeDeleteHooks, featureCvtermHook)
	case boil.BeforeUpsertHook:
		featureCvtermBeforeUpsertHooks = append(featureCvtermBeforeUpsertHooks, featureCvtermHook)
	case boil.AfterInsertHook:
		featureCvtermAfterInsertHooks = append(featureCvtermAfterInsertHooks, featureCvtermHook)
	case boil.AfterSelectHook:
		featureCvtermAfterSelectHooks = append(featureCvtermAfterSelectHooks, featureCvtermHook)
	case boil.AfterUpdateHook:
		featureCvtermAfterUpdateHooks = append(featureCvtermAfterUpdateHooks, featureCvtermHook)
	case boil.AfterDeleteHook:
		featureCvtermAfterDeleteHooks = append(featureCvtermAfterDeleteHooks, featureCvtermHook)
	case boil.AfterUpsertHook:
		featureCvtermAfterUpsertHooks = append(featureCvtermAfterUpsertHooks, featureCvtermHook)
	}
}

// OneP returns a single featureCvterm record from the query, and panics on error.
func (q featureCvtermQuery) OneP() *FeatureCvterm {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureCvterm record from the query.
func (q featureCvtermQuery) One() (*FeatureCvterm, error) {
	o := &FeatureCvterm{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for feature_cvterm")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeatureCvterm records from the query, and panics on error.
func (q featureCvtermQuery) AllP() FeatureCvtermSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeatureCvterm records from the query.
func (q featureCvtermQuery) All() (FeatureCvtermSlice, error) {
	var o FeatureCvtermSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to FeatureCvterm slice")
	}

	if len(featureCvtermAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeatureCvterm records in the query, and panics on error.
func (q featureCvtermQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeatureCvterm records in the query.
func (q featureCvtermQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count feature_cvterm rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureCvtermQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureCvtermQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if feature_cvterm exists")
	}

	return count > 0, nil
}

// CvtermG pointed to by the foreign key.
func (o *FeatureCvterm) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *FeatureCvterm) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// PubG pointed to by the foreign key.
func (o *FeatureCvterm) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *FeatureCvterm) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// FeatureG pointed to by the foreign key.
func (o *FeatureCvterm) FeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Feature(boil.GetDB(), mods...)
}

// Feature pointed to by the foreign key.
func (o *FeatureCvterm) Feature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// FeatureCvtermPubG pointed to by the foreign key.
func (o *FeatureCvterm) FeatureCvtermPubG(mods ...qm.QueryMod) featureCvtermPubQuery {
	return o.FeatureCvtermPub(boil.GetDB(), mods...)
}

// FeatureCvtermPub pointed to by the foreign key.
func (o *FeatureCvterm) FeatureCvtermPub(exec boil.Executor, mods ...qm.QueryMod) featureCvtermPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_cvterm_id=$1", o.FeatureCvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvtermPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvterm_pub\"")

	return query
}

// FeatureCvtermDbxrefG pointed to by the foreign key.
func (o *FeatureCvterm) FeatureCvtermDbxrefG(mods ...qm.QueryMod) featureCvtermDbxrefQuery {
	return o.FeatureCvtermDbxref(boil.GetDB(), mods...)
}

// FeatureCvtermDbxref pointed to by the foreign key.
func (o *FeatureCvterm) FeatureCvtermDbxref(exec boil.Executor, mods ...qm.QueryMod) featureCvtermDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_cvterm_id=$1", o.FeatureCvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvtermDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvterm_dbxref\"")

	return query
}

// FeatureCvtermpropG pointed to by the foreign key.
func (o *FeatureCvterm) FeatureCvtermpropG(mods ...qm.QueryMod) featureCvtermpropQuery {
	return o.FeatureCvtermprop(boil.GetDB(), mods...)
}

// FeatureCvtermprop pointed to by the foreign key.
func (o *FeatureCvterm) FeatureCvtermprop(exec boil.Executor, mods ...qm.QueryMod) featureCvtermpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_cvterm_id=$1", o.FeatureCvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvtermprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvtermprop\"")

	return query
}

// LoadCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermL) LoadCvterm(e boil.Executor, singular bool, maybeFeatureCvterm interface{}) error {
	var slice []*FeatureCvterm
	var object *FeatureCvterm

	count := 1
	if singular {
		object = maybeFeatureCvterm.(*FeatureCvterm)
	} else {
		slice = *maybeFeatureCvterm.(*FeatureCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermR{}
			args[i] = obj.CvtermID
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

	if len(featureCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Cvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvtermID {
				local.R.Cvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermL) LoadPub(e boil.Executor, singular bool, maybeFeatureCvterm interface{}) error {
	var slice []*FeatureCvterm
	var object *FeatureCvterm

	count := 1
	if singular {
		object = maybeFeatureCvterm.(*FeatureCvterm)
	} else {
		slice = *maybeFeatureCvterm.(*FeatureCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermR{}
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

	if len(featureCvtermAfterSelectHooks) != 0 {
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
func (featureCvtermL) LoadFeature(e boil.Executor, singular bool, maybeFeatureCvterm interface{}) error {
	var slice []*FeatureCvterm
	var object *FeatureCvterm

	count := 1
	if singular {
		object = maybeFeatureCvterm.(*FeatureCvterm)
	} else {
		slice = *maybeFeatureCvterm.(*FeatureCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermR{}
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

	if len(featureCvtermAfterSelectHooks) != 0 {
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

// LoadFeatureCvtermPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermL) LoadFeatureCvtermPub(e boil.Executor, singular bool, maybeFeatureCvterm interface{}) error {
	var slice []*FeatureCvterm
	var object *FeatureCvterm

	count := 1
	if singular {
		object = maybeFeatureCvterm.(*FeatureCvterm)
	} else {
		slice = *maybeFeatureCvterm.(*FeatureCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermR{}
		args[0] = object.FeatureCvtermID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermR{}
			args[i] = obj.FeatureCvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_cvterm_pub\" where \"feature_cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureCvtermPub")
	}
	defer results.Close()

	var resultSlice []*FeatureCvtermPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureCvtermPub")
	}

	if len(featureCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureCvtermPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureCvtermID == foreign.FeatureCvtermID {
				local.R.FeatureCvtermPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureCvtermDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermL) LoadFeatureCvtermDbxref(e boil.Executor, singular bool, maybeFeatureCvterm interface{}) error {
	var slice []*FeatureCvterm
	var object *FeatureCvterm

	count := 1
	if singular {
		object = maybeFeatureCvterm.(*FeatureCvterm)
	} else {
		slice = *maybeFeatureCvterm.(*FeatureCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermR{}
		args[0] = object.FeatureCvtermID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermR{}
			args[i] = obj.FeatureCvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_cvterm_dbxref\" where \"feature_cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureCvtermDbxref")
	}
	defer results.Close()

	var resultSlice []*FeatureCvtermDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureCvtermDbxref")
	}

	if len(featureCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureCvtermDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureCvtermID == foreign.FeatureCvtermID {
				local.R.FeatureCvtermDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureCvtermprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureCvtermL) LoadFeatureCvtermprop(e boil.Executor, singular bool, maybeFeatureCvterm interface{}) error {
	var slice []*FeatureCvterm
	var object *FeatureCvterm

	count := 1
	if singular {
		object = maybeFeatureCvterm.(*FeatureCvterm)
	} else {
		slice = *maybeFeatureCvterm.(*FeatureCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureCvtermR{}
		args[0] = object.FeatureCvtermID
	} else {
		for i, obj := range slice {
			obj.R = &featureCvtermR{}
			args[i] = obj.FeatureCvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_cvtermprop\" where \"feature_cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureCvtermprop")
	}
	defer results.Close()

	var resultSlice []*FeatureCvtermprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureCvtermprop")
	}

	if len(featureCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureCvtermprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureCvtermID == foreign.FeatureCvtermID {
				local.R.FeatureCvtermprop = foreign
				break
			}
		}
	}

	return nil
}

// SetCvterm of the feature_cvterm to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.FeatureCvterm.
func (o *FeatureCvterm) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.FeatureCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CvtermID = related.CvtermID

	if o.R == nil {
		o.R = &featureCvtermR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			FeatureCvterm: o,
		}
	} else {
		related.R.FeatureCvterm = o
	}

	return nil
}

// SetPub of the feature_cvterm to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.FeatureCvterm.
func (o *FeatureCvterm) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.FeatureCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &featureCvtermR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			FeatureCvterm: o,
		}
	} else {
		related.R.FeatureCvterm = o
	}

	return nil
}

// SetFeature of the feature_cvterm to the related item.
// Sets o.R.Feature to related.
// Adds o to related.R.FeatureCvterm.
func (o *FeatureCvterm) SetFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeatureCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureID = related.FeatureID

	if o.R == nil {
		o.R = &featureCvtermR{
			Feature: related,
		}
	} else {
		o.R.Feature = related
	}

	if related.R == nil {
		related.R = &featureR{
			FeatureCvterm: o,
		}
	} else {
		related.R.FeatureCvterm = o
	}

	return nil
}

// SetFeatureCvtermPub of the feature_cvterm to the related item.
// Sets o.R.FeatureCvtermPub to related.
// Adds o to related.R.FeatureCvterm.
func (o *FeatureCvterm) SetFeatureCvtermPub(exec boil.Executor, insert bool, related *FeatureCvtermPub) error {
	var err error

	if insert {
		related.FeatureCvtermID = o.FeatureCvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_cvterm_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureCvtermPubPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureCvtermID, related.FeatureCvtermPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureCvtermID = o.FeatureCvtermID

	}

	if o.R == nil {
		o.R = &featureCvtermR{
			FeatureCvtermPub: related,
		}
	} else {
		o.R.FeatureCvtermPub = related
	}

	if related.R == nil {
		related.R = &featureCvtermPubR{
			FeatureCvterm: o,
		}
	} else {
		related.R.FeatureCvterm = o
	}
	return nil
}

// SetFeatureCvtermDbxref of the feature_cvterm to the related item.
// Sets o.R.FeatureCvtermDbxref to related.
// Adds o to related.R.FeatureCvterm.
func (o *FeatureCvterm) SetFeatureCvtermDbxref(exec boil.Executor, insert bool, related *FeatureCvtermDbxref) error {
	var err error

	if insert {
		related.FeatureCvtermID = o.FeatureCvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_cvterm_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureCvtermDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureCvtermID, related.FeatureCvtermDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureCvtermID = o.FeatureCvtermID

	}

	if o.R == nil {
		o.R = &featureCvtermR{
			FeatureCvtermDbxref: related,
		}
	} else {
		o.R.FeatureCvtermDbxref = related
	}

	if related.R == nil {
		related.R = &featureCvtermDbxrefR{
			FeatureCvterm: o,
		}
	} else {
		related.R.FeatureCvterm = o
	}
	return nil
}

// SetFeatureCvtermprop of the feature_cvterm to the related item.
// Sets o.R.FeatureCvtermprop to related.
// Adds o to related.R.FeatureCvterm.
func (o *FeatureCvterm) SetFeatureCvtermprop(exec boil.Executor, insert bool, related *FeatureCvtermprop) error {
	var err error

	if insert {
		related.FeatureCvtermID = o.FeatureCvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_cvtermprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureCvtermpropPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureCvtermID, related.FeatureCvtermpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureCvtermID = o.FeatureCvtermID

	}

	if o.R == nil {
		o.R = &featureCvtermR{
			FeatureCvtermprop: related,
		}
	} else {
		o.R.FeatureCvtermprop = related
	}

	if related.R == nil {
		related.R = &featureCvtermpropR{
			FeatureCvterm: o,
		}
	} else {
		related.R.FeatureCvterm = o
	}
	return nil
}

// FeatureCvtermsG retrieves all records.
func FeatureCvtermsG(mods ...qm.QueryMod) featureCvtermQuery {
	return FeatureCvterms(boil.GetDB(), mods...)
}

// FeatureCvterms retrieves all the records using an executor.
func FeatureCvterms(exec boil.Executor, mods ...qm.QueryMod) featureCvtermQuery {
	mods = append(mods, qm.From("\"feature_cvterm\""))
	return featureCvtermQuery{NewQuery(exec, mods...)}
}

// FindFeatureCvtermG retrieves a single record by ID.
func FindFeatureCvtermG(featureCvtermID int, selectCols ...string) (*FeatureCvterm, error) {
	return FindFeatureCvterm(boil.GetDB(), featureCvtermID, selectCols...)
}

// FindFeatureCvtermGP retrieves a single record by ID, and panics on error.
func FindFeatureCvtermGP(featureCvtermID int, selectCols ...string) *FeatureCvterm {
	retobj, err := FindFeatureCvterm(boil.GetDB(), featureCvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureCvterm retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureCvterm(exec boil.Executor, featureCvtermID int, selectCols ...string) (*FeatureCvterm, error) {
	featureCvtermObj := &FeatureCvterm{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_cvterm\" where \"feature_cvterm_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureCvtermID)

	err := q.Bind(featureCvtermObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from feature_cvterm")
	}

	return featureCvtermObj, nil
}

// FindFeatureCvtermP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureCvtermP(exec boil.Executor, featureCvtermID int, selectCols ...string) *FeatureCvterm {
	retobj, err := FindFeatureCvterm(exec, featureCvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeatureCvterm) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeatureCvterm) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeatureCvterm) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeatureCvterm) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_cvterm provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureCvtermColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureCvtermInsertCacheMut.RLock()
	cache, cached := featureCvtermInsertCache[key]
	featureCvtermInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureCvtermColumns,
			featureCvtermColumnsWithDefault,
			featureCvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureCvtermType, featureCvtermMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureCvtermType, featureCvtermMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_cvterm\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into feature_cvterm")
	}

	if !cached {
		featureCvtermInsertCacheMut.Lock()
		featureCvtermInsertCache[key] = cache
		featureCvtermInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeatureCvterm record. See Update for
// whitelist behavior description.
func (o *FeatureCvterm) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeatureCvterm record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeatureCvterm) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeatureCvterm, and panics on error.
// See Update for whitelist behavior description.
func (o *FeatureCvterm) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeatureCvterm.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeatureCvterm) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureCvtermUpdateCacheMut.RLock()
	cache, cached := featureCvtermUpdateCache[key]
	featureCvtermUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureCvtermColumns, featureCvtermPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update feature_cvterm, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featureCvtermPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureCvtermType, featureCvtermMapping, append(wl, featureCvtermPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update feature_cvterm row")
	}

	if !cached {
		featureCvtermUpdateCacheMut.Lock()
		featureCvtermUpdateCache[key] = cache
		featureCvtermUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureCvtermQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureCvtermQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for feature_cvterm")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureCvtermSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureCvtermSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureCvtermSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureCvtermSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_cvterm\" SET %s WHERE (\"feature_cvterm_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureCvtermPrimaryKeyColumns), len(colNames)+1, len(featureCvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in featureCvterm slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeatureCvterm) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeatureCvterm) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeatureCvterm) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeatureCvterm) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_cvterm provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureCvtermColumnsWithDefault, o)

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

	featureCvtermUpsertCacheMut.RLock()
	cache, cached := featureCvtermUpsertCache[key]
	featureCvtermUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureCvtermColumns,
			featureCvtermColumnsWithDefault,
			featureCvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureCvtermColumns,
			featureCvtermPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert feature_cvterm, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featureCvtermPrimaryKeyColumns))
			copy(conflict, featureCvtermPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_cvterm\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureCvtermType, featureCvtermMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureCvtermType, featureCvtermMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for feature_cvterm")
	}

	if !cached {
		featureCvtermUpsertCacheMut.Lock()
		featureCvtermUpsertCache[key] = cache
		featureCvtermUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeatureCvterm record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureCvterm) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeatureCvterm record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeatureCvterm) DeleteG() error {
	if o == nil {
		return errors.New("chado: no FeatureCvterm provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeatureCvterm record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureCvterm) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeatureCvterm record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeatureCvterm) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeatureCvterm provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureCvtermPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_cvterm\" WHERE \"feature_cvterm_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from feature_cvterm")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureCvtermQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureCvtermQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no featureCvtermQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from feature_cvterm")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureCvtermSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureCvtermSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no FeatureCvterm slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureCvtermSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureCvtermSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeatureCvterm slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureCvtermBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureCvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureCvtermPrimaryKeyColumns), 1, len(featureCvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featureCvterm slice")
	}

	if len(featureCvtermAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeatureCvterm) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeatureCvterm) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeatureCvterm) ReloadG() error {
	if o == nil {
		return errors.New("chado: no FeatureCvterm provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeatureCvterm) Reload(exec boil.Executor) error {
	ret, err := FindFeatureCvterm(exec, o.FeatureCvtermID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureCvtermSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureCvtermSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureCvtermSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty FeatureCvtermSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureCvtermSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureCvterms := FeatureCvtermSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_cvterm\".* FROM \"feature_cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureCvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featureCvtermPrimaryKeyColumns), 1, len(featureCvtermPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureCvterms)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in FeatureCvtermSlice")
	}

	*o = featureCvterms

	return nil
}

// FeatureCvtermExists checks if the FeatureCvterm row exists.
func FeatureCvtermExists(exec boil.Executor, featureCvtermID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_cvterm\" where \"feature_cvterm_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureCvtermID)
	}

	row := exec.QueryRow(sql, featureCvtermID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if feature_cvterm exists")
	}

	return exists, nil
}

// FeatureCvtermExistsG checks if the FeatureCvterm row exists.
func FeatureCvtermExistsG(featureCvtermID int) (bool, error) {
	return FeatureCvtermExists(boil.GetDB(), featureCvtermID)
}

// FeatureCvtermExistsGP checks if the FeatureCvterm row exists. Panics on error.
func FeatureCvtermExistsGP(featureCvtermID int) bool {
	e, err := FeatureCvtermExists(boil.GetDB(), featureCvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureCvtermExistsP checks if the FeatureCvterm row exists. Panics on error.
func FeatureCvtermExistsP(exec boil.Executor, featureCvtermID int) bool {
	e, err := FeatureCvtermExists(exec, featureCvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
