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

// FeaturePhenotype is an object representing the database table.
type FeaturePhenotype struct {
	FeaturePhenotypeID int `boil:"feature_phenotype_id" json:"feature_phenotype_id" toml:"feature_phenotype_id" yaml:"feature_phenotype_id"`
	FeatureID          int `boil:"feature_id" json:"feature_id" toml:"feature_id" yaml:"feature_id"`
	PhenotypeID        int `boil:"phenotype_id" json:"phenotype_id" toml:"phenotype_id" yaml:"phenotype_id"`

	R *featurePhenotypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featurePhenotypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featurePhenotypeR is where relationships are stored.
type featurePhenotypeR struct {
	Phenotype *Phenotype
	Feature   *Feature
}

// featurePhenotypeL is where Load methods for each relationship are stored.
type featurePhenotypeL struct{}

var (
	featurePhenotypeColumns               = []string{"feature_phenotype_id", "feature_id", "phenotype_id"}
	featurePhenotypeColumnsWithoutDefault = []string{"feature_id", "phenotype_id"}
	featurePhenotypeColumnsWithDefault    = []string{"feature_phenotype_id"}
	featurePhenotypePrimaryKeyColumns     = []string{"feature_phenotype_id"}
)

type (
	// FeaturePhenotypeSlice is an alias for a slice of pointers to FeaturePhenotype.
	// This should generally be used opposed to []FeaturePhenotype.
	FeaturePhenotypeSlice []*FeaturePhenotype
	// FeaturePhenotypeHook is the signature for custom FeaturePhenotype hook methods
	FeaturePhenotypeHook func(boil.Executor, *FeaturePhenotype) error

	featurePhenotypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featurePhenotypeType                 = reflect.TypeOf(&FeaturePhenotype{})
	featurePhenotypeMapping              = queries.MakeStructMapping(featurePhenotypeType)
	featurePhenotypePrimaryKeyMapping, _ = queries.BindMapping(featurePhenotypeType, featurePhenotypeMapping, featurePhenotypePrimaryKeyColumns)
	featurePhenotypeInsertCacheMut       sync.RWMutex
	featurePhenotypeInsertCache          = make(map[string]insertCache)
	featurePhenotypeUpdateCacheMut       sync.RWMutex
	featurePhenotypeUpdateCache          = make(map[string]updateCache)
	featurePhenotypeUpsertCacheMut       sync.RWMutex
	featurePhenotypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featurePhenotypeBeforeInsertHooks []FeaturePhenotypeHook
var featurePhenotypeBeforeUpdateHooks []FeaturePhenotypeHook
var featurePhenotypeBeforeDeleteHooks []FeaturePhenotypeHook
var featurePhenotypeBeforeUpsertHooks []FeaturePhenotypeHook

var featurePhenotypeAfterInsertHooks []FeaturePhenotypeHook
var featurePhenotypeAfterSelectHooks []FeaturePhenotypeHook
var featurePhenotypeAfterUpdateHooks []FeaturePhenotypeHook
var featurePhenotypeAfterDeleteHooks []FeaturePhenotypeHook
var featurePhenotypeAfterUpsertHooks []FeaturePhenotypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeaturePhenotype) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePhenotypeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeaturePhenotype) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePhenotypeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeaturePhenotype) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePhenotypeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeaturePhenotype) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePhenotypeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeaturePhenotype) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePhenotypeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeaturePhenotype) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePhenotypeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeaturePhenotype) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePhenotypeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeaturePhenotype) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePhenotypeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeaturePhenotype) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featurePhenotypeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeaturePhenotypeHook registers your hook function for all future operations.
func AddFeaturePhenotypeHook(hookPoint boil.HookPoint, featurePhenotypeHook FeaturePhenotypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featurePhenotypeBeforeInsertHooks = append(featurePhenotypeBeforeInsertHooks, featurePhenotypeHook)
	case boil.BeforeUpdateHook:
		featurePhenotypeBeforeUpdateHooks = append(featurePhenotypeBeforeUpdateHooks, featurePhenotypeHook)
	case boil.BeforeDeleteHook:
		featurePhenotypeBeforeDeleteHooks = append(featurePhenotypeBeforeDeleteHooks, featurePhenotypeHook)
	case boil.BeforeUpsertHook:
		featurePhenotypeBeforeUpsertHooks = append(featurePhenotypeBeforeUpsertHooks, featurePhenotypeHook)
	case boil.AfterInsertHook:
		featurePhenotypeAfterInsertHooks = append(featurePhenotypeAfterInsertHooks, featurePhenotypeHook)
	case boil.AfterSelectHook:
		featurePhenotypeAfterSelectHooks = append(featurePhenotypeAfterSelectHooks, featurePhenotypeHook)
	case boil.AfterUpdateHook:
		featurePhenotypeAfterUpdateHooks = append(featurePhenotypeAfterUpdateHooks, featurePhenotypeHook)
	case boil.AfterDeleteHook:
		featurePhenotypeAfterDeleteHooks = append(featurePhenotypeAfterDeleteHooks, featurePhenotypeHook)
	case boil.AfterUpsertHook:
		featurePhenotypeAfterUpsertHooks = append(featurePhenotypeAfterUpsertHooks, featurePhenotypeHook)
	}
}

