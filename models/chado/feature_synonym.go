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

// FeatureSynonym is an object representing the database table.
type FeatureSynonym struct {
	FeatureSynonymID int  `boil:"feature_synonym_id" json:"feature_synonym_id" toml:"feature_synonym_id" yaml:"feature_synonym_id"`
	SynonymID        int  `boil:"synonym_id" json:"synonym_id" toml:"synonym_id" yaml:"synonym_id"`
	FeatureID        int  `boil:"feature_id" json:"feature_id" toml:"feature_id" yaml:"feature_id"`
	PubID            int  `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`
	IsCurrent        bool `boil:"is_current" json:"is_current" toml:"is_current" yaml:"is_current"`
	IsInternal       bool `boil:"is_internal" json:"is_internal" toml:"is_internal" yaml:"is_internal"`

	R *featureSynonymR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureSynonymL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureSynonymR is where relationships are stored.
type featureSynonymR struct {
	Pub     *Pub
	Synonym *Synonym
	Feature *Feature
}

// featureSynonymL is where Load methods for each relationship are stored.
type featureSynonymL struct{}

var (
	featureSynonymColumns               = []string{"feature_synonym_id", "synonym_id", "feature_id", "pub_id", "is_current", "is_internal"}
	featureSynonymColumnsWithoutDefault = []string{"synonym_id", "feature_id", "pub_id"}
	featureSynonymColumnsWithDefault    = []string{"feature_synonym_id", "is_current", "is_internal"}
	featureSynonymPrimaryKeyColumns     = []string{"feature_synonym_id"}
)

type (
	// FeatureSynonymSlice is an alias for a slice of pointers to FeatureSynonym.
	// This should generally be used opposed to []FeatureSynonym.
	FeatureSynonymSlice []*FeatureSynonym
	// FeatureSynonymHook is the signature for custom FeatureSynonym hook methods
	FeatureSynonymHook func(boil.Executor, *FeatureSynonym) error

	featureSynonymQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureSynonymType                 = reflect.TypeOf(&FeatureSynonym{})
	featureSynonymMapping              = queries.MakeStructMapping(featureSynonymType)
	featureSynonymPrimaryKeyMapping, _ = queries.BindMapping(featureSynonymType, featureSynonymMapping, featureSynonymPrimaryKeyColumns)
	featureSynonymInsertCacheMut       sync.RWMutex
	featureSynonymInsertCache          = make(map[string]insertCache)
	featureSynonymUpdateCacheMut       sync.RWMutex
	featureSynonymUpdateCache          = make(map[string]updateCache)
	featureSynonymUpsertCacheMut       sync.RWMutex
	featureSynonymUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureSynonymBeforeInsertHooks []FeatureSynonymHook
var featureSynonymBeforeUpdateHooks []FeatureSynonymHook
var featureSynonymBeforeDeleteHooks []FeatureSynonymHook
var featureSynonymBeforeUpsertHooks []FeatureSynonymHook

var featureSynonymAfterInsertHooks []FeatureSynonymHook
var featureSynonymAfterSelectHooks []FeatureSynonymHook
var featureSynonymAfterUpdateHooks []FeatureSynonymHook
var featureSynonymAfterDeleteHooks []FeatureSynonymHook
var featureSynonymAfterUpsertHooks []FeatureSynonymHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeatureSynonym) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureSynonymBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeatureSynonym) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureSynonymBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeatureSynonym) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureSynonymBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeatureSynonym) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureSynonymBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeatureSynonym) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureSynonymAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeatureSynonym) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureSynonymAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeatureSynonym) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureSynonymAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeatureSynonym) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureSynonymAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeatureSynonym) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureSynonymAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureSynonymHook registers your hook function for all future operations.
func AddFeatureSynonymHook(hookPoint boil.HookPoint, featureSynonymHook FeatureSynonymHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureSynonymBeforeInsertHooks = append(featureSynonymBeforeInsertHooks, featureSynonymHook)
	case boil.BeforeUpdateHook:
		featureSynonymBeforeUpdateHooks = append(featureSynonymBeforeUpdateHooks, featureSynonymHook)
	case boil.BeforeDeleteHook:
		featureSynonymBeforeDeleteHooks = append(featureSynonymBeforeDeleteHooks, featureSynonymHook)
	case boil.BeforeUpsertHook:
		featureSynonymBeforeUpsertHooks = append(featureSynonymBeforeUpsertHooks, featureSynonymHook)
	case boil.AfterInsertHook:
		featureSynonymAfterInsertHooks = append(featureSynonymAfterInsertHooks, featureSynonymHook)
	case boil.AfterSelectHook:
		featureSynonymAfterSelectHooks = append(featureSynonymAfterSelectHooks, featureSynonymHook)
	case boil.AfterUpdateHook:
		featureSynonymAfterUpdateHooks = append(featureSynonymAfterUpdateHooks, featureSynonymHook)
	case boil.AfterDeleteHook:
		featureSynonymAfterDeleteHooks = append(featureSynonymAfterDeleteHooks, featureSynonymHook)
	case boil.AfterUpsertHook:
		featureSynonymAfterUpsertHooks = append(featureSynonymAfterUpsertHooks, featureSynonymHook)
	}
}