// OneP returns a single featurePhenotype record from the query, and panics on error.
func (q featurePhenotypeQuery) OneP() *FeaturePhenotype {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featurePhenotype record from the query.
func (q featurePhenotypeQuery) One() (*FeaturePhenotype, error) {
	o := &FeaturePhenotype{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for feature_phenotype")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeaturePhenotype records from the query, and panics on error.
func (q featurePhenotypeQuery) AllP() FeaturePhenotypeSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeaturePhenotype records from the query.
func (q featurePhenotypeQuery) All() (FeaturePhenotypeSlice, error) {
	var o FeaturePhenotypeSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to FeaturePhenotype slice")
	}

	if len(featurePhenotypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeaturePhenotype records in the query, and panics on error.
func (q featurePhenotypeQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeaturePhenotype records in the query.
func (q featurePhenotypeQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count feature_phenotype rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featurePhenotypeQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featurePhenotypeQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if feature_phenotype exists")
	}

	return count > 0, nil
}

// PhenotypeG pointed to by the foreign key.
func (o *FeaturePhenotype) PhenotypeG(mods ...qm.QueryMod) phenotypeQuery {
	return o.Phenotype(boil.GetDB(), mods...)
}

// Phenotype pointed to by the foreign key.
func (o *FeaturePhenotype) Phenotype(exec boil.Executor, mods ...qm.QueryMod) phenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_id=$1", o.PhenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype\"")

	return query
}

// FeatureG pointed to by the foreign key.
func (o *FeaturePhenotype) FeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Feature(boil.GetDB(), mods...)
}

// Feature pointed to by the foreign key.
func (o *FeaturePhenotype) Feature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// LoadPhenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurePhenotypeL) LoadPhenotype(e boil.Executor, singular bool, maybeFeaturePhenotype interface{}) error {
	var slice []*FeaturePhenotype
	var object *FeaturePhenotype

	count := 1
	if singular {
		object = maybeFeaturePhenotype.(*FeaturePhenotype)
	} else {
		slice = *maybeFeaturePhenotype.(*FeaturePhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurePhenotypeR{}
		args[0] = object.PhenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &featurePhenotypeR{}
			args[i] = obj.PhenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype\" where \"phenotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phenotype")
	}
	defer results.Close()

	var resultSlice []*Phenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phenotype")
	}

	if len(featurePhenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phenotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PhenotypeID == foreign.PhenotypeID {
				local.R.Phenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featurePhenotypeL) LoadFeature(e boil.Executor, singular bool, maybeFeaturePhenotype interface{}) error {
	var slice []*FeaturePhenotype
	var object *FeaturePhenotype

	count := 1
	if singular {
		object = maybeFeaturePhenotype.(*FeaturePhenotype)
	} else {
		slice = *maybeFeaturePhenotype.(*FeaturePhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featurePhenotypeR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featurePhenotypeR{}
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

	if len(featurePhenotypeAfterSelectHooks) != 0 {
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

// SetPhenotype of the feature_phenotype to the related item.
// Sets o.R.Phenotype to related.
// Adds o to related.R.FeaturePhenotype.
func (o *FeaturePhenotype) SetPhenotype(exec boil.Executor, insert bool, related *Phenotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_phenotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurePhenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.PhenotypeID, o.FeaturePhenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PhenotypeID = related.PhenotypeID

	if o.R == nil {
		o.R = &featurePhenotypeR{
			Phenotype: related,
		}
	} else {
		o.R.Phenotype = related
	}

	if related.R == nil {
		related.R = &phenotypeR{
			FeaturePhenotype: o,
		}
	} else {
		related.R.FeaturePhenotype = o
	}

	return nil
}

// SetFeature of the feature_phenotype to the related item.
// Sets o.R.Feature to related.
// Adds o to related.R.FeaturePhenotype.
func (o *FeaturePhenotype) SetFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_phenotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurePhenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeaturePhenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureID = related.FeatureID

	if o.R == nil {
		o.R = &featurePhenotypeR{
			Feature: related,
		}
	} else {
		o.R.Feature = related
	}

	if related.R == nil {
		related.R = &featureR{
			FeaturePhenotype: o,
		}
	} else {
		related.R.FeaturePhenotype = o
	}

	return nil
}

// FeaturePhenotypesG retrieves all records.
func FeaturePhenotypesG(mods ...qm.QueryMod) featurePhenotypeQuery {
	return FeaturePhenotypes(boil.GetDB(), mods...)
}

// FeaturePhenotypes retrieves all the records using an executor.
func FeaturePhenotypes(exec boil.Executor, mods ...qm.QueryMod) featurePhenotypeQuery {
	mods = append(mods, qm.From("\"feature_phenotype\""))
	return featurePhenotypeQuery{NewQuery(exec, mods...)}
}

// FindFeaturePhenotypeG retrieves a single record by ID.
func FindFeaturePhenotypeG(featurePhenotypeID int, selectCols ...string) (*FeaturePhenotype, error) {
	return FindFeaturePhenotype(boil.GetDB(), featurePhenotypeID, selectCols...)
}

// FindFeaturePhenotypeGP retrieves a single record by ID, and panics on error.
func FindFeaturePhenotypeGP(featurePhenotypeID int, selectCols ...string) *FeaturePhenotype {
	retobj, err := FindFeaturePhenotype(boil.GetDB(), featurePhenotypeID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeaturePhenotype retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeaturePhenotype(exec boil.Executor, featurePhenotypeID int, selectCols ...string) (*FeaturePhenotype, error) {
	featurePhenotypeObj := &FeaturePhenotype{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_phenotype\" where \"feature_phenotype_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featurePhenotypeID)

	err := q.Bind(featurePhenotypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from feature_phenotype")
	}

	return featurePhenotypeObj, nil
}

// FindFeaturePhenotypeP retrieves a single record by ID with an executor, and panics on error.
func FindFeaturePhenotypeP(exec boil.Executor, featurePhenotypeID int, selectCols ...string) *FeaturePhenotype {
	retobj, err := FindFeaturePhenotype(exec, featurePhenotypeID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeaturePhenotype) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeaturePhenotype) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeaturePhenotype) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeaturePhenotype) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_phenotype provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurePhenotypeColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featurePhenotypeInsertCacheMut.RLock()
	cache, cached := featurePhenotypeInsertCache[key]
	featurePhenotypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featurePhenotypeColumns,
			featurePhenotypeColumnsWithDefault,
			featurePhenotypeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featurePhenotypeType, featurePhenotypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featurePhenotypeType, featurePhenotypeMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_phenotype\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into feature_phenotype")
	}

	if !cached {
		featurePhenotypeInsertCacheMut.Lock()
		featurePhenotypeInsertCache[key] = cache
		featurePhenotypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeaturePhenotype record. See Update for
// whitelist behavior description.
func (o *FeaturePhenotype) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeaturePhenotype record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeaturePhenotype) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeaturePhenotype, and panics on error.
// See Update for whitelist behavior description.
func (o *FeaturePhenotype) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeaturePhenotype.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeaturePhenotype) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featurePhenotypeUpdateCacheMut.RLock()
	cache, cached := featurePhenotypeUpdateCache[key]
	featurePhenotypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featurePhenotypeColumns, featurePhenotypePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update feature_phenotype, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_phenotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featurePhenotypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featurePhenotypeType, featurePhenotypeMapping, append(wl, featurePhenotypePrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update feature_phenotype row")
	}

	if !cached {
		featurePhenotypeUpdateCacheMut.Lock()
		featurePhenotypeUpdateCache[key] = cache
		featurePhenotypeUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featurePhenotypeQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featurePhenotypeQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for feature_phenotype")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeaturePhenotypeSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeaturePhenotypeSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeaturePhenotypeSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeaturePhenotypeSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePhenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_phenotype\" SET %s WHERE (\"feature_phenotype_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurePhenotypePrimaryKeyColumns), len(colNames)+1, len(featurePhenotypePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in featurePhenotype slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeaturePhenotype) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeaturePhenotype) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeaturePhenotype) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeaturePhenotype) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_phenotype provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featurePhenotypeColumnsWithDefault, o)

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

	featurePhenotypeUpsertCacheMut.RLock()
	cache, cached := featurePhenotypeUpsertCache[key]
	featurePhenotypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featurePhenotypeColumns,
			featurePhenotypeColumnsWithDefault,
			featurePhenotypeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featurePhenotypeColumns,
			featurePhenotypePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert feature_phenotype, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featurePhenotypePrimaryKeyColumns))
			copy(conflict, featurePhenotypePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_phenotype\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featurePhenotypeType, featurePhenotypeMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featurePhenotypeType, featurePhenotypeMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for feature_phenotype")
	}

	if !cached {
		featurePhenotypeUpsertCacheMut.Lock()
		featurePhenotypeUpsertCache[key] = cache
		featurePhenotypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeaturePhenotype record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeaturePhenotype) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeaturePhenotype record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeaturePhenotype) DeleteG() error {
	if o == nil {
		return errors.New("chado: no FeaturePhenotype provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeaturePhenotype record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeaturePhenotype) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeaturePhenotype record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeaturePhenotype) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeaturePhenotype provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featurePhenotypePrimaryKeyMapping)
	sql := "DELETE FROM \"feature_phenotype\" WHERE \"feature_phenotype_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from feature_phenotype")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featurePhenotypeQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featurePhenotypeQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no featurePhenotypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from feature_phenotype")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeaturePhenotypeSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeaturePhenotypeSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no FeaturePhenotype slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeaturePhenotypeSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeaturePhenotypeSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeaturePhenotype slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featurePhenotypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePhenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_phenotype\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurePhenotypePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurePhenotypePrimaryKeyColumns), 1, len(featurePhenotypePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featurePhenotype slice")
	}

	if len(featurePhenotypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeaturePhenotype) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeaturePhenotype) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeaturePhenotype) ReloadG() error {
	if o == nil {
		return errors.New("chado: no FeaturePhenotype provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeaturePhenotype) Reload(exec boil.Executor) error {
	ret, err := FindFeaturePhenotype(exec, o.FeaturePhenotypeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturePhenotypeSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeaturePhenotypeSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturePhenotypeSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty FeaturePhenotypeSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeaturePhenotypeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featurePhenotypes := FeaturePhenotypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePhenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_phenotype\".* FROM \"feature_phenotype\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurePhenotypePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featurePhenotypePrimaryKeyColumns), 1, len(featurePhenotypePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featurePhenotypes)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in FeaturePhenotypeSlice")
	}

	*o = featurePhenotypes

	return nil
}

// FeaturePhenotypeExists checks if the FeaturePhenotype row exists.
func FeaturePhenotypeExists(exec boil.Executor, featurePhenotypeID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_phenotype\" where \"feature_phenotype_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featurePhenotypeID)
	}

	row := exec.QueryRow(sql, featurePhenotypeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if feature_phenotype exists")
	}

	return exists, nil
}

// FeaturePhenotypeExistsG checks if the FeaturePhenotype row exists.
func FeaturePhenotypeExistsG(featurePhenotypeID int) (bool, error) {
	return FeaturePhenotypeExists(boil.GetDB(), featurePhenotypeID)
}

// FeaturePhenotypeExistsGP checks if the FeaturePhenotype row exists. Panics on error.
func FeaturePhenotypeExistsGP(featurePhenotypeID int) bool {
	e, err := FeaturePhenotypeExists(boil.GetDB(), featurePhenotypeID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeaturePhenotypeExistsP checks if the FeaturePhenotype row exists. Panics on error.
func FeaturePhenotypeExistsP(exec boil.Executor, featurePhenotypeID int) bool {
	e, err := FeaturePhenotypeExists(exec, featurePhenotypeID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