// OneP returns a single featureSynonym record from the query, and panics on error.
func (q featureSynonymQuery) OneP() *FeatureSynonym {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureSynonym record from the query.
func (q featureSynonymQuery) One() (*FeatureSynonym, error) {
	o := &FeatureSynonym{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for feature_synonym")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeatureSynonym records from the query, and panics on error.
func (q featureSynonymQuery) AllP() FeatureSynonymSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeatureSynonym records from the query.
func (q featureSynonymQuery) All() (FeatureSynonymSlice, error) {
	var o FeatureSynonymSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to FeatureSynonym slice")
	}

	if len(featureSynonymAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeatureSynonym records in the query, and panics on error.
func (q featureSynonymQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeatureSynonym records in the query.
func (q featureSynonymQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count feature_synonym rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureSynonymQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureSynonymQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if feature_synonym exists")
	}

	return count > 0, nil
}

// PubG pointed to by the foreign key.
func (o *FeatureSynonym) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *FeatureSynonym) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// SynonymG pointed to by the foreign key.
func (o *FeatureSynonym) SynonymG(mods ...qm.QueryMod) synonymQuery {
	return o.Synonym(boil.GetDB(), mods...)
}

// Synonym pointed to by the foreign key.
func (o *FeatureSynonym) Synonym(exec boil.Executor, mods ...qm.QueryMod) synonymQuery {
	queryMods := []qm.QueryMod{
		qm.Where("synonym_id=$1", o.SynonymID),
	}

	queryMods = append(queryMods, mods...)

	query := Synonyms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"synonym\"")

	return query
}

// FeatureG pointed to by the foreign key.
func (o *FeatureSynonym) FeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Feature(boil.GetDB(), mods...)
}

// Feature pointed to by the foreign key.
func (o *FeatureSynonym) Feature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureSynonymL) LoadPub(e boil.Executor, singular bool, maybeFeatureSynonym interface{}) error {
	var slice []*FeatureSynonym
	var object *FeatureSynonym

	count := 1
	if singular {
		object = maybeFeatureSynonym.(*FeatureSynonym)
	} else {
		slice = *maybeFeatureSynonym.(*FeatureSynonymSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureSynonymR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &featureSynonymR{}
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

	if len(featureSynonymAfterSelectHooks) != 0 {
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

// LoadSynonym allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureSynonymL) LoadSynonym(e boil.Executor, singular bool, maybeFeatureSynonym interface{}) error {
	var slice []*FeatureSynonym
	var object *FeatureSynonym

	count := 1
	if singular {
		object = maybeFeatureSynonym.(*FeatureSynonym)
	} else {
		slice = *maybeFeatureSynonym.(*FeatureSynonymSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureSynonymR{}
		args[0] = object.SynonymID
	} else {
		for i, obj := range slice {
			obj.R = &featureSynonymR{}
			args[i] = obj.SynonymID
		}
	}

	query := fmt.Sprintf(
		"select * from \"synonym\" where \"synonym_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Synonym")
	}
	defer results.Close()

	var resultSlice []*Synonym
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Synonym")
	}

	if len(featureSynonymAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Synonym = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.SynonymID == foreign.SynonymID {
				local.R.Synonym = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureSynonymL) LoadFeature(e boil.Executor, singular bool, maybeFeatureSynonym interface{}) error {
	var slice []*FeatureSynonym
	var object *FeatureSynonym

	count := 1
	if singular {
		object = maybeFeatureSynonym.(*FeatureSynonym)
	} else {
		slice = *maybeFeatureSynonym.(*FeatureSynonymSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureSynonymR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureSynonymR{}
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

	if len(featureSynonymAfterSelectHooks) != 0 {
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

// SetPub of the feature_synonym to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.FeatureSynonym.
func (o *FeatureSynonym) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_synonym\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureSynonymPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.FeatureSynonymID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &featureSynonymR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			FeatureSynonym: o,
		}
	} else {
		related.R.FeatureSynonym = o
	}

	return nil
}

// SetSynonym of the feature_synonym to the related item.
// Sets o.R.Synonym to related.
// Adds o to related.R.FeatureSynonym.
func (o *FeatureSynonym) SetSynonym(exec boil.Executor, insert bool, related *Synonym) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_synonym\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"synonym_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureSynonymPrimaryKeyColumns),
	)
	values := []interface{}{related.SynonymID, o.FeatureSynonymID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SynonymID = related.SynonymID

	if o.R == nil {
		o.R = &featureSynonymR{
			Synonym: related,
		}
	} else {
		o.R.Synonym = related
	}

	if related.R == nil {
		related.R = &synonymR{
			FeatureSynonym: o,
		}
	} else {
		related.R.FeatureSynonym = o
	}

	return nil
}

// SetFeature of the feature_synonym to the related item.
// Sets o.R.Feature to related.
// Adds o to related.R.FeatureSynonym.
func (o *FeatureSynonym) SetFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_synonym\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureSynonymPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeatureSynonymID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureID = related.FeatureID

	if o.R == nil {
		o.R = &featureSynonymR{
			Feature: related,
		}
	} else {
		o.R.Feature = related
	}

	if related.R == nil {
		related.R = &featureR{
			FeatureSynonym: o,
		}
	} else {
		related.R.FeatureSynonym = o
	}

	return nil
}

// FeatureSynonymsG retrieves all records.
func FeatureSynonymsG(mods ...qm.QueryMod) featureSynonymQuery {
	return FeatureSynonyms(boil.GetDB(), mods...)
}

// FeatureSynonyms retrieves all the records using an executor.
func FeatureSynonyms(exec boil.Executor, mods ...qm.QueryMod) featureSynonymQuery {
	mods = append(mods, qm.From("\"feature_synonym\""))
	return featureSynonymQuery{NewQuery(exec, mods...)}
}

// FindFeatureSynonymG retrieves a single record by ID.
func FindFeatureSynonymG(featureSynonymID int, selectCols ...string) (*FeatureSynonym, error) {
	return FindFeatureSynonym(boil.GetDB(), featureSynonymID, selectCols...)
}

// FindFeatureSynonymGP retrieves a single record by ID, and panics on error.
func FindFeatureSynonymGP(featureSynonymID int, selectCols ...string) *FeatureSynonym {
	retobj, err := FindFeatureSynonym(boil.GetDB(), featureSynonymID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureSynonym retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureSynonym(exec boil.Executor, featureSynonymID int, selectCols ...string) (*FeatureSynonym, error) {
	featureSynonymObj := &FeatureSynonym{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_synonym\" where \"feature_synonym_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureSynonymID)

	err := q.Bind(featureSynonymObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from feature_synonym")
	}

	return featureSynonymObj, nil
}

// FindFeatureSynonymP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureSynonymP(exec boil.Executor, featureSynonymID int, selectCols ...string) *FeatureSynonym {
	retobj, err := FindFeatureSynonym(exec, featureSynonymID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeatureSynonym) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeatureSynonym) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeatureSynonym) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeatureSynonym) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_synonym provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureSynonymColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureSynonymInsertCacheMut.RLock()
	cache, cached := featureSynonymInsertCache[key]
	featureSynonymInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureSynonymColumns,
			featureSynonymColumnsWithDefault,
			featureSynonymColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureSynonymType, featureSynonymMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureSynonymType, featureSynonymMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_synonym\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into feature_synonym")
	}

	if !cached {
		featureSynonymInsertCacheMut.Lock()
		featureSynonymInsertCache[key] = cache
		featureSynonymInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeatureSynonym record. See Update for
// whitelist behavior description.
func (o *FeatureSynonym) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeatureSynonym record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeatureSynonym) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeatureSynonym, and panics on error.
// See Update for whitelist behavior description.
func (o *FeatureSynonym) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeatureSynonym.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeatureSynonym) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureSynonymUpdateCacheMut.RLock()
	cache, cached := featureSynonymUpdateCache[key]
	featureSynonymUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureSynonymColumns, featureSynonymPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update feature_synonym, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_synonym\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featureSynonymPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureSynonymType, featureSynonymMapping, append(wl, featureSynonymPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update feature_synonym row")
	}

	if !cached {
		featureSynonymUpdateCacheMut.Lock()
		featureSynonymUpdateCache[key] = cache
		featureSynonymUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureSynonymQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureSynonymQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for feature_synonym")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureSynonymSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureSynonymSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureSynonymSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureSynonymSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureSynonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_synonym\" SET %s WHERE (\"feature_synonym_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureSynonymPrimaryKeyColumns), len(colNames)+1, len(featureSynonymPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in featureSynonym slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeatureSynonym) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeatureSynonym) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeatureSynonym) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeatureSynonym) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_synonym provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureSynonymColumnsWithDefault, o)

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

	featureSynonymUpsertCacheMut.RLock()
	cache, cached := featureSynonymUpsertCache[key]
	featureSynonymUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureSynonymColumns,
			featureSynonymColumnsWithDefault,
			featureSynonymColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureSynonymColumns,
			featureSynonymPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert feature_synonym, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featureSynonymPrimaryKeyColumns))
			copy(conflict, featureSynonymPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_synonym\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureSynonymType, featureSynonymMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureSynonymType, featureSynonymMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for feature_synonym")
	}

	if !cached {
		featureSynonymUpsertCacheMut.Lock()
		featureSynonymUpsertCache[key] = cache
		featureSynonymUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeatureSynonym record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureSynonym) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeatureSynonym record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeatureSynonym) DeleteG() error {
	if o == nil {
		return errors.New("chado: no FeatureSynonym provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeatureSynonym record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureSynonym) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeatureSynonym record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeatureSynonym) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeatureSynonym provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureSynonymPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_synonym\" WHERE \"feature_synonym_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from feature_synonym")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureSynonymQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureSynonymQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no featureSynonymQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from feature_synonym")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureSynonymSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureSynonymSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no FeatureSynonym slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureSynonymSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureSynonymSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeatureSynonym slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureSynonymBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureSynonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_synonym\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureSynonymPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureSynonymPrimaryKeyColumns), 1, len(featureSynonymPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featureSynonym slice")
	}

	if len(featureSynonymAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeatureSynonym) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeatureSynonym) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeatureSynonym) ReloadG() error {
	if o == nil {
		return errors.New("chado: no FeatureSynonym provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeatureSynonym) Reload(exec boil.Executor) error {
	ret, err := FindFeatureSynonym(exec, o.FeatureSynonymID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureSynonymSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureSynonymSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureSynonymSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty FeatureSynonymSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureSynonymSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureSynonyms := FeatureSynonymSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureSynonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_synonym\".* FROM \"feature_synonym\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureSynonymPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featureSynonymPrimaryKeyColumns), 1, len(featureSynonymPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureSynonyms)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in FeatureSynonymSlice")
	}

	*o = featureSynonyms

	return nil
}

// FeatureSynonymExists checks if the FeatureSynonym row exists.
func FeatureSynonymExists(exec boil.Executor, featureSynonymID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_synonym\" where \"feature_synonym_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureSynonymID)
	}

	row := exec.QueryRow(sql, featureSynonymID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if feature_synonym exists")
	}

	return exists, nil
}

// FeatureSynonymExistsG checks if the FeatureSynonym row exists.
func FeatureSynonymExistsG(featureSynonymID int) (bool, error) {
	return FeatureSynonymExists(boil.GetDB(), featureSynonymID)
}

// FeatureSynonymExistsGP checks if the FeatureSynonym row exists. Panics on error.
func FeatureSynonymExistsGP(featureSynonymID int) bool {
	e, err := FeatureSynonymExists(boil.GetDB(), featureSynonymID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureSynonymExistsP checks if the FeatureSynonym row exists. Panics on error.
func FeatureSynonymExistsP(exec boil.Executor, featureSynonymID int) bool {
	e, err := FeatureSynonymExists(exec, featureSynonymID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
